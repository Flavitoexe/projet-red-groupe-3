package game

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
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
	BowState  int
}

type Item struct {
	Name     string
	Quantity int
	Tag      string
}

type Equipment struct {
	Casque string
	Body   string
	Foot   string
}

var basicInventory []Item = []Item{
	{"Lame secrète", 4, "Arme"},
	{"Epée classique", 1, "Arme"},
	{"Epée moyenne", 0, "Arme"},
	{"Epée plus", 0, "Arme"},
	{"Arc", 1, "Arme"},
	{"Bouclier unique", 1, "Armure"},
	{"Potion de vie", 1, "Cons"},
	{"Potion de force", 2, "Cons"},
	{"Flèches classiques", 10, "Cons"},
	{"Flèches empoisonnées", 0, "Cons"},
}

var class1 Character = Character{
	Name:      "",
	Class:     "Humains",
	Level:     1,
	Hp:        50,
	HpMax:     100,
	Inventory: basicInventory,
	Damage:    10,
	Skill:     []string{"Coup de poing"},
}

var class2 Character = Character{
	Name:      "",
	Class:     "Elfes",
	Level:     1,
	Hp:        40,
	HpMax:     80,
	Inventory: basicInventory,
	Damage:    10,
	Skill:     []string{"Coup de poing"},
}

var class3 Character = Character{
	Name:      "",
	Class:     "Nains",
	Level:     1,
	Hp:        60,
	HpMax:     120,
	Inventory: basicInventory,
	Damage:    10,
	Skill:     []string{"Coup de poing"},
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
			fmt.Println("Choix invalide, veuillez réessayer :")
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
	fmt.Printf("\t - Points de vie actuels : %d\n", player.Hp)
	fmt.Printf("\t - Points de vie max : %d\n", player.HpMax)
	fmt.Printf("\t - Dégats : %d\n", player.Damage)

}

func readInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur de lecture, réessayez.")
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			fmt.Println("Choix invalide, veuillez réessayer (Tapez 0, 1 ou 2).")
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Entrée invalide, veuillez taper un nombre (ex: 1, 2, 0).")
			continue
		}
		return n
	}
}

func (player *Character) AddInventory(item Item) {

	// if !(item.Tag == "Cons") {
	// 	item.Quantity = 1
	// }

	for index := range player.Inventory {
		if player.Inventory[index].Name == item.Name {
			player.Inventory[index].Quantity += item.Quantity
			fmt.Printf("Quantité de %s : %d\n", item.Name, player.Inventory[index].Quantity)
			return
		}
	}
}

func (player *Character) RemoveInventory(item Item, quantity int) {

	for index := range player.Inventory {
		if player.Inventory[index].Quantity >= 0 && player.Inventory[index].Quantity-item.Quantity >= 0 {
			player.Inventory[index].Quantity -= quantity
			fmt.Printf("Nouvelle quantité de %s : %d\n", item.Name, player.Inventory[index].Quantity)
			return
		}
	}
	fmt.Printf("Vous n'avez pas de %s.\n", item.Name)
}

func (player Character) AccesInventory() {
	fmt.Println("\n=== Inventaire du personnage ===")

	if len(player.Inventory) == 0 {
		fmt.Println("\n\n\t Inventaire vide\n\n")
		return
	}

	fmt.Println("\n\t1 == Armes ==\n\t2 == Armure ==\n\t3 == Consommables ==\n")
	userChoice := readInt("\nQue souhaitez vous voir ? ")

	switch userChoice {

	case 1:
		fmt.Println("== Armes ==")
		for _, weapon := range player.Inventory {
			if weapon.Tag != "Arme" || weapon.Quantity == 0 {
				continue
			} else if weapon.Tag == "Arme" {
				fmt.Printf("\t - %s\n", weapon.Name)
			}
		}
	case 2:
		fmt.Println("== Armures ==")
		for _, armor := range player.Inventory {
			if armor.Tag != "Armure" || armor.Quantity == 0 {
				continue
			} else if armor.Tag == "Armure" {
				fmt.Printf("\t - %s\n", armor.Name)
			}
		}
	case 3:
		fmt.Println("== Consommables ==")
		for _, cons := range player.Inventory {
			if cons.Tag != "Cons" {
				continue
			} else if cons.Tag == "Cons" {
				fmt.Printf("\t - %s x %d\n", cons.Name, cons.Quantity)
			}
		}

		fmt.Println("\nVoulez vous utiliser un objet ?\n(0 pour sortir, 1 - 2 pour consommer une potion, 3 - 4 pour seléctionner un type de flèche)\n")
		userChoice2 := readInt("Votre choix : ")

		switch userChoice2 {
		case 0:
			player.AccesInventory()
		case 1:
			player.takeHealthPot()
		case 2:
			player.takeStrenghtPot()
		case 3:
			if player.Inventory[8].Quantity <= 0 {
				fmt.Print("Plus de flèches normales. Allez en acheter au marchand.\n")
			} else {
				fmt.Printf("Vous équipez les flèches normales. (%d restantes)\n", player.Inventory[8].Quantity)
				player.BowState = 0
			}
		case 4:
			if player.Inventory[9].Quantity <= 0 {
				fmt.Println("Plus de flèches empoisonnées. Allez en acheter au marchand.\n")
			} else {
				fmt.Printf("Vous équipez les flèches empoisonnées. (%d restantes)\n", player.Inventory[9].Quantity)
				player.BowState = 1
			}
		}
	}
	// for i := 0; i < len(player.Inventory); i++ {

	// }
	// for _, items := range player.Inventory {
	// 	if items.Quantity == 0 {
	// 		continue
	// 	} else {
	// 		fmt.Printf("\t - %s x %d\n", items.Name, items.Quantity)
	// 	}
	// }
}

