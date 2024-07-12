create database box;

create schema
    if not exists public;

create table
    if not exists public.projects
(
    id          bigserial                not null
        constraint projects_pk
            primary key,
    name        varchar(300)             not null,
    description varchar(3000) default '' not null,
    version     varchar(300)  default '' not null
);