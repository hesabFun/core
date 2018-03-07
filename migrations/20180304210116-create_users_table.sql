-- +migrate Up
CREATE TABLE `users` (
  `id`             INT(10) UNSIGNED                    NOT NULL,
  `created_at`     TIMESTAMP                           NULL     DEFAULT CURRENT_TIMESTAMP,
  `updated_at`     TIMESTAMP                           NULL     DEFAULT CURRENT_TIMESTAMP,
  `deleted_at`     TIMESTAMP                           NULL     DEFAULT NULL,
  `name`           VARCHAR(255)                        NOT NULL DEFAULT '',
  `email`          VARCHAR(255)                        NOT NULL DEFAULT '',
  `mobile`         VARCHAR(255)                        NOT NULL DEFAULT '',
  `password`       VARCHAR(255)                                 DEFAULT NULL,
  `status`         ENUM ('pending', 'active', 'block') NOT NULL DEFAULT 'pending',
  `type`           ENUM ('user', 'admin', 'god')       NOT NULL DEFAULT 'user',
  `remember_token` VARCHAR(255)                                 DEFAULT NULL,
  `sms_token`      INT(11)                                      DEFAULT NULL
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`, `mobile`);

ALTER TABLE `users`
  MODIFY `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

-- +migrate Down
DROP TABLE `users`;
