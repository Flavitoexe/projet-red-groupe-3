package combat

import (
	"fmt"
	"projet-red/character"
)

func (enemy *character.Enemy) enemyPattern() {
	if turn%3 == 0 {
		enemy.Damage *= 2
	}
}

func (character character.Character) characterTurn(enemy character.Enemy) {
	fmt.Println("\n=== Menu Combat ===")
	fmt.Println("\t- Attaquer\t 1")
	fmt.Println("\t- Inventaire\t 2")
	fmt.Println("\t- Menu\t 0")
	fmt.Println("\nQue souhaitez vous faire ? : ")
	var userChoice int
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		&enemy.Hp -= character.Damage
		fmt.Printf("%s a infligé %d dégats à %s !", character.Name, character.Damage, enemy.Name)
	case 2:
		character.accesInventory()
	case 3:
		character.MainMenu()
	}
}
