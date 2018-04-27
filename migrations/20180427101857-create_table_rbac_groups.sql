-- +migrate Up
CREATE TABLE `rbac_groups` (
  `id`         INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `company_id` INT(11) UNSIGNED NOT NULL,
  `name`       VARCHAR(64)      NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `company group` (`company_id`),
  CONSTRAINT `company group` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +migrate Down
DROP TABLE `rbac_groups`;