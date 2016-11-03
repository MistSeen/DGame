//+build appengine disableunsafe

package debuging

import (
	"reflect"
)

const (
	//UnsafeReflectDisabled is disable
	UnsafeReflectDisabled = true
)

func unsafeReflectValue(v reflect.Value) reflect.Value {
	return v
}
