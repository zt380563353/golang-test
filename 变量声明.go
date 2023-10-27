package main

import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func defFun1() {
	// 声明一个变量并初始化
	var a = "TEST"
	fmt.Println(a)

	// 没有初始化就为0值
	var b int
	fmt.Println(b)

	// bool 零值为False
	var c bool
	fmt.Println(c)
}

func defFun2() {
	var i int
	var f float64
	var b bool
	var s string
	/* 0 0 false "" */
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func main() {
	defFun1()
	defFun2()

	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
