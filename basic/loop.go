package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	//相当于while。go语言中没有while
	//每次 Scan() 一行
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(convertToBinary(5))  //101
	fmt.Println(convertToBinary(13)) //1101

	const filename = "test.txt"
	printFile(filename)

	//forever()
}
