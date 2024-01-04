CREATE TABLE `contracts` (
	`id` varchar(128) NOT NULL,
	`type` text NOT NULL,
	`name` text NOT NULL,
	`image` text,
	`symbol` varchar(24) NOT NULL,
	`userId` varchar(128) NOT NULL,
	`address` varchar(42) NOT NULL,
	`deletedAt` timestamp,
	`createdAt` timestamp NOT NULL DEFAULT (now()),
	`updatedAt` timestamp ON UPDATE CURRENT_TIMESTAMP,
	`description` varchar(900),
	CONSTRAINT `contracts_id` PRIMARY KEY(`id`),
	CONSTRAINT `contracts_symbol_unique` UNIQUE(`symbol`)
);
--> statement-breakpoint
CREATE TABLE `nfts` (
	`id` varchar(128) NOT NULL,
	`name` text,
	`price` float DEFAULT 0,
	`image` text,
	`quantity` int NOT NULL,
	`deletedAt` timestamp,
	`createdAt` timestamp NOT NULL DEFAULT (now()),
	`updatedAt` timestamp ON UPDATE CURRENT_TIMESTAMP,
	`contractId` varchar(128),
	`properties` json,
	`description` varchar(900),
	CONSTRAINT `nfts_id` PRIMARY KEY(`id`)
);
--> statement-breakpoint
CREATE TABLE `users` (
	`id` varchar(128) NOT NULL,
	`name` text,
	`email` text,
	`phone` text,
	`avatar` text,
	`address` varchar(42) NOT NULL,
	`createdAt` timestamp NOT NULL DEFAULT (now()),
	`updatedAt` timestamp ON UPDATE CURRENT_TIMESTAMP,
	CONSTRAINT `users_id` PRIMARY KEY(`id`)
);
--> statement-breakpoint
CREATE INDEX `symbol_idx` ON `contracts` (`symbol`);--> statement-breakpoint
CREATE INDEX `address_idx` ON `contracts` (`address`);--> statement-breakpoint
CREATE INDEX `address_idx` ON `users` (`address`);--> statement-breakpoint
ALTER TABLE `contracts` ADD CONSTRAINT `contracts_userId_users_id_fk` FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE `nfts` ADD CONSTRAINT `nfts_contractId_contracts_id_fk` FOREIGN KEY (`contractId`) REFERENCES `contracts`(`id`) ON DELETE no action ON UPDATE no action;