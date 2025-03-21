ALTER TABLE feedback 
  ADD COLUMN created_at_str VARCHAR(50),
  ADD COLUMN answered_at_str VARCHAR(50);

UPDATE feedback 
SET 
  created_at_str = created_at::TEXT,
  answered_at_str = answered_at::TEXT;

ALTER TABLE feedback 
  DROP COLUMN created_at,
  DROP COLUMN answered_at;

ALTER TABLE feedback 
  RENAME COLUMN created_at_str TO created_at,
  RENAME COLUMN answered_at_str TO answered_at;

ALTER TABLE feedback 
  ALTER COLUMN created_at SET NOT NULL,
  ALTER COLUMN created_at SET DEFAULT to_char(NOW(), 'YYYY-MM-DD"T"HH24:MI:SS"Z"');
