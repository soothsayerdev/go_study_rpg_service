package classe

type Mage struct {
	Mana int
}

func (k *Mage) Attack() int {
	return 50
}

func NewMage() *Mage {
	return &Mage{
		Mana: 100,
	}
}
