CREATE TABLE IF NOT EXISTS pi_user (
    us_id VARCHAR(36) PRIMARY KEY,
    us_email VARCHAR(255) NOT NULL,
    us_password VARCHAR(255) NOT NULL,
    us_created_at VARCHAR(60) NOT NULL,
    us_updated_at VARCHAR(60) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_us_email ON pi_user(us_email);