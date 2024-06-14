create table
    if not exists env
(
    key   text not null
        on conflict replace
            constraint env_pk
                primary key,
    value text not null default ''
);

insert into env (key, value) values
                                 ('id', ''),
                                 ('title', ''),
                                 ('description', ''),
                                 ('owner', ''),
                                 ('default_lang', '');
