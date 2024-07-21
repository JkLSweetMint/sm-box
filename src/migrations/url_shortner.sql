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

    constraint check_source
        check (source is not null and source ~ '^(?:[-a-zA-Z0-9()*@:%_\+.~#?&\/=]*)$')
);

create type public.url_type as enum ('proxy', 'redirect');

create table
    if not exists public.properties
(
    url            bigint                    not null
        references public.urls(id)
            on delete cascade,
    type           public.url_type           not null,
    number_of_uses integer,
    start_active   timestamptz,
    end_active     timestamptz,

    constraint check_properties
        check ((number_of_uses is not null) or (start_active is not null and end_active is not null))
);

create or replace function public.create_short_url(source varchar(2048), type public.url_type, number_of_uses integer, start_active timestamptz, end_active timestamptz)
    returns void
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
                               url,
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
end;
$$;

create table
    if not exists public.usage_history
(

    url            bigint      not null
        references public.urls(id)
            on delete cascade,
    timestamp      timestamptz not null default now()
);