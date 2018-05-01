-- +migrate Up
CREATE TABLE `companies` (
  `id`         INT(11) UNSIGNED                    NOT NULL AUTO_INCREMENT,
  `name`       VARCHAR(64)                         NOT NULL DEFAULT '',
  `status`     ENUM ('pending', 'active', 'block') NOT NULL DEFAULT 'pending',
  `created_at` TIMESTAMP                           NOT NULL DEFAULT current_timestamp(),
  `deleted_at` TIMESTAMP                           NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  KEY `company categories` (`name`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `companies`;