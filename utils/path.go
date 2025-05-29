package utils

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// getAbsRelativesPaths obtém os caminhos absolutos dos caminhos relativos fornecidos.
// basePath: caminho base a partir do qual os caminhos relativos serão resolvidos.
// paths: lista de caminhos relativos.
// Retorna uma lista de caminhos absolutos e um erro, se houver.
func getAbsRelativesPaths(basePath string, paths []string) ([]string, error) {
	var absPaths []string
	for _, path := range paths {
		absPath, err := filepath.Abs(filepath.Join(basePath, path))
		if err != nil {
			return nil, fmt.Errorf("erro ao obter o caminho absoluto do arquivo: %v", err)
		}
		absPaths = append(absPaths, absPath)
	}
	return absPaths, nil
}

// SanitizePath sanitiza um caminho de arquivo, garantindo que ele esteja dentro do caminho base.
// basePath: caminho base.
// inputPath: caminho de entrada a ser sanitizado.
// Retorna o caminho sanitizado e um erro, se houver.
func SanitizePath(basePath, inputPath string) (string, error) {
	// Limpa o caminho de entrada
	cleanPath := filepath.Clean(inputPath)

	// Junta o caminho base e o caminho de entrada limpo
	fullPath := filepath.Join(basePath, cleanPath)

	// Garante que o caminho completo esteja dentro do caminho base
	if !strings.HasPrefix(fullPath, filepath.Clean(basePath)+string(filepath.Separator)) {
		return "", errors.New("invalid file path")
	}

	return fullPath, nil
}

// EnsureDir garante que um diretório exista, criando-o se necessário.
// path: caminho do diretório.
// perm: permissões do diretório.
// userGroup: lista contendo o usuário e grupo do diretório.
// Retorna um erro, se houver.
func EnsureDir(path string, perm os.FileMode, userGroup []string) error {
	basePath := filepath.Dir(path)
	targetPath := filepath.Base(path)
	pathSSanitized, pathSSanitizedErr := SanitizePath(basePath, targetPath)
	if pathSSanitizedErr != nil {
		return pathSSanitizedErr
	}

	// Verifica se o diretório já existe
	cmdCheckDir := exec.Command("test", "-d", pathSSanitized)
	cmdCheckDirErr := cmdCheckDir.Run()
	if cmdCheckDirErr == nil {
		return nil
	}

	// Cria o diretório
	cmdPathUser := exec.Command("mkdir", "-p", pathSSanitized)
	cmdPathUserErr := cmdPathUser.Run()
	if cmdPathUserErr != nil {
		cmdPathRoot := exec.Command("sudo", "mkdir", "-p", pathSSanitized)
		cmdPathRootErr := cmdPathRoot.Run()
		if cmdPathRootErr != nil {
			return fmt.Errorf("erro ao criar o diretório: %w", cmdPathRootErr)
		}

		// Define o proprietário e as permissões do diretório
		if len(userGroup) > 0 {
			cmdChown := exec.Command("sudo", "chown", strings.Join(userGroup, ":"), pathSSanitized)
			cmdChownErr := cmdChown.Run()
			if cmdChownErr != nil {
				return fmt.Errorf("erro ao definir o proprietário do diretório: %v", cmdChownErr)
			}
		}

		permOp := perm
		if permOp == 0 {
			permOp = os.ModePerm
		}

		// Define as permissões do diretório
		cmdChmod := exec.Command("sudo", "chmod", strconv.Itoa(int(permOp)), pathSSanitized)
		cmdChmodErr := cmdChmod.Run()
		if cmdChmodErr != nil {
			return fmt.Errorf("erro ao definir permissões do diretório: %w", cmdChmodErr)
		}
	}
	return nil
}

