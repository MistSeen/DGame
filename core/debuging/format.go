package debuging

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const supportedFlags = "0-+# "

var (
	panicBytes            = []byte("(PANIC=")
	plusBytes             = []byte("+")
	iBytes                = []byte("i")
	trueBytes             = []byte("true")
	falseBytes            = []byte("false")
	interfaceBytes        = []byte("(interface {})")
	commaNewlineBytes     = []byte(",\n")
	newlineBytes          = []byte("\n")
	openBraceBytes        = []byte("{")
	openBraceNewlineBytes = []byte("{\n")
	closeBraceBytes       = []byte("}")
	asteriskBytes         = []byte("*")
	colonBytes            = []byte(":")
	colonSpaceBytes       = []byte(": ")
	openParenBytes        = []byte("(")
	closeParenBytes       = []byte(")")
	spaceBytes            = []byte(" ")
	pointerChainBytes     = []byte("->")
	nilAngleBytes         = []byte("<nil>")
	maxNewlineBytes       = []byte("<max depth reached>\n")
	maxShortBytes         = []byte("<max>")
	circularBytes         = []byte("<already shown>")
	invalidAngleBytes     = []byte("<invalid>")
	openBracketBytes      = []byte("[")
	closeBracketBytes     = []byte("]")
	percentBytes          = []byte("%")
	precisionBytes        = []byte(".")
	openAngleBytes        = []byte("<")
	closeAngleBytes       = []byte(">")
	openMapBytes          = []byte("map[")
	closeMapBytes         = []byte("]")
	lenEqualsBytes        = []byte("len=")
	capEqualsBytes        = []byte("cap=")
	tooManyBytes          = []byte("......")
)

var (
	// uint8Type is a reflect.Type representing a uint8.  It is used to
	// convert cgo types to uint8 slices for hexformating.
	uint8Type = reflect.TypeOf(uint8(0))

	// cCharRE is a regular expression that matches a cgo char.
	// It is used to detect character arrays to hexformat them.
	regexCharRE = regexp.MustCompile("^.*\\._Ctype_char$")

	// cUnsignedCharRE is a regular expression that matches a cgo unsigned
	// char.  It is used to detect unsigned character arrays to hexformat
	// them.
	regexUnsignedCharRE = regexp.MustCompile("^.*\\._Ctype_unsignedchar$")

	// cUint8tCharRE is a regular expression that matches a cgo uint8_t.
	// It is used to detect uint8_t arrays to hexformat them.
	regexUint8tCharRE = regexp.MustCompile("^.*\\._Ctype_uint8_t$")

	// hexDigits is used to map a decimal value to a hex digit.
	hexDigits = "0123456789abcdef"
)

//========================================================//

//PrintWriter TODO: go
type FormatWriter interface {
	// Write is the function to call to emit formatted output to be printed.
	Write(b []byte) (n int, err error)
	WriteFormat(format string, a ...interface{}) (n int, err error)
	WriteBool(val bool)
	WriteInt(val int64, base int)
	WriteUint(val uint64, base int)
	WriteFloat(val float64, bitSize int)
	WriteComplex(val complex128, bitSize int)
	WriteHexPtr(ptr uintptr)

	// Width returns the value of the width option and whether it has been set.
	Width() (wid int, ok bool)
	// Precision returns the value of the precision option and whether it has been set.
	Precision() (prec int, ok bool)
	// Flag reports whether the flag c, a character, has been set.
	Flag(c int) bool
}

//WriteBool TODO: info
func WriteBool(p io.Writer, val bool) {
	if val {
		p.Write(trueBytes)
	} else {
		p.Write(falseBytes)
	}
}

//WriteInt write TODO: info
func WriteInt(p io.Writer, val int64, base int) {
	p.Write([]byte(strconv.FormatInt(val, base)))
}

//WriteUint TODO: info
func WriteUint(p io.Writer, val uint64, base int) {
	p.Write([]byte(strconv.FormatUint(val, base)))
}

