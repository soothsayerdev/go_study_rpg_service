package classes

import "go_study_rpg_service/internal/tools"

type Knight struct {
	Name      string
	Level     int
	Energy    int
	Equipment tools.Weapon
}