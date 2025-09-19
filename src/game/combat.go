package game

import (
	"fmt"
	"math/rand"
)

func enemyPattern(player *Character, enemy *Enemy, turn int) {
	if turn%3 == 0 {
		fmt.Printf(Red+"\n%s utilise son attaque spéciale et vous inflige %d dégats !\n"+Reset, enemy.Name, 2*10)
		player.Hp -= 2 * enemy.Damage
	} else if turn < 4 {
		fmt.Printf(Red+"\nAttention !! Plus que %d tours avant que %s n'utilise son attaque spéciale !\n"+Reset, 3-turn, enemy.Name)
		player.Hp -= enemy.Damage
	} else {
		fmt.Printf(Red+"\nAttention !! Plus que %d tours avant que %s n'utilise son attaque spéciale !\n"+Reset, 3-(turn/3), enemy.Name)
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
		userChoice := readInt(Blue + "\nQue souhaitez vous faire ?\n" + Reset)

		switch userChoice {
		case 1:
			enemy.Hp -= player.Damage
			fmt.Printf("%s exécute une attaque basique et inflige %d dégats à %s !", player.Name, player.Damage, enemy.Name)
			leave = true
		case 2:
			player.AccesInventory()
		case 0:
			fmt.Println("\nVous fuyez l'ennemi.\nLâche.\n")
			leave = true
			player.MainMenu()
		default:
			fmt.Println(Red + "\nChoix invalide, veuillez réessayer avec une des options affichées." + Reset)
			fmt.Scan(&userChoice)
		}
	}
}

func TrainingFight(player *Character) {

	fmt.Println("\n=== Entraînement ===")
	fmt.Println("\tBienvenue dans le menu d'entraînement !\n\nVous vous battrez contre un Garde supérieur :")

	var enemy Enemy
	enemy.InitSuperiorGuard()

	player.Infight = true
	fmt.Printf(Bold+"Vous rencontrez ce %s et il vous provoque.\nDéfoncez le.\n"+Reset, enemy.Name)
	cptTour := 1

	for player.Hp > 0 && enemy.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf(Cyan+"\nIl vous reste %d/%d point de vies\n"+Reset, player.Hp, player.HpMax)
		fmt.Printf(Cyan+"Il reste %d/%d point de vies à %s\n"+Reset, enemy.Hp, enemy.MaxHp, enemy.Name)
		characterTurn(&enemy, player)

		if enemy.Hp > 0 {
			enemyPattern(player, &enemy, cptTour)
		}

		player.isDead()
		enemy.isDead()
		cptTour++
	}

}

func DuelRandom(player *Character) {

	enemy := Enemy{}
	randomEnemy := rand.Intn(3)

	switch randomEnemy {
	case 0:
		enemy.InitBandit()
	case 1:
		enemy.InitGuard()
	case 2:
		enemy.InitSuperiorGuard()
	}

	player.Infight = true
	fmt.Printf(Bold+"\nVous rencontrez %s et il vous provoque.\n\tDéfoncez le.\n"+Reset, enemy.Name)
	cptTour := 1

	for player.Hp > 0 && enemy.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf(Cyan+"\nIl vous reste %d/%d point de vies\n"+Reset, player.Hp, player.HpMax)
		fmt.Printf(Cyan+"Il reste %d/%d point de vies à %s\n"+Reset, enemy.Hp, enemy.MaxHp, enemy.Name)
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
	fmt.Printf(Bold+"\nVous rencontrez %s et il vous provoque.\n\tDéfoncez le.\n"+Reset, enemy.Name)
	cptTour := 1

	for player.Hp > 0 && enemy.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf(Cyan+"\nIl vous reste %d/%d point de vies\n"+Reset, player.Hp, player.HpMax)
		fmt.Printf(Cyan+"Il reste %d/%d point de vies à %s\n"+Reset, enemy.Hp, enemy.MaxHp, enemy.Name)
		characterTurn(&enemy, player)

		if enemy.Hp > 0 {
			enemyPattern(player, &enemy, cptTour)
		}

		player.isDead()
		enemy.isDead()
		cptTour++
	}
}

func FightPremierGardien(player *Character, firstGuard Enemy) {

	player.Infight = true
	cptTour := 1

	for player.Hp > 0 && firstGuard.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf("\nIl vous reste %d/%d point de vies\n", player.Hp, player.HpMax)
		fmt.Printf("Il reste %d/%d point de vies à %s\n", firstGuard.Hp, firstGuard.MaxHp, firstGuard.Name)
		characterTurn(&firstGuard, player)

		if firstGuard.Hp > 0 {
			enemyPattern(player, &firstGuard, cptTour)
		}

		player.isDead()
		cptTour++
	}
	if firstGuard.Hp <= 0 {
		PremierGardien()
	}
}

