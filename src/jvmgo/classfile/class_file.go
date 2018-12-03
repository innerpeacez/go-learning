package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32          // 魔数
	minorVersion uint16          // 次版本号
	majorVersion uint16          // 主版本号
	constantPool ConstantPool    // 常量池
	accessFlags  uint16          // 类访问标志（判断类或者接口，public，private等）
	thisClass    uint16          // 类名（对应常量池索引）
	superClass   uint16          // 超类名（对应常量池索引）
	interfaces   []uint16        // 类实现的接口表（对应常量池索引）
	fields       []*MemberInfo   // 字段
	methods      []*MemberInfo   // 方法
	attributes   []AttributeInfo //
}

/* 把[]byte解析成ClassFile结构体 */
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/* 一起调用其他方法解析class */
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// 读取class并检查class文件的魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0XCAFEBABE {
		panic("java.lang.ClassFormatError : magic!")
	}
}

// 读取class版本号，并检查版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
func (self *ClassFile) MinorVersion() uint16 { //getter
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 { //getter
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool { //getter
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 { //getter
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo { //getter
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo { //getter
	return self.methods
}

// 从常量池中查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // java中只有java.lang.Object类没有超类
}

// 接口名
func (self *ClassFile) InterfaceName() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
