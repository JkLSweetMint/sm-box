create database i18n;

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

create table
    if not exists public.languages
(
    code   varchar(5)            not null
        constraint languages_pk
            primary key,
    name   varchar(300),
    active boolean default false not null
);

create table
    if not exists public.sections
(
    id     uuid          default gen_random_uuid() not null
        constraint sections_pk
            primary key,
    parent uuid
        references public.sections(id)
            on delete cascade,

    key    varchar(252)                             not null,
    name   varchar(1024) default ''                 not null,

    constraint check_key
        check (key ~ '^[-0-9a-zA-Z_]{2,252}$'),

    unique (key, parent)
);

create schema
    if not exists texts;

create table
    if not exists texts.texts
(
    id         uuid          not null default gen_random_uuid()
        constraint texts_pk
            primary key,
    language   varchar(5)    not null
        references public.languages(code)
            on delete cascade,
    section    uuid          not null
        references public.sections(id)
            on delete cascade,
    key        varchar(252)  not null,
    value      varchar(4096) not null,

    unique (language, section, key)
);

create table
    if not exists texts.drafts
(
    id         uuid          not null default gen_random_uuid()
        constraint draft_texts_pk
            primary key,
    language   varchar(5)    not null
        references public.languages(code)
            on delete cascade,
    section    uuid          not null
        references public.sections(id)
            on delete cascade,
    key        varchar(252)  not null,
    value      varchar(4096) not null
);

create or replace function public.create_sections_fn(path varchar(1024))
    returns uuid
    language plpgsql as
$$
declare
    keys varchar[];
    sectionID uuid = null;
    sectionKey varchar;
    sectionParent uuid = null;
begin
    select
        into keys string_to_array(path, '.');

    for index in 1..(array_length(keys, 1)) loop
            sectionID = null;
            sectionKey = keys[index];

            if sectionParent is null then
                select
                    into sectionID id
                from
                    public.sections
                where
                    key = sectionKey and
                    (parent is null and sectionParent is null);
            else
                select
                    into sectionID id
                from
                    public.sections
                where
                    key = sectionKey and
                    parent = sectionParent;
            end if;

            if sectionID is null then
                insert into
                    public.sections(key, parent)
                values
                    (sectionKey, sectionParent)
                returning id into sectionID;
            end if;

            sectionParent = sectionID;
        end loop;

    return sectionID;
end;
$$;

create or replace function public.create_texts_for_language_fn()
    returns trigger
    language plpgsql as
$$
begin
    insert into
        texts.texts(
            language,
            section,
            key,
            value
        )
        select
            new.code as language,
            section,
            key,
            '' as value
        from
            texts.texts
        where
            language = public.get_default_language()
        group by key, section;

    return new;
end;
$$;

create trigger create_texts_for_language
    after insert on public.languages
    for each row
execute procedure public.create_texts_for_language_fn();

create or replace function public.create_texts_other_languages_fn()
    returns trigger
    language plpgsql as
$$
begin
    if new.language = public.get_default_language() then
        insert into
            texts.texts(
                language,
                section,
                key,
                value
            )
            select
                languages.code as language,
                new.section as section,
                new.key as key,
                '' as value
            from
                public.languages as languages
            where
                languages.code != new.language and
                (
                    select
                        count(*)
                    from
                        texts.texts
                    where
                        language = languages.code and
                        section = new.section and
                        key = new.key
                ) = 0;
    end if;

    return new;
end;
$$;

create trigger create_texts_other_languages
    after insert on texts.texts
    for each row
execute procedure public.create_texts_other_languages_fn();

create or replace function public.write_text(lang varchar(5), path varchar(1024), val varchar(4096))
    returns void
    language plpgsql as
$$
declare
    keys varchar[];
    key_ varchar(256);
    sectionID uuid;
begin
    select
        into keys string_to_array(path, '.');

    key_ = keys[array_length(keys, 1)];
    path = array_to_string(keys[:array_length(keys, 1)-1],'.');

    select
        into sectionID public.create_sections_fn(path);

    insert into
        texts.texts(
            language,
            section,
            key,
            value
        ) values (
                     lang,
                     sectionID,
                     key_,
                     val
                 ) on conflict (language, section, key) do
            update set
            value = val;
end;
$$;

