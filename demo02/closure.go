package main

import "fmt"

// go的闭包：
// 变量f是一个函数并且它引用了其外部作用域中的x变量，此时y就是一个闭包
// 在f的生命周期内 变量x也一直有效
func add() func(y int) int {

	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	// f代表一个生命周期
	f := add()
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)
	r = f(30)
	fmt.Printf("r: %v\n", r)

	fmt.Println("===================")
	// f2是一个新的生命周期
	f2 := add()
	r = f2(100)
	fmt.Printf("r: %v\n", r)
	r = f2(100)
	fmt.Printf("r: %v\n", r)

}
