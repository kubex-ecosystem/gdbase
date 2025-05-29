package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// acquireHostKey adquire a chave do host.
// host: o endereço do host.
// Retorna a chave pública do host e um erro, se houver.
func acquireHostKey(host string) (ssh.PublicKey, error) {
	var pubKey ssh.PublicKey

	pubKey = getKnownHostKey(host, "")
	if pubKey != nil {
		return pubKey, nil
	}

	pubKey = getKnownHostKey(host, filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa.pub"))
	if pubKey != nil {
		return pubKey, nil
	}

	return nil, fmt.Errorf("chave do host não encontrada")
}

// saveHostKey salva a chave do host no arquivo known_hosts.
// host: o endereço do host.
// pubKey: a chave pública do host.
// Retorna um erro, se houver.
func saveHostKey(host string, pubKey ssh.PublicKey) error {
	knownHostsPath := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
	file, err := os.OpenFile(knownHostsPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("falha ao abrir o arquivo known_hosts: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s %s\n", host, strings.TrimSpace(string(ssh.MarshalAuthorizedKey(pubKey)))))
	if err != nil {
		return fmt.Errorf("falha ao escrever no arquivo known_hosts: %v", err)
	}

	return nil
}

// SshTunnel configura um túnel SSH com base nos parâmetros fornecidos.
// sshUser: o nome de usuário SSH.
// sshCert: o certificado SSH.
// sshPassword: a senha SSH.
// sshAddress: o endereço do servidor SSH.
// sshPort: a porta do servidor SSH.
// tunnels: os túneis a serem configurados.
// Retorna um erro, se houver.
func SshTunnel(sshUser string, sshCert string, sshPassword string, sshAddress string, sshPort string, tunnels ...string) error {
	if sshUser == "" || sshAddress == "" || sshPort == "" || len(tunnels) == 0 {
		return fmt.Errorf("argumentos inválidos")
	}

	var sshHost string
	var sshConfig *ssh.ClientConfig

	sshHost = fmt.Sprintf("%s:%s", sshAddress, sshPort)

	if sshCert != "" {
		sshSigner, sshSignerErr := ssh.ParsePrivateKey([]byte(sshCert))
		if sshSignerErr != nil {
			return fmt.Errorf("falha ao analisar a chave privada: %v", sshSignerErr)
		}
		sshConfig = &ssh.ClientConfig{
			User: sshUser,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(sshSigner),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Adiciona um HostKeyCallback
		}
	} else {
		sshConfig = &ssh.ClientConfig{
			User: sshUser,
			Auth: []ssh.AuthMethod{
				ssh.Password(sshPassword),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Adiciona um HostKeyCallback
		}
	}

	sshConn, sshConnErr := ssh.Dial("tcp", sshHost, sshConfig)
	if sshConnErr != nil {
		return fmt.Errorf("falha ao conectar ao servidor SSH: %v", sshConnErr)
	}
	defer sshConn.Close()

	var localListeners []*net.Listener
	var localListener net.Listener
	var localListenerErr error
	for _, tunnel := range tunnels {
		localListener, localListenerErr = net.Listen("tcp", tunnel)
		if localListenerErr != nil {
			return fmt.Errorf("falha ao iniciar o listener local: %v", localListenerErr)
		}
		localListeners = append(localListeners, &localListener)
		defer localListener.Close()
		log.Printf("Túnel local iniciado em %s", tunnel)
	}

	for {
		localConn, localConnErr := localListener.Accept()
		if localConnErr != nil {
			log.Printf("Erro ao aceitar a conexão local: %v", localConnErr)
			continue
		}

		remoteConn, remoteConnErr := sshConn.Dial("tcp", localConn.RemoteAddr().String())
		if remoteConnErr != nil {
			log.Printf("Erro ao conectar ao servidor remoto: %v", remoteConnErr)
			continue
		}

		go forwardData(localConn, remoteConn)
		go forwardData(remoteConn, localConn)
	}
}

// forwardData encaminha dados entre as conexões.
// src: a conexão de origem.
// dest: a conexão de destino.
func forwardData(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()

	_, err := io.Copy(src, dest)
	if err != nil {
		log.Printf("Erro ao transferir dados: %v", err)
	}
}

// getKnownHostKey obtém a chave de host conhecida para validação.
// host: o endereço do host.
// publicKeyPath: o caminho para a chave pública.
// Retorna a chave pública do host.
func getKnownHostKey(host string, publicKeyPath string) ssh.PublicKey {
	if publicKeyPath != "" {
		key, err := os.ReadFile(publicKeyPath)
		if err != nil {
			log.Fatalf("falha ao ler a chave pública: %v", err)
			return nil
		}
		pubKey, _, _, _, parseAuthorizedKeyErr := ssh.ParseAuthorizedKey(key)
		if parseAuthorizedKeyErr != nil {
			log.Fatalf("falha ao analisar a chave pública: %v", parseAuthorizedKeyErr)
			return nil
		}
		return pubKey
	}

	// Tenta usar o ssh-copy-id para obter a chave pública, se não der certo, retorna nil
	copyIDCmd, copyIDCmdErr := exec.Command("bash", "-c", fmt.Sprintf("ssh-copy-id %s", host)).Output()
	if copyIDCmdErr != nil {
		log.Printf("falha ao executar o comando ssh-copy-id: %v", copyIDCmdErr)
		return nil
	}
	pubKey, _, _, _, parseAuthorizedKeyErr := ssh.ParseAuthorizedKey(copyIDCmd)
	if parseAuthorizedKeyErr != nil {
		log.Fatalf("falha ao analisar a chave pública: %v", parseAuthorizedKeyErr)
		return nil
	}
	return pubKey // Substitua pela chave pública apropriada para o seu caso.
}
