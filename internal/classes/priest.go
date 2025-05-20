package classes

import (
	"go_study_rpg_service/internal/skills"
	"go_study_rpg_service/internal/tools"
)

type Priest struct {
	Name      string
	Level     int
	Manapool  int
	Equipment tools.Weapon
	Skill     skills.SkillMagic
}

func NewPadre(name string, level int, manapool int, equipment tools.Weapon, skill skills.SkillMagic) *Priest {
	return &Priest{
		Name: name,
		Level: level,
		Manapool: manapool,
		Equipment: equipment,
		Skill: skill,
	}
}

func (p *Priest) Heal() int {
	if p.Manapool < 20 {
		return 0
	}
	p.Manapool -= 20
	return 50
}
