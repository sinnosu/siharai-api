-- 企業テーブルの作成
CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL,
    representative_name VARCHAR(255),
    phone_number VARCHAR(20),
    postal_code VARCHAR(10),
    address TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ユーザーテーブルの作成
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    company_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_company FOREIGN KEY (company_id) REFERENCES companies(id)
);

-- 取引先テーブルの作成
CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    company_id INT NOT NULL,
    client_name VARCHAR(255) NOT NULL,
    representative_name VARCHAR(255),
    phone_number VARCHAR(20),
    postal_code VARCHAR(10),
    address TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_client_company FOREIGN KEY (company_id) REFERENCES companies(id)
);

-- 取引先銀行口座テーブルの作成
CREATE TABLE client_bank_accounts (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    branch_name VARCHAR(255),
    account_number VARCHAR(50),
    account_holder_name VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_account_client FOREIGN KEY (client_id) REFERENCES clients(id)
);

-- 請求書データテーブルの作成
CREATE TABLE invoices (
    id SERIAL PRIMARY KEY,
    company_id INT NOT NULL,
    client_id INT NOT NULL,
    issue_date DATE NOT NULL,
    payment_amount DECIMAL(15, 2) NOT NULL,
    fee DECIMAL(15, 2),
    fee_rate DECIMAL(5, 2),
    tax DECIMAL(15, 2),
    tax_rate DECIMAL(5, 2),
    billed_amount DECIMAL(15, 2) NOT NULL,
    payment_due_date DATE NOT NULL,
    status VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_invoice_company FOREIGN KEY (company_id) REFERENCES companies(id),
    CONSTRAINT fk_invoice_client FOREIGN KEY (client_id) REFERENCES clients(id)
);
