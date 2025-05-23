package classes

import (
	"fmt"
	"go_study_rpg_service/internal/skills"
	"go_study_rpg_service/internal/tools"
)

type Mage struct {
	Name      string
	Level     int
	Manapool  int
	Equipment tools.Weapon
	Skill     skills.SkillMagic
}

func NewMage(name string, level int, manapool int, equipment tools.Weapon, skill skills.SkillMagic) *Mage {
	return &Mage{
		Name:      name,
		Level:     level,
		Manapool:  manapool,
		Equipment: equipment,
		Skill:     skill,
	}
}

func (m *Mage) UseSkill() string {
	if m.Manapool >= m.Skill.GetManaCost() {
		m.Manapool -= m.Skill.GetManaCost()

		return fmt.Sprintf("%s %d (Mana remaining: %d)", m.Name, m.Skill.GetManaCost(), m.Manapool)
	}
	return "Not enough mana!"
}

func (m *Mage) GetClassName() string {
	return "Mage"
}
