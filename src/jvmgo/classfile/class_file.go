package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
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

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0XCAFEBABE {
		panic("java.lang.ClassFormatError : magic!")
	}
}
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {

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

func (self *ClassFile) InterfaceName() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
