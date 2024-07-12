create database users;

create schema
    if not exists public;

create schema
    if not exists access_system;

create table
    if not exists access_system.roles
(
    id         bigserial    not null
        constraint access_system_roles_pk
            primary key,
    project_id bigint,
    name       varchar(300) not null,
    is_system  boolean      not null default false,

    check (
        id > 2
            or (id = 1 and name = 'root')
            or (id = 2 and name = 'user')
        )
);

create table
    if not exists access_system.role_inheritance
(
    parent bigint not null
        references access_system.roles (id),
    heir   bigint not null
        references access_system.roles (id),

    check (parent != heir)
);

insert into
    access_system.roles (project_id, name, is_system)
values
    (null, 'root', true),
    (null, 'user', true);

insert into
    access_system.role_inheritance (parent, heir)
values
    (1, 2);

create schema
    if not exists users;

create table
    if not exists users.users
(
    id         bigserial     not null
        constraint users_pk
            primary key,
    email      varchar(1024) not null
        constraint users_email_uq
            unique,
    username   varchar(1024) not null
        constraint users_username_uq
            unique,
    password   varchar(1024) not null,

    constraint check_username
        check (username ~ '^[-0-9a-zA-Z_]{3,16}$'),

    constraint check_email
        check (email is null or email ~ '^[-a-zA-Z0-9._%+]+@[-a-zA-Z0-9.]+\.[a-zA-Z]{2,}$')
);

create table
    if not exists users.accesses
(
    user_id bigint not null
        references users.users (id)
            on delete cascade ,
    role_id bigint not null,

    unique (user_id, role_id)
);

create or replace function access_system.get_user_access(userID bigint)
    returns setof record
    language plpgsql as
$$
declare
    result record;

begin
    for result in
        WITH RECURSIVE cte_roles (id, project_id, name, parent) AS (
            select
                roles.id,
                roles.project_id,
                roles.name,
                0::bigint as parent
            from
                access_system.roles as roles
            where
                roles.id in (
                    select
                        accesses.role_id as id
                    from
                        users.users as users
                            left join users.accesses accesses on users.id = accesses.user_id

                    where
                        users.id = userID
                )

            UNION ALL

            select
                roles.id,
                roles.project_id,
                roles.name,
                role_inheritance.parent as parent
            from
                access_system.roles as roles
                    left join access_system.role_inheritance role_inheritance on (role_inheritance.heir = roles.id)
                    JOIN cte_roles cte ON cte.id = role_inheritance.parent
        )

select
    distinct id,
             coalesce(project_id, 0) as project_id,
             name,
             coalesce(parent, 0) as parent
from
    cte_roles
        loop
    return next result;
end loop;
end;
$$;
