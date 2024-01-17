CREATE TABLE IF NOT EXISTS employees.employees
(
    id integer NOT NULL DEFAULT nextval('employees.employees_id_seq'::regclass),
    first_name text COLLATE pg_catalog."default" NOT NULL,
    last_name text COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    hire_date date NOT NULL,
    created_at date NOT NULL,
    updated_at date NOT NULL,
    CONSTRAINT employees_pkey PRIMARY KEY (id)
)