-- +migrate Up
INSERT INTO `users` (`id`, `name`, `email`, `mobile`, `password`, `status`, `type`, `remember_token`, `sms_token`)
VALUES (1, 'عرفان', 'test1@hesabfun.com', '09111111111', '25d55ad283aa400af464c76d713c07ad', 'active', 'user', '', 0);

-- +migrate Down
DELETE FROM `users`
WHERE `users`.`id` = 1;
