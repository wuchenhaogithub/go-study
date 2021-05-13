package main

import "fmt"

/**
结构体 可以由相同类型的数据或不同类型的数据组成的集合
定义结构体
 需要type  struct 语法
 struct	语句定义一个新的数据类型，结构体中有一个或多个成员，type 语句设定了结构体的名称

	 type struct_variable_type struct {
	   member definition
	   member definition
	   ...
	   member definition
	}
	一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

	variable_name := structure_variable_type {value1, value2...valuen}
	或
	variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
*/

func main() {
	useStruct()
}

type Books struct {
	name   string
	author string
	bookId int
}

func useStruct() {
	var b Books
	b = Books{"php", "php", 1}
	javaName := Books{name: "java"} // 忽略的字段为 0 或 空
	fmt.Println(b.name, javaName)
	b.changeBooksName("go")
	changeBookIds(&b)
	printBook(b)
}

//结构体作为函数参数
func printBook(book Books) {

	fmt.Printf("Book name : %s\n", book.name)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}

//结构体指针
func (b *Books) changeBooksName(name string) {
	b.name = name
}

func changeBookIds(book *Books) {
	book.bookId = 10
}
