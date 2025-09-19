package main

import "projet-red/game"

func main() {
	p1 := game.Character{}
	game.Intro()
	p1.CharacterCreation()
	p1.MainMenu()
}
