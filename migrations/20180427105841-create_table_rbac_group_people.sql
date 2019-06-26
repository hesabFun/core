-- +migrate Up
CREATE SEQUENCE public.rbac_group_people_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.rbac_group_people_id_seq
    OWNER TO postgres;

CREATE TABLE public.rbac_group_people
(
    id       bigint DEFAULT nextval('public.rbac_group_people_id_seq'::regclass) PRIMARY KEY NOT NULL,
    group_id bigint                                                                          NOT NULL,
    user_id  bigint                                                                          NOT NULL,

    CONSTRAINT group_people FOREIGN KEY (group_id) REFERENCES public.rbac_groups (id) ON DELETE CASCADE,
    CONSTRAINT "people in group" FOREIGN KEY (user_id) REFERENCES public.users (id) ON DELETE CASCADE
);
ALTER TABLE public.rbac_group_people
    OWNER TO postgres;

CREATE INDEX idx_group_people_group_id_and_user_id ON public.rbac_group_people USING btree (group_id, user_id);

-- +migrate Down
DROP TABLE public.rbac_group_people;
DROP SEQUENCE public.rbac_group_people_id_seq;