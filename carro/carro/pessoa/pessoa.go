package pessoa

type Pessoa struct {
	Nome   string
	classe Classe
}

type Classe interface {
	Attack() int
}

func NewPessoa(nome string, classe Classe) *Pessoa {
	return &Pessoa{
		Nome:   nome,
		classe: classe,
	}
}

func (p *Pessoa) Attack() int {
	return p.classe.Attack()
}
