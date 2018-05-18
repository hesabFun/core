-- +migrate Up
INSERT INTO `products` (`id`, `company_id`, `category_id`, `name`, `description`, `price`)
VALUES (1, 1, 1, 'product 1', 'description for this product', 12000);

-- +migrate Down
DELETE FROM `products`
WHERE `id` = 1;