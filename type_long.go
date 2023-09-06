package gqlgen_type

//import (
//	"fmt"
//	"io"
//)
//
//type CustomLong int64
//
//// UnmarshalGQL implements the graphql.Unmarshaler interface
//func (c *CustomLong) UnmarshalGQL(v interface{}) error {
//	value, ok := v.(int64)
//	if !ok {
//		return fmt.Errorf("CustomLong must be int64")
//	}
//	*c = CustomLong(value)
//	return nil
//}
//
//// MarshalGQL implements the graphql.Marshaler interface
//func (c CustomLong) MarshalGQL(w io.Writer) {
//	w.Write([]byte(fmt.Sprintf("%v", int64(c))))
//}

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalMyCustomLongScalar(v int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(fmt.Sprintf("%v", v)))
	})
}

func UnmarshalMyCustomLongScalar(v interface{}) (int64, error) {
	value, ok := v.(int64)
	if !ok {
		return 0, fmt.Errorf("MyCustomLong must be int64")
	}
	return value, nil
}

