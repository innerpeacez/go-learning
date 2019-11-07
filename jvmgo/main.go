package main

import (
	"fmt"
	"github.com/innerpeacez/go-learning/jvmgo/classfile"
	"github.com/innerpeacez/go-learning/jvmgo/classpath"
	"strings"
)

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//fmt.Printf("classpath:%s class:%s args:%s\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	//classData, _, err := cp.ReadClass(className)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
	//if err != nil {
	//	fmt.Printf("Could not find or load main class %s\n", cmd.class)
	//}
	//fmt.Printf("class data : %v\n", classData)
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count : %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags : %v\n", cf.AccessFlags())
	fmt.Printf("this count : %v\n", cf.ClassName())
	fmt.Printf("super count : %v\n", cf.SuperClassName())
	fmt.Printf("interface count : %v\n", cf.InterfaceNames())
	fmt.Printf("fields count : %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("   %s\n", f.Name())
	}
	fmt.Printf("methods count : %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("   %s\n", m.Name())

	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
