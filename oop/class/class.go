package main

import "fmt"

// Class
type FootballPlayer struct { 
	Name string // Attribute
}

//  Method
func (f FootballPlayer)  Dribbing(){
	fmt.Printf("%s: Dribbing\n", f.Name)
}

func main() {
	cr := FootballPlayer{Name: "CR7"}
	lm := FootballPlayer{Name: "Messi"}

	cr.Dribbing()
	lm.Dribbing()
}