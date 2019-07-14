module quarxlab

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.37.4
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190121172915-509febef88a4
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190301231843-5614ed5bae6f
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190226205417-e64efc72b421
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20181108054448-85acf8d2951c
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190312170243-e65039ee4138
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.3.1
	google.golang.org/appengine => github.com/golang/appengine v1.4.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190404172233-64821d5d2107
	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.0
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dchest/captcha v0.0.0-20170622155422-6a29415a8364
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/jinzhu/gorm v1.9.8
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
)
