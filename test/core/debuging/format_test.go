package debuging

import (
	"bytes"
	"fmt"
	"reflect"
	. "server/core/debuging"
	"testing"
)

type testFmtState struct {
	bytes.Buffer
}

func (f *testFmtState) Flag(flag int) bool {
	if flag == int('+') {
		return true
	}
	return false
}
func (f *testFmtState) Precision() (int, bool) {
	return 0, false
}
func (f *testFmtState) Width() (int, bool) {
	return 0, false
}
func (d *testFmtState) Write(b []byte) (n int, err error) {
	return d.Buffer.Write(b)
}
func (d *testFmtState) WriteFormat(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(&d.Buffer, format, a)
}
func (d *testFmtState) WriteBool(val bool) {
	WriteBool(&d.Buffer, val)
}
func (d *testFmtState) WriteInt(val int64, base int) {
	WriteInt(&d.Buffer, val, base)
}
func (d *testFmtState) WriteUint(val uint64, base int) {
	WriteUint(&d.Buffer, val, base)
}
func (d *testFmtState) WriteFloat(val float64, bitSize int) {
	WriteFloat(&d.Buffer, val, bitSize)
}
func (d *testFmtState) WriteComplex(val complex128, bitSize int) {
	WriteComplex(&d.Buffer, val, bitSize)
}
func (d *testFmtState) WriteHexPtr(ptr uintptr) {
	WriteHexPtr(&d.Buffer, ptr)
}

func TestFormatInvalidReflectValue(t *testing.T) {
	v := new(reflect.Value)

	buf := new(testFmtState)
	f := &ParamFormat{Config: &Config, Writer: buf, Value: *v}
	f.FormatValue(*v)

	outStr := buf.String()
	expected := "<invalid>"

	if outStr != expected {
		t.Errorf("InvalidReflectValue \n expected: %s \n got: %s", expected, outStr)
	}
}