// EnsureFile garante que um arquivo exista, criando-o se necessário.
// path: caminho do arquivo.
// perm: permissões do arquivo.
// userGroup: lista contendo o usuário e grupo do arquivo.
// Retorna um erro, se houver.
func EnsureFile(path string, perm os.FileMode, userGroup []string) error {
	basePath := filepath.Dir(path)
	targetPath := filepath.Base(path)
	pathSSanitized, pathSSanitizedErr := SanitizePath(basePath, targetPath)
	if pathSSanitizedErr != nil {
		return pathSSanitizedErr
	}

	// Verifica se o diretório base existe
	cmdCheckDir := exec.Command("ls", "-d", basePath)
	cmdCheckDirErr := cmdCheckDir.Run()
	if cmdCheckDirErr != nil {
		ensureDirErr := EnsureDir(basePath, perm, userGroup)
		if ensureDirErr != nil {
			return ensureDirErr
		}
	}

	// Verifica se o arquivo já existe
	cmdCheckFile := exec.Command("ls", pathSSanitized)
	cmdCheckFileErr := cmdCheckFile.Run()
	if cmdCheckFileErr == nil {
		return nil
	}

	// Cria o arquivo
	cmdPathUser := exec.Command("touch", pathSSanitized)
	cmdPathUserErr := cmdPathUser.Run()
	if cmdPathUserErr != nil {
		cmdPathRoot := exec.Command("sudo", "touch", pathSSanitized)
		cmdPathRootErr := cmdPathRoot.Run()
		if cmdPathRootErr != nil {
			return fmt.Errorf("erro ao criar o arquivo: %v", cmdPathRootErr)
		}

		// Define o proprietário e as permissões do arquivo
		if len(userGroup) > 0 {
			cmdChown := exec.Command("sudo", "chown", strings.Join(userGroup, ":"), pathSSanitized)
			cmdChownErr := cmdChown.Run()
			if cmdChownErr != nil {
				return fmt.Errorf("erro ao definir o proprietário do arquivo: %v", cmdChownErr)
			}
		}
		permOp := perm
		if permOp == 0 {
			permOp = os.ModePerm
		}

		// Define as permissões do arquivo
		cmdChmod := exec.Command("sudo", "chmod", strconv.Itoa(int(permOp)), pathSSanitized)
		cmdChmodErr := cmdChmod.Run()
		if cmdChmodErr != nil {
			return fmt.Errorf("erro ao definir permissões do arquivo: %v", cmdChmodErr)
		}
	}
	return nil
}

// createTempDirTree cria a estrutura de diretórios temporários.
// Retorna um erro, se houver.
func createTempDirTree() error {
	tempDir := os.Getenv("XDG_CACHE_HOME")
	if tempDir == "" {
		tempDir = os.Getenv("HOME")
		if tempDir == "" {
			tempDir = "/tmp"
		}
		tempDir = filepath.Join(tempDir, ".cache")
	}

	primaryUser, primaryUserErr := GetPrimaryUser()
	if primaryUserErr != nil {
		return primaryUserErr
	}
	primaryGroup, primaryGroupErr := GetPrimaryGroup()
	if primaryGroupErr != nil {
		return primaryGroupErr
	}
	owners := []string{primaryUser, primaryGroup}

	// Garantir que o diretório .cache exista
	tempDirBaseErr := EnsureDir(tempDir, 0777, owners)
	if tempDirBaseErr != nil {
		return tempDirBaseErr
	}

	// Garantir que o diretório kubex exista dentro de .cache
	tempDir = filepath.Join(tempDir, "kubex")
	tempDirKubexErr := EnsureDir(tempDir, 0777, owners)
	if tempDirKubexErr != nil {
		return tempDirKubexErr
	}

	// Garantir que o diretório kbx exista dentro de kubex
	tempDir = filepath.Join(tempDir, "kbx")
	tempDirKbxErr := EnsureDir(tempDir, 0777, owners)
	if tempDirKbxErr != nil {
		return tempDirKbxErr
	}

	_ = os.Setenv("KBX_TEMP_DIR", tempDir)

	return nil
}

