CREATE TABLE IF NOT EXISTS `professores` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
  `nome` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL UNIQUE,
  `idade` INT NOT NULL,
  `descricao` TEXT NOT NULL,
  `valor_hora` DECIMAL(10,2) NOT NULL,
  `foto_perfil` VARCHAR(255) NULL,
  `password` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
