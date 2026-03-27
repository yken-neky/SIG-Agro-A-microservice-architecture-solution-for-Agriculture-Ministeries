module github.com/sig-agro/services/user-service

go 1.22

require (
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/lib/pq v1.10.9
	github.com/sig-agro/api v0.0.0
	google.golang.org/grpc v1.63.2
)

replace github.com/sig-agro/api => ../../api

require (
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
