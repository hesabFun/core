-- +migrate Up
INSERT INTO `employees` (`id`, `user_id`, `company_id`, `status`, `type`)
VALUES (1, 1, 1, 'active', 'manager');

-- +migrate Down
DELETE FROM `employees`
WHERE `id` = 1;