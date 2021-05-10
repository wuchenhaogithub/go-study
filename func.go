package main

import "fmt"

/**
Go语言最少有个main()函数、
可以通过函数来划分不同的功能，逻辑上没个函数执行的是指定的任务
函数声明告诉了编译器函数的名称，返回类型，和参数

func function_name( [parameter list] ) [return_types] {
   函数体
}

 */

func max(num1 int,num2 int) int {
	var result  int
	if num1 > num2 {
		result = num1
	}else {
		result = num2
	}
	return result
}
/**
函数返回多个值
*/
func swap(x,y string)(string,string)  {
	return x,y
}







func main() {
	//fmt.Println(max(1,2))
	a,b:=swap("Google","Runoob")
	fmt.Println(a,b)
}


