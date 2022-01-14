package core

type Link struct {
	x         int
	y         int
	direction string
	name      string
	toLevel   string
}

func (l Link) GetPosition() (int, int) {
	return l.x, l.y
}
