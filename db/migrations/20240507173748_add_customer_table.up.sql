CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX customers_name_trgm_idx ON customers USING gin (name gin_trgm_ops);
CREATE INDEX customers_phone_trgm_idx ON customers USING gin (phone gin_trgm_ops);