package game

import (
	"fmt"
)

func enemyPattern(player *Character, enemy *Enemy, turn int) {
	if turn%3 == 0 {
		fmt.Printf("\n%s utilise son attaque spéciale et vous inflige %d dégats !", enemy.Name, 2*10)
		player.Hp -= 2 * 10
		turn = 0
	} else {
		fmt.Printf("\nAttention !! Plus que %d tours avant que %s n'utilise son attaque spéciale !", turn/3, enemy.Name)
		player.Hp -= 10
	}
}

func characterTurn(enemy *Enemy, player *Character) {
	fmt.Println("\n== Menu Combat ==\n")
	fmt.Println("\t- Attaquer\t 1")
	fmt.Println("\t- Inventaire\t 2")
	fmt.Println("\t- Menu\t 0")
	fmt.Println("\nQue souhaitez vous faire ? : ")
	var userChoice int
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		enemy.Hp -= 10
		fmt.Printf("%s a infligé %d dégats à %s !", player.Name, 10, enemy.Name)
	case 2:
		player.AccesInventory()
		// fmt.Println("Voulez vous utiliser un objet ?\n\t0 : Non\n\t1 : Oui")
		// var userChoice2 int

	case 3:
		player.MainMenu()
	default:
		fmt.Println("\nChoix invalide, veuillez réessayer avec une des options affichées.")
		fmt.Scan(&userChoice)
	}
}

func TrainingFight(player Character, enemy Enemy) {
	fmt.Printf("\nVous rencontrez %s et il vous provoque.\n\tDéfoncez le.", enemy.Name)
	cptTour := 1

	for player.Hp > 0 && enemy.Hp > 0 {
		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf("\nIl vous reste %d/%d point de vies\n", player.Hp, player.HpMax)
		fmt.Printf("Il reste %d/%d point de vies à %s\n", enemy.Hp, enemy.MaxHp, enemy.Name)
		characterTurn(&enemy, &player)
		if enemy.Hp > 0 {
			enemyPattern(&player, &enemy, cptTour)
			cptTour++
		}
	}

}

// func main() {
// 	e1 := enemies.Enemy{}
// 	e1.initEnemy()
// 	p1 := character.Character{}
// 	p1.initCharacter()
// 	trainingFight(p1, e1)
// }
