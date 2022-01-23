module github.com/sonr-io/sonr/node/highway

go 1.17

require (
	github.com/kataras/golog v0.1.7
	github.com/sonr-io/sonr v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.9.0
	github.com/tendermint/starport v0.19.2
	go.buf.build/grpc/go/sonr-io/highway v1.2.7
	go.buf.build/grpc/go/sonr-io/sonr v1.2.5
	google.golang.org/grpc v1.43.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/sonr-io/sonr => ../../
)
