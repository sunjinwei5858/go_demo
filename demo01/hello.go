package main

import (
	"fmt"
	"math"
)

// 基本类型转string：使用fmt.Sprintln()
func int2String(a int) string {
	return fmt.Sprintln(a)
}

func bol2String(a bool) string {
	return fmt.Sprintln(a)
}

func main() {

	fmt.Println("hello world")

	// https://tour.go-zh.org/basics/2
	result := math.Abs(-19)
	fmt.Println(result)

	s := int2String(2)
	fmt.Println("int格式化的字符串=", s)

	b := bol2String(true)
	fmt.Println("布尔格式化的字符串=", b)

}
