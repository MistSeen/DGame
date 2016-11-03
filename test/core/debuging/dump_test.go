package debuging

import (
	"bytes"
	"fmt"
	"reflect"
	. "server/core/debuging"
	"strconv"
	"testing"
	"unsafe"
)

func TestDumpInvalidReflectValue(t *testing.T) {
	v := new(reflect.Value)

	outStr := DumpValue(*v)
	expected := "(reflect.Value)<invalid Value>"

	if outStr != expected {
		t.Errorf("InvalidReflectValue \n expected: [%s] \n got: [%s]", expected, outStr)
	}
}

type dumpData struct {
	in interface{}
	on []string
}

func (d *dumpData) onString() string {
	s := ""
	for i, on := range d.on {
		if i > 1 {
			s += fmt.Sprintf("#%d: %s", i+1, on)
		}
		s += on
	}
	return s
}

var dumpDatas = make([]dumpData, 0)

func addToDumpDatas(in interface{}, on ...string) {
	d := dumpData{in, on}
	dumpDatas = append(dumpDatas, d)
}

func TestDumpInt(t *testing.T) {

	//Max int8
	v := int8(127)
	v_nil := (*int8)(nil)
	v_ptr := &v

	v_str := fmt.Sprintf("%p", v_ptr)
	v_ptr_str := fmt.Sprintf("%p", &v_ptr)

	v_type := "int8"
	v_value := "127"

	addToDumpDatas(v, "("+v_type+")"+v_value)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_str+")"+v_value+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value)
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	//Max int16
	v2 := int16(32767)
	v2_nil := (*int16)(nil)
	v2_ptr := &v2

	v2_str := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_str := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "int16"
	v2_value := "32767"

	addToDumpDatas(v2, "("+v2_type+")"+v2_value)
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_str+")"+v2_value+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_str+"->"+v2_str+")"+v2_value)
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	//Max int32
	v3 := int32(2147483647)
	v3_nil := (*int32)(nil)
	v3_ptr := &v3

	v3_str := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_str := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "int32"
	v3_value := "2147483647"

	addToDumpDatas(v3, "("+v3_type+")"+v3_value)
	addToDumpDatas(v3_ptr, "(*"+v3_type+")("+v3_str+")"+v3_value+"")
	addToDumpDatas(&v3_ptr, "(**"+v3_type+")("+v3_ptr_str+"->"+v3_str+")"+v3_value)
	addToDumpDatas(v3_nil, "(*"+v3_type+")<nil>")

	//Max int64
	v4 := int64(9223372036854775807)
	v4_nil := (*int64)(nil)
	v4_ptr := &v4

	v4_str := fmt.Sprintf("%p", v4_ptr)
	v4_ptr_str := fmt.Sprintf("%p", &v4_ptr)

	v4_type := "int64"
	v4_value := "9223372036854775807"

	addToDumpDatas(v4, "("+v4_type+")"+v4_value)
	addToDumpDatas(v4_ptr, "(*"+v4_type+")("+v4_str+")"+v4_value+"")
	addToDumpDatas(&v4_ptr, "(**"+v4_type+")("+v4_ptr_str+"->"+v4_str+")"+v4_value)
	addToDumpDatas(v4_nil, "(*"+v4_type+")<nil>")

	//Max int
	v5 := int(2147483647)
	v5_nil := (*int)(nil)
	v5_ptr := &v5

	v5_str := fmt.Sprintf("%p", v5_ptr)
	v5_ptr_str := fmt.Sprintf("%p", &v5_ptr)

	v5_type := "int"
	v5_value := "2147483647"

	addToDumpDatas(v5, "("+v5_type+")"+v5_value)
	addToDumpDatas(v5_ptr, "(*"+v5_type+")("+v5_str+")"+v5_value+"")
	addToDumpDatas(&v5_ptr, "(**"+v5_type+")("+v5_ptr_str+"->"+v5_str+")"+v5_value)
	addToDumpDatas(v5_nil, "(*"+v5_type+")<nil>")

	doDumpTesting(t)
}

