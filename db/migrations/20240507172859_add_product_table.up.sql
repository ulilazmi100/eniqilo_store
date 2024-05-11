CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(30) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    notes VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    stock INT NOT NULL,
    location VARCHAR(255) NOT NULL,
    is_avail BOOLEAN NOT NULL, DEFAULT 1
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX products_name_trgm_idx ON products USING gin (name gin_trgm_ops);
CREATE INDEX products_sku_idx ON products (sku);
CREATE INDEX products_price_idx ON products (price);
CREATE INDEX products_created_at_idx ON products (created_at);