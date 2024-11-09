CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT PRIMARY KEY
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"username" TEXT NOT NULL,
"email" TEXT NOT NULL,
"admin" bool NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
, "password_hash" TEXT NOT NULL DEFAULT '');
CREATE TABLE IF NOT EXISTS "products" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"price" decimal NOT NULL,
"description" TEXT NOT NULL,
"stock" INTEGER NOT NULL,
"category_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "categories" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"description" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "carts" (
"id" TEXT PRIMARY KEY,
"user_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "cart_items" (
"id" TEXT PRIMARY KEY,
"cart_id" char(36) NOT NULL,
"product_id" char(36) NOT NULL,
"quantity" INTEGER NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "wishlists" (
"id" TEXT PRIMARY KEY,
"user_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "wishlist_items" (
"id" TEXT PRIMARY KEY,
"wishlist_id" char(36) NOT NULL,
"product_id" char(36) NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
