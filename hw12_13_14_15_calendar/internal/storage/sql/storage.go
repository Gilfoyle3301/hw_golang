package sqlstorage

import (
	"context"
	"fmt"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	logger *logger.Logger
	DB     *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, connectSchema string) error {
	db, err := sqlx.ConnectContext(ctx, "postgres", connectSchema)
	if err != nil {
		s.logger.Error("failed connection database")
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	s.DB = db
	return db.Ping()
}

func (s *Storage) Close(ctx context.Context) error {
	return s.DB.Close()
}

func (s *Storage) AddEvent(ctx context.Context, event storage.Event) error {
	_, err := s.DB.NamedExecContext(ctx, `INSERT INTO events (
			id, 
			title, 
			description,
			event_at,
			start_at,
			end_at,
			notify_at,
			is_notify 
		) VALUE (
			:id, 
			:title, 
			:description,
			:event_at,
			:start_at,
			:end_at,
			:notify_at,
			:is_notify
		)`, event)
	return err
}

func (s *Storage) UpdateEvent(ctx context.Context, event storage.Event) error {
	_, err := s.DB.ExecContext(ctx, `UPDATE events SET 
		title = $2,
		description = $3,
		event_at = $4,
		start_at = $5,
		end_at = $6,
		notify_at = $7,
		is_notify = $8 WHERE id = $1`, event.ID, event.Title, event.Description, event.EventAt, event.StartAt, event.StartAt, event.StartAt, event.NotifyAt, event.IsNotify)
	return err
}

func (s *Storage) DeleteEvent(ctx context.Context, event storage.Event) error {
	_, err := s.DB.ExecContext(ctx, `DELETE FROM events WHERE id =  $1;`, event.ID)
	return err
}

func (s *Storage) ListEvents(ctx context.Context) (*[]storage.Event, error) {
	listEvents := []storage.Event{}
	query := `SELECT * FROM events`
	err := s.DB.SelectContext(ctx, &listEvents, query)
	if err != nil {
		return nil, err
	}
	return &listEvents, nil
}
