CREATE TYPE "protocol" AS ENUM (
  'local',
  'activitypub'
);

CREATE TABLE "users" (
  "id" varchar(255) PRIMARY KEY,
  "username" varchar(255) NOT NULL,
  "host" varchar(255) NOT NULL,
  "protocol" protocol,
  "hashed_password" varchar(255),
  "display_name" varchar(255),
  "profile" text,
  "icon" varchar(255),
  "created_at" timestamp DEFAULT 'CURRENT_TIMESTAMP',
  "updated_at" timestamp DEFAULT 'CURRENT_TIMESTAMP'
);

CREATE TABLE "ap_user_identifiers" (
  "user_id" vatchar(255) PRIMARY KEY,
  "public_key" text NOT NULL,
  "private_key" text NOT NULL,
  "created_at" timestamp DEFAULT 'CURRENT_TIMESTAMP',
  "updated_at" timestamp DEFAULT 'CURRENT_TIMESTAMP'
);

CREATE TABLE "assets" (
  "id" varchar(255) PRIMARY KEY,
  "url" varchar(255) NOT NULL,
  "created_at" timestamp DEFAULT 'CURRENT_TIMESTAMP',
  "updated_at" timestamp DEFAULT 'CURRENT_TIMESTAMP'
);

CREATE TABLE "follows" (
  "id" varchar(255) PRIMARY KEY,
  "follower_id" varchar(255) NOT NULL,
  "followee_id" varchar(255) NOT NULL,
  "created_at" timestamp DEFAULT 'CURRENT_TIMESTAMP',
  "updated_at" timestamp DEFAULT 'CURRENT_TIMESTAMP'
);

ALTER TABLE "users" ADD FOREIGN KEY ("icon") REFERENCES "assets" ("id");

ALTER TABLE "ap_user_identifiers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("followee_id") REFERENCES "users" ("id");
