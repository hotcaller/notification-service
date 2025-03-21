ALTER TABLE subscriptions 
  ADD COLUMN token_text TEXT;

UPDATE subscriptions 
SET token_text = token::TEXT;

ALTER TABLE subscriptions 
  DROP COLUMN token;

ALTER TABLE subscriptions 
  RENAME COLUMN token_text TO token;

ALTER TABLE subscriptions 
  ALTER COLUMN token SET NOT NULL;