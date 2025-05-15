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




type UltimateKnight struct {
}
