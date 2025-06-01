CREATE TABLE IF NOT EXISTS pi_session (
    se_id VARCHAR(36) PRIMARY KEY,
    se_email VARCHAR(255) NOT NULL,
    se_created_at VARCHAR(60) NOT NULL,
    se_updated_at VARCHAR(60) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_se_email ON pi_session(se_email);