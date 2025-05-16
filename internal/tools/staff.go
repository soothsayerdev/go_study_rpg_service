package tools

import "fmt"

type Staff struct {
	Name    string
	Element string
	Power   int
}

func (s *Staff) Attack() string {
	return fmt.Sprintf("Channeling %s elemental power through %s", s.Element, s.Name)
}

func (s *Staff) GetName() string {
	return s.Name
}
