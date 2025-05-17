BEGIN;

CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- для gen_random_uuid()

CREATE TABLE tariffs
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT NOT NULL,
    description TEXT
);


COMMIT;