package main

import (
	"errors"
	"fmt"
	"os"
)

// go错误处理
func main() {

	_, err := os.Open("a.txt")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("err.Error(): %v\n", err.Error())
	}

	_, err2 := div(2, 0)

	if err2 != nil {
		fmt.Println(err2)
	}

}

// 自定义异常方式1： errors.New()
func div(a int, b int) (int, error) {
	if b == 0 {
		return b, errors.New("除数不能为0")
	}
	return a / b, nil
}

// 自定义异常方式2：定义结构体struct
