-- +goose Up
-- Add organizations table for B2B multi-tenancy
CREATE TABLE "organizations" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(255) NOT NULL,
  "domain" VARCHAR(255) UNIQUE,
  "api_key_prefix" VARCHAR(10) UNIQUE,
  "status" VARCHAR(20) NOT NULL DEFAULT 'active',
  "plan" VARCHAR(50) NOT NULL DEFAULT 'basic',
  "max_users" INT NOT NULL DEFAULT 10,
  "max_api_requests_per_month" BIGINT NOT NULL DEFAULT 10000,
  "settings" JSONB,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- Add organization_id to existing tables
ALTER TABLE "users" ADD COLUMN "organization_id" UUID REFERENCES "organizations"("id");
ALTER TABLE "wallets" ADD COLUMN "organization_id" UUID REFERENCES "organizations"("id");
ALTER TABLE "transactions" ADD COLUMN "organization_id" UUID REFERENCES "organizations"("id");

-- Create API keys table
CREATE TABLE "api_keys" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "organization_id" UUID NOT NULL REFERENCES "organizations"("id"),
  "name" VARCHAR(255) NOT NULL,
  "key_hash" VARCHAR(255) NOT NULL UNIQUE,
  "permissions" JSONB NOT NULL DEFAULT '[]',
  "status" VARCHAR(20) NOT NULL DEFAULT 'active',
  "last_used_at" TIMESTAMP,
  "expires_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- Create organization members table
CREATE TABLE "organization_members" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "organization_id" UUID NOT NULL REFERENCES "organizations"("id"),
  "user_id" UUID NOT NULL REFERENCES "users"("id"),
  "role" VARCHAR(50) NOT NULL DEFAULT 'member',
  "permissions" JSONB NOT NULL DEFAULT '[]',
  "status" VARCHAR(20) NOT NULL DEFAULT 'active',
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  UNIQUE("organization_id", "user_id")
);

-- Create usage tracking table
CREATE TABLE "api_usage" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "organization_id" UUID NOT NULL REFERENCES "organizations"("id"),
  "api_key_id" UUID REFERENCES "api_keys"("id"),
  "endpoint" VARCHAR(255) NOT NULL,
  "method" VARCHAR(10) NOT NULL,
  "status_code" INT NOT NULL,
  "response_time_ms" INT,
  "created_at" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- Create indexes for performance
CREATE INDEX idx_users_organization_id ON "users"("organization_id");
CREATE INDEX idx_wallets_organization_id ON "wallets"("organization_id");
CREATE INDEX idx_transactions_organization_id ON "transactions"("organization_id");
CREATE INDEX idx_api_keys_organization_id ON "api_keys"("organization_id");
CREATE INDEX idx_api_usage_organization_id ON "api_usage"("organization_id");
CREATE INDEX idx_api_usage_created_at ON "api_usage"("created_at");

-- +goose Down
DROP INDEX IF EXISTS idx_api_usage_created_at;
DROP INDEX IF EXISTS idx_api_usage_organization_id;
DROP INDEX IF EXISTS idx_api_keys_organization_id;
DROP INDEX IF EXISTS idx_transactions_organization_id;
DROP INDEX IF EXISTS idx_wallets_organization_id;
DROP INDEX IF EXISTS idx_users_organization_id;

DROP TABLE IF EXISTS "api_usage";
DROP TABLE IF EXISTS "organization_members";
DROP TABLE IF EXISTS "api_keys";

ALTER TABLE "transactions" DROP COLUMN IF EXISTS "organization_id";
ALTER TABLE "wallets" DROP COLUMN IF EXISTS "organization_id";
ALTER TABLE "users" DROP COLUMN IF EXISTS "organization_id";

DROP TABLE IF EXISTS "organizations"; 