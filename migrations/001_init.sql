-- Create Table Accounts
CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(36) PRIMARY KEY,
    account_number VARCHAR(10) UNIQUE NOT NULL,
    balance DECIMAL(15,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL 
);

-- Create Table Transactions
CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(36) PRIMARY KEY,
    from_account_id VARCHAR(36) REFERENCES accounts(id),
    to_account_id VARCHAR(36) REFERENCES accounts(id),
    amount DECIMAL(15,2) NOT NULL,
    transaction_type VARCHAR(20) NOT NULL,
    qr_code TEXT,
    created_at TIMESTAMP NOT NULL
);

-- Create Indexes
CREATE INDEX idx_account_number ON account(account_number);
CREATE INDEX idx_transaction_from_acccount ON transaction(from_account_id);
CREATE INDEX idx_transaction_to_account ON transaction(to_account_id);