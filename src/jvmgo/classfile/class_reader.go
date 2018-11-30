package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// uint8 类似java中无符号 byte，1个字节
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	// 跳过已经读取过的数据
	self.data = self.data[1:]
	return val
}

// uint16 类似java中无符号 char，2个字节
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// uint32 类似java中无符号 int，4个字节
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// uint64 类似java中无符号 long，8个字节
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

// 读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
