package classes

import (
	"go_study_rpg_service/internal/skills"
	"go_study_rpg_service/internal/tools"
)

type Knight struct {
	Name      string
	Level     int
	Energy    int
	Equipment tools.Weapon
	Skill skills.SkillPhys
}

func NewKnight(name string, level int, energy int, equipment tools.Weapon, skill skills.SkillPhys) *Knight{
	return &Knight{
		Name: name,
		Level: level,
		Energy: energy,
		Equipment: equipment,
		Skill: skill,
	}
}