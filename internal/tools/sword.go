package tools

type Sword struct {
	Name     string
	Material string
	Damage   int
}

func (s *Sword) Attack() string   { return "Sword" }
func (s *Sword) DamageBonus() int { return 10 }
func (s *Sword) GetName() string  { return "Elendil's grace" }

// func (s *Sword) Attack() string {
// 	return fmt.Sprintf("Slashing with %s dealing %d damage!", s.Name, s.Damage)
// }

// func (s *Sword) GetName() string {
// 	return s.Name
// }
