package main

import "fmt"

func main() {
	// Implicit type กำหนด type ตั้งแต่แรก
	var a int = 10

	fmt.Println(a)

	// Dynamic type กำหนดแค่รูปตัวแปร ส่วน type เดี๋ยวมัน auto ให้
	b := "20"
	fmt.Println(b)
}
