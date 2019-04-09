package test

import "testing"

func TestSwitch(t *testing.T) {
	// 每个case结束，默认break ,不存在switch穿透
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2: // case 中支持多个条件匹配
			t.Log("偶数")
		case 1, 3:
			t.Log("奇数")
		default:
			t.Log("不是0-3")
		}

	}

}