// EnsureTempDir garante que o diretório temporário exista.
// Retorna um erro, se houver.
func EnsureTempDir() error {
	tempDir := os.Getenv("KBX_TEMP_DIR")
	if tempDir == "" {
		createTempDirTreeErr := createTempDirTree()
		if createTempDirTreeErr != nil {
			return createTempDirTreeErr
		}
		tempDir = os.Getenv("KBX_TEMP_DIR")
		if tempDir == "" {
			return errors.New("erro ao garantir o diretório temporário")
		}
	}

	SocketPath := filepath.Join(tempDir, "kbx.sock")
	_ = os.Setenv("KBX_UNIX_SOCK_FILE", SocketPath)

	SocketPidPath := filepath.Join(tempDir, "unixz.pid")
	_ = os.Setenv("KBX_UNIX_PID_FILE", SocketPidPath)

	SrvLogFilePath := filepath.Join(tempDir, "unix.log")
	_ = os.Setenv("KBX_UNIX_LOG_FILE", SrvLogFilePath)

	TcpLogFilePath := filepath.Join(tempDir, "tcp.log")
	_ = os.Setenv("KBX_TCP_LOG_FILE", TcpLogFilePath)

	CliLogFilePath := filepath.Join(tempDir, "cli.log")
	_ = os.Setenv("KBX_CLI_LOG_FILE", CliLogFilePath)

	return nil
}

// GetTempDir obtém o diretório temporário.
// Retorna o caminho do diretório temporário e um erro, se houver.
func GetTempDir() (string, error) {
	tempDir := os.Getenv("KBX_TEMP_DIR")
	if tempDir == "" {
		tempDirErr := EnsureTempDir()
		if tempDirErr != nil {
			return "", tempDirErr
		}
		tempDir = os.Getenv("KBX_TEMP_DIR")
		if tempDir == "" {
			return "", errors.New("erro ao obter o diretório temporário")
		}
	}
	return tempDir, nil
}

// GetWorkDir obtém o diretório de trabalho.
// Retorna o caminho do diretório de trabalho e um erro, se houver.
func GetWorkDir() (string, error) {
	cwd := os.Getenv("KBX_CWD")
	if cwd == "" {
		homeDir, _ := os.UserHomeDir()
		if homeDir == "" {
			homedirCmd := exec.Command("echo", "$HOME")
			homedirOut, _ := homedirCmd.Output()
			homeDir = strings.TrimSpace(string(homedirOut))
			if homeDir == "" {
				homeDir, _ = os.Getwd()
			}
		}
		homeDir = filepath.Join(homeDir, ".kubex", "kbx")
		ensureHomeDirErr := EnsureDir(homeDir, 0777, []string{})
		if ensureHomeDirErr != nil {
			return "", ensureHomeDirErr
		}
		cwd = homeDir
	}
	return cwd, nil
}

func GetKubexDir() (string, error) {
	cwd := os.Getenv("KBX_CWD")
	if cwd == "" {
		homeDir, _ := os.UserHomeDir()
		if homeDir == "" {
			homedirCmd := exec.Command("echo", "$HOME")
			homedirOut, _ := homedirCmd.Output()
			homeDir = strings.TrimSpace(string(homedirOut))
			if homeDir == "" {
				homeDir, _ = os.Getwd()
			}
		}
		homeDir = filepath.Join(homeDir, ".kubex")
		ensureHomeDirErr := EnsureDir(homeDir, 0777, []string{})
		if ensureHomeDirErr != nil {
			return "", ensureHomeDirErr
		}
		cwd = homeDir
	}
	return cwd, nil
}

// GetHomeDir obtém o diretório home do usuário.
// Retorna o caminho do diretório home e um erro, se houver.
func GetHomeDir() (string, error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDirCmd := exec.Command("echo", "$HOME")
		homeDirOut, _ := homeDirCmd.Output()
		homeDir = strings.TrimSpace(string(homeDirOut))
		if homeDir == "" {
			homeDir, _ = os.UserHomeDir()
			if homeDir == "" {
				primaryUserName, primaryUserNameErr := GetPrimaryUser()
				if primaryUserNameErr != nil {
					return "", primaryUserNameErr
				}
				homeDir = filepath.Join("/home", primaryUserName)
				if _, err := os.Stat(homeDir); os.IsNotExist(err) {
					homeDir = filepath.Join("/Users", primaryUserName)
					if _, err := os.Stat(homeDir); os.IsNotExist(err) {
						homeDir = filepath.Join("/root")
						if _, err := os.Stat(homeDir); os.IsNotExist(err) {
							return "", errors.New("erro ao obter o diretório home")
						}
					}
				}
			}
		}
	}
	return homeDir, nil
}

