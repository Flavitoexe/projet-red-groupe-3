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
	Money     int
	BowState  int
	Infight   bool
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

func (player *Character) classChoice() {
	var validClass bool
	for !validClass {

		fmt.Println("• Voici les différentes classes :")
		fmt.Printf("\nClasse 1 : %s", class1.Class)
		fmt.Printf("\n\tPv : %d\n\tDégats : %d\n\tCapacités : %s\n", class1.HpMax, class1.Damage, class1.Skill[0])
		fmt.Println("Guerrier polyvalent, a un bon nombre de points de vie et inflige de bons dégats.\n")

		fmt.Printf("\nClasse 2 : %s", class2.Class)
		fmt.Printf("\n\tPv : %d\n\tDégats : %d\n\tCapacités : %s\n", class2.HpMax, class2.Damage, class2.Skill[0])
		fmt.Println("Guerrier léger, a moins de points de vie que sa version classique, mais inflige plus de dégats.\n")

		fmt.Printf("\nClasse 3 : %s ", class3.Class)
		fmt.Printf("\n\tPv : %d\n\tDégats : %d\n\tCapacités : %s\n", class3.HpMax, class3.Damage, class3.Skill[0])
		fmt.Println("Guerrier lourd, a plus de points de vie que ses autres versions, mais inflige moins de dégats.\n")

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
			fmt.Println("\nChoix invalide, veuillez réessayer !\n")
		}
	}
}

func (player *Character) CharacterCreation() {
	player.classChoice()
	var validName bool
	for !validName {
		fmt.Println("\nChoisissez un nom pour votre perso : ")
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
	fmt.Printf("\t - Argent : %d\n", player.Money)

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

	if item.Quantity > 0 && !(item.Tag == "Cons") {
		fmt.Println(" item.Quantity > 0 && !(item.Tag == Cons)")
		fmt.Println("Vous avez déjà cet objet.")
		return
	}

	for index := range player.Inventory {
		if player.Inventory[index].Name == item.Name {
			player.Inventory[index].Quantity++
			fmt.Printf("\nQuantité de %s : %d\n", item.Name, player.Inventory[index].Quantity)
			return
		}
	}
}

func (player *Character) RemoveInventory(item Item) {

	for index := range player.Inventory {
		if player.Inventory[index].Quantity >= 0 && player.Inventory[index].Quantity-1 >= 0 {
			player.Inventory[index].Quantity--
			fmt.Printf("Nouvelle quantité de %s : %d\n", item.Name, player.Inventory[index].Quantity)
			return
		}
	}
	fmt.Printf("Vous n'avez pas de %s.\n", item.Name)
}

func (player Character) AccesInventory() {

	var leave bool
	for !leave {
		fmt.Println("\n=== Inventaire du personnage ===")
		fmt.Println("\t1 == Armes ==\n\t2 == Armure ==\n\t3 == Consommables ==\n\t0 - Quitter\n")
		userChoice := readInt("Que souhaitez vous voir ?  ")

		switch userChoice {

		case 0:
			fmt.Println("Vous quittez l'inventaire.")
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
			}

			userChoice2 := readInt("0 - Quitter\n")
			if userChoice2 == 0 {
				player.AccesInventory()
				leave = true
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
			}

			userChoice2 := readInt("0 - Quitter\n")
			if userChoice2 == 0 {
				player.AccesInventory()
				leave = true
			}

		case 3:
			fmt.Println("== Consommables ==")
			var temp int

			for _, cons := range player.Inventory {
				if cons.Tag != "Cons" || cons.Quantity == 0 {
					continue
				} else if cons.Tag == "Cons" {
					temp++
					fmt.Printf("\t - %s x %d\n", cons.Name, cons.Quantity)
				}
			}

			if temp == 0 {
				fmt.Println("\n\t Inventaire des consommables vide\n")
				userChoice2 := readInt("0 - Quitter\n")
				if userChoice2 == 0 {
					player.AccesInventory()
					leave = true
					return
				}
			}

			fmt.Println("\nVoulez vous utiliser un objet ?\n(0 pour sortir, 1 - 2 pour consommer une potion, 3 - 4 pour seléctionner un type de flèche)\n")
			userChoice2 := readInt("Votre choix : ")

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
			}
		}
	}
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

	player.RemoveInventory(healthPot)
	// player.Inventory[potIndex].Quantity -= 1

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

func (player Character) MainMenu() {
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("\t 1 - Informations du personnage")
		fmt.Println("\t 2 - Inventaire")
		fmt.Println("\t 3 - Marchand")
		fmt.Println("\t 4 - Entraînement")
		fmt.Println("\t 0 - Quitter")

		userChoice := readInt("\nQue souhaitez vous faire ? ")

		switch userChoice {
		case 1:
			player.displayInfo()
		case 2:
			player.AccesInventory()
		case 3:
			player.AccessShop()
		case 4:
			TrainingFight(&player)
		case 0:
			fmt.Println("\nVous quittez l'aventure.\nMerci pour votre participation !\n(Prochaine fois c'est 10 balles si tu veux lancer le jeu)")
			return
		default:
			fmt.Println("\nChoix invalide, veuillez réessayer (Tapez 0, 1 ou 2).")
		}
	}
}

func (player *Character) isDead() {
	if player.Hp == 0 {
		fmt.Println("\n\nTu es mort, repenses-y à deux fois la prochaine fois.\n")
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
