package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// writeFileIfNotExists creates a file and writes default content if it does not exist.
func writeFileIfNotExists(filePath, content string) error {
	// Check if the file already exists
	if _, err := os.Stat(filePath); err == nil {
		return nil // File already exists, do nothing
	}

	// Create the file and write the content
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// appendToShellConfig adds a snippet to the shell configuration file if it's not already present.
func appendToShellConfig(shellConfig, snippet string) error {
	// Read the file content if it exists
	content, err := os.ReadFile(shellConfig)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Check if the snippet is already present
	if strings.Contains(string(content), snippet) {
		return nil // Snippet is already in the file, no need to modify
	}

	// Open the file for appending
	file, err := os.OpenFile(shellConfig, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append the snippet at the end of the file
	_, err = file.WriteString("\n" + snippet + "\n")
	return err
}

// removeSnippetFromShellConfig removes a specific snippet from the shell configuration file.
func removeSnippetFromShellConfig(shellConfig, snippet string) error {
	// Read the file content
	content, err := os.ReadFile(shellConfig)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File does not exist, nothing to remove
		}
		return err
	}

	// Convert content to a string and remove the snippet
	newContent := strings.ReplaceAll(string(content), "\n"+snippet+"\n", "")

	// If no changes, return early
	if newContent == string(content) {
		return nil
	}

	// Write the modified content back to the file
	return os.WriteFile(shellConfig, []byte(newContent), 0644)
}

// InstallBashHelpers creates the necessary scripts and adds them to the shell startup files.
func InstallBashHelpers() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Environment files to be created
	filesToCreate := map[string]string{
		filepath.Join(homeDir, ".mori_utils_env"):       cdxBashEnvTemplate,
		filepath.Join(homeDir, ".mori_logging_env"):     cdxBashLogTemplate,
		filepath.Join(homeDir, ".mori_yes_no_question"): cdxYesNoQyestionTemplate,
	}

	// Create files if they do not exist
	for filePath, content := range filesToCreate {
		if err := writeFileIfNotExists(filePath, content); err != nil {
			fmt.Println("Error creating file:", filePath, err)
		}
	}

	// Snippet to be added to .bashrc or .zshrc
	snippet := `if test -f "$HOME/.mori_utils_env"; then
	. "$HOME/.mori_utils_env"
fi

if test -f "$HOME/.mori_logging_env"; then
	. "$HOME/.mori_logging_env"
fi
`
	// Detect which shell is being used
	shell := os.Getenv("SHELL")
	shellConfig := filepath.Join(homeDir, ".bashrc") // Default for Bash

	if strings.HasSuffix(shell, "zsh") {
		shellConfig = filepath.Join(homeDir, ".zshrc")
	}

	// Add the snippet to the shell configuration file
	if err := appendToShellConfig(shellConfig, snippet); err != nil {
		fmt.Println("Error modifying shell configuration file:", err)
	} else {
		fmt.Println("Configuration successfully added to", shellConfig)
	}

	// Determina o arquivo de configuração correto do shell
	shellRC := fmt.Sprintf("$HOME/.$(basename $SHELL)rc")

	// Comando a ser injetado
	command := fmt.Sprintf(". %s\n", shellRC)

	// Obtém o terminal do usuário (stdin ativo)
	tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	if err != nil {
		fmt.Println("⚠️ Não foi possível acessar o terminal:", err)
		fmt.Printf("Por favor, execute manualmente: . %s\n", shellRC)
		return
	}
	defer tty.Close()

	// Aguarda um breve momento para garantir que a saída foi impressa antes da injeção
	time.Sleep(500 * time.Millisecond)

	// Escreve o comando diretamente no terminal do usuário
	_, wrtErr := tty.WriteString(command)
	if wrtErr != nil {
		return
	}
}

// UninstallBashHelpers removes the scripts and their references from the shell startup files.
func UninstallBashHelpers() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Files to be removed
	filesToRemove := []string{
		filepath.Join(homeDir, ".mori_utils_env"),
		filepath.Join(homeDir, ".mori_logging_env"),
		filepath.Join(homeDir, ".mori_yes_no_question"),
		filepath.Join(filepath.Join(homeDir, ".cache"), ".mori_install_log"),
		filepath.Join(filepath.Join(homeDir, ".cache"), ".mri_usage_warning"),
	}

	// Remove files
	for _, filePath := range filesToRemove {
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			fmt.Println("Error removing file:", filePath, err)
		}
	}

	// Snippet to be removed from .bashrc or .zshrc
	snippet := `if test -f "$HOME/.mori_utils_env"; then
	. "$HOME/.mori_utils_env"
fi

if test -f "$HOME/.mori_logging_env"; then
	. "$HOME/.mori_logging_env"
fi
`
	// Detect which shell is being used
	shell := os.Getenv("SHELL")
	shellConfig := filepath.Join(homeDir, ".bashrc") // Default for Bash

	if strings.HasSuffix(shell, "zsh") {
		shellConfig = filepath.Join(homeDir, ".zshrc")
	}

	// Remove the snippet from the shell configuration file
	if err := removeSnippetFromShellConfig(shellConfig, snippet); err != nil {
		fmt.Println("Error modifying shell configuration file:", err)
	} else {
		fmt.Println("Configuration successfully removed from", shellConfig)
	}

	fmt.Println("Bash helpers successfully uninstalled.")
}
