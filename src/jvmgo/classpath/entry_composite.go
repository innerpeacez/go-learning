package classpath

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {

}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {

}

func (self CompositeEntry) String() string {
}
