package main

import "fmt"

// golang中指针类型
func main() {

	// 基本数据类型在内存中的布局
	var n int = 10

	// n的地址是什么 &n
	fmt.Println("n的地址是 ", n)

	// 下面的 var ptr *int = &n
	// 1 ptr 是一个指针类型
	// 2 ptr 的类型 *int
	// 3 ptr 本身的值 &n
	var ptr *int = &n
	fmt.Printf("prt=%v\n", ptr)
	fmt.Printf("prt 的地址=%v", &ptr)

	// 获取指针类型所指向的值 使用*
	fmt.Printf("prt 指向的值=%v", *ptr)

}
