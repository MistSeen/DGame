package debuging

import (
	"bytes"
	"fmt"
	. "server/core/debuging"
	"testing"
	"unsafe"
)

type printData struct {
	format string
	in     interface{}
	on     []string
}

func (f *printData) onString() string {
	s := ""
	for i, on := range f.on {
		if i > 1 {
			s += fmt.Sprintf("#%d: %s", i+1, on)
		}
		s += on
	}
	return s
}

var printDatas = make([]printData, 0)

func addToPrintDatas(f string, in interface{}, on ...string) {
	v := printData{f, in, on}
	printDatas = append(printDatas, v)
}
func TestPrintInt(t *testing.T) {

	v := int8(127)
	v_nil := (*int8)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "int8"
	v_value := "127"

	addToPrintDatas("%v", v, v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)
	addToPrintDatas("%v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, v_value)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	v2 := int16(32767)
	v2_nil := (*int16)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "int16"
	v2_v2alue := "32767"

	addToPrintDatas("%v", v2, v2_v2alue)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_v2alue)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_v2alue)
	addToPrintDatas("%v", v2_nil, "<nil>")
	addToPrintDatas("%+v", v2, v2_v2alue)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", v2_nil, "<nil>")
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_nil, "(*"+v2_type+")"+"<nil>")
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_nil, "(*"+v2_type+")"+"<nil>")

	v3 := int32(2147483647)
	v3_nil := (*int32)(nil)
	v3_ptr := &v3

	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_addr := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "int32"
	v3_v3alue := "2147483647"

	addToPrintDatas("%v", v3, v3_v3alue)
	addToPrintDatas("%v", v3_ptr, "<*>"+v3_v3alue)
	addToPrintDatas("%v", &v3_ptr, "<**>"+v3_v3alue)
	addToPrintDatas("%v", v3_nil, "<nil>")
	addToPrintDatas("%+v", v3, v3_v3alue)
	addToPrintDatas("%+v", v3_ptr, "<*>("+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%+v", &v3_ptr, "<**>("+v3_ptr_addr+"->"+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%+v", v3_nil, "<nil>")
	addToPrintDatas("%#v", v3, "("+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", v3_ptr, "(*"+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", &v3_ptr, "(**"+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", v3_nil, "(*"+v3_type+")"+"<nil>")
	addToPrintDatas("%#+v", v3, "("+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#+v", v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%#+v", &v3_ptr, "(**"+v3_type+")("+v3_ptr_addr+"->"+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%#+v", v3_nil, "(*"+v3_type+")"+"<nil>")

	v4 := int64(9223372036854775807)
	v4_nil := (*int64)(nil)
	v4_ptr := &v4

	v4_addr := fmt.Sprintf("%p", v4_ptr)
	v4_ptr_addr := fmt.Sprintf("%p", &v4_ptr)

	v4_type := "int64"
	v4_v4alue := "9223372036854775807"

	addToPrintDatas("%v", v4, v4_v4alue)
	addToPrintDatas("%v", v4_ptr, "<*>"+v4_v4alue)
	addToPrintDatas("%v", &v4_ptr, "<**>"+v4_v4alue)
	addToPrintDatas("%v", v4_nil, "<nil>")
	addToPrintDatas("%+v", v4, v4_v4alue)
	addToPrintDatas("%+v", v4_ptr, "<*>("+v4_addr+")"+v4_v4alue)
	addToPrintDatas("%+v", &v4_ptr, "<**>("+v4_ptr_addr+"->"+v4_addr+")"+v4_v4alue)
	addToPrintDatas("%+v", v4_nil, "<nil>")
	addToPrintDatas("%#v", v4, "("+v4_type+")"+v4_v4alue)
	addToPrintDatas("%#v", v4_ptr, "(*"+v4_type+")"+v4_v4alue)
	addToPrintDatas("%#v", &v4_ptr, "(**"+v4_type+")"+v4_v4alue)
	addToPrintDatas("%#v", v4_nil, "(*"+v4_type+")"+"<nil>")
	addToPrintDatas("%#+v", v4, "("+v4_type+")"+v4_v4alue)
	addToPrintDatas("%#+v", v4_ptr, "(*"+v4_type+")("+v4_addr+")"+v4_v4alue)
	addToPrintDatas("%#+v", &v4_ptr, "(**"+v4_type+")("+v4_ptr_addr+"->"+v4_addr+")"+v4_v4alue)
	addToPrintDatas("%#+v", v4_nil, "(*"+v4_type+")"+"<nil>")

	v5 := int(9223372036854775807)
	v5_nil := (*int)(nil)
	v5_ptr := &v5

	v5_addr := fmt.Sprintf("%p", v5_ptr)
	v5_ptr_addr := fmt.Sprintf("%p", &v5_ptr)

	v5_type := "int"
	v5_v5alue := "9223372036854775807"

	addToPrintDatas("%v", v5, v5_v5alue)
	addToPrintDatas("%v", v5_ptr, "<*>"+v5_v5alue)
	addToPrintDatas("%v", &v5_ptr, "<**>"+v5_v5alue)
	addToPrintDatas("%v", v5_nil, "<nil>")
	addToPrintDatas("%+v", v5, v5_v5alue)
	addToPrintDatas("%+v", v5_ptr, "<*>("+v5_addr+")"+v5_v5alue)
	addToPrintDatas("%+v", &v5_ptr, "<**>("+v5_ptr_addr+"->"+v5_addr+")"+v5_v5alue)
	addToPrintDatas("%+v", v5_nil, "<nil>")
	addToPrintDatas("%#v", v5, "("+v5_type+")"+v5_v5alue)
	addToPrintDatas("%#v", v5_ptr, "(*"+v5_type+")"+v5_v5alue)
	addToPrintDatas("%#v", &v5_ptr, "(**"+v5_type+")"+v5_v5alue)
	addToPrintDatas("%#v", v5_nil, "(*"+v5_type+")"+"<nil>")
	addToPrintDatas("%#+v", v5, "("+v5_type+")"+v5_v5alue)
	addToPrintDatas("%#+v", v5_ptr, "(*"+v5_type+")("+v5_addr+")"+v5_v5alue)
	addToPrintDatas("%#+v", &v5_ptr, "(**"+v5_type+")("+v5_ptr_addr+"->"+v5_addr+")"+v5_v5alue)
	addToPrintDatas("%#+v", v5_nil, "(*"+v5_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintUInt(t *testing.T) {

	v := uint8(255)
	v_nil := (*uint8)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "uint8"
	v_value := "255"

	addToPrintDatas("%v", v, v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)
	addToPrintDatas("%v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, v_value)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	v2 := uint16(65535)
	v2_nil := (*uint16)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "uint16"
	v2_v2alue := "65535"

	addToPrintDatas("%v", v2, v2_v2alue)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_v2alue)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_v2alue)
	addToPrintDatas("%v", v2_nil, "<nil>")
	addToPrintDatas("%+v", v2, v2_v2alue)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", v2_nil, "<nil>")
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_nil, "(*"+v2_type+")"+"<nil>")
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_nil, "(*"+v2_type+")"+"<nil>")

	v3 := uint32(4294967295)
	v3_nil := (*uint32)(nil)
	v3_ptr := &v3

	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_addr := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "uint32"
	v3_v3alue := "4294967295"

	addToPrintDatas("%v", v3, v3_v3alue)
	addToPrintDatas("%v", v3_ptr, "<*>"+v3_v3alue)
	addToPrintDatas("%v", &v3_ptr, "<**>"+v3_v3alue)
	addToPrintDatas("%v", v3_nil, "<nil>")
	addToPrintDatas("%+v", v3, v3_v3alue)
	addToPrintDatas("%+v", v3_ptr, "<*>("+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%+v", &v3_ptr, "<**>("+v3_ptr_addr+"->"+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%+v", v3_nil, "<nil>")
	addToPrintDatas("%#v", v3, "("+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", v3_ptr, "(*"+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", &v3_ptr, "(**"+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", v3_nil, "(*"+v3_type+")"+"<nil>")
	addToPrintDatas("%#+v", v3, "("+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#+v", v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%#+v", &v3_ptr, "(**"+v3_type+")("+v3_ptr_addr+"->"+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%#+v", v3_nil, "(*"+v3_type+")"+"<nil>")

	v4 := uint64(18446744073709551615)
	v4_nil := (*uint64)(nil)
	v4_ptr := &v4

	v4_addr := fmt.Sprintf("%p", v4_ptr)
	v4_ptr_addr := fmt.Sprintf("%p", &v4_ptr)

	v4_type := "uint64"
	v4_v4alue := "18446744073709551615"

	addToPrintDatas("%v", v4, v4_v4alue)
	addToPrintDatas("%v", v4_ptr, "<*>"+v4_v4alue)
	addToPrintDatas("%v", &v4_ptr, "<**>"+v4_v4alue)
	addToPrintDatas("%v", v4_nil, "<nil>")
	addToPrintDatas("%+v", v4, v4_v4alue)
	addToPrintDatas("%+v", v4_ptr, "<*>("+v4_addr+")"+v4_v4alue)
	addToPrintDatas("%+v", &v4_ptr, "<**>("+v4_ptr_addr+"->"+v4_addr+")"+v4_v4alue)
	addToPrintDatas("%+v", v4_nil, "<nil>")
	addToPrintDatas("%#v", v4, "("+v4_type+")"+v4_v4alue)
	addToPrintDatas("%#v", v4_ptr, "(*"+v4_type+")"+v4_v4alue)
	addToPrintDatas("%#v", &v4_ptr, "(**"+v4_type+")"+v4_v4alue)
	addToPrintDatas("%#v", v4_nil, "(*"+v4_type+")"+"<nil>")
	addToPrintDatas("%#+v", v4, "("+v4_type+")"+v4_v4alue)
	addToPrintDatas("%#+v", v4_ptr, "(*"+v4_type+")("+v4_addr+")"+v4_v4alue)
	addToPrintDatas("%#+v", &v4_ptr, "(**"+v4_type+")("+v4_ptr_addr+"->"+v4_addr+")"+v4_v4alue)
	addToPrintDatas("%#+v", v4_nil, "(*"+v4_type+")"+"<nil>")

	v5 := uint(4294967295)
	v5_nil := (*uint)(nil)
	v5_ptr := &v5

	v5_addr := fmt.Sprintf("%p", v5_ptr)
	v5_ptr_addr := fmt.Sprintf("%p", &v5_ptr)

	v5_type := "uint"
	v5_v5alue := "4294967295"

	addToPrintDatas("%v", v5, v5_v5alue)
	addToPrintDatas("%v", v5_ptr, "<*>"+v5_v5alue)
	addToPrintDatas("%v", &v5_ptr, "<**>"+v5_v5alue)
	addToPrintDatas("%v", v5_nil, "<nil>")
	addToPrintDatas("%+v", v5, v5_v5alue)
	addToPrintDatas("%+v", v5_ptr, "<*>("+v5_addr+")"+v5_v5alue)
	addToPrintDatas("%+v", &v5_ptr, "<**>("+v5_ptr_addr+"->"+v5_addr+")"+v5_v5alue)
	addToPrintDatas("%+v", v5_nil, "<nil>")
	addToPrintDatas("%#v", v5, "("+v5_type+")"+v5_v5alue)
	addToPrintDatas("%#v", v5_ptr, "(*"+v5_type+")"+v5_v5alue)
	addToPrintDatas("%#v", &v5_ptr, "(**"+v5_type+")"+v5_v5alue)
	addToPrintDatas("%#v", v5_nil, "(*"+v5_type+")"+"<nil>")
	addToPrintDatas("%#+v", v5, "("+v5_type+")"+v5_v5alue)
	addToPrintDatas("%#+v", v5_ptr, "(*"+v5_type+")("+v5_addr+")"+v5_v5alue)
	addToPrintDatas("%#+v", &v5_ptr, "(**"+v5_type+")("+v5_ptr_addr+"->"+v5_addr+")"+v5_v5alue)
	addToPrintDatas("%#+v", v5_nil, "(*"+v5_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintBool(t *testing.T) {

	v := bool(true)
	v_nil := (*bool)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "bool"
	v_value := "true"

	addToPrintDatas("%v", v, v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)
	addToPrintDatas("%v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, v_value)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	v2 := bool(false)
	v2_nil := (*bool)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "bool"
	v2_v2alue := "false"

	addToPrintDatas("%v", v2, v2_v2alue)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_v2alue)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_v2alue)
	addToPrintDatas("%v", v2_nil, "<nil>")
	addToPrintDatas("%+v", v2, v2_v2alue)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", v2_nil, "<nil>")
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_nil, "(*"+v2_type+")"+"<nil>")
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_nil, "(*"+v2_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintFloat(t *testing.T) {

	v := float32(3.1415)
	v_nil := (*float32)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "float32"
	v_value := "3.1415"

	addToPrintDatas("%v", v, v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)
	addToPrintDatas("%v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, v_value)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	v2 := float64(3.1415926)
	v2_nil := (*float64)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "float64"
	v2_v2alue := "3.1415926"

	addToPrintDatas("%v", v2, v2_v2alue)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_v2alue)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_v2alue)
	addToPrintDatas("%v", v2_nil, "<nil>")
	addToPrintDatas("%+v", v2, v2_v2alue)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", v2_nil, "<nil>")
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_nil, "(*"+v2_type+")"+"<nil>")
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_nil, "(*"+v2_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintComplex(t *testing.T) {

	v := complex(float32(-6), 2)
	v_nil := (*complex64)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "complex64"
	v_value := "(-6+2i)"

	addToPrintDatas("%v", v, v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)
	addToPrintDatas("%v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, v_value)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	v2 := complex(float64(-1), 3)

	v2_nil := (*complex128)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "complex128"
	v2_v2alue := "(-1+3i)"

	addToPrintDatas("%v", v2, v2_v2alue)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_v2alue)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_v2alue)
	addToPrintDatas("%v", v2_nil, "<nil>")
	addToPrintDatas("%+v", v2, v2_v2alue)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%+v", v2_nil, "<nil>")
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_nil, "(*"+v2_type+")"+"<nil>")
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_nil, "(*"+v2_type+")"+"<nil>")

	doPrintTesting(t)
}
func TestPrintString(t *testing.T) {

	//String true
	v := "debuging"
	v_nil := (*string)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "string"
	v_value := "\"debuging\""
	v_value_1 := "(len=8) \"debuging\""

	addToPrintDatas("%v", v, v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, v_value)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+v_value_1)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value_1)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value_1)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value_1)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value_1)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value_1)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	doPrintTesting(t)
}

