package count

import (
	"context"
	"errors"
	"log/slog"
)

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

type CountStoreService struct {
	Log        *slog.Logger
	CountStore *CountStore
}

func NewService(log *slog.Logger, cs *CountStore) CountStoreService {
	return CountStoreService{
		Log:        log,
		CountStore: cs,
	}
}

func (cs CountStoreService) Increment(ctx context.Context, incrementType IncrementType, sessionID string) (counts Counts, err error) {
	println("Increment count")
	return counts, nil
}

func (cs CountStoreService) Get(ctx context.Context, sessionID string) (counts Counts, err error) {
	println("Get count")
	counts.Session = 4
	return counts, nil
}
