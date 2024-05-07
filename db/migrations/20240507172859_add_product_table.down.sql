DROP TABLE IF EXISTS products;

DROP INDEX IF EXISTS products_name_trgm_idx;
DROP INDEX IF EXISTS products_sku_idx;
DROP INDEX IF EXISTS products_price_idx;
DROP INDEX IF EXISTS products_created_at_idx;