// ListFiles lista os arquivos em um diretório que correspondem a um padrão.
// path: caminho do diretório.
// pattern: padrão a ser procurado nos nomes dos arquivos.
// Retorna uma lista de arquivos que correspondem ao padrão e um erro, se houver.
func ListFiles(path string, pattern string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.Contains(info.Name(), pattern) {
			files = append(files, filePath)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// CheckFilePathExists verifica se um caminho de arquivo existe.
// path: caminho do arquivo.
// Retorna true se o arquivo existir, caso contrário, false e um erro, se houver.
func CheckFilePathExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	}
	if err != nil {
		return false, err
	}
	if info.IsDir() {
		return false, errors.New("path is a directory, not a file")
	}
	return true, nil
}

// CheckPathLastAccessTime verifica o tempo de último acesso de um caminho.
// path: caminho do arquivo ou diretório.
// Retorna o tempo de último acesso e um erro, se houver.
func CheckPathLastAccessTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

// unmountTmpfsOrSecuredTempDir desmonta um diretório tmpfs ou seguro.
// tempDir: caminho do diretório temporário.
// Retorna um erro, se houver.
func unmountTmpfsOrSecuredTempDir(tempDir string) error {
	defer func() {
		// Ensure umount at the end, or remove all contents forcibly if umount fails
		// Check if tmpfs is mounted
		mountCmd := exec.Command("mount")
		mountCmdOut, mountCmdOutErr := mountCmd.Output()
		if mountCmdOutErr != nil {
			return
		}
		if strings.Contains(string(mountCmdOut), tempDir) {
			// Remove all contents forcibly
			rmCmd := exec.Command("rm", "-rf", tempDir)
			rmCmdErr := rmCmd.Run()
			if rmCmdErr != nil {
				return
			}
		}
	}()

	// Unmount tmpfs
	umountCmd := exec.Command("umount", tempDir)
	if umountCmdErr := umountCmd.Run(); umountCmdErr != nil {
		return fmt.Errorf("erro ao desmontar tmpfs: %v", umountCmdErr)
	}
	return nil
}

// checkTmpfsOrSecuredTempDirUsage verifica o uso de um diretório tmpfs ou seguro.
// tempDir: caminho do diretório temporário.
// Retorna um erro, se houver.
func checkTmpfsOrSecuredTempDirUsage(tempDir string) error {
	// Check tmpfs usage and unmount if unused for 10 minutes
	time.Sleep(10 * time.Minute)
	lastAccessTime, lastAccessTimeErr := CheckPathLastAccessTime(tempDir)
	if lastAccessTimeErr != nil {
		unmountTmpfsOrSecuredTempDirErr := unmountTmpfsOrSecuredTempDir(tempDir)
		if unmountTmpfsOrSecuredTempDirErr != nil {
			return unmountTmpfsOrSecuredTempDirErr
		}
	}
	if time.Since(lastAccessTime) > 10*time.Minute {
		unmountTmpfsOrSecuredTempDirErr := unmountTmpfsOrSecuredTempDir(tempDir)
		if unmountTmpfsOrSecuredTempDirErr != nil {
			return unmountTmpfsOrSecuredTempDirErr
		} else {
			return nil
		}
	}
	return checkTmpfsOrSecuredTempDirUsage(tempDir)
}

