package check

import (
	"reflect"
	"strings"
)

// Inside fails the test if the container (string, slice, or array) does not include the item.
func Inside(t T, report Report, item, container any, message ...any) {
	t.Helper()

	found, ok := inside(t, item, container)
	if !ok || found {
		return
	}

	custom := prefix(message...)
	report("%sExpected collection to contain '%v'", custom, item)
}

// NotInside fails the test if the container includes the item.
func NotInside(t T, report Report, item, container any, message ...any) {
	t.Helper()

	found, ok := inside(t, item, container)
	if !ok || !found {
		return
	}

	custom := prefix(message...)
	report("%sExpected collection NOT to contain '%v'", custom, item)
}

// inside reports whether container holds item; ok is false when the container
// type is not supported, after failing the test with t.Fatalf.
func inside(t T, item, container any) (found, ok bool) {
	t.Helper()

	val := reflect.ValueOf(container)

	switch val.Kind() {
	case reflect.String:
		strItem, isString := item.(string)
		if !isString {
			return false, true
		}
		return strings.Contains(val.String(), strItem), true
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(val.Index(i).Interface(), item) {
				return true, true
			}
		}
	case reflect.Map:
		itemVal := reflect.ValueOf(item)
		if !itemVal.IsValid() {
			return false, true
		}

		if itemVal.Type().AssignableTo(val.Type().Key()) {
			return val.MapIndex(itemVal).IsValid(), true
		}
	default:
		t.Fatalf("Inside does not support type %T", container)
		return false, false
	}

	return false, true
}
