/* 
  Versão 1.0
  Author: Rafael Mori
  Description: Script de inicialização do banco de dados para o serviços diversos (comercial, MCP, etc.)
 */
 
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";  -- Para gerar UUIDs
CREATE EXTENSION IF NOT EXISTS "pg_trgm";  -- Para buscas de texto eficientes
CREATE EXTENSION IF NOT EXISTS "btree_gist";  -- Para índices GIN em tipos de dados não nativos
CREATE EXTENSION IF NOT EXISTS "fuzzystrmatch";  -- Para comparação de strings
CREATE EXTENSION IF NOT EXISTS "hstore";  -- Para armazenar pares chave-valor
-- COMMIT;

-- Criação de roles e usuários
CREATE ROLE readonly;
CREATE ROLE readwrite;
CREATE ROLE admin;
-- COMMIT;

-- Criação de usuários e atribuição de roles
CREATE USER user_readonly WITH PASSWORD 'readonlypass';
CREATE USER user_readwrite WITH PASSWORD 'readwritepass';
CREATE USER user_admin WITH PASSWORD 'adminpass';
-- COMMIT;

GRANT readonly TO user_readonly;
GRANT readwrite TO user_readwrite;
GRANT admin TO user_admin;
-- COMMIT;

-- Permissões para roles
GRANT CONNECT ON DATABASE kubex_db TO readonly, readwrite, admin;
GRANT USAGE ON SCHEMA public TO readonly, readwrite, admin;
-- COMMIT;

GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO readwrite;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin;
-- COMMIT;

-- Enums
CREATE TYPE inventory_status AS ENUM ('available', 'reserved', 'damaged', 'expired');
CREATE TYPE order_status AS ENUM ('draft', 'pending', 'confirmed', 'processing', 'shipped', 'delivered', 'cancelled');
CREATE TYPE payment_status AS ENUM ('pending', 'paid', 'failed', 'refunded');
CREATE TYPE confidence_level AS ENUM ('high', 'medium', 'low');
CREATE TYPE address_type AS ENUM ('billing', 'shipping', 'both');
CREATE TYPE address_status AS ENUM ('active', 'inactive', 'archived');
CREATE TYPE cron_type AS ENUM('cron', 'interval');
CREATE TYPE http_method AS ENUM('GET', 'POST', 'PUT', 'DELETE');
CREATE TYPE last_run_status AS ENUM('success', 'failure', 'pending', 'running', 'completed');
CREATE TYPE last_run_message AS ENUM('success', 'failure', 'pending', 'running', 'completed');
CREATE TYPE job_status AS ENUM('SUCCESS', 'FAILED', 'PENDING', 'RUNNING', 'COMPLETED');
CREATE TYPE job_type AS ENUM('cron', 'interval');
CREATE TYPE job_command_type AS ENUM('shell', 'api', 'script');
-- COMMIT;

-- Tabela de roles
CREATE TABLE IF NOT EXISTS roles (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(50) NOT NULL UNIQUE,
    description text,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

-- Tabela de permissões
CREATE TABLE IF NOT EXISTS permissions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(50) NOT NULL UNIQUE,
    description text,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

-- Tabela de permissões por role
CREATE TABLE IF NOT EXISTS role_permissions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id uuid NOT NULL REFERENCES roles(id),
    permission_id uuid NOT NULL REFERENCES permissions(id),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    UNIQUE (role_id, permission_id)
);

-- Tabela de usuários
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(255) NOT NULL,
    username varchar(50) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    email varchar(100) NOT NULL UNIQUE,
    phone varchar(30),
    document varchar(50),
    role_id uuid REFERENCES roles(id),
    active boolean NOT NULL DEFAULT true,   
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    last_login timestamp without time zone
);

