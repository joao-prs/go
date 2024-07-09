// go mod init game2
// go get golang.org/x/term

package main

import (
	//"fmt"
	//"os"
	//"bufio"
	//"strconv"
)

func main() {
	// criando personagem

	jogador := panel_create_player()
	panel_introducao_0(jogador)
	panel_manual_do_game()

	// encerrar quando nao houver mais funcoes
	quitando()
}