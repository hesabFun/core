-- +migrate Up
CREATE TABLE `employees` (
  `id`         INT(11) UNSIGNED                                                             NOT NULL AUTO_INCREMENT,
  `user_id`    INT(11) UNSIGNED                                                             NOT NULL,
  `company_id` INT(11) UNSIGNED                                                             NOT NULL,
  `status`     ENUM ('pending', 'active', 'block')                                          NOT NULL DEFAULT 'pending',
  `type`       ENUM ('none', 'manager', 'accountant', 'headmaster_accountant', 'technical') NOT NULL DEFAULT 'none',
  PRIMARY KEY (`id`),
  KEY `company_id` (`company_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `componay employe` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `user employe` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `employees`;