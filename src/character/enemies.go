package character

type enemy struct {
	name   string
	maxHp  int
	Hp     int
	damage int
}

func (enemy1 *enemy) initEnemy() {
	*enemy1 = enemy{"Recrue", 100, 80, 5}
}
