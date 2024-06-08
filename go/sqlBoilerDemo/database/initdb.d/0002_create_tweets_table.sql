CREATE TABLE tweets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tweet_id CHAR(36) NOT NULL UNIQUE,
    user_id VARCHAR(255) NOT NULL,
    text TEXT NOT NULL,
    likes_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (tweet_id, user_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

