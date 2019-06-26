-- +migrate Up
CREATE SEQUENCE public.product_discounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.product_discounts_id_seq
    OWNER TO postgres;

CREATE TABLE public.product_discounts
(
    id         bigint DEFAULT nextval('public.product_discounts_id_seq'::regclass) PRIMARY KEY NOT NULL,
    product_id bigint                                                                          NOT NULL,
    discount   bigint                                                                          NOT NULL,
    start_date timestamp with time zone,
    end_date   timestamp with time zone,

    CONSTRAINT "product discounts" FOREIGN KEY (product_id) REFERENCES public.products (id) ON DELETE CASCADE
);
ALTER TABLE public.product_discounts
    OWNER TO postgres;

CREATE INDEX idx_product_discounts_product_id ON public.product_discounts USING btree (product_id);

-- +migrate Down
DROP TABLE public.product_discounts;
DROP SEQUENCE public.product_discounts_id_seq;