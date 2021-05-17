package main

import "fmt"

/**
map 是一种无序的键值对的集合。map最重要的一点是可以通过key 来快速检索数据，key 类似于索引，指向数据的值
map是集合，可以像迭代数组和切片那样迭代它，不过map 是无序的，我们无法决定返回顺序，map是hash表来实现的

申明变量 默认map 是 nil
var map_variable map[key_data_type]value_data_type
或
map_variable := make(map[key_data_type]value_data_type)
*/

func main() {
	useMap()
}

func useMap() {
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "巴黎"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	//输出地图值
	for country, capital := range countryCapitalMap {
		fmt.Printf("%s 首都是 %s\n", country, capital)
	}

	capital, ok := countryCapitalMap["American"] //如果存在则返回，否则返回false
	fmt.Println(capital)
	fmt.Println(ok)

	delete(countryCapitalMap, "India") //删除元素
	fmt.Println(countryCapitalMap)
}
