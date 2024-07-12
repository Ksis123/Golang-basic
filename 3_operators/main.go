package main

import "fmt"

func main() {
	// Dynamic variable
	a := 9
	b := 2

	fmt.Println("Add : ", a+b)
	fmt.Println("Subtract", a-b)
	fmt.Println("Multiply", a*b)
	fmt.Println("Divide", a/b) // ได้ผลละพ์
	fmt.Println("Mod", a%b)    //คืนค่าเศษที่เหลือจากการหาร:

	fmt.Println()

	fmt.Println("Is", a == b)
	fmt.Println("Not", a != b)
	fmt.Println("More than", a > b)
	fmt.Println("Less than", a < b)
	fmt.Println("More than or equal to", a >= b)
	fmt.Println("Less than or equal to", a <= b)

	fmt.Println()

	fmt.Println("And", true && true)
	fmt.Println("Or", true || false)
}
