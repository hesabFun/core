-- +migrate Up
CREATE SEQUENCE public.product_categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.product_categories_id_seq
    OWNER TO postgres;

CREATE TABLE public.product_categories
(
    id         bigint                DEFAULT nextval('public.product_categories_id_seq'::regclass) PRIMARY KEY NOT NULL,
    company_id bigint,
    parent_id  bigint                DEFAULT '0'::bigint,
    name       character varying(64) DEFAULT ''::character varying                                             NOT NULL,
    "order"    bigint                DEFAULT '0'::bigint                                                       NOT NULL,

    CONSTRAINT "company categories" FOREIGN KEY (company_id) REFERENCES public.companies (id) ON DELETE CASCADE
);
ALTER TABLE public.product_categories
    OWNER TO postgres;

CREATE INDEX idx_product_categories_company_id ON public.product_categories USING btree (company_id);
CREATE INDEX idx_product_categories_parent_id ON public.product_categories USING btree (parent_id);

-- +migrate Down
DROP TABLE public.product_categories;
DROP SEQUENCE public.product_categories_id_seq;