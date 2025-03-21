ALTER TABLE notifications 
  ADD COLUMN created_at_text TEXT;

UPDATE notifications 
SET created_at_text = created_at::TEXT;

ALTER TABLE notifications 
  DROP COLUMN created_at;

ALTER TABLE notifications 
  RENAME COLUMN created_at_text TO created_at;