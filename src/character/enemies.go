package character

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
type Enemy struct {
	Name   string
	MaxHp  int
	Hp     int
	Damage int
}

func (enemy1 *Enemy) initEnemy() {
	*enemy1 = Enemy{"Recrue", 100, 80, 5}
}
