package gqlgen_type

import (
	"fmt"
	"io"
)

type LongContext int64

func (l LongContext) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("%v", l)))
}

func (l *LongContext) UnmarshalGQL(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("LongContext must be int64")
	}
	*l = LongContext(value)
	return nil
}
