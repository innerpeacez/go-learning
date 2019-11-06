package classpath

import (
	"os"
	"strings"
)

// 存放路径分隔符
const pathListSeparator = string(os.PathListSeparator)

// 定义Entry接口
type Entry interface {
	// 寻找和加载class文件的方法
	readClass(className string /*class 文件的相对路径*/) ([]byte /*读取到的字节数据*/, Entry /*定位到class文件的Entry*/, error /*错误信息*/)
	// 相当于JAVA中的toString(),返回变量的字符串表示
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
