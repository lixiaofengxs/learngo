package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating Slice!")
	// 定义一个slice，此时还没有底层array。
	// 但是go语言每个类型都有一个ZeroValue
	// ZeroValue for slice is nil
	var s []int

	// ZeroValue = nil，此时没有底层数组，但是可以append元素
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i*2+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8} //len(s1)=4 cap(s1)=4
	printSlice(s1)

	s2 := make([]int, 16) //len(s2)=16, cap(s2)=16
	printSlice(s2)
	s3 := make([]int, 10, 32) //len(s3)=10, cap(s3)=32
	printSlice(s3)

	fmt.Println("Copying Slice!")
	copy(s2, s1) //s1->s2
	printSlice(s2)

	fmt.Println("Deleting element from slice!")
	//删除s2中下标为3的元素(8)
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front!")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front =", front)
	printSlice(s2)

	fmt.Println("Popping from back!")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("tail =", tail)
	printSlice(s2)
}
