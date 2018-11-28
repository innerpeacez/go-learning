package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string // 用于存放目录的绝对路径
}

func newDirEntry(path string) *DirEntry {
	// 将path参数转换成据对路径
	absDir, err := filepath.Abs(path)
	// 如果转换出错，调用panic函数终止程序
	if err != nil {
		panic(err)
	}
	// 转换成功返回DirEntry
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 将目录和class文件名拼接成一个完成的路径
	fileName := filepath.Join(self.absDir, className)
	// 调用ioutil包提供的读取文件内容的方法读取文件内容
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
