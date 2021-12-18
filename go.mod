module github.com/brettmostert/trple

go 1.17

require (
	github.com/brettmostert/trple-proto-product-go v0.0.0-20211218095504-3b905f266b95
	google.golang.org/grpc v1.43.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace google.golang.org/genproto => github.com/brettmostert/go-genproto v0.0.0-20211208223120-3a66f561d7aa
