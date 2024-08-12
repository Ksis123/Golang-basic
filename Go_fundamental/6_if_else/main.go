package main

import "fmt"

func main() {
	ratings := 12.0

	if ratings < 5.0 {
		fmt.Println("Disappointed : ", ratings)
	}else if (ratings >= 5.0 ) && (ratings < 7.0) {
		fmt.Println("Normal : ", ratings);
	}else if (ratings >= 7.0 ) && (ratings < 10.0) {
		fmt.Println("Good : ", ratings);
	} else {
		fmt.Println("Invalid rating");
	}
}