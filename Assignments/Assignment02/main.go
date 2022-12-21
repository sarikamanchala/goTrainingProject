package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"GO/Assignments/Assignment02/calculator"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the First Integer")
	int1, _ := reader.ReadString('\n')
	// fmt.Println("Integer 1 is:", int1)
	int11, _ := strconv.Atoi(strings.TrimSpace(int1))

	fmt.Println("Enter the Second Integer")
	int2, _ := reader.ReadString('\n')
	int22, _ := strconv.Atoi(strings.TrimSpace(int2))

	fmt.Println(int11, int22)
	a := calculator.Add(int11, int22)
	fmt.Println("Addition is :", a)

	b := calculator.Sub(int11, int22)
	fmt.Println("Subtraction is :", b)

	c := calculator.Mul(int11, int22)
	fmt.Println("Multiplication is :", c)

	q, r := calculator.Div(int11, int22)
	fmt.Println("Qutient is :", q)
	fmt.Println("Remainder is :", r)

}
