package packet

import (
	"reflect"
	"server/core/debuging"
)

//BinStream is BinStream
type BinStream struct {
	pos  int
	data []byte
}

func ReaderBinStream(data []byte) *BinStream {
	return &BinStream{data: data}
}

func WriterBinStream() *BinStream {
	return &BinStream{data: make([]byte, 0, 512)}
}

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////
func (s *BinStream) Data() []byte {
	return s.data
}
func (s *BinStream) lenght() int {
	if s.data == nil {
		return 0
	}
	return len(s.data)
}
func (s *BinStream) Packet(v reflect.Value) {

}

func (s *BinStream) packing(v *reflect.Value) {
	if v.IsNil() {
		return
	}
	switch v.Kind() {
	case reflect.Bool:
		s.WriteBool(v.Bool())
	case reflect.Int8:
		s.WriteInt8(int8(v.Int()))
	case reflect.Int16:
		s.WriteInt16(int16(v.Int()))
	case reflect.Int32:
		s.WriteInt32(int32(v.Int()))
	case reflect.Int64:
		s.WriteInt64(v.Int())
	case reflect.Uint8:
		s.WriteUInt8(uint8(v.Uint()))
	case reflect.Uint16:
		s.WriteUint16(uint16(v.Uint()))
	case reflect.Uint32:
		s.WriteUint32(uint32(v.Uint()))
	case reflect.Uint64:
		s.WriteUint64(v.Uint())
	case reflect.Float32:
		s.WriteFloat32(float32(v.Float()))
	case reflect.Float64:
		s.WriteFloat64(float64(v.Float()))
	case reflect.String:
		s.WriteString(v.String())
	case reflect.Array, reflect.slice:
		s.packSlice(v)
	case reflect.Interface, reflect.Ptr:
		s.packing(v.Elem())
	case reflect.Struct:
		s.packStruct(v)
	default:
		debuging.Tracef("BinStream is not support type :", v)
	}
}

func (s *BinStream) packSlice(v *reflect.Value) {
	if bs, ok := v.Interface().([]byte); ok { // special treat for []bytes
		writer.WriteBytes(bs)
	} else {
		len := v.Len()
		s.WriteUint16(uint16(len))
		for i := 0; i < len; i++ {
			s.packing(v.Index(i))
		}
	}
}

func (s *BinStream) packStruct(v *reflect.Value) {
	len := v.NumField()
	for index := 0; index < len; index++ {
		s.packing(v.Field(index))
	}
}
