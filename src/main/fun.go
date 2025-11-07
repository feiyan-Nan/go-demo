package main

func add(a int, b int) int {
	return a + b
}

func sums(callback func(int), nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	callback(sum)
	return sum
}

func returns(a int, b int) (c int) {
	c = a + b
	return
}
