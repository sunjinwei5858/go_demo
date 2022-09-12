package main

import (
	"fmt"
	"math"
	"strconv"
)

// 基本类型转string：使用fmt.Sprintln()或者使用strconv包的函数Itoa()
func int2String(a int) string {
	return fmt.Sprintln(a)
}

func bol2String(a bool) string {
	return fmt.Sprintln(a)
}

// string转基本数据类型：使用strconv包的函数
// string转基本数据类型的注意事项：要确保string类型能够转化为有效类型 否则golang会将其转化为0
func string2Int(a string) int {
	i, _ := strconv.ParseInt(a, 10, 64)
	return int(i)
}

func string2Bol(a string) bool {

	// 说明：
	// 1 strconv.ParseBool(a)该函数会返回两个值 bool和error
	// 2 因为只想获取到value bool，不想获取error 所以使用忽略 下划线_
	b, _ := strconv.ParseBool(a)
	return b
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

	bol := string2Bol("true")
	fmt.Println("字符串格式化布尔=", bol)

	i := string2Int("2222")
	fmt.Println("字符串格式化int=", i)

}
