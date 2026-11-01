package telemetry

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

const (
	physicsName  = `Local\acpmf_physics`
	graphicsName = `Local\acpmf_graphics`
	staticName   = `Local\acpmf_static`
	retryDelay   = time.Second
	staleAfter   = 2 * time.Second
)

var (
	errUnsupported = errors.New("telemetry: ACC shared memory is only available on Windows")
	errStale       = errors.New("telemetry: data is stale")
)

type conn struct {
	physics  *pagePhysics
	graphics *pageGraphics
	static   *pageStatic
	closers  []func()
}

func (c *conn) release() {
	for i := len(c.closers) - 1; i >= 0; i-- {
		c.closers[i]()
	}
	c.closers = nil
}

type Reader struct {
	interval   time.Duration
	staleAfter time.Duration
	frames     chan Frame
	dropped    atomic.Uint64
	mu         sync.RWMutex
	session    Session
	haveSes    bool
}

func NewReader(rate int) *Reader {
	if rate <= 0 {
		rate = 60
	}
	return &Reader{
		interval:   time.Second / time.Duration(rate),
		staleAfter: staleAfter,
		frames:     make(chan Frame, rate*2),
	}
}

func (r *Reader) Frames() <-chan Frame { return r.frames }

func (r *Reader) Session() (Session, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.session, r.haveSes
}

func (r *Reader) Dropped() uint64 { return r.dropped.Load() }

func (r *Reader) setSession(s Session) {
	r.mu.Lock()
	r.session, r.haveSes = s, true
	r.mu.Unlock()
}

func (r *Reader) Run(ctx context.Context) error {
	defer close(r.frames)
	for {
		c, err := connect()
		if errors.Is(err, errUnsupported) {
			return err
		}
		if err == nil {
			_ = r.pump(ctx, c)
			c.release()
		}
		if !sleep(ctx, retryDelay) {
			return ctx.Err()
		}
	}
}

type pumpState struct {
	lastID     int32
	lastChange time.Time
	live       bool
}

func (r *Reader) pump(ctx context.Context, c *conn) error {
	t := time.NewTicker(r.interval)
	defer t.Stop()
	st := pumpState{lastID: -1, lastChange: time.Now()}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case now := <-t.C:
			if err := r.tick(c, &st, now); err != nil {
				return err
			}
		}
	}
}

func (r *Reader) tick(c *conn, st *pumpState, now time.Time) error {
	ph := *c.physics
	if c.physics.PacketID != ph.PacketID {
		return nil
	}

	gr := *c.graphics
	if gr.AccStatus != statusLive {
		st.live = false
		st.lastChange = now
		return nil
	}
	if !st.live {
		st.live = true
		r.setSession(toSession(c.static))
	}
	if ph.PacketID == st.lastID {
		if now.Sub(st.lastChange) > r.staleAfter {
			return errStale
		}
		return nil
	}
	st.lastID, st.lastChange = ph.PacketID, now

	r.emit(toFrame(&ph, &gr, now))
	return nil
}

func (r *Reader) emit(f Frame) {
	select {
	case r.frames <- f:
	default:
		r.dropped.Add(1)
	}
}

func sleep(ctx context.Context, d time.Duration) bool {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-t.C:
		return true
	}
}
