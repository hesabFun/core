-- +migrate Up
CREATE TABLE `product_categories` (
  `id`         INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `company_id` INT(11) UNSIGNED          DEFAULT NULL,
  `parent_id`  INT(11) UNSIGNED          DEFAULT 0,
  `name`       VARCHAR(64)      NOT NULL DEFAULT '',
  `order`      INT(11)          NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `company_id` (`company_id`),
  KEY `parent` (`parent_id`),
  CONSTRAINT `company categories` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `product_categories`;