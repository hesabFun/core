-- +migrate Up
INSERT INTO `users` (`created_at`, `updated_at`, `deleted_at`, `name`, `email`, `mobile`, `password`, `status`, `type`, `remember_token`, `sms_token`)
VALUES
  ('2018-02-04 18:10:59', '2018-02-04 18:10:59', NULL, 'عرفان', '', '09111111111', '25d55ad283aa400af464c76d713c07ad',
                          'active', 'user', '', 0);

-- +migrate Down
DELETE FROM `users`
WHERE `users`.`mobile` = '09111111111';
