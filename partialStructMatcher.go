package gomockmatchers

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
)

const PartialStructMatcherDescription = "callback function returns true"

// Matches only on the fields specified by the input. Matching is performed using
func PartialStructMatcher(input interface{}) gomock.Matcher {
	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Struct {
		panic("argument to PartialStructMatcher must be a struct")
	}

	return partialStructMatcher{input}
}

type partialStructMatcher struct {
	// The function to run to test for matches. Provided by caller.
	input interface{}
}

func (m partialStructMatcher) Matches(x interface{}) bool {
	wantType := reflect.TypeOf(m.input)
	gotType := reflect.TypeOf(x)

	if !wantType.AssignableTo(gotType) {
		// Types must be compatible. We use AssignableTo() rather than
		// equality to account for embedded structs.
		return false
	}

	wantValue := reflect.ValueOf(m.input)
	gotValue := reflect.ValueOf(x)

	for i := 0; i < wantValue.NumField(); i++ {
		wantField := wantType.Field(i)
		gotFieldValue := gotValue.Field(i).Interface()

		fmt.Printf("%v: value: %v\n", wantField.Name, wantValue.Field(i).Interface())
		fmt.Printf("%v: type: %v\n", wantField.Name, wantField.Type)
		fmt.Printf("%v: zero? %v\n", wantField.Name, wantValue.Field(i).IsZero())
		if wantValue.Field(i).IsZero() {
			continue
		}

		fmt.Printf("%v: eqtest: %v and %v\n", wantField.Name, wantValue.Field(i).Interface(), gotFieldValue)
		if !reflect.DeepEqual(wantValue.Field(i).Interface(), gotFieldValue) {
			return false
		}
	}

	return true
}

func (o partialStructMatcher) String() string {
	return PartialStructMatcherDescription
}
