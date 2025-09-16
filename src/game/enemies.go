package game

type Enemy struct {
	Name   string
	MaxHp  int
	Hp     int
	Damage int
}

func (enemy1 *Enemy) InitEnemy() {
	*enemy1 = Enemy{"Recrue", 50, 50, 5}
}
