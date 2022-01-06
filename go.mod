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

replace v2ray.com/core => github.com/v2fly/v2ray-core v1.24.5-0.20200531043819-9dc12961fac5
