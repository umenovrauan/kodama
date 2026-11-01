package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/umenovrauan/kodama/internal/telemetry"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	r := telemetry.NewReader(60)
	errc := make(chan error, 1)
	go func() { errc <- r.Run(ctx) }()

	fmt.Println("kodama: жду ACC… (Ctrl+C — выход)")

	var shown telemetry.Session
	for f := range r.Frames() {
		if s, ok := r.Session(); ok && s != shown {
			fmt.Printf("== %s / %s, maxRPM %d, SM %s\n", s.Track, s.Car, s.MaxRPM, s.SMVersion)
			shown = s
		}
		fmt.Printf("lap %2d pos %.3f %5.1f km/h gas %.2f brake %.2f steer %+.2f gear %2d %v\n", f.Lap, f.TrackPos, f.Speed*3.6, f.Throttle, f.Brake, f.Steer, f.Gear, f.LapTime.Round(10*time.Millisecond))
	}
	if err := <-errc; err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
	fmt.Printf("выход, выброшено кадров: %d\n", r.Dropped())
}
