package main

import "fmt"

func main() {
	slice1 := [5]int{1, 2, 3, 4, 5}
	slice2 := []int{}
	for i := len(slice1) - 1; i >= 0; i-- {
		slice2 = append(slice2, slice1[i])
	}
	fmt.Println(slice2)
}
