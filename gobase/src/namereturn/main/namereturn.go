package main

// 分割两位数
func split(num int) (dig2, dig1 int) {
	dig1 = num * 4 / 9
	dig2 = num - dig1
	return dig2, dig1
}

func main() {
	dig2, dig1 := split(17)
	println("十位数:", dig2, "个位数：", dig1)
}
