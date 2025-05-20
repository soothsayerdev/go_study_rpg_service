package tools

type Wand struct {
	Name     string
	Material string
	Power    int
}

func (w *Wand) Attack() string   { return "Wand" }
func (w *Wand) DamageBonus() int { return 8 }
func (w *Wand) GetName() string  { return "Wand of wands (elder wand)" }

// func (w *Wand) Attack() string {
// 	return fmt.Sprintf("Casting magic blast with %s made of %s!", w.Name, w.Material)
// }

// func (w *Wand) GetName() string {
// 	return w.Name
// }
