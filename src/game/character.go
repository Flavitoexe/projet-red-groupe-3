package game

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

func (player *Character) initCharacter() {
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
		Skill:  []string{"Coup de poing"}}
}

func (player *Character) characterCreation() {
	var validName bool
	for !validName {
		fmt.Println("Choisissez un nom pour votre perso.")
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
	fmt.Println("Voulez vous utiliser un objet ?\n\t1 - 9 : Utilise l'objet\n\t0 : Quitter l'inventaire")
	fmt.Scan(&userChoice)
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
	fmt.Printf("Vous utilisez une potion de soin.\nNouveau Hp : %d\n", player.Hp)

	player.Inventory[potIndex].Quantity -= 1

	if player.Inventory[potIndex].Quantity <= 0 {
		player.Inventory = append(player.Inventory[:potIndex], player.Inventory[potIndex+1:]...)
	}
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
			fmt.Println("Veuillez entrer un nombre.")
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Entrée invalide — tapez un nombre (ex: 1, 2, 0).")
			continue
		}
		return n
	}
}

func (player Character) MainMenu() {
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Printf("\t 1 - Informations du personnage\n")
		fmt.Printf("\t 2 - Inventaire\n")
		fmt.Printf("\t 0 - Quitter\n")

		userChoice := readInt("à vous: ")

		switch userChoice {
		case 1:
			player.displayInfo()
		case 2:
			player.AccesInventory()
		case 0:
			fmt.Println("Merci au revoir.")
			return
		default:
			fmt.Println("Votre choix n'est pas valide.")
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

func main() {
	p1 := Character{}
	p1.initCharacter()
	p1.characterCreation()
	p1.displayInfo()
	p1.MainMenu()
}
