insert into
    users.users(
        email,
        username,
        password
    )
values
    ('jklgreentea@gmail.com', 'root', 'toor');

insert into
    users.accesses (
        user_id,
        role_id
    )
values (1, 1);

insert into
    access_system.roles (project_id, name, is_system)
values
    (1, 'root', false),
    (1, 'user', true);

insert into
    access_system.role_inheritance (parent_id, heir_id)
values
    (1, 3),
    (3, 4);