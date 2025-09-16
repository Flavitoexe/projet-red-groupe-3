package game

import (
	"fmt"
)

func accessMarchand() {
	var aRmes map[string]int = map[string]int{"Epée classique ": 15, "Epée rare": 35, "Epée ultra rare": 60, "Arc Oeil de l'Aube": 80}
	var dEfenses map[string]int = map[string]int{"Potion de neuil(gratuit)": 0, "Potion de vie ": 15, "Potion de force": 45, "Bouclier unique": 30, "Bouclier Spartiate": 60}
	fmt.Println("\n ==== Menu du Marchand ====")
	fmt.Println("\t 1. Inventaire ?")
	fmt.Println("\t 2. Que souhaitez vous acheter ?")
	fmt.Println("Armes\n")
	for a, Listarms := range aRmes {
		fmt.Println("\t", a, Listarms)
	}
	fmt.Println("")
	fmt.Println("Défenses\n")
	for d, Listdefense := range dEfenses {
		fmt.Println("\t", d, Listdefense)
	}
	fmt.Println("\n\t 3. ==== Quitter le magasin (tu vas le regretter) ====")
}
func addInventory() {

}
func removeInventor() {

}
