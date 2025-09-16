package game

import (
	"fmt"
	"os"
	"slices"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	Hp        int
	HpMax     int
	Inventory []Item
	Damage    int
	Skill     []string
}

type Item struct {
	Name     string
	Quantity int
}

var class1 Character = Character{
	Name:      "",
	Class:     "Humains",
	Level:     1,
	Hp:        50,
	HpMax:     100,
	Inventory: []Item{},
	Damage:    10,
	Skill:     []string{"Coup de poing"},
}
var class2 Character = Character{
	Name:      "",
	Class:     "Elfes",
	Level:     1,
	Hp:        40,
	HpMax:     80,
	Inventory: []Item{},
	Damage:    10,
	Skill:     []string{"Coup de poing"},
}
var class3 Character = Character{
	Name:      "",
	Class:     "Nains",
	Level:     1,
	Hp:        60,
	HpMax:     120,
	Inventory: []Item{},
	Damage:    10,
	Skill:     []string{"Coup de poing"},
}

func (player *Character) InitCharacter() {
	*player = Character{
		Name:  "",
		Class: "Humain",
		Level: 1,
		Hp:    100,
		HpMax: 150,
		Inventory: []Item{
			{"Item", 1},
			{"Potion de vie", 3},
			{"Lame secrète", 4},
			{"Epée classique", 1},
			{"Epée moyenne", 1},
			{"Epée plus", 1},
			{"Bouclier unique", 1},
			{"Potion force", 2},
			{"Arc", 1},
			{"Flèches classiques", 10},
			{"Flèches empoisonnées", 5},
		},
		Damage: 5,
		Skill:  []string{"Coup de poing"},
	}
}

func (player *Character) classChoice() {
	var validClass bool
	for !validClass {

		fmt.Println("Voici les différentes classes : ")
		fmt.Printf("\nClasse 1 : %s", class1.Class)
		fmt.Printf("\n\tPv : %d\n\tDégats : %d\n\tCapacités : %s", class1.HpMax, class1.Damage, class1.Skill[0])
		fmt.Printf("\nClasse 2 : %s", class2.Class)
		fmt.Printf("\n\tPv : %d\n\tDégats : %d\n\tCapacités : %s", class2.HpMax, class2.Damage, class2.Skill[0])
		fmt.Printf("\nClasse 3 : %s ", class3.Class)
		fmt.Printf("\n\tPv : %d\n\tDégats : %d\n\tCapacités : %s\n", class3.HpMax, class3.Damage, class3.Skill[0])

		var classChoice int
		fmt.Println("\nFaites votre choix (1, 2 ou 3): ")
		fmt.Scan(&classChoice)

		switch classChoice {
		case 1:
			*player = class1
			validClass = true
		case 2:
			*player = class2
			validClass = true
		case 3:
			*player = class3
			validClass = true
		default:
			fmt.Println("Choix invalide, veuillez réessayer.\n")
		}
	}
}

func (player *Character) CharacterCreation() {
	player.classChoice()
	var validName bool
	for !validName {
		fmt.Println("Choisissez un nom pour votre perso : ")
		var nameChoice string = ""
		fmt.Scan(&nameChoice)

		for i := 0; i < len(nameChoice); i++ {
			if !((65 <= nameChoice[i] && nameChoice[i] <= 90) || (97 <= nameChoice[i] && nameChoice[i] <= 122)) {
				fmt.Println("Oups erreur, votre nom n'est pas correct, veuillez recommencer.")
				validName = false
				player.Name = ""
				break
			} else {
				if (65 <= nameChoice[i] && nameChoice[i] <= 90) && i != 0 {
					player.Name += string(nameChoice[i] + 32)
					validName = true
				} else if (97 <= nameChoice[i] && nameChoice[i] <= 122) && i != 0 {
					player.Name += string(nameChoice[i])
					validName = true
				} else if (97 <= nameChoice[i] && nameChoice[i] <= 122) && i == 0 {
					validName = true
					player.Name += string(nameChoice[i] - 32)
				} else if (65 <= nameChoice[i] && nameChoice[i] <= 90) && i == 0 {
					validName = true
					player.Name += string(nameChoice[i])
				}
			}
		}
	}
}

func (player Character) displayInfo() {
	fmt.Println("\n=== Information du personnage ===")
	fmt.Printf("\t - Nom : %s\n", player.Name)
	fmt.Printf("\t - Classe : %s\n", player.Class)
	fmt.Printf("\t - Niveau : %d\n", player.Level)
	fmt.Printf("\t - Pv : %d\n", player.Hp)
	fmt.Printf("\t - Pv Max : %d\n", player.HpMax)

}

func (player Character) AccesInventory() {
	fmt.Println("\n=== Inventaire du personnage ===")

	if len(player.Inventory) == 0 {
		fmt.Println("\n\n\t Inventaire vide\n")
	}

	for _, items := range player.Inventory {
		fmt.Printf("\t - %s x %d\n", items.Name, items.Quantity)
	}

	var userChoice int
	fmt.Println("Voulez vous utiliser un objet ?\n\t1 - 2 : Consommer la potion\n\t3 - 4 : Sélectionner le type de flèche\n\t0 : Quitter l'inventaire")
	fmt.Scan(&userChoice)
}

func (player *Character) takeHealthPot() {
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
	fmt.Printf("Vous utilisez une potion de soin.\nNouveau Hp : %d\n", player.Hp)

	player.Inventory[potIndex].Quantity -= 1

	if player.Inventory[potIndex].Quantity <= 0 {
		player.Inventory = append(player.Inventory[:potIndex], player.Inventory[potIndex+1:]...)
	}
}

func (player Character) MainMenu() {
	for {
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
			player.AccesInventory()
		case 0:
			os.Exit(02)
			return
		default:
			fmt.Println("Erreur : Choix non valide")
		}
	}
}

func (player *Character) isDead() {
	if player.Hp == 0 {
		fmt.Println("\nTu es mort, repenses-y à deux fois la prochaine fois.\n")
		fmt.Println("Tu viens de ressusciter, bonne chance à toi !")
		player.Hp = player.HpMax / 2
		fmt.Printf("Ton nouveau Hp est : %d\n", player.Hp)
	}
}

func (player *Character) spellBook() {
	for _, element := range player.Skill {
		if !(element == "Boule de feu") {
			player.Skill = append(player.Skill, "Boule de feu")
		} else {
			continue
		}
	}
}
