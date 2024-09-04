# Run go file ./test/sample-server.go on two different processes
go run ./test/sample-server/sample-server.go &
go run ./test/balancer-test.go

curl localhost:8080