func TestDumpUInt(t *testing.T) {

	//Max uint8
	v := uint8(0xff)
	v_nil := (*uint8)(nil)
	v_ptr := &v

	v_str := fmt.Sprintf("%p", v_ptr)
	v_ptr_str := fmt.Sprintf("%p", &v_ptr)

	v_type := "uint8"
	v_value := strconv.FormatUint(uint64(v), 10)

	addToDumpDatas(v, "("+v_type+")"+v_value)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_str+")"+v_value+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value)
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	//Max uint16
	v2 := uint16(0xffff)
	v2_nil := (*uint16)(nil)
	v2_ptr := &v2

	v2_str := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_str := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "uint16"
	v2_value := strconv.FormatUint(uint64(v2), 10)

	addToDumpDatas(v2, "("+v2_type+")"+v2_value)
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_str+")"+v2_value+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_str+"->"+v2_str+")"+v2_value)
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	//Max uint32
	v3 := uint32(0xffffffff)
	v3_nil := (*uint32)(nil)
	v3_ptr := &v3

	v3_str := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_str := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "uint32"
	v3_value := strconv.FormatUint(uint64(v3), 10)

	addToDumpDatas(v3, "("+v3_type+")"+v3_value)
	addToDumpDatas(v3_ptr, "(*"+v3_type+")("+v3_str+")"+v3_value+"")
	addToDumpDatas(&v3_ptr, "(**"+v3_type+")("+v3_ptr_str+"->"+v3_str+")"+v3_value)
	addToDumpDatas(v3_nil, "(*"+v3_type+")<nil>")

	//Max uint64
	v4 := uint64(0xffffffffffffffff)
	v4_nil := (*uint64)(nil)
	v4_ptr := &v4

	v4_str := fmt.Sprintf("%p", v4_ptr)
	v4_ptr_str := fmt.Sprintf("%p", &v4_ptr)

	v4_type := "uint64"
	v4_value := strconv.FormatUint(uint64(v4), 10)

	addToDumpDatas(v4, "("+v4_type+")"+v4_value)
	addToDumpDatas(v4_ptr, "(*"+v4_type+")("+v4_str+")"+v4_value+"")
	addToDumpDatas(&v4_ptr, "(**"+v4_type+")("+v4_ptr_str+"->"+v4_str+")"+v4_value)
	addToDumpDatas(v4_nil, "(*"+v4_type+")<nil>")

	//Max int
	v5 := uint(0xffffffffffffffff)
	v5_nil := (*uint)(nil)
	v5_ptr := &v5

	v5_str := fmt.Sprintf("%p", v5_ptr)
	v5_ptr_str := fmt.Sprintf("%p", &v5_ptr)

	v5_type := "uint"
	v5_value := strconv.FormatUint(uint64(v5), 10)

	addToDumpDatas(v5, "("+v5_type+")"+v5_value)
	addToDumpDatas(v5_ptr, "(*"+v5_type+")("+v5_str+")"+v5_value+"")
	addToDumpDatas(&v5_ptr, "(**"+v5_type+")("+v5_ptr_str+"->"+v5_str+")"+v5_value)
	addToDumpDatas(v5_nil, "(*"+v5_type+")<nil>")

	doDumpTesting(t)
}

func TestDumpBool(t *testing.T) {

	//Boolean true
	v := bool(true)
	v_nil := (*bool)(nil)
	v_ptr := &v

	v_str := fmt.Sprintf("%p", v_ptr)
	v_ptr_str := fmt.Sprintf("%p", &v_ptr)

	v_type := "bool"
	v_value := strconv.FormatBool(v)

	addToDumpDatas(v, "("+v_type+")"+v_value)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_str+")"+v_value+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value)
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	//Boolean false
	v2 := bool(false)
	v2_nil := (*bool)(nil)
	v2_ptr := &v2

	v2_str := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_str := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "bool"
	v2_value := strconv.FormatBool(v2)

	addToDumpDatas(v2, "("+v2_type+")"+v2_value)
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_str+")"+v2_value+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_str+"->"+v2_str+")"+v2_value)
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	doDumpTesting(t)
}

func TestDumpFloat(t *testing.T) {

	//Float 3.14159
	v := float32(3.14159)
	v_nil := (*float32)(nil)
	v_ptr := &v

	v_str := fmt.Sprintf("%p", v_ptr)
	v_ptr_str := fmt.Sprintf("%p", &v_ptr)

	v_type := "float32"
	v_value := strconv.FormatFloat(float64(v), 'g', 6, 32)

	addToDumpDatas(v, "("+v_type+")"+v_value)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_str+")"+v_value+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value)
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	//float64 false
	v2 := float64(3.1415926)
	v2_nil := (*float64)(nil)
	v2_ptr := &v2

	v2_str := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_str := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "float64"
	v2_value := strconv.FormatFloat(float64(v2), 'g', 8, 64)

	addToDumpDatas(v2, "("+v2_type+")"+v2_value)
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_str+")"+v2_value+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_str+"->"+v2_str+")"+v2_value)
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	doDumpTesting(t)
}

func TestDumpComplex64(t *testing.T) {

	//Boolean true
	v := complex(float32(8), -99)
	v_nil := (*complex64)(nil)
	v_ptr := &v

	v_str := fmt.Sprintf("%p", v_ptr)
	v_ptr_str := fmt.Sprintf("%p", &v_ptr)

	v_type := "complex64"
	v_value := "(8-99i)"

	addToDumpDatas(v, "("+v_type+")"+v_value)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_str+")"+v_value+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value)
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	//Boolean false
	v2 := complex(float64(-28), 199)
	v2_nil := (*complex128)(nil)
	v2_ptr := &v2

	v2_str := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_str := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "complex128"
	v2_value := "(-28+199i)"

	addToDumpDatas(v2, "("+v2_type+")"+v2_value)
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_str+")"+v2_value+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_str+"->"+v2_str+")"+v2_value)
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	doDumpTesting(t)
}

func TestDumpString(t *testing.T) {

	//String true
	v := "debuging"
	v_nil := (*string)(nil)
	v_ptr := &v

	v_str := fmt.Sprintf("%p", v_ptr)
	v_ptr_str := fmt.Sprintf("%p", &v_ptr)

	v_type := "string"
	v_value := "(len=8) \"debuging\""

	addToDumpDatas(v, "("+v_type+")"+v_value)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_str+")"+v_value+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value)
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	doDumpTesting(t)
}

type dString string

func (s dString) String() string {
	return string(s)
}

type dPtrString string

func (p *dPtrString) String() string {
	return string(*p)
}