type pString string

func (s pString) String() string {
	return string(s)
}

type pPtrString string

func (p *pPtrString) String() string {
	return string(*p)
}

func TestPrintArray(t *testing.T) {

	v := [3]int{1, 2, 3}
	v_nil := (*[3]int)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "[3]int"
	v_value := "(len=3 cap=3) [1 2 3 ]"
	v_value_1 := "(len=3 cap=3) [\n(int)1,\n(int)2,\n(int)3,\n]"

	addToPrintDatas("%v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)
	addToPrintDatas("%v", v_nil, "<nil>")

	addToPrintDatas("%+v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")

	addToPrintDatas("%#v", v, "("+v_type+")"+v_value_1)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value_1)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value_1)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")

	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value_1)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value_1)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value_1)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintArray2(t *testing.T) {
	v2 := [3]pString{"1", "2", "3"}

	v2_nil := (*[3]pString)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "[3]debuging.pString"
	v2_v2alue := "(len=3 cap=3) [\n(debuging.pString)(len=1) 1,\n(debuging.pString)(len=1) 2,\n(debuging.pString)(len=1) 3,\n]"
	v2_v2alue_1 := "(len=3 cap=3) [1 2 3 ]"
	if IsDisabledUnsafeReflect {
		v2_v2alue = "(len=3 cap=3) [1 2 3 ]"
	}

	addToPrintDatas("%v", v2, "("+v2_type+")"+v2_v2alue_1)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_v2alue_1)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_v2alue_1)
	addToPrintDatas("%v", v2_nil, "<nil>")

	addToPrintDatas("%+v", v2, "("+v2_type+")"+v2_v2alue_1)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_v2alue_1)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue_1)
	addToPrintDatas("%+v", v2_nil, "<nil>")

	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_nil, "(*"+v2_type+")"+"<nil>")

	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_nil, "(*"+v2_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintArray3(t *testing.T) {

	v3 := [3]interface{}{"one", int(2), uint(3)}

	v3_nil := (*[3]interface{})(nil)
	v3_ptr := &v3

	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_addr := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "[3]interface {}"
	v3_v3alue := "(len=3 cap=3) [\n(string)(len=3) \"one\",\n(int)2,\n(uint)3,\n]"
	v3_v3alue_1 := "(len=3 cap=3) [\"one\" 2 3 ]"
	if IsDisabledUnsafeReflect {
		v3_v3alue = "(len=3 cap=3) [\"one\" 2 3 ]"
		v3_v3alue_1 = "(len=3 cap=3) [\"one\" 2 3 ]"
	}

	addToPrintDatas("%v", v3, "("+v3_type+")"+v3_v3alue_1)
	addToPrintDatas("%v", v3_ptr, "<*>"+v3_v3alue_1)
	addToPrintDatas("%v", &v3_ptr, "<**>"+v3_v3alue_1)
	addToPrintDatas("%v", v3_nil, "<nil>")

	addToPrintDatas("%+v", v3, "("+v3_type+")"+v3_v3alue_1)
	addToPrintDatas("%+v", v3_ptr, "<*>("+v3_addr+")"+v3_v3alue_1)
	addToPrintDatas("%+v", &v3_ptr, "<**>("+v3_ptr_addr+"->"+v3_addr+")"+v3_v3alue_1)
	addToPrintDatas("%+v", v3_nil, "<nil>")

	addToPrintDatas("%#v", v3, "("+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", v3_ptr, "(*"+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", &v3_ptr, "(**"+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", v3_nil, "(*"+v3_type+")"+"<nil>")
	addToPrintDatas("%#+v", v3, "("+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#+v", v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%#+v", &v3_ptr, "(**"+v3_type+")("+v3_ptr_addr+"->"+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%#+v", v3_nil, "(*"+v3_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintArray4(t *testing.T) {
	v4 := [34]byte{
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30,
		0x31, 0x32,
	}

	v4_nil := (*[34]byte)(nil)

	v4_ptr := &v4

	v4_len_str := fmt.Sprintf("%d", len(v4))
	v4_cap_str := fmt.Sprintf("%d", cap(v4))
	v4_addr := fmt.Sprintf("%p", v4_ptr)
	v4_ptr_addr := fmt.Sprintf("%p", &v4_ptr)

	v4_type := "[34]uint8"
	v4_value := "(len=" + v4_len_str + " cap=" + v4_cap_str + ") " +
		"[\n" +
		" 00000000  11 12 13 14 15 16 17 18  19 1a 1b 1c 1d 1e 1f 20  |............... |\n" +
		" 00000010  21 22 23 24 25 26 27 28  29 2a 2b 2c 2d 2e 2f 30  |!\"#$%&'()*+,-./0|\n" +
		" 00000020  31 32                                             |12|\n" +
		"]"

	addToPrintDatas("%v", v4, "("+v4_type+")"+v4_value)
	addToPrintDatas("%v", v4_ptr, "<*>"+v4_value)
	addToPrintDatas("%v", &v4_ptr, "<**>"+v4_value)
	addToPrintDatas("%v", v4_nil, "<nil>")

	addToPrintDatas("%+v", v4, "("+v4_type+")"+v4_value)
	addToPrintDatas("%+v", v4_ptr, "<*>("+v4_addr+")"+v4_value)
	addToPrintDatas("%+v", &v4_ptr, "<**>("+v4_ptr_addr+"->"+v4_addr+")"+v4_value)

	addToPrintDatas("%+v", v4_nil, "<nil>")
	addToPrintDatas("%#v", v4, "("+v4_type+")"+v4_value)
	addToPrintDatas("%#v", v4_ptr, "(*"+v4_type+")"+v4_value)
	addToPrintDatas("%#v", &v4_ptr, "(**"+v4_type+")"+v4_value)
	addToPrintDatas("%#v", v4_nil, "(*"+v4_type+")"+"<nil>")
	addToPrintDatas("%#+v", v4, "("+v4_type+")"+v4_value)
	addToPrintDatas("%#+v", v4_ptr, "(*"+v4_type+")("+v4_addr+")"+v4_value)
	addToPrintDatas("%#+v", &v4_ptr, "(**"+v4_type+")("+v4_ptr_addr+"->"+v4_addr+")"+v4_value)
	addToPrintDatas("%#+v", v4_nil, "(*"+v4_type+")"+"<nil>")
	doPrintTesting(t)
}

func TestPrintSlice(t *testing.T) {

	v := []int{1, 2, 3}
	v_nil := (*[]int)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "[]int"
	v_value := "(len=3 cap=3) [1 2 3 ]"
	v_value_1 := "(len=3 cap=3) [\n(int)1,\n(int)2,\n(int)3,\n]"

	addToPrintDatas("%v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)
	addToPrintDatas("%v", v_nil, "<nil>")

	addToPrintDatas("%+v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")

	addToPrintDatas("%#v", v, "("+v_type+")"+v_value_1)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value_1)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value_1)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")

	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value_1)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value_1)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value_1)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintSlice2(t *testing.T) {
	v2 := []pString{"1", "2", "3"}

	v2_nil := (*[]pString)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "[]debuging.pString"
	v2_v2alue := "(len=3 cap=3) [\n(debuging.pString)(len=1) 1,\n(debuging.pString)(len=1) 2,\n(debuging.pString)(len=1) 3,\n]"
	v2_v2alue_1 := "(len=3 cap=3) [1 2 3 ]"
	if IsDisabledUnsafeReflect {
		v2_v2alue = "(len=3 cap=3) [1 2 3 ]"
	}

	addToPrintDatas("%v", v2, "("+v2_type+")"+v2_v2alue_1)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_v2alue_1)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_v2alue_1)
	addToPrintDatas("%v", v2_nil, "<nil>")

	addToPrintDatas("%+v", v2, "("+v2_type+")"+v2_v2alue_1)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_v2alue_1)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue_1)
	addToPrintDatas("%+v", v2_nil, "<nil>")

	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#v", v2_nil, "(*"+v2_type+")"+"<nil>")

	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_v2alue)
	addToPrintDatas("%#+v", v2_nil, "(*"+v2_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintSlice3(t *testing.T) {

	v3 := []interface{}{"one", int(2), uint(3)}

	v3_nil := (*[]interface{})(nil)
	v3_ptr := &v3

	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_addr := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "[]interface {}"
	v3_v3alue := "(len=3 cap=3) [\n(string)(len=3) \"one\",\n(int)2,\n(uint)3,\n]"
	v3_v3alue_1 := "(len=3 cap=3) [\"one\" 2 3 ]"
	if IsDisabledUnsafeReflect {
		v3_v3alue = "(len=3 cap=3) [\"one\" 2 3 ]"
		v3_v3alue_1 = "(len=3 cap=3) [\"one\" 2 3 ]"
	}

	addToPrintDatas("%v", v3, "("+v3_type+")"+v3_v3alue_1)
	addToPrintDatas("%v", v3_ptr, "<*>"+v3_v3alue_1)
	addToPrintDatas("%v", &v3_ptr, "<**>"+v3_v3alue_1)
	addToPrintDatas("%v", v3_nil, "<nil>")

	addToPrintDatas("%+v", v3, "("+v3_type+")"+v3_v3alue_1)
	addToPrintDatas("%+v", v3_ptr, "<*>("+v3_addr+")"+v3_v3alue_1)
	addToPrintDatas("%+v", &v3_ptr, "<**>("+v3_ptr_addr+"->"+v3_addr+")"+v3_v3alue_1)
	addToPrintDatas("%+v", v3_nil, "<nil>")

	addToPrintDatas("%#v", v3, "("+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", v3_ptr, "(*"+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", &v3_ptr, "(**"+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#v", v3_nil, "(*"+v3_type+")"+"<nil>")
	addToPrintDatas("%#+v", v3, "("+v3_type+")"+v3_v3alue)
	addToPrintDatas("%#+v", v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%#+v", &v3_ptr, "(**"+v3_type+")("+v3_ptr_addr+"->"+v3_addr+")"+v3_v3alue)
	addToPrintDatas("%#+v", v3_nil, "(*"+v3_type+")"+"<nil>")

	doPrintTesting(t)
}

func TestPrintSlice4(t *testing.T) {
	v4 := []byte{
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30,
		0x31, 0x32,
	}

	v4_nil := (*[]byte)(nil)

	v4_ptr := &v4

	v4_len_str := fmt.Sprintf("%d", len(v4))
	v4_cap_str := fmt.Sprintf("%d", cap(v4))
	v4_addr := fmt.Sprintf("%p", v4_ptr)
	v4_ptr_addr := fmt.Sprintf("%p", &v4_ptr)

	v4_type := "[]uint8"
	v4_value := "(len=" + v4_len_str + " cap=" + v4_cap_str + ") " +
		"[\n" +
		" 00000000  11 12 13 14 15 16 17 18  19 1a 1b 1c 1d 1e 1f 20  |............... |\n" +
		" 00000010  21 22 23 24 25 26 27 28  29 2a 2b 2c 2d 2e 2f 30  |!\"#$%&'()*+,-./0|\n" +
		" 00000020  31 32                                             |12|\n" +
		"]"

	addToPrintDatas("%v", v4, "("+v4_type+")"+v4_value)
	addToPrintDatas("%v", v4_ptr, "<*>"+v4_value)
	addToPrintDatas("%v", &v4_ptr, "<**>"+v4_value)
	addToPrintDatas("%v", v4_nil, "<nil>")

	addToPrintDatas("%+v", v4, "("+v4_type+")"+v4_value)
	addToPrintDatas("%+v", v4_ptr, "<*>("+v4_addr+")"+v4_value)
	addToPrintDatas("%+v", &v4_ptr, "<**>("+v4_ptr_addr+"->"+v4_addr+")"+v4_value)

	addToPrintDatas("%+v", v4_nil, "<nil>")
	addToPrintDatas("%#v", v4, "("+v4_type+")"+v4_value)
	addToPrintDatas("%#v", v4_ptr, "(*"+v4_type+")"+v4_value)
	addToPrintDatas("%#v", &v4_ptr, "(**"+v4_type+")"+v4_value)
	addToPrintDatas("%#v", v4_nil, "(*"+v4_type+")"+"<nil>")
	addToPrintDatas("%#+v", v4, "("+v4_type+")"+v4_value)
	addToPrintDatas("%#+v", v4_ptr, "(*"+v4_type+")("+v4_addr+")"+v4_value)
	addToPrintDatas("%#+v", &v4_ptr, "(**"+v4_type+")("+v4_ptr_addr+"->"+v4_addr+")"+v4_value)
	addToPrintDatas("%#+v", v4_nil, "(*"+v4_type+")"+"<nil>")
	doPrintTesting(t)
}

func TestPrintInterface(t *testing.T) {
	// Nil interface.
	var v interface{}
	v_nil := (*interface{})(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "interface {}"
	v_value := "<nil>"

	addToPrintDatas("%v", v, v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, v_value)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	//Sub interface.
	v2 := interface{}(uint16(65535))
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "uint16"
	v2_value := "65535"

	addToPrintDatas("%v", v2, v2_value)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_value)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_value)
	addToPrintDatas("%+v", v2, v2_value)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_value)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_value)
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_value)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_value)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_value)
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_value)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_value)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_value)
	doPrintTesting(t)
}

