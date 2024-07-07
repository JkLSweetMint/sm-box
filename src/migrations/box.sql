create database box;

create schema
    if not exists public;

create table
    if not exists public.env
(
    key varchar(300) not null
        constraint env_pk
            primary key,
    value varchar(1000) not null
);

insert into
    public.env(key, value)
values
    ('default_language', 'en-US');

create or replace function public.get_default_language()
    returns varchar(5)
    language plpgsql as
$$
begin
    return (
            select
                value
            from
                public.env
            where key = 'default_language'
        );
end;
$$;

create table
    if not exists public.projects
(
    id          bigserial                not null
        constraint projects_pk
            primary key,
    owner_id    bigint                   not null,
    name        varchar(300)             not null,
    description varchar(3000) default '' not null,
    version     varchar(300)  default '' not null
);
