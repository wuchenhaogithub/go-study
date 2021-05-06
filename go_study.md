<h1>Go 数据类型</h1>
数据类型的出现，是为了把数据分成所需内存大小不同的数据。编程的时候需要用大数据的时候才能申请大内存，就可以充分利用内存
<table>
<tr>
    <td>类型</td>
    <td>描述</td>
</tr>
<tr>
    <td>布尔型</td>
    <td>true/false</td>
</tr>
<tr>
    <td>数字类型</td>
    <td>int/float32/float64,Go还支持复数，其中位运算采用补码</td>
</tr>
<tr>
    <td>字符串类型</td>
    <td>字符串类型:
字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。</td>
</tr>
<tr>
    <td>派生类型</td>
    <td>(a) 指针类型（Pointer）<br>
(b) 数组类型<br>
(c) 结构化类型(struct)<br>
(d) Channel 类型<br>
(e) 函数类型<br>
(f) 切片类型<br>
(g) 接口类型（interface）<br>
(h) Map 类型</td>
</tr>
</table>
<h2>Go 语言变量</h2>
Go语音变量名由字母，数字，下划线组成，不能数字开头
声明变量一般形式是使用var 关键字

```gotemplate
var identifier type
//可以一次声明多个变量
var identifier1 ,identifier2 type
```
 
<h2>变量声明</h2>
第一种，指定变量类型，如果没有初始值，则变量默认为零值

```gotemplate
package main
import "fmt"
func main() {

// 声明一个变量并初始化
var a = "RUNOOB"
fmt.Println(a)

// 没有初始化就为零值
var b int
fmt.Println(b)

// bool 零值为 false
var c bool
fmt.Println(c)
}
```
数值类型（包括complex64/128）为 0

布尔类型为 false

字符串为 ""（空字符串）

以下几种类型为 nil：
```gotemplate
var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口
```

第二种，根据值自动判定变量的类型
第三种，省略var，注意 := 左侧如果没有声明新的变量，就会产生编译错误，格式如下
```gotemplate
v_name := value

var inVal int
inval :=1 //这时会产生编译错误，因为inVal 已经声明，不需要重现声明

inVal1 := 1 即可
```

<h2>多变量声明</h2>

```gotemplate
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)

```
