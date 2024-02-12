-- Write your migrate up statements here

ALTER TABLE todo ADD COLUMN item_order INTEGER UNSIGNED;

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

ALTER TABLE todo DROP COLUMN item_order;