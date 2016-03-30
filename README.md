# sorturl
SortURL Service Written by GoLang

# get
go get github.com/caoxiaolin/sorturl

# usage
http服务启动后
post一个url过去，返回短链接
请求短链接，返回原始的URL

# e.g.
curl -d "url=http://www.github.com" "http://192.168.245.128:4000"

http://192.168.245.128:4000/oM1F


curl "http://192.168.245.128:4000/oM1F"

http://www.github.com
