package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	first_name  string
	last_name   string
	fav_country []string
}
type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle   vehicle
	fourWheel bool
}
type sedan struct {
	vehicle vehicle
	luxury  bool
}
type SQUARE struct {
	length int
}
type CIRCLE struct {
	radius int
}
type SHAPE interface {
	AREA() float32
}

func Hello() string {
	return "Hello"
}
func INFO(shape SHAPE) {
	fmt.Println(shape.AREA())
}
func (square SQUARE) AREA() float32 {
	return float32(square.length * square.length)
}
func (circle CIRCLE) AREA() float32 {
	return PI * float32(circle.radius*circle.radius)
}

var mapvariable map[string]person = make(map[string]person)

const PI float32 = 3.14

func main() {
	Hello()

	person1 := person{"John", "Snow", []string{"USA", "Austrilia", "Italy"}}
	person2 := person{"Tyrion", "Lannister", []string{"Austria", "Turkey", "Dubai"}}

	//1
	fmt.Println("First person struct ", person1)
	fmt.Println("Person First Name ", person1.first_name)
	fmt.Println("Looping over the favourite Country")
	for i, v := range person1.fav_country {
		fmt.Println(i, " ", v)
	}

	fmt.Println("Second person struct ", person2)
	fmt.Println("Person First Name ", person2.first_name)
	fmt.Println("Looping over the favourite Country")
	for i, v := range person2.fav_country {
		fmt.Println(i, " ", v)
	}
	fmt.Println("------------------------------------")

	//2
	mapvariable[person1.last_name] = person1
	mapvariable[person2.last_name] = person2
	for _, v := range mapvariable {
		fmt.Println(v)
	}
	fmt.Println("-------------------------------------")
	//3
	truck1 := truck{vehicle{2, "Orange"}, true}
	sedan1 := sedan{vehicle{4, "Black"}, true}
	fmt.Println("Truck1 Details : ", truck1)
	fmt.Println("Sedan Details : ", sedan1)
	fmt.Println("Truck1 is of color : ", truck1.vehicle.color)
	fmt.Println("Sedan1 number of doors : ", sedan1.vehicle.doors)
	fmt.Println("-------------------------------------")
	//4
	square := SQUARE{3}
	circle := CIRCLE{2}
	fmt.Print("SQUARE AREA : ")
	INFO(square)
	fmt.Print("CIRCLE AREA : ")
	INFO(circle)
	fmt.Println("-------------------------------------")
	//5
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the String:")
	inp1, _ := reader.ReadString('\n')
	inp11 := strings.TrimSpace(inp1)
	fmt.Println("input is :", inp11)
	input := strings.Split(inp11, " ")
	//input := strings.SplitAfter(inp11, " ")
	fmt.Println(input)
	var m1 map[string]int = make(map[string]int)
	for _, v := range input {
		//key := strings.TrimSpace(v)
		//fmt.Println(i, " ", key, " before", m1[key])
		if m1[v] == 0 {
			m1[v] = 1
		} else {
			m1[v] = m1[v] + 1
		}
		//fmt.Println(i, " ", key, " after", m1[key])
	}
	fmt.Println("Ocuurances of strings:", m1)
}
