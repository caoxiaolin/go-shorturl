su - postgres
psql

CREATE USER rdtest WITH PASSWORD '123456';
CREATE DATABASE shorturl OWNER rdtest;
GRANT ALL PRIVILEGES ON DATABASE shorturl TO rdtest;

psql -U rdtest -d shorturl

CREATE TABLE url (
      id bigserial,
      url character varying,
      hits bigint,
      create_time timestamp,
      last_access_time timestamp,
      PRIMARY KEY (id)
);

TRUNCATE url RESTART IDENTITY;
