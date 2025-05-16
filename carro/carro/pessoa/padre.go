package pessoa

type Padre struct {
	Nome   string
	titulo Titulo
}

type Titulo interface {
	Heal() int
}

func NewPadre(nome string, titulo Titulo) *Padre {
	return &Padre{
		Nome:   nome,
		titulo: titulo,
	}
}

func (p *Padre) Heal() int {
	return p.titulo.Heal()
}
