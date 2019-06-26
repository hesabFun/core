-- +migrate Up
CREATE SEQUENCE public.product_variants_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.product_variants_id_seq
    OWNER TO postgres;

CREATE TABLE public.product_variants
(
    id          bigint                DEFAULT nextval('public.product_variants_id_seq'::regclass) PRIMARY KEY NOT NULL,
    company_id  bigint                                                                                        NOT NULL,
    category_id bigint                                                                                        NOT NULL,
    name        character varying(64) DEFAULT ''::character varying                                           NOT NULL,

    CONSTRAINT "category variants" FOREIGN KEY (category_id) REFERENCES public.product_categories (id) ON DELETE CASCADE,
    CONSTRAINT "company variants" FOREIGN KEY (company_id) REFERENCES public.companies (id) ON DELETE CASCADE
);
ALTER TABLE public.product_variants
    OWNER TO postgres;

CREATE INDEX idx_product_variants_category_id_and_company_id ON public.product_variants USING btree (category_id, company_id);

-- +migrate Down
DROP TABLE public.product_variants;
DROP SEQUENCE public.product_variants_id_seq;