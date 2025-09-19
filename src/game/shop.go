package game

import (
	"fmt"
)

func (player *Character) AccessShop() {

	var leave bool
	for !leave {
		fmt.Println("\n=== Marchand ===")
		fmt.Println("\t1 - Armes")
		fmt.Println("\t2 - Armures")
		fmt.Println("\t3 - Consommables")
		fmt.Println("\t4 - Matériaux")
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
			fmt.Printf("\n%d H disponibles.", player.Money)
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
					fmt.Printf("\n%s ajouté à votre inventaire !", weaponArr[userChoice2-1].Name)
					fmt.Printf("\n- %d H. %d restants.", weaponMap[weaponArr[userChoice2-1]], player.Money)
					player.AddInventory(weaponArr[userChoice2-1], 1)
				}
			}

		case 2:
			i := 1
			fmt.Println("\n== Armures ==")
			for _, defense := range defenseArr {
				fmt.Printf("\t%d - %s : %d H\n", i, defense.Name, defenseMap[defense])
				i++
			}
			fmt.Printf("\n%d H disponibles.", player.Money)
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
					fmt.Printf("\n%s ajouté à votre inventaire !", defenseArr[userChoice2-1].Name)
					fmt.Printf("\n- %d H. %d restants.", defenseMap[defenseArr[userChoice2-1]], player.Money)
					player.AddInventory(defenseArr[userChoice2-1], 1)
				}
			}

		case 3:
			i := 1
			fmt.Println("\n== Consommables ==")
			for _, consommables := range consArr {
				fmt.Printf("\t%d - %s : %d H\n", i, consommables.Name, consMap[consommables])
				i++
			}
			fmt.Printf("\n%d H disponibles.", player.Money)
			fmt.Println("\nQue souhaitez vous acheter ?")

			userChoice2 := readInt("(1-9 pour sélectionner un consommable, 0 pour quitter)\n")

			if userChoice2 == 0 {
				player.AccessShop()
				leave = true
			} else {
				quantityChoice := readInt("Quelle quantité souhaitez vous acheter ? ")

				if player.Money-(consMap[consArr[userChoice2-1]]*quantityChoice) < 0 {
					fmt.Println("\nVous n'avez pas assez d'argent.")
				} else {
					player.Money -= consMap[consArr[userChoice2-1]] * quantityChoice
					fmt.Printf("\n%s x %d ajouté à votre inventaire !", consArr[userChoice2-1].Name, quantityChoice)
					fmt.Printf("\n- %d H. %d restants.", consMap[consArr[userChoice2-1]]*quantityChoice, player.Money)
					player.AddInventory(consArr[userChoice2-1], quantityChoice)
				}
			}

		case 4:
			i := 1
			fmt.Println("\n== Matériaux ==")
			for _, misc := range miscArr {
				fmt.Printf("\t%d - %s : %d H\n", i, misc.Name, miscMap[misc])
				i++
			}
			fmt.Printf("\n%d H disponibles.", player.Money)
			fmt.Println("\nQue souhaitez vous acheter ?")

			userChoice2 := readInt("(1-3 pour sélectionner un matériau, 0 pour quitter)\n")

			if userChoice2 == 0 {
				player.AccessShop()
				leave = true
			} else {
				quantityChoice := readInt("Quelle quantité souhaitez vous acheter ? ")

				if player.Money-(miscMap[miscArr[userChoice2-1]]*quantityChoice) < 0 {
					fmt.Println("\nVous n'avez pas assez d'argent.")
				} else {
					player.Money -= miscMap[miscArr[userChoice2-1]] * quantityChoice
					// if player.AddInventory(miscArr[userChoice2-1], quantityChoice) {
					fmt.Printf("\n%s x %d ajouté à votre inventaire !", miscArr[userChoice2-1].Name, quantityChoice)
					fmt.Printf("\n- %d H. %d restants.", miscMap[miscArr[userChoice2-1]]*quantityChoice, player.Money)
				}
			}
		case 0:
			fmt.Println("Vous quittez le marchand.")
			leave = true

		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func (player *Character) MenuForgeron() {
	var valid bool
	for !valid {
		fmt.Println("\n=== Menu Forgeron ===")
		fmt.Println("\t1 - Linothorax du Guerrier d'Olympe")
		fmt.Println("\t2 - Cuirasse de Fer des Hoplites Éternels")
		fmt.Println("\t3 - Égide d'Or d'Hélios")
		fmt.Println("\t0 - Quitter")
		fmt.Println("\nQuel set d'armure souhaitez vous faire fabriquer ?")

		userChoice := readInt("")

		switch userChoice {
		case 0:
			fmt.Println("\nVous quittez le forgeron.")
			player.MainMenu()
			valid = true

		case 1:

			fmt.Println("== Linothorax du Guerrier d'Olympe ==")
			fmt.Println("\t1 - Heaume du Regard de l'Hydre")
			fmt.Printf("\t\t7 Cuirs, 15 H\n")
			fmt.Println("\t2 - Linothorax du Sang de l'Hydre")
			fmt.Printf("\t\t10 Cuirs, 20 H\n")
			fmt.Println("\t3 - Grèves de la Bête Serpentine")
			fmt.Printf("\t\t5 Cuirs, 10 H\n")
			fmt.Println("\t0 - Quitter")

			fmt.Println("\nQuelle pièce souhaitez vous faire fabriquer ?")
			userChoice2 := readInt("")

			switch userChoice2 {
			case 0:
				player.MenuForgeron()
				valid = true

			case 1:
				if player.Inventory[24].Quantity < 7 {
					fmt.Printf("Pas assez de cuir (%d/7).\n", player.Inventory[24].Quantity)
				} else if player.Money < 15 {
					fmt.Printf("Pas assez de H (%d, 15).", player.Money)
				} else {

					player.Inventory[24].Quantity -= 7
					player.Money -= 15
					fmt.Println("Le forgeron vous a fabriqué un Heaume du Regard de l'Hydre.")
					player.AddInventory(leatherCap, 1)
				}

			case 2:
				if player.Inventory[24].Quantity < 7 {
					fmt.Printf("Pas assez de cuir (%d/10).\n", player.Inventory[24].Quantity)
				} else if player.Money < 15 {
					fmt.Printf("Pas assez de H (%d, 20).", player.Money)
				} else {

					player.Inventory[24].Quantity -= 10
					player.Money -= 20
					fmt.Println("Le forgeron vous a fabriqué un Linothorax du Sang de l'Hydre.")
					player.AddInventory(leatherChest, 1)
				}

			case 3:
				if player.Inventory[24].Quantity < 7 {
					fmt.Printf("Pas assez de cuir (%d/5).\n", player.Inventory[24].Quantity)
				} else if player.Money < 15 {
					fmt.Printf("Pas assez de H (%d, 10).", player.Money)
				} else {

					player.Inventory[24].Quantity -= 5
					player.Money -= 10
					fmt.Println("Le forgeron vous a fabriqué des Grèves de la Bête Serpentine.")
					player.AddInventory(leatherBoots, 1)
				}

			default:
				fmt.Println(Red + "\nChoix invalide, veuillez taper un nombre proposé." + Reset)
			}

		case 2:
			fmt.Println("== Cuirasse de Fer de l'Hécatonchire ==")
			fmt.Println("\t1 - Heaume des Cent Yeux")
			fmt.Printf("\t\t10 Fers, 20 H\n")
			fmt.Println("\t2 - Cuirasse de Fer de l'Hécatonchire")
			fmt.Printf("\t\t15 Fers, 30 H\n")
			fmt.Println("\t3 - Grèves des Cent Pas")
			fmt.Printf("\t\t7 Fers, 15 H\n")
			fmt.Println("\t0 - Quitter")

			fmt.Println("\nQuelle pièce souhaitez vous faire fabriquer ?")
			userChoice2 := readInt("")

			switch userChoice2 {
			case 0:
				player.MenuForgeron()
				valid = true

			case 1:
				if player.Inventory[25].Quantity < 10 {
					fmt.Printf("Pas assez de fer (%d/10).\n", player.Inventory[25].Quantity)
				} else if player.Money < 20 {
					fmt.Printf("Pas assez de H (%d/20).", player.Money)
				} else {

					player.Inventory[25].Quantity -= 10
					player.Money -= 20
					fmt.Println("Le forgeron vous a fabriqué un Heaume des Cent Yeux.")
					player.AddInventory(ironCap, 1)
				}

			case 2:
				if player.Inventory[25].Quantity < 15 {
					fmt.Printf("Pas assez de fer (%d/15).\n", player.Inventory[25].Quantity)
				} else if player.Money < 30 {
					fmt.Printf("Pas assez de H (%d/30).", player.Money)
				} else {

					player.Inventory[25].Quantity -= 15
					player.Money -= 30
					fmt.Println("Le forgeron vous a fabriqué une Cuirasse de Fer de l'Hécatonchire.")
					player.AddInventory(ironChest, 1)
				}

			case 3:
				if player.Inventory[25].Quantity < 7 {
					fmt.Printf("Pas assez de fer (%d/7).\n", player.Inventory[25].Quantity)
				} else if player.Money < 15 {
					fmt.Printf("Pas assez de H (%d/15).", player.Money)
				} else {

					player.Inventory[25].Quantity -= 7
					player.Money -= 15
					fmt.Println("Le forgeron vous a fabriqué des Grèves des Cent Pas.")
					player.AddInventory(ironBoots, 1)
				}

			default:
				fmt.Println(Red + "\nChoix invalide, veuillez taper un nombre proposé." + Reset)
			}

		case 3:
			fmt.Println("== Égide d'Or d'Hélios ==")
			fmt.Println("\t1 - Couronne Solaire d’Hélios")
			fmt.Printf("\t\t15 Ors, 30 H\n")
			fmt.Println("\t2 - Égide d'Hélios")
			fmt.Printf("\t\t20 Ors, 40 H\n")
			fmt.Println("\t3 - Sandales Flamboyantes du Soleil")
			fmt.Printf("\t\t10 Ors, 20 H\n")
			fmt.Println("\t0 - Quitter")

			fmt.Println("\nQuelle pièce souhaitez vous faire fabriquer ?")
			userChoice2 := readInt("")

			switch userChoice2 {
			case 0:
				player.MenuForgeron()
				valid = true

			case 1:
				if player.Inventory[26].Quantity < 15 {
					fmt.Printf("Pas assez d'or (%d/15).\n", player.Inventory[26].Quantity)
				} else if player.Money < 30 {
					fmt.Printf("Pas assez de H (%d/30).", player.Money)
				} else {

					player.Inventory[26].Quantity -= 15
					player.Money -= 30
					fmt.Println("Le forgeron vous a fabriqué un Heaume des Cent Yeux.")
					player.AddInventory(goldCap, 1)
				}

			case 2:
				if player.Inventory[26].Quantity < 20 {
					fmt.Printf("Pas assez d'or (%d/20).\n", player.Inventory[26].Quantity)
				} else if player.Money < 40 {
					fmt.Printf("Pas assez de H (%d/40).", player.Money)
				} else {

					player.Inventory[26].Quantity -= 20
					player.Money -= 40
					fmt.Println("Le forgeron vous a fabriqué une Égide d'Hélios.")
					player.AddInventory(goldChest, 1)
				}

			case 3:
				if player.Inventory[26].Quantity < 10 {
					fmt.Printf("Pas assez d'or (%d/10).\n", player.Inventory[26].Quantity)
				} else if player.Money < 205 {
					fmt.Printf("Pas assez de H (%d/20).", player.Money)
				} else {

					player.Inventory[26].Quantity -= 10
					player.Money -= 20
					fmt.Println("Le forgeron vous a fabriqué des Sandales Flamboyantes du Soleil.")
					player.AddInventory(goldBoots, 1)
				}
			}

			// fmt.Printf("\t 1 - Chapeau de l'aventurier\n")
			// fmt.Printf("\t 2 - Tunique de l'aventurier\n")
			// fmt.Printf("\t 3 - Bottes de l'aventurier\n")
			// fmt.Printf("\t 4 - Inventaire\n")
			// fmt.Printf("\t 0 - Quitter\n")

			// var menuForgeron int
			// fmt.Println(Blue + "\nFaites votre choix (0, 1, 2, 3 ou 4): " + Reset)
			// fmt.Scan(&menuForgeron)

			// switch menuForgeron {
			// case 1:
			// 	if player.Money >= 5 {
			// 		player.Money -= 5
			// 		player.AddInventory(adventurerCap, 1)
			// 		fmt.Println(Green + "\nYoupi! Le forgeron a fabriqué le chapeau de l'aventurier.\n" + Reset)

			// 	} else {
			// 		fmt.Print(Red + "Oups, vous n'avez pas assez d'argent pour le fabriquer !" + Reset)
			// 	}
			// case 2:
			// 	if player.Money >= 5 {
			// 		player.Money -= 5
			// 		player.AddInventory(adventurerChest, 1)
			// 		fmt.Println(Green + "\nYoupi! Le forgeron a fabriqué la tunique de l'aventurier.\n" + Reset)
			// 	} else {
			// 		fmt.Print(Red + "Oups, vous n'avez pas assez d'argent pour le fabriquer !" + Reset)
			// 	}
			// case 3:
			// 	if player.Money >= 5 {
			// 		player.Money -= 5
			// 		player.AddInventory(adventurerBoots, 1)
			// 		fmt.Println(Green + "\nYoupi! Le forgeron a fabriqué les bottes de l'aventurier.\n" + Reset)
			// 	} else {
			// 		fmt.Print(Red + "Oups, vous n'avez pas assez d'argent pour le fabriquer !" + Reset)
			// 	}
			// case 4:
			// 	player.AccesInventory()
			// case 0:
			// 	return
			// default:
			// 	fmt.Println(Red + "\nChoix invalide, veuillez réessayer !\n" + Reset)
			// }
		}
	}
}
