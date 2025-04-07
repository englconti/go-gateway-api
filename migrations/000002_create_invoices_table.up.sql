CREATE TABLE IF NOT EXISTS invoices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(id),
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(255) NOT NULL DEFAULT 'pending',
    description VARCHAR(255) NOT NULL,
    payment_type VARCHAR(255) NOT NULL,
    card_last_digits VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE INDEX IF NOT EXISTS idx_invoices_account_id ON invoices (account_id);
CREATE INDEX IF NOT EXISTS idx_invoices_status ON invoices (status);
CREATE INDEX IF NOT EXISTS idx_invoices_created_at ON invoices (created_at);