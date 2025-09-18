package game

import (
	"fmt"
)

func enemyPattern(player *Character, enemy *Enemy, turn int) {
	if turn%3 == 0 {
		fmt.Printf("\n%s utilise son attaque spéciale et vous inflige %d dégats !", enemy.Name, 2*10)
		player.Hp -= 2 * enemy.Damage
	} else if turn < 4 {
		fmt.Printf("\nAttention !! Plus que %d tours avant que %s n'utilise son attaque spéciale !\n", 3-turn, enemy.Name)
		player.Hp -= enemy.Damage
	} else {
		fmt.Printf("\nAttention !! Plus que %d tours avant que %s n'utilise son attaque spéciale !\n", 3-(turn/3), enemy.Name)
		player.Hp -= enemy.Damage
	}
}

func characterTurn(enemy *Enemy, player *Character) {

	var leave bool
	for !leave {
		fmt.Println("\n== Menu Combat ==")
		fmt.Println("\t1 - Attaque basique")
		fmt.Println("\t2 - Inventaire")
		fmt.Println("\t0 - Fuir")
		userChoice := readInt("\nQue souhaitez vous faire ? : ")

		switch userChoice {
		case 1:
			enemy.Hp -= 10
			fmt.Printf("%s exécute une attaque basique et inflige %d dégats à %s !", player.Name, 10, enemy.Name)
			leave = true
		case 2:
			player.AccesInventory()
		case 0:
			fmt.Println("\nVous fuyez l'ennemi.\nLâche.")
			leave = true
			player.MainMenu()
		default:
			fmt.Println("\nChoix invalide, veuillez réessayer avec une des options affichées.")
			fmt.Scan(&userChoice)
		}
	}
}

func TrainingFight(player *Character) {

	fmt.Println("=== Entraînement ===")
	fmt.Println("\tBienvenue dans le menu d'entraînement !\nVous vous battrez contre un Garde supérieur :")

	var enemy Enemy
	enemy.InitSuperiorGuard()

	player.Infight = true
	fmt.Printf("\nVous rencontrez %s et il vous provoque.\n\tDéfoncez le.", enemy.Name)
	cptTour := 1

	for player.Hp > 0 && enemy.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf("\nIl vous reste %d/%d point de vies\n", player.Hp, player.HpMax)
		fmt.Printf("Il reste %d/%d point de vies à %s\n", enemy.Hp, enemy.MaxHp, enemy.Name)
		characterTurn(&enemy, player)

		if enemy.Hp > 0 {
			enemyPattern(player, &enemy, cptTour)
		}

		player.isDead()
		enemy.isDead()
		cptTour++
	}

}

func Duel(player *Character, enemy Enemy) {

	player.Infight = true
	fmt.Printf("\nVous rencontrez %s et il vous provoque.\n\tDéfoncez le.", enemy.Name)
	cptTour := 1

	for player.Hp > 0 && enemy.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf("\nIl vous reste %d/%d point de vies\n", player.Hp, player.HpMax)
		fmt.Printf("Il reste %d/%d point de vies à %s\n", enemy.Hp, enemy.MaxHp, enemy.Name)
		characterTurn(&enemy, player)

		if enemy.Hp > 0 {
			enemyPattern(player, &enemy, cptTour)
		}

		player.isDead()
		enemy.isDead()
		cptTour++
	}
}
