#!/bin/bash

dokku plugin:install https://github.com/dokku/dokku-letsencrypt.git || dokku plugin:update letsencrypt
dokku plugin:install https://github.com/dokku/dokku-redis.git redis || dokku plugin:update redis
dokku plugin:install https://github.com/dokku/dokku-postgres.git postgres || dokku plugin:update postgres

dokku apps:create hesabfun
dokku redis:create hesabfun_redis
dokku redis:link hesabfun_redis hesabfun
dokku postgres:create hesabfun_postgres
dokku postgres:link hesabfun_postgres hesabfun

dokku config:set --no-restart hesabfun DOKKU_LETSENCRYPT_EMAIL=erfun@hesab.fun
dokku config:set --no-restart hesabfun E_SERVICES_SENTRY_ENABLED=true
dokku config:set --no-restart hesabfun E_SERVICES_SENTRY_PROJECT=2
dokku config:set --no-restart hesabfun E_SERVICES_SENTRY_URL=https://hesab.fun
dokku config:set --no-restart hesabfun E_SERVICES_SENTRY_SECRET=${SENTRY_KEY}

// PUSH

dokku letsencrypt hesabfun