create or replace function public.write_draft_text(lang varchar(5), path varchar(1024), val varchar(4096))
    returns void
    language plpgsql as
$$
declare
    keys varchar[];
    key_ varchar(256);
    sectionID uuid;
begin
    select
        into keys string_to_array(path, '.');

    key_ = keys[array_length(keys, 1)];
    path = array_to_string(keys[:array_length(keys, 1)-1],'.');

    select
        into sectionID public.create_sections_fn(path);

    insert into
        texts.drafts(
            language,
            section,
            key,
            value
        ) values (
                     lang,
                     sectionID,
                     key_,
                     val
                 );
end;
$$;

create or replace function public.update_draft_text(textID uuid, val varchar(4096))
    returns void
    language plpgsql as
$$
begin
    update
        texts.drafts
    set
        value = val
    where
        id = textID;
end;
$$;

create or replace function public.assemble_dictionary(lang varchar(5), path varchar(1024))
    returns setof record
    language plpgsql as
$$
declare
    result record;

    keys varchar[];
    sectionID uuid = null;
    sectionKey varchar;
    sectionParent uuid = null;

begin
    select
        into keys string_to_array(path, '.');

    for index in 1..(array_length(keys, 1)) loop
            sectionID = null;
            sectionKey = keys[index];

            if sectionParent is null then
                select
                    into sectionID id
                from
                    public.sections
                where
                    key = sectionKey and
                    (parent is null and sectionParent is null);
            else
                select
                    into sectionID id
                from
                    public.sections
                where
                    key = sectionKey and
                    parent = sectionParent;
            end if;

            sectionParent = sectionID;
        end loop;

    for result in
        WITH RECURSIVE cte_sections (id, parent, key, full_key) AS (
            select
                sections.id,
                sections.parent,
                sections.key,
                path as full_key
            from
                public.sections as sections
            where
                sections.id = sectionID

            UNION ALL

            select
                sections.id,
                sections.parent,
                sections.key,
                cte.full_key || '.' || sections.key as full_key
            from
                public.sections as sections
                    JOIN cte_sections cte ON cte.id = sections.parent
        )

        select
            (cte.full_key || '.' || texts.key)::varchar(1024) as key,
            texts.value as value
        from
            texts.texts as texts
                left join cte_sections cte ON texts.section = cte.id
        where
            cte.id = texts.section and
            texts.language = lang

        loop
            return next result;
        end loop;
end;
$$;

create or replace function public.assemble_dictionary_for_edit(lang varchar(5), path varchar(1024))
    returns setof record
    language plpgsql as
$$
declare
    result record;

    keys varchar[];
    sectionID uuid = null;
    sectionKey varchar;
    sectionParent uuid = null;

begin
    select
        into keys string_to_array(path, '.');

    for index in 1..(array_length(keys, 1)) loop
            sectionID = null;
            sectionKey = keys[index];

            if sectionParent is null then
                select
                    into sectionID id
                from
                    public.sections
                where
                    key = sectionKey and
                    (parent is null and sectionParent is null);
            else
                select
                    into sectionID id
                from
                    public.sections
                where
                    key = sectionKey and
                    parent = sectionParent;
            end if;

            sectionParent = sectionID;
        end loop;

    for result in
        WITH RECURSIVE cte_sections (id, parent, key, full_key) AS (
            select
                sections.id,
                sections.parent,
                sections.key,
                path as full_key
            from
                public.sections as sections
            where
                sections.id = sectionID

            UNION ALL

            select
                sections.id,
                sections.parent,
                sections.key,
                cte.full_key || '.' || sections.key as full_key
            from
                public.sections as sections
                    JOIN cte_sections cte ON cte.id = sections.parent
        )

        select
            texts.id as id,
            texts.language as language,
            (cte.full_key || '.' || texts.key)::varchar(1024) as key,
            texts.value as value,
            false as is_draft
        from
            texts.texts as texts
                left join cte_sections cte ON texts.section = cte.id
        where
            cte.id = texts.section and
            texts.language = lang

        UNION ALL

        select
            texts.id as id,
            texts.language as language,
            (cte.full_key || '.' || texts.key)::varchar(1024) as key,
            texts.value as value,
            true as is_draft
        from
            texts.drafts as texts
                left join cte_sections cte ON texts.section = cte.id
        where
            cte.id = texts.section and
            texts.language = lang


        loop
            return next result;
        end loop;
