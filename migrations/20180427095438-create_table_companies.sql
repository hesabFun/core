-- +migrate Up
CREATE TABLE `companies` (
  `id`   INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(64)      NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `company categories` (`name`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `companies`;