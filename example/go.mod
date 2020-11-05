module github.com/Code-Hex/grpc-gate/example

go 1.15

require (
	github.com/Code-Hex/grpc-gate v1.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.33.1
)

replace github.com/Code-Hex/grpc-gate => ../
