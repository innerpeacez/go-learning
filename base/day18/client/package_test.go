package client

import (
	my_init "github.com/innerpeacez/go-learning/base/day18/init"
	"github.com/innerpeacez/go-learning/base/day18/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibSerice(8))

	my_init.HelloInit()
}
