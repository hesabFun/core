-- +migrate Up
CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.products_id_seq
    OWNER TO postgres;

CREATE TABLE public.products
(
    id          bigint                DEFAULT nextval('public.products_id_seq'::regclass) PRIMARY KEY NOT NULL,
    company_id  bigint                DEFAULT '0'::bigint                                             NOT NULL,
    category_id bigint                DEFAULT '0'::bigint                                             NOT NULL,
    name        character varying(64) DEFAULT ''::character varying                                   NOT NULL,
    description text                  DEFAULT ''::text                                                NOT NULL,
    price       bigint                                                                                NOT NULL,

    CONSTRAINT "company products" FOREIGN KEY (company_id) REFERENCES public.companies (id) ON DELETE CASCADE,
    CONSTRAINT "product category" FOREIGN KEY (category_id) REFERENCES public.product_categories (id) ON DELETE CASCADE
);
ALTER TABLE public.products
    OWNER TO postgres;

CREATE INDEX idx_products_category_id_and_company_id ON public.products USING btree (category_id, company_id);
CREATE INDEX idx_products_name ON public.products USING gin (to_tsvector('simple'::regconfig, (name)::text));

-- +migrate Down
DROP TABLE public.products;
DROP SEQUENCE public.products_id_seq;