// mountTmpfsOrSecuredTempDir monta um diretório tmpfs ou seguro.
// tempDir: caminho do diretório temporário.
// path: caminho do arquivo temporário.
// Retorna o caminho do diretório temporário e um erro, se houver.
func mountTmpfsOrSecuredTempDir(tempDir string, path string) (string, error) {
	// Mount tmpfs
	mountCmd := exec.Command("mount", "-t", "tmpfs", "-o", "size=64M", "tmpfs", tempDir)
	if mountCmdErr := mountCmd.Run(); mountCmdErr != nil {
		logzExtCmdA := exec.Command("kbx", "logz", "warn", "erro ao montar tmpfs: "+mountCmdErr.Error(), "utils")
		logzExtCmdB := exec.Command("kbx", "logz", "warn", "tentando montar diretório seguro: "+tempDir, "utils")
		_ = logzExtCmdA.Run()
		_ = logzExtCmdB.Run()
	} else {
		// Monitor tmpfs, if it has been unused for 10 minutes, unmount it
		go func() {
			_ = checkTmpfsOrSecuredTempDirUsage(tempDir)
		}()
		return tempDir, nil
	}
	// Getting user and group for ensure ownership of tempDir all contents
	primaryUser, primaryUserErr := GetPrimaryUser()
	if primaryUserErr != nil {
		return "", primaryUserErr
	}
	primaryGroup, primaryGroupErr := GetPrimaryGroup()
	if primaryGroupErr != nil {
		return "", primaryGroupErr
	}
	owners := []string{primaryUser, primaryGroup}
	// Ensure dir
	ensureDirErr := EnsureDir(tempDir, 0700, owners)
	if ensureDirErr != nil {
		return "", ensureDirErr
	}
	// Ensure file
	ensureFileErr := EnsureFile(path, 0600, owners)
	if ensureFileErr != nil {
		return "", ensureFileErr
	}
	return tempDir, nil
}

// AcctTempSsCtx cria um arquivo temporário seguro para um token.
// token: o token a ser armazenado.
// Retorna o caminho do arquivo temporário e um erro, se houver.
func AcctTempSsCtx(token string) (string, error) {
	tempDir, tempDirErr := GetTempDir()
	if tempDirErr != nil {
		return "", tempDirErr
	}
	tempDir = filepath.Join(tempDir, ".kbx_tmpfs")
	ensureDirErr := EnsureDir(tempDir, 0700, []string{})
	if ensureDirErr != nil {
		return "", ensureDirErr
	}

	// Remove todos os conteúdos antigos do tempDir que contêm "acxtkb" em qualquer parte do nome
	oldFiles, oldFilesErr := ListFiles(tempDir, "acxtkb")
	if oldFilesErr != nil {
		return "", oldFilesErr
	}
	for _, oldFile := range oldFiles {
		commonRmErr := os.Remove(oldFile)
		if commonRmErr != nil {
			rootRmCmd := exec.Command("sudo", "rm", "-rf", oldFile)
			rootRmCmdErr := rootRmCmd.Run()
			if rootRmCmdErr != nil {
				return "", fmt.Errorf("erro ao remover arquivo antigo: %v", rootRmCmdErr)
			}
		}
	}

	// Cria um arquivo temporário seguro para o token
	mkSsCtxFilePth, mkSsCtxFilePthErr := mountTmpfsOrSecuredTempDir(tempDir, filepath.Join(tempDir, "acxtkb"))
	if mkSsCtxFilePthErr != nil {
		return "", mkSsCtxFilePthErr
	}

	// Criptografa o token e escreve no arquivo tmpfs, assina o arquivo com hash e mantém o hash na memória
	tokenFilePath, tokenFilePathErr := SetSecureSignedTempData(mkSsCtxFilePth, token, "")
	if tokenFilePathErr != nil {
		return "", tokenFilePathErr
	}

	return tokenFilePath, nil
}

// AcctTempSsCtxCheck verifica se o token é válido.
// Retorna o conteúdo do token e um erro, se houver.
func AcctTempSsCtxCheck() (string, error) {
	tempDir, tempDirErr := GetTempDir()
	if tempDirErr != nil {
		return "", tempDirErr
	}
	tempDir = filepath.Join(tempDir, ".kbx_tmpfs")
	ensureDirErr := EnsureDir(tempDir, 0700, []string{})
	if ensureDirErr != nil {
		return "", ensureDirErr
	}
	tokenCnt, tokenCntErr := CheckSecureSignedTempData(tempDir, "")
	if tokenCntErr != nil {
		return "", tokenCntErr
	}

	return tokenCnt, nil
}

