package const_test

import "testing"
import "fmt"

const (
	one = 1 << iota
	two
	three
)

func TestConst(t *testing.T) {
	fmt.Print(one)
}