-- Tabela de cron jobs
-- Esta tabela é responsável por armazenar as tarefas agendadas
CREATE TABLE IF NOT EXISTS cron_jobs (
    id               uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name             VARCHAR(255), -- Nome da tarefa
    description      TEXT, -- Descrição da tarefa

    cron_type         cron_type DEFAULT 'cron', -- Tipo de agendamento (cron ou intervalo)
    cron_expression  TEXT DEFAULT '2 * * * *', -- Expressão cron (se for cron)

    starts_at         TIMESTAMP DEFAULT NOW(), -- Hora de início
    ends_at           TIMESTAMP DEFAULT NULL, -- Hora de término

    command          TEXT, -- Comando a ser executado (caso seja via CLI)
    method           http_method, -- Tipo de requisição (se for API)
    api_endpoint     VARCHAR(255), -- URL do endpoint (se for API)
    payload          JSONB, -- Dados que precisam ser enviados na request
    headers          JSONB, -- Cabeçalhos que precisam ser enviados na request
    retries          INTEGER DEFAULT 0, -- Número de tentativas

    exec_timeout     INTEGER DEFAULT 30, -- Tempo máximo de execução (em segundos)
    max_retries      INTEGER DEFAULT 3, -- Número máximo de tentativas
    retry_interval   INTEGER DEFAULT 10, -- Intervalo entre tentativas (em segundos)
    max_execution_time INTEGER DEFAULT 300, -- Tempo máximo de execução (em segundos)

    last_run_status   last_run_status DEFAULT 'pending', -- Status da última execução
    last_run_message TEXT DEFAULT NULL, -- Mensagem da última execução
    last_run_time    TIMESTAMP DEFAULT NULL, -- Hora da última execução

    is_recurring     BOOLEAN DEFAULT FALSE, -- Se a tarefa é recorrente
    is_active        BOOLEAN DEFAULT TRUE, -- Status do job (ativo ou pausado)

    created_at       TIMESTAMP DEFAULT NOW(), -- Data de criação
    updated_at       TIMESTAMP DEFAULT NOW(), -- Última modificação
    last_executed_at TIMESTAMP DEFAULT NULL, -- Última vez que foi executado com sucesso

    user_id         uuid REFERENCES users(id), -- Referência ao usuário que criou o job
    created_by      uuid REFERENCES users(id), -- Referência ao usuário que criou o job
    updated_by      uuid REFERENCES users(id), -- Referência ao usuário que atualizou o job
    last_executed_by uuid REFERENCES users(id), -- Referência ao usuário que executou o job pela última vez

    metadata        JSONB -- Metadados adicionais
);

CREATE TABLE IF NOT EXISTS execution_logs (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    cronjob_id    UUID REFERENCES cron_jobs(id),
    execution_time TIMESTAMP DEFAULT NOW(),
    status       job_status DEFAULT 'PENDING',
    output       TEXT DEFAULT NULL,
    error_message TEXT DEFAULT NULL,
    retry_count  INTEGER DEFAULT 0,
    created_at   TIMESTAMP DEFAULT NOW(), 
    updated_at   TIMESTAMP DEFAULT NOW(), 
    user_id     uuid REFERENCES users(id),
    created_by   uuid REFERENCES users(id),
    updated_by   uuid REFERENCES users(id),
    metadata     JSONB 
);

CREATE TABLE IF NOT EXISTS job_queue (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    cronjob_id    UUID REFERENCES cron_jobs(id),
    status        job_status DEFAULT 'PENDING', 
    scheduled_time TIMESTAMP DEFAULT NOW(),
    execution_time TIMESTAMP DEFAULT NULL,
    error_message TEXT DEFAULT NULL,
    retry_count   INTEGER DEFAULT 0,
    next_run_time TIMESTAMP DEFAULT NULL, 
    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP DEFAULT NOW(),
    metadata      JSONB DEFAULT NULL,
    user_id      uuid REFERENCES users(id), 
    created_by    uuid REFERENCES users(id),
    updated_by    uuid REFERENCES users(id),
    last_executed_by uuid REFERENCES users(id),
    job_type       job_type DEFAULT 'cron',
    job_expression TEXT DEFAULT '2 * * * *',
    job_command   TEXT,
    job_method    http_method,
    job_api_endpoint VARCHAR(255),
    job_payload   JSONB,
    job_headers   JSONB,
    job_retries   INTEGER DEFAULT 0,
    job_timeout   INTEGER DEFAULT 0
);


