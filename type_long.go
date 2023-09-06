package gqlgen_type

import (
	"fmt"
	"io"
)

type CustomLong int64

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (c *CustomLong) UnmarshalGQL(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("CustomLong must be int64")
	}
	*c = CustomLong(value)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (c CustomLong) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("%v", int64(c))))
}

