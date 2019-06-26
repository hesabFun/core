-- +migrate Up
SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';
SET default_with_oids = false;

-- CREATE SCHEMA public;
-- ALTER SCHEMA public OWNER TO postgres;

CREATE TYPE public.users_status AS ENUM (
    'pending',
    'active',
    'block'
    );
ALTER TYPE public.users_status OWNER TO postgres;

CREATE TYPE public.users_type AS ENUM (
    'user',
    'admin',
    'god'
    );
ALTER TYPE public.users_type OWNER TO postgres;

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.users_id_seq
    OWNER TO postgres;

CREATE TABLE public.users
(
    id             bigint                   DEFAULT nextval('public.users_id_seq'::regclass) PRIMARY KEY NOT NULL,
    name           character varying(255)   DEFAULT '':: character varying                               NOT NULL,
    email          character varying(255)   DEFAULT '':: character varying UNIQUE,
    mobile         character varying(255)   DEFAULT '':: character varying UNIQUE,
    password       character varying(255)   DEFAULT '':: character varying,
    status         public.users_status      DEFAULT 'pending'::public.users_status                       NOT NULL,
    type           public.users_type        DEFAULT 'user'::public.users_type                            NOT NULL,
    remember_token character varying(255)   DEFAULT '':: character varying,
    sms_token      bigint,
    created_at     timestamp with time zone DEFAULT now()                                                NOT NULL,
    deleted_at     timestamp with time zone
);
ALTER TABLE public.users
    OWNER TO postgres;

CREATE UNIQUE INDEX idx_users_email_mobile_name ON public.users USING btree (email, mobile, name);

-- +migrate Down
DROP TABLE public.users;
DROP TYPE public.users_type;
DROP TYPE public.users_status;
DROP SEQUENCE public.users_id_seq;