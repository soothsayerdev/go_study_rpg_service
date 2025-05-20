package tools

type Dagger struct {
	Name    string
	Poison  bool
	Stealth int
}

func (d *Dagger) Attack() string   { return "Strinking" }
func (d *Dagger) DamageBonus() int { return 15 }
func (d *Dagger) GetName() string  { return "Sting" }

// func (d *Dagger) Attack() string {
// 	stealthAttack := "normal attack"
// 	if d.Stealth > 5 {
// 		stealthAttack = "critical strike"

// 	}
// 	return fmt.Sprintf("Strinking from shadows with %s - %s!", d.Name, stealthAttack)
// }

// func (d *Dagger) GetName() string {
// 	return d.Name
// }
