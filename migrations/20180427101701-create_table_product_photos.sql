-- +migrate Up
CREATE TABLE `product_photos` (
  `id`         INT(11) UNSIGNED                            NOT NULL AUTO_INCREMENT,
  `product_id` INT(11) UNSIGNED                            NOT NULL,
  `file`       VARCHAR(255)                                NOT NULL DEFAULT ''
  COMMENT 'File name',
  `file_size`  INT(11)                                     NOT NULL,
  `file_type`  ENUM ('jpg', 'jpeg', 'png', 'gif', 'other') NOT NULL DEFAULT 'other',
  `width`      INT(11)                                     NOT NULL,
  `length`     INT(11)                                     NOT NULL,
  `order`      INT(2)                                      NOT NULL,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `product photos` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `product_photos`;