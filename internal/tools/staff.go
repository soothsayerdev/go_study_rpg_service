package tools

type Staff struct {
	Name    string
	Element string
	Power   int
}

func (s *Staff) Attack() string   { return "Channeling..." }
func (s *Staff) DamageBonus() int { return 6 }
func (s *Staff) GetName() string  { return "Staff of perdition" }

// func (s *Staff) Attack() string {
// 	return fmt.Sprintf("Channeling %s elemental power through %s", s.Element, s.Name)
// }

// func (s *Staff) GetName() string {
// 	return s.Name
// }
