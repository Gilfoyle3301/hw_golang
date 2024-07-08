package memorystorage

import (
	"sync"

	"github.com/Gilfoyle3301/hw_golang/hw12_13_14_15_calendar/internal/logger"
	"github.com/Gilfoyle3301/hw_golang/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	mu     sync.RWMutex
	event  map[string]*storage.Event
	logger *logger.Logger
}

func New() *Storage {
	return &Storage{
		event:  make(map[string]*storage.Event),
		logger: logger.NewLogger(),
		mu:     sync.RWMutex{},
	}
}

func (s *Storage) AddEvent(event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.event[event.ID] != nil {
		s.logger.Warn("an event with this ID already exists")
	} else {
		s.logger.Info("event added succsseful")
	}
	s.event[event.ID] = &event
	return nil
}

func (s *Storage) UpdateEvent(event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.event[event.ID] == nil {
		s.logger.Error("event not found")
	} else {
		s.logger.Info("event update succsseful")
	}
	return nil
}

func (s *Storage) DeleteEvent(event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.event[event.ID] == nil {
		s.logger.Error("event not found")
	} else {
		s.logger.Info("event deleted succsseful")
	}
	delete(s.event, event.ID)
	return nil
}

func (s *Storage) ListEvents() ([]storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.event) != 0 {
		result := make([]storage.Event, 0, len(s.event))
		for _, v := range s.event {
			result = append(result, *v)
		}

		return result, nil
	}
	s.logger.Warn("no events were found")
	return []storage.Event{}, nil

}
