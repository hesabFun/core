-- +migrate Up
CREATE SEQUENCE public.product_variant_values_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.product_variants_id_seq
    OWNER TO postgres;

CREATE TABLE public.product_variant_values
(
    id         bigint                DEFAULT nextval('public.product_variant_values_id_seq'::regclass) PRIMARY KEY NOT NULL,
    variant_id bigint                                                                                              NOT NULL,
    product_id bigint                                                                                              NOT NULL,
    value      character varying(11) DEFAULT ''::character varying                                                 NOT NULL,

    CONSTRAINT "product variant" FOREIGN KEY (product_id) REFERENCES public.products (id) ON DELETE CASCADE,
    CONSTRAINT "variant value" FOREIGN KEY (variant_id) REFERENCES public.product_variants (id) ON DELETE CASCADE
);
ALTER TABLE public.product_variant_values
    OWNER TO postgres;

CREATE INDEX idx_product_variant_values_product_id_and_variant_id ON public.product_variant_values USING btree (product_id, variant_id);

-- +migrate Down
DROP TABLE public.product_variant_values;
DROP SEQUENCE public.product_variant_values_id_seq;