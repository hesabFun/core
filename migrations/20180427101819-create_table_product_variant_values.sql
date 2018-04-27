-- +migrate Up
CREATE TABLE `product_variant_values` (
  `id`         INT(11)          NOT NULL,
  `variant_id` INT(11) UNSIGNED NOT NULL,
  `product_id` INT(11) UNSIGNED NOT NULL,
  `value`      VARCHAR(11)      NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `variant value` (`variant_id`),
  KEY `product variant` (`product_id`),
  CONSTRAINT `product variant` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `variant value` FOREIGN KEY (`variant_id`) REFERENCES `product_variants` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `product_variant_values`;