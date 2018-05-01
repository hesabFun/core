-- +migrate Up
CREATE TABLE `rbac_roles` (
  `id`        INT(11) UNSIGNED                      NOT NULL AUTO_INCREMENT,
  `alias`     VARCHAR(64)                           NOT NULL DEFAULT '',
  `path`      VARCHAR(64)                           NOT NULL DEFAULT '',
  `method`    ENUM ('get', 'post', 'put', 'delete') NOT NULL DEFAULT 'get',
  `menu`      ENUM ('no', 'yes')                    NOT NULL DEFAULT 'no',
  `order`     INT(11)                               NOT NULL,
  `parent_id` INT(11) UNSIGNED                      NOT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `rbac_roles`;