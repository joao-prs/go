package main

import (
	"fmt"
	//"float"
)

// Pessoa
type Pessoa struct {
    nome string
    idade int
}

// Método de pessoa
func (p *Pessoa) ExibirInfo() {
    fmt.Println("Nome:", a.nome)
    fmt.Println("Idade:", a.idade)
}

// funcionario
type Funcionario struct {
	Pessoa
	cargo string
	salario float32
}

// Método de funcionario
func (f *Funcionario) trabalhar() {
	fmt.Printf("\ntrabalhando...\n")
}