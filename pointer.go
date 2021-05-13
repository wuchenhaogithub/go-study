package main

import "fmt"

/**
go语言&取地址符，获取变量的内存地址
*/
func main() {
	var ptr *int
	a := 10
	var ip *int
	var b int
	ip = &a //指针变量的存储地址
	b = *ip

	fmt.Println(a, b, &a, ip, *ip)
	ptr = &b
	fmt.Println(ptr) //空指针
	if ptr == nil {
		fmt.Printf("空指针")
	} else {
		fmt.Printf("not")
	}
	fmt.Println()
	usePointer()
	array()

	x := 100
	y := 200
	usePointerSwap(&x, &y) //指针传参
	fmt.Println(x, y)

}

//指针数值
func array() {
	const max = 3
	var i int
	var a = [3]int{1, 2, 3}
	var ptr [max]*int
	for i = 0; i < max; i++ {
		ptr[i] = &a[i]
	}
	for i = 0; i < max; i++ {
		fmt.Printf("ptr[%d] = %d\n", i, *ptr[i])
	}

}

func usePointerSwap(x *int, y *int) {
	*x, *y = *y, *x
}

//指向指针的指针
func usePointer() {
	var a int
	var ptr *int
	var pptr **int

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

}
