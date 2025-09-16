package main

import (
	character "game/character.go"
	enemies "game/enemies.go"
)

func main() {
	p1 := character.Character{}
	p1.initCharacter()
	e1 := enemies.Enemy{}
	e1.initEnemy()
	p1.characterCreation()
	p1.MainMenu()
}
