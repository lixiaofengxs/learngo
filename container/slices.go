package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func extendSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] // {2, 3, 4, 5}
	// a = arr[4] 报错
	// slice可以向后扩展，但不可以向前扩展
	// s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
	s2 := s1[3:5] //注意：居然不报错

	fmt.Println("extending slice")
	fmt.Println("arr =", arr)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}

func appendElemToSlice() {
	fmt.Println("append element to slice!")

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] // {2, 3, 4, 5}
	s2 := s1[3:5]  // {5, 6}
	// s2后面还剩一个元素7，append(s2, 10)时，10替换了7。s3底层仍然是arr。
	s3 := append(s2, 10) // {5, 6, 10}

	// 添加元素时如果超过cap，系统会重新分配更大的底层数组。
	// 由于值传递的关系，必须接收append的返回值 s = append(s, val)

	// s4, s5 no longer view arr。
	// go会在内存中创建一个新的，更大的数组作为底层array
	s4 := append(s3, 11) // {5, 6, 10, 11}
	s5 := append(s4, 12) // {5, 6, 10, 11, 12}

	fmt.Println("arr =", arr)
	fmt.Println("s3 =", s3)
	fmt.Println("s4 =", s4)
	fmt.Println("s5 =", s5)
}

/* Slice的底层实现：
 *		array: slice的底层数组。
 *		ptr: 指向slice开头的元素。
 *		len: slice的长度是多少，即使用下标[]取值时的范围。下标 >= len 的都报错(越界)。
 *		capacity: 表示从ptr开始到数组结束的整个长度。因此，只要不超过cap时就会进行slice扩展。
 */

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7} //数组

	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	// slice本身没有数据，是对底层array的一个view
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println("s1 =", s1)
	fmt.Println("arr =", arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println("s1 =", s2)
	fmt.Println("arr =", arr)

	// reslice
	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	extendSlice()

	appendElemToSlice()
}
