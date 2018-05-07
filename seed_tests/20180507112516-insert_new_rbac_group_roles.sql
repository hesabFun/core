-- +migrate Up
INSERT INTO `rbac_group_roles` (`id`, `group_id`, `role_id`)
VALUES (1, 1, 1), (2, 1, 2), (3, 1, 3), (4, 1, 4), (5, 1, 5), (6, 1, 6);

-- +migrate Down
DELETE FROM `rbac_group_roles`
WHERE `id` = 1 OR 2 OR 3 OR 4;