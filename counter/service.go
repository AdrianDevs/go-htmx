package count

import (
	"context"
	"errors"
	"log/slog"
)

type StoreInterface interface {
	GetCount(ctx context.Context, id string) (count int, err error)
}

type Counts struct {
	Global  int
	Session int
}

type IncrementType int

const (
	IncrementTypeUnknown IncrementType = iota
	IncrementTypeGlobal
	IncrementTypeSession
)

var ErrUnknownIncrementType = errors.New("unknown IncrementType")

type Service struct {
	Log   *slog.Logger
	Store StoreInterface
}

func NewService(log *slog.Logger, s StoreInterface) *Service {
	return &Service{
		Log:   log,
		Store: s,
	}
}

func (s Service) Increment(ctx context.Context, incrementType IncrementType, sessionID string) (counts Counts, err error) {
	println("Increment count")
	return counts, nil
}

func (s Service) Get(ctx context.Context, sessionID string) (counts Counts, err error) {
	println("Get count")
	counts.Session = 4
	return counts, nil
}
