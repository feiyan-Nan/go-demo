package main

const repeatCount = 3

const englishHelloPrefix = "EN"
const frenchHelloPrefix = "french"
const spanishHelloPrefix = "spanish"

func main() {
	//a := 9
	//var b *int
	//b = &a
	//*b = 10
	//println(a, *b)

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	addPoint(&arr)
	println(arr[0])
	s := make([]int, 10, 12)
	println(cap(s))
	println(len(s))
	//for i := range arr {
	//	arr[i] += 1
	//	fmt.Printf("%p  %d\n", &arr[i], arr[i])
	//}
	////println(arr)
	//fmt.Println(arr)
	//println(hello(frenchHelloPrefix) + "world")
	//println(repeat("liguigong "))
	////numbers := [5]int{1, 2, 3, 4, 5}
	//println(sum([]int{1, 2, 3, 4, 5}))
	//add()
	//s := "明天a会更改好"
	//for _, ele := range s {
	//	fmt.Printf("%c\n", ele)
	//}
	//
	//a := int('a')
	//fmt.Printf("%d\n", a)
	//println(a)
	//arr1 := []byte(s)
	//s1 := string(arr1)
	//println(s1)
	//println("134r")
	//arr := strings.Split(s, "")
	//
	//println(strings.HasPrefix(s, "a"))
	//println(strings.Index(s, "v"))
	//println(strings.Clone(s))
	//
	//println(strings.Join([]string{"a", "b
	//}
	//arr1 := []byte(s)
	//s1 := string(arr1)
	//println(s1)
	//println("134r")
	//arr := strings.Split(s, "")
	//
	//println(strings.HasPrefix(s, "a"))
	//println(strings.Index(s, "v"))
	//println(strings.Clone(s))
	//
	//println(strings.Join([]string{"a", "b", "c"}, " "))
	//
	//for _, i2 := range arr {
	//	println(i2)
	//}

	//for i := 0; i < len(arr); i++ {
	//	println(arr[i])
	//}
	//字符串拼接
	//println(arr)
}

func addPoint(arr *[9]int) {
	//arr[0] = 200
	(*arr)[0] = 300
}

func sum(numbers []int) int {
	var sum int
	//for i := 0; i < len(numbers); i++ {
	//	sum += numbers[i]
	//}
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func repeat(s string) string {
	var repeater string
	for i := 0; i < repeatCount; i++ {
		repeater += s
	}
	return repeater
}

func hello(language string) (prefix string) {
	switch language {
	case englishHelloPrefix:
		prefix = "英语, "
	case frenchHelloPrefix:
		prefix = "法语, "
	case spanishHelloPrefix:
		prefix = "西班牙, "
	default:
		prefix = "英语, "
	}
	return prefix
}
