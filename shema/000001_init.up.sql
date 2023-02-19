CREATE TABLE users
(
    id serial not null unique primary key,
    name varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null
);

CREATE TABLE admins
(
    id serial not null unique primary key,
    username varchar(255) not null unique,
    password varchar(255) not null
);

CREATE TABLE services
(
    id serial not null unique primary key,
    cost int not null default 0,
    description varchar(255),
    name varchar(255) not null
);

CREATE TABLE applies
(
    id serial not null unique primary key,
    user_id int references users("id") on delete cascade not null,
    service_id int references services("id") on delete cascade not null
);

CREATE TABLE employers
(
    id serial not null unique primary key,
    name varchar(255) not null
);

CREATE TABLE employers_services
(
    id serial not null unique primary key,
    employer_id int references employers(id) on delete cascade not null,
    service_id int references services(id) on delete cascade not null
);

CREATE TABLE operations
(
    id serial not null unique primary key,
    summ int not null
);

CREATE TABLE reports
(
    id serial not null unique primary key,
    created_date timestamp not null default now() not null,
    updated_date timestamp not null default now() not null,
    leads int
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_date = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON reports
FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();
