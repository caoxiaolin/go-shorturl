# shorturl
ShortURL Service Written by GoLang

# get
go get github.com/caoxiaolin/go-shorturl

# database
<pre>
> su - postgres
> psql
postgres=# CREATE USER rdtest WITH PASSWORD '123456';
postgres=# CREATE DATABASE shorturl OWNER rdtest;
postgres=# GRANT ALL PRIVILEGES ON DATABASE shorturl TO rdtest;
postgres=# \q

> psql -U rdtest -d shorturl
shorturl=> CREATE TABLE url (
      id bigserial,
      url character varying,
      hits bigint,
      create_time timestamp,
      last_access_time timestamp,
      PRIMARY KEY (id)
);
shorturl=> \q
</pre>
# usage
http服务启动后
post一个url过去，返回短链接
请求短链接，返回原始的URL，链接不存在，则返回404

# e.g.
<pre>
curl -d "url=http://www.github.com" "http://127.0.0.1:4000"
http://127.0.0.1:4000/oM1F
curl "http://127.0.0.1:4000/oM1F"
http://www.github.com
</pre>
请求不存在的短链时：
<pre>
curl "http://127.0.0.1:4000/xx79y" -i
HTTP/1.1 404 Not Found
Date: Tue, 05 Jun 2018 10:18:12 GMT
Content-Length: 0
</pre>
