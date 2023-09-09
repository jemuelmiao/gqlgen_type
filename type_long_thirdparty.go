package gqlgen_type

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
)

func MarshalLongThirdParty(v int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(fmt.Sprintf("%v", v)))
	})
}

func UnmarshalLongThirdParty(v interface{}) (int64, error) {
	value, ok := v.(int64)
	if !ok {
		return 0, fmt.Errorf("LongThirdParty must be int64")
	}
	return value, nil
}
