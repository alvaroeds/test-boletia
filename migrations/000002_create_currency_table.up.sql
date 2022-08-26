CREATE TABLE IF NOT EXISTS currency (
    id                  SERIAL    NOT NULL,
    code                VARCHAR(5)      NOT NULL,
    value               NUMERIC(13, 8)  DEFAULT 0 NOT NULL,
    created_at          TIMESTAMP       DEFAULT NOW(),
    updated_at          TIMESTAMP,
    CONSTRAINT currency_id_pk PRIMARY KEY (id)
);