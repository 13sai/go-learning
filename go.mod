module github.com/13sai/go-learing

go 1.15

require (
	github.com/13sai/gohelper v0.0.0-20210412100048-ca02c3b5351c
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.7.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.9.1
	github.com/muesli/cache2go v0.0.0-20210519043705-f6c4b2d7bc5d
	github.com/nsqio/go-nsq v1.0.8
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cast v1.3.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.4.0
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1
