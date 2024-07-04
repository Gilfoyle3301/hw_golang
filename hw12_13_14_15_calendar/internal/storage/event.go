package storage

import "time"

type Event struct {
	ID          string
	Title       string
	EventAt     time.Time
	StartAt     time.Time
	EndAt       time.Time
	Description string
	NotifyAt    time.Time
	IsNotify    time.Time
}
