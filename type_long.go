package gqlgen_type

import (
	"fmt"
	"io"
)

type Long int64

func (l Long) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("%v", l)))
}

func (l *Long) UnmarshalGQL(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("Long must be int64")
	}
	*l = Long(value)
	return nil
}
