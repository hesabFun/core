-- +migrate Up
CREATE TABLE `product_variants` (
  `id`          INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `company_id`  INT(11) UNSIGNED NOT NULL,
  `category_id` INT(11) UNSIGNED NOT NULL,
  `name`        VARCHAR(64)      NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `company variants` (`company_id`),
  KEY `category variants` (`category_id`),
  CONSTRAINT `category variants` FOREIGN KEY (`category_id`) REFERENCES `product_categories` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `company variants` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `product_variants`;