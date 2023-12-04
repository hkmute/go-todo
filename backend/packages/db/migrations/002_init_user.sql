-- Write your migrate up statements here

CREATE TABLE app_user (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE todo ADD COLUMN user_id INTEGER NOT NULL REFERENCES app_user(id);

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

ALTER TABLE todo DROP COLUMN user_id;
DROP TABLE app_user;
