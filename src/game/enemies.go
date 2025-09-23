package game

import (
	"fmt"
	"math/rand"
)

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

func (enemy *Enemy) InitFirstGuard() {
	*enemy = Enemy{Hp: 100, Name: "Premier Gardien", MaxHp: 100, Damage: 15}
}

func (enemy *Enemy) InitSecondGuard() {
	*enemy = Enemy{Hp: 150, Name: "Second Gardien", MaxHp: 150, Damage: 20}
}

func (enemy *Enemy) InitArtemis() {
	*enemy = Enemy{"Artémis", 250, 250, 30, gold}
}

func (enemy *Enemy) InitHephaistos() {
	*enemy = Enemy{"Héphaïstos", 300, 300, 40, gold}
}

func (enemy *Enemy) InitAres() {
	*enemy = Enemy{"Arès", 350, 350, 50, gold}
}

func (enemy *Enemy) InitHades() {
	*enemy = Enemy{"Hadès", 400, 400, 60, gold}
}

func (player *Character) EnemyDead(enemy Enemy) {
	if enemy.Hp == 0 {
		fmt.Printf("\nVous avez vaincu %s !", enemy.Name)
		fmt.Printf("\nVous récoltez : %d !", rand.Intn(10))
		player.Infight = false
	}
}

func (player *Character) GuardDead(guard Enemy) {
	if guard.Hp == 0 {
		fmt.Printf("\nVous avez vaincu %s !", guard.Name)
		fmt.Printf("\nVous récoltez : %d !", rand.Intn(50-25+1)+25)
		player.Infight = false

	}
}

func (player *Character) BossDead(boss Enemy) {
	if boss.Hp == 0 {
		fmt.Printf("\nVous avez vaincu %s !", boss.Name)
		fmt.Printf("\nVous récoltez : %d !", 100)
		player.Infight = false

	}
}
