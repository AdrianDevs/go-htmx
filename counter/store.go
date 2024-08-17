package count

type CountStore struct {
	db    string
	table string
}

func NewStore(dbName string, tableName string) (cs *CountStore, err error) {
	cs = &CountStore{db: dbName, table: tableName}
	return cs, nil
}
