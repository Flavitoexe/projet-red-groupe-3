package game

// Section Armes (séparée épées + arcs)
var secretBlade Item = Item{"Lame secrète", 0, "Arme"}
var classicSword Item = Item{"Epée classique", 0, "Arme"}
var mediumSword Item = Item{"Epée moyenne", 0, "Arme"}
var plusSword Item = Item{"Epée plus", 0, "Arme"}
var rareSword Item = Item{"Epée rare", 0, "Arme"}
var ultraRareSword Item = Item{"Epée ultra rare", 0, "Arme"}

var bow Item = Item{"Arc", 0, "Arme"}
var eyeOfDawnBow Item = Item{"Arc Oeil de l'Aube", 0, "Arme"}

// Section Défenses
var shield Item = Item{"Bouclier unique", 0, "Armure"}
var uniqueShield Item = Item{"Bouclier unique", 0, "Armure"}
var spartanShield Item = Item{"Bouclier Spartiate", 0, "Armure"}
var chapeauAventurier Item = Item{"Chapeau de l'aventurier", 0, "Armure"}
var tuniqueAventurier Item = Item{"Tunique de l'aventurier", 0, "Armure"}
var BottesAventurier Item = Item{"Bottes de l'aventurier", 0, "Armure"}

// Section Consommables
var healthPot Item = Item{"Potion de vie", 0, "Cons"}

// var neuilPot Item = Item{"Potion de neuil", 0, "Cons"}
var strenghtPot Item = Item{"Potion de force", 0, "Cons"}
var arrow Item = Item{"Flèches classiques", 0, "Cons"}
var poisonedArrow Item = Item{"Flèches empoisonnées", 0, "Cons"}

// Inventaire du joueur
var inventory []Item = []Item{secretBlade, classicSword, mediumSword, plusSword, rareSword, ultraRareSword, bow, eyeOfDawnBow,
	shield, uniqueShield, spartanShield,
	healthPot, strenghtPot, arrow, poisonedArrow}

// Maps et Arrays pour le shop
var weaponMap map[Item]int = map[Item]int{classicSword: 15, mediumSword: 20, rareSword: 35, ultraRareSword: 60, eyeOfDawnBow: 80}
var weaponArr []Item = []Item{classicSword, mediumSword, rareSword, ultraRareSword, eyeOfDawnBow}

var defenseMap map[Item]int = map[Item]int{uniqueShield: 30, spartanShield: 60}
var defenseArr []Item = []Item{uniqueShield, spartanShield}

var consMap map[Item]int = map[Item]int{healthPot: 10, strenghtPot: 15, arrow: 5, poisonedArrow: 10}
var consArr []Item = []Item{healthPot, strenghtPot, arrow, poisonedArrow}
