package services

import (
	"fmt"
	"os"

	"github.com/docker/go-connections/nat"
	glb "github.com/rafa-mori/gdbase/internal/globals"
	gl "github.com/rafa-mori/gdbase/logger"
	t "github.com/rafa-mori/gdbase/types"
)

func SetupRabbitMQ(config *t.RabbitMQ, dockerService IDockerService) error {
	if config == nil || !config.Enabled {
		gl.Log("debug", "RabbitMQ está desabilitado na configuração. Ignorando inicialização.")
		return nil
	}

	// Verifica se o serviço já está rodando
	if IsServiceRunning(config.Reference.Name) {
		gl.Log("info", fmt.Sprintf("✅ RabbitMQ (%s) já está rodando!", config.Reference.Name))
		return nil
	}

	// Configura valores padrão, caso estejam ausentes
	if config.Username == "" {
		config.Username = "guest"
	}
	if config.Password == "" {
		config.Password = "guest"
	}
	if config.Port == nil || config.Port == "" {
		config.Port = "5672"
	}
	if config.ManagementPort == "" {
		config.ManagementPort = "15672"
	}
	if config.ErlangCookie == "" {
		config.ErlangCookie = "defaultcookie"
	}
	if config.Volume == "" {
		config.Volume = os.ExpandEnv(glb.DefaultRabbitMQVolume)
	}

	// Cria o volume, se necessário
	if err := dockerService.CreateVolume(config.Reference.Name, config.Volume); err != nil {
		gl.Log("error", fmt.Sprintf("❌ Erro ao criar volume do RabbitMQ: %v", err))
		return err
	}

	// Mapeia as portas
	portMap := []nat.PortMap{
		dockerService.MapPorts(fmt.Sprintf("%s", config.ManagementPort), "15672/tcp"),
		dockerService.MapPorts(fmt.Sprintf("%s", config.Port), "5672/tcp"),
	}

	// Configura as variáveis de ambiente
	envVars := []string{
		"RABBITMQ_DEFAULT_USER=" + config.Username,
		"RABBITMQ_DEFAULT_PASS=" + config.Password,
		"RABBITMQ_ERLANG_COOKIE=" + config.ErlangCookie,
	}

	// Inicializa o container do RabbitMQ
	service := dockerService.AddService(
		config.Reference.Name,
		"rabbitmq:3-management",
		envVars,
		portMap,
		map[string]struct{}{
			fmt.Sprintf("%s:/var/lib/rabbitmq", config.Volume): {},
		},
	)
	if service == nil {
		err := fmt.Errorf("serviço não encontrado: %s", config.Reference.Name)
		gl.Log("error", fmt.Sprintf("❌ Erro ao iniciar o RabbitMQ: %v", err))
		return err
	}

	gl.Log("success", fmt.Sprintf("✅ RabbitMQ (%s) iniciado com sucesso!", config.Reference.Name))
	return nil
}
