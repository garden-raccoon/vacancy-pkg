module github.com/garden-raccoon/orders-admin-package

go 1.23.2

require (
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/misnaged/scriptorium v0.0.0-20231207043744-47446928a2b9
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.1
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.67.1
