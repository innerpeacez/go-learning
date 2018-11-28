package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string // 用于存放目录的绝对路径
}

func newZipEntry(path string) *ZipEntry {
	// 将path参数转换成据对路径
	absPath, err := filepath.Abs(path)
	// 如果转换出错，调用panic函数终止程序
	if err != nil {
		panic(err)
	}
	// 转换成功返回ZipEntry
	return &ZipEntry{absPath}
}

// 从zip文件中提取class文件
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 使用zip包中的方法打开zip压缩包
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	// 使用defer方法关闭打开的zip压缩包
	defer r.Close()

	// 遍历zip压缩包，寻找class文件
	for _, f := range r.File {
		if f.Name == className {
			// 如果找到对应的class文件，打开文件
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			// 使用defer方法关闭打开的文件
			defer rc.Close()

			// 读取打开的文件内容
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	// 搜索不到class文件，报错
	return nil, nil, errors.New("class not found : " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