// AcctTempSsCtxClear remove todos os conteúdos antigos do tempDir que contêm "acxtkb" em qualquer parte do nome.
// Retorna um erro, se houver.
func AcctTempSsCtxClear() error {
	tempDir, tempDirErr := GetTempDir()
	if tempDirErr != nil {
		return tempDirErr
	}
	tempDir = filepath.Join(tempDir, ".kbx_tmpfs")
	ensureDirErr := EnsureDir(tempDir, 0700, []string{})
	if ensureDirErr != nil {
		return ensureDirErr
	}
	// Remove todos os conteúdos antigos do tempDir que contêm "acxtkb" em qualquer parte do nome, recursivamente
	oldFiles, oldFilesErr := ListFiles(tempDir, "acxtkb")
	if oldFilesErr != nil {
		return oldFilesErr
	}
	for _, oldFile := range oldFiles {
		commonRmErr := os.Remove(oldFile)
		if commonRmErr != nil {
			rootRmCmd := exec.Command("sudo", "rm", "-rf", oldFile)
			rootRmCmdErr := rootRmCmd.Run()
			if rootRmCmdErr != nil {
				return fmt.Errorf("erro ao remover arquivo antigo: %v", rootRmCmdErr)
			}
		}
	}
	return nil
}

// SetSecureSignedTempData criptografa o token e escreve no arquivo tmpfs, assina o arquivo com hash e mantém o hash na memória.
// tempDir: diretório temporário.
// data: dados a serem armazenados.
// tag: tag opcional para o arquivo.
// Retorna o nome do arquivo temporário e um erro, se houver.
func SetSecureSignedTempData(tempDir string, data string, tag string) (string, error) {
	tokenFilePrefix := "acxtkb"
	if tag != "" {
		tokenFilePrefix = tag
	}
	tokenFileTimeSufix := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String())))
	tokenFileNameHash := tokenFilePrefix + tokenFileTimeSufix
	tokenFilePath := filepath.Join(tempDir, tokenFileNameHash)
	signTokenHash := createHash("acctkn" + data + tokenFileTimeSufix)
	encryptedToken, encryptedTokenErr := EncryptData(data, signTokenHash)
	if encryptedTokenErr != nil {
		return "", fmt.Errorf("erro ao criptografar token: %w", encryptedTokenErr)
	}
	if writeFileErr := os.WriteFile(tokenFilePath, []byte(encryptedToken), 0600); writeFileErr != nil {
		return "", fmt.Errorf("erro ao escrever token: %w", writeFileErr)
	}
	if writeSignHashErr := os.WriteFile(filepath.Join("/dev/shm", tokenFileNameHash), []byte(signTokenHash), 0600); writeSignHashErr != nil {
		return "", fmt.Errorf("erro ao escrever hash do token: %w", writeSignHashErr)
	}
	return tokenFileNameHash, nil
}

// CheckSecureSignedTempData verifica se o token armazenado é válido.
// tempDir: diretório temporário.
// tag: tag opcional para o arquivo.
// Retorna o conteúdo do token e um erro, se houver.
func CheckSecureSignedTempData(tempDir string, tag string) (string, error) {
	tokenFilePrefix := "acxtkb"
	if tag != "" {
		tokenFilePrefix = tag
	}
	tokenFiles, tokenFilesErr := ListFiles(tempDir, tokenFilePrefix)
	if tokenFilesErr != nil {
		return "", fmt.Errorf("erro ao listar arquivos de token: %w", tokenFilesErr)
	}
	if len(tokenFiles) == 0 {
		return "", errors.New("nenhum arquivo de token encontrado")
	}
	tokenFileName := tokenFiles[0]
	tokenFileHash, tokenFileHashErr := os.ReadFile(filepath.Join("/dev/shm", tokenFileName))
	if tokenFileHashErr != nil {
		return "", fmt.Errorf("erro ao ler hash do arquivo de token: %w", tokenFileHashErr)
	}
	tokenFileData, tokenFileDataErr := os.ReadFile(tokenFileName)
	if tokenFileDataErr != nil {
		return "", fmt.Errorf("erro ao ler arquivo de token: %w", tokenFileDataErr)
	}
	tokenFileDataStr := string(tokenFileData)
	tokenFileDataHash := createHash("acctkn" + tokenFileDataStr + tokenFileName)
	if string(tokenFileHash) != tokenFileDataHash {
		return "", errors.New("hash do arquivo de token inválido")
	}
	decryptedToken, decryptedTokenErr := DecryptData(tokenFileDataStr, string(tokenFileHash))
	if decryptedTokenErr != nil {
		return "", fmt.Errorf("erro ao descriptografar token: %w", decryptedTokenErr)
	}
	return decryptedToken, nil
}

