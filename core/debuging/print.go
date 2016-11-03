package debuging

import (
	"fmt"
	"io"
	"reflect"
)

type FmtPrintWriter struct {
	ParamFormat
	fs fmt.State
}

func NewFmtPrintWriter(config *DebugConfig, v interface{}) fmt.Formatter {
	p := &FmtPrintWriter{}
	p.Config = config
	p.pointers = make(map[uintptr]int)
	p.Value = v
	p.Writer = p
	return p
}

func (f *FmtPrintWriter) Format(fs fmt.State, verb rune) {

	f.fs = fs

	if verb != 'v' {
		format := f.buildStdFormat(verb)
		f.Writer.WriteFormat(format, f.Value)
		return
	}
	if f.Value == nil {
		if f.Writer.Flag('#') {
			f.Writer.Write(interfaceBytes)
		}
		f.Writer.Write(nilAngleBytes)
		return
	}

	f.format(reflect.ValueOf(f.Value))
}

//-----------------------------------------------------//

func FmtFprint(w io.Writer, args ...interface{}) (n int, err error) {

	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Fprint(w, formatters...)
}
func FmtFPrintf(w io.Writer, format string, args ...interface{}) (n int, err error) {

	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Fprintf(w, format, formatters...)
}
func FmtPrint(args ...interface{}) (n int, err error) {
	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Print(formatters...)
}

func FmtErrorf(format string, args ...interface{}) (err error) {
	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Errorf(format, formatters...)
}

func FmtFprintf(w io.Writer, format string, args ...interface{}) (n int, err error) {
	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Fprintf(w, format, formatters...)
}
func FmtFprintln(w io.Writer, args ...interface{}) (n int, err error) {
	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Fprintln(w, formatters...)
}

func FmtPrintf(format string, args ...interface{}) (n int, err error) {
	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Printf(format, formatters)
}

func FmtSprintln(args ...interface{}) string {
	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Sprintln(formatters...)
}
func FmtPrintln(args ...interface{}) (n int, err error) {
	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Println(formatters...)
}
func FmtSprintf(format string, args ...interface{}) string {
	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Sprintf(format, formatters...)
}
func FmtSprint(args ...interface{}) string {
	formatters := make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFmtPrintWriter(&Config, arg)
	}
	return fmt.Sprint(formatters...)
}

//-----------------------------------------------------//

func (f *FmtPrintWriter) Write(b []byte) (n int, err error) {
	return f.fs.Write(b)
}
func (f *FmtPrintWriter) WriteFormat(format string, a ...interface{}) (n int, err error) {

	return fmt.Fprintf(f.fs, format, a)
}
func (f *FmtPrintWriter) WriteBool(val bool) {
	WriteBool(f.fs, val)
}
func (f *FmtPrintWriter) WriteInt(val int64, base int) {
	WriteInt(f.fs, val, base)
}
func (f *FmtPrintWriter) WriteUint(val uint64, base int) {
	WriteUint(f.fs, val, base)
}
func (f *FmtPrintWriter) WriteFloat(val float64, bitSize int) {
	WriteFloat(f.fs, val, bitSize)
}
func (f *FmtPrintWriter) WriteComplex(val complex128, bitSize int) {
	WriteComplex(f.fs, val, bitSize)
}
func (f *FmtPrintWriter) WriteHexPtr(ptr uintptr) {
	WriteHexPtr(f.fs, ptr)
}

// Width returns the value of the width option and whether it has been set.
func (f *FmtPrintWriter) Width() (wid int, ok bool) {
	return f.fs.Width()
}

// Precision returns the value of the precision option and whether it has been set.
func (f *FmtPrintWriter) Precision() (prec int, ok bool) {
	return f.fs.Precision()
}

// Flag reports whether the flag c, a character, has been set.
func (f *FmtPrintWriter) Flag(c int) bool {
	return f.fs.Flag(c)
}
