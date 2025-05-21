package skills

import (
	"go_study_rpg_service/internal/tools"
)

type Spell struct {
	Name     string
	CostMana int
	Weapon   tools.Weapon
	Element  string
}

type Ability struct {
	Name       string
	CostEnergy int
	Weapon     tools.Weapon
}

// Evitar usar ponteiro para interface, pois interface ja é referencia
type SkillMagic interface {
	Execute() string
	GetManaCost() int
}

type SkillPhys interface {
	Execute() string
	GetEnergyCost() int
}
