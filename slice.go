package main

import "fmt"

/**
切片
切片是数组的抽象
Go数组长度不能改变，在特定的场景中这样的集合就不能用了
与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片容量增大
*/

func main() {
	createSlice()
}

func createSlice() {
	/**
		语法 var identifier []type
	     或make() 函数来创建切片:
			var slice1 []type = make([]type, len)
			也可以简写
			slice1 := make([]type, len) //这里 len 是数组的长度并且也是切片的初始长度。
			也可以指定容量，其中 capacity 为可选参数。
			make([]T, length, capacity)
	*/

	var number = make([]int, 3, 5)
	printSlice(number)
	s := []int{1, 2, 3}
	printSlice(s)
	var slice []int
	printSlice(slice)
	if slice == nil {
		fmt.Printf("切片是空")
	}
	fmt.Println()
	useSplit()
	useSliceFunc()

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v \n ", len(x), cap(x), x)
}

//切片截取
func useSplit() {
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(num[0:2])
	printSlice(num[2:])
	num1 := make([]int, 2)
	printSlice(num1)
}

func useSliceFunc() {
	var number []int
	printSlice(number)
	number = append(number, 1, 2, 3, 4, 5) //追加元素,可以同时追加多个
	printSlice(number)

	//len 获取切片长度，cap获取切片容器长度
	number1 := make([]int, len(number)*2, cap(number)*2) /* 创建切片 numbers1 是之前切片的两倍容量*/
	printSlice(number1)
	copy(number1, number) //cope 复制number 中的值到number1
	printSlice(number1)

}
