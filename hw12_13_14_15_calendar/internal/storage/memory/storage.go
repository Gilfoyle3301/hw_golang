package memorystorage

import (
	"sync"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage"
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
