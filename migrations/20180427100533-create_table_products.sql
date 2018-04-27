-- +migrate Up
CREATE TABLE `products` (
  `id`          INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `compony_id`  INT(11) UNSIGNED NOT NULL DEFAULT 0,
  `category_id` INT(11) UNSIGNED NOT NULL DEFAULT 0,
  `name`        VARCHAR(64)      NOT NULL DEFAULT '',
  `description` TEXT             NOT NULL DEFAULT '',
  `price`       INT(11)          NOT NULL,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  KEY `company products` (`compony_id`),
  FULLTEXT KEY `product name` (`name`),
  CONSTRAINT `company products` FOREIGN KEY (`compony_id`) REFERENCES `companies` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `product category` FOREIGN KEY (`category_id`) REFERENCES `product_categories` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `products`;