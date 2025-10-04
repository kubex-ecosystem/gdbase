DO $$ BEGIN IF EXISTS (SELECT 1 FROM pg_database WHERE datname = 'gdbase_db') THEN EXECUTE 'GRANT CONNECT ON DATABASE gdbase_db TO readonly, readwrite, admin'; END IF; END$$;

ALTER SCHEMA public OWNER TO admin;
REVOKE CREATE ON SCHEMA public FROM PUBLIC;
GRANT USAGE ON SCHEMA public TO readonly, readwrite, admin;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO readonly, readwrite, admin;
GRANT SELECT ON ALL SEQUENCES IN SCHEMA public TO readonly;
GRANT UPDATE ON ALL SEQUENCES IN SCHEMA public TO readwrite, admin;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO readwrite, admin;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly;

DO $$
DECLARE r record;

BEGIN
FOR r IN (SELECT n.nspname AS nsp FROM pg_namespace n WHERE n.nspname IN ('public')) LOOP

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT SELECT ON TABLES TO readonly',
    r.nsp
);

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT SELECT,INSERT,UPDATE,DELETE ON TABLES TO readwrite',
    r.nsp
);

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT ALL ON TABLES TO admin',
    r.nsp
);

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT USAGE ON SEQUENCES TO readonly',
    r.nsp
);

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT USAGE,UPDATE ON SEQUENCES TO readwrite',
    r.nsp
);

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT ALL ON SEQUENCES TO admin',
    r.nsp
);

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT EXECUTE ON FUNCTIONS TO readwrite, readonly',
    r.nsp
);

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT ALL ON FUNCTIONS TO admin',
    r.nsp
);

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT USAGE ON TYPES TO readonly, readwrite',
    r.nsp
);

EXECUTE format(
    'ALTER DEFAULT PRIVILEGES IN SCHEMA %I GRANT ALL ON TYPES TO admin',
    r.nsp
);

END LOOP;

END$$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'uuid-ossp') THEN
        BEGIN
            EXECUTE format('ALTER EXTENSION %I OWNER TO admin', 'uuid-ossp');
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'skipping ALTER EXTENSION %: %', 'uuid-ossp', SQLERRM;
        END;
    END IF;
END$$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'pgcrypto') THEN
        BEGIN
            EXECUTE format('ALTER EXTENSION %I OWNER TO admin', 'pgcrypto');

EXCEPTION WHEN OTHERS THEN RAISE NOTICE 'skipping ALTER EXTENSION %: %',
'pgcrypto',
SQLERRM;

END;

END IF;

END$$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'pg_trgm') THEN
        BEGIN
            EXECUTE format('ALTER EXTENSION %I OWNER TO admin', 'pg_trgm');
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'skipping ALTER EXTENSION %: %', 'pg_trgm', SQLERRM;
        END;
    END IF;
END$$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'btree_gist') THEN
        BEGIN
            EXECUTE format('ALTER EXTENSION %I OWNER TO admin', 'btree_gist');

EXCEPTION WHEN OTHERS THEN RAISE NOTICE 'skipping ALTER EXTENSION %: %',
'btree_gist',
SQLERRM;

END;

END IF;

END$$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'fuzzystrmatch') THEN
        BEGIN
            EXECUTE format('ALTER EXTENSION %I OWNER TO admin', 'fuzzystrmatch');
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'skipping ALTER EXTENSION %: %', 'fuzzystrmatch', SQLERRM;
        END;
    END IF;
END$$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'hstore') THEN
        BEGIN
            EXECUTE format('ALTER EXTENSION %I OWNER TO admin', 'hstore');

EXCEPTION WHEN OTHERS THEN RAISE NOTICE 'skipping ALTER EXTENSION %: %',
'hstore',
SQLERRM;

END;

END IF;

END$$;

CREATE EXTENSION IF NOT EXISTS citext;

ALTER TABLE users ALTER COLUMN email TYPE citext;

ALTER TABLE users ALTER COLUMN username TYPE citext;

ALTER ROLE user_readonly PASSWORD NULL;

ALTER ROLE user_readwrite PASSWORD NULL;
