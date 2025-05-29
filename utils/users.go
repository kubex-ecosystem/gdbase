package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetPrimaryUser retorna o nome do usuário principal do sistema.
// Executa o comando `id -un` para obter o nome do usuário.
// Retorna o nome do usuário como string e um erro, se houver.
func GetPrimaryUser() (string, error) {
	cmd := exec.Command("id", "-un")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("erro ao obter o usuário principal: %v", err)
	}
	user := strings.TrimSpace(string(output))
	return user, nil
}

// GetPrimaryGroup retorna o nome do grupo principal do sistema.
// Executa o comando `id -gn` para obter o nome do grupo.
// Retorna o nome do grupo como string e um erro, se houver.
func GetPrimaryGroup() (string, error) {
	cmd := exec.Command("id", "-gn")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("erro ao obter o grupo principal: %v", err)
	}
	group := strings.TrimSpace(string(output))
	return group, nil
}

// GetGroups retorna uma lista de IDs dos grupos aos quais o usuário pertence.
// Executa o comando `id -G` para obter os IDs dos grupos.
// Retorna uma lista de strings com os IDs dos grupos e um erro, se houver.
func GetGroups() ([]string, error) {
	cmd := exec.Command("id", "-G")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter os grupos: %v", err)
	}
	groups := strings.Split(strings.TrimSpace(string(output)), " ")
	return groups, nil
}
