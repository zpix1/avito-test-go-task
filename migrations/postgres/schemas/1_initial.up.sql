CREATE DOMAIN SLUG_NAME as VARCHAR(255) NOT NULL CONSTRAINT non_empty CHECK(length(value)>0);

CREATE TABLE slugs
(
    id   BIGSERIAL PRIMARY KEY,
    name SLUG_NAME UNIQUE
);

CREATE TABLE slugs_users
(
    slug_name SLUG_NAME REFERENCES slugs(name) ON DELETE CASCADE,
    user_id INTEGER,
--     NULL means no timelimit present
    valid_until TIMESTAMP,
    CONSTRAINT slugs_users_pk PRIMARY KEY (slug_name, user_id)
);

CREATE TABLE slugs_history
(
    user_id INTEGER,
    slug_name SLUG_NAME,
    removed boolean,
    created_at TIMESTAMP DEFAULT current_timestamp
)