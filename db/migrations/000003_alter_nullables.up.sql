BEGIN;
ALTER TABLE "users" ALTER COLUMN "full_name" SET NOT NULL;
ALTER TABLE "monetary_accounts" ALTER COLUMN "user_uuid" SET NOT NULL;
ALTER TABLE "monetary_accounts" ALTER COLUMN "balance" SET NOT NULL;
ALTER TABLE "monetary_accounts" ALTER COLUMN "currency" SET NOT NULL;
ALTER TABLE "monetary_accounts" ALTER COLUMN "service_id" SET NOT NULL;
ALTER TABLE "transaction_histories" ALTER COLUMN "service_id" SET NOT NULL;
ALTER TABLE "transaction_histories" ALTER COLUMN "mutation" SET NOT NULL;
ALTER TABLE "transaction_histories" ALTER COLUMN "currency" SET NOT NULL;
ALTER TABLE "transaction_histories" ALTER COLUMN "status" SET NOT NULL;
ALTER TABLE "recurring_payments" ALTER COLUMN "account_uuid" SET NOT NULL;
ALTER TABLE "recurring_payments" ALTER COLUMN "scheduler_type" SET NOT NULL;
COMMIT;
