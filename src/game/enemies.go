package game

import "fmt"

type Enemy struct {
	Name   string
	MaxHp  int
	Hp     int
	Damage int
}

func (enemy1 *Enemy) InitEnemy() {
	*enemy1 = Enemy{"Recrue", 50, 50, 5}
}

func (enemy Enemy) isDead() {
	if enemy.Hp == 0 {
		fmt.Printf("\nVous avez vaincu %s !", enemy.Name)
	}
}
