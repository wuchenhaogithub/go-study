package main

import "fmt"

func main()  {
    //getIf()
    getSwitch()
}
/**
if 布尔表达式 {
      在布尔表达式为 true 时执行
}
```
 */
func getIf()  {
    a:=5

    /* 判断布尔表达式 */
    if a < 20 {
        /* 如果条件为 true 则执行以下语句 */
        fmt.Printf("a 小于 20\n" )
    } else {
        /* 如果条件为 false 则执行以下语句 */
        fmt.Printf("a 不小于 20\n" )
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
func getSwitch()  {
    var grade string
    var marks int = 90
    switch marks {
    case 90: grade = "A"
    case 80: grade = "B"
    case 50,60,70 : grade = "C"
    default: grade = "D"
    }
    fmt.Printf("你的等级是 %s\n", grade );


}