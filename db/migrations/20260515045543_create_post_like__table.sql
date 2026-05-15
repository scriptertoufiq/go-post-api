-- migrate:up
CREATE TABLE IF NOT EXISTS post_likes (
    id int AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    post_id int NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id_post_likes FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_post_id_post_likes FOREIGN KEY (post_id) REFERENCES posts(id)
);

-- migrate:down
DROP TABLE IF EXISTS "post_likes";