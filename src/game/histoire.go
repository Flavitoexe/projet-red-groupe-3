package game

import "fmt"

func Intro() {
	Logo()
	fmt.Print("\nBienvenue héros, ton voyage commence ici car le monde a besoin de ton courage. \nEn effet, les dieux, autrefois protecteurs de l'humanité, ont été corrompus par une force maléfique. Notre monde est désormais plongé dans le chaos et seul un héros choisi par le destin peut le sauver. Pour cela il lui faudra les affronter et les tuer. C'est le seul moyen de ramener la paix.\n\n")
}

func GardinensArtemis() {
	fmt.Print("\nTu pénètres dans une vaste forêt ancienne, baignée d'une lumière lunaire surnaturelle. Ce territoire sacré est le domaine d'Artémis, déesse de la chasse. Avant d'atteindre la déesse, tu dois affronter deux gardiens.\n\n")
}

func DieuArtémis() {
	fmt.Print("\nDevant toi se tient Artémis, déesse de la chasse, autrefois protectrice. Son regard glacé promet un combat sans merci. Gagneras-tu ?\n\n")
}

func GardiensHéphaïstos() {
	fmt.Print("\nTu pénètres un vaste atelier sacré où résonnent les coups de marteau divins. Ce lieu est dédié à Héphaïstos, dieu du feu et de la forge. Avant d'atteindre la déesse, tu dois affronter deux gardiens.\n\n")
}

func DieuHéphaïstos() {
	fmt.Print("\nDevant toi, Héphaïstos, dieu de la forge, autrefois artisan des merveilles divines. Ses outils brûlants annoncent un combat brutal et sans pitié. Gagneras-tu ?\n\n")
}

func GardiensArès() {
	fmt.Print("\nTu pénètres dans un fort militaire, vaste et imposant, construit sur un plateau battu par le vent. Ce lieu est dédié à Arès, dieu de la guerre. Pour atteindre Arès, tu dois traverser le camp contrôlé par ses gardiens et les affronter.\n\n")
}

func DieuArès() {
	fmt.Print("\nDevant toi se dresse Arès, dieu de la guerre, incarnation de la rage et du combat. Son regard brûle d'une soif de bataille sans fin. Prépare-toi à un affrontement féroce. Bonne chance à toi.\n\n")
}

func GardiensHadès() {
	fmt.Print("\nLe joueur s'enfonce dans les profondeurs sombres du Royaume des morts, un monde ténébreux fait de pierre noire, de flammes froides et de brumes éternelles. Ce territoire funeste est gouverné par Hadès, dieu des enfers. Avant d'atteindre Hadès, tu dois affronter ses gardiens fous.\n\n")
}

func DieuHadès() {
	fmt.Print("\nDevant toi se tient Hadès, souverain des enfers, maître des âmes perdues. Son regard froid et impénétrable plonge dans les ténèbres de ton âme. Le jugement final approche.\n\n")
}

func PremierGardien() {
	GainRound()
	fmt.Print("\nBravo, guerrier. Ce gardien est tombé… mais un autre te défie à son tour.\n\n")
}

func DeuxiemeGardien() {
	GainRound()
	fmt.Print("\nTu as vaincu le dernier gardien... le dieu t'attend.\n\n")
}

func VictoirePerso() {
	FinJeu()
	fmt.Print("Tu as gagné. Par ton courage et ta force, tu as libéré les dieux.\n Grâce à toi, la Terre est sauvée.\n Nous te remercions, héros, car c'est toi qui as changé le destin du monde.\n\n")
}

func PerduCombat() {
	Perdu()
	fmt.Print("\nTu as été vaincu… mais ce n'est que partie remise. À la prochaine, héros déchu.\n")
}

func PremierCombat() {
	fmt.Print("\nLe premier gardien se dresse devant vous !")
}

func VaincuDieu() {
	GainRound()
	fmt.Println("\nBravo, tu as vaincu ce dieu mais ta mission n'est pas terminée...")
}
