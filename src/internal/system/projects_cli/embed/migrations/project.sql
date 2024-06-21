create table
    if not exists env
(
    key   text not null
        on conflict replace
            constraint env_pk
                primary key,
    value text not null default '',

    constraint check_uuid
        check (key regexp '^([0-9a-zA-Z_]+)$')
);

insert into env (key, value) values
                                 ('id', ''),
                                 ('uuid', ''),
                                 ('name', ''),
                                 ('version', ''),
                                 ('description', ''),
                                 ('owner', '1'),
                                 ('default_lang', 'ru');
