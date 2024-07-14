CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password1 VARCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    eco_points INTEGER DEFAULT 100,
    bio VARCHAR DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);


CREATE TABLE refresh_tokens (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid REFERENCES users(id) not null,
    token text UNIQUE not null
);

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
