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
	Skill     skills.Skill
}

func (m *Mage) UseSkill(skill skills.Skill) string {
	if m.Manapool >= skill.GetManaCost() {
		m.Manapool -= skill.GetManaCost()

		return fmt.Sprintf("%s %s (Mana remaining: %d)", m.Name, skill., m.Manapool)
	}
	return "Not enough mana!"
}

func (m *Mage) GetClassName() string {
	return "Mage"
}