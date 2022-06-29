module github.com/13sai/go-learing

go 1.15

require (
	github.com/13sai/gohelper v0.0.0-20210412100048-ca02c3b5351c
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/boltdb/bolt v1.3.1
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.7.0
	github.com/goccy/go-json v0.9.7
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/influxdata/influxdb-client-go/v2 v2.5.1
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.9.1
	github.com/muesli/cache2go v0.0.0-20210519043705-f6c4b2d7bc5d
	github.com/nsqio/go-nsq v1.0.8
	github.com/pkg/errors v0.9.1
	github.com/shirou/gopsutil/v3 v3.21.8
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/viper v1.12.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.1
	google.golang.org/genproto v0.0.0-20220525015930-6ca3db687a9d
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.4
	gorm.io/gorm v1.23.5
	gorm.io/plugin/dbresolver v1.2.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1
