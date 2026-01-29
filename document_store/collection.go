package document_store

type CollectionConfig struct {
	PrimaryKey string
}

type Collection struct {
	CollectionConfig
	Data map[string]*Document
}

func (s *Collection) Put(doc Document) {
	keyField, ok := doc.Fields[s.PrimaryKey]

	if !ok {
		return
	}

	if keyField.Type != DocumentFieldTypeString {
		return
	}

	if newDocKey, isString := keyField.Value.(string); isString {
		s.Data[newDocKey] = &doc
	}

}

func (s *Collection) Get(key string) (*Document, bool) {
	document, ok := s.Data[key]

	return document, ok
}

func (s *Collection) Delete(key string) bool {
	_, ok := s.Data[key]

	if !ok {
		return false
	}

	delete(s.Data, key)

	return true
}

func (s *Collection) List() []Document {
	listOfDocuments := make([]Document, 0, len(s.Data))

	for _, doc := range s.Data {
		listOfDocuments = append(listOfDocuments, *doc)
	}

	return listOfDocuments
}
