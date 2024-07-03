insert into
    users.users(
                 email,
                 username,
                 password
    )
values
    ('jklgreentea@gmail.com', 'root', 'toor');

update
    users.accesses
set
    role_id = 1
where
    user_id = 1;

insert into
    public.projects(
                    owner_id,
                    name,
                    description,
                    version
    )
    values (
            1,
            'System',
            '',
            ''
    );

insert into
    users.users(
        project_id,
        email,
        username,
        password
    )
values
    (1, 'manager1@gmail.com', 'manager1', 'manager1'),
    (2, 'manager@gmail2.com', 'manager2', 'manager2');

insert into
    public.projects(
    owner_id,
    name,
    description,
    version
)
values (
           2,
           'Test',
           '',
           ''
       );

insert into
    public.projects(
    owner_id,
    name,
    description,
    version
)
values (
           1,
           'Test 2',
           '',
           ''
       );

