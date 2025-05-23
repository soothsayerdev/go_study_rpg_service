package classes

import (
	"fmt"
	"go_study_rpg_service/internal/skills"
	"go_study_rpg_service/internal/tools"
)

type Archer struct {
	Name      string
	Level     int
	Energy    int
	Equipment tools.Weapon
	Skill     skills.SkillPhys
}

func NewArcher(name string, level int, energy int, equipment tools.Weapon, skill skills.SkillPhys) *Archer {
	return &Archer{
		Name:      name,
		Level:     level,
		Energy:    energy,
		Equipment: equipment,
		Skill:     skill,
	}
}

func (a *Archer) UseSkill() string {
	if a.Energy >= a.Skill.GetEnergyCost() {
		a.Energy -= a.Skill.GetEnergyCost()

		return fmt.Sprintf("%s %d (Energy remaining: %d)", a.Name, 
		a.Skill.GetEnergyCost(), a.Energy)
	}
	return "Not enough energy!"
}

func (a *Archer) GetClassName() string {
	return a.Name
}
