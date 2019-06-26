-- +migrate Up
CREATE SEQUENCE public.companies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.companies_id_seq
    OWNER TO postgres;

CREATE TYPE public.companies_status AS ENUM (
    'pending',
    'active',
    'block'
    );
ALTER TYPE public.companies_status
    OWNER TO postgres;

CREATE TABLE public.companies
(
    id         bigint                   DEFAULT nextval('public.companies_id_seq'::regclass) PRIMARY KEY NOT NULL,
    name       character varying(64)    DEFAULT ''::character varying                                    NOT NULL,
    status     public.companies_status  DEFAULT 'pending'::public.companies_status                       NOT NULL,
    created_at timestamp with time zone DEFAULT now()                                                    NOT NULL,
    deleted_at timestamp with time zone
);
ALTER TABLE public.companies
    OWNER TO postgres;

CREATE INDEX idx_companies_name ON public.companies USING btree (name);

-- +migrate Down
DROP TABLE public.companies;
DROP SEQUENCE public.companies_id_seq;
DROP TYPE public.companies_status;