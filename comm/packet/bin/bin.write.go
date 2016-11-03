package packet

import "math"

func (s *BinStream) WriteZeros(n int) {
	for i := 0; i < n; i++ {
		s.data = append(s.data, byte(0))
	}
}

func (s *BinStream) WriteBool(v bool) {
	if v {
		s.data = append(s.data, byte(1))
	} else {
		s.data = append(s.data, byte(0))
	}
}

func (s *BinStream) WriteByte(v byte) {
	s.data = append(s.data, v)
}

func (s *BinStream) WriteBytes(v []byte) {
	s.WriteUint16(uint16(len(v)))
	s.data = append(s.data, v...)
}

func (s *BinStream) WriteRawBytes(v []byte) {
	s.data = append(s.data, v...)
}

func (s *BinStream) WriteString(v string) {
	bytes := []byte(v)
	s.WriteUint16(uint16(len(bytes)))
	s.data = append(s.data, bytes...)
}

func (s *BinStream) WriteInt8(v int8) {
	s.WriteByte(byte(v))
}
func (s *BinStream) WriteUInt8(v uint8) {
	s.WriteByte(byte(v))
}

func (s *BinStream) WriteUint16(v uint16) {
	s.data = append(s.data, byte(v>>8), byte(v))
}

func (s *BinStream) WriteInt16(v int16) {
	s.WriteUint16(uint16(v))
}

func (s *BinStream) WriteUint24(v uint32) {
	s.data = append(s.data, byte(v>>16), byte(v>>8), byte(v))
}

func (s *BinStream) WriteUint32(v uint32) {
	s.data = append(s.data, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

func (s *BinStream) WriteInt32(v int32) {
	s.WriteUint32(uint32(v))
}

func (s *BinStream) WriteUint64(v uint64) {
	s.data = append(s.data, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

func (s *BinStream) WriteInt64(v int64) {
	s.WriteUint64(uint64(v))
}

func (s *BinStream) WriteFloat32(f float32) {
	v := math.Float32bits(f)
	s.WriteUint32(v)
}

func (s *BinStream) WriteFloat64(f float64) {
	v := math.Float64bits(f)
	s.WriteUint64(v)
}
