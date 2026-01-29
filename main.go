package main

import (
	"errors"
	"fmt"
	"lesson5/document_store"
	"reflect"
)

type MyStract struct {
	Primary string
	Info    string
}

var ErrNilInput = errors.New("nil input")
var ErrNonStructType = errors.New("non struct type input")
var ErrDocumentNotFound = errors.New("document not found")
var ErrCollectionAlreadyExists = errors.New("collection already exists")
var ErrCollectionNotFound = errors.New("collection not found")
var ErrUnsupportedDocumentField = errors.New("unsupported document field")

var TypicalErrorBuilderMarshaling = errors.New("error in marshaling Document: %v")

func MarshalDocument(input any) (*document_store.Document, error) {

	if input == nil {
		return nil, fmt.Errorf(TypicalErrorBuilderMarshaling.Error(), ErrNilInput)
	}

	var typeOfInput reflect.Type = reflect.TypeOf(input)
	var valueOfInput reflect.Value = reflect.Indirect(reflect.ValueOf(input))

	var doc *document_store.Document

	if typeOfInput.Kind() != reflect.Struct {
		return nil, fmt.Errorf(TypicalErrorBuilderMarshaling.Error(), ErrNonStructType)
	}

	if fieldsCount := typeOfInput.NumField(); fieldsCount > 0 {
		doc := &document_store.Document{
			Fields: map[string]document_store.DocumentField{},
		}

		for i := range fieldsCount {
			fieldInput := typeOfInput.Field(i)

			fieldValue := valueOfInput.FieldByName(fieldInput.Name).Addr().Interface()

			switch fieldInput.Type.Kind() {
			case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
				doc.Fields[fieldInput.Name] = document_store.DocumentField{
					Type:  document_store.DocumentFieldTypeNumber,
					Value: fieldValue.(int64),
				}
			case reflect.String:
				doc.Fields[fieldInput.Name] = document_store.DocumentField{
					Type:  document_store.DocumentFieldTypeString,
					Value: fieldValue.(string),
				}
			case reflect.Bool:
				doc.Fields[fieldInput.Name] = document_store.DocumentField{
					Type:  document_store.DocumentFieldTypeBool,
					Value: fieldValue.(bool),
				}
			case reflect.Array, reflect.Slice:
				doc.Fields[fieldInput.Name] = document_store.DocumentField{
					Type:  document_store.DocumentFieldTypeArray,
					Value: fieldValue.([]any),
				}
			case reflect.Struct:
				doc.Fields[fieldInput.Name] = document_store.DocumentField{
					Type:  document_store.DocumentFieldTypeObject,
					Value: fieldValue,
				}
			default:
				return nil, fmt.Errorf(TypicalErrorBuilderMarshaling.Error(), ErrUnsupportedDocumentField)
			}
		}
	}

	return doc, nil
}

func main() {
	s := &MyStract{
		Primary: "Xsdsad",
		Info:    "Info",
	}

	v := *s

	doc, err := MarshalDocument(v)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(doc)
}
