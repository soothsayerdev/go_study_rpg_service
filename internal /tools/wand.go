package tools

import "fmt"

type Wand struct {
	Name string
	Material string
	Power int
}

type Staff struct {
	Name string
	Element string
	Power int
}

type Sword struct {
	Name string
	Material string
	Damage int
}

type Dagger struct {
	Name string
	Poison bool
	Stealth int
}

type Weapon interface {
	Attack() string
	GetName() string
}

func (w *Wand) Attack() string {
	return fmt.Sprintf("Casting magic blast with %s made of %s!", w.Name, w.Material)
}

func (w *Wand) GetName() string {
	return w.Name
}


func (s *Staff) Attack() string {
	return fmt.Sprintf("Channeling %s elemental power through %s", s.Element, s.Name)
}
 
func (s *Staff) GetName() string {
	return s.Name
}


func (s *Sword) Attack() string {
	return fmt.Sprintf("Slashing with %s dealing %d damage!", s.Name, s.Damage)
}

func (s *Sword) GetName() string {
	return s.Name
}


func (d *Dagger) Attack() string {
	stealthAttack := "normal attack"
	if d.Stealth > 5 {
		stealthAttack = "critical strike"
	
	}
	return fmt.Sprintf("Strinking from shadows with %s - %s!", d.Name, stealthAttack)
}

func (d *Dagger) GetName() string {
	return d.Name
}