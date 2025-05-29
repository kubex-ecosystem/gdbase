# 🏛️ **GDBASE - Infraestrutura de Banco de Dados Modular** 

</br>

## 🔥 **Visão Geral**  

**GDBASE** é uma solução de gerenciamento de bancos de dados desenvolvida em Go, projetada para ser **modular, escalável e automática**. Ele permite **configuração zero**, mas também suporta customizações avançadas via arquivos de configuração.  

Além de gerenciar bancos locais e baseados em Docker, **ele suporta conexões simultâneas com múltiplos bancos de dados**, tornando-se uma solução robusta para sistemas distribuídos.  

## 🔗 **Recursos Principais**  

✅ **Configuração dinâmica e automática** → Senhas são geradas randômicamente e armazenadas no keyring, enquanto portas ocupadas são ajustadas automaticamente.  
✅ **Compatível com múltiplos DBs** → Redis, RabbitMQ, MongoDB, PostgreSQL e SQLite já configurados para uso imediato.  
✅ **Estrutura arquitetural bem definida** → Models seguem padrão `Model → Repo → Service`, garantindo modularidade e organização.  
✅ **Túnel SSH para bancos externos** → `gdbase ssh tunnel` permite conectar bancos remotos via SSH com segurança total.  
✅ **Customização avançada** → Pode ser configurado manualmente via arquivos ou rodar sem necessidade de intervenção.  
✅ **Orquestração via Docker** → Geração automática de containers, garantindo portabilidade e fácil implantação.  
✅ **Monitoramento e eventos** → Implementação de event bus para rastreamento interno de ações no sistema.  

## 📝 **Instalação**  

Clone o repositório e compile:  

```sh
git clone https://github.com/rafa-mori/gdbase.git
cd gdbase
go build -o gdbase .
```

## 🚀 **Iniciando o Servidor**  

Para inicializar **GDBASE**, basta rodar:  

```sh
./gdbase start
```

Isso **configura os bancos de dados, ajusta conexões e inicializa todos os serviços automaticamente**.  

## 🔎 **Comandos da CLI**  

A CLI integrada permite comandos avançados, incluindo **gerenciamento de bancos de dados e configurações de rede**:  

```sh
./gdbase --help
```

## 💡 **Resumo dos comandos essenciais:**  

| Comando           | Função                                             |
|-------------------|----------------------------------------------------|
| `start`           | Inicializa `gdbase` e configura todos os serviços  |
| `status`          | Exibe status dos bancos de dados ativos            |
| `config`          | Cria um arquivo de configuração para customização  |
| `ssh tunnel`      | Cria um túnel seguro para bancos externos via SSH  |
| `docker`          | Gerencia containers Docker para bancos de dados    |

## 📂 **Estrutura do Projeto (Implementação central)**

``` plaintext

./
│
├── cmd
│   ├── cli
│   │   ├── cmds_ssh.go
│   │   ├── cmds_utils.go
│   │   ├── common.go
│   │   ├── configs.go
│   │   ├── database.go
│   │   ├── docker.go
│   │   ├── serverdata_types.go
│   │   ├── webprocess_types.go
│   │   ├── webserver_types.go
│   │   ├── websrvproc_types.go
│   │   └── websrvstatus_types.go
│   │
│   ├── gen_models.go # Extrai do banco de dados as informações necessárias para criar  as structs.
│   ├── models.go # Esse é o arquivo gerado pelo anterior com as definições das structs.(sem erros. roda lisinho, porém não uso ainda)
│   │
│   ├── main.go
│   ├── usage.go
│   └── wrpr.go
│
├── docs
│   ├── assets
│   │   └── top_banner.png
│   └── CONTRIBUTING.md
│
├── go.mod
├── go.sum
│
├── internal
│   ├── events
│   │   ├── dispatcher.go
│   │   ├── event_bus.go
│   │   ├── event.go
│   │   └── screenning.go
│   ├── models
│   │   ├── clients
│   │   │   ├── client_model.go
│   │   │   ├── clients_repo.go
│   │   │   └── clients_service.go
│   │   ├── cron
│   │   │   ├── common.go
│   │   │   ├── cronjob_model.go
│   │   │   ├── cronjob_repo.go
│   │   │   └── cronjob_service.go
│   │   ├── job_queue
│   │   │   ├── job_queue.go
│   │   │   ├── job_repo.go
│   │   │   └── job_service.go
│   │   ├── notification
│   │   │   ├── notification_model.go
│   │   │   ├── notification_repo.go
│   │   │   └── notification_service.go
│   │   ├── orders
│   │   │   ├── order_model.go
│   │   │   ├── order_repo.go
│   │   │   ├── order_service.go
│   │   │   └── order_status.go
│   │   ├── products
│   │   │   ├── product_model.go
│   │   │   ├── products_repo.go
│   │   │   └── products_service.go
│   │   ├── users
│   │   │   ├── user_model.go
│   │   │   ├── user_repo.go
│   │   │   └── user_service.go
│   │   └── webhooks
│   │       ├── webhook_model.go
│   │       ├── webhook_repo.go
│   │       └── webhook_service.go
│   │
│   └── services
│       ├── assets
│       │   ├── ddl_full_model.sql
│       │   └── init-db.sql
│       ├── broker_info.go
│       ├── broker_manager.go
│       ├── broker_server.go
│       ├── db_service.go
│       ├── docker_client.go
│       ├── docker_service.go
│       ├── execution_log_service.go
│       ├── rabbitmq.go
│       └── utils.go
│
├── tests
│   ├── client_model_test.go
│   ├── client_repo_test.go
│   ├── database_test.go
│   ├── dkr_abs_test.go
│   ├── execution_log_test.go
│   ├── order_repo_service_test.go
│   ├── product_model_test.go
│   ├── product_repo_test.go
│   ├── user_model_crud_test.go
│   ├── user_model_test.go
│   ├── user_repo_test.go
│   └── user_service_test.go
│
└── version
    ├── CLI_VERSION
    └── semantic.go

```
