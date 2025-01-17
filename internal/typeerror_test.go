package internal_test

import (
	"testing"

	"github.com/lyraproj/dgo/dgo"

	require "github.com/lyraproj/dgo/dgo_test"
	"github.com/lyraproj/dgo/newtype"
	"github.com/lyraproj/dgo/typ"
	"github.com/lyraproj/dgo/vf"
)

func TestIllegalAssignment(t *testing.T) {
	v := newtype.IllegalAssignment(typ.String, vf.Integer(3))

	require.Equal(t, v, newtype.IllegalAssignment(typ.String, vf.Integer(3)))
	require.NotEqual(t, v, newtype.IllegalAssignment(typ.String, vf.Integer(4)))
	require.NotEqual(t, v, `oops`)

	require.Instance(t, v.Type(), v)
	require.NotEqual(t, 0, v.HashCode())
	require.Equal(t, v.HashCode(), v.HashCode())
	require.Equal(t, `the value 3 cannot be assigned to a variable of type string`, v.String())
}

func TestIllegalSize(t *testing.T) {
	v := newtype.IllegalSize(newtype.String(1, 10), 12)

	require.Equal(t, v, newtype.IllegalSize(newtype.String(1, 10), 12))
	require.NotEqual(t, v, newtype.IllegalSize(newtype.String(1, 10), 11))
	require.NotEqual(t, v, `oops`)

	require.Instance(t, v.Type(), v)
	require.NotEqual(t, 0, v.HashCode())
	require.Equal(t, v.HashCode(), v.HashCode())
	require.Equal(t, `size constraint violation on type string[1,10] when attempting resize to 12`, v.String())
}

func TestIllegalMapKey(t *testing.T) {
	tp := newtype.Parse(`{a:string}`).(dgo.StructType)
	v := newtype.IllegalMapKey(tp, vf.String(`b`))

	require.Equal(t, v, newtype.IllegalMapKey(tp, vf.String(`b`)))
	require.NotEqual(t, v, newtype.IllegalMapKey(tp, vf.String(`c`)))
	require.NotEqual(t, v, `oops`)

	require.Instance(t, v.Type(), v)
	require.NotEqual(t, 0, v.HashCode())
	require.Equal(t, v.HashCode(), v.HashCode())
	require.Equal(t, `key "b" cannot added to type {"a":string}`, v.String())
}
