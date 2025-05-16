package classe

type Priest struct {
	Mana int
}

func (k *Priest) Heal() int {
	return 50
}

func NewPriest() *Priest {
	return &Priest{
		Mana: 100,
	}
}
