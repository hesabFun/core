-- +migrate Up
CREATE SEQUENCE public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.transactions_id_seq
    OWNER TO postgres;

CREATE TABLE public.transactions
(
    id         bigint                   DEFAULT nextval('public.transactions_id_seq'::regclass) PRIMARY KEY NOT NULL,
    title      character varying(255)   DEFAULT ''::character varying                                       NOT NULL,
    company_id bigint,
    product_id bigint,
    user_id    bigint,
    amount     numeric                                                                                      NOT NULL,
    date       timestamp with time zone,
    created_at timestamp with time zone DEFAULT now()                                                       NOT NULL,

    CONSTRAINT "company transactions" FOREIGN KEY (company_id) REFERENCES public.companies (id) ON DELETE RESTRICT,
    CONSTRAINT "product transactions" FOREIGN KEY (product_id) REFERENCES public.products (id) ON DELETE RESTRICT,
    CONSTRAINT "user transactions" FOREIGN KEY (user_id) REFERENCES public.users (id) ON DELETE RESTRICT
);
ALTER TABLE public.transactions
    OWNER TO postgres;

CREATE INDEX idx_transaction_company_id_product_id_user_id ON public.transactions USING btree (company_id, product_id, user_id);

-- +migrate Down
DROP TABLE public.transactions;
DROP SEQUENCE public.transactions_id_seq;