set http_proxy=http://127.0.0.1:1080 & set https_proxy=http://127.0.0.1:1080

set GOPROXY=
set CGO_ENABLED=0

go get -u github.com/issueye/lichee-js

go mod tidy

go mod vendor