-- COMMIT;

-- Tabela de endereços (abstrata, reutilizável)
CREATE TABLE IF NOT EXISTS addresses (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    external_id varchar(255),
    external_code varchar(255),
    street varchar(255) NOT NULL,
    number varchar(50) NOT NULL,
    complement varchar(100),
    district varchar(100),
    city varchar(100) NOT NULL,
    state varchar(50) NOT NULL,
    country varchar(50) NOT NULL,
    zip_code varchar(20) NOT NULL,
    is_main boolean,
    is_active boolean NOT NULL DEFAULT true,
    notes text,
    latitude numeric(10,6),
    longitude numeric(10,6),
    address_type varchar(20),
    address_status varchar(20),
    address_category varchar(20),
    address_tags text[],
    is_default boolean,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    last_sync_at timestamp without time zone
);
-- COMMIT;

-- Tabela de logs de auditoria
CREATE TABLE IF NOT EXISTS audit_logs (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid REFERENCES users(id),
    action varchar(50) NOT NULL,
    entity_type varchar(50) NOT NULL,
    entity_id uuid NOT NULL,
    changes jsonb,
    created_at timestamp without time zone NOT NULL DEFAULT now()
);
-- COMMIT;

-- Tabela de logs de erro
CREATE TABLE IF NOT EXISTS error_logs (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    error_message text NOT NULL,
    stack_trace text,
    created_at timestamp without time zone NOT NULL DEFAULT now()
);
-- COMMIT;

-- Tabela de logs de acesso
CREATE TABLE IF NOT EXISTS access_logs (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid REFERENCES users(id),
    action varchar(50) NOT NULL,
    ip_address varchar(50),
    user_agent text,
    created_at timestamp without time zone NOT NULL DEFAULT now()
);
-- COMMIT;

