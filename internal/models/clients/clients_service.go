package clients

import (
	"errors"
	"fmt"
)

// IClientService define os métodos disponíveis para gerenciar clientes.
type IClientService interface {
	// Cria um cliente.
	CreateClient(client *ClientDetailed) (*ClientDetailed, error)
	// Obtém um cliente pelo ID.
	GetClientByID(id string) (*ClientDetailed, error)
	// Atualiza um cliente.
	UpdateClient(client *ClientDetailed) (*ClientDetailed, error)
	// Exclui um cliente pelo ID.
	DeleteClient(id string) error
	// Lista todos os clientes.
	ListClients() ([]*ClientDetailed, error)
}

// ClientService é a implementação de IClientService.
type ClientService struct {
	repo IClientRepo
}

// NewClientService cria uma nova instância de ClientService.
func NewClientService(repo IClientRepo) IClientService {
	return &ClientService{repo: repo}
}

func (cs *ClientService) CreateClient(client *ClientDetailed) (*ClientDetailed, error) {
	if client == nil {
		return nil, errors.New("ClientService: client is nil")
	}
	// Aqui você pode adicionar validações específicas do domínio.
	if client.ID == "" {
		return nil, errors.New("ClientService: client ID is empty")
	}
	createdClient, err := cs.repo.Create(client)
	if err != nil {
		return nil, fmt.Errorf("ClientService: error creating client: %w", err)
	}
	return createdClient, nil
}

func (cs *ClientService) GetClientByID(id string) (*ClientDetailed, error) {
	client, err := cs.repo.FindOne("id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("ClientService: error fetching client: %w", err)
	}
	return client, nil
}

func (cs *ClientService) UpdateClient(client *ClientDetailed) (*ClientDetailed, error) {
	updatedClient, err := cs.repo.Update(client)
	if err != nil {
		return nil, fmt.Errorf("ClientService: error updating client: %w", err)
	}
	return updatedClient, nil
}

func (cs *ClientService) DeleteClient(id string) error {
	err := cs.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("ClientService: error deleting client: %w", err)
	}
	return nil
}

func (cs *ClientService) ListClients() ([]*ClientDetailed, error) {
	clients, err := cs.repo.FindAll("1 = 1")
	if err != nil {
		return nil, fmt.Errorf("ClientService: error listing clients: %w", err)
	}
	return clients, nil
}