func TestDumpArray(t *testing.T) {

	//Boolean true
	v := [3]int{3, 1, 2}
	v_nil := (*[3]int)(nil)
	v_ptr := &v

	v_len := fmt.Sprintf("%d", len(v))
	v_cap := fmt.Sprintf("%d", cap(v))

	v_str := fmt.Sprintf("%p", v_ptr)
	v_ptr_str := fmt.Sprintf("%p", &v_ptr)

	v_type := "int"
	v_value := fmt.Sprintf("(len=%s cap=%s) [\n(%s)3,\n(%s)1,\n(%s)2,\n]",
		v_len, v_cap, v_type, v_type, v_type)

	addToDumpDatas(v, "([3]"+v_type+")"+v_value)
	addToDumpDatas(v_ptr, "(*[3]"+v_type+")("+v_str+")"+v_value+"")
	addToDumpDatas(&v_ptr, "(**[3]"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value)
	addToDumpDatas(v_nil, "(*[3]"+v_type+")<nil>")

	v2_tStr_01 := dString("1")
	v2_tStr_02 := dString("2")
	v2_tStr_03 := dString("3")

	v2_tStr_01_len_str := fmt.Sprintf("%d", len(v2_tStr_01))
	v2_tStr_02_len_str := fmt.Sprintf("%d", len(v2_tStr_02))
	v2_tStr_03_len_str := fmt.Sprintf("%d", len(v2_tStr_03))

	v2 := [3]dString{v2_tStr_01, v2_tStr_02, v2_tStr_03}

	v2_len := fmt.Sprintf("%d", len(v2))
	v2_cap := fmt.Sprintf("%d", cap(v2))

	v2_nil := (*[3]dString)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_addr_ptr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "debuging.dString"

	v2_value := fmt.Sprintf("(len=%s cap=%s) [\n(%s)(len=%s) 1,\n(%s)(len=%s) 2,\n(%s)(len=%s) 3,\n]",
		v2_len, v2_cap,
		v2_type, v2_tStr_01_len_str,
		v2_type, v2_tStr_02_len_str,
		v2_type, v2_tStr_03_len_str)
	v2_value_str := v2_value

	if IsDisabledUnsafeReflect {
		v2_value_str = fmt.Sprintf("(len=%s cap=%s) {\n(%s)(len=%s) \"1\",\n(%s)(len=%s) \"2\",\n(%s)(len=%s) \"3\",\n}",
			v2_len, v2_cap,
			v2_type, v2_tStr_01_len_str,
			v2_type, v2_tStr_02_len_str,
			v2_type, v2_tStr_03_len_str)
	}

	addToDumpDatas(v2, "([3]"+v2_type+")"+v2_value)
	addToDumpDatas(v2_ptr, "(*[3]"+v2_type+")("+v2_addr+")"+v2_value_str+"")
	addToDumpDatas(&v2_ptr, "(**[3]"+v2_type+")("+v2_addr_ptr+"->"+v2_addr+")"+v2_value_str+"")
	addToDumpDatas(v2_nil, "(*[3]"+v2_type+")<nil>")

	//Array containing interfaces
	v3_01_str := "1st"
	v3_02_int := int(2)
	v3_03_uint := uint(3)
	v3_01_len_str := fmt.Sprintf("%d", len(v3_01_str))

	v3 := [3]interface{}{v3_01_str, v3_02_int, v3_03_uint}

	v3_len := fmt.Sprintf("%d", len(v3))
	v3_cap := fmt.Sprintf("%d", cap(v3))
	v3_ptr := &v3
	v3_nil := (*[3]interface{})(nil)

	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_addr_ptr := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "[3]interface {}"
	v3_01_type := "string"
	v3_02_type := "int"
	v3_03_type := "uint"

	v3_value := fmt.Sprintf("(len=%s cap=%s) [\n(%s)(len=%s) \"1st\",\n(%s)2,\n(%s)3,\n]",
		v3_len, v3_cap,
		v3_01_type, v3_01_len_str,
		v3_02_type, v3_03_type)

	v3_value_str := v3_value

	addToDumpDatas(v3, "("+v3_type+")"+v3_value)
	addToDumpDatas(v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_value_str+"")
	addToDumpDatas(&v3_ptr, "(**"+v3_type+")("+v3_addr_ptr+"->"+v3_addr+")"+v3_value_str+"")
	addToDumpDatas(v3_nil, "(*"+v3_type+")<nil>")

	v4 := [34]byte{
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30,
		0x31, 0x32,
	}
	v4_len_str := fmt.Sprintf("%d", len(v4))
	v4_cap_str := fmt.Sprintf("%d", cap(v4))
	v4_nil := (*[34]byte)(nil)
	v4_ptr := &v4
	v4_addr := fmt.Sprintf("%p", v4_ptr)
	v4_ptr_addr := fmt.Sprintf("%p", &v4_ptr)
	v4_type := "[34]uint8"

	v4_value_str := "(len=" + v4_len_str + " cap=" + v4_cap_str + ") " +
		"[\n 00000000  11 12 13 14 15 16 17 18  19 1a 1b 1c 1d 1e 1f 20" +
		"  |............... |\n" +
		" 00000010  21 22 23 24 25 26 27 28  29 2a 2b 2c 2d 2e 2f 30" +
		"  |!\"#$%&'()*+,-./0|\n" +
		" 00000020  31 32                                           " +
		"  |12|\n]"

	addToDumpDatas(v4, "("+v4_type+")"+v4_value_str)
	addToDumpDatas(v4_ptr, "(*"+v4_type+")("+v4_addr+")"+v4_value_str+"")
	addToDumpDatas(&v4_ptr, "(**"+v4_type+")("+v4_ptr_addr+"->"+v4_addr+")"+v4_value_str+"")
	addToDumpDatas(v4_nil, "(*"+v4_type+")<nil>")
	doDumpTesting(t)
}

