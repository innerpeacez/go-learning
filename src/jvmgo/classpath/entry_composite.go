package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	// 创建Entry数组，用于保存pathList中的小Entry
	var compositeEntry []Entry
	// 根据分隔符，切割字符串中的，小entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		// 将每一个小的路径，创建成entry
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// 遍历所有的路径
	for _, entry := range self {
		// 搜索class
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found : " + className)
}

func (self CompositeEntry) String() string {
	// 调用调用每一个子路经的String方法
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
