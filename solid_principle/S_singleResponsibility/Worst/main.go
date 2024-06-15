package main

import "fmt"

type Player struct{}

func (p *Player) Move() {
	fmt.Println("Can Move.")
}

func (p *Player) Attack() {
	fmt.Println("Can Attack.")
}

func (p *Player) DisplayScoreBoard() {
	fmt.Println("DisplayScoreBoard.")
}


func main() {
	player := &Player{}

	player.Move()
	player.Attack()
	player.DisplayScoreBoard()
}