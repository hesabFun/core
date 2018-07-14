-- +migrate Up
INSERT INTO `employees` (`id`, `user_id`, `company_id`, `status_by_employee`, `status_by_company`, `type`)
VALUES (1, 1, 1, 'active', 'active', 'manager');

-- +migrate Down
DELETE FROM `employees`
WHERE `id` = 1;