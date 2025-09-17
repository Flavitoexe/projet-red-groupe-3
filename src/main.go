package main

import "projet-red/game"

func main() {
	p1 := game.Character{}
	// p1.InitCharacter()
	game.Intro()
	p1.CharacterCreation()
	//  p1.MainMenu()
	// e1 := game.Enemy{}
	// e1.InitEnemy()
	// game.TrainingFight(p1, e1)
	// game.AccessShop()
	// p1.AccesInventory()
	p1.AddInventory(game.Item{"Epée plus", 1, "Arme"})
	// p1.AccesInventory()
	p1.RemoveInventory(game.Item{"Flèches classiques", 10, "Cons"}, 4)
	p1.AccesInventory()
}
