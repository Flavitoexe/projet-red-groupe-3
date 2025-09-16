package combat

import (
	"fmt"
	"projet-red/character"
)

func (enemy *character.Enemy) enemyPattern() {
	if turn%3 == 0 {
		enemy.damage *= 2
	}
}

func (character character.Character) characterTurn() {
	fmt.Println("\n=== Menu Combat ===")
	fmt.Println("\t- Attaquer\t 1")
	fmt.Println("\t- Inventaire\t 2")
	fmt.Println("\t- Menu\t 0")
	fmt.Println("\nQue souhaitez vous faire ? : ")
	var userChoice int
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:

	}
}
