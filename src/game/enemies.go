package game

import "fmt"

type Enemy struct {
	Name   string
	MaxHp  int
	Hp     int
	Damage int
}

func (enemy *Enemy) InitBandit() {
	*enemy = Enemy{"Bandit", 30, 30, 3}
}

func (enemy *Enemy) InitGuard() {
	*enemy = Enemy{"Garde", 50, 50, 5}
}

func (enemy *Enemy) InitSuperiorGuard() {
	*enemy = Enemy{"Garde supérieur", 75, 75, 10}
}

func (enemy *Enemy) InitArtemis() {
	*enemy = Enemy{"Artémis", 250, 250, 30}
}

func (enemy *Enemy) InitHephaistos() {
	*enemy = Enemy{"Héphaïstos", 400, 400, 40}
}

func (enemy *Enemy) InitAres() {
	*enemy = Enemy{"Arès", 550, 550, 50}
}

func (enemy *Enemy) InitHades() {
	*enemy = Enemy{"Hadès", 700, 700, 70}
}

func (enemy Enemy) isDead() {
	if enemy.Hp == 0 {
		fmt.Printf("\nVous avez vaincu %s !", enemy.Name)
	}
}
