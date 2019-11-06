package classfile

// LocalVariableTable属性存放方法中的局部变量
type LocalVariableTableAttribute struct {
	LocalVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	LocalVariableTableLength := reader.readUint16()
	self.LocalVariableTable = make([]*LocalVariableTableEntry, LocalVariableTableLength)
	for i := range self.LocalVariableTable {
		self.LocalVariableTable[i] = &LocalVariableTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
