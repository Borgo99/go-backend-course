CREATE TABLE accounts (
  id bigserial PRIMARY KEY,
  owner varchar NOT NULL,
  balance bigint NOT NULL DEFAULT 0,
  currency varchar NOT NULL  DEFAULT 'EUR',
  created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE entries (
  id bigserial PRIMARY KEY,
  account bigserial NOT NULL REFERENCES accounts (id),
  amount bigint NOT NULL,
  created_at timestamptz DEFAULT (now())
);

CREATE TABLE transfers (
  id bigserial PRIMARY KEY,
  sender bigserial NOT NULL REFERENCES accounts (id),
  receiver bigserial NOT NULL REFERENCES accounts (id),
  amount bigint NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now()),
  CHECK (amount > 0),
  CHECK (sender <> receiver)
);

CREATE INDEX ON accounts (owner);
CREATE INDEX ON entries (account, created_at);
CREATE INDEX ON transfers (sender);
CREATE INDEX ON transfers (receiver);