func TestPrintMap(t *testing.T) {

	v_key_1 := "1st"
	v_key_2 := "2nd"
	v := map[string]int{v_key_1: 1, v_key_2: 2}
	v_nil := map[string]int(nil)
	v_ptr_nil := (*map[string]int)(nil)
	v_ptr := &v

	v_len := fmt.Sprintf("%d", len(v))

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "map[string]int"

	v_value_01 := fmt.Sprintf("map[\"1st\":1 \"2nd\":2]")
	v_value_10 := fmt.Sprintf("map[\"2nd\":2 \"1st\":1]")

	v_value_23 := fmt.Sprintf("(len=%s) {\n\"1st\":1,\n\"2nd\":2,\n}", v_len)
	v_value_32 := fmt.Sprintf("(len=%s) {\n\"2nd\":2,\n\"1st\":1,\n}", v_len)

	addToPrintDatas("%v", v, v_value_01, v_value_10)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value_01, "<*>"+v_value_10)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value_01, "<**>"+v_value_10)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%+v", v_ptr_nil, "<nil>")
	addToPrintDatas("%+v", v, v_value_01, v_value_10)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value_01, "<*>("+v_addr+")"+v_value_10)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value_01, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value_10)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%+v", v_ptr_nil, "<nil>")

	addToPrintDatas("%#v", v, "("+v_type+")"+v_value_32, "("+v_type+")"+v_value_23)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value_32, "(*"+v_type+")"+v_value_23)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value_32, "(**"+v_type+")"+v_value_23)
	addToPrintDatas("%#v", v_nil, "("+v_type+")"+"<nil>")
	addToPrintDatas("%#v", v_ptr_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value_32, "("+v_type+")"+v_value_23)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value_32, "(*"+v_type+")("+v_addr+")"+v_value_23)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value_32, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value_23)
	addToPrintDatas("%#+v", v_nil, "("+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v_ptr_nil, "(*"+v_type+")"+"<nil>")

	v2_key_1 := pPtrString("1st")
	v2_key_2 := pPtrString("2nd")

	v2 := map[pPtrString]int{v2_key_1: 1, v2_key_2: 2}
	v2_nil := map[pPtrString]int(nil)
	v2_ptr_nil := (*map[pPtrString]int)(nil)
	v2_ptr := &v2

	v2_len := fmt.Sprintf("%d", len(v2))

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "map[debuging.pPtrString]int"

	v2_value_01 := fmt.Sprintf("map[1st:1 2nd:2]")
	v2_value_10 := fmt.Sprintf("map[2nd:2 1st:1]")

	v2_value_23 := fmt.Sprintf("(len=%s) {\n1st:1,\n2nd:2,\n}", v2_len)
	v2_value_32 := fmt.Sprintf("(len=%s) {\n2nd:2,\n1st:1,\n}", v2_len)

	addToPrintDatas("%v", v2, v2_value_01, v2_value_10)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_value_01, "<*>"+v2_value_10)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_value_01, "<**>"+v2_value_10)
	addToPrintDatas("%+v", v2_nil, "<nil>")
	addToPrintDatas("%+v", v2, v2_value_01, v2_value_10)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_value_01, "<*>("+v2_addr+")"+v2_value_10)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_value_01, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_value_10)
	addToPrintDatas("%+v", v2_nil, "<nil>")

	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_value_23, "("+v2_type+")"+v2_value_32)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_value_23, "(*"+v2_type+")"+v2_value_32)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_value_23, "(**"+v2_type+")"+v2_value_32)
	addToPrintDatas("%#v", v2_ptr_nil, "(*"+v2_type+")"+"<nil>")
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_value_23, "("+v2_type+")"+v2_value_32)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_value_23, "(*"+v2_type+")("+v2_addr+")"+v2_value_32)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_value_23, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_value_32)
	addToPrintDatas("%#+v", v2_ptr_nil, "(*"+v2_type+")"+"<nil>")

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

	v3_value := fmt.Sprintf("map[\"1st\":1]")

	v3_value2 := fmt.Sprintf("(len=%s) {\n(%s)(len=%s) \"1st\":(%s)1,\n}", v3_len, v3_key_type, v3_key_len, v3_value_type)

	addToPrintDatas("%v", v3, v3_value)
	addToPrintDatas("%v", v3_ptr, "<*>"+v3_value)
	addToPrintDatas("%v", &v3_ptr, "<**>"+v3_value)
	addToPrintDatas("%+v", v3_nil, "<nil>")
	addToPrintDatas("%+v", v3, v3_value)
	addToPrintDatas("%+v", v3_ptr, "<*>("+v3_addr+")"+v3_value)
	addToPrintDatas("%+v", &v3_ptr, "<**>("+v3_ptr_addr+"->"+v3_addr+")"+v3_value)
	addToPrintDatas("%+v", v3_ptr_nil, "<nil>")
	addToPrintDatas("%#v", v3, "("+v3_type+")"+v3_value2)
	addToPrintDatas("%#v", v3_ptr, "(*"+v3_type+")"+v3_value2)
	addToPrintDatas("%#v", &v3_ptr, "(**"+v3_type+")"+v3_value2)
	addToPrintDatas("%#v", v3_ptr_nil, "(*"+v3_type+")"+"<nil>")
	addToPrintDatas("%#+v", v3, "("+v3_type+")"+v3_value2)
	addToPrintDatas("%#+v", v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_value2)
	addToPrintDatas("%#+v", &v3_ptr, "(**"+v3_type+")("+v3_ptr_addr+"->"+v3_addr+")"+v3_value2)
	addToPrintDatas("%#+v", v3_ptr_nil, "(*"+v3_type+")"+"<nil>")

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

	v4_value := fmt.Sprintf("map[\"nil\":<nil>]")
	v4_value2 := fmt.Sprintf("(len=%s) {\n\"nil\":(%s)<nil>,\n}", v4_len, v4_value_type)

	addToPrintDatas("%v", v4, v4_value)
	addToPrintDatas("%v", v4_ptr, "<*>"+v4_value)
	addToPrintDatas("%v", &v4_ptr, "<**>"+v4_value)
	addToPrintDatas("%+v", v4_nil, "<nil>")
	addToPrintDatas("%+v", v4, v4_value)
	addToPrintDatas("%+v", v4_ptr, "<*>("+v4_addr+")"+v4_value)
	addToPrintDatas("%+v", &v4_ptr, "<**>("+v4_ptr_addr+"->"+v4_addr+")"+v4_value)
	addToPrintDatas("%+v", v4_nil, "<nil>")
	addToPrintDatas("%#v", v4, "("+v4_type+")"+v4_value2)
	addToPrintDatas("%#v", v4_ptr, "(*"+v4_type+")"+v4_value2)
	addToPrintDatas("%#v", &v4_ptr, "(**"+v4_type+")"+v4_value2)
	addToPrintDatas("%#v", v4_ptr_nil, "(*"+v4_type+")"+"<nil>")
	addToPrintDatas("%#+v", v4, "("+v4_type+")"+v4_value2)
	addToPrintDatas("%#+v", v4_ptr, "(*"+v4_type+")("+v4_addr+")"+v4_value2)
	addToPrintDatas("%#+v", &v4_ptr, "(**"+v4_type+")("+v4_ptr_addr+"->"+v4_addr+")"+v4_value2)
	addToPrintDatas("%#+v", v4_ptr_nil, "(*"+v4_type+")"+"<nil>")
	doPrintTesting(t)
}

