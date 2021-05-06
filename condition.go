package main

import (
	"fmt"
	"time"
)

func main() {
	//getIf()
	//getSwitch()
	getSelect()
}

/**
if 布尔表达式 {
      在布尔表达式为 true 时执行
}
```
*/
func getIf() {
	a := 5

	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

/**

switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}

*/
func getSwitch() {
	var grade string
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

/**
  select 是go的一个控制结构，类似于通信的switch语句。每个case必须是一个通信操作（channel）。要么发送要么是接受。
  select随机执行一个可运行的case。如果case不能运行，他将会阻塞，知道case可运行
  每个 case 都必须是一个通信
  所有 channel 表达式都会被求值
  所有被发送的表达式都会被求值
  如果任意某个通信可以进行，它就执行，其他被忽略。
  如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
  否则：
  如果有 default 子句，则执行该语句。
  如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
*/
func getSelect() {
	ch := make(chan int)
	c := 0
	stopCh := make(chan int)
	go Chan(ch, stopCh)
	for {
		select {
		case c = <-ch:
			fmt.Printf("receive", c)
			fmt.Printf("channel")
			fmt.Println()
		case s := <-ch:
			fmt.Printf("receive--", s)
			fmt.Println()
		case _ = <-stopCh:
			goto end
		}
	}
end:
}

func Chan(ch chan int, stopCh chan int) {
	var i int
	i = 10
	for j := 0; j < i; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- 0
}
