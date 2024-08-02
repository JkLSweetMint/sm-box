create database notifications;

create schema
    if not exists public;

create schema
    if not exists users;

create table
    if not exists users.notifications (
    id   bigserial not null
           constraint notifications_pk
               primary key,
    type varchar(256),

    sender_id    bigint,
    recipient_id bigint not null,

    title      varchar(300),
    title_i18n uuid,

    text      varchar(300),
    text_i18n uuid,

    created_timestamp timestamptz not null default now(),
    read_timestamp    timestamptz,

    constraint check_title
        check (title is not null or title_i18n is not null),

    constraint check_text
        check (text is not null or text_i18n is not null)
);