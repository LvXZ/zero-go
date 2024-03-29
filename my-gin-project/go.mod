module my-gin-project

go 1.12

require (
	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.5.0
	github.com/google/wire v0.4.0
	go-common v1.7.0
	go.uber.org/automaxprocs v1.2.0
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc v1.27.1
	google.golang.org/protobuf v1.27.1
)

replace go-common => git.bilibili.co/platform/go-common v1.7.0
