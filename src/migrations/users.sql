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

    name      varchar(300),
    name_i18n uuid,

    description      varchar(2048),
    description_i18n uuid,

    is_system  boolean not null default false,

    constraint check_name
        check (name is not null or name_i18n is not null),

    check (
        id > 3
            or (id = 1 and name = 'root')
            or (id = 2 and name = 'user')
            or (id = 3 and name = 'guest')
        )
);

create table
    if not exists access_system.role_inheritance
(
    parent_id bigint not null
        references access_system.roles (id)
            on delete cascade,
    heir_id   bigint not null
        references access_system.roles (id)
            on delete cascade,

    check (parent_id != heir_id)
);

create table
    if not exists access_system.permissions
(
    id         bigserial    not null
        constraint access_system_permissions_pk
            primary key,
    project_id bigint,

    name      varchar(300),
    name_i18n uuid,

    description      varchar(2048),
    description_i18n uuid,

    is_system  boolean not null default false

    constraint check_name
        check (name is not null or name_i18n is not null)
);

create table
    if not exists access_system.role_permissions
(
    role_id         bigint not null
            references access_system.roles (id)
                on delete cascade,
    permission_id   bigint not null
            references access_system.permissions (id)
                on delete cascade
);

insert into
    access_system.roles (project_id, name, is_system)
values
    (null, 'root', true),
    (null, 'user', false),
    (null, 'guest', false);

insert into
    access_system.role_inheritance (parent_id, heir_id)
values
    (1, 2),
    (2, 3);

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
    user_id       bigint not null
        references users.users (id)
            on delete cascade,
    role_id       bigint
        references access_system.roles (id)
            on delete cascade,
    permission_id bigint
        references access_system.permissions (id)
            on delete cascade,

    constraint check_accesses
        check ((role_id is not null and permission_id is null) or (role_id is null and permission_id is not null)),

    unique (user_id, permission_id, role_id)
);

create or replace function access_system.get_user_roles(userID bigint)
    returns setof record
    language plpgsql as
$$
declare
    result record;

begin
    for result in
        WITH RECURSIVE cte_roles (id, project_id, parent_id, name, name_i18n, description, description_i18n, is_system) AS (
            select
                roles.id,
                roles.project_id,
                0::bigint as parent_id,
                roles.name,
                roles.name_i18n,
                roles.description,
                roles.description_i18n,
                roles.is_system
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
                role_inheritance.parent_id as parent_id,
                roles.name,
                roles.name_i18n,
                roles.description,
                roles.description_i18n,
                roles.is_system
            from
                access_system.roles as roles
                    left join access_system.role_inheritance role_inheritance on (role_inheritance.heir_id = roles.id)
                    JOIN cte_roles cte ON cte.id = role_inheritance.parent_id
        )

        select
            distinct id,
                     coalesce(project_id, 0) as project_id,
                     coalesce(parent_id, 0) as parent_id,
                     name,
                     name_i18n,
                     description,
                     description_i18n,
                     is_system
        from
            cte_roles
        loop
    return next result;
end loop;
end;
$$;

create or replace function access_system.get_user_permissions(userID bigint)
    returns setof record
    language plpgsql as
$$
declare
    result record;

begin
    for result in
        WITH RECURSIVE cte_roles (id, project_id, parent_id) AS (
            select
                roles.id,
                roles.project_id,
                0::bigint as parent_id
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
                role_inheritance.parent_id as parent_id
            from
                access_system.roles as roles
                    left join access_system.role_inheritance role_inheritance on (role_inheritance.heir_id = roles.id)
                    JOIN cte_roles cte ON cte.id = role_inheritance.parent_id
        )

        select
            distinct permissions.id,
                     permissions.project_id,
                     cte.id,
                     permissions.name,
                     permissions.name_i18n,
                     permissions.description,
                     permissions.description_i18n,
                     permissions.is_system
        from
            cte_roles as cte
                left join access_system.role_permissions as role_permissions on role_permissions.role_id = cte.id
                left join access_system.permissions as permissions on permissions.id = role_permissions.permission_id
        where
            permissions.id is not null

        UNION ALL

        select
            permissions.id,
            permissions.project_id,
            accesses.role_id,
            permissions.name,
            permissions.name_i18n,
            permissions.description,
            permissions.description_i18n,
            permissions.is_system
        from
            users.accesses as accesses
                left join access_system.permissions as permissions on permissions.id = accesses.permission_id
        where
            accesses.user_id = userID and
            permissions.id is not null

        loop
            return next result;
        end loop;
end;
$$;

create function access_system.get_all_roles() returns SETOF record
    language plpgsql
as
$$
declare
    result record;

begin
    for result in
        WITH RECURSIVE cte_roles (id, project_id, parent_id, name, name_i18n, description, description_i18n, is_system) AS (
            select
                roles.id,
                roles.project_id,
                0::bigint as parent_id,
                roles.name,
                roles.name_i18n,
                roles.description,
                roles.description_i18n,
                roles.is_system
            from
                access_system.roles as roles
            where
                roles.id in (
                    select
                        accesses.role_id as id
                    from
                        users.users as users
                            left join users.accesses accesses on users.id = accesses.user_id
                )

            UNION ALL

            select
                roles.id,
                roles.project_id,
                role_inheritance.parent_id as parent_id,
                roles.name,
                roles.name_i18n,
                roles.description,
                roles.description_i18n,
                roles.is_system
            from
                access_system.roles as roles
                    left join access_system.role_inheritance role_inheritance on (role_inheritance.heir_id = roles.id)
                    JOIN cte_roles cte ON cte.id = role_inheritance.parent_id
        )

        select
            distinct id,
                     coalesce(project_id, 0) as project_id,
                     coalesce(parent_id, 0) as parent_id,
                     name,
                     name_i18n,
                     description,
                     description_i18n,
                     is_system
        from
            cte_roles
        loop
            return next result;
        end loop;
end;
$$;

create or replace function access_system.get_all_permissions()
    returns setof record
    language plpgsql as
$$
declare
    result record;

begin
    for result in
        WITH RECURSIVE cte_roles (id, project_id, parent_id) AS (
            select
                roles.id,
                roles.project_id,
                0::bigint as parent_id
            from
                access_system.roles as roles
            where
                roles.id in (
                    select
                        accesses.role_id as id
                    from
                        users.users as users
                            left join users.accesses accesses on users.id = accesses.user_id
                )

            UNION ALL

            select
                roles.id,
                roles.project_id,
                role_inheritance.parent_id as parent_id
            from
                access_system.roles as roles
                    left join access_system.role_inheritance role_inheritance on (role_inheritance.heir_id = roles.id)
                    JOIN cte_roles cte ON cte.id = role_inheritance.parent_id
        )

        select
            distinct permissions.id,
                     permissions.project_id,
                     cte.id,
                     permissions.name,
                     permissions.name_i18n,
                     permissions.description,
                     permissions.description_i18n,
                     permissions.is_system
        from
            cte_roles as cte
                left join access_system.role_permissions as role_permissions on role_permissions.role_id = cte.id
                left join access_system.permissions as permissions on permissions.id = role_permissions.permission_id
        where
            permissions.id is not null

        UNION ALL

        select
            permissions.id,
            permissions.project_id,
            accesses.role_id,
            permissions.name,
            permissions.name_i18n,
            permissions.description,
            permissions.description_i18n,
            permissions.is_system
        from
            users.accesses as accesses
                left join access_system.permissions as permissions on permissions.id = accesses.permission_id
        where
            permissions.id is not null

        loop
            return next result;
        end loop;
end;
$$;