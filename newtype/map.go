package newtype

import (
	"github.com/lyraproj/dgo/dgo"
	"github.com/lyraproj/dgo/internal"
)

// Map returns a type that represents an Map value
func Map(args ...interface{}) dgo.MapType {
	return internal.MapType(args...)
}

func StructEntry(key string, valueType dgo.Type, required bool) dgo.MapEntryType {
	return internal.StructEntry(key, valueType, required)
}

func Struct(entries ...dgo.MapEntryType) dgo.StructType {
	return internal.Struct(entries)
}