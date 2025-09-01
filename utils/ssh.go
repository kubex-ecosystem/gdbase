// Package utils fornece utilitários diversos.
package utils

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

// ForwardMode: L = local forward (listen local -> dial remoto via SSH)
//
//	R = remote forward (listen remoto -> dial local)

type ForwardMode byte

const (
	LocalForward  ForwardMode = 'L'
	RemoteForward ForwardMode = 'R'
)

// ForwardSpec descreve um túnel.
// Ex.: "L:127.0.0.1:15432->127.0.0.1:5432"  (listen local 15432; alvo remoto 5432)
//
//	"R:0.0.0.0:5432->127.0.0.1:5432"     (listen remoto 5432; alvo local 5432)
type ForwardSpec struct {
	Mode   ForwardMode
	Listen string // host:port onde escuta (local p/ L; remoto p/ R)
	Target string // host:port destino (remoto p/ L; local p/ R)
}

func ParseForwardSpec(s string) (ForwardSpec, error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return ForwardSpec{}, fmt.Errorf("spec inválida: %q", s)
	}
	mode := ForwardMode(parts[0][0])
	rest := parts[1]
	lr := strings.Split(rest, "->")
	if len(lr) != 2 {
		return ForwardSpec{}, fmt.Errorf("esperado 'listen->target' em %q", s)
	}
	return ForwardSpec{Mode: mode, Listen: lr[0], Target: lr[1]}, nil
}

type SSHCred struct {
	User       string // ex.: "ubuntu"
	Password   string // opcional se usar chave
	PrivateKey []byte // opcional; PEM
	// HostKey: se vazio, tenta known_hosts; se falhar, recusa (sem Insecure)
}

type Tunnel struct {
	client *ssh.Client
}

// SSHConnect abre a conexão SSH segura validando host key via known_hosts.
func SSHConnect(addr string, cred SSHCred, timeout time.Duration) (*Tunnel, error) {
	var auths []ssh.AuthMethod
	if len(cred.PrivateKey) > 0 {
		signer, err := ssh.ParsePrivateKey(cred.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("chave privada inválida: %w", err)
		}
		auths = append(auths, ssh.PublicKeys(signer))
	}
	if cred.Password != "" {
		auths = append(auths, ssh.Password(cred.Password))
	}

	khPath := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
	hostKeyCb, err := knownhosts.New(khPath)
	if err != nil {
		return nil, fmt.Errorf("falha ao carregar known_hosts: %w", err)
	}

	cfg := &ssh.ClientConfig{
		User:            cred.User,
		Auth:            auths,
		HostKeyCallback: hostKeyCb,
		Timeout:         timeout,
	}
	cli, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		return nil, fmt.Errorf("SSH dial falhou: %w", err)
	}
	return &Tunnel{client: cli}, nil
}

func (t *Tunnel) Close() error { return t.client.Close() }

// Start inicia N túneis, cada qual numa goroutine.
// Aceita tanto L quanto R. Retorna função de teardown.
func (t *Tunnel) Start(specs ...ForwardSpec) (func(), error) {
	stopFns := make([]func(), 0, len(specs))
	for _, sp := range specs {
		switch sp.Mode {
		case LocalForward:
			ln, err := net.Listen("tcp", sp.Listen)
			if err != nil {
				return nil, fmt.Errorf("listen local %s: %w", sp.Listen, err)
			}
			stop := make(chan struct{})
			go func(listener net.Listener, target string) {
				for {
					conn, err := listener.Accept()
					if err != nil {
						select {
						case <-stop:
							return
						default:
						}
						continue
					}
					go t.handleLocal(conn, target)
				}
			}(ln, sp.Target)
			stopFns = append(stopFns, func() { close(stop); _ = ln.Close() })

		case RemoteForward:
			// Listen REMOTO via SSH (RFC 4254)
			rln, err := t.client.Listen("tcp", sp.Listen)
			if err != nil {
				return nil, fmt.Errorf("listen remoto %s: %w", sp.Listen, err)
			}
			stop := make(chan struct{})
			go func(rlistener net.Listener, target string) {
				for {
					rconn, err := rlistener.Accept()
					if err != nil {
						select {
						case <-stop:
							return
						default:
						}
						continue
					}
					// para cada conexão remota, disca LOCAL no target
					go t.handleRemote(rconn, target)
				}
			}(rln, sp.Target)
			stopFns = append(stopFns, func() { close(stop); _ = rln.Close() })

		default:
			return nil, fmt.Errorf("modo desconhecido em %v", sp)
		}
	}
	return func() {
		for i := len(stopFns) - 1; i >= 0; i-- {
			stopFns[i]()
		}
	}, nil
}

func (t *Tunnel) handleLocal(localConn net.Conn, remoteTarget string) {
	defer localConn.Close()
	remoteConn, err := t.client.Dial("tcp", remoteTarget)
	if err != nil {
		_ = localConn.Close()
		return
	}
	pipeBoth(localConn, remoteConn)
}

func (t *Tunnel) handleRemote(remoteConn net.Conn, localTarget string) {
	defer remoteConn.Close()
	localConn, err := net.Dial("tcp", localTarget)
	if err != nil {
		_ = remoteConn.Close()
		return
	}
	pipeBoth(remoteConn, localConn)
}

func pipeBoth(a, b net.Conn) {
	// a<->b com half-close em cada direção
	go copyClose(a, b)
	go copyClose(b, a)
}

func copyClose(dst, src net.Conn) {
	defer func() {
		// half-close no sentido de escrita, se suportado
		type closeWriter interface{ CloseWrite() error }
		if cw, ok := dst.(closeWriter); ok {
			_ = cw.CloseWrite()
		} else {
			_ = dst.Close()
		}
	}()
	_, _ = io.Copy(dst, src)
}
