//go:build !windows

package telemetry

func connect() (*conn, error) { return nil, errUnsupported }
