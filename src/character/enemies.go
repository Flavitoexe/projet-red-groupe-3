package character

type Enemy struct {
	Name   string
	MaxHp  int
	Hp     int
	Damage int
}

func (enemy1 *Enemy) initEnemy() {
	*enemy1 = Enemy{"Recrue", 100, 80, 5}
}
