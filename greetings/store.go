package greet

import "context"

type Store struct {
	db    string
	table string
}

func NewStore(dbName string, tableName string) (s *Store, err error) {
	s = &Store{db: dbName, table: tableName}
	return s, nil
}

func (s Store) GetGreetings(ctx context.Context) (greetings []string, err error) {
	greetings = []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return greetings, nil
}
