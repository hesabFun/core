-- +migrate Up
CREATE SEQUENCE public.product_photos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.product_photos_id_seq
    OWNER TO postgres;

CREATE TABLE public.product_photos
(
    id         bigint                 DEFAULT nextval('public.product_photos_id_seq'::regclass) PRIMARY KEY NOT NULL,
    product_id bigint                                                                                       NOT NULL,
    file       character varying(255) DEFAULT ''::character varying                                         NOT NULL,
    file_size  bigint                                                                                       NOT NULL,
    width      bigint                                                                                       NOT NULL,
    length     bigint                                                                                       NOT NULL,
    "order"    integer                DEFAULT 0::integer                                                    NOT NULL,

    CONSTRAINT "product photos" FOREIGN KEY (product_id) REFERENCES public.products (id) ON DELETE CASCADE
);
COMMENT ON COLUMN public.product_photos.file IS 'File address and name';
ALTER TABLE public.product_photos
    OWNER TO postgres;

CREATE INDEX idx_product_photos_product_id ON public.product_photos USING btree (product_id);

-- +migrate Down
DROP TABLE public.product_photos;
DROP SEQUENCE public.product_photos_id_seq;