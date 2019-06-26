-- +migrate Up
CREATE TYPE public.employees_status_by_employee AS ENUM (
    'pending',
    'active',
    'block'
    );
ALTER TYPE public.employees_status_by_employee OWNER TO postgres;

CREATE TYPE public.employees_status_by_company AS ENUM (
    'pending',
    'active',
    'block'
    );
ALTER TYPE public.employees_status_by_company OWNER TO postgres;

CREATE SEQUENCE public.employees_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.employees_id_seq
    OWNER TO postgres;

CREATE TABLE public.employees
(
    id                 bigint                              DEFAULT nextval('public.employees_id_seq'::regclass) PRIMARY KEY NOT NULL,
    user_id            bigint                                                                                               NOT NULL,
    company_id         bigint                                                                                               NOT NULL,
    status_by_employee public.employees_status_by_employee DEFAULT 'pending'::public.employees_status_by_employee           NOT NULL,
    status_by_company  public.employees_status_by_company  DEFAULT 'pending'::public.employees_status_by_company            NOT NULL,
    created_at         timestamp with time zone            DEFAULT now()                                                    NOT NULL,
    deleted_at         timestamp with time zone,

    CONSTRAINT "company employees" FOREIGN KEY (company_id) REFERENCES public.companies (id) ON DELETE CASCADE,
    CONSTRAINT "user employee" FOREIGN KEY (user_id) REFERENCES public.users (id) ON DELETE CASCADE
);


ALTER TABLE public.employees
    OWNER TO postgres;

CREATE INDEX idx_employees_company_id_and_user_id ON public.employees USING btree (company_id, user_id);

-- +migrate Down
DROP TABLE public.employees;
DROP SEQUENCE public.employees_id_seq;
DROP type public.employees_status_by_employee;
DROP type public.employees_status_by_company;