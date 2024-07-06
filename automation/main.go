package main

import (
    "fmt"
    "net"
)

func verificarPortaAberta(porta string) bool {
    // Tenta conectar-se à porta
    _, err := net.Dial("tcp", "localhost:"+porta)

    // Verifica se a conexão foi bem sucedida
    if err == nil {
        return true
    } else {
        return false
    }
}

func main() {
    // Porta que você deseja testar
    porta := "7070"

    // Verifica se a porta está aberta
    if verificarPortaAberta(porta) {
        fmt.Printf("A porta %s está aberta\n", porta)
    } else {
        fmt.Printf("A porta %s está fechada ou inacessível\n", porta)
    }
}