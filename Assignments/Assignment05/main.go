package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//1
	for i := 1; i <= 100; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	//2
	fmt.Println("Odd numbers are :")
	j := 1
	for j <= 50 {
		if j%2 != 0 {
			fmt.Print(j, " ")
		}
		j = j + 1
	}
	fmt.Println()
	//3
	fmt.Println("Even numbers are :")
	k := 1
	for {
		if k%2 == 0 {
			fmt.Print(k, " ")
		}
		if k == 50 {
			break
		}
		if k < 50 {
			k = k + 1
		}
	}
	fmt.Println()
	//4
	fmt.Println("Quotient's are:")
	for i := 50; i <= 150; i++ {
		fmt.Print(i/6, " ")
	}
	fmt.Println()
	//5
	fmt.Println("Checking the string with input:")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the String")
	inp1, _ := reader.ReadString('\n')
	inp1 = strings.TrimSpace(inp1)
	if inp1 == "Golang tutorial" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}
	fmt.Println()
	//6
	fmt.Println("Question 06:")
	for i := 1; i <= 80; i++ {
		if i%2 == 0 && i%4 == 0 {
			fmt.Println("Golang tutorial")
		} else if i%2 == 0 && i%4 != 0 {
			fmt.Println("Golang")
		} else if i%4 == 0 {
			fmt.Println("tutorial")
		} else {
			fmt.Println(i)
		}
	}
}
