-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE aaa.companies
(
    id         bigserial                NOT NULL,
    name       varchar                  NOT NULL,
    status     int                      NOT NULL DEFAULT 1,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT aaa_companies_id_primary PRIMARY KEY (id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE public.companies;