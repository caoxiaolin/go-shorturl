[![Build Status](https://www.travis-ci.org/caoxiaolin/go-shorturl.svg?branch=master)](https://www.travis-ci.org/caoxiaolin/go-shorturl)
[![codecov.io](https://codecov.io/github/caoxiaolin/go-shorturl/coverage.svg?branch=master)](https://codecov.io/github/caoxiaolin/go-shorturl?branch=master)
# shorturl
ShortURL Service Written by GoLang

# get
go get github.com/caoxiaolin/go-shorturl

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
