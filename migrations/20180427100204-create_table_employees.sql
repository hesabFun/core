-- +migrate Up
CREATE TABLE `employees` (
  `id`                 INT(11) UNSIGNED                                                             NOT NULL AUTO_INCREMENT,
  `user_id`            INT(11) UNSIGNED                                                             NOT NULL,
  `company_id`         INT(11) UNSIGNED                                                             NOT NULL,
  `status_by_employee` ENUM ('pending', 'active', 'block')                                          NOT NULL DEFAULT 'pending',
  `status_by_company`  ENUM ('pending', 'active', 'block')                                          NOT NULL DEFAULT 'pending',
  `type`               ENUM ('none', 'manager', 'accountant', 'headmaster_accountant', 'technical') NOT NULL DEFAULT 'none',
  `created_at`         TIMESTAMP                                                                    NOT NULL DEFAULT current_timestamp(),
  `deleted_at`         TIMESTAMP                                                                    NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  KEY `company_id` (`company_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `componay employee` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `user employee` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `employees`;