-- Tabela de categorias de produto
CREATE TABLE IF NOT EXISTS product_category (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(255) NOT NULL,
    parent_id uuid REFERENCES product_category(id),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

-- Tabela de produtos
CREATE TABLE IF NOT EXISTS products (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    external_id varchar(255),
    sku varchar(100) NOT NULL UNIQUE,
    barcode varchar(50),
    default_vol_type varchar(50),
    name varchar(255) NOT NULL,
    description text,
    category_id uuid NOT NULL REFERENCES product_category(id),
    manufacturer varchar(255),
    image_url varchar(255),
    image text,
    brand varchar(100),
    price numeric(18,2) NOT NULL,
    cost numeric(18,2),
    weight numeric(10,3),
    length numeric(10,3),
    width numeric(10,3),
    height numeric(10,3),
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    last_sync_at timestamp without time zone,
    min_stock_threshold integer,
    max_stock_threshold integer,
    reorder_point integer,
    lead_time_days integer,
    shelf_life_days integer,
    search_vector tsvector
);

-- COMMIT;

-- Índices para produtos
CREATE INDEX IF NOT EXISTS idx_product_sku ON products(sku);
CREATE INDEX IF NOT EXISTS idx_product_barcode ON products(barcode);
CREATE INDEX IF NOT EXISTS idx_product_name ON products(name);
CREATE INDEX IF NOT EXISTS idx_product_category_id ON products(category_id);
CREATE INDEX IF NOT EXISTS idx_product_manufacturer ON products(manufacturer);
CREATE INDEX IF NOT EXISTS idx_product_search_vector ON products USING GIN(search_vector);
-- COMMIT;

-- Trigger para atualizar o campo search_vector
CREATE OR REPLACE FUNCTION update_product_search_vector() RETURNS TRIGGER AS $$
BEGIN
    NEW.search_vector :=
        setweight(to_tsvector('portuguese', COALESCE(NEW.name, '')), 'A') ||
        setweight(to_tsvector('portuguese', COALESCE(NEW.sku, '')), 'A') ||
        setweight(to_tsvector('portuguese', COALESCE(NEW.barcode, '')), 'A') ||
        setweight(to_tsvector('portuguese', COALESCE(NEW.description, '')), 'C') ||
        setweight(to_tsvector('portuguese', COALESCE(NEW.brand, '')), 'B') ||
        setweight(to_tsvector('portuguese', COALESCE(NEW.manufacturer, '')), 'B');
    RETURN NEW;
END
$$ LANGUAGE plpgsql;
-- COMMIT;

DROP TRIGGER IF EXISTS trigger_update_product_search_vector ON products;
CREATE TRIGGER trigger_update_product_search_vector
    BEFORE INSERT OR UPDATE ON products
    FOR EACH ROW EXECUTE FUNCTION update_product_search_vector();
-- COMMIT;

-- Tabela de parceiros
CREATE TABLE IF NOT EXISTS partners (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    external_id varchar(255),
    code varchar(100) NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    trade_name varchar(255),
    document varchar(50) NOT NULL,
    type varchar(20) NOT NULL CHECK (type IN ('individual','company')),
    category varchar(50) CHECK (category IN ('SUPERMERCADO','LOJA_DE_COSMETICOS','FARMACIA','ATACAREJO')),
    status varchar(20) NOT NULL CHECK (status IN ('ACTIVE','INACTIVE','BLOCKED','ARCHIVED')),
    region varchar(100),
    segment varchar(100),
    size varchar(10),
    address_ids uuid[] NOT NULL,
    credit_limit numeric(18,2),
    current_debt numeric(18,2),
    payment_terms text[],
    last_purchase_date timestamp without time zone,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    is_active boolean NOT NULL DEFAULT true
);

-- Tabela de contatos do parceiro
CREATE TABLE IF NOT EXISTS partner_contact (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    partner_id uuid NOT NULL REFERENCES partners(id) ON DELETE CASCADE,
    name varchar(100) NOT NULL,
    email varchar(100),
    phone varchar(30),
    position varchar(50),
    is_primary boolean NOT NULL DEFAULT false
);

-- Tabela de histórico de vendas do parceiro
CREATE TABLE IF NOT EXISTS partner_sales_history (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    partner_id uuid NOT NULL REFERENCES partners(id) ON DELETE CASCADE,
    year integer NOT NULL,
    q1 integer NOT NULL DEFAULT 0,
    q2 integer NOT NULL DEFAULT 0,
    q3 integer NOT NULL DEFAULT 0,
    q4 integer NOT NULL DEFAULT 0
);
-- COMMIT;

-- Tabela de armazéns
CREATE TABLE IF NOT EXISTS warehouses (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(255) NOT NULL,
    location varchar(255),
    capacity integer,
    current_stock integer,
    manager varchar(100),
    contact varchar(100),
    address_id uuid REFERENCES addresses(id),
    external_id varchar(255),
    external_code varchar(255),
    notes text,
    tags text[],
    status varchar(50),
    created_by uuid REFERENCES users(id),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_by uuid REFERENCES users(id),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    last_sync_at timestamp without time zone,
    is_active boolean NOT NULL DEFAULT true
);

-- Tabela de estoque
CREATE TABLE IF NOT EXISTS inventory (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id uuid NOT NULL REFERENCES products(id),
    warehouse_id uuid NOT NULL REFERENCES warehouses(id),
    quantity numeric(18,3) NOT NULL DEFAULT 0,
    minimum_level numeric(18,3),
    maximum_level numeric(18,3),
    reorder_point numeric(18,3),
    last_count_date timestamp without time zone,
    status varchar(50),
    vol_type varchar(50),
    lot_control varchar(100),
    expiration_date timestamp without time zone,
    location_code varchar(100),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    last_sync_at timestamp without time zone,
    is_active boolean NOT NULL DEFAULT true
);
-- COMMIT;

-- Tabela de previsões de estoque
CREATE TABLE IF NOT EXISTS stock_predictions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id uuid NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    warehouse_id uuid NOT NULL REFERENCES warehouses(id) ON DELETE CASCADE,
    current_level NUMERIC(18,3) NOT NULL,
    predicted_level NUMERIC(18,3) NOT NULL,
    days_until_stockout INTEGER,
    confidence_level VARCHAR(10) CHECK (confidence_level IN ('high','medium','low')) NOT NULL,
    suggested_reorder_quantity NUMERIC(18,3),
    prediction_date TIMESTAMP NOT NULL DEFAULT NOW(),
    prediction_horizon_days INTEGER NOT NULL DEFAULT 30,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_prediction UNIQUE (product_id, warehouse_id, prediction_date, prediction_horizon_days)
);
-- COMMIT;

