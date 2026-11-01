package lap

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/umenovrauan/kodama/internal/telemetry"
	"io"
	"os"
	"slices"
	"strconv"
	"time"
)

var header = []string{
	"t_ms",
	"lap",
	"pos",
	"speed",
	"throttle",
	"brake",
	"steer",
	"gear",
	"lap_time_ms",
}

func Write(w io.Writer, frames []telemetry.Frame) error {
	cw := csv.NewWriter(w)
	if err := cw.Write(header); err != nil {
		return err
	}

	var start time.Time
	if len(frames) > 0 {
		start = frames[0].At
	}
	for _, f := range frames {
		row := []string{
			strconv.FormatInt(f.At.Sub(start).Milliseconds(), 10),
			strconv.Itoa(f.Lap),
			formatFloat(f.TrackPos, 6),
			formatFloat(f.Speed, 3),
			formatFloat(f.Throttle, 4),
			formatFloat(f.Brake, 4),
			formatFloat(f.Steer, 4),
			strconv.Itoa(f.Gear),
			strconv.FormatInt(f.LapTime.Milliseconds(), 10),
		}
		if err := cw.Write(row); err != nil {
			return err
		}
	}

	cw.Flush()
	return cw.Error()
}

func Read(r io.Reader) ([]telemetry.Frame, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = len(header)

	rows, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("lap: %w", err)
	}
	if len(rows) == 0 || !slices.Equal(rows[0], header) {
		return nil, errors.New("lap: not a lap file (bad header)")
	}

	frames := make([]telemetry.Frame, 0, len(rows)-1)
	for i, row := range rows[1:] {
		f, err := parseRow(row)
		if err != nil {
			return nil, fmt.Errorf("lap: line %d: %w", i+2, err)
		}
		frames = append(frames, f)
	}
	return frames, nil
}

func ReadFile(path string) ([]telemetry.Frame, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Read(f)
}

func parseRow(row []string) (telemetry.Frame, error) {
	var p parser
	f := telemetry.Frame{
		At:       time.Time{}.Add(time.Duration(p.int(row[0])) * time.Millisecond),
		Lap:      int(p.int(row[1])),
		TrackPos: p.float(row[2]),
		Speed:    p.float(row[3]),
		Throttle: p.float(row[4]),
		Brake:    p.float(row[5]),
		Steer:    p.float(row[6]),
		Gear:     int(p.int(row[7])),
		LapTime:  time.Duration(p.int(row[8])) * time.Millisecond,
	}
	return f, p.err
}

type parser struct{ err error }

func (p *parser) int(s string) int64 {
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 64)
	p.err = err
	return v
}

func (p *parser) float(s string) float64 {
	if p.err != nil {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	p.err = err
	return v
}

func formatFloat(v float64, digits int) string {
	return strconv.FormatFloat(v, 'f', digits, 64)
}
