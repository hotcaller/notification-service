CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    header VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    type VARCHAR(20) NOT NULL DEFAULT 'news',
    target_id BIGINT NOT NULL,
    org_token TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS feedback (
    id SERIAL PRIMARY KEY,
    header VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    answer TEXT,
    user_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    answered_at TIMESTAMP
);