func (player *Character) takeHealthPot() {
	potIndex := slices.IndexFunc(player.Inventory, func(items Item) bool {
		return (items.Name == "Potion de vie" && items.Quantity > 0)
	})
	if potIndex == -1 {
		fmt.Println("Plus de Potion de vie. Allez en acheter au marchand.")
		return
	}

	player.Hp += 50
	if player.Hp > player.HpMax {
		player.Hp = player.HpMax
	}
	fmt.Printf("Vous utilisez une potion de soin.\nNouveau Hp : %d\n", player.Hp)

	player.Inventory[potIndex].Quantity -= 1

	// if player.Inventory[potIndex].Quantity <= 0 {
	// 	player.Inventory = append(player.Inventory[:potIndex], player.Inventory[potIndex+1:]...)
	// }
}

func (player *Character) takeStrenghtPot() {
	// potIndex := slices.IndexFunc(player.Inventory, func(items Item) bool {
	// 	return (items.Name == "Potion de force" && items.Quantity > 0)
	// })
	if player.Inventory[7].Quantity == 0 {
		fmt.Println("Plus de Potion de force. Allez en acheter au marchand.")
		return
	}

	newDamage := &player.Damage
	*newDamage += 10
	fmt.Printf("Vous utilisez une potion de force d'une durée de 2 minutes.\nNouveaux dégats : %d\n", player.Damage)

	go func() {
		time.Sleep(2 * time.Minute)
		player.Damage -= 10
		fmt.Printf("\nL'effet de la potion de force s'est dissipé.\nNouveaux dégats : %d\n", player.Damage)
	}()

	player.Inventory[7].Quantity -= 1

	// if player.Inventory[potIndex].Quantity <= 0 {
	// 	player.Inventory = append(player.Inventory[:potIndex], player.Inventory[potIndex+1:]...)
	// }
}

// func readInt(prompt string) int {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print(prompt)
// 		line, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("Erreur de lecture, réessayez.")
// 			continue
// 		}
// 		line = strings.TrimSpace(line)
// 		if line == "" {
// 			fmt.Println("Choix invalide, veuillez réessayer (Tapez 0, 1 ou 2).")
// 			continue
// 		}
// 		n, err := strconv.Atoi(line)
// 		if err != nil {
// 			fmt.Println("Entrée invalide, veuillez taper un nombre (ex: 1, 2, 0).")
// 			continue
// 		}
// 		return n
// 	}
// }

func (player Character) MainMenu() {
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Printf("\t 1 - Informations du personnage\n")
		fmt.Printf("\t 2 - Inventaire\n")
		fmt.Printf("\t 3 - Hagraah\n")
		fmt.Printf("\t 0 - Quitter\n")

		// var validChoice bool

		// for !validChoice {
		// 	fmt.Println("\nQue souhaitez vous faire ?")
		// 	var userChoice int
		// 	fmt.Scan(userChoice)

		// 	switch userChoice {
		// 	case 1:
		// 		player.displayInfo()
		// 		validChoice = true
		// 	case 2:
		// 		player.AccesInventory()
		// 		validChoice = true
		// 	case 0:
		// 		fmt.Println("\nVous quittez l'aventure.\nMerci pour votre participation !")
		// 		os.Exit(02)
		// 	default:
		// 		fmt.Println("Choix invalide, veuillez réessayer :")
		// 	}
		// }

		userChoice := readInt("\nQue souhaitez vous faire ? ")

		switch userChoice {
		case 1:
			player.displayInfo()
		case 2:
			player.AccesInventory()
		case 0:
			fmt.Println("\nVous quittez l'aventure.\nMerci pour votre participation !\n(Prochaine fois c'est 10 balles si tu veux lancer le jeu)")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer (Tapez 0, 1 ou 2).")
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
