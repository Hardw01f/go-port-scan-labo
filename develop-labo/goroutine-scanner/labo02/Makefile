#/bin/sh
build : PortScanner_check_CPU.go
	GOOS=darwin GOARCH=amd64 go build -o ./bin/PortScan
	GOOS=linux GOARCH=amd64 go build -o ./bin/PortScan_Cents

