package main

import "fmt"

/*
go语言使用值传递？引用传递？
	go语言只有值传递一种方式。
	值传递每次都要拷贝一份，效率不高，如何使用指针？
*/

func pointer(pa *int) {
	*pa = 10
}

func swapByPointer(a, b *int) {
	*b, *a = *a, *b
}

func swapByValue(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	var k int = 3
	// 指针不能远算
	var pk *int = &k
	*pk = 4
	fmt.Println(k)

	a, b := 3, 4
	swapByPointer(&a, &b)
	fmt.Println(a, b)

	a, b = swapByValue(3, 4)
	fmt.Println(a, b)
}
