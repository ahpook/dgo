package dgo

import (
	"encoding/json"

	"github.com/lyraproj/dgo/util"
	"gopkg.in/yaml.v3"
)

type (
	// MapEntry is a key-value association in a Map
	MapEntry interface {
		Value
		Freezable

		Key() Value

		Value() Value
	}

	// StructEntry describes a MapEntry
	StructEntry interface {
		MapEntry

		Required() bool
	}

	// EntryDoer performs some task on behalf of a caller
	EntryDoer func(entry MapEntry)

	// EntryMapper maps produces the value of an entry to a new value
	EntryMapper func(entry MapEntry) interface{}

	// EntryPredicate returns true of false based on the given entry
	EntryPredicate func(entry MapEntry) bool

	// Map represents an ordered set of key-value associations. The Map preserves the order by which the entries
	// were added. Associations retain their order even if their value change. When creating a Map from a go map
	// the associations will be sorted based on the natural order of the keys.
	Map interface {
		Value
		Freezable
		util.Indentable
		json.Marshaler
		json.Unmarshaler
		yaml.Marshaler
		yaml.Unmarshaler

		// All returns true if the predicate returns true for all entries of this Map.
		All(predicate EntryPredicate) bool

		// AllKeys returns true if the predicate returns true for all keys of this Map.
		AllKeys(predicate Predicate) bool

		// AllValues returns true if the predicate returns true for all values of this Map.
		AllValues(predicate Predicate) bool

		// Any returns true if the predicate returns true for any entry of this Map.
		Any(doer EntryPredicate) bool

		// AnyKey returns true if the predicate returns true for any key of this Map.
		AnyKey(doer Predicate) bool

		// AnyValue returns true if the predicate returns true for any value of this Map.
		AnyValue(doer Predicate) bool

		// Copy returns a copy of the Map. The copy is frozen or mutable depending on
		// the given argument. A request to create a frozen copy of an already frozen Map
		// is a no-op that returns the receiver
		//
		// If a frozen copy is requested from a non-frozen Map, then all non-frozen keys and
		// values will be copied and frozen recursively.
		//
		// A Copy of a map that contains back references to itself will result in a stack
		// overflow panic.
		Copy(frozen bool) Map

		// Each calls the given doer with each entry of this Map
		Each(doer EntryDoer)

		// EachKey calls the given doer with each key of this Map
		EachKey(doer Doer)

		// EachValue calls the given doer with each value of this Map
		EachValue(doer Doer)

		// Entries returns a frozen snapshot of the entries in this map.
		Entries() Array

		// Get returns the value for the given key. The method will return nil when the key is not present
		// in the map. Use NilValue to bind a key to nil
		Get(key interface{}) Value

		// Keys returns frozen snapshot of all the keys of this map
		Keys() Array

		// Len returns the number of associations in this map
		Len() int

		// Map returns a new map with the same keys where each value has been replaced using the
		// given mapper function.
		Map(mapper EntryMapper) Map

		// Merge returns a Map where all associations from this and the given Map are merged. The associations of the
		// given map have priority.
		Merge(associations Map) Map

		// Put adds an association between the given key and value. The old value for the key or nil is returned. The
		// method will panic if the map is immutable
		Put(key, value interface{}) Value

		// PutAll adds all associations from the given Map, overwriting any that has the same key. It will panic if the
		// map is immutable.
		PutAll(associations Map)

		// Remove returns a Map that is guaranteed to have no value associated with the given key. The previous value
		// associated with the key or nil is returned. The method will panic if the map is immutable.
		Remove(key interface{}) Value

		// RemoveAll returns a Map that is guaranteed to have no values associated with any of the given keys. It will
		// panic if the map is immutable.
		RemoveAll(keys Array)

		// SetType sets the type for this Map to the given argument which must be a MapType or a string that evaluates
		// to a MapType. The Map must be mutable and an instance of the given type
		SetType(t interface{})

		// Values returns snapshot of all the values of this map.
		Values() Array

		// With creates a copy of this Map containing an association between the given key and value.
		With(key, value interface{}) Map

		// Without returns a Map that is guaranteed to have no value associated with the given key.
		Without(key interface{}) Map

		// WithoutAll returns a Map that is guaranteed to have no values associated with any of the given keys.
		WithoutAll(keys Array) Map
	}

	// MapType is implemented by types representing implementations of the Map value
	MapType interface {
		SizedType

		// KeyType returns the type of the keys for instances of this type
		KeyType() Type

		// ValueType returns the type of the values for instances of this type
		ValueType() Type
	}

	// StructType represents maps with explicitly defined typed entries.
	StructType interface {
		MapType

		// Additional returns true if the maps that is described by this type are allowed to
		// have additional entries.
		Additional() bool

		// Each iterates over each entry of the StructType
		Each(doer func(StructEntry))

		// Get returns the MapEntry that is identified with the given key
		Get(key interface{}) MapEntry

		// Len returns the number of StructEntrys in this StructType
		Len() int

		// Validate checks that the given value represents a Map which is an instance of this struct and returns a
		// possibly empty slice of errors explaining why that's not the case. Errors are generated if a required key
		// is missing, not recognized, or if it is of incorrect type.
		//
		// The keyLabel argument is an optional function that produces a suitable label for a key. If it is nil,
		// then a default function that produces the string "parameter '<key>'" will be used. The function
		// is called when errors are produced.
		//
		// An empty slice indicates a successful validation
		Validate(keyLabel func(key Value) string, value interface{}) []error

		// ValidateVerbose checks that the given value represents a Map which is an instance of this struct and returns
		// a boolean result. During validation, both successful and failing errors are verbosely explained on the given
		// Indenter.
		ValidateVerbose(value interface{}, out util.Indenter) bool
	}
)
