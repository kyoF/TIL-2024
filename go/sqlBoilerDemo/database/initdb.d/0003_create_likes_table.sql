CREATE TABLE likes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    tweet_id CHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, tweet_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (tweet_id) REFERENCES tweets(tweet_id)
);

