package classes

import (
	"fmt"
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
		Name:      name,
		Level:     level,
		Manapool:  manapool,
		Equipment: equipment,
		Skill:     skill,
	}
}

func (p *Priest) UseSkill() string {
	if p.Manapool >= p.Skill.GetManaCost() {
		p.Manapool -= p.Skill.GetManaCost()

		return fmt.Sprintf("%s %d (Mana remaining: %d)", p.Name, p.Skill.GetManaCost(), p.Manapool)
	}
	return "Not enough mana!"
}

func (p *Priest) Heal() int {
	if p.Manapool < 20 {
		return 0
	}
	p.Manapool -= 20
	return 50
}

func (p *Priest) GetClassName() string {
	return p.Name
}