-- Tabela de pedidos
-- CREATE TABLE IF NOT EXISTS orders (
--     id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
--     external_id varchar(255),
--     order_number varchar(100),
--     partner_id uuid NOT NULL REFERENCES partners(id),
--     status varchar(30) NOT NULL,
--     order_date timestamp without time zone NOT NULL,
--     estimated_delivery_date timestamp without time zone,
--     actual_delivery_date timestamp without time zone,
--     shipping_address_id uuid REFERENCES addresses(id),
--     payment_method varchar(30),
--     payment_status varchar(20),
--     notes text,
--     total_amount numeric(18,2) NOT NULL,
--     discount_amount numeric(18,2) NOT NULL DEFAULT 0,
--     tax_amount numeric(18,2),
--     shipping_amount numeric(18,2),
--     final_amount numeric(18,2) NOT NULL,
--     is_automatically_generated boolean DEFAULT false,
--     created_at timestamp without time zone NOT NULL DEFAULT now(),
--     updated_at timestamp without time zone NOT NULL DEFAULT now(),
--     last_sync_at timestamp without time zone,
--     --prediction_id uuid REFERENCES stock_predictions(id),
--     priority integer, 
--     expected_margin numeric(18,2)
-- );
-- COMMIT;

-- Tabela de configurações de sincronização
CREATE TABLE IF NOT EXISTS sync_config (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_name VARCHAR(100) NOT NULL,
    last_sync_timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    sync_interval_minutes INTEGER NOT NULL DEFAULT 60,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    error_count INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Tabela de logs de sincronização
CREATE TABLE IF NOT EXISTS sync_logs (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_name VARCHAR(100) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    status VARCHAR(50) NOT NULL,
    records_processed INTEGER,
    records_created INTEGER,
    records_updated INTEGER,
    records_failed INTEGER,
    error_message TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- COMMIT;

-- Tabela de dados de previsão diária (para armazenar séries temporais)
CREATE TABLE IF NOT EXISTS prediction_daily_data (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    --prediction_id uuid NOT NULL REFERENCES stock_predictions(id) ON DELETE CASCADE,
    day_date DATE NOT NULL,
    predicted_demand NUMERIC(18,3) NOT NULL,
    predicted_stock NUMERIC(18,3) NOT NULL,
    lower_bound NUMERIC(18,3),
    upper_bound NUMERIC(18,3),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
    -- CONSTRAINT unique_prediction_day UNIQUE (prediction_id, day_date)
);

-- Tabela de configurações de usuários
CREATE TABLE IF NOT EXISTS user_preferences (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    preference_type VARCHAR(50) NOT NULL CHECK (preference_type IN ('notification', 'theme', 'language')),
    preference_value_type VARCHAR(50) NOT NULL CHECK (preference_value_type IN ('string', 'boolean', 'integer', 'float')),
    preference_key VARCHAR(100) NOT NULL,
    preference_value TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_user_preference UNIQUE (user_id, preference_key)
);

-- Tabela de eventos de auditoria
CREATE TABLE IF NOT EXISTS audit_events (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    entity_type VARCHAR(100) NOT NULL,
    entity_id uuid NOT NULL,
    action VARCHAR(50) NOT NULL, -- create, update, delete
    user_id uuid NOT NULL REFERENCES users(id),
    changes JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- COMMIT;

-- Tabela de tokens de sessão
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_id varchar(255) NOT NULL UNIQUE,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);
-- COMMIT;

-- Índices para tabelas auxiliares
CREATE INDEX IF NOT EXISTS idx_sync_logs_entity_name ON sync_logs(entity_name);
-- CREATE INDEX IF NOT EXISTS idx_stock_predictions_product_id ON stock_predictions(product_id);
-- CREATE INDEX IF NOT EXISTS idx_stock_predictions_warehouse_id ON stock_predictions(warehouse_id);
-- CREATE INDEX IF NOT EXISTS idx_stock_predictions_days_until_stockout ON stock_predictions(days_until_stockout);
-- CREATE INDEX IF NOT EXISTS idx_stock_predictions_confidence_level ON stock_predictions(confidence_level);
-- CREATE INDEX IF NOT EXISTS idx_prediction_daily_data_prediction_id ON prediction_daily_data(prediction_id);
-- CREATE INDEX IF NOT EXISTS idx_prediction_daily_data_day_date ON prediction_daily_data(day_date);
CREATE INDEX IF NOT EXISTS idx_user_preferences_user_id ON user_preferences(user_id);
CREATE INDEX IF NOT EXISTS idx_audit_events_entity_type_id ON audit_events(entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_audit_events_created_at ON audit_events(created_at);
CREATE INDEX IF NOT EXISTS idx_audit_events_user_id ON audit_events(user_id);
-- COMMIT;

-- Artefatos básicos iniciais
INSERT INTO product_category (id, name, created_at, updated_at) VALUES (uuid_generate_v4(), 'Default', now(), now()) ON CONFLICT DO NOTHING;
INSERT INTO addresses (id, street, number, city, state, country, zip_code, is_active, created_at, updated_at) VALUES (uuid_generate_v4(), 'Rua Exemplo', '100', 'Cidade', 'UF', 'Brasil', '00000-000', true, now(), now()) ON CONFLICT DO NOTHING;
INSERT INTO partners (id, code, name, document, type, status, address_ids, is_active, created_at, updated_at) VALUES (uuid_generate_v4(), 'P0001', 'Parceiro Exemplo', '00.000.000/0001-00', 'company', 'ACTIVE', ARRAY[(SELECT id FROM addresses LIMIT 1)], true, now(), now()) ON CONFLICT DO NOTHING;
INSERT INTO warehouses (id, name, address_id, is_active, created_at, updated_at) VALUES (uuid_generate_v4(), 'Armazém Central', (SELECT id FROM addresses LIMIT 1), true, now(), now()) ON CONFLICT DO NOTHING;
-- COMMIT;


-- Inserting default roles
INSERT INTO roles (id, name, description, created_at, updated_at) VALUES
(uuid_generate_v4(), 'admin', 'Administrator with full access', now(), now()),
(uuid_generate_v4(), 'editor', 'Editor with modification permissions', now(), now()),
(uuid_generate_v4(), 'viewer', 'User with read-only access', now(), now());

-- Inserting default permissions
INSERT INTO permissions (id, name, description, created_at, updated_at) VALUES
(uuid_generate_v4(), 'user_create', 'Permission to create users', now(), now()),
(uuid_generate_v4(), 'user_edit', 'Permission to edit users', now(), now()),
(uuid_generate_v4(), 'user_delete', 'Permission to delete users', now(), now()),
(uuid_generate_v4(), 'post_create', 'Permission to create posts', now(), now()),
(uuid_generate_v4(), 'post_edit', 'Permission to edit posts', now(), now()),
(uuid_generate_v4(), 'post_delete', 'Permission to delete posts', now(), now());

-- Associating permissions with roles
INSERT INTO role_permissions (id, role_id, permission_id, created_at, updated_at) VALUES
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM permissions WHERE name = 'user_create'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM permissions WHERE name = 'user_edit'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM permissions WHERE name = 'user_delete'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM permissions WHERE name = 'post_create'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM permissions WHERE name = 'post_edit'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM permissions WHERE name = 'post_delete'), now(), now());

INSERT INTO role_permissions (id, role_id, permission_id, created_at, updated_at) VALUES
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'editor'), (SELECT id FROM permissions WHERE name = 'post_create'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'editor'), (SELECT id FROM permissions WHERE name = 'post_edit'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'editor'), (SELECT id FROM permissions WHERE name = 'post_delete'), now(), now());

