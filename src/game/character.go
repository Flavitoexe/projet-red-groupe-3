package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Character struct {
	Name       string
	Class      string
	Level      int
	Hp         int
	HpMax      int
	Inventory  []Item
	Damage     int
	Skill      []string
	Money      int
	Equipment  []Equipment
	BowState   int
	Infight    bool
	Resurected bool
}

type Item struct {
	Name     string
	Quantity int
	Tag      string
}

type Equipment struct {
	Casque     Item
	Body       Item
	Foot       Item
	IsEquipped bool
}

var class1 Character = Character{
	Name:      "",
	Class:     "Guerrier",
	Level:     1,
	Hp:        50,
	HpMax:     100,
	Inventory: inventory,
	Damage:    10,
	Skill:     []string{"Coup de poing"},
	Money:     100,
}

var class2 Character = Character{
	Name:      "",
	Class:     "Guerrier léger",
	Level:     1,
	Hp:        40,
	HpMax:     80,
	Inventory: inventory,
	Damage:    15,
	Skill:     []string{"Coup de poing"},
	Money:     100,
}

var class3 Character = Character{
	Name:      "",
	Class:     "Guerrier lourd",
	Level:     1,
	Hp:        60,
	HpMax:     120,
	Inventory: inventory,
	Damage:    5,
	Skill:     []string{"Coup de poing"},
	Money:     100,
}

var class4 Character = Character{
	Name:      "",
	Class:     "cheat",
	Level:     1,
	Hp:        1000,
	HpMax:     1000,
	Inventory: inventory,
	Damage:    250,
	Skill:     []string{"Coup de poing"},
	Money:     10000,
}

func (player *Character) classChoice() {
	var validClass bool
	for !validClass {

		fmt.Println(Underline + Bold + " Voici les différentes classes :" + Reset)
		fmt.Printf(Bold+"\n• Classe 1 : %s"+Reset, class1.Class)
		fmt.Printf("\n\tPv : %d\n\tDégats : %d\n\tCapacités : %s\n", class1.HpMax, class1.Damage, class1.Skill[0])
		fmt.Println("Guerrier polyvalent, a un bon nombre de points de vie et inflige de bons dégats.\n")

		fmt.Printf(Bold+"\n• Classe 2 : %s"+Reset, class2.Class)
		fmt.Printf("\n\tPv : %d\n\tDégats : %d\n\tCapacités : %s\n", class2.HpMax, class2.Damage, class2.Skill[0])
		fmt.Println("Guerrier léger, a moins de points de vie que sa version classique, mais inflige plus de dégats.\n")

		fmt.Printf(Bold+"\n• Classe 3 : %s "+Reset, class3.Class)
		fmt.Printf("\n\tPv : %d\n\tDégats : %d\n\tCapacités : %s\n", class3.HpMax, class3.Damage, class3.Skill[0])
		fmt.Println("Guerrier lourd, a plus de points de vie que ses autres versions, mais inflige moins de dégats.\n")

		var classChoice int
		fmt.Println(Blue + "\nFaites votre choix (1, 2 ou 3): \n " + Reset)
		fmt.Scan(&classChoice)

		switch classChoice {
		case 1:
			*player = class1
			validClass = true
			return
		case 2:
			*player = class2
			validClass = true
			return
		case 3:
			*player = class3
			validClass = true
		case 3003:
			*player = class4
			validClass = true
		default:
			fmt.Println(Red + "\nChoix invalide, veuillez réessayer !\n" + Reset)
			return
		}
	}
}

