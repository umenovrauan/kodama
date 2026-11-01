package lap

import (
	"context"
	"time"

	"github.com/umenovrauan/kodama/internal/telemetry"
)

type Replay struct {
	frames []telemetry.Frame
	speed  float64
	out    chan telemetry.Frame
}

func NewReplay(frames []telemetry.Frame, speed float64) *Replay {
	return &Replay{
		frames: frames,
		speed:  speed,
		out:    make(chan telemetry.Frame, 120),
	}
}

func (p *Replay) Frames() <-chan telemetry.Frame { return p.out }

func (p *Replay) Run(ctx context.Context) error {
	defer close(p.out)
	if len(p.frames) == 0 {
		return nil
	}

	start := time.Now()
	base := p.frames[0].At

	for _, f := range p.frames {
		at := time.Now()
		if p.speed > 0 {
			offset := time.Duration(float64(f.At.Sub(base)) / p.speed)
			at = start.Add(offset)
			if !sleepUntil(ctx, at) {
				return ctx.Err()
			}
		}

		f.At = at
		select {
		case p.out <- f:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

func sleepUntil(ctx context.Context, t time.Time) bool {
	d := time.Until(t)
	if d <= 0 {
		return ctx.Err() == nil
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}
