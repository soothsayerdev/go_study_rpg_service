package classes

import (
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
		Name: name,
		Level: level,
		Energy: energy,
		Equipment: equipment,
		Skill: skill,
	}
}
