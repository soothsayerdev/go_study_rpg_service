package classes

import "go_study_rpg_service/internal/tools"

type Archer struct {
	Name      string
	Level     int
	Energy    int
	Equipment tools.Weapon
}
