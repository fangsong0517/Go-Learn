package main

import "fmt"

//要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。⽆论哪种转
//换，都会重新分配内存，并复制字节数组
func main2() {
	s := "abcd汉字"
	bs := []byte(s)
	bs[0] = 'c'
	println(string(bs))
	for _, r := range s {
		fmt.Printf("%c, ", r)
	}
}
