package lap

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/umenovrauan/kodama/internal/telemetry"
)

const minLapFrames = 600

type Recorder struct {
	dir     string
	session telemetry.Session
	lap     int
	clean   bool
	lastPos float64
	frames  []telemetry.Frame
}

func NewRecorder(dir string) *Recorder {
	return &Recorder{dir: dir}
}

func (r *Recorder) Add(f telemetry.Frame, s telemetry.Session) (string, error) {
	if s != r.session {
		r.session = s
		r.start(f, false)
		return "", nil
	}

	if f.Lap == r.lap {
		r.append(f)
		return "", nil
	}

	crossed := f.Lap == r.lap+1

	var path string
	var err error
	if crossed && r.clean && len(r.frames) >= minLapFrames {
		path, err = r.save()
	}
	r.start(f, crossed)
	return path, err
}

func (r *Recorder) start(f telemetry.Frame, clean bool) {
	r.lap = f.Lap
	r.clean = clean
	r.lastPos = f.TrackPos
	r.frames = append(r.frames[:0], f)
}

func (r *Recorder) append(f telemetry.Frame) {
	wrap := r.lastPos > 0.9 && f.TrackPos < 0.1
	if f.TrackPos < r.lastPos-0.01 && !wrap {
		r.clean = false
	}
	r.lastPos = f.TrackPos
	r.frames = append(r.frames, f)
}

func (r *Recorder) save() (string, error) {
	first, last := r.frames[0], r.frames[len(r.frames)-1]

	dir := filepath.Join(r.dir, safeName(r.session.Track), safeName(r.session.Car))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("lap: %w", err)
	}

	name := fmt.Sprintf("%s_lap%02d_%s.csv",
		first.At.Format("2006-01-02_150405"), r.lap, formatLapTime(last.LapTime))
	path := filepath.Join(dir, name)

	file, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("lap: %w", err)
	}
	if err := Write(file, r.frames); err != nil {
		file.Close()
		os.Remove(path)
		return "", fmt.Errorf("lap: write %s: %w", path, err)
	}
	if err := file.Close(); err != nil {
		os.Remove(path)
		return "", fmt.Errorf("lap: close %s: %w", path, err)
	}
	return path, nil
}

func formatLapTime(d time.Duration) string {
	m := d / time.Minute
	s := d % time.Minute / time.Second
	ms := d % time.Second / time.Millisecond
	return fmt.Sprintf("%d-%02d.%03d", m, s, ms)
}

func safeName(s string) string {
	s = strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '_', r == '-', r == '.':
			return r
		default:
			return '_'
		}
	}, s)
	if s == "" {
		return "unknown"
	}
	return s
}
