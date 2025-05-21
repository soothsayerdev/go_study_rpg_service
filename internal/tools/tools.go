package tools

type Weapon interface {
	Attack() string
	DamageBonus() int
	GetName() string
}
