-- +migrate Up
INSERT INTO `users` (`id`, `name`, `email`, `mobile`, `password`, `status`, `type`, `remember_token`, `sms_token`)
VALUES (1, 'عرفان', 'test1@hesabfun.com', '09111111111', '25d55ad283aa400af464c76d713c07ad', 'active', 'user', '', 0),
  (2, 'employee test', 'employee1@hesabfun.com', '09111111112', '25d55ad283aa400af464c76d713c07ad', 'active', 'user',
   '', 0),
  (3, 'verify user by sms', '', '09111111113', '25d55ad283aa400af464c76d713c07ad', 'pending', 'user', '', 1234);

-- +migrate Down
DELETE FROM users
WHERE id = 1 OR 2 OR 3;
