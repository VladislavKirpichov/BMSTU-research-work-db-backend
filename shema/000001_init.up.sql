CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE services
(
    id serial not null unique,
    name varchar(255) not null
);

CREATE TABLE users_services
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    service_id int references services(id) on delete cascade not null
);

CREATE TABLE employers
(
    id serial not null unique,
    name varchar(255) not null
);

CREATE TABLE employers_services
(
    id serial not null unique,
    employer_id int references employers(id) on delete cascade not null,
    service_id int references services(id) on delete cascade not null
);

CREATE TABLE operations
( 
    id serial not null unique,
    summ int not null
);

CREATE TABLE reports
(
    id serial not null unique,
    created_date timestamp not null
);
