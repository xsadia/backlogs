CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users ( 
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  bio VARCHAR(160),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),
  deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS backlogs (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  user_id uuid CONSTRAINT backlog_user_fk REFERENCES users (id),
  description VARCHAR(160),
  genre VARCHAR(10) NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),
  deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS backlog_items (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  backlog_id uuid CONSTRAINT item_backlog_fk REFERENCES backlogs (id),
  avg_time_completion INTEGER,
  review TEXT
);
