package main

import (
	"fmt"
	"main/cmd/carro/classe"
	"main/cmd/carro/pessoa"
)

func main() {

	mage := classe.NewMage()
	knight := classe.NewKnight()
	priest := classe.NewPriest()

	jose := pessoa.NewPessoa("Jose", mage)
	joao := pessoa.NewPessoa("João", knight)

	//Por que não aceita?
	pedro := pessoa.NewPessoa("Pedro", priest)

	// Por que aceita?
	bispo := pessoa.NewPadre("Pedro", priest)

	fmt.Printf("%s attack is: %d\n", jose.Nome, jose.Attack())
	fmt.Printf("%s attack is: %d\n", joao.Nome, joao.Attack())
	fmt.Printf("%s attack is: %d\n", pedro.Nome, pedro.Attack())

	fmt.Printf("%s attack is: %d\n", bispo.Nome, bispo.Heal())
}
