CREATE TABLE "sharedlinks" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "urlhash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
