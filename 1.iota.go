package main

import "fmt"

//关键字 iota 定义常量组中从 0 开始按⾏计数的⾃增枚举值。
const (
	a = iota
	b
	c
	d
	e
	f
)

//KB MB GB TB
const (
	_        = iota             //iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB
	GB
	TB
)

//如果iota自增被打断，须显示恢复
const (
	A = iota //0
	B        //1
	C = "c"  //c
	D        //c
	E = iota //4
	F        //5
)

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {}

func main1() {
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(KB, MB, GB, TB)
	fmt.Println(A, B, C, D, E, F)
	c := Black
	test(c)
	x := 1
	//test(x) //Error test(Color(x))
	test(Color(x))
	test(1) //常量会被编译器自动转换
}
