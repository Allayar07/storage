### psql command

create table users
(
    id serial not null primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar not null
);

create table file
(
    id serial not null primary key,
    bucket varchar not null,
    file varchar not null,
    key varchar not null
    
);

create table files
(
    id serial not null primary key,
    data json,
    key varchar not null

);