-- +migrate Up
CREATE SEQUENCE public.rbac_groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.rbac_groups_id_seq
    OWNER TO postgres;

CREATE TABLE public.rbac_groups
(
    id         bigint                DEFAULT nextval('public.rbac_groups_id_seq'::regclass) PRIMARY KEY NOT NULL,
    company_id bigint                                                                                   NOT NULL,
    name       character varying(64) DEFAULT ''::character varying                                      NOT NULL,

    CONSTRAINT "company groups" FOREIGN KEY (company_id) REFERENCES public.companies (id) ON DELETE CASCADE
);
ALTER TABLE public.rbac_groups
    OWNER TO postgres;

CREATE INDEX idx_rbac_groups_company_id ON public.rbac_groups USING btree (company_id);

-- +migrate Down
DROP TABLE public.rbac_groups;
DROP SEQUENCE public.rbac_groups_id_seq;