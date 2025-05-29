package utils

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

// CheckPortOpen verifica se uma porta está aberta.
// port: a porta a ser verificada.
// Retorna true se a porta estiver aberta, caso contrário, false. Retorna um erro, se houver.
func CheckPortOpen(port string) (bool, error) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return false, err // Porta já está em uso ou bloqueada
	}
	_ = ln.Close()
	return true, nil // Porta disponível
}

// ListOpenPorts lista todas as portas abertas no sistema.
// Executa o comando `netstat` para obter as portas abertas.
// Retorna uma lista de strings com as portas abertas e um erro, se houver.
func ListOpenPorts() ([]string, error) {
	cmd := exec.Command("sh", "-c", "netstat -tuln | grep LISTEN | awk '{print $4}' | sed 's/.*://' | sort -n | uniq")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	ports := strings.Split(strings.TrimSpace(string(output)), "\n")
	return ports, nil
}

// ClosePort fecha uma porta específica.
// port: a porta a ser fechada.
// Retorna um erro, se houver.
func ClosePort(port string) error {
	cmd := exec.Command("fuser", "-k", port+"/tcp")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("falha ao fechar a porta %s: %v", port, err)
	}
	return nil
}

// OpenPort abre uma porta específica.
// port: a porta a ser aberta.
// Retorna um erro, se houver.
func OpenPort(port string) error {
	cmd := exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--dport", port, "-j", "ACCEPT")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("falha ao abrir a porta %s: %v", port, err)
	}
	return nil
}

// IsIPv6 verifica se um endereço IP é um endereço IPv6.
// ip: o endereço IP a ser verificado.
// Retorna true se o endereço IP for um endereço IPv6, caso contrário, false.
func IsIPv6(ip string) bool {
	var sanitizedIP string
	if strings.Contains(ip, ":") {
		sanitizedIP = strings.Replace(ip, ":", "", -1)

		// If the sanitized IP is still the same as the original, then it's an IPv6
		if sanitizedIP == ip {
			return true
		}
	}
	return false
}
