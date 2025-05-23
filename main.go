package main

import (
	"fmt"
	"go_study_rpg_service/internal/classes"
	"go_study_rpg_service/internal/person"
	"go_study_rpg_service/internal/skills"
	"go_study_rpg_service/internal/tools"
)

func main() {
	// Weapons
	fireStaff := &tools.Staff{
		Name:    "Phoenix Staff",
		Element: "Fire",
		Power:   100,
	}

	daggerPoisoned := &skills.Ability{
		Name: "Dagger Poisoned",
		CostEnergy: 40,
		Weapon: &tools.Dagger{
			Name: "Sting",
			Poison: false,
			Stealth: 10,
		},
	}

	poisonDagger := &tools.Dagger{
		Name:    "Shadowfang",
		Poison:  true,
		Stealth: 8,
	}

	
	// Chars
	gandalf := &classes.Mage{
		Name:     "Gandalf",
		Level:    50,
		Manapool: 1000,
	}

		// Skills
	fireball := &skills.FireballSpell{
		Spell: &skills.Spell{
			Name:    "Fireball",
			CostMana:    50,
			Weapon:  fireStaff,
			Element: "Fire",
		},
		BurnDamage: 25,
	}



	// new daggerPoisoned
	rogue := classes.NewRogue("Sam", 19, 60, &tools.Dagger{}, *daggerPoisoned)
	bk_mage := classes.NewMage("BK", 29, 200, &tools.Staff{}, fireball)
	rogue2 := classes.NewRogue("Sam", 19, 60, poisonDagger, *daggerPoisoned)

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
	
	// rogue := &classes.Rogue{
	// 	Name:      "Garret",
	// 	Level:     30,
	// 	Energy:    100,
	// 	Equipment: poisonDagger,
	// 	Ability: skills.DaggerPoisoned{
	// 		Name: "Dagger Poisened",
	// 		DamagePerSecond: 10,
	// 		Min_Lvl: 10,
	// 	},
	// }

	bk := person.NewPerson("bk", bk_mage)
	
	fmt.Println(bk)
	fmt.Println(bk.UseSkill())


	// Bag
	bagMage := &tools.Bag{}

	bagMage.Add("Poçao de mana")
	bagMage.Add("Staff of perdition")
	bagMage.Add("The One Ring")

	fmt.Println(bagMage.Items)

	fmt.Println(bagMage.Items)
	fmt.Println(gandalf.UseSkill())
	fmt.Println(rogue.Equipment.Attack())
	fmt.Println(rogue2.Equipment.Attack())
	fmt.Println(legolas.Equipment.Attack())
}
