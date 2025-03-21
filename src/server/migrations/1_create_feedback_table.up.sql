CREATE TABLE IF NOT EXISTS feedback (
    id SERIAL PRIMARY KEY,
    header VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    answer TEXT,
    user_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    answered_at TIMESTAMP
);