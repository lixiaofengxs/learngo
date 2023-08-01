package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我是周米粒!"

	//获得字节长度
	fmt.Println(len(s))

	//utf-8编码每个中文占3个字节
	//使用[]byte获得字节
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	//ch is a rune. unicode编码
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	//每个中文只计数一次。获得字符数量
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	//按照rune编码，中文3字节编码为一个字符
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	//转成 rune slice
	// rune <-> int32 每个rune 4个字节
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()

	//其他字符串操作
	// Fields, Split, Join
	// Contains, Index
	// ToLower, ToUpper
	// Trim, TrimRight, TrimLeft

}
