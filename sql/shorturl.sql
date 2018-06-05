CREATE TABLE url (
      id bigserial,
      url character varying,
      hits bigint,
      create_time timestamp,
      last_access_time timestamp,
      PRIMARY KEY (id)
)
