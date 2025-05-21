package person

type Person struct {
	Name string
	classe classe
}

type classe interface {
	GetClassName() string
	UseSkill() string
}

func NewPerson(name string, classe classe) *Person {
	return &Person{
		Name: name,
		classe: classe,
	}
}

func (p *Person) UseSkill() string {
	return p.classe.UseSkill()
}