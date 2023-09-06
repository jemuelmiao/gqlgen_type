package gqlgen_type

import (
	"github.com/99designs/gqlgen/graphql"
)

type CustomLong int64

func MarshalCustomLong(i CustomLong) graphql.Marshaler {
	return graphql.MarshalInt64(int64(i))
}

func UnmarshalCustomLong(v interface{}) (CustomLong, error) {
	if value, ok := v.(int64); ok {
		return CustomLong(value), nil
	}
	return 0, fmt.Errorf("could not unmarshal CustomLong")
}