//WriteFloat TODO: info
func WriteFloat(p io.Writer, val float64, bitSize int) {
	p.Write([]byte(strconv.FormatFloat(val, 'g', -1, bitSize)))
}

//WriteComplex TODO: info
func WriteComplex(p io.Writer, val complex128, bitSize int) {
	r := real(val)
	i := imag(val)
	p.Write([]byte(openParenBytes))
	p.Write([]byte(strconv.FormatFloat(r, 'g', -1, bitSize)))
	if i >= 0 {
		p.Write(plusBytes)
	}
	p.Write([]byte(strconv.FormatFloat(i, 'g', -1, bitSize)))
	p.Write(iBytes)
	p.Write(closeParenBytes)
}

//WriteHexPtr TODO: info
func WriteHexPtr(p io.Writer, ptr uintptr) {
	num := uint64(ptr)
	if num == 0 {
		p.Write(nilAngleBytes)
		return
	}

	buf := make([]byte, 18)

	base := uint64(16)
	i := len(buf) - 1
	for num >= base {
		buf[i] = hexDigits[num%base]
		num /= base
		i--
	}
	buf[i] = hexDigits[num]
	i--
	buf[i] = 'x'
	i--
	buf[i] = '0'

	buf = buf[i:]
	p.Write(buf)
}

//========================================================//
type ingoreNextState uint16

const (
	ingore_next_none   ingoreNextState = 0
	ingore_next_type   ingoreNextState = 1 << iota
	ingore_next_lencap ingoreNextState = 1 << iota

	ingore_next_all = ingore_next_type | ingore_next_lencap
)

//========================================================//
type ParamFormat struct {
	Config      *DebugConfig
	Value       interface{}
	Writer      FormatWriter
	pointers    map[uintptr]int
	depth       int
	ingoreState ingoreNextState
}

//NewParamFormat is
func NewParamFormat(v interface{}) *ParamFormat {
	return NewParamFormatWithConfig(&Config, v)
}

//NewParamFormatWithConfig is config
func NewParamFormatWithConfig(config *DebugConfig, v interface{}) (formater *ParamFormat) {
	formater = &ParamFormat{Value: v, Config: config}
	formater.pointers = make(map[uintptr]int)
	return formater
}

func (f *ParamFormat) FormatValue(v reflect.Value) {
	f.format(v)
}

func (f *ParamFormat) defaultFormater() (format string) {
	buf := bytes.NewBuffer(percentBytes)
	for _, flag := range supportedFlags {
		if f.Writer.Flag(int(flag)) {
			buf.WriteRune(flag)
		}
	}
	buf.WriteRune('v')
	format = buf.String()
	return format
}

func (f *ParamFormat) buildStdFormat(verb rune) (format string) {
	buf := bytes.NewBuffer(percentBytes)
	for _, flag := range supportedFlags {
		if f.Writer.Flag(int(flag)) {
			buf.WriteRune(flag)
		}
	}
	if width, ok := f.Writer.Width(); ok {
		buf.WriteString(strconv.Itoa(width))
	}
	if p, ok := f.Writer.Precision(); ok {
		buf.Write(precisionBytes)
		buf.WriteString(strconv.Itoa(p))
	}
	buf.WriteRune(verb)
	format = buf.String()
	return format
}
func (f *ParamFormat) isIngoreNone() bool {
	return f.ingoreState == ingore_next_none
}
func (f *ParamFormat) setIngoreNone() {
	f.ingoreState = ingore_next_none
}
func (f *ParamFormat) isNextIngore(state ingoreNextState) bool {
	return ((f.ingoreState & state) != 0)
}
func (f *ParamFormat) setNextIngoreState(state ingoreNextState, isIngore bool) {
	if isIngore {
		f.ingoreState = f.ingoreState | state
	} else {
		if f.isNextIngore(state) {
			f.ingoreState = f.ingoreState ^ state
		}
	}
}

