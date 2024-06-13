package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // ประกาศตัวแปร wg ของประเภท WaitGroup
	wg.Add(3)             // กำหนดให้ WaitGroup รอการทำงานของ 3 Goroutines

	for i := 1; i <= 3; i++ { // วนลูปตั้งแต่ i = 1 ถึง i = 3
		go func(i int, wg *sync.WaitGroup) { // สร้าง Goroutine สำหรับแต่ละค่า i
			defer wg.Done() // เมื่อ Goroutine นี้เสร็จสิ้น จะลดจำนวน counter ของ WaitGroup ลง 1
			fmt.Println(i)  // พิมพ์ค่าของ i ออกมา
		}(i, &wg) // ส่งค่า i และตัวชี้ไปยัง wg ให้กับฟังก์ชัน Goroutine
	}

	wg.Wait() // รอจนกว่า Goroutines ทั้งหมดจะเสร็จสิ้น
}

// func hello() {
// 	for {
// 		fmt.Println("Hello from another thread")
// 	}
// }
