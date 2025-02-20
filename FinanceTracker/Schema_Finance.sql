CREATE TABLE "users" (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR(100),
  "email" VARCHAR(255) UNIQUE,
  "password_hash" TEXT,
  "currency" VARCHAR(10),
  "created_at" TIMESTAMP DEFAULT 'now()'
);

CREATE TABLE "income" (
  "id" UUID PRIMARY KEY,
  "user_id" UUID,
  "amount" DECIMAL(10,2),
  "category" VARCHAR(50),
  "source" VARCHAR(100),
  "date" DATE,
  "created_at" TIMESTAMP DEFAULT 'now()'
);

CREATE TABLE "expenses" (
  "id" UUID PRIMARY KEY,
  "user_id" UUID,
  "amount" DECIMAL(10,2),
  "category" VARCHAR(50),
  "description" TEXT,
  "date" DATE,
  "created_at" TIMESTAMP DEFAULT 'now()'
);

CREATE TABLE "budgets" (
  "id" UUID PRIMARY KEY,
  "user_id" UUID,
  "category" VARCHAR(50),
  "amount" DECIMAL(10,2),
  "start_date" DATE,
  "end_date" DATE,
  "created_at" TIMESTAMP DEFAULT 'now()'
);

CREATE TABLE "goals" (
  "id" UUID PRIMARY KEY,
  "user_id" UUID,
  "name" VARCHAR(100),
  "target_amount" DECIMAL(10,2),
  "current_amount" DECIMAL(10,2) DEFAULT 0,
  "deadline" DATE,
  "created_at" TIMESTAMP DEFAULT 'now()'
);

CREATE TABLE "transactions" (
  "id" UUID PRIMARY KEY,
  "user_id" UUID,
  "type" VARCHAR(10),
  "amount" DECIMAL(10,2),
  "category" VARCHAR(50),
  "description" TEXT,
  "date" DATE,
  "created_at" TIMESTAMP DEFAULT 'now()'
);

CREATE TABLE "recurring_payments" (
  "id" UUID PRIMARY KEY,
  "user_id" UUID,
  "name" VARCHAR(100),
  "amount" DECIMAL(10,2),
  "category" VARCHAR(50),
  "frequency" VARCHAR(10),
  "next_payment_date" DATE,
  "created_at" TIMESTAMP DEFAULT 'now()'
);

CREATE TABLE "alerts" (
  "id" UUID PRIMARY KEY,
  "user_id" UUID,
  "type" VARCHAR(10),
  "message" TEXT,
  "triggered_at" TIMESTAMP DEFAULT 'now()'
);

CREATE TABLE "reports" (
  "id" UUID PRIMARY KEY,
  "user_id" UUID,
  "type" VARCHAR(10),
  "report_data" JSON,
  "created_at" TIMESTAMP DEFAULT 'now()'
);

ALTER TABLE "income" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "budgets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "goals" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "recurring_payments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "alerts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reports" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