func (f *ParamFormat) format(v reflect.Value) {
	kind := v.Kind()
	if kind == reflect.Invalid {
		f.Writer.Write(invalidAngleBytes)
		return
	}

	if kind == reflect.Ptr {
		f.formatPtr(v)
		return
	}
	isShowType := f.Writer.Flag('#')
	isArray := (kind == reflect.Array || kind == reflect.Slice)

	if !f.isNextIngore(ingore_next_type) && (isArray || isShowType) {
		f.Writer.Write(openParenBytes)
		f.Writer.Write([]byte(v.Type().String()))
		f.Writer.Write(closeParenBytes)
	}
	if !f.isNextIngore(ingore_next_lencap) && (isArray || isShowType) {
		valueLen, valueCap := 0, 0
		switch v.Kind() {
		case reflect.Array, reflect.Slice, reflect.Chan:
			valueLen, valueCap = v.Len(), v.Cap()
		case reflect.Map, reflect.String:
			valueLen = v.Len()
		}
		if valueLen != 0 || valueCap != 0 {
			f.Writer.Write(openParenBytes)
			if valueLen != 0 {
				f.Writer.Write(lenEqualsBytes)
				f.Writer.WriteInt(int64(valueLen), 10)
			}
			if valueCap != 0 {
				if valueLen != 0 {
					f.Writer.Write(spaceBytes)
				}
				f.Writer.Write(capEqualsBytes)
				f.Writer.WriteInt(int64(valueCap), 10)
			}
			f.Writer.Write(closeParenBytes)
			f.Writer.Write(spaceBytes)
		}
	}
	f.setIngoreNone()
	if !f.Config.DisableMethods {
		if kind != reflect.Invalid && kind != reflect.Interface {
			if handled := handleMethods(f.Config, f.Writer, v); handled {
				return
			}
		}
	}
	switch kind {
	case reflect.Invalid: /*Do nothing*/
	case reflect.Bool:
		f.Writer.WriteBool(v.Bool())
	case reflect.Float32:
		f.Writer.WriteFloat(v.Float(), 32)
	case reflect.Float64:
		f.Writer.WriteFloat(v.Float(), 64)
	case reflect.Complex64:
		f.Writer.WriteComplex(v.Complex(), 32)
	case reflect.Complex128:
		f.Writer.WriteComplex(v.Complex(), 64)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		f.Writer.WriteInt(v.Int(), 10)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		f.Writer.WriteUint(v.Uint(), 10)
	case reflect.String:
		f.Writer.Write([]byte(strconv.Quote(v.String())))
	case reflect.Interface:
		if v.IsNil() {
			f.Writer.Write(nilAngleBytes)
		}
	case reflect.Ptr: /*Do nothing*/
	case reflect.Uintptr:
		f.Writer.WriteHexPtr(uintptr(v.Uint()))
	case reflect.UnsafePointer, reflect.Chan, reflect.Func:
		f.Writer.WriteHexPtr(v.Pointer())
	case reflect.Slice:
		if v.IsNil() {
			f.Writer.Write(nilAngleBytes)
			break
		}
		fallthrough
	case reflect.Array:
		f.formatDepth(&v, openBracketBytes, closeBracketBytes,
			func(vv *reflect.Value) {
				f.formatArray(vv)
			})
	case reflect.Struct:
		f.formatDepth(&v, openBraceBytes, closeBraceBytes,
			func(vv *reflect.Value) {
				f.formatStruct(vv)
			})
	case reflect.Map:
		if v.IsNil() {
			f.Writer.Write(nilAngleBytes)
			break
		}
		openBytes := openBraceBytes
		closeBytes := closeBraceBytes
		if !isShowType {
			openBytes = openMapBytes
			closeBytes = closeMapBytes
		}
		f.formatDepth(&v, openBytes, closeBytes, func(vv *reflect.Value) {
			f.formatMap(vv)
		})
	default:
		format := f.defaultFormater()
		if v.CanInterface() {
			fmt.Fprint(f.Writer, format, v.Interface())
		} else {
			fmt.Fprint(f.Writer, format, v.String())
		}
	}
}

