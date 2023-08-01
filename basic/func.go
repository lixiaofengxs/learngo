package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	case "%":
		_, r := div(a, b) //不想要的参数用 _
		return r, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// 13 / 3 = 4...1
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}

// op是一个函数。func(int, int) int
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3, 4, "*"))

	q, r := div(13, 3)
	fmt.Println(q, "...", r)

	fmt.Println(eval(13, 3, "%"))

	if result, err := eval(3, 4, "-"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, 3, 4))

	//匿名函数
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
