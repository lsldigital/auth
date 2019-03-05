package auth

import (
	"go/constant"
	"go/token"
	"reflect"

	"github.com/asdine/storm/q"
)

// NewTimeMatcher creates a Matcher for a given field1 and field2.
func NewTimeMatcher(field1, field2 string, tok token.Token) q.Matcher {
	return timeMatcherDelegate{Field1: field1, Field2: field2, Tok: tok}
}

type timeMatcherDelegate struct {
	Field1, Field2 string
	Tok            token.Token
}

func (r timeMatcherDelegate) Match(i interface{}) (bool, error) {
	v := reflect.Indirect(reflect.ValueOf(i))
	return r.MatchValue(&v)
}

func (r timeMatcherDelegate) MatchValue(v *reflect.Value) (bool, error) {
	field1 := v.FieldByName(r.Field1)
	if !field1.IsValid() {
		return false, q.ErrUnknownField
	}
	field2 := v.FieldByName(r.Field2)
	if !field2.IsValid() {
		return false, q.ErrUnknownField
	}
	return compare(field1.Interface(), field2.Interface(), r.Tok), nil
}

func compare(a, b interface{}, tok token.Token) bool {
	typea, typeb := reflect.TypeOf(a), reflect.TypeOf(b)

	if typea != nil && (typea.String() == "time.Time" || typea.String() == "*time.Time") &&
		typeb != nil && (typeb.String() == "time.Duration" || typeb.String() == "*time.Duration") {

		vala, valb := reflect.ValueOf(a), reflect.ValueOf(b)

		if typea.String() == "*time.Time" && vala.IsNil() {
			return true
		}

		if typeb.String() == "*time.Duration" {
			if valb.IsNil() {
				return true
			}
			valb = valb.Elem()
		}

		valc := vala.MethodByName("Add").Call([]reflect.Value{valb})[0]

		var x, y int64
		x = 1
		if vala.MethodByName("Equal").Call([]reflect.Value{valc})[0].Bool() {
			y = 1
		} else if vala.MethodByName("Before").Call([]reflect.Value{valc})[0].Bool() {
			y = 2
		}
		return constant.Compare(constant.MakeInt64(x), tok, constant.MakeInt64(y))
	}

	return false
}