func TestDumpSlice(t *testing.T) {

	// Slice containing standard float32 values.
	v := []float32{3.14, 6.28, 12.56}
	v_len := fmt.Sprintf("%d", len(v))
	v_cap := fmt.Sprintf("%d", cap(v))
	v_nil := (*[]float32)(nil)

	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "[]float32"
	vt_type := "float32"
	v_value_str := fmt.Sprintf("(len=%s cap=%s) [\n(%s)3.14,\n(%s)6.28,\n(%s)12.56,\n]",
		v_len, v_cap, vt_type, vt_type, vt_type)

	addToDumpDatas(v, "("+v_type+")"+v_value_str)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_addr+")"+v_value_str+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value_str+"")
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	v2_tStr_01 := dString("1")
	v2_tStr_02 := dString("2")
	v2_tStr_03 := dString("3")

	v2_tStr_01_len_str := fmt.Sprintf("%d", len(v2_tStr_01))
	v2_tStr_02_len_str := fmt.Sprintf("%d", len(v2_tStr_02))
	v2_tStr_03_len_str := fmt.Sprintf("%d", len(v2_tStr_03))

	v2 := []dString{v2_tStr_01, v2_tStr_02, v2_tStr_03}

	v2_len := fmt.Sprintf("%d", len(v2))
	v2_cap := fmt.Sprintf("%d", cap(v2))

	v2_nil := (*[]dString)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_addr_ptr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "debuging.dString"

	v2_value := fmt.Sprintf("(len=%s cap=%s) [\n(%s)(len=%s) 1,\n(%s)(len=%s) 2,\n(%s)(len=%s) 3,\n]",
		v2_len, v2_cap,
		v2_type, v2_tStr_01_len_str,
		v2_type, v2_tStr_02_len_str,
		v2_type, v2_tStr_03_len_str)
	v2_value_str := v2_value

	if IsDisabledUnsafeReflect {
		v2_value_str = fmt.Sprintf("(len=%s cap=%s) [\n(%s)(len=%s) \"1\",\n(%s)(len=%s) \"2\",\n(%s)(len=%s) \"3\",\n]",
			v2_len, v2_cap,
			v2_type, v2_tStr_01_len_str,
			v2_type, v2_tStr_02_len_str,
			v2_type, v2_tStr_03_len_str)
	}

	addToDumpDatas(v2, "([]"+v2_type+")"+v2_value)
	addToDumpDatas(v2_ptr, "(*[]"+v2_type+")("+v2_addr+")"+v2_value_str+"")
	addToDumpDatas(&v2_ptr, "(**[]"+v2_type+")("+v2_addr_ptr+"->"+v2_addr+")"+v2_value_str+"")
	addToDumpDatas(v2_nil, "(*[]"+v2_type+")<nil>")

	//Array containing interfaces
	v3_01_str := "1st"
	v3_02_int := int(2)
	v3_03_uint := uint(3)
	v3_01_len_str := fmt.Sprintf("%d", len(v3_01_str))

	v3 := [3]interface{}{v3_01_str, v3_02_int, v3_03_uint}

	v3_len := fmt.Sprintf("%d", len(v3))
	v3_cap := fmt.Sprintf("%d", cap(v3))
	v3_ptr := &v3
	v3_nil := (*[3]interface{})(nil)

	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_addr_ptr := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "[3]interface {}"
	v3_01_type := "string"
	v3_02_type := "int"
	v3_03_type := "uint"

	v3_value := fmt.Sprintf("(len=%s cap=%s) [\n(%s)(len=%s) \"1st\",\n(%s)2,\n(%s)3,\n]",
		v3_len, v3_cap,
		v3_01_type, v3_01_len_str,
		v3_02_type, v3_03_type)

	v3_value_str := v3_value

	addToDumpDatas(v3, "("+v3_type+")"+v3_value)
	addToDumpDatas(v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_value_str+"")
	addToDumpDatas(&v3_ptr, "(**"+v3_type+")("+v3_addr_ptr+"->"+v3_addr+")"+v3_value_str+"")
	addToDumpDatas(v3_nil, "(*"+v3_type+")<nil>")

	v4 := []byte{
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30,
		0x31, 0x32,
	}
	v4_len_str := fmt.Sprintf("%d", len(v4))
	v4_cap_str := fmt.Sprintf("%d", cap(v4))
	v4_nil := (*[]byte)(nil)
	v4_ptr := &v4
	v4_addr := fmt.Sprintf("%p", v4_ptr)
	v4_ptr_addr := fmt.Sprintf("%p", &v4_ptr)
	v4_type := "[]uint8"

	v4_value_str := "(len=" + v4_len_str + " cap=" + v4_cap_str + ") " +
		"[\n 00000000  11 12 13 14 15 16 17 18  19 1a 1b 1c 1d 1e 1f 20" +
		"  |............... |\n" +
		" 00000010  21 22 23 24 25 26 27 28  29 2a 2b 2c 2d 2e 2f 30" +
		"  |!\"#$%&'()*+,-./0|\n" +
		" 00000020  31 32                                           " +
		"  |12|\n]"

	addToDumpDatas(v4, "("+v4_type+")"+v4_value_str)
	addToDumpDatas(v4_ptr, "(*"+v4_type+")("+v4_addr+")"+v4_value_str+"")
	addToDumpDatas(&v4_ptr, "(**"+v4_type+")("+v4_ptr_addr+"->"+v4_addr+")"+v4_value_str+"")
	addToDumpDatas(v4_nil, "(*"+v4_type+")<nil>")
	doDumpTesting(t)

}

