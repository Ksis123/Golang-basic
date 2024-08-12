package main

import "fmt"


func emote(ratings float64) string {
	switch{
		case ratings < 5.0:
			return ("Disappointed")
		case (ratings >= 5.0 ) && (ratings < 7.0):
			return "Normal"
		case (ratings >= 7.0 ) && (ratings < 10.0):
			return "Good"
		default:
			return "Invalid rating"
	}
}

func main() {
	fmt.Println(emote(12.0))
	fmt.Println(emote(4.0))
	fmt.Println(emote(6.0))
	fmt.Println(emote(8.0))
	fmt.Println(emote(10.0))

}