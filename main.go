package main

import (
	"fmt"
	"go_study_rpg_service/internal/classes"
	"go_study_rpg_service/internal/skills"
	"go_study_rpg_service/internal/tools"
)

func main() {
	fireStaff := &tools.Staff{
		Name:    "Phoenix Staff",
		Element: "Fire",
		Power:   100,
	}

	poisonDagger := &tools.Dagger{
		Name:    "Shadowfang",
		Poison:  true,
		Stealth: 8,
	}

	gandalf := &classes.Mage{
		Name:     "Gandalf",
		Level:    50,
		Manapool: 1000,
	}

	rogue := &classes.Rogue{
		Name:      "Garret",
		Level:     30,
		Energy:    100,
		Equipment: poisonDagger,
		Ability: skills.DaggerPoisoned{
			Name: "Dagger Poisened",
			DamagePerSecond: 10,
			Min_Lvl: 10,
		},
	}

	daggerPoisoned := &skills.DaggerPoisoned{
		Ability: &skills.Ability{
			Name: "Arac's ruin",
			CostEnergy: 15,
			Weapon: &tools.Dagger{
				Name: "Sith's Dagger",
				Poison: false,
				Damage: 20,
			},
		},
		DamagePerSecond: 8,
		Min_Lvl: 10,
	}

	fireball := &skills.FireballSpell{
		Spell: &skills.Spell{
			Name:    "Fireball",
			CostMana:    50,
			Weapon:  fireStaff,
			Element: "Fire",
		},
		BurnDamage: 25,
	}

	legolas := &classes.Archer{
		Name:   "Legolas",
		Level:  40,
		Energy: 170,
		Equipment: &tools.Bow{
			Name:       "Bow of Rivendell",
			Material:   "Mithrill",
			Damage:     60,
			CountArrow: 0,
		},
	}

	fmt.Println(rogue.Ability)
	fmt.Println(gandalf.UseSkill(fireball))
	fmt.Println(rogue.Equipment.Attack())
	fmt.Println(legolas.Equipment.Attack())
}
