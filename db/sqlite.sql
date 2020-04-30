CREATE TABLE `users` (
	`user_id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`email`	TEXT NOT NULL UNIQUE,
	`password`	TEXT NOT NULL,
	`username`	TEXT NOT NULL UNIQUE,
	`email_confirmation_string`	TEXT NOT NULL UNIQUE,
	`created_at`	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at`	TIMESTAMP,
	`is_email_confirmed`	INTEGER NOT NULL DEFAULT 0
);
CREATE TABLE `forgot_password` (
	`forgot_password_id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`user_id`	INTEGER NOT NULL UNIQUE,
	`random_string`	TEXT NOT NULL UNIQUE,
	`created_at`	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at`	TIMESTAMP
);
CREATE TABLE `exchanges` (
	`exchange_id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`name`	TEXT NOT NULL UNIQUE,
	`logo_image_name`	TEXT,
	`created_at`	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at`	TIMESTAMP,
	`is_active`	INTEGER NOT NULL DEFAULT 1
);
CREATE TABLE `exchange_accounts` (
	`exchange_account_id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`user_id`	INTEGER NOT NULL,
	`exchange_id`	INTEGER NOT NULL,
	`name`	TEXT,
	`key`	TEXT NOT NULL,
	`secret`	TEXT NOT NULL,
	`created_at`	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at`	TIMESTAMP,
	`is_active`	INTEGER NOT NULL DEFAULT 1
);
CREATE TABLE `balances` (
	`balance_id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`user_id`	INTEGER NOT NULL,
	`exchange_account_id`	INTEGER NOT NULL,
	`total_btc`	REAL,
	`total_usdt`	REAL,
	`created_at`	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at`	TIMESTAMP
);
