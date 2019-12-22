module github.com/frankrap/goalgo

go 1.13

require (
	github.com/Workiva/go-datastructures v1.0.50
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/golang/protobuf v1.3.1
	github.com/hashicorp/go-plugin v1.0.0
	github.com/nntaoli-project/GoEx v1.0.7
	github.com/sirupsen/logrus v1.4.2
	github.com/sony/sonyflake v0.0.0-20181109022403-6d5bd6181009
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	golang.org/x/net v0.0.0-20190420063019-afa5a82059c6
	google.golang.org/grpc v1.20.1
)

replace github.com/nntaoli-project/GoEx v1.0.7 => github.com/frankrap/GoEx v0.0.0-20191024031732-875e548a111b
