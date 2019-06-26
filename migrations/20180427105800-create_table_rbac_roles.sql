-- +migrate Up
CREATE TYPE public.rbac_roles_method AS ENUM (
    'get',
    'post',
    'put',
    'delete'
    );
ALTER TYPE public.rbac_roles_method OWNER TO postgres;

CREATE TYPE public.rbac_roles_menu AS ENUM (
    'no',
    'yes'
    );
ALTER TYPE public.rbac_roles_menu OWNER TO postgres;

CREATE SEQUENCE public.rbac_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.rbac_roles_id_seq
    OWNER TO postgres;

CREATE TABLE public.rbac_roles
(
    id        bigint                   DEFAULT nextval('public.rbac_roles_id_seq'::regclass) PRIMARY KEY NOT NULL,
    alias     character varying(64)    DEFAULT ''::character varying                                     NOT NULL,
    path      character varying(64)    DEFAULT ''::character varying                                     NOT NULL,
    method    public.rbac_roles_method DEFAULT 'get'::public.rbac_roles_method                           NOT NULL,
    menu      public.rbac_roles_menu   DEFAULT 'no'::public.rbac_roles_menu                              NOT NULL,
    "order"   bigint                                                                                     NOT NULL,
    parent_id bigint                                                                                     NOT NULL
);
ALTER TABLE public.rbac_roles
    OWNER TO postgres;

-- +migrate Down
DROP TABLE public.rbac_roles;
DROP TYPE public.rbac_roles_method;
DROP TYPE public.rbac_roles_menu;
DROP SEQUENCE public.rbac_roles_id_seq;