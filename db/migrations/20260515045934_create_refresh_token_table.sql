-- migrate:up
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id int AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id_refresh_tokens FOREIGN KEY (user_id) REFERENCES users(id)
);

-- migrate:down
DROP TABLE IF EXISTS "refresh_tokens";-- migrate:up


