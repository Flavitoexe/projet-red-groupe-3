package game

import (
	"fmt"
)

func (enemy *Enemy) enemyPattern(player *Character, turn int) {
	if turn%3 == 0 {
		player.Hp -= 2 * enemy.Damage
	} else {

		player.Hp -= enemy.Damage
	}
}

func (character *Character) characterTurn(enemy *Enemy) {
	fmt.Println("\n=== Menu Combat ===\n")
	fmt.Println("\t- Attaquer\t 1")
	fmt.Println("\t- Inventaire\t 2")
	fmt.Println("\t- Menu\t 0")
	fmt.Println("\nQue souhaitez vous faire ? : ")
	var userChoice int
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		enemy.Hp -= Damage
		fmt.Printf("%s a infligé %d dégats à %s !", character.Name, character.Damage, enemy.Name)
	case 2:
		AccesInventory()
		// fmt.Println("Voulez vous utiliser un objet ?\n\t0 : Non\n\t1 : Oui")
		// var userChoice2 int

	case 3:
		MainMenu()
	}
}

func (player *Character) trainingFight(enemy Enemy) {
	fmt.Printf("\nVous rencontrez %s et il vous provoque.\n\tDéfoncez le.", enemy.Name)
	cptTour := 1

	for player.Hp > 0 && enemy.Hp > 0 {
		player.isDead()
		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		player.characterTurn(enemy)
		if enemy.Hp > 0 {
			enemy.enemyPattern(player, cptTour)
			cpt++
		}
	}
}
