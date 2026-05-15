-- migrate:up
CREATE TABLE IF NOT EXISTS comment_likes (
    id int AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    comment_id int NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id_comment_likes FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_comment_id_comment_likes FOREIGN KEY (comment_id) REFERENCES comments(id)
);

-- migrate:down
DROP TABLE IF EXISTS "comment_likes";