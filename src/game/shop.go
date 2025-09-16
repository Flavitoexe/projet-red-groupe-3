package game

import (
	"fmt"
)

func AccessShop() {
	var weaponsMap map[string]int = map[string]int{"Epée classique ": 15, "Epée rare": 35, "Epée ultra rare": 60, "Arc Oeil de l'Aube": 80}
	var defenseMap map[string]int = map[string]int{"Potion de neuil ": 0, "Potion de vie ": 15, "Potion de force": 45, "Bouclier unique": 30, "Bouclier Spartiate": 60}

	fmt.Println("\n ==== Menu du Marchand ====")
	fmt.Println("\t 1. Inventaire ")
	fmt.Println("\t 2. Que souhaitez vous acheter ")
	fmt.Println("\n= Armes =")
	for weapon, price := range weaponsMap {
		fmt.Printf("\t - %s : %d H\n", weapon, price)
	}
	// fmt.Println("")
	fmt.Println("\n= Défenses =")
	for defense, price := range defenseMap {
		fmt.Printf("\t - %s : %d H\n", defense, price)
	}
	fmt.Println("\n 3. ==== Quitter le magasin (tu vas le regretter) ====")
}