func TestPrintStruct(t *testing.T) {

	//Struct with primitives
	type p1 struct {
		a int8
		b uint8
	}
	v := p1{126, 255}
	v_nil := (*p1)(nil)
	v_ptr := &v

	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)

	v_type := "debuging.p1"
	v_f1_type := "int8"
	v_f2_type := "uint8"

	v_value := fmt.Sprintf("{\n 126,\n 255,\n}")
	v_value_0 := fmt.Sprintf("{\n a:126,\n b:255,\n}")
	v_value_1 := fmt.Sprintf("{\n a:(%s)126,\n b:(%s)255,\n}", v_f1_type, v_f2_type)

	addToPrintDatas("%v", v, v_value)
	addToPrintDatas("%v", v_ptr, "<*>"+v_value)
	addToPrintDatas("%v", &v_ptr, "<**>"+v_value)

	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, v_value_0)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+v_value_0)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+v_value_0)
	addToPrintDatas("%+v", v_nil, "<nil>")

	addToPrintDatas("%#v", v, "("+v_type+")"+v_value_1)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+v_value_1)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+v_value_1)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+v_value_1)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+v_value_1)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+v_value_1)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	//Struct contains another Struct
	type p2 struct {
		A p1
		B bool
	}
	v2 := p2{p1{127, 255}, true}
	v2_nil := (*p2)(nil)
	v2_ptr := &v2

	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)

	v2_type := "debuging.p2"
	v2_f1_type := "debuging.p1"
	v2_f1_f1_type := "int8"
	v2_f1_f2_type := "uint8"
	v2_f2_type := "bool"

	v2_value := fmt.Sprintf("{\n {\n  127,\n  255,\n },\n true,\n}")
	v2_value_0 := fmt.Sprintf("{\n A:{\n  a:127,\n  b:255,\n },\n B:true,\n}")
	v2_value_1 := fmt.Sprintf("{\n A:(%s){\n  a:(%s)127,\n  b:(%s)255,\n },\n B:(%s)true,\n}", v2_f1_type, v2_f1_f1_type, v2_f1_f2_type, v2_f2_type)

	addToPrintDatas("%v", v2, v2_value)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2_value)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2_value)

	addToPrintDatas("%+v", v2_nil, "<nil>")
	addToPrintDatas("%+v", v2, v2_value_0)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2_value_0)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2_value_0)
	addToPrintDatas("%+v", v2_nil, "<nil>")

	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2_value_1)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2_value_1)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2_value_1)
	addToPrintDatas("%#v", v2_nil, "(*"+v2_type+")"+"<nil>")
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2_value_1)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2_value_1)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2_value_1)
	addToPrintDatas("%#+v", v2_nil, "(*"+v2_type+")"+"<nil>")

	doPrintTesting(t)
}
func TestPrintStruct2(t *testing.T) {

	//Struct contains another Struct
	type p3 struct {
		C pPtrString
		D pPtrString
	}
	v3 := p3{("1st"), ("2nd")}
	v3_nil := (*p3)(nil)
	v3_ptr := &v3

	v3_addr := fmt.Sprintf("%p", v3_ptr)
	v3_ptr_addr := fmt.Sprintf("%p", &v3_ptr)

	v3_type := "debuging.p3"
	v3_f1_type := "debuging.pPtrString"
	v3_f2_type := "debuging.pPtrString"
	v3_f1_len := fmt.Sprintf("%d", len(v3.C))
	v3_f2_len := fmt.Sprintf("%d", len(v3.D))

	v3_value := fmt.Sprintf("{\n 1st,\n 2nd,\n}")
	v3_value_0 := fmt.Sprintf("{\n C:1st,\n D:2nd,\n}")
	v3_value_1 := fmt.Sprintf("{\n C:(%s)(len=%s) 1st,\n D:(%s)(len=%s) 2nd,\n}", v3_f1_type, v3_f1_len, v3_f2_type, v3_f2_len)

	if IsDisabledUnsafeReflect {
		v3_value = fmt.Sprintf("{\n 1st,\n 2nd,\n}")
		v3_value_0 = fmt.Sprintf("{\n C: 1st,\n D: 2nd,\n}")
		v3_value_1 = fmt.Sprintf("{\n C:(%s)(len=%s) 1st,\n D:(%s)(len=%s) 2nd,\n}", v3_f1_type, v3_f1_len, v3_f2_type, v3_f2_len)
	}

	addToPrintDatas("%v", v3, v3_value)
	addToPrintDatas("%v", v3_ptr, "<*>"+v3_value)
	addToPrintDatas("%v", &v3_ptr, "<**>"+v3_value)
	addToPrintDatas("%+v", v3_nil, "<nil>")

	addToPrintDatas("%+v", v3, v3_value_0)
	addToPrintDatas("%+v", v3_ptr, "<*>("+v3_addr+")"+v3_value_0)
	addToPrintDatas("%+v", &v3_ptr, "<**>("+v3_ptr_addr+"->"+v3_addr+")"+v3_value_0)
	addToPrintDatas("%+v", v3_nil, "<nil>")

	addToPrintDatas("%#v", v3, "("+v3_type+")"+v3_value_1)
	addToPrintDatas("%#v", v3_ptr, "(*"+v3_type+")"+v3_value_1)
	addToPrintDatas("%#v", &v3_ptr, "(**"+v3_type+")"+v3_value_1)
	addToPrintDatas("%#v", v3_nil, "(*"+v3_type+")"+"<nil>")

	addToPrintDatas("%#+v", v3, "("+v3_type+")"+v3_value_1)
	addToPrintDatas("%#+v", v3_ptr, "(*"+v3_type+")("+v3_addr+")"+v3_value_1)
	addToPrintDatas("%#+v", &v3_ptr, "(**"+v3_type+")("+v3_ptr_addr+"->"+v3_addr+")"+v3_value_1)
	addToPrintDatas("%#+v", v3_nil, "(*"+v3_type+")"+"<nil>")

	type p5 struct {
		a string
	}
	type p6 struct {
		*p5
		e *p5
	}
	// Struct that contains embedded struct and field to same struct.
	e := p5{"p5str"}
	v4 := p6{p5: &e, e: &e}
	v_nil4 := (*p6)(nil)
	v_type4 := &v4
	eAddr := fmt.Sprintf("%p", &e)
	v4Addr := fmt.Sprintf("%p", v_type4)
	v_type4Addr := fmt.Sprintf("%p", &v_type4)
	v4t := "debuging.p6"
	v4t2 := "debuging.p5"
	v4t3 := "string"

	v4s := "{\n <*>{\n  \"p5str\",\n },\n <*>{\n  \"p5str\",\n },\n}"
	v4p2 := "{\n p5:<*>(" + eAddr + "){\n  a:\"p5str\",\n },\n e:<*>(" + eAddr + "){\n  a:\"p5str\",\n },\n}"
	v4p3 := "{\n p5:(*" + v4t2 + "){\n  a:(" + v4t3 + ")(len=5) \"p5str\",\n },\n e:(*" + v4t2 + "){\n  a:(" + v4t3 + ")(len=5) \"p5str\",\n },\n}"
	v4s4 := "{\n p5:(*" + v4t2 + ")(" + eAddr + "){\n  a:(" + v4t3 + ")(len=5) \"p5str\",\n },\n e:(*" + v4t2 + ")(" + eAddr + "){\n  a:(" + v4t3 + ")(len=5) \"p5str\",\n },\n}"

	addToPrintDatas("%v", v4, v4s)
	addToPrintDatas("%v", v_type4, "<*>"+v4s)
	addToPrintDatas("%v", &v_type4, "<**>"+v4s)
	addToPrintDatas("%+v", v_nil4, "<nil>")
	addToPrintDatas("%+v", v4, v4p2)
	addToPrintDatas("%+v", v_type4, "<*>("+v4Addr+")"+v4p2)
	addToPrintDatas("%+v", &v_type4, "<**>("+v_type4Addr+"->"+v4Addr+")"+v4p2)
	addToPrintDatas("%+v", v_nil4, "<nil>")

	addToPrintDatas("%#v", v4, "("+v4t+")"+v4p3)
	addToPrintDatas("%#v", v_type4, "(*"+v4t+")"+v4p3)
	addToPrintDatas("%#v", &v_type4, "(**"+v4t+")"+v4p3)
	addToPrintDatas("%#v", v_nil4, "(*"+v4t+")"+"<nil>")

	addToPrintDatas("%#+v", v4, "("+v4t+")"+v4s4)
	addToPrintDatas("%#+v", v_type4, "(*"+v4t+")("+v4Addr+")"+v4s4)
	addToPrintDatas("%#+v", &v_type4, "(**"+v4t+")("+v_type4Addr+"->"+v4Addr+")"+v4s4)
	addToPrintDatas("%#+v", v_nil4, "(*"+v4t+")"+"<nil>")

	doPrintTesting(t)
}
func TestPrintUintptr(t *testing.T) {

	//Struct contains another Struct

	// Null pointer.
	v := uintptr(0)
	v_nil := (*uintptr)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "uintptr"
	vs := "<nil>"

	addToPrintDatas("%v", v, vs)
	addToPrintDatas("%v", v_ptr, "<*>"+vs)
	addToPrintDatas("%v", &v_ptr, "<**>"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, vs)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+vs)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+vs)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+vs)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+vs)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")
	// Address of real variable.
	i := 1
	v2 := uintptr(unsafe.Pointer(&i))
	//v2_nil := (*uintptr)(nil)
	v2_ptr := &v2
	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)
	v2_type := "uintptr"
	v2s := fmt.Sprintf("%p", &i)

	addToPrintDatas("%v", v2, v2s)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2s)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2s)
	addToPrintDatas("%+v", v2, v2s)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2s)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2s)
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2s)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2s)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2s)
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2s)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2s)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2s)

	doPrintTesting(t)
}

