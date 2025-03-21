-- When target_id = 0, this is a broadcast notification to all users with the given org_token
CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    header VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    type VARCHAR(20) NOT NULL DEFAULT 'news',
    target_id BIGINT NOT NULL, -- When 0, broadcast to all users with matching org_token
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