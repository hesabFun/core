-- +migrate Up
CREATE TABLE `users` (
  `id`             INT(11) UNSIGNED                    NOT NULL AUTO_INCREMENT,
  `name`           VARCHAR(255)                        NOT NULL DEFAULT '',
  `email`          VARCHAR(255)                        NOT NULL DEFAULT '',
  `mobile`         VARCHAR(255)                        NOT NULL DEFAULT '',
  `password`       VARCHAR(255)                                 DEFAULT '',
  `status`         ENUM ('pending', 'active', 'block') NOT NULL DEFAULT 'pending',
  `type`           ENUM ('user', 'admin', 'god')       NOT NULL DEFAULT 'user',
  `remember_token` VARCHAR(255)                                 DEFAULT '',
  `sms_token`      INT(11)                                      DEFAULT NULL,
  `created_at`     TIMESTAMP                           NOT NULL DEFAULT current_timestamp(),
  `updated_at`     TIMESTAMP                           NOT NULL DEFAULT current_timestamp(),
  `deleted_at`     TIMESTAMP                           NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`, `mobile`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `users`;
