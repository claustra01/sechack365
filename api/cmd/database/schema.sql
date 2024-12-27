CREATE TYPE "protocol" AS ENUM (
  'local',
  'activitypub',
  'nostr'
);

CREATE TABLE "users" (
  "id" varchar(255) PRIMARY KEY,
  "username" varchar(255) NOT NULL DEFAULT '',
  "protocol" protocol NOT NULL,
  "hashed_password" varchar(255) NOT NULL DEFAULT '',
  "display_name" varchar(255) NOT NULL DEFAULT '',
  "profile" text,
  "icon" varchar(255),
  "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "ap_user_identifiers" (
  "user_id" varchar(255) PRIMARY KEY,
  "local_username" varchar(255) NOT NULL,
  "host" varchar(255) NOT NULL,
  "public_key" text NOT NULL DEFAULT '',
  "private_key" text,
  "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "nostr_user_identifiers" (
  "user_id" varchar(255) PRIMARY KEY,
  "public_key" text NOT NULL,
  "private_key" text,
  "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "follows" (
  "id" varchar(255) UNIQUE,
  "follower_id" varchar(255),
  "target_id" varchar(255),
  "is_accepted" boolean NOT NULL DEFAULT false,
  "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("follower_id", "target_id")
);

CREATE TABLE "posts" (
  "id" varchar(255) PRIMARY KEY,
  "user_id" varchar(255),
  "content" text NOT NULL,
  "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "reactions" (
  "id" varchar(255) PRIMARY KEY,
  "type" varchar(255) NOT NULL,
  "user_id" varchar(255) NOT NULL,
  "post_id" varchar(255) NOT NULL,
  "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "nostr_relays" (
  "id" varchar(255) PRIMARY KEY,
  "url" varchar(255) UNIQUE NOT NULL,
  "is_enable" boolean NOT NULL DEFAULT true,
  "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "ap_user_identifiers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "nostr_user_identifiers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("target_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reactions" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

