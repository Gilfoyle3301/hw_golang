package memorystorage

import (
	"testing"

	"github.com/Gilfoyle3301/hw_golang/hw12_13_14_15_calendar/internal/storage"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/suite"
)

type MemStoreTest struct {
	suite.Suite
}

func TestMemStorage(t *testing.T) {
	suite.Run(t, new(MemStoreTest))
}

func (s *MemStoreTest) TestAddAndListEvent() {
	memStorage := New()
	for i := 0; i < 300; i++ {
		newEvent := storage.Event{}
		err := faker.FakeData(&newEvent)
		s.Require().NoError(err)
		err = memStorage.AddEvent(newEvent)
		s.Require().NoError(err)
	}
	ev, _ := memStorage.ListEvents()
	s.Require().Equal(300, len(ev))
}

func (s *MemStoreTest) TestDeleteEvent() {
	newStorage := New()
	newEvent := storage.Event{}
	err := faker.FakeData(&newEvent)
	s.Require().NoError(err)
	err = newStorage.AddEvent(newEvent)
	s.Require().NoError(err)
	newStorage.DeleteEvent(newEvent)
	s.Require().NoError(err)
	s.Require().Equal(0, len(newStorage.event))

}
