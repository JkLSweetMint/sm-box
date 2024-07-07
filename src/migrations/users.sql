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
    project_id bigint,
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
            access_system.roles as roles
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
