package skills

import (
	"fmt"
	"go_study_rpg_service/internal/tools"
)

type Spell struct {
	Name    string
	Cost    int
	Weapon  tools.Weapon
	Element string
}

type Skill interface {
	Execute() string
	GetManaCost() int
}

func (s *Skill) Execute() string {
	return fmt.Sprintf("Casting %s using %s - Cost: %d mana",
		s.Name, s.Weapon.Attack(), s.Cost)
}

func (s *Spell) GetManaCost() int {
	return s.Cost
}

type FireballSpell struct {
	Spell *Spell
	BurnDamage int
}

type StealthStrike struct {
	Spell *Spell
	CriticalChance float64
}

type HolySmite struct {
	Spell *Spell
	HealAmount int
}

// type Tool interface {
// 	Action()
// }

// type tools struct {
// 	Hand hand
// 	Staff Staff
// 	Wand wand
// }

// func (spell *Spell) CostSpell() string {
// 	return fmt.Sprintf("I casting %s using %s mana", spell.Name, spell.Cost)
// }

// func Action() Tool {
// 	return fmt.Sprintf("I casting with ...")
// }
