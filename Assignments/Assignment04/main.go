package main

import (
	"GO/Assignments/Assignment04/calculator"
	"fmt"
)

func main() {
	//1
	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("The type of the array1 is %T\n", array1)
	//2
	slice1 := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("The type of the slice1 is %T\n", slice1)
	//3
	slice2 := slice1[:5]
	slice3 := slice1[5:]
	slice4 := slice1[2:7]
	slice5 := slice1[1:6]
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	//4
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println("Updated slice :", x)
	x = append(x, 53, 54, 55)
	fmt.Println("Updated slice :", x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println("Added y slice with x is: ", x)
	//5
	array2 := [4]string{"Missouri", "Omaha", "Chicago", "Ohio"}
	fmt.Println("Before updating the array: ", array2)
	updateThirdElement(&array2)
	fmt.Println("After updating the 3rd element: ", array2)
	//6
	a := 100
	b := 10
	add := calculator.Add(&a, &b)
	sub := calculator.Sub(&a, &b)
	mul := calculator.Mul(&a, &b)
	q, r := calculator.Div(&a, &b)
	fmt.Println("Addition : ", *add)
	fmt.Println("Subtraction : ", *sub)
	fmt.Println("Multiplication : ", *mul)
	fmt.Printf("Quotient : %v, Remainder : %v ", *q, *r)
}

// 5
func updateThirdElement(array2 *[4]string) {
	array2[2] = "Texas"
}
