package game

import (
	"fmt"
)

func (player *Character) AccessShop() {
	// var weaponsMap map[string]int = map[string]int{"Epée classique ": 15, "Epée rare": 35, "Epée ultra rare": 60, "Arc Oeil de l'Aube": 80}
	// var defenseMap map[string]int = map[string]int{"Potion de neuil ": 0, "Potion de vie ": 15, "Potion de force": 45, "Bouclier unique": 30, "Bouclier Spartiate": 60}
	// var consMap map[Item]int = map[Item]int{}

	// var validShop bool
	// for !validShop {
	// 	fmt.Println("\n ==== Menu du Marchand ====")
	// 	fmt.Println("\t1 - Inventaire ")
	// 	fmt.Println("\t2 - Magasin ")
	// 	fmt.Println("\t0 - Quitter le magasin (tu vas le regretter) ")
	// 	// userChoice := readInt("Que souhaitez vous faire ? ")

	// 	var userChoice int
	// 	fmt.Println("\nQue souhaitez vous faire ?")
	// 	fmt.Scan(&userChoice)

	// 	switch userChoice {
	// 	case 1:
	// 		player.AccesInventory()
	// 	case 2:
	// 		zdv
	// 	case 0:
	// 		player.MainMenu()
	// 	default:
	// 		fmt.Println("Choix invalide, veuillez réessayer (Tapez 0, 1 ou 2).")
	// 	}
	// }

	var leave bool
	for !leave {
		fmt.Println("\n=== Marchand ===")
		fmt.Println("\t1 - Armes")
		fmt.Println("\t2 - Armures")
		fmt.Println("\t3 - Consommables")
		fmt.Println("\t0 - Quitter")
		userChoice := readInt("\nQue souhaitez vous acheter ? ")

		switch userChoice {
		case 1:
			i := 1
			fmt.Println("\n== Armes ==")
			for _, weapon := range weaponArr {
				fmt.Printf("\t%d - %s : %d H\n", i, weapon.Name, weaponMap[weapon])
				i++
			}
			fmt.Println("\nQue souhaitez vous acheter ?")

			userChoice2 := readInt("(1-9 pour sélectionner une arme, 0 pour quitter)\n")
			if userChoice2 == 0 {
				player.AccessShop()
				leave = true
			} else {
				if player.Money-weaponMap[weaponArr[userChoice2-1]] < 0 {
					fmt.Println("\nVous n'avez pas assez d'argent.")

				} else {
					player.Money -= weaponMap[weaponArr[userChoice2-1]]
					fmt.Printf("\nObjet '%s' ajouté à votre inventaire !", weaponArr[userChoice2-1].Name)
					fmt.Printf("\n- %d H. %d restants.", weaponMap[weaponArr[userChoice2-1]], player.Money)
					player.AddInventory(weaponArr[userChoice2-1])
				}
			}

		case 2:
			i := 1
			fmt.Println("\n== Armures ==")
			for _, defense := range defenseArr {
				fmt.Printf("\t%d - %s : %d H\n", i, defense.Name, defenseMap[defense])
				i++
			}
			fmt.Println("\nQue souhaitez vous acheter ?")

			userChoice2 := readInt("(1-9 pour sélectionner une défense, 0 pour quitter)\n")
			if userChoice2 == 0 {
				player.AccessShop()
				leave = true
			} else {
				if player.Money-defenseMap[defenseArr[userChoice2-1]] < 0 {
					fmt.Println("\nVous n'avez pas assez d'argent.")
				} else {
					player.Money -= defenseMap[defenseArr[userChoice2-1]]
					fmt.Printf("\nObjet '%s' ajouté à votre inventaire !", defenseArr[userChoice2-1].Name)
					fmt.Printf("\n- %d H. %d restants.", defenseMap[defenseArr[userChoice2-1]], player.Money)
					player.AddInventory(defenseArr[userChoice2-1])
				}
			}

		case 3:
			i := 1
			fmt.Println("\n== Consommables ==")
			for _, consommables := range consArr {
				fmt.Printf("\t%d - %s : %d H\n", i, consommables.Name, consMap[consommables])
				i++
			}
			fmt.Println("\nQue souhaitez vous acheter ?")

			userChoice2 := readInt("(1-9 pour sélectionner une arme, 0 pour quitter)\n")
			if userChoice2 == 0 {
				player.AccessShop()
				leave = true
			} else {
				if player.Money-consMap[consArr[userChoice2-1]] < 0 {
					fmt.Println("\nVous n'avez pas assez d'argent.")
				} else {
					player.Money -= consMap[consArr[userChoice2-1]]
					fmt.Printf("\nObjet '%s' ajouté à votre inventaire !", consArr[userChoice2-1].Name)
					fmt.Printf("\n- %d H. %d restants.", consMap[consArr[userChoice2-1]], player.Money)
					player.AddInventory(consArr[userChoice2-1])
				}
			}

		case 0:
			fmt.Println("Vous quittez le marchand.")
			leave = true

		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
	// fmt.Println("\n= Armes =")

	// for weapon, price := range weaponsMap {
	// 	fmt.Printf("\t - %s : %d H\n", weapon, price)
	// }
	// // fmt.Println("")
	// fmt.Println("\n= Défenses =")
	// for defense, price := range defenseMap {
	// 	fmt.Printf("\t - %s : %d H\n", defense, price)
	// }
}
