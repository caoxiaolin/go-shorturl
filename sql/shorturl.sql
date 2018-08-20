psql -c "CREATE DATABASE shorturl;" -U postgres
psql -c "CREATE TABLE url (id bigserial, url character varying, hits bigint, create_time timestamp, last_access_time timestamp, PRIMARY KEY (id));" -U postgres -d shorturl

# 清空表
TRUNCATE url RESTART IDENTITY;
