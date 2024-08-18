package count

import "context"

type Store struct {
	db    string
	table string
}

func NewStore(dbName string, tableName string) (s *Store, err error) {
	s = &Store{db: dbName, table: tableName}
	return s, nil
}

func (s Store) GetCount(ctx context.Context, id string) (count int, err error) {
	return 1, nil
}
