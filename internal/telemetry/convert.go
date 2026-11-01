package telemetry

import (
	"time"
	"unicode/utf16"
)


func toFrame(p *pagePhysics, g *pageGraphics, at time.Time) Frame {
	return Frame{
		At:	at,
		TrackPos: float64(g.NormalizedCarPosition),
		Speed:	float64(g.SpeedKmh) / 3.6,
		Throttle:	float64(p.Gas),
		Brake:	float64(p.Brake),
		Steer: float64(p.SteerAngle),
		Gear: int(p.Gear) - 1,
		Lap: int(p.CompletedLaps) + 1,
		LapTime: time.Duration(g.ICurrentTime) * time.Millisecond,
	}
}

func toSession(s *pageStatis) Session {
	return Session{
		Track: wstr(s.Track[:]),
		Car: wstr(s.CarModel[:]),
		MaxRPM: int(s.MaxRPM),
		SMVersion: wrst(s.SmVersion[:]),
	}
}

func wstr(b []uint16) string {
	for i, c := range b {
		if c == 0 {
			b = b[:i]
			break
		}
	}
	return string(utf16.Decode(b))
}