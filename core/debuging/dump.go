package debuging

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

type DumpWriter struct {
	Writer io.Writer
}

func DumpWithConfigOfWriter(config *DebugConfig, ioWriter io.Writer, a ...interface{}) {
	printWriter := &DumpWriter{Writer: ioWriter}
	for _, arg := range a {
		if arg == nil {
			printWriter.Write(interfaceBytes)
			printWriter.Write(nilAngleBytes)
			continue
		}
		d := ParamFormat{Config: config, Writer: printWriter}
		d.pointers = make(map[uintptr]int)
		d.format(reflect.ValueOf(arg))
	}
}
func DumapWithWriter(ioWriter io.Writer, a ...interface{}) {
	DumpWithConfigOfWriter(&Config, ioWriter, a...)
}
func DumpValue(a ...interface{}) string {
	var buf bytes.Buffer
	DumpWithConfigOfWriter(&Config, &buf, a...)
	return buf.String()
}
func DumpPrint(a ...interface{}) {
	DumpWithConfigOfWriter(&Config, os.Stdout, a...)
}

//=============================================//
func (d *DumpWriter) Write(b []byte) (n int, err error) {
	return d.Writer.Write(b)
}
func (d *DumpWriter) WriteFormat(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(d.Writer, format, a)
}
func (d *DumpWriter) WriteBool(val bool) {
	WriteBool(d.Writer, val)
}
func (d *DumpWriter) WriteInt(val int64, base int) {
	WriteInt(d.Writer, val, base)
}
func (d *DumpWriter) WriteUint(val uint64, base int) {
	WriteUint(d.Writer, val, base)
}
func (d *DumpWriter) WriteFloat(val float64, bitSize int) {
	WriteFloat(d.Writer, val, bitSize)
}
func (d *DumpWriter) WriteComplex(val complex128, bitSize int) {
	WriteComplex(d.Writer, val, bitSize)
}
func (d *DumpWriter) WriteHexPtr(ptr uintptr) {
	WriteHexPtr(d.Writer, ptr)
}

// Width returns the value of the width option and whether it has been set.
func (d *DumpWriter) Width() (wid int, ok bool) {
	return 0, true
}

// Precision returns the value of the precision option and whether it has been set.
func (d *DumpWriter) Precision() (prec int, ok bool) {
	return 0, true
}

// Flag reports whether the flag c, a character, has been set.
func (d *DumpWriter) Flag(c int) bool {
	return true
}
