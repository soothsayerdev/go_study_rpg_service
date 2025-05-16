package skills

import "fmt"

type FireballSpell struct {
	Spell      *Spell
	BurnDamage int
}

func (f *FireballSpell) Execute() string {
	return fmt.Sprintf("Casting %s with %s - Burn Dmg : %d - Cost: %d mana",
						f.Spell.Name, f.Spell.Weapon.Attack(), f.BurnDamage, f.Spell.CostMana)
}

func (f *FireballSpell) GetManaCost() int {
	return f.Spell.CostMana
}