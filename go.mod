module main

go 1.17

require (
	github.com/aliyun/fc-runtime-go-sdk v0.0.3
	github.com/aws/aws-sdk-go-v2 v1.11.2
	github.com/aws/aws-sdk-go-v2/config v1.11.1
	github.com/aws/aws-sdk-go-v2/credentials v1.6.5
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.13.2
	github.com/v2fly/vmessping v0.3.4
	v2ray.com/core v4.19.1+incompatible
)

require (
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.8.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.0.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.5.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.7.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.12.0 // indirect
	github.com/aws/smithy-go v1.9.0 // indirect
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/refraction-networking/utls v0.0.0-20190909200633-43c36d3c1f57 // indirect
	go.starlark.net v0.0.0-20190919145610-979af19b165c // indirect
	golang.org/x/crypto v0.0.0-20191029031824-8986dd9e96cf // indirect
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 // indirect
	golang.org/x/sys v0.0.0-20190412213103-97732733099d // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20180831171423-11092d34479b // indirect
	google.golang.org/grpc v1.24.0 // indirect
)

replace v2ray.com/core => github.com/v2fly/v2ray-core v1.24.5-0.20200531043819-9dc12961fac5
