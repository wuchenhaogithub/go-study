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

func max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/**
函数返回多个值
*/
func swap(x, y int) (int, int) {
	return x, y
}

/**
值传递	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
引用传递	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数
*/

func swaps(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

/**
函数当实参
*/
// 声明一个函数类型
type cb func(int) int

func callBack(x int) int {
	fmt.Printf("this is callBack")
	return x
}

func test() {
	testCallBack(1, callBack)
	fmt.Println()
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
}

func testCallBack(x int, f cb) {
	f(x)
}

/**
闭包函数,匿名函数是一个内联语句或表达式，可以直接使用函数内的变量，不必声明
*/

func getSequence(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return x1, x2 + i
	}
}

func main() {
	//闭包函数
	//addFunc :=getSequence(1,2)
	//fmt.Println(addFunc())
	//fmt.Println(addFunc())
	//fmt.Println(addFunc())
	//fmt.Println(addFunc())

	useFunc()
}

// Circle 定义结构体 /**
type Circle struct {
	radius float64
}

func useFunc() {
	var c Circle
	fmt.Println(c.radius)
	c.radius = 10.00
	fmt.Println(c.getArea())
	c.changeRadius(15)
	fmt.Println(c.radius)
}

//求面积
func (c Circle) getArea() float64 {
	return c.radius * c.radius
}

// 注意如果想要更改成功c的值，这里需要传指针
func (c *Circle) changeRadius(radius float64) {
	c.radius = radius
}
