CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "role_id" bigint NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY,
  "role_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "title" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "content" text NOT NULL,
  "category_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "category_name" varchar NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("role_id");

CREATE INDEX ON "posts" ("username");

CREATE INDEX ON "posts" ("category_id");

CREATE INDEX ON "posts" ("username", "category_id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "posts" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

INSERT INTO "roles" (id, role_name) VALUES (1, 'SUPERUSER'), (2, 'USER');