func WatchKubexFiles() error {
	// watch --no-title --color -n 1 'bash -c "tput -x clear && tree $HOME/.cache/kubex -s --du -C -h && tree $HOME/.kubex -s --du -C -h"'
	watchCmd := exec.Command("watch", "--no-title", "--color", "-n", "1", "bash", "-c", "'tput -x clear && tree $HOME/.cache/kubex -s --du -C -h && tree $HOME/.kubex -s --du -C -h'")
	watchCmd.Stdout = os.Stdout
	watchCmd.Stderr = os.Stderr
	if watchCmdErr := watchCmd.Run(); watchCmdErr != nil {
		return fmt.Errorf("erro ao assistir arquivos Kubex: %v", watchCmdErr)
	}
	return nil
}

func CreateTarGz(archivePath string, files []string) error {
	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo tar.gz: %v", err)
	}
	defer func(archiveFile *os.File) {
		_ = archiveFile.Close()
	}(archiveFile)

	gw := gzip.NewWriter(archiveFile)
	defer func(gw *gzip.Writer) {
		_ = gw.Close()
	}(gw)

	tw := tar.NewWriter(gw)
	defer func(tw *tar.Writer) {
		_ = tw.Close()
	}(tw)

	for _, file := range files {
		addErr := AddFileToTar(tw, file)
		if addErr != nil {
			return addErr
		}
	}

	return nil
}

func AddFileToTar(tw *tar.Writer, filePath string) error {
	tempDir, tempDirErr := GetTempDir()
	if tempDirErr != nil {
		return tempDirErr
	}

	// Validar e sanitizar o caminho do arquivo
	cleanPath := filepath.Clean(filePath)
	if !strings.HasPrefix(cleanPath, tempDir) {
		return fmt.Errorf("caminho do arquivo inválido: %s", filePath)
	}

	file, err := os.Open(cleanPath)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo para o tar: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("erro ao obter informações do arquivo: %v", err)
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return fmt.Errorf("erro ao criar o cabeçalho do tar: %v", err)
	}
	header.Name = filepath.Base(cleanPath)

	err = tw.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("erro ao escrever o cabeçalho do tar: %v", err)
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return fmt.Errorf("erro ao copiar o arquivo para o tar: %v", err)
	}

	return nil
}

func ArchiveOldFiles() error {
	tempDir, tempDirErr := GetTempDir()
	if tempDirErr != nil {
		return tempDirErr
	}

	archiveName := fmt.Sprintf("logs_archive_%s.zip", time.Now().Format("20060102_150405"))
	archivePath := filepath.Join(tempDir, archiveName)

	zipFile, err := os.Create(archivePath)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo zip: %v", err)
	}
	defer func(zipFile *os.File) {
		_ = zipFile.Close()
	}(zipFile)

	zipWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) {
		_ = zipWriter.Close()
	}(zipWriter)

	files, err := os.ReadDir(tempDir)
	if err != nil {
		return fmt.Errorf("erro ao ler o diretório de logs: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".log") {
			filePath := filepath.Join(tempDir, file.Name())
			err := AddFileToZip(zipWriter, filePath)
			if err != nil {
				return err
			}
		}
	}

	fmt.Println("Logs arquivados com sucesso em:", archivePath)
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("erro ao obter informações do arquivo: %v", err)
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("erro ao criar o cabeçalho do arquivo zip: %v", err)
	}
	header.Name = filepath.Base(filePath)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("erro ao criar o cabeçalho do arquivo zip: %v", err)
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		return fmt.Errorf("erro ao copiar o arquivo para o zip: %v", err)
	}

	return nil
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo de origem: %v", err)
	}
	defer func(sourceFile *os.File) {
		_ = sourceFile.Close()
	}(sourceFile)

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo de destino: %v", err)
	}
	defer func(destFile *os.File) {
		_ = destFile.Close()
	}(destFile)

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("erro ao copiar o arquivo: %v", err)
	}

	return nil
}

func RemoveFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("erro ao remover o arquivo: %v", err)
	}
	return nil
}

func RemoveFiles(paths []string) error {
	for _, path := range paths {
		err := RemoveFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}
