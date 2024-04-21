create table if not exists projects
(
    id          integer not null
        constraint projects_pk
            primary key autoincrement,
    title       text    not null,
    description TEXT
);

create table if not exists system_access_roles
(
    id   integer not null
        constraint system_access_roles_pk
            primary key autoincrement,
    project_id integer
        references system_access_roles (id),
    title text    not null

        check (id > 2 or (id = 1 and title = 'root') or (id = 2 and title = 'user'))
);

create table if not exists system_access_roles_inheritance
(
    parent integer not null
        references system_access_roles (id),
    heir integer not null
        references system_access_roles (id)

        check (parent != heir)
);

insert into system_access_roles(id, project_id, title) values (1, null, 'root');
insert into system_access_roles(id, project_id, title) values (2, null, 'user');

insert into system_access_roles_inheritance(parent,  heir) values (1, 2);

create table if not exists users
(
    id       integer not null
        constraint users_pk
            primary key autoincrement,
    email    text    not null
        constraint users_email_un
            unique,
    username text    not null
        constraint users_username_un
            unique,
    password text    not null

        check (username regexp '^[0-9a-zA-Z-_]{3,16}$')
        check (email regexp '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

create table if not exists users_roles (
                                           user_id integer not null
                                               references users(id),
                                           role_id integer not null
                                               references system_access_roles (id)
);

create table if not exists projects_owners (
                                               project_id integer not null
                                                   references projects (id),
                                               owner_id integer not null
                                                   references users (id)
);