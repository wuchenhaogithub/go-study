package main

import "fmt"

/**
接口
*/

// Phone 定义接口
type Phone interface {
	call()
}

// NokiaPhone 结构体
type NokiaPhone struct {
}

// 方法调用
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("I am iphone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
