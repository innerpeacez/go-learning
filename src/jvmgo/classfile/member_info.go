package classfile

type MemberInfo struct {
	cp              ConstantPool    // 常量池
	accessFlags     uint16          // 字段或者方法的访问标志
	nameIndex       uint16          // 字段名或者方法名
	descriptorIndex uint16          // 字段或者方法的描述符
	attributes      []AttributeInfo // 方法的参数（属性）表
}

// 获取字段表或者方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 从常量池字段或者方法的具体数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp), // 获取属性表
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

// 从常量池中查找字段或者方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// 从常量池中查找字段或者方法的描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