func (f *ParamFormat) formatPtr(v reflect.Value) {
	isShowTypes := f.Writer.Flag('#')
	if v.IsNil() && (!isShowTypes || !f.isIngoreNone()) {
		f.Writer.Write(nilAngleBytes)
		return
	}

	for k, depth := range f.pointers {
		if f.depth <= depth {
			delete(f.pointers, k)
		}
	}

	var pointerChain []uintptr
	var isNilFound = false
	var isCycleFound = false
	var indirects = 0
	var ve = v

	for ve.Kind() == reflect.Ptr {
		if ve.IsNil() {
			isNilFound = true
			break
		}
		indirects++
		addr := ve.Pointer()
		pointerChain = append(pointerChain, addr)

		if depth, ok := f.pointers[addr]; ok && depth < f.depth {
			isCycleFound = true
			indirects--
			break
		}

		f.pointers[addr] = f.depth

		ve = ve.Elem()

		if ve.Kind() == reflect.Interface {
			if ve.IsNil() {
				isNilFound = true
				break
			}
			ve = ve.Elem()
		}
	}
	if isShowTypes && f.isIngoreNone() {
		f.Writer.Write(openParenBytes)
		f.Writer.Write(bytes.Repeat(asteriskBytes, indirects))
		f.Writer.Write([]byte(ve.Type().String()))
		f.Writer.Write(closeParenBytes)
	} else {
		if isNilFound || isCycleFound {
			indirects += strings.Count(ve.Type().String(), "*")
		}
		f.Writer.Write(openAngleBytes)
		f.Writer.Write([]byte(strings.Repeat("*", indirects)))
		f.Writer.Write(closeAngleBytes)
	}
	if f.Writer.Flag('+') && (len(pointerChain) > 0) {
		f.Writer.Write(openParenBytes)
		for index, addr := range pointerChain {
			if index > 0 {
				f.Writer.Write(pointerChainBytes)
			}
			f.Writer.WriteHexPtr(addr)
		}
		f.Writer.Write(closeParenBytes)
	}
	switch {
	case isNilFound == true:
		f.Writer.Write(nilAngleBytes)
	case isCycleFound == true:
		f.Writer.Write(circularBytes)
	default:
		f.setNextIngoreState(ingore_next_type, true)
		f.format(ve)
	}

}

func (f *ParamFormat) formatArray(v *reflect.Value) {

	var buf []uint8
	doConvert := false
	doHexFormat := false
	numEntries := v.Len()
	isShowType := f.Writer.Flag('#')

	if numEntries > 0 {
		v_type := v.Index(0).Type()
		v_type_str := v_type.String()
		switch {
		case regexCharRE.MatchString(v_type_str):
			fallthrough
		case regexUnsignedCharRE.MatchString(v_type_str):
			fallthrough
		case regexUint8tCharRE.MatchString(v_type_str):
			doConvert = true
		case v_type.Kind() == reflect.Uint8:
			vs := *v
			if !vs.CanInterface() || !vs.CanAddr() {
				vs = unsafeReflectValue(vs)
			}
			if !IsDisabledUnsafeReflect {
				vs = vs.Slice(0, numEntries)
				iface := vs.Interface()

				if slice, ok := iface.([]uint8); ok {
					buf = slice
					doHexFormat = true
					break
				}
			}
			doConvert = true
		}

		if doConvert && v_type.ConvertibleTo(uint8Type) {
			buf = make([]uint8, numEntries)
			for i := 0; i < numEntries; i++ {
				vv := v.Index(i)
				buf[i] = uint8(vv.Convert(uint8Type).Uint())
			}
			doHexFormat = true
		}
	}

	if doHexFormat {
		indent := strings.Repeat(" ", f.depth)
		str := indent + hex.Dump(buf)
		str = strings.Replace(str, "\n", "\n"+indent, -1)
		str = strings.TrimRight(str, " ")
		f.Writer.Write(newlineBytes)
		f.Writer.Write([]byte(str))
		return
	}
	if isShowType {
		f.Writer.Write(newlineBytes)
	}
	for i := 0; i < numEntries; i++ {
		f.format(f.unpack(v.Index(i)))
		if isShowType {
			if i < numEntries {
				f.Writer.Write(commaNewlineBytes)
			} else {
				f.Writer.Write(newlineBytes)
			}
		} else {
			f.Writer.Write(spaceBytes)
		}
	}
}

