package main

import "fmt"

// 数组是值类型
// [3]int 和 [5]int 是不同的数据类型，不能传递[3]int
// 调用 func f(arr [10]int) 会拷贝数组
// go语言中一般不直接使用数组，用切片
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func passByArrayPointer(arr *[5]int) {
	arr[0] = 100
}

func main() {
	var arr1 [5]int

	// 使用这种方式时，必须要给数组初值
	arr2 := [3]int{1, 3, 5}

	// 让编译器自己计算数组元素个数
	// 注意：并不是[],这种是切片，需要写成[...]
	arr3 := [...]int{2, 4, 6, 8, 10}

	// 二维数组
	var grid [4][5]int

	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 使用range关键字可以同时获得数组的索引和元素值
	for index, value := range arr2 {
		fmt.Println(index, value)
	}

	for _, value := range arr3 {
		fmt.Println(value)
	}

	fmt.Println(arr1, arr2, arr3, grid)

	printArray(arr3)
	fmt.Println(arr3)

	passByArrayPointer(&arr3)
	fmt.Println(arr3)
}
