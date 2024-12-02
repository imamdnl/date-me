-- init.sql
CREATE SCHEMA IF NOT EXISTS date_me_schema;

CREATE TABLE IF NOT EXISTS date_me_schema.users
(
    id                  bigserial not null
    CONSTRAINT users_pk_2
    PRIMARY KEY,
    email               varchar not null
    CONSTRAINT users_pk
    UNIQUE,
    password            varchar not null,
    name                varchar(50),
    age                 integer,
    gender              smallint,
    bio                 text,
    profile_picture_url varchar(255),
    is_verified         boolean default false,
    is_premium          boolean default false,
    created_at          timestamp default now() not null,
    updated_at          timestamp default now() not null
    );

ALTER TABLE date_me_schema.users
    OWNER TO root;