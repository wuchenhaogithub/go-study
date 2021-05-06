package main

import "fmt"

func main() {
	useFor()
	useForLoop()
}

/**
 for 循环有3种形式，只有一种使用分号
	语法：
		for init; condition; post { }
		for condition { }   省略 init  post 类似于while语句
		for { }  无限循环

*/
func useFor() {
	sum := 1
	for {
		sum++
		if sum == 100 {
			goto end
		}
	}
end:
	sum = 1
	fmt.Println(sum)
}

/**
循环嵌套用法,99乘法表
*/
func useForLoop() {
	var i, j int
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}

/**
循环控制语句
控制语句	描述
break 语句	经常用于中断当前 for 循环或跳出 switch 语句
continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
goto 语句	将控制转移到被标记的语句。
*/
