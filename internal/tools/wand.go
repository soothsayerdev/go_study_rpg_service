package tools

import "fmt"

type Wand struct {
	Name     string
	Material string
	Power    int
}

func (w *Wand) Attack() string {
	return fmt.Sprintf("Casting magic blast with %s made of %s!", w.Name, w.Material)
}

func (w *Wand) GetName() string {
	return w.Name
}