func TestDumpInterface(t *testing.T) {

	// Nil interface.
	var v interface{}
	v_nil := (*interface{})(nil)
	v_ptr := &v

	v_str := fmt.Sprintf("%p", v_ptr)
	v_ptr_str := fmt.Sprintf("%p", &v_ptr)

	v_type := "interface {}"
	v_value := "<nil>"

	addToDumpDatas(v, "("+v_type+")"+v_value)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_str+")"+v_value+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value)
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	//Sub-interface.
	v2 := interface{}(uint16(65535))
	v2_ptr := &v2

	v2_str := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_str := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "uint16"
	v2_value := "65535"

	addToDumpDatas(v2, "("+v2_type+")"+v2_value)
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_str+")"+v2_value+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_str+"->"+v2_str+")"+v2_value)

	doDumpTesting(t)
}

func TestDumpMap(t *testing.T) {

	v_key_1 := "1st"
	v_key_2 := "2nd"
	v := map[string]int{v_key_1: 1, v_key_2: 2}
	v_nil := map[string]int(nil)
	v_ptr_nil := (*map[string]int)(nil)
	v_ptr := &v

	v_len := fmt.Sprintf("%d", len(v))

	v_str := fmt.Sprintf("%p", v_ptr)
	v_ptr_str := fmt.Sprintf("%p", &v_ptr)

	v_type := "map[string]int"
	v_value_1 := fmt.Sprintf("(len=%s) {\n\"1st\":1,\n\"2nd\":2,\n}",
		v_len)
	v_value_2 := fmt.Sprintf("(len=%s) {\n\"2nd\":2,\n\"1st\":1,\n}",
		v_len)
	addToDumpDatas(v, "("+v_type+")"+v_value_1, "("+v_type+")"+v_value_2)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_str+")"+v_value_1+"", "(*"+v_type+")("+v_str+")"+v_value_2+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value_1, "(**"+v_type+")("+v_ptr_str+"->"+v_str+")"+v_value_2)
	addToDumpDatas(v_nil, "("+v_type+")<nil>")
	addToDumpDatas(v_ptr_nil, "(*"+v_type+")<nil>")

	v2_key_1 := dPtrString("1st")
	v2_key_2 := dPtrString("2nd")

	v2 := map[dPtrString]int{v2_key_1: 1, v2_key_2: 2}
	v2_nil := map[dPtrString]int(nil)
	v2_ptr_nil := (*map[dPtrString]int)(nil)
	v2_ptr := &v2

	v2_len := fmt.Sprintf("%d", len(v2))

	v2_str := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_str := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "map[debuging.dPtrString]int"
	v2_value_1 := fmt.Sprintf("(len=%s) {\n1st:1,\n2nd:2,\n}",
		v2_len)
	v2_value_2 := fmt.Sprintf("(len=%s) {\n2nd:2,\n1st:1,\n}",
		v2_len)
	addToDumpDatas(v2, "("+v2_type+")"+v2_value_1, "("+v2_type+")"+v2_value_2)
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_str+")"+v2_value_1+"", "(*"+v2_type+")("+v2_str+")"+v2_value_2+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_str+"->"+v2_str+")"+v2_value_1, "(**"+v2_type+")("+v2_ptr_str+"->"+v2_str+")"+v2_value_2)
	addToDumpDatas(v2_nil, "("+v2_type+")<nil>")
	addToDumpDatas(v2_ptr_nil, "(*"+v2_type+")<nil>")

	//map with interface key and value
	v3_key := "1st"
	v3_key_len := fmt.Sprintf("%d", len(v3_key))

	v3 := map[interface{}]interface{}{v3_key: 1}
	v3_len := fmt.Sprintf("%d", len(v3))
	v3_nil := map[interface{}]interface{}(nil)
	v3_ptr_nil := (*map[interface{}]interface{})(nil)

	v3_ptr := &v3
	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_addr := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "map[interface {}]interface {}"
	v3_key_type := "string"
	v3_value_type := "int"

	v3_value := fmt.Sprintf("(len=%s) {\n(%s)(len=%s) \"1st\":(%s)1,\n}",
		v3_len, v3_key_type, v3_key_len, v3_value_type)

	addToDumpDatas(v3, "("+v3_type+")"+v3_value)
	addToDumpDatas(v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_value)
	addToDumpDatas(&v3_ptr, "(**"+v3_type+")("+v3_ptr_addr+"->"+v3_addr+")"+v3_value)
	addToDumpDatas(v3_nil, "("+v3_type+")<nil>")
	addToDumpDatas(v3_ptr_nil, "(*"+v3_type+")<nil>")

	//map with nil interface value
	v4_key := "nil"

	v4 := map[string]interface{}{v4_key: nil}
	v4_len := fmt.Sprintf("%d", len(v4))
	v4_nil := (map[string]interface{})(nil)
	v4_ptr_nil := (*map[string]interface{})(nil)

	v4_ptr := &v4
	v4_addr := fmt.Sprintf("%p", v4_ptr)
	v4_ptr_addr := fmt.Sprintf("%p", &v4_ptr)

	v4_type := "map[string]interface {}"
	v4_value_type := "interface {}"

	v4_value := fmt.Sprintf("(len=%s) {\n\"nil\":(%s)<nil>,\n}",
		v4_len, v4_value_type)

	addToDumpDatas(v4, "("+v4_type+")"+v4_value)
	addToDumpDatas(v4_ptr, "(*"+v4_type+")("+v4_addr+")"+v4_value)
	addToDumpDatas(&v4_ptr, "(**"+v4_type+")("+v4_ptr_addr+"->"+v4_addr+")"+v4_value)
	addToDumpDatas(v4_nil, "("+v4_type+")<nil>")
	addToDumpDatas(v4_ptr_nil, "(*"+v4_type+")<nil>")
	doDumpTesting(t)
}

