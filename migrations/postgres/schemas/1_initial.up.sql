CREATE TABLE slugs
(
    id   BIGSERIAL PRIMARY KEY ,
    name VARCHAR(255) UNIQUE NOT NULL CONSTRAINT non_empty CHECK(length(name)>0)
);

CREATE TABLE slugs_users
(
    slug_name VARCHAR(255) REFERENCES slugs(name),
    user_id INTEGER,
    CONSTRAINT slugs_users_pk PRIMARY KEY (slug_name, user_id)
);