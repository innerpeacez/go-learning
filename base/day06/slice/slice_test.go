package slice

import "testing"

func TestSlice(t *testing.T) {
	var slice0 []int
	t.Log(len(slice0), cap(slice0))
	slice0 = append(slice0, 1)
	t.Log(len(slice0), cap(slice0))

	slice1 := []int{1, 2, 3, 4, 5}
	t.Log(len(slice1), cap(slice1))
}

func TestSlice2(t *testing.T) {
	slice0 := make([]int, 3, 5)
	t.Log(slice0[0], slice0[1], slice0[2])
	t.Log(len(slice0), cap(slice0))
	slice0 = append(slice0, 1)
	t.Log(slice0[0], slice0[1], slice0[2], slice0[3])
	t.Log(len(slice0), cap(slice0))
}

func TestSlice3(t *testing.T) {
	// 切片声明的方式
	// 1 只声明不初始化，使用append函数添加元素
	var slice0 []int
	slice4 := []int{}
	t.Log(len(slice0), cap(slice0))
	t.Log(slice4)

	// 2 声明并初始化数据
	var slice1 = []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}
	t.Log(slice1, slice2) // 为了不报错

	// 3 使用make函数，声明slice，制定len 和 cap ，并初始化 len 个数据类型的默认值
	slice3 := make([]int, 3, 6)
	t.Log(slice3)
}

func TestSlice4(t *testing.T) {
	// slice 的可变长特性,cap 不够时自动扩容为原cap的两倍
	var slice0 []int
	for i := 0; i < 10; i++ {
		slice0 = append(slice0, i) // 扩容时，内存空间会改变，复制原有的数据到新的内存空间
		t.Log(len(slice0), cap(slice0))
	}
}

func TestSlice5(t *testing.T) {
	month := []string{"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December"}

	// Q2 第二季度 阶段
	Q2 := month[3:6] // 从(3,6]等索引位置截取
	t.Log(len(Q2), cap(Q2), Q2)

	// 共享存储任何位置修改都将影响所有截取的子切片
	summer := month[3:6]
	Q2[0] = "四月"
	t.Log(len(summer), cap(summer), summer)
}

func TestSlice6(t *testing.T) {
	// slice 不能比较，只能判断是否为 nil
	slice0 := []int{1, 2, 3, 4}
	slice1 := []int{1, 2, 3, 4}

	//if slice0 == slice1 {
	//	t.Log("equal")
	//}

	if slice0 == nil {
		t.Log("equal")
	}

	if slice1 == nil {
		t.Log("equal")
	}
}
