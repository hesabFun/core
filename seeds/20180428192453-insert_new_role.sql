-- +migrate Up
INSERT INTO `rbac_roles` (`id`, `alias`, `path`, `method`, `menu`, `order`, `parent_id`)
VALUES
  (1, 'index', '/', 'get', 'no', 0, 0),
  (2, 'login', '/v1/auth/login', 'post', 'no', 0, 0),
  (3, 'get profile', '/v1/auth/profile', 'get', 'no', 0, 0),
  (4, 'get companies', '/v1/companies', 'get', 'yes', 0, 0),
  (5, 'get company details', '/v1/companies/:id', 'get', 'no', 0, 0),
  (6, 'get menu', '/v1/companies/:id/menu', 'get', 'no', 0, 0),
  (7, 'insert new category', '/v1/companies/:id/categories', 'post', 'no', 0, 0),
  (8, 'insert new product', '/v1/companies/:id/products', 'post', 'no', 0, 0),
  (9, 'insert new transaction', '/v1/companies/:id/transactions', 'post', 'no', 0, 0),
  (10, 'get transactions', '/v1/companies/:id/transactions', 'get', 'yes', 0, 0),
  (11, 'get products', '/v1/companies/:id/products', 'get', 'yes', 0, 0),
  (12, 'insert new company', '/v1/companies', 'post', 'no', 0, 0),
  (13, 'add new employee', '/v1/companies/:id/employees', 'post', 'no', 0, 0),
  (14, 'get employees', '/v1/companies/:id/employees', 'get', 'yes', 0, 0);

-- +migrate Down
DELETE FROM `rbac_roles`
WHERE `id` = 1 OR 2 OR 3 OR 4 OR 5 OR 6 OR 7 OR 8 OR 9 OR 10 OR 11 OR 12 OR 13;