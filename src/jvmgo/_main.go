package main

//
//import (
//	"fmt"
//	"jvmgo/classpath"
//	"strings"
//)
//
//func startJVM(cmd *Cmd) {
//	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
//	fmt.Printf("classpath:%s class:%s args:%s\n", cp, cmd.class, cmd.args)
//
//	className := strings.Replace(cmd.class, ".", "/", -1)
//	classData, _, err := cp.ReadClass(className)
//	if err != nil {
//		fmt.Printf("Could not find or load main class %s\n", cmd.class)
//	}
//	fmt.Printf("class data : %v\n", classData)
//}
//
//func main() {
//	cmd := parseCmd()
//	if cmd.versionFlag {
//		fmt.Println("version 0.0.1")
//	} else if cmd.helpFlag || cmd.class == "" {
//		printUsage()
//	} else {
//		startJVM(cmd)
//	}
//}
