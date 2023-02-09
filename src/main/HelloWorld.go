package main

import "fmt"

type add func(a, b int) int

func main() {
	m := map[string]int{"语文": 0, "数学": 39}
	m["英文"] = 56

	m3 := map[string]string{
		"guang": "toqiang",
		"hss":   "sasa",
	}
	println(m3)
	fmt.Println(m3)
}