func TestDumpStruct(t *testing.T) {

	//Struct with primitives
	type d1 struct {
		a int8
		b uint8
	}
	v := d1{127, 255}
	v_nil := (*d1)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "debuging.d1"
	v_f1_type := "int8"
	v_f2_type := "uint8"

	v_value := fmt.Sprintf("{\n a:(%s)127,\n b:(%s)255,\n}", v_f1_type, v_f2_type)

	addToDumpDatas(v, "("+v_type+")"+v_value)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_addr+")"+v_value)
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	//Struct contains another Struct
	type d2 struct {
		A d1
		B bool
	}
	v2 := d2{d1{127, 255}, true}
	v2_nil := (*d2)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "debuging.d2"
	v2_f1_type := "debuging.d1"
	v2_f1_f1_type := "int8"
	v2_f1_f2_type := "uint8"
	v2_f2_type := "bool"

	v2_value := fmt.Sprintf("{\n A:(%s){\n  a:(%s)127,\n  b:(%s)255,\n },\n B:(%s)true,\n}",
		v2_f1_type, v2_f1_f1_type, v2_f1_f2_type, v2_f2_type)

	addToDumpDatas(v2, "("+v2_type+")"+v2_value)
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_value)
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_value)
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	doDumpTesting(t)
}
func TestDumpStruct2(t *testing.T) {

	//Struct contains another Struct
	type d3 struct {
		C dPtrString
		D dPtrString
	}
	v3 := d3{("1st"), ("2nd")}
	v3_nil := (*d3)(nil)
	v3_ptr := &v3

	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_addr := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "debuging.d3"
	v3_f1_type := "debuging.dPtrString"
	v3_f2_type := "debuging.dPtrString"
	v3_f1_len := fmt.Sprintf("%d", len(v3.C))
	v3_f2_len := fmt.Sprintf("%d", len(v3.D))

	v3_value := fmt.Sprintf("{\n C:(%s)(len=%s) 1st,\n D:(%s)(len=%s) 2nd,\n}", v3_f1_type, v3_f1_len, v3_f2_type, v3_f2_len)
	v3_value_s := v3_value

	if IsDisabledUnsafeReflect {
		v3_value = fmt.Sprintf("{\n C:(%s)(len=%s) \"1st\",\n D:(%s)(len=%s) \"2nd\",\n}", v3_f1_type, v3_f1_len, v3_f2_type, v3_f2_len)
		v3_value_s = fmt.Sprintf("{\n C:(%s)(len=%s) \"1st\",\n D:(%s)(len=%s) \"2nd\",\n}", v3_f1_type, v3_f1_len, v3_f2_type, v3_f2_len)
	}

	addToDumpDatas(v3, "("+v3_type+")"+v3_value)
	addToDumpDatas(v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_value)
	addToDumpDatas(&v3_ptr, "(**"+v3_type+")("+v3_ptr_addr+"->"+v3_addr+")"+v3_value_s)
	addToDumpDatas(v3_nil, "(*"+v3_type+")<nil>")
	doDumpTesting(t)
}
func TestDumpUintptr(t *testing.T) {

	//Struct contains another Struct

	// Null pointer.
	v := uintptr(0)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "uintptr"
	nil_type := "<nil>"
	addToDumpDatas(v, "("+v_type+")"+nil_type)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_addr+")"+nil_type+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+nil_type+"")

	// Address of real variable.
	i := 1
	v2 := uintptr(unsafe.Pointer(&i))
	v2_nil := (*uintptr)(nil)
	v2_ptr := &v2
	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)
	v2_type := "uintptr"
	v2_str_str := fmt.Sprintf("%p", &i)
	addToDumpDatas(v2, "("+v2_type+")"+v2_str_str+"")
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_str_str+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_str_str+"")
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	doDumpTesting(t)
}

func TestDumpUnsafePointer(t *testing.T) {

	// Null pointer.
	v := unsafe.Pointer(uintptr(0))
	v_nil := (*unsafe.Pointer)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "unsafe.Pointer"
	nil_type := "<nil>"
	addToDumpDatas(v, "("+v_type+")"+nil_type)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_addr+")"+nil_type+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+nil_type+"")
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	// Address of real variable.
	i := 1
	v2 := unsafe.Pointer(&i)
	v2_nil := (*unsafe.Pointer)(nil)
	v2_ptr := &v2
	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)
	v2_type := "unsafe.Pointer"
	v2_str_str := fmt.Sprintf("%p", &i)
	addToDumpDatas(v2, "("+v2_type+")"+v2_str_str+"")
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_str_str+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_str_str+"")
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	doDumpTesting(t)
}

