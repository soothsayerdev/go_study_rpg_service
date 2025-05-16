package classe

type Knight struct {
	Stamina int
}

func (k *Knight) Attack() int {
	return 10
}

func NewKnight() *Knight {
	return &Knight{
		Stamina: 100,
	}
}
