package classes

import (
	"go_study_rpg_service/internal/skills"
)

type Character interface {
	UseSkill(skill skills.Skill) string
	GetClassName() string
}

type UltimateKnight struct {
}
