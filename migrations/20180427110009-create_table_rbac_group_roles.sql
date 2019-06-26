-- +migrate Up
CREATE SEQUENCE public.rbac_group_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.rbac_group_roles_id_seq
    OWNER TO postgres;

CREATE TABLE public.rbac_group_roles
(
    id       bigint DEFAULT nextval('public.rbac_group_roles_id_seq'::regclass) PRIMARY KEY NOT NULL,
    group_id bigint                                                                         NOT NULL,
    role_id  bigint                                                                         NOT NULL,

    CONSTRAINT "group roles" FOREIGN KEY (group_id) REFERENCES public.rbac_groups (id) ON DELETE CASCADE,
    CONSTRAINT "role group" FOREIGN KEY (role_id) REFERENCES public.rbac_roles (id) ON DELETE CASCADE
);
ALTER TABLE public.rbac_group_roles
    OWNER TO postgres;

CREATE INDEX idx_rbac_group_roles_group_id_and_role_id ON public.rbac_group_roles USING btree (group_id, role_id);

-- +migrate Down
DROP TABLE public.rbac_group_roles;
DROP SEQUENCE public.rbac_group_roles_id_seq;