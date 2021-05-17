package main

import "fmt"

/**
函数递归-recursion
*/

func main() {
	num := 15
	fmt.Printf("%d 的阶乘是 %d\n", num, Factorial(uint64(num)))
	fmt.Println("斐波那契数列")
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

// Factorial 乘阶
func Factorial(n uint64) uint64 {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}

/**
斐波那契数 0    1    1    2    3    5    8    13    21    34

fibonacci f(n-1)+f(n-2)
*/
func fibonacci(n int) int {
	if n > 2 {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return n
}
