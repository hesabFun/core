-- +migrate Up
CREATE TABLE `rbac_group_roles` (
  `id`       INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_id` INT(11) UNSIGNED NOT NULL,
  `role_id`  INT(11) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  KEY `group roles` (`group_id`),
  KEY `role group` (`role_id`),
  CONSTRAINT `group roles` FOREIGN KEY (`group_id`) REFERENCES `rbac_groups` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `role group` FOREIGN KEY (`role_id`) REFERENCES `rbac_roles` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `rbac_group_roles`;