package skills

import (
	"fmt"
)

type DaggerPoisoned struct {
	Ability         *Ability
	DamagePerSecond int
	Min_Lvl         int
}
// new DaggerP

func (d *DaggerPoisoned) Execute() string {
	// if classes.Rogue.Level < d.Min_Lvl {
	// 	return "Not enough minimum level to cast this ability!"  -> regra de negocio - service
	// }
	return fmt.Sprintf("Throw %s, giving %d damage per second - Cost : %d - Minimum Level to cast: %d",
		d.Ability.Name, d.DamagePerSecond, d.Ability.CostEnergy, d.Min_Lvl)
}
