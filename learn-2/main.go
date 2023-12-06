package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {
	fmt.Print("无换行，正常输出\n")
	fmt.Println("自带换行")
	fmt.Printf("格式化输出，字符串 %s \n", "test")
	fmt.Printf("格式化输出，数字 %d \n", 123)
	fmt.Printf("格式化输出，浮点数 %f \n", 123.456)
	fmt.Printf("格式化输出，浮点数，保留1位小数 %.1f \n", 123.456)
	fmt.Printf("格式化输出，类型 %T \n", "test")
	fmt.Printf("格式化输出，布尔值 %t \n", true)
	fmt.Printf("格式化输出, 任何类型 %v \n", "test")
	person := Person{name: "test"}
	fmt.Printf("格式化输出, 结构体值 %v \n", person)
	fmt.Printf("格式化输出, 结构体字段名和值 %+v \n", person)
	fmt.Printf("格式化输出, 结构体字段名、值和类型 %#+v \n", person)

	a := fmt.Sprint("不添加换行", "test")
	fmt.Print(a)
	b := fmt.Sprintln("自动添加换行")
	fmt.Print(b)
	c := fmt.Sprintf("格式转换 %s", "字符串")
	fmt.Print(c)

	err := fmt.Errorf("格式化错误 %s", "test")
	fmt.Println(err)

	fmt.Println("请输入你的名称")
	var name int
	// output, err := fmt.Scan(&name)
	// output, err := fmt.Scanf("test%d", &name)
	output, err := fmt.Scanln(&name)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output, name)
}