func (f *ParamFormat) formatStruct(v *reflect.Value) {
	numFields := v.NumField()
	vt := v.Type()
	isShowType := f.Writer.Flag('+') || f.Writer.Flag('#')

	if numFields > 0 {
		f.Writer.Write(newlineBytes)
	}

	for index := 0; index < numFields; index++ {
		f.formatIndent()
		vtf := vt.Field(index)
		if isShowType {
			f.Writer.Write([]byte(vtf.Name))
			f.Writer.Write(colonBytes)
		}
		f.format(f.unpack(v.Field(index)))

		if index < numFields {
			f.Writer.Write(commaNewlineBytes)
		} else {
			f.Writer.Write(newlineBytes)
		}
	}
}

func (f *ParamFormat) formatMap(v *reflect.Value) {
	keys := v.MapKeys()
	if f.Config.SortKeys {
		SortValue(keys, f.Config)
	}
	isShowType := f.Writer.Flag('#')
	keyLen := len(keys)
	if isShowType {
		f.Writer.Write(newlineBytes)
	}
	for index, key := range keys {
		f.setNextIngoreState(ingore_next_all, true)
		f.format(f.unpack(key))

		f.Writer.Write(colonBytes)

		f.setNextIngoreState(ingore_next_all, true)
		f.format(f.unpack(v.MapIndex(key)))
		if isShowType {
			f.Writer.Write(commaNewlineBytes)
		} else {
			if index < keyLen-1 {
				f.Writer.Write(spaceBytes)
			}
		}
	}
}

type depthEach func(v *reflect.Value)

func (f *ParamFormat) formatDepth(v *reflect.Value, beginBytes []byte, endBytes []byte, eachFunc depthEach) {
	f.Writer.Write(beginBytes)
	f.depth++
	if (f.Config.MaxDepth != 0) && f.depth > f.Config.MaxDepth {
		f.Writer.Write(tooManyBytes)
		f.depth--
	} else {
		eachFunc(v)
		f.depth--
		f.formatIndent()
	}
	f.Writer.Write(endBytes)
}

func (f *ParamFormat) unpack(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface {
		f.setIngoreNone()
		if !v.IsNil() {
			v = v.Elem()
		}
	}
	return v
}
func (f *ParamFormat) formatIndent() {
	if f.depth > 0 {
		f.Writer.Write(bytes.Repeat(spaceBytes, f.depth))
	}
}

//=======================================================================//
func handleMethods(config *DebugConfig, writer io.Writer, v reflect.Value) (handled bool) {

	if !v.CanInterface() {
		if IsDisabledUnsafeReflect {
			return false
		}
		v = unsafeReflectValue(v)
	}

	if (!config.DisablePointerMethods) && (!IsDisabledUnsafeReflect) && (!v.CanAddr()) {
		v = unsafeReflectValue(v)
	}
	if v.CanAddr() {
		v = v.Addr()
	}
	switch iface := v.Interface().(type) {
	case error:
		defer onPanicError(writer, v)
		handleContinueMethods(config.ContinueOnMethod, writer, iface.Error())
		return !config.ContinueOnMethod
	case fmt.Stringer:
		defer onPanicError(writer, v)
		handleContinueMethods(config.ContinueOnMethod, writer, iface.String())
		return !config.ContinueOnMethod
	}
	return false
}

func handleContinueMethods(isContinue bool, writer io.Writer, errStr string) {
	if isContinue {
		writer.Write(openParenBytes)
		writer.Write([]byte(errStr))
		writer.Write(closeParenBytes)
		writer.Write(spaceBytes)
		return
	}
	writer.Write([]byte(errStr))
}
func onPanicError(writer io.Writer, v reflect.Value) {
	if err := recover(); err != nil {
		writer.Write(panicBytes)
		fmt.Fprintf(writer, "%v", err)
		writer.Write(closeParenBytes)
	}

}

//=======================================================================//
