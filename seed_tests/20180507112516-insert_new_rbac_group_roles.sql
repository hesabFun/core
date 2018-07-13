-- +migrate Up
INSERT INTO `rbac_group_roles` (`id`, `group_id`, `role_id`)
VALUES (1, 1, 1), (2, 1, 2), (3, 1, 3), (4, 1, 4), (5, 1, 5), (6, 1, 6), (7, 1, 7), (8, 1, 8), (9, 1, 9), (10, 1, 10),
  (11, 1, 11), (12, 1, 12), (13, 1, 13);

-- +migrate Down
DELETE FROM `rbac_group_roles`
WHERE `id` = 1 OR 2 OR 3 OR 4 OR 5 OR 6 OR 7 OR 8 OR 9 OR 10 OR 11 OR 12 OR 13;