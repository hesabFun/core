-- +migrate Up
CREATE TABLE `rbac_group_people` (
  `id`       INT(11)          NOT NULL,
  `group_id` INT(11) UNSIGNED NOT NULL,
  `user_id`  INT(11) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  KEY `people in group` (`user_id`),
  KEY `group_people` (`group_id`),
  CONSTRAINT `group_people` FOREIGN KEY (`group_id`) REFERENCES `rbac_groups` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `people in group` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `rbac_group_people`;