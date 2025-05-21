package tools

import "fmt"

type Dagger struct {
	Name    string
	Poison  bool
	Damage int
}

func (d *Dagger) Attack() string {
	stealthAttack := "normal attack"
	if d.Damage > 5 {
		stealthAttack = "critical strike"

	}
	return fmt.Sprintf("Strinking from shadows with %s - %s!", d.Name, stealthAttack)
}

func (d *Dagger) GetName() string {
	return d.Name
}
