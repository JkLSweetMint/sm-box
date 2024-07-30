create database dashboard;

create schema
    if not exists public;

create schema
    if not exists sidebar;

create table
    if not exists sidebar.navigation
(
    id         bigserial not null
        constraint sidebar_navigation_pk
            primary key,
    owner_id   bigint,

    title      varchar(300),
    title_i18n uuid,

    active boolean not null default false,

    constraint check_title
        check (title is not null or title_i18n is not null)
);

create table
    if not exists sidebar.navigation_sections
(
    id             bigserial not null
        constraint sidebar_navigation_sections_pk
            primary key,
    navigation_id  bigint
        references sidebar.navigation(id)
            on delete cascade,

    title      varchar(300),
    title_i18n uuid,

    active boolean not null default false,

    constraint check_title
        check (title is not null or title_i18n is not null)
);

create table
    if not exists sidebar.navigation_elements
(
    id             bigserial not null
        constraint sidebar_navigation_elements_pk
            primary key,
    parent_id      bigint
        references sidebar.navigation_elements(id)
            on delete cascade,
    navigation_id  bigint    not null
        references sidebar.navigation(id)
            on delete cascade,
    section_id  bigint
        references sidebar.navigation_sections(id)
            on delete cascade,

    index integer not null,
    icon  varchar(128),

    title      varchar(300),
    title_i18n uuid,

    route varchar(2048),
    link  varchar(2048),

    active boolean not null default false,

    check (route is not null or link is not null),

    constraint check_title
        check (title is not null or title_i18n is not null),

    unique (navigation_id, section_id, parent_id, index)
);

create or replace function sidebar.sidebar_navigation_elements_create_index_fn()
    returns trigger
    language plpgsql as
$$
begin
    if (new.index is null) then
        select
            into new.index index + 1 as index
        from
            sidebar.navigation_elements
        where
            navigation_id = new.navigation_id and
            parent_id = new.parent_id and
            section_id = new.section_id
        order by index desc limit 1;
    end if;

    if (new.index is null) then
        new.index = 0;
    end if;

    return new;
end;
$$;

create trigger sidebar_navigation_elements_create_index
    before insert on sidebar.navigation_elements
    for each row
execute procedure sidebar.sidebar_navigation_elements_create_index_fn();

create schema
    if not exists navbar;

create table
    if not exists navbar.shortcuts
(
    id         bigserial      not null
        constraint navbar_shortcuts_pk
            primary key,
    owner_id   bigint         not null,
    icon       varchar(2048),

    title      varchar(300),
    title_i18n uuid,

    active boolean not null default false,

    constraint check_title
        check (title is not null or title_i18n is not null)
);
