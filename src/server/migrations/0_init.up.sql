CREATE TABLE notifications (
                        id SERIAL PRIMARY KEY,
                        message TEXT NOT NULL,
                        target_id BIGINT NOT NULL,
                        org_token TEXT NOT NULL,
                        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);