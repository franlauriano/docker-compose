----------------------------------------------------------------------------------------------------
-- VERSION 1.0.0
----------------------------------------------------------------------------------------------------

-- create tables

CREATE TABLE beaches
(
    id BIGSERIAL PRIMARY KEY,
    ranking int,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name varchar (256) NOT NULL CHECK (name <> ''),
    state varchar (256) NOT NULL CHECK (state <> '')
);
