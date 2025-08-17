CREATE TABLE IF NOT EXISTS receivables (
    id TEXT PRIMARY KEY,
    customer_id TEXT NOT NULL REFERENCES customers(id),
    amount NUMERIC(18,2) NOT NULL,
    received NUMERIC(18,2) DEFAULT 0,
    due_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
