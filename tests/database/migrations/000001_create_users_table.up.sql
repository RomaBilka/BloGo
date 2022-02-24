BEGIN;

CREATE TABLE IF NOT EXISTS public.users
(
    id serial NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    phone text NOT NULL,
    PRIMARY KEY (id)
    );

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;

COMMIT;