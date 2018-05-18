-- +migrate Up
INSERT INTO `product_categories` (`id`, `company_id`, `parent_id`, `name`, `order`)
VALUES (1, 1, 0, 'category 1', 0);

-- +migrate Down
DELETE FROM `rbac_group_people`
WHERE `id` = 1;