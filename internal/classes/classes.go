package classes

import (
	"fmt"
	"go_study_rpg_service/internal/skills"
	"go_study_rpg_service/internal/tools"
)

type Character interface {
	UseSkill(skill skills.Skill) string
	GetClassName() string
}

type Mage struct {
	Name      string
	Level     int
	Manapool  int
	Equipment tools.Weapon
	Skill 	skills.Skill
}

func (m *Mage) UseSkill(skill skills.Skill) string {
	if m.Manapool >= skill.GetManaCost() {
		m.Manapool -= skill.GetManaCost()

		return fmt.Sprintf("%s %s (Mana remaining: %d)", m.Name, skill.Execute(), m.Manapool)
	}
	return "Not enough mana!"
}

func (m *Mage) GetClassName() string {
	return "Mage"
}

type Rogue struct {
	Name      string
	Level     int
	Energy    int
	Equipment tools.Weapon
}

func (r *Rogue) UseSkill(skill skills.Skill) string {
	return fmt.Sprintf("%s performs a sneaky %s", r.Name, skill.Execute())
}

func (r *Rogue) GetClassName() string {
	return "Rogue"
}

type Knight struct {
	Name      string
	Level     int
	Energy    int
	Equipment tools.Weapon
}

type Archer struct {
	Name  string 
	Level int
	Energy int
	Equipment tools.Weapon
}

type UltimateKnight struct {
}