end;
$$;

create or replace function public.draft_to_text(textID uuid)
    returns void
    language plpgsql as
$$
declare
begin
    insert into
        texts.texts(
        language,
        section,
        key,
        value
    )
        (
            select
                language,
                section,
                key,
                value
            from
                texts.drafts as draft
            where
                id = textID
        )
    on conflict (language, section, key)
        do update
        set
            value = (select value from texts.drafts where id = textID);
end;
$$;

create or replace function public.text_to_draft(textID uuid)
    returns void
    language plpgsql as
$$
declare
begin
    insert into
        texts.drafts(
        language,
        section,
        key,
        value
    )
    select
        texts.language,
        texts.section,
        texts.key,
        texts.value
    from
        texts.texts as texts
    where
        texts.id = textID;
end;
$$;

insert into
    public.languages(code, name, active)
values
    ('ru-RU', 'Русский', true),
    ('en-US', 'English', true),
    ('zh-CN', '中文', true);

select
    public.write_text('en-US', 'toasts.error.title', 'An error occured');

select
    public.write_text('ru-RU', 'toasts.error.title', 'Произошла ошибка');

select
    public.write_text('zh-CN', 'toasts.error.title', '发生错误');

select
    public.write_text('en-US', 'auth.form.title', 'Welcome to SM-Box'),
    public.write_text('en-US', 'auth.form.description', 'Please, provide your authorization credentials to proceed. '),
    public.write_text('en-US', 'auth.form.inputs.username.label', 'Username'),
    public.write_text('en-US', 'auth.form.inputs.password.label', 'Password'),
    public.write_text('en-US', 'auth.form.buttons.log_in.text', 'Log in'),
    public.write_text('en-US', 'auth.form.errors.field_is_required', 'Field is required. '),
    public.write_text('en-US', 'auth.form.errors.invalid_value', 'Invalid value. ');

select
    public.write_text('ru-RU', 'auth.form.title', 'Добро пожаловать в SM-Box'),
    public.write_text('ru-RU', 'auth.form.description', 'Пожалуйста, укажите свои учетные данные для авторизации, чтобы продолжить. '),
    public.write_text('ru-RU', 'auth.form.inputs.username.label', 'Имя пользователя'),
    public.write_text('ru-RU', 'auth.form.inputs.password.label', 'Пароль'),
    public.write_text('ru-RU', 'auth.form.buttons.log_in.text', 'Войти'),
    public.write_text('ru-RU', 'auth.form.errors.field_is_required', 'Это поле обязательное. '),
    public.write_text('ru-RU', 'auth.form.errors.invalid_value', 'Недопустимое значение. ');

select
    public.write_text('zh-CN', 'auth.form.title', '欢迎来到SM-Box'),
    public.write_text('zh-CN', 'auth.form.description', '请提供您的登录凭据继续。'),
    public.write_text('zh-CN', 'auth.form.inputs.username.label', '用户名称'),
    public.write_text('zh-CN', 'auth.form.inputs.password.label', '密码'),
    public.write_text('zh-CN', 'auth.form.buttons.log_in.text', '进入'),
    public.write_text('zh-CN', 'auth.form.errors.field_is_required', '这个字段是必需的。'),
    public.write_text('zh-CN', 'auth.form.errors.invalid_value', '无效值。');

select
    public.write_text('en-US', 'auth.project-select.form.title', 'Select a project'),
    public.write_text('en-US', 'auth.project-select.form.inputs.select.label', 'Search options...'),
    public.write_text('en-US', 'auth.project-select.form.buttons.confirm.text', 'Confirm');

select
    public.write_text('ru-RU', 'auth.project-select.form.title', 'Выберите проект'),
    public.write_text('ru-RU', 'auth.project-select.form.inputs.select.label', 'Параметры поиска...'),
    public.write_text('ru-RU', 'auth.project-select.form.buttons.confirm.text', 'Подтвердить');

select
    public.write_text('zh-CN', 'auth.project-select.form.title', '选择项目'),
    public.write_text('zh-CN', 'auth.project-select.form.inputs.select.label', '搜索选项。'),
    public.write_text('zh-CN', 'auth.project-select.form.buttons.confirm.text', '确认');