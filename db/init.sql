CREATE TABLE IF NOT EXISTS repository_tb
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    author     VARCHAR(255) NOT NULL,
    api_url    TEXT         NOT NULL,
    gh_url     TEXT         NOT NULL,
    clone_url  TEXT         NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (name, author)
);
