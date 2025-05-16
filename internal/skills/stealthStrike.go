package skills

import "fmt"

type StealthStrike struct {
	Ability        *Ability
	CriticalChance float64
}

func (s *StealthStrike) Execute() string {
	return fmt.Sprintf("Casting %s with %s - Critical Chance : %d - Cost: %d",
		s.Ability.Name, s.Ability.Weapon.Attack(), s.CriticalChance, s.Ability.CostEnergy)
}

func (s *StealthStrike) GetManaCost() int {
	return s.Ability.CostEnergy
}
