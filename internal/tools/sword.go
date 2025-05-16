package tools

import "fmt"

type Sword struct {
	Name     string
	Material string
	Damage   int
}

func (s *Sword) Attack() string {
	return fmt.Sprintf("Slashing with %s dealing %d damage!", s.Name, s.Damage)
}

func (s *Sword) GetName() string {
	return s.Name
}
