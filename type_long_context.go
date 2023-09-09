package gqlgen_type

import (
	"context"
	"fmt"
	"io"
)

type LongContext int64

func (l LongContext) MarshalGQLContext(ctx context.Context, w io.Writer) error {
	w.Write([]byte(fmt.Sprintf("%v", l)))
	return nil
}

func (l *LongContext) UnmarshalGQLContext(ctx context.Context, v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("LongContext must be int64")
	}
	*l = LongContext(value)
	return nil
}
