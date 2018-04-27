-- +migrate Up
CREATE TABLE `product_discounts` (
  `id`         INT(11) UNSIGNED         NOT NULL AUTO_INCREMENT,
  `product_id` INT(11) UNSIGNED         NOT NULL,
  `discount`   INT(11)                  NOT NULL,
  `type`       ENUM ('percent', 'cash') NOT NULL DEFAULT 'percent',
  `start_date` TIMESTAMP                NOT NULL DEFAULT '0000-00-00 00:00:00',
  `end_date`   TIMESTAMP                NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  KEY `product discount` (`product_id`),
  CONSTRAINT `product discount` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `product_discounts`;