INSERT INTO role_permissions (id, role_id, permission_id, created_at, updated_at) VALUES
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'viewer'), (SELECT id FROM permissions WHERE name = 'user_create'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'viewer'), (SELECT id FROM permissions WHERE name = 'user_edit'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'viewer'), (SELECT id FROM permissions WHERE name = 'user_delete'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'viewer'), (SELECT id FROM permissions WHERE name = 'post_create'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'viewer'), (SELECT id FROM permissions WHERE name = 'post_edit'), now(), now()),
(uuid_generate_v4(), (SELECT id FROM roles WHERE name = 'viewer'), (SELECT id FROM permissions WHERE name = 'post_delete'), now(), now());

-- Criando um usuário de exemplo
INSERT INTO "users" ("id","name","username","password","email","phone","role_id","document","active","created_at","updated_at","last_login") 
VALUES (
    uuid_generate_v4(),
    'Test User',
    'testUser',
    '$2a$10$24gpz0aVeuDarfmgwZlZoeJufrxAVKUsw5MjpfHlFN576I.gz.oSW',
    'abcdef',
    '9898989898',
    CASE WHEN (SELECT id FROM roles WHERE name = 'admin') IS NOT NULL THEN
        (SELECT id FROM roles WHERE name = 'admin') 
    ELSE
        (SELECT id FROM roles WHERE name = 'editor')
    END,
    '22CBCA1346796431',
    true,
    now(),
    now(),
    now()
)
ON CONFLICT ("username") DO UPDATE
SET "name" = 'TestUser',
    "email" = 'abcdef@test.com',
    "phone" = '9898989898',
    "role_id" = CAST('06ccc24a-4385-4f66-b528-5f8098c8e22d' as uuid),
    "document" = '22CBCA1346796431',
    "active" = true,
    "updated_at" = now()
