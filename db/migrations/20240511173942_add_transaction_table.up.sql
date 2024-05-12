CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    customer_id VARCHAR(255),
    paid INT NOT NULL,
    change INT NOT NULL,
    product_details JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX transactions_customer_id_idx ON transactions (customer_id);
CREATE INDEX transactions_created_at_idx ON transactions (created_at);
