package tools

import "fmt"

type Bow struct {
	Name       string
	Material   string
	Damage     int
	CountArrow int
}

func (b *Bow) Attack() string {
	if b.CountArrow == 0 {
		return "Not enough arrow suficient, reload bow!"
	}
	return "Shooting..."
}

func (b *Bow) DamageBonus() int { return 13 }
func (b *Bow) GetName() string  { return b.Name }

func (b *Bow) ReloadingBow() string { // qm carrega o arco é o arqueiro ou o arco? e pq?
	b.CountArrow += 1
	return fmt.Sprintf("Reloading bow, %d arrows now!", b.CountArrow)
}