func TestDumpChan(t *testing.T) {

	// Null pointer.
	var v chan int
	v_nil := (*chan int)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "chan int"
	nil_type := "<nil>"
	addToDumpDatas(v, "("+v_type+")"+nil_type)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_addr+")"+nil_type+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+nil_type+"")
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	// Address of real variable.
	v2 := make(chan int)
	v2_nil := (*chan int)(nil)
	v2_ptr := &v2
	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)
	v2_type := "chan int"
	v2_str_str := fmt.Sprintf("%p", v2)
	addToDumpDatas(v2, "("+v2_type+")"+v2_str_str+"")
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_str_str+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_str_str+"")
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	doDumpTesting(t)
}
func testFunc0() {
	fmt.Print("GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG")
}
func TestDumpFunc(t *testing.T) {

	// Null pointer.
	// Function with no params and no returns.
	v := testFunc0
	v_nil := (*func())(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "func()"
	v_str_str := fmt.Sprintf("%p", v)
	addToDumpDatas(v, "("+v_type+")"+v_str_str)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_addr+")"+v_str_str+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_str_str+"")
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	// Address of real variable.
	v2 := TestDumpFunc
	v2_nil := (*func(*testing.T))(nil)
	v2_ptr := &v2
	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)
	v2_type := "func(*testing.T)"
	v2_str_str := fmt.Sprintf("%p", v2)
	addToDumpDatas(v2, "("+v2_type+")"+v2_str_str+"")
	addToDumpDatas(v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_str_str+"")
	addToDumpDatas(&v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_str_str+"")
	addToDumpDatas(v2_nil, "(*"+v2_type+")<nil>")

	// Function with multiple params and multiple returns.
	var v3 = func(i int, s string) (b bool, err error) {
		return true, nil
	}
	v3_nil := (*func(int, string) (bool, error))(nil)
	v3_ptr := &v3
	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_addr := fmt.Sprintf("%p", &v3_ptr)
	v3_type := "func(int, string) (bool, error)"
	v3_str_str := fmt.Sprintf("%p", v3)
	addToDumpDatas(v3, "("+v3_type+")"+v3_str_str+"")
	addToDumpDatas(v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_str_str+"")
	addToDumpDatas(&v3_ptr, "(**"+v3_type+")("+v3_ptr_addr+"->"+v3_addr+")"+v3_str_str+"")
	addToDumpDatas(v3_nil, "(*"+v3_type+")<nil>")
	doDumpTesting(t)
}

type dError int

func (e dError) Error() string {
	return fmt.Sprintf("error: %d", int(e))
}
func TestDumpError(t *testing.T) {

	// Null pointer.
	// Function with no params and no returns.
	v := dError(999)
	v_nil := (*dError)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "debuging.dError"
	v_str_str := "error: 999"
	addToDumpDatas(v, "("+v_type+")"+v_str_str)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_addr+")"+v_str_str+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_str_str+"")
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	doDumpTesting(t)
}

type dPanicer int

func (p dPanicer) String() string {
	panic("dPanicer")
}
func TestDumpPanic(t *testing.T) {

	// Null pointer.
	// Function with no params and no returns.
	v := dPanicer(999)
	v_nil := (*dPanicer)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "debuging.dPanicer"
	v_str_str := "(PANIC=dPanicer)999"
	addToDumpDatas(v, "("+v_type+")"+v_str_str)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_addr+")"+v_str_str+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_str_str+"")
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	doDumpTesting(t)
}
func TestDumpSelfRef(t *testing.T) {

	// Struct that is circular through self referencing.
	type selfRef struct {
		ref *selfRef
	}
	v := selfRef{nil}
	v.ref = &v

	v_nil := (*selfRef)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "debuging.selfRef"
	v_str_str := fmt.Sprintf("{\n ref:(*%s)(%s){\n"+
		"  ref:(*%s)(%s)<already shown>,\n },\n}", v_type, v_addr, v_type, v_addr)

	v_ptr_str_str := fmt.Sprintf("{\n ref:(*%s)(%s)<already shown>,\n}", v_type, v_addr)

	addToDumpDatas(v, "("+v_type+")"+v_str_str)
	addToDumpDatas(v_ptr, "(*"+v_type+")("+v_addr+")"+v_ptr_str_str+"")
	addToDumpDatas(&v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_ptr_str_str+"")
	addToDumpDatas(v_nil, "(*"+v_type+")<nil>")

	doDumpTesting(t)
}

type dRefRef0 struct {
	ref1 *dRefRef1
}
type dRefRef1 struct {
	ref0 *dRefRef0
}

func TestDumpRefRef(t *testing.T) {

	v0 := dRefRef0{nil}
	v1 := dRefRef1{nil}

	v0.ref1, v1.ref0 = &v1, &v0

	v0_nil := (*dRefRef0)(nil)
	v0_ptr := &v0
	v0_addr := fmt.Sprintf("%p", v0_ptr)
	v0_ptr_addr := fmt.Sprintf("%p", &v0_ptr)
	v0_type := "debuging.dRefRef0"
	v1_type := "debuging.dRefRef1"

	v1_ptr := &v1
	v1_addr := fmt.Sprintf("%p", v1_ptr)

	v0_str_str := fmt.Sprintf(
		"{\n ref1:(*%s)(%s){\n"+
			"  ref0:(*%s)(%s){\n"+
			"   ref1:(*%s)(%s)<already shown>,\n"+
			"  },\n"+
			" },\n"+
			"}",
		v1_type, v1_addr,
		v0_type, v0_addr,
		v1_type, v1_addr)

	v0_ptr_str_str := fmt.Sprintf(
		"{\n ref1:(*%s)(%s){\n"+
			"  ref0:(*%s)(%s)<already shown>,\n"+
			" },\n"+
			"}",
		v1_type, v1_addr,
		v0_type, v0_addr)

	addToDumpDatas(v0, "("+v0_type+")"+v0_str_str)
	addToDumpDatas(v0_ptr, "(*"+v0_type+")("+v0_addr+")"+v0_ptr_str_str+"")
	addToDumpDatas(&v0_ptr, "(**"+v0_type+")("+v0_ptr_addr+"->"+v0_addr+")"+v0_ptr_str_str+"")
	addToDumpDatas(v0_nil, "(*"+v0_type+")<nil>")

	doDumpTesting(t)
}

type dIndirCir0 struct {
	ref1 *dIndirCir1
}
type dIndirCir1 struct {
	ref2 *dIndirCir2
}
type dIndirCir2 struct {
	ref0 *dIndirCir0
}

