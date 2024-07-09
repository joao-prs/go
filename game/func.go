// go mod init game
// go get golang.org/x/term

package main

import (
	"fmt"
    "os"
    "os/exec"
    "runtime"
	"golang.org/x/term"
	"time"
)



// system functions #############################################
func clearScreen() {
    var clearCmd *exec.Cmd
    if runtime.GOOS == "windows" {
        clearCmd = exec.Command("cmd", "/c", "cls")
    } else {
        clearCmd = exec.Command("clear")
    }
    clearCmd.Stdout = os.Stdout
    clearCmd.Run()
}

func screen_size() {
    width, height, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Erro ao obter tamanho do terminal: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("Tamanho do terminal: %d x %d (largura x altura)\n", width, height)
}

func screen_size_block() {
    for {
		width, height, err := term.GetSize(int(os.Stdout.Fd()))
		
		// erro na função
        if err != nil {
            fmt.Fprintf(os.Stderr, "Erro ao obter tamanho do terminal: %v\n", err)
            os.Exit(1)
        }

		// quebra o loop, e sai da função
        if width >= 100 && height >= 30 {
            clearScreen()
            fmt.Println("Tamanho do terminal adequado!")
            break
        }

		// realimenta o loop até que a condição a cima seja atendida
        clearScreen()
        fmt.Printf("Resolução inferior a 100 x 30. Atual: %d x %d\n", width, height)
        time.Sleep(1 * time.Second)
    }
}

// custom functions #############################################
func show_main_lore() {
	fmt.Printf("aaaaaaaaaaaaaaaaa\n")
	fmt.Printf("aaaaaaaaaaaaaaaaa\n")
	fmt.Printf("aaaaaaaaaaaaaaaaa\n")
	fmt.Printf("aaaaaaaaaaaaaaaaa\n")
	fmt.Printf("aaaaaaaaaaaaaaaaa\n")
}