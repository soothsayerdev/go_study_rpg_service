package tools

import "fmt"

type Bow struct {
	Name       string
	Material   string
	Damage     int
	CountArrow int
}

func (b *Bow) ReloadingBow() string {
	b.CountArrow += 1
	return fmt.Sprintf("Reloading bow, %d arrows now!", b.CountArrow)
}

func (b *Bow) Attack() string {
	if b.CountArrow == 0 {
		return "Not enough arrow suficient, reload bow!"
	}
	return fmt.Sprintf("Shooting arrow with %s", b.Name)
}

func (b *Bow) GetName() string {
	return b.Name
}
