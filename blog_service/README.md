## blog_gin

```text
1.gin
go get -u github.com/gin-gonic/gin@latest

2.viper
go get -u github.com/spf13/viper@latest

3.gorm
go get -u github.com/jinzhu/gorm@latest

4.log
go get -u gopkg.in/natefinch/lumberjack.v2

5.swagger

6.playground
go get -u github.com/go-playground/validator/v10

7.jwt
go get -u github.com/dgrijalva/jwt-go@v3.2.0

8.gomail
go get -u gopkg.in/gomail.v2

9.dingTalk
go get -u github.com/blinkbean/dingtalk

10.ratelimit
go get -u github.com/juju/ratelimit@1.0.2

11.jarger
docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
jaegertracing/all-in-one

12.opentracing
go get -u github.com/opentracing/opentracing-go
go get -u github.com/uber/jaeger-client-go

13.sql tracing
go get -u github.com/eddycjy/opentracing-gorm
go get -u github.com/smacker/opentracing-gorm

14.go-bindata

15.fsnotify
go get -u golang.org/x/sys/...
go get -u github.com/gsnotify/fsnotify

16.zap
go get -u go.uber.org/zap
```