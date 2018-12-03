package classfile

type ConstantPool []ConstantInfo

// 获取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantPool, cpCount)

	for i := 1; i < cpCount; i++ { // 注意索引从1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //占两个位置
		}
	}
	return cp
}

// 根据索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// 根据索引查找字段或者方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	descriptor := self.getUtf8(ntInfo.descriptorIndex)
	return name, descriptor
}

// 根据索引从常量池中查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getClassName(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

// 从常量池中查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
