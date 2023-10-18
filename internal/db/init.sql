CREATE DATABASE IF NOT EXISTS DB_NAME;
USE DB_NAME;

-- DROP TABLES
-- DROP TABLE IF EXISTS subscriptions;
-- DROP TABLE IF EXISTS messages;
-- DROP TABLE IF EXISTS rooms;
-- DROP TABLE IF EXISTS users;

-- CREATE TABLES
CREATE TABLE IF NOT EXISTS rooms (
  id UUID PRIMARY KEY,
  name varchar(45) NOT NULL,
  updated_at timestamptz NOT NULL,
  created_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  name varchar(45) NOT NULL,
  email varchar(45) NOT NULL,
  password varchar(45) NOT NULL,
  updated_at timestamptz NOT NULL,
  created_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS subscriptions (
  id UUID PRIMARY KEY,
  room_id UUID NOT NULL REFERENCES rooms (id),
  user_id UUID NOT NULL REFERENCES users (id),
  updated_at timestamptz NOT NULL,
  created_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS messages (
  id UUID PRIMARY KEY,
  msg TEXT NOT NULL,
  room_id UUID NOT NULL REFERENCES rooms (id),
  user_id UUID NOT NULL REFERENCES users (id),
  updated_at timestamptz NOT NULL,
  created_at timestamptz NOT NULL
);

-- INDEXING
CREATE INDEX ON rooms (updated_at);
CREATE INDEX ON subscriptions (room_id, user_id);
CREATE INDEX ON subscriptions (user_id, updated_at);
CREATE INDEX ON messages (room_id, user_id);
CREATE INDEX ON messages (created_at);
