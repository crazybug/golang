package main

import "fmt"

func main() {
	i := 6
	switch i {
	case 2:
		fmt.Println("i is equal to 2")
		//fallthrough
	case 6:
		fmt.Println("i is equal to 6")
		//fallthrough
	default:
		fmt.Println("use fallthrough for default")
	}
}
