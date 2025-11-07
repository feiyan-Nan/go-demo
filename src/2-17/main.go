package main

const (
	Version = "0.1"
	Build   = iota
	Name
	Author
)

type Checkout func(int, int) int

func GetTotal(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	//goods := []string{"apple", "orange"}
	//for _, v := range goods {
	//	println(v)
	//}
	//goods2 := map[string]string{}
	//goods2["apple"] = "red"
	//println(goods2["apple"])
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run(":8080")
	//var checkout Checkout = func(a, b int) int {
	//	return a + b
	//}
	//println(checkout(1, 2))
	//println(checkout(1, 2))
	//fmt.Println()
	var total = GetTotal(78)
	println(total(34))

	//println(Version, Build, Name, Author)
}
