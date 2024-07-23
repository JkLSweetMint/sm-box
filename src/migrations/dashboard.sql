create database dashboard;

create schema
    if not exists public;

create schema
    if not exists sidebar;

create table
    if not exists sidebar.navigation
(
    id         bigserial not null
        constraint sidebar_navigation_pk
            primary key,
    owner_id   bigint    not null,
    title_id   uuid      not null default gen_random_uuid()
);

create table
    if not exists sidebar.navigation_elements
(
    id             uuid      not null default gen_random_uuid()
        constraint sidebar_navigation_elements_pk
            primary key,
    navigation     bigint
        references sidebar.navigation(id)
            on delete cascade,
    parent         uuid
        references sidebar.navigation_elements(id)
            on delete cascade,
    text_id        uuid      not null default gen_random_uuid(),
    icon           varchar(128),
    route          varchar(2048),
    link           varchar(2048)
);

create schema
    if not exists navbar;

create table
    if not exists navbar.shortcuts
(
    id         bigserial                   not null
        constraint navbar_shortcuts_pk
            primary key,
    owner_id   bigint                      not null,
    text_id uuid default gen_random_uuid() not null,
    icon       varchar(2048)
);