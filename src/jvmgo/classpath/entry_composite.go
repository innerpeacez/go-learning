package classpath

import "strings"

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	// 创建Entry数组，用于保存pathList中的小Entry
	compositeEntry := []Entry{}
	// 根据分隔符，切割字符串中的，小entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		// 将每一个小的路径，创建成entry
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {

}

func (self CompositeEntry) String() string {
}
