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

//import (
//	"fmt"
//	"io"
//
//	"github.com/99designs/gqlgen/graphql"
//)
//
//func MarshalMyCustomLongScalar(v int64) graphql.Marshaler {
//	return graphql.WriterFunc(func(w io.Writer) {
//		w.Write([]byte(fmt.Sprintf("%v", v)))
//	})
//}
//
//func UnmarshalMyCustomLongScalar(v interface{}) (int64, error) {
//	value, ok := v.(int64)
//	if !ok {
//		return 0, fmt.Errorf("MyCustomLong must be int64")
//	}
//	return value, nil
//}


import (
	"fmt"
	"io"
	"strings"
)

type Banned bool

func (b Banned) MarshalGQL(w io.Writer) {
	if b {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func (b *Banned) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*b = strings.ToLower(v) == "true"
		return nil
	case bool:
		*b = Banned(v)
		return nil
	default:
		return fmt.Errorf("%T is not a bool", v)
	}
}

