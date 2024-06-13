package main

import (
	"fmt"
	"time"
)

// chan ประกาศ buffer กี่ตัว ต้องใช้ให้ครบ ไม่งั้นจะ down
func main() {
	jobCh := make(chan int, 10)
	resultCh := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		jobCh <- i + 1
	}
	close(jobCh)

	for i := 1; i <= 2; i++ {
		go double(jobCh, resultCh)
	}

	for i := 1; i <= 10; i++ {
		result := <-resultCh
		fmt.Println(result)
	}
}

func double(jobCh <-chan int, resultCh chan<- int) {
	for j := 1; j <= <-jobCh; j++ {
		time.Sleep(1 * time.Second)
		resultCh <- j * 2
	}
}
