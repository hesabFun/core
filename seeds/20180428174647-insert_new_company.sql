-- +migrate Up
INSERT INTO `companies` (`id`, `name`, `status`)
VALUES (1, 'حساب‌فان', 'active');

-- +migrate Down
DELETE FROM `companies`
WHERE `id` = 1;