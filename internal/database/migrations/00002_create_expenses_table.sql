-- +goose Up
CREATE TABLE "expenses" (
  "id" integer primary key autoincrement not null,
  "user_id" integer not null,
  "category_id" integer not null,
  "amount" integer not null,
  "description" text,
  "spent_at" datetime not null,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  foreign key ("user_id") references "users" ("id") on delete cascade,
  foreign key ("category_id") references "categories" ("id") on delete cascade
);

-- +goose Down
DROP TABLE "expenses";
