package main

import "fmt"

func main() {
	
	individuo := Funcionario{
		nome: "Jordan",
		idade: 26,
		cargo: "devops",
		salario: 4000.0,
	}

	fmt.Println("Nome do cara: ", individuo.nome)

	individuo.ExibirInfo()

	individuo.trabalhar()
}

