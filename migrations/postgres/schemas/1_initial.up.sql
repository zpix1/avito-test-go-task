CREATE DOMAIN SLUG_NAME as VARCHAR(255) NOT NULL CONSTRAINT non_empty CHECK (length(value) > 0);
-- uint32
CREATE DOMAIN AUTO_ADD_WEIGHT as BIGINT CONSTRAINT non_negative CHECK (value >= 0 and value < '4294967296'::BIGINT);

CREATE TABLE slugs
(
    id              BIGSERIAL PRIMARY KEY,
    name            SLUG_NAME UNIQUE,
    auto_add_weight AUTO_ADD_WEIGHT
);

CREATE TABLE slugs_users
(
    slug_name   SLUG_NAME REFERENCES slugs (name) ON DELETE CASCADE,
    user_id     INTEGER,
--     NULL means no timelimit present
    valid_until TIMESTAMP,
    CONSTRAINT slugs_users_pk PRIMARY KEY (slug_name, user_id)
);

CREATE TABLE slugs_history
(
    user_id    INTEGER,
    slug_name  SLUG_NAME,
    removed    boolean,
    created_at TIMESTAMP DEFAULT current_timestamp
)