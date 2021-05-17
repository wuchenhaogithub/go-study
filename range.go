package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums { //在数组中使用range将传入index 和值两个变量。不需要该元素的序号，可以使用空白符“_”省略了
		sum += num
	}
	fmt.Println(sum)

	for index, val := range nums {
		fmt.Printf("index = %d  num = %d \n", index, val)
	}

	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range 可以枚举Unicode字符串
	for i, c := range "go" {
		fmt.Printf("index = %d  Unicode = %d \n", i, c)

	}
}
