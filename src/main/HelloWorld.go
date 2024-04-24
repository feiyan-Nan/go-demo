package main

import (
	"fmt"
)

const repeatCount = 3

const englishHelloPrefix = "EN"
const frenchHelloPrefix = "french"
const spanishHelloPrefix = "spanish"

type User struct {
	Name string
	Age  int
}

type Paper struct {
	Title string
	User  *User
}

func main() {
	//arr := []int{1, 2, 3}
	//for i := 0; i < len(arr); i++ {
	//	println(arr[i])
	//}
	//str := "南飞雁"
	//for i, val := range str {
	//	//println(val)
	//	fmt.Printf("%d %c \n", i, val)
	//}
	//brr := []byte(str)
	//for i, b := range brr {
	//	fmt.Printf("%d %d \n", i, b)
	//}
	//for {
	//	fmt.Println("Hello World")
	//	time.Sleep(time.Second * 3)
	//	time.Sleep(time.Second * 3)
	//	fmt.Println("Hello World1")
	//	break
	//}

	if_goto()

	//ch := make(chan int, 10)
	//for i := 0; i < 10; i++ {
	//	ch <- rand.Intn(100)
	//}
	//close(ch)
	//for i := range ch {
	//	println(i)
	//}

	//for val := range ch {
	//	println(val)
	//}
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//do(10)
	//do("Hello, World!") // 输出：This is a string value
	//do(true)            // 输出：This is a bool value
	//paper := new(Paper)
	//paper.Title = "Go语言"
	//paper.User = &User{
	//	Name: "张三",
	//	Age:  18,
	//}
	//color := "green"
	//println(time.Now().Month())
	//switch color {
	//case "red":
	//	println("red")
	//	fallthrough
	//case "green":
	//	println("green")
	//	fallthrough
	//case "yellow":
	//	println("yellow")
	//default:
	//	println("other")
	//}

	//var a int = 10
	//a := 10
	//b := 20
	//c := 30
	//d := 40
	//e := 50
	//f := 60
	//g := 70
	//h := 80
	//i := 90
	//j := 100

}

func if_goto() {
	i := 4
	if i%2 == 0 {
		println("AA")
		goto L2
	} else {
		println("BB")
		goto L1
	}
L1:
	println("L1")
	i += 10
	println(i, "L1")
L2:
	println("L2")
	i *= 10
	println(i, "L2")
}

func do(v interface{}) {
	switch value := v.(type) {
	case int:
		println("int", value)
	case string:
		println("string", value)
	default:
		fmt.Println("other", value)
	}
}

//a := 9
//var b *int
//b = &a
//*b = 10
//println(a, *b)

//arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//addPoint(&arr)
//println(arr[0])
//s := make([]int, 10, 12)
//println(cap(s))
//println(len(s))

//aa := map[string]int{"a": 10000, "b": 2}
//aa["H"] = 10
//delete(aa, "a")
//println(aa)
//fmt.Println(aa)
//
//var ch chan int
//ch <- 10
//println(ch)

//ch1 := make(chan int, 4)
//ch1 <- 30
//ch1 <- 40
//ch1 <- 50
//ch2 := ch1
//ch2 <- 100
//<-ch1
//<-ch1
//<-ch1
//d := <-ch1
//fmt.Println(d)

//value, exist := aa["a"]
//println(aa["a"])
//println(len(aa))
//if value, exists := aa["H"]; exists {
//	println(value)
//} else {
//	println("不存在")
//}
//
//for key, value := range aa {
//	println(key, value)
//}

//m := map[string]string{"a": "23"}
//n := make(map[string]int, 10)
//n["a"] = 10
//println(len(n))

//p := make(map[string]int)

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
//}

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
