package document_store

type Store struct {
	Storage map[string]*Collection
}

func NewStore() *Store {
	return &Store{
		Storage: make(map[string]*Collection),
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	if _, ok := s.Storage[name]; ok {
		return false, nil
	}

	newCollection := &Collection{
		Data: make(map[string]*Document),
	}

	newCollection.PrimaryKey = cfg.PrimaryKey

	s.Storage[name] = newCollection

	return true, newCollection
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	if collection, ok := s.Storage[name]; ok {
		return collection, true
	}

	return nil, false
}

func (s *Store) DeleteCollection(name string) bool {
	if _, ok := s.Storage[name]; ok {
		delete(s.Storage, name)

		return true
	}

	return false
}
