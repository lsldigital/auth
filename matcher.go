package auth

import (
	"reflect"
	"time"

	"github.com/asdine/storm/q"
)

// NewExpiredTimeMatcher creates a Matcher for a given field1 and field2.
func NewExpiredTimeMatcher(field1, field2 string) q.Matcher {
	return timeMatcherDelegate{Field1: field1, Field2: field2}
}

type timeMatcherDelegate struct {
	Field1, Field2 string
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
	return compare(field1.Interface(), field2.Interface()), nil
}

func compare(a, b interface{}) bool {
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

		timec, ok := vala.MethodByName("Add").Call([]reflect.Value{valb})[0].Interface().(time.Time)
		if !ok {
			return true
		}

		return time.Now().Truncate(time.Hour).After(timec)

	}

	return false
}