func TestPrintUnsafePointer(t *testing.T) {

	// Null pointer.
	v := unsafe.Pointer(uintptr(0))
	v_nil := (*unsafe.Pointer)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "unsafe.Pointer"
	vs := "<nil>"

	addToPrintDatas("%v", v, vs)
	addToPrintDatas("%v", v_ptr, "<*>"+vs)
	addToPrintDatas("%v", &v_ptr, "<**>"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, vs)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+vs)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+vs)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+vs)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+vs)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	// Address of real variable.
	i := 1
	v2 := unsafe.Pointer(&i)
	//v2_nil := (*unsafe.Pointer)(nil)
	v2_ptr := &v2
	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)
	v2_type := "unsafe.Pointer"
	v2s := fmt.Sprintf("%p", &i)

	addToPrintDatas("%v", v2, v2s)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2s)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2s)
	addToPrintDatas("%+v", v2, v2s)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2s)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2s)
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2s)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2s)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2s)
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2s)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2s)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2s)

	doPrintTesting(t)
}

func TestPrintChan(t *testing.T) {

	// Null pointer.
	var v chan int
	v_nil := (*chan int)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "chan int"
	vs := "<nil>"
	addToPrintDatas("%v", v, vs)
	addToPrintDatas("%v", v_ptr, "<*>"+vs)
	addToPrintDatas("%v", &v_ptr, "<**>"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, vs)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+vs)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+vs)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+vs)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+vs)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	// Address of real variable.
	v2 := make(chan int)
	//v2_nil := (*chan int)(nil)
	v2_ptr := &v2
	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)
	v2_type := "chan int"
	v2s := fmt.Sprintf("%p", v2)

	addToPrintDatas("%v", v2, v2s)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2s)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2s)
	addToPrintDatas("%+v", v2, v2s)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2s)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2s)
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2s)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2s)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2s)
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2s)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2s)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2s)

	doPrintTesting(t)
}
func testPrintFunc0() {
	fmt.Print("GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG")
}
func TestPrintFunc(t *testing.T) {
	// Function with no params and no returns.
	v := testPrintFunc0
	v_nil := (*func())(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "func()"
	vs := fmt.Sprintf("%p", v)
	addToPrintDatas("%v", v, vs)
	addToPrintDatas("%v", v_ptr, "<*>"+vs)
	addToPrintDatas("%v", &v_ptr, "<**>"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, vs)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+vs)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+vs)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+vs)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+vs)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")

	// Function with param and no returns.
	v2 := TestPrintFunc
	nv2 := (*func(*testing.T))(nil)
	v2_ptr := &v2
	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)
	v2_type := "func(*testing.T)"
	v2s := fmt.Sprintf("%p", v2)
	addToPrintDatas("%v", v2, v2s)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2s)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2s)
	addToPrintDatas("%+v", nv2, "<nil>")
	addToPrintDatas("%+v", v2, v2s)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2s)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2s)
	addToPrintDatas("%+v", nv2, "<nil>")
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2s)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2s)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2s)
	addToPrintDatas("%#v", nv2, "(*"+v2_type+")"+"<nil>")
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2s)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2s)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2s)
	addToPrintDatas("%#+v", nv2, "(*"+v2_type+")"+"<nil>")

	// Function with multiple params and multiple returns.
	var v3 = func(i int, s string) (b bool, err error) {
		return true, nil
	}
	nv3 := (*func(int, string) (bool, error))(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "func(int, string) (bool, error)"
	v3s := fmt.Sprintf("%p", v3)
	addToPrintDatas("%v", v3, v3s)
	addToPrintDatas("%v", pv3, "<*>"+v3s)
	addToPrintDatas("%v", &pv3, "<**>"+v3s)
	addToPrintDatas("%+v", nv3, "<nil>")
	addToPrintDatas("%+v", v3, v3s)
	addToPrintDatas("%+v", pv3, "<*>("+v3Addr+")"+v3s)
	addToPrintDatas("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s)
	addToPrintDatas("%+v", nv3, "<nil>")
	addToPrintDatas("%#v", v3, "("+v3t+")"+v3s)
	addToPrintDatas("%#v", pv3, "(*"+v3t+")"+v3s)
	addToPrintDatas("%#v", &pv3, "(**"+v3t+")"+v3s)
	addToPrintDatas("%#v", nv3, "(*"+v3t+")"+"<nil>")
	addToPrintDatas("%#+v", v3, "("+v3t+")"+v3s)
	addToPrintDatas("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s)
	addToPrintDatas("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s)
	addToPrintDatas("%#+v", nv3, "(*"+v3t+")"+"<nil>")
}

type pRefRef0 struct {
	ref1 *pRefRef1
}
type pRefRef1 struct {
	ref0 *pRefRef0
}

type pIndirCir1 struct {
	ref2 *pIndirCir2
}
type pIndirCir2 struct {
	ref3 *pIndirCir3
}
type pIndirCir3 struct {
	ref1 *pIndirCir1
}

func TestPrintCircular(t *testing.T) {
	// Struct that is circular through self referencing.
	type circular struct {
		c *circular
	}
	v := circular{nil}
	v.c = &v
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "debuging.circular"
	vs := "{<*>{<*><shown>}}"
	vs2 := "{<*><shown>}"
	vs3 := "{c:<*>(" + v_addr + "){c:<*>(" + v_addr + ")<shown>}}"
	vs4 := "{c:<*>(" + v_addr + ")<shown>}"
	vs5 := "{c:(*" + v_type + "){c:(*" + v_type + ")<shown>}}"
	vs6 := "{c:(*" + v_type + ")<shown>}"
	vs7 := "{c:(*" + v_type + ")(" + v_addr + "){c:(*" + v_type + ")(" + v_addr + ")<shown>}}"
	vs8 := "{c:(*" + v_type + ")(" + v_addr + ")<shown>}"
	addToPrintDatas("%v", v, vs)
	addToPrintDatas("%v", v_ptr, "<*>"+vs2)
	addToPrintDatas("%v", &v_ptr, "<**>"+vs2)
	addToPrintDatas("%+v", v, vs3)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+vs4)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+vs4)
	addToPrintDatas("%#v", v, "("+v_type+")"+vs5)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+vs6)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+vs6)
	addToPrintDatas("%#+v", v, "("+v_type+")"+vs7)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+vs8)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+vs8)

	// Structs that are circular through cross referencing.
	v2 := pRefRef0{nil}
	ts2 := pRefRef1{&v2}
	v2.ref1 = &ts2
	v2_ptr := &v2
	ts2Addr := fmt.Sprintf("%p", &ts2)
	v2_addr := fmt.Sprintf("%p", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%p", &v2_ptr)
	v2_type := "debuging.xref1"
	v2t2 := "debuging.xref2"
	v2s := "{<*>{<*>{<*><shown>}}}"
	v2s2 := "{<*>{<*><shown>}}"
	v2s3 := "{ps2:<*>(" + ts2Addr + "){ps1:<*>(" + v2_addr + "){ps2:<*>(" + ts2Addr + ")<shown>}}}"
	v2s4 := "{ps2:<*>(" + ts2Addr + "){ps1:<*>(" + v2_addr + ")<shown>}}"
	v2s5 := "{ps2:(*" + v2t2 + "){ps1:(*" + v2_type + "){ps2:(*" + v2t2 + ")<shown>}}}"
	v2s6 := "{ps2:(*" + v2t2 + "){ps1:(*" + v2_type + ")<shown>}}"
	v2s7 := "{ps2:(*" + v2t2 + ")(" + ts2Addr + "){ps1:(*" + v2_type + ")(" + v2_addr + "){ps2:(*" + v2t2 + ")(" + ts2Addr + ")<shown>}}}"
	v2s8 := "{ps2:(*" + v2t2 + ")(" + ts2Addr + "){ps1:(*" + v2_type + ")(" + v2_addr + ")<shown>}}"
	addToPrintDatas("%v", v2, v2s)
	addToPrintDatas("%v", v2_ptr, "<*>"+v2s2)
	addToPrintDatas("%v", &v2_ptr, "<**>"+v2s2)
	addToPrintDatas("%+v", v2, v2s3)
	addToPrintDatas("%+v", v2_ptr, "<*>("+v2_addr+")"+v2s4)
	addToPrintDatas("%+v", &v2_ptr, "<**>("+v2_ptr_addr+"->"+v2_addr+")"+v2s4)
	addToPrintDatas("%#v", v2, "("+v2_type+")"+v2s5)
	addToPrintDatas("%#v", v2_ptr, "(*"+v2_type+")"+v2s6)
	addToPrintDatas("%#v", &v2_ptr, "(**"+v2_type+")"+v2s6)
	addToPrintDatas("%#+v", v2, "("+v2_type+")"+v2s7)
	addToPrintDatas("%#+v", v2_ptr, "(*"+v2_type+")("+v2_addr+")"+v2s8)
	addToPrintDatas("%#+v", &v2_ptr, "(**"+v2_type+")("+v2_ptr_addr+"->"+v2_addr+")"+v2s8)

	// Structs that are indirectly circular.
	v3 := pIndirCir1{nil}
	tic2 := pIndirCir2{nil}
	tic3 := pIndirCir3{&v3}
	tic2.ref3 = &tic3
	v3.ref2 = &tic2
	pv3 := &v3
	tic2Addr := fmt.Sprintf("%p", &tic2)
	tic3Addr := fmt.Sprintf("%p", &tic3)
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "debuging.pIndirCir1"
	v3t2 := "debuging.pIndirCir2"
	v3t3 := "debuging.pIndirCir3"
	v3s := "{<*>{<*>{<*>{<*><shown>}}}}"
	v3s2 := "{<*>{<*>{<*><shown>}}}"
	v3s3 := "{ps2:<*>(" + tic2Addr + "){ps3:<*>(" + tic3Addr + "){ps1:<*>(" + v3Addr + "){ps2:<*>(" + tic2Addr + ")<shown>}}}}"
	v3s4 := "{ps2:<*>(" + tic2Addr + "){ps3:<*>(" + tic3Addr + "){ps1:<*>(" + v3Addr + ")<shown>}}}"
	v3s5 := "{ps2:(*" + v3t2 + "){ps3:(*" + v3t3 + "){ps1:(*" + v3t + "){ps2:(*" + v3t2 + ")<shown>}}}}"
	v3s6 := "{ps2:(*" + v3t2 + "){ps3:(*" + v3t3 + "){ps1:(*" + v3t + ")<shown>}}}"
	v3s7 := "{ps2:(*" + v3t2 + ")(" + tic2Addr + "){ps3:(*" + v3t3 + ")(" + tic3Addr + "){ps1:(*" + v3t + ")(" + v3Addr + "){ps2:(*" + v3t2 + ")(" + tic2Addr + ")<shown>}}}}"
	v3s8 := "{ps2:(*" + v3t2 + ")(" + tic2Addr + "){ps3:(*" + v3t3 + ")(" + tic3Addr + "){ps1:(*" + v3t + ")(" + v3Addr + ")<shown>}}}"

	addToPrintDatas("%v", v3, v3s)
	addToPrintDatas("%v", pv3, "<*>"+v3s2)
	addToPrintDatas("%v", &pv3, "<**>"+v3s2)
	addToPrintDatas("%+v", v3, v3s3)
	addToPrintDatas("%+v", pv3, "<*>("+v3Addr+")"+v3s4)
	addToPrintDatas("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s4)
	addToPrintDatas("%#v", v3, "("+v3t+")"+v3s5)
	addToPrintDatas("%#v", pv3, "(*"+v3t+")"+v3s6)
	addToPrintDatas("%#v", &pv3, "(**"+v3t+")"+v3s6)
	addToPrintDatas("%#+v", v3, "("+v3t+")"+v3s7)
	addToPrintDatas("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s8)
	addToPrintDatas("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s8)
}

type pPanicer int

func (p pPanicer) String() string {
	panic("pPanicer")
}
func TestPrintPanic(t *testing.T) {
	// Type that panics in its Stringer interface.
	v := pPanicer(127)
	v_nil := (*pPanicer)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "debuging.pPanicer"
	vs := "(PANIC=pPanicer)127"
	addToPrintDatas("%v", v, vs)
	addToPrintDatas("%v", v_ptr, "<*>"+vs)
	addToPrintDatas("%v", &v_ptr, "<**>"+vs)
	addToPrintDatas("%v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, vs)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+vs)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+vs)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+vs)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+vs)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")
}

type pError int

func (e pError) Error() string {
	return fmt.Sprintf("error: %d", int(e))
}
func TestPrintError(t *testing.T) {
	// Type that has a custom Error interface.
	v := pError(127)
	v_nil := (*pError)(nil)
	v_ptr := &v
	v_addr := fmt.Sprintf("%p", v_ptr)
	v_ptr_addr := fmt.Sprintf("%p", &v_ptr)
	v_type := "debuging.pError"
	vs := "error: 127"
	addToPrintDatas("%v", v, vs)
	addToPrintDatas("%v", v_ptr, "<*>"+vs)
	addToPrintDatas("%v", &v_ptr, "<**>"+vs)
	addToPrintDatas("%v", v_nil, "<nil>")
	addToPrintDatas("%+v", v, vs)
	addToPrintDatas("%+v", v_ptr, "<*>("+v_addr+")"+vs)
	addToPrintDatas("%+v", &v_ptr, "<**>("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%+v", v_nil, "<nil>")
	addToPrintDatas("%#v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#v", v_ptr, "(*"+v_type+")"+vs)
	addToPrintDatas("%#v", &v_ptr, "(**"+v_type+")"+vs)
	addToPrintDatas("%#v", v_nil, "(*"+v_type+")"+"<nil>")
	addToPrintDatas("%#+v", v, "("+v_type+")"+vs)
	addToPrintDatas("%#+v", v_ptr, "(*"+v_type+")("+v_addr+")"+vs)
	addToPrintDatas("%#+v", &v_ptr, "(**"+v_type+")("+v_ptr_addr+"->"+v_addr+")"+vs)
	addToPrintDatas("%#+v", v_nil, "(*"+v_type+")"+"<nil>")
}

func TestPrintPassthrough(t *testing.T) {
	// %x passthrough with uint.
	v := uint(4294967295)
	v_ptr := &v
	v_addr := fmt.Sprintf("%x", v_ptr)
	v_ptr_addr := fmt.Sprintf("%x", &v_ptr)
	vs := "ffffffff"
	addToPrintDatas("%x", v, vs)
	addToPrintDatas("%x", v_ptr, v_addr)
	addToPrintDatas("%x", &v_ptr, v_ptr_addr)

	// %#x passthrough with uint.
	v2 := int(2147483647)
	v2_ptr := &v2
	v2_addr := fmt.Sprintf("%#x", v2_ptr)
	v2_ptr_addr := fmt.Sprintf("%#x", &v2_ptr)
	v2s := "0x7fffffff"
	addToPrintDatas("%#x", v2, v2s)
	addToPrintDatas("%#x", v2_ptr, v2_addr)
	addToPrintDatas("%#x", &v2_ptr, v2_ptr_addr)

	// %f passthrough with precision.
	addToPrintDatas("%.2f", 3.1415, "3.14")
	addToPrintDatas("%.3f", 3.1415, "3.142")
	addToPrintDatas("%.4f", 3.1415, "3.1415")

	// %f passthrough with width and precision.
	addToPrintDatas("%5.2f", 3.1415, " 3.14")
	addToPrintDatas("%6.3f", 3.1415, " 3.142")
	addToPrintDatas("%7.4f", 3.1415, " 3.1415")

	// %d passthrough with width.
	addToPrintDatas("%3d", 127, "127")
	addToPrintDatas("%4d", 127, " 127")
	addToPrintDatas("%5d", 127, "  127")

	// %q passthrough with string.
	addToPrintDatas("%q", "test", "\"test\"")
}

type testPrintStruct struct {
	x int
}

func (ts testPrintStruct) String() string {
	return fmt.Sprintf("ts.%d", ts.x)
}

type testPrintStructP struct {
	x int
}

func (ts *testPrintStructP) String() string {
	return fmt.Sprintf("ts.%d", ts.x)
}
func TestPrintSortedKeys(t *testing.T) {

	in := map[int]string{1: "1", 3: "3", 2: "2"}
	want := "map[1:\"1\" 2:\"2\" 3:\"3\"]"
	cfg := DebugConfig{SortKeys: true}
	got := cfg.Sprint(in)
	if got != want {
		t.Errorf("Reverse(%v) want:\n%v \ngot:\n%v,", in, want, got)
	}

}
func TestPrintSortedKeys1(t *testing.T) {

	for _, c := range []struct {
		in   map[pString]int
		want string
	}{
		{map[pString]int{"1": 1, "3": 3, "2": 2},
			"map[1:1 2:2 3:3]",
		},
	} {
		cfg := DebugConfig{SortKeys: true}
		got := cfg.Sprint(c.in)
		if got != c.want {
			t.Errorf("Reverse(%v) want:\n%v got:\n%v,", c.in, c.want, got)
		}
	}
}

func TestPrintSortedKeys2(t *testing.T) {

	in := map[pPtrString]int{pPtrString("1"): 1, pPtrString("3"): 3, pPtrString("2"): 2}
	want := "map[1:1 2:2 3:3]"
	if !IsDisabledUnsafeReflect {
		want = "map[1:1 2:2 3:3]"
	}
	cfg := DebugConfig{SortKeys: true}
	got := cfg.Sprint(in)
	if got != want {
		t.Errorf("Reverse(%v) want:\n%v got:\n%v,", in, want, got)
	}
}

func TestPrintSortedKeys3(t *testing.T) {

	in := map[pError]int{pError(1): 1, pError(3): 3, pError(2): 2}
	want := "map[error: 1:1 error: 2:2 error: 3:3]"
	if !IsDisabledUnsafeReflect {
		want = "map[error: 1:1 error: 2:2 error: 3:3]"
	}
	cfg := DebugConfig{SortKeys: true}
	got := cfg.Sprint(in)
	if got != want {
		t.Errorf("Reverse(%v) want:\n%v got:\n%v,", in, want, got)
	}
}

func doPrintTesting(t *testing.T) {

	for i, d := range printDatas {
		buf := new(bytes.Buffer)

		FmtFPrintf(buf, d.format, d.in)
		out := buf.String()

		if !testInPrintOn(out, d.on) {
			t.Errorf("Formatter Error  #%d \n ned:[%v] \n got:[%s] ", i, d.onString(), out)
		}
	}
}

func testInPrintOn(result string, wants []string) bool {
	for _, item := range wants {
		if result == item {
			return true
		}
	}
	return false
}
