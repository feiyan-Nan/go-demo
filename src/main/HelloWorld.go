package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Printf("slice_str", slice_str)
	var a = 12
	println(strconv.Itoa(a))
	switch_condition()
}
func add(a int) int {
	return a + 10
}
func switch_condition() {
	color := "yellow"
	switch color {
	case "green":
		fmt.Println("go")
	case "red", "yellow": //用逗号分隔多个condition，它们之间是“或”的关系，只需要有一个condition满足就行
		fmt.Println("stop")
	}

	//switch后带表达式时，switch-case只能模拟相等的情况；如果switch后不带表达式，case后就可以跟任意的条件表达式
	switch {
	case add(5) > 10:
		fmt.Println("right")
	default:
		fmt.Println("wrong")
	}
}
