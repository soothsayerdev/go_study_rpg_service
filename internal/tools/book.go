package tools

import (
	"fmt"
)

type Book struct {
	Name          string
	Type          string
	Min_Lvl       int
	LearningPages int
}

func (b *Book) toLearn() string {
	b.LearningPages += 1
	return fmt.Sprintf("To learn spells and invocations, %d", b.LearningPages)
}

func (b *Book) invokeCreature() string {
	if b.LearningPages > 5 {
		return ("Invoking a magma-dog")
	}
	if b.LearningPages > 10 {
		return ("Invoking a imp-wizard")
	}
	if b.LearningPages > 15 {
		return ("Invoking a troll")
	}
	return ("Not enough learning pages")
}

func (b *Book) GetName() string {
	return b.Name
}
