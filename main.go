package main

import (
	"fmt"
	"go_study_rpg_service/internal/tools"
    "go_study_rpg_service/internal/skills"
    "go_study_rpg_service/internal/classes"
)

func main() {
	fireStaff := &tools.Staff{
		Name: "Phoenix Staff",
		Element: "Fire",
		Power: 100,
	}

	poisonDagger := &tools.Dagger{
		Name: "Shadowfang",
		Poison: true,
		Stealth: 8,
	}

	gandalf := &classes.Mage{
		Name: "Gandalf",
		Level: 50,
		Manapool: 1000,
	}

	rogue := &classes.Rogue{
		Name: "Garret",
		Level: 30,
		Energy: 100,
		Equipment: poisonDagger,
	}

	fireball := &skills.FireballSpell{
		Spell: skills.Spell{
			Name: "Fireball",
			Cost: 50,
			Weapon: fireStaff,
			Element: "Fire",
		},
		BurnDamage: 25,
	}

	fmt.Println(gandalf.UseSkill(fireball))
	fmt.Println(rogue.Equipment.Attack())
}