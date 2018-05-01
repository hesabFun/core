-- +migrate Up
INSERT INTO `rbac_roles` (`id`, `alias`, `path`, `method`, `menu`, `order`, `parent_id`)
VALUES
  (1, 'index', '/', 'get', 'no', 0, 0),
  (2, 'login', '/v1/auth/login', 'post', 'no', 0, 0),
  (3, 'profile', '/v1/auth/profile', 'get', 'no', 0, 0),
  (4, 'menu', '/v1/menu', 'get', 'no', 0, 0);

-- +migrate Down
DELETE FROM `rbac_roles`
WHERE `id` = 1 OR 2 OR 3 OR 4;