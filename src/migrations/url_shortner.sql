create database url_shortner;

create schema
    if not exists public;

create or replace function public.generate_short_url(len integer)
    returns varchar
    language plpgsql as
$$
declare
    chars varchar = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    url varchar;
    exist bool;
begin
    loop
        SELECT array_to_string(array(select substr(chars,((random()*(length(chars)-1)+1)::integer),1) from generate_series(1,len)),'') into url;

        select count(urls.*) > 0 into exist from public.urls as urls where urls.reduction = url;

        if not exist then
            return url;
        end if;
    end loop;
end;
$$;

create table
    if not exists public.urls
(
    id        bigserial     not null
                constraint urls_pk
                    primary key,
    source    varchar(2048) not null,
    reduction varchar(16)   not null default public.generate_short_url(16),

    active boolean not null default false,

    constraint check_source
        check (source is not null and source ~ '^(?:[-a-zA-Z0-9()*@:%_\+.~#?&\/=]*)$')
);

create type public.url_type as enum ('proxy', 'redirect');

create table
    if not exists public.properties
(
    url_id         bigint                    not null
        references public.urls(id)
            on delete cascade,
    type           public.url_type           not null,
    number_of_uses integer,
    start_active   timestamptz,
    end_active     timestamptz,

    constraint check_properties
        check ((number_of_uses is not null) or (start_active is not null) or (end_active is not null)),

    constraint check_number_of_uses
        check (number_of_uses >= 0)
);

create or replace function public.create_short_url(source varchar(2048), type public.url_type, number_of_uses integer, start_active timestamptz, end_active timestamptz)
    returns bigint
    language plpgsql as
$$
declare
    urlID bigint;
begin
    insert into
        public.urls(
                    source
        )
    values (
            source
           )
    returning id into urlID;

    insert into
        public.properties(
           url_id,
           type,
           number_of_uses,
           start_active,
           end_active
        )
    values (
            urlID,
            type,
            number_of_uses,
            start_active,
            end_active
   );

    return urlID;
end;
$$;

create type public.usage_history_status as enum ('success', 'failed', 'forbidden');

create table
    if not exists public.usage_history
(
    url_id         bigint                      not null
        references public.urls(id)
            on delete cascade,
    status         public.usage_history_status not null,
    timestamp      timestamptz                 not null default now(),
    token_info     jsonb                       not null
);

create table
    if not exists public.accesses
(

    url_id         bigint not null
        references public.urls(id)
            on delete cascade,
    role_id       bigint,
    permission_id bigint,

    constraint check_accesses
        check (role_id is not null or permission_id is not null)
);
