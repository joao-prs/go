package main

import (
	"fmt"
	"os/exec"
)

// verifica se o pacote ta instalado
func isPackageInstalled(pkgName string) bool {
	cmd := exec.Command("dpkg", "-s", pkgName)
	err := cmd.Run()
	return err == nil
}

// instalar pacote
func installPackage(pkgName string) error {
	cmd := exec.Command("sudo", "apt-get", "install", "-y", pkgName)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}


func main() {
	packages := []string{"apache2", "php8.0"}

	for _, pkg := range packages {
		if isPackageInstalled(pkg) {
			fmt.Printf("O pacote %s já está instalado.\n", pkg)
		} else {
			fmt.Printf("O pacote %s não está instalado. Instalando...\n", pkg)
			err := installPackage(pkg)
			if err != nil {
				fmt.Printf("Falha ao instalar o pacote %s: %v\n", pkg, err)
			} else {
				fmt.Printf("Pacote %s instalado com sucesso.\n", pkg)
			}
		}
	}
}