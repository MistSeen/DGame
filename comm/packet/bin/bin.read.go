package packet

import (
	"errors"
	"math"
)

func (s *BinStream) ReadByte() (ret byte, err error) {
	if s.pos >= len(s.data) {
		err = errors.New("read byte failed")
		return
	}

	ret = s.data[s.pos]
	s.pos++
	return
}

func (s *BinStream) ReadBool() (ret bool, err error) {
	b, _err := s.ReadByte()
	if b != byte(1) {
		return false, _err
	}

	return true, _err
}

func (s *BinStream) ReadBytes() (ret []byte, err error) {
	dataLen := len(s.data)
	if s.pos+2 > dataLen {
		err = errors.New("read bytes failed.not data")
		return
	}
	size := 0
	if size = s.readDataSize(); size < 0 {
		err = errors.New("read bytes failed.date size error")
		return
	}

	if s.pos+size > dataLen {
		err = errors.New("read bytes faile. to long")
		return
	}

	ret = s.data[s.pos : s.pos+size]
	s.pos += size

	return
}

func (s *BinStream) ReadString() (ret string, err error) {
	dataLen := len(s.data)
	if s.pos+2 > dataLen {
		err = errors.New("read bytes failed.not data")
		return
	}
	size := 0
	if size = s.readDataSize(); size < 0 {
		err = errors.New("read bytes failed.date size error")
		return
	}

	if s.pos+size > dataLen {
		err = errors.New("read bytes faile. to long")
		return
	}

	bytes := s.data[s.pos : s.pos+size]
	s.pos += size
	ret = string(bytes)
	return
}

func (s *BinStream) ReadUint8() (ret uint8, err error) {
	_ret, _err := s.ReadByte()
	ret = uint8(_ret)
	err = _err
	return
}
func (s *BinStream) ReadInt8() (ret int8, err error) {
	_ret, _err := s.ReadByte()
	ret = int8(_ret)
	err = _err
	return
}

func (s *BinStream) ReadUint16() (ret uint16, err error) {
	dataLen := len(s.data)
	if s.pos+2 > dataLen {
		err = errors.New("read uint 16 failed")
		return
	}

	buf := s.data[s.pos : s.pos+2]
	ret = uint16(buf[0])<<8 | uint16(buf[1])

	s.pos += 2

	return
}

func (b *BinStream) ReadInt16() (ret int16, err error) {
	_ret, _err := b.ReadUint16()

	ret = int16(_ret)
	err = _err
	return
}

func (s *BinStream) ReadUint24() (ret uint32, err error) {
	dataLen := len(s.data)
	if s.pos+3 > dataLen {
		err = errors.New("read uint 16 failed")
		return
	}

	buf := s.data[s.pos : s.pos+3]
	ret = uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])

	s.pos += 3

	return
}

func (b *BinStream) ReadInt24() (ret int32, err error) {
	_ret, _err := b.ReadUint24()

	ret = int32(_ret)
	err = _err
	return
}

func (s *BinStream) ReadUint32() (ret uint32, err error) {
	dataLen := len(s.data)
	if s.pos+4 > dataLen {
		err = errors.New("read uint 16 failed")
		return
	}

	buf := s.data[s.pos : s.pos+4]
	ret = uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])

	s.pos += 4

	return
}

func (b *BinStream) ReadInt32() (ret int32, err error) {
	_ret, _err := b.ReadUint32()

	ret = int32(_ret)
	err = _err
	return
}

func (s *BinStream) ReadUint64() (ret uint64, err error) {
	dataLen := len(s.data)
	if s.pos+8 > dataLen {
		err = errors.New("read uint 16 failed")
		return
	}

	ret = 0
	buf := s.data[s.pos : s.pos+8]
	for i, v := range buf {
		ret |= uint64(v) << uint((7-i)*8)
	}

	s.pos += 8

	return
}

func (b *BinStream) ReadInt64() (ret int64, err error) {
	_ret, _err := b.ReadUint64()

	ret = int64(_ret)
	err = _err
	return
}

func (b *BinStream) ReadFloat32() (ret float32, err error) {
	bits, _err := b.ReadUint32()
	if _err != nil {
		return float32(0), _err
	}

	ret = math.Float32frombits(bits)
	if math.IsNaN(float64(ret)) || math.IsInf(float64(ret), 0) {
		return 0, nil
	}

	return ret, nil
}

func (b *BinStream) ReadFloat64() (ret float64, err error) {
	bits, _err := b.ReadUint64()
	if _err != nil {
		return float64(0), _err
	}

	ret = math.Float64frombits(bits)
	if math.IsNaN(ret) || math.IsInf(ret, 0) {
		return 0, nil
	}

	return ret, nil
}
func (s *BinStream) readDataSize() (ret int) {
	dataLen := len(s.data)
	//这里不是大于等于
	//因为data[a,b]操作中,取值范围是[a,b), b的值是可以取到len值
	//而 data[b] 中的值为[0~(len-1)] 属于索引值, 因此b的值只能取到len-1
	if s.pos+2 > dataLen {
		return -1
	}
	buf := s.data[s.pos : s.pos+2]
	ret = int(buf[0])<<8 | int(buf[1])
	s.pos += 2
	return
}
