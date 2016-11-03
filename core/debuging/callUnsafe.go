//+build !appengine !disableunsafe

package debuging

import (
	"reflect"
	"unsafe"
)

//G
const (
	// UnsafeReflectDisabled is a build-time constant which specifies whether or
	// not access to the unsafe package is available.
	IsDisabledUnsafeReflect = false
	//ptrSize is the size of a pointer on the current arch.
	ptrSize = unsafe.Sizeof((*byte)(nil))
)

var (
	// offsetPtr, offsetScalar, and offsetFlag are the offsets for the
	// internal reflect.Value fields.  These values are valid before golang
	// commit ecccf07e7f9d which changed the format.  The are also valid
	// after commit 82f48826c6c7 which changed the format again to mirror
	// the original format.  Code in the init function updates these offsets
	// as necessary.
	offsetPtr    = uintptr(ptrSize)
	offsetScalar = uintptr(0)
	offsetFlag   = uintptr(ptrSize * 2)

	//flagKindWith and flagKindShift indicate various bits that the
	// reflect package user internally to track kind information
	//
	//flagRO indicates whether or not the value field of a reflect.Value is read-only
	//
	//flagIndir indicates whether the value field of a reflect.Value is the actual
	//data or a pointer to the data
	//
	//There values are valid before golong conmit 90a7c3c86944 which chagned their positions.
	//Code in the init function updates there flag as necessary
	flagKindWith  = uintptr(5)
	flagKindShift = uintptr(flagKindWith - 1)
	flagRO        = uintptr(1 << 0)
	flagIndir     = uintptr(1 << 1)
)

// unsafeReflectValue converts the passed reflect.Value into a one that callUnsaft
// the typical safety restrictions preventing access to unaddressable and
// unexported data.  It works by digging the raw pointer to the underlying
// value out of the protected value and generating a new unprotected (unsafe)
// reflect.Value to it.
//
// This allows us to check for implementations of the Stringer and error
// interfaces to be used for pretty printing ordinarily unaddressable and
// inaccessible values such as unexported struct fields.
func unsafeReflectValue(v reflect.Value) (rv reflect.Value) {
	offset := 1

	vt := v.Type()
	upv := unsafe.Pointer(uintptr(unsafe.Pointer(&v)) + offsetPtr)
	rvf := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&v)) + offsetFlag))

	if (rvf & flagIndir) != 0 {
		vt = reflect.PtrTo(v.Type())
		offset++
	} else if offsetScalar != 0 {
		switch vt.Kind() {
		case reflect.Uintptr, reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer: /*Do nonthing*/
		default:
			upv = unsafe.Pointer(uintptr(unsafe.Pointer(&v)) + offsetScalar)
		}
	}

	pv := reflect.NewAt(vt, upv)
	rv = pv
	for i := 0; i < offset; i++ {
		rv = rv.Elem()
	}
	return rv
}

func init() {
	// This code constructs a new reflect.Value from a known small integer
	// and checks if the size of the reflect.Value struct indicates it has
	// the scalar field. When it does, the offsets are updated accordingly.
	vv := reflect.ValueOf(0xf00)
	if unsafe.Sizeof(vv) == ptrSize*4 {
		offsetScalar = ptrSize * 3
		offsetFlag = ptrSize * 3
	}
	// Commit 90a7c3c86944 changed the flag positions such that the low
	// order bits are the kind.  This code extracts the kind from the flags
	// field and ensures it's the correct type.  When it's not, the flag
	// order has been changed to the newer format, so the flags are updated
	// accordingly.
	upf := unsafe.Pointer(uintptr(unsafe.Pointer(&vv)) + offsetFlag)
	upfv := *(*uintptr)(upf)
	flagKindMark := uintptr((1<<flagKindWith - 1) << flagKindShift)

	if (upfv&flagKindMark)>>flagKindShift != uintptr(reflect.Int) {
		flagKindShift = 0
		flagRO = 1 << 5
		flagIndir = 1 << 6
		// Commit adf9b30e5594 modified the flags to separate the
		// flagRO flag into two bits which specifies whether or not the
		// field is embedded.  This causes flagIndir to move over a bit
		// and means that flagRO is the combination of either of the
		// original flagRO bit and the new bit.
		//
		// This code detects the change by extracting what used to be
		// the indirect bit to ensure it's set.  When it's not, the flag
		// order has been changed to the newer format, so the flags are
		// updated accordingly.
		if (upfv & flagIndir) == 0 {
			flagRO = 3 << 5
			flagIndir = 1 << 7
		}
	}
}
