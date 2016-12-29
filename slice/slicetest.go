package main

import "fmt"

func main() {
	arr := [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	a, b, c := arr[2:5], arr[3:10], arr[:]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	d := c[:5]
	d[0] = "A"
	fmt.Println(d)
	fmt.Printf("d capacity is %d\n", cap(d))
	fmt.Println(c)

}