RETURNING id;
-- COMMIT;

-----------------------------------------------------------------------------------

-- Tabela para LLM Models
CREATE TABLE IF NOT EXISTS mcp_llm_models (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    enabled boolean DEFAULT true,
    provider TEXT NOT NULL,
    model TEXT NOT NULL,
    temperature REAL DEFAULT 0.7,
    max_tokens INTEGER DEFAULT 2048,
    top_p REAL DEFAULT 1.0,
    frequency_penalty REAL DEFAULT 0.0,
    presence_penalty REAL DEFAULT 0.0,
    stop_sequences TEXT[],
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    created_by uuid REFERENCES users(id),
    updated_by uuid REFERENCES users(id),
    UNIQUE(provider, model)
);
-- COMMIT;

-- Tabela para Tokens
-- CREATE TABLE IF NOT EXISTS mcp_tokens (
--     id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
--     provider TEXT NOT NULL,
--     user_identifier TEXT NOT NULL,
--     access_token TEXT NOT NULL,
--     refresh_token TEXT,
--     expires_at TIMESTAMP,
--     scopes TEXT[],
--     created_at TIMESTAMP DEFAULT now(),
--     updated_at TIMESTAMP DEFAULT now(),
--     UNIQUE(provider, user_identifier)
-- );
-- COMMIT;

-- Tabela de preferências (flexível e armazenada em JSONB)
CREATE TABLE IF NOT EXISTS mcp_user_preferences (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    scope TEXT NOT NULL DEFAULT 'defaults',
    config JSONB NOT NULL,
    updated_at TIMESTAMP DEFAULT now(),
    updated_by uuid REFERENCES users(id),
    created_at TIMESTAMP DEFAULT now(),
    created_by uuid REFERENCES users(id),
    UNIQUE(scope)
);
-- COMMIT;

-- Tabela para provider_configurations (por ex: GitHub, GitLab, etc.)
CREATE TABLE IF NOT EXISTS mcp_provider_configs (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    provider TEXT NOT NULL,
    org_or_group TEXT NOT NULL,
    config JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    created_by uuid REFERENCES users(id),
    updated_at TIMESTAMP DEFAULT now(),
    updated_by uuid REFERENCES users(id),
    UNIQUE(provider, org_or_group)
);
-- COMMIT;