func (player *Character) CharacterCreation() {
	player.classChoice()
	var validName bool
	for !validName {
		fmt.Println(Blue + "\nChoisissez un nom pour votre perso : " + Reset)
		var nameChoice string = ""
		fmt.Scan(&nameChoice)

		for i := 0; i < len(nameChoice); i++ {
			if !((65 <= nameChoice[i] && nameChoice[i] <= 90) || (97 <= nameChoice[i] && nameChoice[i] <= 122)) {
				fmt.Println(Red + "Oups erreur, votre nom n'est pas correct, veuillez recommencer." + Reset)
				validName = false
				player.Name = ""
				break
			} else {
				if (65 <= nameChoice[i] && nameChoice[i] <= 90) && i != 0 {
					player.Name += string(nameChoice[i] + 32)
					validName = true
					break
				} else if (97 <= nameChoice[i] && nameChoice[i] <= 122) && i != 0 {
					player.Name += string(nameChoice[i])
					validName = true
					break
				} else if (97 <= nameChoice[i] && nameChoice[i] <= 122) && i == 0 {
					validName = true
					player.Name += string(nameChoice[i] - 32)
					break
				} else if (65 <= nameChoice[i] && nameChoice[i] <= 90) && i == 0 {
					validName = true
					player.Name += string(nameChoice[i])
					break
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
	fmt.Printf("\t - Argent : %d\n", player.Money)

}

func readInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(Red + "Erreur de lecture, réessayez.\n" + Reset)
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			fmt.Println(Red + "Choix invalide, veuillez réessayer.\n" + Reset)
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(Red + "Entrée invalide, veuillez taper un nombre.\n" + Reset)
			continue
		}
		return n
	}
}

func (player *Character) limitInventory(item Item) {
	if !(item.Quantity <= 10) {
		fmt.Println(Red + "Oups, vous ne pouvez pas avoir d'autres objets.\n" + Reset)
		return
	}
}

func (player *Character) AddInventory(item Item, quantity int) {
	// player.limitInventory(item)
	if item.Quantity > 0 && !(item.Tag == "Cons") {
		fmt.Println("Vous avez déjà cet objet.\n")
		return
	}

	for index := range player.Inventory {
		if player.Inventory[index].Name == item.Name {
			player.Inventory[index].Quantity += quantity
			fmt.Printf("\nQuantité de %s : %d\n", item.Name, player.Inventory[index].Quantity)
			return
		}
	}
}

func (player *Character) RemoveInventory(item Item) {

	for index := range player.Inventory {
		// if player.Inventory[index].Quantity >= 0 && player.Inventory[index].Quantity-1 >= 0 {
		if player.Inventory[index].Quantity != 0 && player.Inventory[index].Name == item.Name {
			player.Inventory[index].Quantity--
			fmt.Printf("Nouvelle quantité de %s : %d\n", item.Name, player.Inventory[index].Quantity)
			return
		}
	}
	fmt.Printf("Vous n'avez pas de %s.\n", item.Name)
}

func (player *Character) AccesInventory() {
	var leave bool
	for !leave {
		fmt.Println("\n=== Inventaire du personnage ===")
		fmt.Println("\t1 == Armes ==\n\t2 == Armure ==\n\t3 == Consommables ==\n\t4 == Matériaux ==\n\t0 - Quitter\n")
		userChoice := readInt(Blue + "Que souhaitez vous voir ?  " + Reset)

		switch userChoice {

		case 0:
			fmt.Println("Vous quittez l'inventaire.\n")
			leave = true
		case 1:
			fmt.Println("== Armes ==")
			var temp int

			for _, weapon := range player.Inventory {
				if weapon.Tag != "Arme" || weapon.Quantity == 0 {
					continue
				} else if weapon.Tag == "Arme" {
					temp++
					fmt.Printf("\t - %s\n", weapon.Name)
				}
			}

			if temp == 0 {
				fmt.Println("\n\t Inventaire des armes vide\n")

				userChoice2 := readInt("0 - Quitter\n")
				if userChoice2 == 0 {
					player.AccesInventory()
					leave = true
				}

				if userChoice2 == 0 {
					player.AccesInventory()
					leave = true
					return
				} else {
					fmt.Println(Red + "\nVous avez pas le choix. Vous quittez un point c'est tout." + Reset)
					player.AccesInventory()
					leave = true
					return
				}
			}

		case 2:
			fmt.Println("== Armures ==")
			var temp int

			for _, armor := range player.Inventory {
				if armor.Tag != "Armure" || armor.Quantity == 0 {
					continue
				} else if armor.Tag == "Armure" {
					temp++
					fmt.Printf("\t - %s\n", armor.Name)
				}
			}

			if temp == 0 {
				fmt.Println("\n\t Inventaire des défenses vide\n")

				userChoice2 := readInt("0 - Quitter\n")
				if userChoice2 == 0 {
					player.AccesInventory()
					leave = true

					if userChoice2 == 0 {
						player.AccesInventory()
						leave = true
						return
					} else {
						fmt.Println(Red + "\nVous avez pas le choix. Vous quittez un point c'est tout." + Reset)
						player.AccesInventory()
						leave = true
						return
					}
				}
			}

		case 3:
			fmt.Println("== Consommables ==")
			var temp int
			// var cpt int = 1

			for _, cons := range player.Inventory {
				if cons.Tag != "Cons" || cons.Quantity == 0 {
					continue
				} else if cons.Tag == "Cons" {
					temp++
					fmt.Printf("\t- %s x %d\n", cons.Name, cons.Quantity)
				}
			}

			if temp == 0 {
				fmt.Println("\n\t Inventaire des consommables vide\n")
				userChoice2 := readInt("0 - Quitter\n")

				if userChoice2 == 0 {
					player.AccesInventory()
					leave = true
					return
				} else {
					fmt.Println(Red + "\nVous avez pas le choix. Vous quittez un point c'est tout." + Reset)
					player.AccesInventory()
					leave = true
					return
				}
			}

			fmt.Println(Blue + "\nVoulez vous utiliser un objet ?\n" + Reset + "(0 pour sortir, 1 - 2 pour consommer une potion, 3 - 4 pour seléctionner un type de flèche)\n")
			userChoice2 := readInt(Blue + "Votre choix : \n" + Reset)

			switch userChoice2 {

			case 0:
				player.AccesInventory()
				leave = true

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
			default:
				fmt.Println(Red + "\nChoix invalide, taper un nombre proposé.\n" + Reset)
			}

		case 4:
			fmt.Println("== Matériaux ==")
			var temp int

			for _, misc := range player.Inventory {
				if misc.Tag != "Misc" || misc.Quantity == 0 {
					continue
				} else if misc.Tag == "Misc" {
					temp++
					fmt.Printf("\t- %s x %d\n", misc.Name, misc.Quantity)
				}
			}

			if temp == 0 {
				fmt.Println("\n\t Inventaire des matériaux vide\n")
				userChoice2 := readInt("0 - Quitter\n")

				if userChoice2 == 0 {
					player.AccesInventory()
					leave = true
					return
				} else {
					fmt.Println(Red + "\nVous avez pas le choix. Vous quittez un point c'est tout." + Reset)
					player.AccesInventory()
					leave = true
					return
				}
			}

		default:
			fmt.Println(Red + "\nChoix invalide, veuillez taper un nombre proposé." + Reset)
		}
	}
}

func (player *Character) takeHealthPot() {

	for index := range player.Inventory {
		if player.Inventory[index].Name == "Potion de vie" {
			if player.Inventory[index].Quantity < 1 {
				fmt.Println("Plus de Potion de vie. Allez en acheter au marchand.")
				return
			}

			player.Hp += 20
			if player.Hp > player.HpMax {

				player.Hp = player.HpMax
			}

			fmt.Printf("Vous venez de consommer une potion de soins !\nNouveaux points de vie : %d.\n", player.Hp)
			player.RemoveInventory(healthPot)
			break
		}
	}
}

func (player *Character) takeStrenghtPot() {

	for index := range player.Inventory {
		if player.Inventory[index].Name == "Potion de force" {
			if player.Inventory[index].Quantity < 1 {
				fmt.Println("Plus de Potion de force. Allez en acheter au marchand.")
				return
			}

			newDamage := &player.Damage
			*newDamage += 10
			fmt.Printf("Vous consommez une potion de force d'une durée de 2 minutes.\nNouveaux dégats : %d\n", player.Damage)

			go func() {
				time.Sleep(15 * time.Second)
				player.Damage -= 10
				fmt.Printf("\nL'effet de la potion de force s'est dissipé.\nNouveaux dégats : %d\n", player.Damage)
			}()
			player.RemoveInventory(strenghtPot)
			break
		}
	}
}

func (player Character) MainMenu() {
	var validChoice bool
	for !validChoice {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Printf("\t 1 - Informations du personnage\n")
		fmt.Printf("\t 2 - Inventaire\n")
		fmt.Printf("\t 3 - Marchand\n")
		fmt.Printf("\t 4 - Forgeron\n")
		fmt.Printf("\t 5 - Entraînement\n")
		fmt.Printf("\t 0 - Quitter\n")

		var menuChoice int
		fmt.Println(Blue + "\nFaites votre choix (0, 1, 2, 3, 4 ou 5): " + Reset)
		fmt.Scan(&menuChoice)

		switch menuChoice {
		case 1:
			player.displayInfo()
		case 2:
			player.AccesInventory()
		case 3:
			player.AccessShop()
		case 4:
			player.MenuForgeron()
		case 5:
			TrainingFight(&player)
		case 0:
			fmt.Println(Magenta + "\nVous quittez l'aventure.\nMerci pour votre participation !\n(Prochaine fois c'est 10 balles si tu veux lancer le jeu)\n" + Reset)
			return
		default:
			fmt.Println(Red + "Choix invalide, veuillez réessayer (tapez une des propositions ci-dessus)." + Reset)
		}
	}
}

func (player *Character) isDead() {
	if player.Hp == 0 && !player.Resurected {
		fmt.Println(Cyan + "\n\nTu es mort, repenses-y à deux fois la prochaine fois.\n" + Reset)
		fmt.Println(Green + "Tu viens de ressusciter, bonne chance à toi !" + Reset)
		player.Hp = player.HpMax / 2
		fmt.Printf("Ton nouveau Hp est : %d\n", player.Hp)
	} else if player.Hp == 0 && player.Resurected {
		fmt.Println("C'est fini ciao")
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

// func (player Character) MenuForgeron() {
// 	var valid bool
// 	for !valid {
// 		fmt.Println("\n=== Menu Forgeron ===")
// 		fmt.Println("\t1 - Linothorax du Guerrier d'Olympe")
// 		fmt.Println("\t2 - Cuirasse de Fer des Hoplites Éternels")
// 		fmt.Println("\t3 - Égide d'Or d'Hélios")
// 		fmt.Println("\t0 - Quitter")
// 		fmt.Println("\nQuel set d'armure souhaitez vous faire fabriquer ?")

// 		userChoice := readInt("")

// 		switch userChoice {
// 		case 0:
// 			player.MainMenu()
// 			valid = true
// 		case 1:
// 			fmt.Println("== Linothorax du Guerrier d'Olympe ==")
// 			fmt.Println("\t1 - Heaume du Regard de l'Hydre")
// 			fmt.Printf("\t\t7 Cuirs, 15 H\n")
// 			fmt.Println("\t2 - Linothorax du Sang de l'Hydre")
// 			fmt.Printf("\t\t10 Cuirs, 20 H\n")
// 			fmt.Println("\t3 - Grèves de la Bête Serpentine")
// 			fmt.Printf("\t\t5 Cuirs, 10 H\n")
// 			fmt.Println("\t0 - Quitter")

// 			fmt.Println("Quelle pièce souhaitez vous faire fabriquer ?")
// 			userChoice2 := readInt("")
// 		}

// 		fmt.Printf("\t 1 - Chapeau de l'aventurier\n")
// 		fmt.Printf("\t 2 - Tunique de l'aventurier\n")
// 		fmt.Printf("\t 3 - Bottes de l'aventurier\n")
// 		fmt.Printf("\t 4 - Inventaire\n")
// 		fmt.Printf("\t 0 - Quitter\n")

// 		var menuForgeron int
// 		fmt.Println(Blue + "\nFaites votre choix (0, 1, 2, 3 ou 4): " + Reset)
// 		fmt.Scan(&menuForgeron)

// 		switch menuForgeron {
// 		case 1:
// 			if player.Money >= 5 {
// 				player.Money -= 5
// 				player.AddInventory(adventurerCap, 1)
// 				fmt.Println(Green + "\nYoupi! Le forgeron a fabriqué le chapeau de l'aventurier.\n" + Reset)

// 			} else {
// 				fmt.Print(Red + "Oups, vous n'avez pas assez d'argent pour le fabriquer !" + Reset)
// 			}
// 		case 2:
// 			if player.Money >= 5 {
// 				player.Money -= 5
// 				player.AddInventory(adventurerChest, 1)
// 				fmt.Println(Green + "\nYoupi! Le forgeron a fabriqué la tunique de l'aventurier.\n" + Reset)
// 			} else {
// 				fmt.Print(Red + "Oups, vous n'avez pas assez d'argent pour le fabriquer !" + Reset)
// 			}
// 		case 3:
// 			if player.Money >= 5 {
// 				player.Money -= 5
// 				player.AddInventory(adventurerBoots, 1)
// 				fmt.Println(Green + "\nYoupi! Le forgeron a fabriqué les bottes de l'aventurier.\n" + Reset)
// 			} else {
// 				fmt.Print(Red + "Oups, vous n'avez pas assez d'argent pour le fabriquer !" + Reset)
// 			}
// 		case 4:
// 			player.AccesInventory()
// 		case 0:
// 			return
// 		default:
// 			fmt.Println(Red + "\nChoix invalide, veuillez réessayer !\n" + Reset)
// 		}
// 	}
// }

// func (player Character) IsEquipped()

func (player Character) upgradeInventorySlot() {

}

const (
	Reset     = "\033[0m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
)
