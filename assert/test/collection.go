package assert

import (
	"reflect"
	"strings"
	"testing"
)

// Inside fails the test if the container (string, slice, or array) does not include the item.
func Inside(t *testing.T, item, container any, message ...any) {
	t.Helper()

	if inside(t, item, container) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection to contain '%v'", custom, item)
}

// NotInside fails the test if the container includes the item.
func NotInside(t *testing.T, item, container any, message ...any) {
	t.Helper()

	if !inside(t, item, container) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection NOT to contain '%v'", custom, item)
}

func inside(t *testing.T, item, container any) bool {
	t.Helper()

	val := reflect.ValueOf(container)

	switch val.Kind() {
	case reflect.String:
		strItem, ok := item.(string)
		if !ok {
			return false
		}
		return strings.Contains(val.String(), strItem)

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(val.Index(i).Interface(), item) {
				return true
			}
		}
	case reflect.Map:
		itemVal := reflect.ValueOf(item)
		if !itemVal.IsValid() {
			return false
		}

		if itemVal.Type().AssignableTo(val.Type().Key()) {
			return val.MapIndex(itemVal).IsValid()
		}
	default:
		t.Fatalf("Inside does not support type %T", container)
	}
	return false
}

// Deprecated: Use Inside instead.
//
// Contains fails the test if the container (string, slice, or array) does not include the item.
func Contains(t *testing.T, container any, item any, message ...any) {
	t.Helper()

	if inside(t, item, container) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection to contain '%v'", custom, item)
}

// Deprecated: Use NotInside instead.
//
// NotContains fails the test if the container includes the item.
func NotContains(t *testing.T, container any, item any, message ...any) {
	t.Helper()

	if !inside(t, item, container) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection NOT to contain '%v'", custom, item)
}
