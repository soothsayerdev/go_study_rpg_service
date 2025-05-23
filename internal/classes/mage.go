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

<<<<<<< HEAD
func (m *Mage) UseSkill(skill skills.SkillMagic) string {
	if m.Manapool >= skill.GetManaCost() {
		m.Manapool -= skill.GetManaCost()
=======
func NewMage(name string, level int, manapool int, equipment tools.Weapon, skill skills.SkillMagic) *Mage {
	return &Mage{
		Name:      name,
		Level:     level,
		Manapool:  manapool,
		Equipment: equipment,
		Skill:     skill,
	}
}
>>>>>>> 3c00d569d9563e921d37ebc65d02139138aaca48

func (m *Mage) UseSkill() string {
	if m.Manapool >= m.Skill.GetManaCost() {
		m.Manapool -= m.Skill.GetManaCost()

		return fmt.Sprintf("%s %d (Mana remaining: %d)", m.Name, m.Skill.GetManaCost(), m.Manapool)
	}
	return "Not enough mana!"
}

<<<<<<< HEAD
func (m *Mage) Execute() string {
	if m.Manapool >= skills.Spell.CostMana {
		m.Manapool -= m.Skill.GetManaCost()

		return fmt.Sprintf("%s %s (Mana remaining: %d)", m.Name, skills.SkillMagic.GetManaCost(m.Skill), m.Manapool)
	}
	return "Not enough mana!"	
}

=======
>>>>>>> 3c00d569d9563e921d37ebc65d02139138aaca48
func (m *Mage) GetClassName() string {
	return "Mage"
}
