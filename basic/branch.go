package main

import (
	"fmt"
	"os"
)

func grade(score int) string {
	g := ""
	// switch会自动break，除非使用fallthrough
	// switch后面可以没有表达式
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func loop() {
	sum := 0
	//for的条件里不需要括号()，加了报错
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	const filename = "test.txt"
	//contents, err := os.ReadFile(filename)

	// if 的条件没有()。条件后面紧跟 {}
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	//更简洁的写法
	if contents, err := os.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(70),
		grade(88),
		grade(99),
		grade(100),
		//grade(101),
	)

	loop()

}
