package main

import "projet-red/game"

func main() {
	p1 := game.Character{}
	p1.InitCharacter()
	p1.CharacterCreation()
	e1 := game.Enemy{}
	e1.InitEnemy()
	// p1.characterCreation()
	//p1.MainMenu()
	game.TrainingFight(p1, e1)
}
