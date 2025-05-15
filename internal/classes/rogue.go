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
	Skill skills.Skill
}

func (r *Rogue) UseSkill(skill skills.Skill) string {
	return fmt.Sprintf("%s performs a sneaky %s", r.Name, skill.Execute())
}

func (r *Rogue) GetClassName() string {
	return "Rogue"
}