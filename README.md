# shorturl
ShortURL Service Written by GoLang

# get
go get github.com/caoxiaolin/go-shorturl

# database
<pre>
su - postgres
psql
CREATE USER rdtest WITH PASSWORD '123456';
CREATE DATABASE shorturl OWNER rdtest;
GRANT ALL PRIVILEGES ON DATABASE shorturl TO rdtest;
</pre>
psql -U rdtest -d shorturl

# usage
http服务启动后
post一个url过去，返回短链接
请求短链接，返回原始的URL

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
