package main

import (
    "fmt"
	"os"
	"bufio"
	"time"
	"strings"
)

func menu_inicial() {

	for {
		clear_screen()
		linha_1()
		fmt.Println("Escolha uma das 3 opções:")
		fmt.Println("	1 - começar o jogo.")
		fmt.Println("	2 - ler o manual do jogo.")
		fmt.Println("	q - sair do jogo.")
		linha_1()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Digite sua escolha: ")
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		switch escolha {
		case "1":
			fmt.Println("Você escolheu a começar o jogo..")
			time.Sleep(1 * time.Second)

			// o jogo inteiro ta aqui kkkkkkkkkkkkkkkkkk vou arrumar depois
			jogador := panel_create_player()
			panel_introducao_0(jogador)

		case "2":
			fmt.Println("Você escolheu ler o manual do jogo.")
			time.Sleep(1 * time.Second)
			panel_manual_do_game()
			continue

		case "q":
			quitando()

		default:
			fmt.Println("Escolha inválida. Por favor, escolha uma das opções válidas.")
			time.Sleep(1 * time.Second) 
			continue // Recarrega o menu
		}

		// Se chegou até aqui, a escolha foi válida, então sai do loop
		break
	}
}