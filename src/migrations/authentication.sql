create database authentication;

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

create schema
    if not exists transports;

create type transports.http_method as enum ('GET', 'HEAD', 'POST', 'PUT', 'DELETE', 'CONNECT', 'OPTIONS', 'TRACE', 'PATCH');

create type transports.http_protocol as enum ('http', 'https', 'ws', 'wss');

create table
    if not exists transports.http_routes
(
    id            bigserial                  not null
        constraint transports_http_routes_pk
            primary key,
    system_name   varchar(1024)              not null,
    name          varchar(1024)              not null,
    description   varchar(4096)              not null default '',
    protocols     transports.http_protocol[] not null default '{}',
    method        transports.http_method     not null,
    path          varchar(4096),
    regexp_path   varchar(4096),
    active        boolean                    not null default false,
    authorize     boolean                    not null default false

    constraint check_path
        check (((path is not null and path ~ '^(?:[-a-zA-Z0-9()*@:%_\+.~#?&\/=]*)$') or regexp_path is not null)),

    unique (system_name, protocols, method, path)
);

create table
    if not exists transports.http_route_accesses
(
    route_id       bigint not null
            references transports.http_routes (id)
                on delete cascade ,
    role_id        bigint,
    permission_id  bigint,

    constraint check_http_route_accesses
        check ((role_id is not null and permission_id is null) or (role_id is null and permission_id is not null)),

    unique (route_id, role_id, permission_id)
);

create or replace function transports.create_http_route_accesses_for_root_fn()
    returns trigger
    language plpgsql as
$$
begin
    insert into
        transports.http_route_accesses(
            route_id,
            role_id
        )
        values (
                   new.id,
                   1
               );

        return new;
end;
$$;

create trigger transports_create_http_route_accesses_for_root
    after insert on transports.http_routes
    for each row
execute procedure transports.create_http_route_accesses_for_root_fn();
