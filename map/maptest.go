package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["a"] = "A"
	m["b"] = "B"
	m["c"] = "C"
	fmt.Println(m)
	cm, ok := m["c"]
	fmt.Println("cm is", cm, ok)

	if ok {
		fmt.Println("delete C")
		delete(m, "c")
	}

	fmt.Println(m)
}
