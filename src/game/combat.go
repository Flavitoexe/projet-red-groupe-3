package game

import (
	"fmt"
	"game/character"
)

func (enemy *character.Enemy) enemyPattern(player character.Character, turn int) {
	if turn%3 == 0 {
		*player.Hp -= 2 * enemy.Damage
	} else {

		*player.Hp -= enemy.Damage
	}
}

func (character *character.Character) characterTurn(enemy character.Enemy) {
	fmt.Println("\n=== Menu Combat ===\n")
	fmt.Println("\t- Attaquer\t 1")
	fmt.Println("\t- Inventaire\t 2")
	fmt.Println("\t- Menu\t 0")
	fmt.Println("\nQue souhaitez vous faire ? : ")
	var userChoice int
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		*enemy.Hp -= character.Damage
		fmt.Printf("%s a infligé %d dégats à %s !", character.Name, character.Damage, enemy.Name)
	case 2:
		character.AccesInventory()
		// fmt.Println("Voulez vous utiliser un objet ?\n\t0 : Non\n\t1 : Oui")
		// var userChoice2 int

	case 3:
		character.MainMenu()
	}
}

func (player *Character) trainingFight(enemy Enemy) {
	fmt.Printf("\nVous rencontrez %s et il vous provoque.\n\tDéfoncez le.", enemy.Name)
	cptTour := 1
	for character.Hp > 0 && enemy.Hp > 0 {
		fmt.Println("\n=== Tour 1 ===\n")
		player.characterTurn(enemy)
		if enemy.Hp > 0 {
			enemy.enemyPattern(player, cptTour)
			cpt++
		}
	}

}
