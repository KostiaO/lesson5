package main

import (
	"fmt"
	"lesson5/document_store"
)

type Str struct {
	I int
}

type MyStract struct {
	Primary    string
	Info       string
	IsAssigned bool
	Summaries  []int
	Structa    Str
}

func main() {
	doc := &document_store.Document{
		Fields: map[string]document_store.DocumentField{
			"Primary": {
				Type:  document_store.DocumentFieldTypeString,
				Value: "DOCINDEX",
			},
			"Info": {
				Type:  document_store.DocumentFieldTypeString,
				Value: "Info",
			},
			"IsAssigned": {
				Type:  document_store.DocumentFieldTypeBool,
				Value: true,
			},
			"Summaries": {
				Type:  document_store.DocumentFieldTypeArray,
				Value: []int{2, 6, 7, 10},
			},
			"Structa": {
				Type: document_store.DocumentFieldTypeObject,
				Value: Str{
					I: 1,
				},
			},
		},
	}

	s := &MyStract{}

	document_store.UnmarshalDocument(doc, s)

	fmt.Println(s.Structa)
}
