package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 12
	arr[9] = 50
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] is %d\n", i, arr[i])
	}
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := [...]string{"a", "b", "c"}

	fmt.Printf("arr1 length is %d ,arr2 length is %d\n", len(arr1), len(arr2))

	arr2d := [...][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Printf("arr2d length is %d", len(arr2d))

}
