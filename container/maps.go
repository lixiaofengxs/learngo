package main

import "fmt"

/* map的key:
 *	map使用哈希表，必须可以比较相等
 *	除了slice, map, function的内建类型都可以作为key
 *	struct类型不包含上述字段，也可以作为key
 */

func main() {
	m := map[string]string{
		"name":   "zml",
		"age":    "12",
		"gender": "man",
	}
	fmt.Println(m)

	//m2 is empty map, not nil
	m2 := make(map[string]int)
	fmt.Println(m2)

	//m3 = nil
	var m3 map[string]int
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	//只打印key
	for k := range m {
		fmt.Println(k)
	}

	//只打印value
	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("Getting values")
	name := m["name"]
	fmt.Println(name)

	//假设name拼错了，会出现什么情况
	nameErr := m["nam"]
	//从map里面取一个不存在的key时，默认拿到的是ZeroValue，string即空串
	fmt.Printf("%q\n", nameErr)

	//判断key是否在map中
	if gender, ok := m["gender"]; ok {
		fmt.Println(gender, ok)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)

}
