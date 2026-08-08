-- +goose Up
CREATE TABLE "users" (
  "id" integer primary key autoincrement not null,
  "role" integer not null,
  "name" varchar not null,
  "email" varchar not null,
  "email_verified_at" datetime,
  "password" varchar not null,
  "remember_token" varchar,
  "created_at" datetime,
  "updated_at" datetime
);

-- +goose Down
DROP TABLE "users";
