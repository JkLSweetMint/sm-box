create table
    if not exists projects
(
    id          integer not null
        constraint projects_pk
            primary key autoincrement,
    title       text    not null,
    description text
);

create table
    if not exists system_access_roles (
                                          id integer not null
                                              constraint system_access_roles_pk
                                                  primary key autoincrement,
                                          project_id integer
                                              references projects(id),
                                          title text not null

                                              check (
                                                  id > 2
                                                      or (id = 1 and title = 'root')
                                                      or (id = 2 and title = 'user')
                                                  )
);

create table
    if not exists system_access_role_inheritance
(
    parent integer not null
        references system_access_roles (id),
    heir   integer not null
        references system_access_roles (id)
        check (parent != heir)
);

create table
    if not exists system_access_jwt_tokens
(
    id         integer not null
        constraint system_access_jwt_tokens_pk
            primary key autoincrement,
    user_id    integer not null
        references users (id),
    data       varchar(4096),

    created_at text    not null default (datetime('now', 'localtime')),
    expired_at text    not null default (datetime('now', 'localtime', '+8 hours'))
);

create table
    if not exists system_access_jwt_token_params
(
    token_id    integer     not null
        references system_access_jwt_tokens (id) unique,
    remote_addr varchar(21) not null,
    user_agent  text        not null,

    constraint check_remote_addr
        check (remote_addr regexp '^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5})$')
);

insert into
    system_access_roles (id, project_id, title)
values
    (1, null, 'root');

insert into
    system_access_roles (id, project_id, title)
values
    (2, null, 'user');

insert into
    system_access_role_inheritance (parent, heir)
values
    (1, 2);

create table
    if not exists users
(
    id         integer      not null
        constraint users_pk primary key autoincrement,
    project_id integer
        references projects (id),
    email      text
        constraint users_email_un unique,
    username   text         not null
        constraint users_username_un unique,
    password   varchar(512) not null

        constraint check_username
            check (username regexp '^[0-9a-za-z-_]{3,16}$'),

    constraint check_email
        check (email is null or email regexp '^[a-za-z0-9._%+-]+@[a-za-z0-9.-]+\.[a-za-z]{2,}$')
);

create table
    if not exists user_accesses
(
    user_id integer not null
        references users (id),
    role_id integer not null
        references system_access_roles (id),

    unique (user_id, role_id)
);

create table
    if not exists projects_owners
(
    project_id integer not null
        references projects (id) unique,
    owner_id   integer not null
        references users (id),

    unique (project_id, owner_id)
);

create table
    if not exists transports_http_routes
(
    id            integer           not null
        constraint transports_http_routes_pk
            primary key autoincrement,
    active        integer default 0 not null,
    method        text              not null,
    path          text              not null,
    register_time text              not null default (datetime('now', 'localtime')),


    constraint check_active
        check (active = 0 or active = 1),

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
        check (path regexp '^(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$'),

    unique (method, path)
);

create trigger
    if not exists transports_http_routes_to_upper_method
    after insert on transports_http_routes
begin
    insert into
        transports_http_route_accesses(
        route_id,
        role_id
    )
    values (
               new.id,
               1
           );
end;


create trigger
    if not exists transports_http_routes_delete_access
    before delete on transports_http_routes
begin
    delete from
        transports_http_route_accesses
    where
        route_id = old.id;
end;

create table
    if not exists transports_http_route_accesses
(
    route_id integer not null
        references transports_http_routes (id),
    role_id  integer not null
        references system_access_roles (id),

    unique (route_id, role_id)
)