-- Tabela para agendamento/sincronização
CREATE TABLE IF NOT EXISTS mcp_sync_tasks (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    provider TEXT NOT NULL,
    target TEXT NOT NULL,
    last_synced TIMESTAMP DEFAULT now(),
    created_at TIMESTAMP DEFAULT now(),
    created_by uuid REFERENCES users(id),
    updated_at TIMESTAMP DEFAULT now(),
    updated_by uuid REFERENCES users(id),
    task_type TEXT NOT NULL CHECK (task_type IN ('pull', 'push', 'sync')),
    task_schedule job_type DEFAULT 'cron',
    task_expression TEXT DEFAULT '2 * * * *',
    task_command_type job_command_type DEFAULT 'api',
    task_method http_method DEFAULT 'POST',
    task_api_endpoint TEXT,
    task_payload JSONB,
    task_headers JSONB,
    task_retries INTEGER DEFAULT 0,
    task_timeout INTEGER DEFAULT 0,
    task_status job_status DEFAULT 'PENDING',
    task_next_run TIMESTAMP DEFAULT now() + INTERVAL '1 hour',
    task_last_run TIMESTAMP,
    task_last_run_status last_run_status DEFAULT 'pending',
    task_last_run_message last_run_message DEFAULT 'pending',
    task_command TEXT,
    task_activated BOOLEAN DEFAULT true,
    task_config JSONB NOT NULL DEFAULT '{}',
    task_tags TEXT[],
    task_priority INTEGER DEFAULT 0,
    task_notes TEXT,
    task_created_at TIMESTAMP DEFAULT now(),
    task_updated_at TIMESTAMP DEFAULT now(),
    task_created_by uuid REFERENCES users(id),
    task_updated_by uuid REFERENCES users(id),
    task_last_executed_by uuid REFERENCES users(id),
    task_last_executed_at TIMESTAMP,
    config JSONB NOT NULL DEFAULT '{}',
    active BOOLEAN DEFAULT true,
    UNIQUE(provider, target, task_type, task_schedule, task_expression, task_api_endpoint, task_method, task_payload, task_headers, task_retries, task_timeout, task_status, task_next_run, task_last_run_status, task_last_run_message)
);
-- COMMIT;

-- Tabela de jobs de sincronização
CREATE TABLE IF NOT EXISTS mcp_sync_jobs (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    task_id uuid NOT NULL REFERENCES mcp_sync_tasks(id) ON DELETE CASCADE,
    job_type TEXT NOT NULL CHECK (job_type IN ('pull', 'push', 'sync')),
    job_target TEXT NOT NULL,
    job_status job_status NOT NULL DEFAULT 'PENDING',
    last_run TIMESTAMP DEFAULT now(),
    last_run_status last_run_status DEFAULT 'pending',
    last_run_message last_run_message DEFAULT 'pending',
    next_run TIMESTAMP DEFAULT now() + INTERVAL '1 hour',
    retries INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    job_timeout INTEGER DEFAULT 0,
    job_command TEXT,
    job_method http_method DEFAULT 'POST',
    api_endpoint TEXT,
    payload JSONB,
    headers JSONB,
    tags TEXT[],
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_by uuid REFERENCES users(id),
    updated_by uuid REFERENCES users(id),
    last_executed_by uuid REFERENCES users(id),
    last_executed_at TIMESTAMP,
    UNIQUE(task_id, target, job_type, next_run, status),
    UNIQUE(task_id, target, job_type, last_run, status),
    UNIQUE(task_id, target, job_type, last_run_status, last_run_message, next_run, status)
);

-- Tabela de logs de sincronização
CREATE TABLE IF NOT EXISTS mcp_sync_logs (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    job_id uuid NOT NULL REFERENCES mcp_sync_jobs(id) ON DELETE CASCADE,
    status TEXT NOT NULL CHECK (status IN ('success', 'failure', 'pending')),
    message TEXT,
    started_at TIMESTAMP NOT NULL DEFAULT now(),
    ended_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    created_by uuid REFERENCES users(id),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_by uuid REFERENCES users(id),
    UNIQUE(job_id, started_at)
);
-- COMMIT;

-----------------------------------------------------------------------------------

COMMIT;
