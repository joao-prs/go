package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// ------------------------------>
//   funcoes ferramentas
// ------------------------------>

func time_3_sec() {
	
	segundos := 3
    for segundos >= 0 {
        fmt.Printf("\rsec: %2d", segundos)
        time.Sleep(1 * time.Second)
        segundos--
    }
    // fmt.Println() 
}
func time_12_sec() {
	
	segundos := 12
    for segundos >= 0 {
        fmt.Printf("\rsec: %2d", segundos)
        time.Sleep(1 * time.Second)
        segundos--
    }
    // fmt.Println() 
}


func clear_screen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// ------------------------------>
//   controles
//   Enter, y, n, q
// ------------------------------>

func confirma_enter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\r\033[0;33mprosseguir? [enter]: \033[0m")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	if confirm == "" {
		confirm = "y"
	} else if confirm == "q" {
		confirm = "q"
		quitando()
	}
}

func confirma_y_ou_n(CONFIRM bool) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[0;33m[Y/n]: \033[0m")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(confirm)

	// Confirma entrada vazia
	if confirm == "" {
		confirm = "y"
	}
	// Filtra em Y ou N
	switch strings.ToLower(confirm) {
	case "n":
		return false
	case "y":
		return true
	case "q":
		quitando()
	}

	return false
}

// ------------------------------>
//   controladores
// ------------------------------>

/*
func quitando() {
	// Implementa a função quitando conforme necessário
	fmt.Println("Quitando...")
	os.Exit(0)
}*/

func quitando() {
	frase := "Saindo do jogo....."
	for _, char := range frase {
		fmt.Printf("%c", char)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
	
	// Limpa a tela antes de encerrar
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	
	// Encerra o game
	os.Exit(0)
}


func linha_0() {
	fmt.Printf("************************************************************\n");
}
func linha_1() {
	fmt.Printf("------------------------------------------------------------\n");
}