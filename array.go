package main

import "fmt"

/**
Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：
var variable_name [SIZE] variable_type
*/

func main() {
	useArray()
	agv := getAverage([]float64{1.0, 10.2, 5.3}, 3)
	fmt.Println(agv)
}

func useArray() {
	var arr [10]int
	fmt.Println(arr)

	var balance = [5]float64{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)
	//数组长度不确定可以用...代替
	unable := [...]int{1, 2, 3, 4}
	fmt.Println(unable)
	/**
	初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

	如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
	*/
	arr[1] = 2
	fmt.Println(arr) //数组元素可以通过索引读取
	/**
	Go 语言支持多维数组，以下为常用的多维数组声明方式：
	var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
	*/
	var threedim [4][4][4]int
	fmt.Println(threedim)

	//二维数组是最简单的多维数组  var arrayName [ x ][ y ] variable_type

	tree := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}} //注意：以上代码中倒数第二行的 } 必须要有逗号，因为最后一行的 } 不能单独一行，也可以写成这样：

	fmt.Println(tree)
	row := []int{10, 11, 13}
	tree = append(tree, row) //数组追加
	fmt.Println(tree)
	for i := range tree {
		fmt.Println(tree[i])
	}
}

func getAverage(arr []float64, size int) float64 {
	var avg, sum float64
	for i := range arr {
		sum += arr[i]
	}
	avg = sum / float64(size)
	return avg
}
