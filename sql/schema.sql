CREATE DATABASE demo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `demo`;

DROP TABLE IF EXISTS `data`;

CREATE TABLE `data` (
  `id` varchar(36) NOT NULL,
  `latitude` decimal(10, 8),
  `longitude` decimal(11, 8),
  `date_added` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