func FightDeuxiemeGardien(player *Character, secondGuard Enemy) {

	player.Infight = true
	cptTour := 1

	for player.Hp > 0 && secondGuard.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf("\nIl vous reste %d/%d point de vies\n", player.Hp, player.HpMax)
		fmt.Printf("Il reste %d/%d point de vies à %s\n", secondGuard.Hp, secondGuard.MaxHp, secondGuard.Name)
		characterTurn(&secondGuard, player)

		if secondGuard.Hp > 0 {
			enemyPattern(player, &secondGuard, cptTour)
		}

		player.isDead()
		cptTour++
	}
	if secondGuard.Hp <= 0 {
		DeuxiemeGardien()
	}
}

func BossFightArtemis(player *Character, artemis Enemy) {
	// FightPremierGardien()
	DieuArtémis()
	player.Infight = true
	cptTour := 1

	for player.Hp > 0 && artemis.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf("\nIl vous reste %d/%d point de vies\n", player.Hp, player.HpMax)
		fmt.Printf("Il reste %d/%d point de vies à %s\n", artemis.Hp, artemis.MaxHp, artemis.Name)
		characterTurn(&artemis, player)

		if artemis.Hp > 0 {
			enemyPattern(player, &artemis, cptTour)
		}

		player.isDead()
		artemis.isDead()
		cptTour++
	}
}

func BossFightHephaistos(player *Character, hephaistos Enemy) {
	DieuHéphaïstos()
	player.Infight = true
	cptTour := 1

	for player.Hp > 0 && hephaistos.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf("\nIl vous reste %d/%d point de vies\n", player.Hp, player.HpMax)
		fmt.Printf("Il reste %d/%d point de vies à %s\n", hephaistos.Hp, hephaistos.MaxHp, hephaistos.Name)
		characterTurn(&hephaistos, player)

		if hephaistos.Hp > 0 {
			enemyPattern(player, &hephaistos, cptTour)
		}

		player.isDead()
		hephaistos.isDead()
		cptTour++
	}
}

func BossFightAres(player *Character, ares Enemy) {
	DieuArès()
	player.Infight = true
	cptTour := 1

	for player.Hp > 0 && ares.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf("\nIl vous reste %d/%d point de vies\n", player.Hp, player.HpMax)
		fmt.Printf("Il reste %d/%d point de vies à %s\n", ares.Hp, ares.MaxHp, ares.Name)
		characterTurn(&ares, player)

		if ares.Hp > 0 {
			enemyPattern(player, &ares, cptTour)
		}

		player.isDead()
		ares.isDead()
		cptTour++
	}
}

func BossFightHades(player *Character, hades Enemy) {
	DieuHadès()
	player.Infight = true
	cptTour := 1

	for player.Hp > 0 && hades.Hp > 0 && player.Infight {

		fmt.Printf("\n=== Tour %d ===\n", cptTour)
		fmt.Printf("\nIl vous reste %d/%d point de vies\n", player.Hp, player.HpMax)
		fmt.Printf("Il reste %d/%d point de vies à %s\n", hades.Hp, hades.MaxHp, hades.Name)
		characterTurn(&hades, player)

		if hades.Hp > 0 {
			enemyPattern(player, &hades, cptTour)
		}

		player.isDead()
		hades.isDead()
		cptTour++
	}
}

func DeroulementCombat1(player Character) {
	GardinensArtemis()
	PremierCombat()
	FightPremierGardien(player)
	if player.Hp > 0 {
		PremierGardien()
		FightDeuxiemeGardien(player)
		if player.Hp > 0 {
			DeuxiemeGardien()
			DieuArtémis()
			BossFightArtemis(player)
			if player.Hp > 0 {
				VaincuDieu()
			}
		}
	}
}

func DeroulementCombat2(player Character) {
	GardiensHéphaïstos()
	PremierCombat()
	FightPremierGardien(player)
	if player.Hp > 0 {
		PremierGardien()
		FightDeuxiemeGardien(player)
		if player.Hp > 0 {
			DeuxiemeGardien()
			DieuHéphaïstos()
			BossFightHephaistos(player)
			if player.Hp > 0 {
				VaincuDieu()
			}
		}
	}
}

func DeroulementCombat3(player Character) {
	GardiensArès()
	PremierCombat()
	FightPremierGardien(player)
	if player.Hp > 0 {
		PremierGardien()
		FightDeuxiemeGardien(player)
		if player.Hp > 0 {
			DeuxiemeGardien()
			DieuArès()
			BossFightAres(player)
			if player.Hp > 0 {
				VaincuDieu()
			}
		}
	}
}

func DeroulementCombat4(player Character) {
	GardiensHadès()
	PremierCombat()
	FightPremierGardien(player)
	if player.Hp > 0 {
		PremierGardien()
		FightDeuxiemeGardien(player)
		if player.Hp > 0 {
			DeuxiemeGardien()
			DieuHadès()
			BossFightHades(player)
			if player.Hp > 0 {
				VictoirePerso()
				go PlaySong()
			}
		}
	}
}
