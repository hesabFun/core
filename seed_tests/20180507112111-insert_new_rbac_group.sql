-- +migrate Up
INSERT INTO `rbac_groups` (`id`, `company_id`, `name`)
VALUES (1, 1, 'گروه فنی');

-- +migrate Down
DELETE FROM `rbac_groups`
WHERE `id` = 1;