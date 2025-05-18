CREATE TABLE IF NOT EXISTS pi_google_integration (
    gi_id VARCHAR(36) PRIMARY KEY,
    gi_user_email VARCHAR(255) NOT NULL,
    gi_google_access_token VARCHAR(255) NOT NULL,
    gi_google_token_type VARCHAR(255) NOT NULL,
    gi_google_refresh_token VARCHAR(255) NOT NULL,
    us_created_at VARCHAR(60) NOT NULL,
    us_updated_at VARCHAR(60) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_gi_user_email ON pi_google_integration(gi_user_email);