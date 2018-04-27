-- +migrate Up
CREATE TABLE `rbac_roles` (
  `id`        INT(11) UNSIGNED                      NOT NULL AUTO_INCREMENT,
  `name`      VARCHAR(64)                           NOT NULL DEFAULT '',
  `role`      VARCHAR(64)                           NOT NULL DEFAULT '',
  `pre_role`  VARCHAR(255)                          NOT NULL DEFAULT '',
  `type`      ENUM ('get', 'post', 'put', 'delete') NOT NULL DEFAULT 'get',
  `menu`      ENUM ('no', 'yes')                    NOT NULL DEFAULT 'no',
  `oeder`     INT(11)                               NOT NULL,
  `parent_id` INT(11) UNSIGNED                      NOT NULL,
  PRIMARY KEY (`id`),
  KEY `parent role` (`parent_id`),
  CONSTRAINT `parent role` FOREIGN KEY (`parent_id`) REFERENCES `rbac_roles` (`id`)
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `rbac_roles`;