-- Create "subscriptions" table
CREATE TABLE "subscriptions" (
  "id" bigserial NOT NULL,
  "service_name" character varying(50) NOT NULL,
  "price" bigint NOT NULL,
  "user_id" character varying(50) NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_subscriptions_user_service_date" to table: "subscriptions"
CREATE INDEX "idx_subscriptions_user_service_date" ON "subscriptions" ("user_id", "service_name", "start_date", "end_date");
