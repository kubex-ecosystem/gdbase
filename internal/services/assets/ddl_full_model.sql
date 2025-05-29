-- Tabela de endereços (abstrata, reutilizável)
CREATE TABLE address (
    id uuid PRIMARY KEY,
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

-- Tabela de parceiros
CREATE TABLE partner (
    id uuid PRIMARY KEY,
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
    address_ids uuid[] NOT NULL, -- array de ids de address
    credit_limit numeric(18,2),
    current_debt numeric(18,2),
    payment_terms text[],
    last_purchase_date timestamp without time zone,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    is_active boolean NOT NULL DEFAULT true
);

-- Tabela de contatos do parceiro
CREATE TABLE partner_contact (
    id uuid PRIMARY KEY,
    partner_id uuid NOT NULL REFERENCES partner(id) ON DELETE CASCADE,
    name varchar(100) NOT NULL,
    email varchar(100),
    phone varchar(30),
    position varchar(50),
    is_primary boolean NOT NULL DEFAULT false
);

-- Tabela de histórico de vendas do parceiro
CREATE TABLE partner_sales_history (
    id uuid PRIMARY KEY,
    partner_id uuid NOT NULL REFERENCES partner(id) ON DELETE CASCADE,
    year integer NOT NULL,
    q1 integer NOT NULL DEFAULT 0,
    q2 integer NOT NULL DEFAULT 0,
    q3 integer NOT NULL DEFAULT 0,
    q4 integer NOT NULL DEFAULT 0
);

-- Tabela de produtos
CREATE TABLE product (
    id uuid PRIMARY KEY,
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
    shelf_life_days integer
);

-- Tabela de categorias de produto
CREATE TABLE product_category (
    id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    parent_id uuid REFERENCES product_category(id),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

-- Tabela de armazéns
CREATE TABLE warehouse (
    id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    location varchar(255),
    capacity integer,
    current_stock integer,
    manager varchar(100),
    contact varchar(100),
    address_id uuid REFERENCES address(id),
    external_id varchar(255),
    external_code varchar(255),
    notes text,
    tags text[],
    status varchar(50),
    created_by uuid,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    last_sync_at timestamp without time zone,
    is_active boolean NOT NULL DEFAULT true
);

-- Tabela de estoque
CREATE TABLE inventory (
    id uuid PRIMARY KEY,
    product_id uuid NOT NULL REFERENCES product(id),
    warehouse_id uuid NOT NULL REFERENCES warehouse(id),
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

-- Tabela de tabela de preços
CREATE TABLE price_table (
    id uuid PRIMARY KEY,
    price_table_id uuid NOT NULL,
    external_id varchar(255),
    external_code varchar(255),
    name varchar(255),
    description text,
    product_id uuid NOT NULL REFERENCES product(id),
    price numeric(18,2) NOT NULL,
    discount numeric(18,2),
    cost numeric(18,2),
    start_date timestamp without time zone,
    end_date timestamp without time zone,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    last_sync_at timestamp without time zone,
    is_active boolean NOT NULL DEFAULT true
);

-- Tabela de referências externas de produto
CREATE TABLE external_reference (
    id uuid PRIMARY KEY,
    product_id uuid NOT NULL REFERENCES product(id),
    system_name varchar(100) NOT NULL,
    external_id varchar(255) NOT NULL,
    external_code varchar(100),
    notes text,
    last_sync_date timestamp without time zone,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

-- Tabela de histórico de vendas de produto
CREATE TABLE sales_history (
    id uuid PRIMARY KEY,
    product_id uuid NOT NULL REFERENCES product(id),
    period varchar(10) NOT NULL CHECK (period IN ('daily','weekly','monthly','quarterly','yearly')),
    date date NOT NULL,
    quantity integer NOT NULL,
    revenue numeric(18,2),
    created_at timestamp without time zone NOT NULL DEFAULT now()
);

-- Tabela de pedidos
CREATE TABLE "order" (
    id uuid PRIMARY KEY,
    external_id varchar(255),
    order_number varchar(100),
    partner_id uuid NOT NULL REFERENCES partner(id),
    status varchar(30) NOT NULL,
    order_date timestamp without time zone NOT NULL,
    estimated_delivery_date timestamp without time zone,
    actual_delivery_date timestamp without time zone,
    shipping_address_id uuid REFERENCES address(id),
    payment_method varchar(30),
    payment_status varchar(20),
    notes text,
    total_amount numeric(18,2) NOT NULL,
    discount_amount numeric(18,2) NOT NULL DEFAULT 0,
    tax_amount numeric(18,2),
    shipping_amount numeric(18,2),
    final_amount numeric(18,2) NOT NULL,
    is_automatically_generated boolean DEFAULT false,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    last_sync_at timestamp without time zone,
    prediction_id uuid,
    priority integer,
    expected_margin numeric(18,2)
);

-- Tabela de itens do pedido
CREATE TABLE order_item (
    id uuid PRIMARY KEY,
    order_id uuid NOT NULL REFERENCES "order"(id) ON DELETE CASCADE,
    product_id uuid NOT NULL REFERENCES product(id),
    quantity numeric(18,3) NOT NULL,
    unit_price numeric(18,2) NOT NULL,
    discount numeric(18,2) NOT NULL DEFAULT 0,
    vol_type varchar(50),
    lot_control varchar(100),
    total numeric(18,2) NOT NULL,
    is_suggested boolean DEFAULT false,
    suggestion_reason text,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

-- Tabela de pagamentos do pedido
CREATE TABLE order_payment (
    id uuid PRIMARY KEY,
    order_id uuid NOT NULL REFERENCES "order"(id) ON DELETE CASCADE,
    method varchar(30) NOT NULL,
    installments integer NOT NULL DEFAULT 1,
    due_date timestamp without time zone NOT NULL,
    value numeric(18,2) NOT NULL,
    status varchar(20) NOT NULL
);

-- Tabela de endereços do pedido (endereços customizados para cada pedido)
CREATE TABLE order_address (
    id uuid PRIMARY KEY,
    order_id uuid NOT NULL REFERENCES "order"(id) ON DELETE CASCADE,
    street varchar(255) NOT NULL,
    number varchar(20) NOT NULL,
    complement varchar(100),
    district varchar(100),
    city varchar(100) NOT NULL,
    state varchar(50) NOT NULL,
    country varchar(50) NOT NULL,
    zip_code varchar(20) NOT NULL,
    type varchar(20) NOT NULL CHECK (type IN ('billing','shipping')),
    is_default boolean NOT NULL DEFAULT false
);

-- Índices essenciais
CREATE INDEX idx_partner_code ON partner(code);
CREATE INDEX idx_partner_status ON partner(status);
CREATE INDEX idx_partner_category ON partner(category);
CREATE INDEX idx_partner_created_at ON partner(created_at);
CREATE INDEX idx_partner_address_ids ON partner USING GIN(address_ids);
CREATE INDEX idx_partner_contact_partner_id ON partner_contact(partner_id);
CREATE INDEX idx_partner_sales_history_partner_id ON partner_sales_history(partner_id);
CREATE INDEX idx_product_category_id ON product(category_id);
CREATE INDEX idx_product_is_active ON product(is_active);
CREATE INDEX idx_product_created_at ON product(created_at);
CREATE INDEX idx_inventory_product_id ON inventory(product_id);
CREATE INDEX idx_inventory_warehouse_id ON inventory(warehouse_id);
CREATE INDEX idx_warehouse_address_id ON warehouse(address_id);
CREATE INDEX idx_order_partner_id ON "order"(partner_id);
CREATE INDEX idx_order_status ON "order"(status);
CREATE INDEX idx_order_order_date ON "order"(order_date);
CREATE INDEX idx_order_payment_status ON "order"(payment_status);
CREATE INDEX idx_order_item_order_id ON order_item(order_id);
CREATE INDEX idx_order_item_product_id ON order_item(product_id);
CREATE INDEX idx_order_payment_order_id ON order_payment(order_id);
CREATE INDEX idx_order_address_order_id ON order_address(order_id);
