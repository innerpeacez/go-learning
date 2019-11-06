package classfile

// LineNumberTable属性存放方法的行号信息
type LineNumberTableAttribute struct {
	LineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	LineNumberTableLength := reader.readUint16()
	self.LineNumberTable = make([]*LineNumberTableEntry, LineNumberTableLength)
	for i := range self.LineNumberTable {
		self.LineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
