ALTER TABLE feedback 
  ADD COLUMN created_at_text TEXT,
  ADD COLUMN answered_at_text TEXT;

UPDATE feedback 
SET 
  created_at_text = created_at::TEXT,
  answered_at_text = answered_at::TEXT;

ALTER TABLE feedback 
  DROP COLUMN created_at,
  DROP COLUMN answered_at;

ALTER TABLE feedback 
  RENAME COLUMN created_at_text TO created_at,
  RENAME COLUMN answered_at_text TO answered_at;

ALTER TABLE feedback 
  ALTER COLUMN created_at SET NOT NULL;

ALTER TABLE feedback 
  ALTER COLUMN created_at SET DEFAULT to_char(NOW(), 'YYYY-MM-DD"T"HH24:MI:SS"Z"');