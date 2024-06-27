package app

import (
	"context"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage"
)

type App struct { // TODO
}

type Logger interface {
	Error(msg string)
	Warn(msg string)
	Info(msg string)
	Debug(msg string)
}

type Storage interface {
	AddEvent(storage.Event) error
	UpdateEvent(storage.Event) error
	DeleteEvent(storage.Event) error
	ListEvents() (*[]storage.Event, error)
}

func New(logger Logger, storage Storage) *App {
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
