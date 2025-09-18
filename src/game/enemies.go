package game

import "fmt"

type Enemy struct {
	Name   string
	MaxHp  int
	Hp     int
	Damage int
	Loot   Item
}

func (enemy *Enemy) InitBandit() {
	*enemy = Enemy{"Bandit", 30, 30, 3, leather}
}

func (enemy *Enemy) InitGuard() {
	*enemy = Enemy{"Garde", 50, 50, 5, leather}
}

func (enemy *Enemy) InitSuperiorGuard() {
	*enemy = Enemy{"Garde supérieur", 75, 75, 10, iron}
}

func (enemy *Enemy) InitArtemis() {
	*enemy = Enemy{"Artémis", 250, 250, 30, gold}
}

func (enemy *Enemy) InitHephaistos() {
	*enemy = Enemy{"Héphaïstos", 400, 400, 40, gold}
}

func (enemy *Enemy) InitAres() {
	*enemy = Enemy{"Arès", 550, 550, 50, gold}
}

func (enemy *Enemy) InitHades() {
	*enemy = Enemy{"Hadès", 700, 700, 70, gold}
}

func (enemy Enemy) isDead() {
	if enemy.Hp == 0 {
		fmt.Printf("\nVous avez vaincu %s !", enemy.Name)
		fmt.Printf("\nVous récoltez : '%s' x %d !", enemy.Name, enemy.Loot)
	}
}
