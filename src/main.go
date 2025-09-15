package main

import (
	"fmt"
	"os"
	"slices"
)

//test

type Character struct {
	Name      string
	Class     string
	Level     int
	Hp        int
	HpMax     int
	Inventory []Item
}

type Item struct {
	Name     string
	Quantity int
}

func (player *Character) initCharacter() {
	*player = Character{"Flavio", "Assassin", 1, 50, 100, []Item{{"Potion de vie", 1}}}
}

func (player Character) displayInfo() {
	fmt.Println("\n=== Information du personnage ===")
	fmt.Printf("\t - Nom : %s\n", player.Name)
	fmt.Printf("\t - Classe : %s\n", player.Class)
	fmt.Printf("\t - Niveau : %d\n", player.Level)
	fmt.Printf("\t - Pv : %d\n", player.Hp)
	fmt.Printf("\t - Pv Max : %d\n", player.HpMax)
	//test

}

func (player Character) accesInventory() {
	fmt.Println("\n=== Inventaire du personnage ===")
	if len(player.Inventory) == 0 {
		fmt.Println("\n\n\t Inventaire vide\n\n")
	}
	for _, items := range player.Inventory {
		fmt.Printf("\t - %s x %d\n", items.Name, items.Quantity)
	}
}

func (player *Character) takePot() {
	potIndex := slices.IndexFunc(player.Inventory, func(items Item) bool {
		return (items.Name == "Potion de vie" && items.Quantity > 0)
	})
	if potIndex == -1 {
		fmt.Println("Pas de Potion de vie")
		return
	}

	player.Hp += 50
	if player.Hp > player.HpMax {
		player.Hp = player.HpMax
	}
	fmt.Printf("Nouveau Hp : %d\n", player.Hp)

	player.Inventory[potIndex].Quantity -= 1

	if player.Inventory[potIndex].Quantity <= 0 {
		player.Inventory = append(player.Inventory[:potIndex], player.Inventory[potIndex+1:]...)
	}
}

func (player Character) MainMenu() {
	for true {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Printf("\t 1 - Informations du personnage\n")
		fmt.Printf("\t 2 - Inventaire\n")
		fmt.Printf("\t 0 - Quitter\n")
		fmt.Println("\nSelectionner un choix : ")
		var userChoice int
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			player.displayInfo()
		case 2:
			player.accesInventory()
		case 0:
			os.Exit(02)
			return
		default:
			fmt.Println("Erreur : Choix non valide\n")

		}
	}
}

func main() {
	p1 := Character{}
	p1.initCharacter()
	p1.MainMenu()
}
