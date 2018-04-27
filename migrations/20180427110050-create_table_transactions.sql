-- +migrate Up
CREATE TABLE `transactions` (
  `id`         BIGINT(20) UNSIGNED      NOT NULL AUTO_INCREMENT,
  `title`      VARCHAR(255)             NOT NULL DEFAULT '',
  `company_id` INT(11) UNSIGNED         NOT NULL,
  `product_id` INT(11) UNSIGNED         NOT NULL,
  `user_id`    INT(11) UNSIGNED         NOT NULL,
  `amount`     BIGINT(20)               NOT NULL,
  `type`       ENUM ('input', 'output') NOT NULL DEFAULT 'input',
  `date`       TIMESTAMP                NOT NULL DEFAULT '0000-00-00 00:00:00',
  `created_at` TIMESTAMP                NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `company transaction` (`company_id`),
  KEY `product transaction` (`product_id`),
  KEY `user transaction` (`user_id`),
  CONSTRAINT `company transaction` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)
    ON UPDATE NO ACTION,
  CONSTRAINT `product transaction` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
    ON UPDATE NO ACTION,
  CONSTRAINT `user transaction` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `transactions`;