func TestDumpIndirCir(t *testing.T) {

	v0 := dIndirCir0{nil}
	v1 := dIndirCir1{nil}
	v2 := dIndirCir2{nil}

	v0.ref1, v1.ref2, v2.ref0 = &v1, &v2, &v0

	v0_nil := (*dIndirCir0)(nil)
	v0_ptr := &v0
	v0_addr := fmt.Sprintf("%p", v0_ptr)
	v0_ptr_addr := fmt.Sprintf("%p", &v0_ptr)

	v0_type := "debuging.dIndirCir0"
	v1_type := "debuging.dIndirCir1"
	v2_type := "debuging.dIndirCir2"

	v1_ptr := &v1
	v2_ptr := &v2
	v1_addr := fmt.Sprintf("%p", v1_ptr)
	v2_addr := fmt.Sprintf("%p", v2_ptr)

	v0_str_str := fmt.Sprintf(
		"{\n ref1:(*%s)(%s){\n"+
			"  ref2:(*%s)(%s){\n"+
			"   ref0:(*%s)(%s){\n"+
			"    ref1:(*%s)(%s)<already shown>,\n"+
			"   },\n"+
			"  },\n"+
			" },\n"+
			"}",
		v1_type, v1_addr,
		v2_type, v2_addr,
		v0_type, v0_addr,
		v1_type, v1_addr)

	v0_ptr_str_str := fmt.Sprintf(
		"{\n ref1:(*%s)(%s){\n"+
			"  ref2:(*%s)(%s){\n"+
			"   ref0:(*%s)(%s)<already shown>,\n"+
			"  },\n"+
			" },\n"+
			"}",
		v1_type, v1_addr,
		v2_type, v2_addr,
		v0_type, v0_addr)

	addToDumpDatas(v0, "("+v0_type+")"+v0_str_str)
	addToDumpDatas(v0_ptr, "(*"+v0_type+")("+v0_addr+")"+v0_ptr_str_str+"")
	addToDumpDatas(&v0_ptr, "(**"+v0_type+")("+v0_ptr_addr+"->"+v0_addr+")"+v0_ptr_str_str+"")
	addToDumpDatas(v0_nil, "(*"+v0_type+")<nil>")

	doDumpTesting(t)
}

func TestDumpSortedKeys(t *testing.T) {

	for _, c := range []struct {
		in   map[int]string
		want string
	}{
		{map[int]string{1: "1", 3: "3", 2: "2"},
			"(map[int]string)(len=3) {\n" +
				"1:\"1\",\n" +
				"2:\"2\",\n" +
				"3:\"3\",\n" +
				"}"},
	} {
		cfg := DebugConfig{SortKeys: true}
		buf := new(bytes.Buffer)
		DumpWithConfigOfWriter(&cfg, buf, c.in)
		got := buf.String()
		if got != c.want {
			t.Errorf("Reverse(%v) want:\n%v got:\n%v,", c.in, c.want, got)
		}
	}
}
func TestDumpSortedKeys1(t *testing.T) {

	for _, c := range []struct {
		in   map[dString]int
		want string
	}{
		{map[dString]int{"1": 1, "3": 3, "2": 2},
			"(map[debuging.dString]int)(len=3) {\n" +
				"1:1,\n" +
				"2:2,\n" +
				"3:3,\n" +
				"}"},
	} {
		cfg := DebugConfig{SortKeys: true}
		buf := new(bytes.Buffer)
		DumpWithConfigOfWriter(&cfg, buf, c.in)
		got := buf.String()
		if got != c.want {
			t.Errorf("Reverse(%v) want:\n%v got:\n%v,", c.in, c.want, got)
		}
	}
}

func TestDumpSortedKeys2(t *testing.T) {

	in := map[dPtrString]int{dPtrString("1"): 1, dPtrString("3"): 3, dPtrString("2"): 2}
	want := "(map[debuging.dPtrString]int)(len=3) {\n" +
		"1:1,\n" +
		"3:3,\n" +
		"2:2,\n" +
		"}"
	if !IsDisabledUnsafeReflect {
		want = "(map[debuging.dPtrString]int)(len=3) {\n" +
			"1:1,\n" +
			"2:2,\n" +
			"3:3,\n" +
			"}"
	}
	cfg := DebugConfig{SortKeys: true}
	buf := new(bytes.Buffer)
	DumpWithConfigOfWriter(&cfg, buf, in)
	got := buf.String()
	if got != want {
		t.Errorf("Reverse(%v) want:\n%v got:\n%v,", in, want, got)
	}
}

func TestDumpSortedKeys3(t *testing.T) {

	in := map[dError]int{dError(1): 1, dError(3): 3, dError(2): 2}
	want := "(map[debuging.dError]int)(len=3) {\n" +
		"1:1,\n" +
		"2:2,\n" +
		"3:3,\n" +
		"}"
	if !IsDisabledUnsafeReflect {
		want = "(map[debuging.dError]int)(len=3) {\n" +
			"error: 1:1,\n" +
			"error: 2:2,\n" +
			"error: 3:3,\n" +
			"}"
	}
	cfg := DebugConfig{SortKeys: true}
	buf := new(bytes.Buffer)
	DumpWithConfigOfWriter(&cfg, buf, in)
	got := buf.String()
	if got != want {
		t.Errorf("Reverse(%v) want:\n%v got:\n%v,", in, want, got)
	}
}
func doDumpTesting(t *testing.T) {
	for i, d := range dumpDatas {
		buf := new(bytes.Buffer)
		DumapWithWriter(buf, d.in)
		out := buf.String()

		if !testInDumpOn(out, d.on) {
			t.Errorf("Dump Error  #%d \n ned:[%v] \n got:[%s] ", i, d.onString(), out)
		}
	}
}

func testInDumpOn(result string, wants []string) bool {
	for _, item := range wants {
		if result == item {
			return true
		}
	}
	return false
}
