package main

import "net"
import "fmt"
import "flag"
import "os"
import "strconv"

func main() {

	port := "36367"

	ln, err := net.Listen("tcp", ":" + port)

	if err != nil {
	fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s", port, err)
	os.Exit(1)
	}

	_ := ln.Close()
	fmt.Printf("TCP Port %q is available", port)
	os.Exit(0)
}