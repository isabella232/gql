// Code generated by " ../../../../github.com/grailbio/base/gtl/generate.py --PREFIX=simpleStruct8 --package=gql --output=struct8.go -DNN=8 struct.go.tpl ". DO NOT EDIT.

package gql

// A template for a struct with a fixed number of fields.
// Template parameters:
//
// 8: number of fields, as an integer.
//
// Example:
// $GRAIL/go/src/github.com/base/grailbio/base/gtl/generate.py --prefix=struct5 --package=gql --output=struct5.go -D8=5 struct.go.tpl

import "github.com/grailbio/gql/symbol"

// Len implements Struct
func (s *simpleStruct8Impl) Len() int { return s.nFields }

// Field implements Struct
func (s *simpleStruct8Impl) Field(i int) StructField { return StructField{s.names[i], s.values[i]} }

// Value implements Struct
func (s *simpleStruct8Impl) Value(colName symbol.ID) (Value, bool) {
	for i := 0; i < s.nFields; i++ {
		if s.names[i] == colName {
			return s.values[i], true
		}
	}
	return Value{}, false
}

var _ Struct = &simpleStruct8Impl{}

type simpleStruct8Impl struct {
	StructImpl
	nFields int
	names   [8]symbol.ID // symbol.ID is 32bit, so pack it from values.
	values  [8]Value
}
