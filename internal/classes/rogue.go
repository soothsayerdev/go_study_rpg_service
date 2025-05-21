package classes

import (
	"fmt"
	"go_study_rpg_service/internal/skills"
	"go_study_rpg_service/internal/tools"
)

type Rogue struct {
	Name      string
	Level     int
	Energy    int
	Equipment tools.Weapon
	Ability   skills.Ability
}

func NewRogue(name string, level int, energy int, equipment tools.Weapon, ability skills.Ability) *Rogue {
	return &Rogue{
		Name: name,
		Level: level,
		Energy: energy,
		Equipment: equipment,
		Ability: ability,
	}
}

func (r *Rogue) UseAbility(skill skills.SkillPhys) string {
	if r.Energy < r.Ability.CostEnergy {
		return "Not enough energy to cast!"
	}
	return fmt.Sprintf("%s performs a sneaky %s", r.Name, skill.Execute())
}

func (r *Rogue) GetClassName() string {
	return "Rogue"
}
