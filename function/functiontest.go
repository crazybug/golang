package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumAndMultiplie(a, b int) (add, mul int) {
	add = a + b
	mul = a * b
	return
}

func main() {
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y)
	max_xz := max(x, z)

	fmt.Printf("max(%d,%d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d,%d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d,%d) = %d\n", y, z, max(y, z))

	fmt.Println(SumAndMultiplie(5, 6))
	add, mul := SumAndMultiplie(5, 6)
	fmt.Printf("add is %d, multiplie is %d\n", add, mul)
}
