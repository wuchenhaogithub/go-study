package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = "RUNOOB"
	fmt.Println(a)
	var b int
	fmt.Println(b)
	var c error
	fmt.Println(c)

	var st string = "120q"
	num, err := strconv.Atoi(st) //字符串转数字
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
