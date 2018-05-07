-- +migrate Up
INSERT INTO `rbac_group_people` (`id`, `group_id`, `user_id`)
VALUES (1, 1, 1);

-- +migrate Down
DELETE FROM `rbac_group_people`
WHERE `id` = 1;