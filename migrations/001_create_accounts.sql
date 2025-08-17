CREATE TABLE IF NOT EXISTS accounts (
    id TEXT PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    type TEXT NOT NULL, -- Harta, Hutang, Modal, Pendapatan, Beban
    parent_id TEXT
);
