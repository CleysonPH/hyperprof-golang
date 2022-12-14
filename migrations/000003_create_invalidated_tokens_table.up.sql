CREATE TABLE IF NOT EXISTS `invalidated_tokens` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
  `token` TEXT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
