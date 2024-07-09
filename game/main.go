// go mod init
// go get golang.org/x/term

package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type Pessoa struct {
	Nome string
	Idade int
}

func main() {
	
	var p Pessoa
	screen_size_block()
	fmt.Printf("Qual seu nome?\n")

	// lendo o nome da pessoa
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		p.Nome = scanner.Text()
	}
	
	// lendo a idade da pessoa
	fmt.Printf("digite um idade\n")
	if scanner.Scan() {
		idadeStr := scanner.Text()
		idade, err := strconv.Atoi(idadeStr)
		if err != nil {
			fmt.Println("Erro ao converter idade:", err)
			return
		}
		p.Idade = idade
	}
	// exibindo os dados das pessoas
	fmt.Printf("Nome: %s, Idade: %d\n", p.Nome, p.Idade)

	clearScreen() 
	show_main_lore()
	screen_size()
}