package main

import "fmt"

func main() {
	var o1, o2, o3, o4, o5 = "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true

	//  o1, o2, o3, o4, o5 := "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true
	//  การประกาศแบบนี้ จะเป็นการประกาศแบบ short declaration ซึ่งไม่สามารถใช้ใน function ที่ไม่ได้ประกาศตัวแปรไว้แล้ว

	var r rune = '👍'
	//  ตัวอย่าง rune  หรือ char ('') ค่าที่เก็บจริง คือ int

	fmt.Println("เรื่อง :", o1)
	fmt.Println("ปี :", o2)
	fmt.Println("เรตติ้ง :", o3)
	fmt.Println("ประเภท :", o4)
	fmt.Println("ซุปเปอร์ฮีโร่ :", o5)
	fmt.Printf("r: %c\n", r)

	fmt.Println("--------------------------------------------------")

	// ใช้ Printf %v แสดงค่าของตัวแปร
	fmt.Printf("เรื่อง : %v\n", o1)
	fmt.Printf("ปี : %v\n", o2)
	fmt.Printf("เรตติ้ง : %v\n", o3)
	fmt.Printf("ประเภท : %v\n", o4)
	fmt.Printf("ซุปเปอร์ฮีโร่ : %v\n", o5)
	fmt.Printf("รีวิว: %c", r)

}
