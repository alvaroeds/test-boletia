CREATE TABLE IF NOT EXISTS request (
    id                  SERIAL          NOT NULL ,
    method              VARCHAR(50)     NOT NULL,
    path             VARCHAR(50)     NOT NULL,
    code                INTEGER         NOT NULL,
    "time"              NUMERIC(13, 8)  DEFAULT 0 NOT NULL,
    created_at          TIMESTAMP       DEFAULT NOW(),
    updated_at          TIMESTAMP,
    CONSTRAINT request_id_pk PRIMARY KEY (id)
);