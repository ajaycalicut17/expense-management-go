-- +goose Up
CREATE TABLE "categories" (
  "id" integer primary key autoincrement not null,
  "name" varchar not null,
  "created_at" datetime,
  "updated_at" datetime
);

-- +goose Down
DROP TABLE "categories";
