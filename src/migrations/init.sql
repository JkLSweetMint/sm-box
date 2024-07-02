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
    if not exists i18n;

create table
    if not exists i18n.languages
(
    code varchar(5) not null
        constraint i18n_languages_pk
            primary key,
    name varchar(300)
);

insert into
    i18n.languages(code, name)
values
    ('ru-RU', 'Русский'),
    ('en-US', 'English');

create schema
    if not exists users;

create table
    if not exists users.users
(
    id         bigserial     not null
        constraint users_pk
            primary key,
    project_id bigint,
    email      varchar(1024) not null
        constraint users_email_uq
            unique,
    username   varchar(1024) not null
        constraint users_username_uq
            unique,
    password   varchar(1024) not null,

    constraint check_username
        check (username ~ '^[-0-9a-za-z_]{3,16}$'),

    constraint check_email
        check (email is null or email ~ '^[-a-za-z0-9._%+]+@[-a-za-z0-9.]+\.[a-za-z]{2,}$')
);

create table
    if not exists public.projects
(
    id          bigserial                not null
        constraint projects_pk
            primary key,
    owner_id    bigint                   not null
        references users.users (id),
    name        varchar(300)             not null,
    description varchar(3000) default '' not null,
    version     varchar(300)  default '' not null
);

create schema
    if not exists system_access;

create table
    if not exists system_access.roles
(
    id         bigserial    not null
        constraint system_access_roles_pk
            primary key,
    project_id bigint
        references public.projects (id),
    name       varchar(300) not null,
    is_system  boolean      not null default false,

    check (
        id > 2
            or (id = 1 and name = 'root')
            or (id = 2 and name = 'user')
        )
);

create table
    if not exists system_access.role_inheritance
(
    parent bigint not null
        references system_access.roles (id),
    heir   bigint not null
        references system_access.roles (id),

    check (parent != heir)
);

create or replace function system_access.get_user_access(userID bigint)
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
                system_access.roles as roles
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
                system_access.roles as roles
                    left join system_access.role_inheritance role_inheritance on (role_inheritance.heir = roles.id)
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

    return result;
end;
$$;

insert into
    system_access.roles (project_id, name, is_system)
values
    (null, 'root', true),
    (null, 'user', true);

insert into
    system_access.role_inheritance (parent, heir)
values
    (1, 2);

create table
    if not exists system_access.jwt_tokens
(
    id         bigserial     not null
        constraint system_access_jwt_tokens_pk
            primary key,
    user_id    bigint
        references users.users (id),
    project_id bigint
        references public.projects (id),

    language varchar(5) default public.get_default_language() not null
        references i18n.languages (code),
    data       varchar(4096) not null,

    issued_at  timestamptz   not null,
    not_before timestamptz   not null,
    expires_at timestamptz   not null
);

create table
    if not exists system_access.jwt_token_params
(
    token_id    bigint        not null
        constraint system_access_jwt_token_token_id_uq
            unique,
    remote_addr varchar(1024) not null,
    user_agent  varchar(4096) not null,

    constraint check_remote_addr
        check (remote_addr ~ '^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5})$')
);

create table
    if not exists users.accesses
(
    user_id bigint not null
        references users.users (id),
    role_id bigint not null
        references system_access.roles (id),

    unique (user_id, role_id)
);

create or replace function users.assign_default_role_to_user_fn()
    returns trigger
    language plpgsql as
$$
declare
    projectRoleID bigint;

begin
    if not exists(select * from users.accesses where user_id = new.id) then
        insert into
            users.accesses(
            user_id,
            role_id
        )
        values (
                   new.id,
                   2
               );
    end if;

    if new.project_id is not null then
        select
            into projectRoleID roles.id
        from
            system_access.roles as roles
        where
            roles.project_id = new.project_id and
            roles.name = 'user';

        if projectRoleID is not null and projectRoleID != 0 then
            insert into
                users.accesses(
                user_id,
                role_id
            )
            values (
                       new.id,
                       projectRoleID
                   );
        end if;
    end if;

    return new;
end;
$$;

create trigger assign_default_role_to_user
    after insert on users.users
    for each row
execute procedure users.assign_default_role_to_user_fn();

create or replace function public.create_roles_for_project_fn()
    returns trigger
    language plpgsql as
$$
declare
    ownerRoleID bigint;
    adminRoleID bigint;
    userRoleID bigint;
begin
    insert into
        system_access.roles (project_id, name, is_system)
    values
        (new.id, 'owner', true)
    returning id into ownerRoleID;

    insert into
        system_access.roles (project_id, name, is_system)
    values
        (new.id, 'admin', true)
    returning id into adminRoleID;

    insert into
        system_access.roles (project_id, name, is_system)
    values
        (new.id, 'user', true)
    returning id into userRoleID;

    insert into
        system_access.role_inheritance (parent, heir)
    values
        (1, ownerRoleID),
        (ownerRoleID, adminRoleID),
        (adminRoleID, userRoleID);

    if new.owner_id is not null then
        insert into
            users.accesses
        values
            (new.owner_id, ownerRoleID);
    end if;

    return new;
end;
$$;

create trigger create_roles_for_project
    after insert on public.projects
    for each row
execute procedure public.create_roles_for_project_fn();

create schema
    if not exists transports;

create table
    if not exists transports.http_routes
(
    id            bigserial                 not null
        constraint transports_http_routes_pk
            primary key,
    name          varchar(1024)             not null,
    description   varchar(4096)             not null default '',
    method        varchar(10)               not null,
    path          varchar(4096)             not null,
    register_time timestamptz default now() not null,
    active        boolean     default false not null,
    authorize     boolean     default false not null

    constraint check_active
        check (active is true or active is false),

    constraint check_authorize
        check (authorize is true or authorize is false),

    constraint check_method
        check (
            method = 'GET'
                or method = 'HEAD'
                or method = 'POST'
                or method = 'PUT'
                or method = 'DELETE'
                or method = 'CONNECT'
                or method = 'OPTIONS'
                or method = 'TRACE'
                or method = 'PATCH'
            ),

    constraint check_path
        check (path ~ '^(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$'),

    unique (method, path)
);

create table
    if not exists transports.http_route_accesses
(
    route_id bigint not null
        references transports.http_routes (id),
    role_id  bigint not null
        references system_access.roles (id),

    unique (route_id, role_id)
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

create or replace function transports.http_routes_delete_access_fn()
    returns trigger
    language plpgsql as
$$
begin
    delete from
        transports.http_route_accesses
    where
        route_id = old.id;

    return old;
end;
$$;

create trigger transports_http_routes_delete_access
    before delete on transports.http_routes
    for each row
execute procedure transports.http_routes_delete_access_fn();
