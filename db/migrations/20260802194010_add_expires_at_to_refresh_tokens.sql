-- migrate:up
ALTER TABLE refresh_tokens
    ADD COLUMN expires_at TIMESTAMP NOT NULL AFTER refresh_token,
    ADD INDEX idx_refresh_tokens_user_id_expires_at (user_id, expires_at);

-- migrate:down
ALTER TABLE refresh_tokens
    DROP INDEX idx_refresh_tokens_user_id_expires_at,
    DROP COLUMN expires_at;
