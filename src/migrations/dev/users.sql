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