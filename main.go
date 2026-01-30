package main

import (
	"errors"
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

var ErrNilInput = errors.New("nil input")
var ErrNonStructType = errors.New("non struct type input")

var ErrUnsupportedDocumentField = errors.New("unsupported document field")

var TypicalErrorBuilderMarshaling = errors.New("error in marshaling Document: %v")

func main() {}
