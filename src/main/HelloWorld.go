package main

import "fmt"

type add func(a, b int) int

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr[1] address %v\n", arr)
	sub_slice := arr[1:3] //从数组创造子切片，len=cap=2
	fmt.Printf("len %d cap %d\n", len(sub_slice), cap(sub_slice))
}

func expansion() {
	s := make([]int, 0, 3)
	prevCap := cap(s)
	for i := 0; i < 100; i++ {
		s = append(s, i)
		currCap := cap(s)
		if currCap > prevCap {
			//每次扩容都是扩到原先的2倍
			fmt.Printf("capacity从%d变成%d\n", prevCap, currCap)
			prevCap = currCap
		}
	}
}
