package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	//go say("world")
	//say("hello")

	//useChar()
	useCloe()
}

/**
管道是用来传递数据的一个数据结构
管道可以通过两个goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符->用于指定通道方向，如果未指定方向，则为双向通道
管道的声明,管道在使用之前必须声明
ch:=make(char int)


ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //把sum 发送到管道 c
}

func useChar() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make(chan int, 10)
	c <- 1
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z, x+y+z)
}

func charVal(val int, c chan int) {
	c <- val
}

func useFibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func useCloe() {
	c := make(chan int, 10)
	go useFibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}
