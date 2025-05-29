package utils

import (
	"fmt"
	"os/exec"
)

// ClearScreen limpa a tela do terminal.
func ClearScreen() {
	fmt.Println("\033[H\033[2J")
}

// Figlet exibe um título estilizado usando o comando `figlet`.
// title: o título a ser exibido.
// Retorna um erro, se houver.
func Figlet(title string) error {
	if !CommandExists("figlet") {
		cmdFigletErr := exec.Command("kbx", "pkg", "install", "figlet").Run()
		if cmdFigletErr != nil {
			return cmdFigletErr
		}
		return nil
	}
	return exec.Command("figlet", "-W", "-c", "-t", "-X", title).Run()
}

// CommandExists verifica se um comando existe no sistema.
// cmd: o comando a ser verificado.
// Retorna true se o comando existir, caso contrário, false.
func CommandExists(cmd string) bool {
	if _, err := exec.LookPath(cmd); err != nil {
		if dpkgErr := exec.Command("dpkg", "-l", cmd).Run(); dpkgErr != nil {
			return false
		}
	}
	return true
}

// PrintTitle exibe um título formatado no terminal.
// title: o título a ser exibido.
func PrintTitle(title string) {
	fmt.Println("========================================")
	fmt.Println(title)
	fmt.Println("========================================")
}

// PrintSection exibe uma seção formatada no terminal.
// section: a seção a ser exibida.
func PrintSection(section string) {
	fmt.Println("----------------------------------------")
	fmt.Println(section)
	fmt.Println("----------------------------------------")
}

// WaitEnter aguarda o usuário pressionar ENTER para continuar.
func WaitEnter() {
	fmt.Print("Pressione ENTER para continuar...")
	_, scanlnErr := fmt.Scanln()
	if scanlnErr != nil {
		return // Ignora erro
	}
}

// WaitEnterClear aguarda o usuário pressionar ENTER e limpa a tela do terminal.
func WaitEnterClear() {
	WaitEnter()
	ClearScreen()
}
