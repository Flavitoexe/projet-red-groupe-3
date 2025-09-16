package main

import "projet-red/game"

func main() {
	p1 := game.Character{}
	// p1.InitCharacter()
	p1.CharacterCreation()
	p1.MainMenu()
	// e1 := game.Enemy{}
	// e1.InitEnemy()
	// game.TrainingFight(p1, e1)
}
