module github.com/wkcw/shippy-cli-consignment

go 1.14

require (
	github.com/micro/go-micro v1.18.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
)

require (
	github.com/micro/go-micro/v2 v2.9.1
	github.com/wkcw/shippy/shippy-service-consignment v0.0.0
)

replace github.com/wkcw/shippy/shippy-service-consignment => ../shippy-service-consignment
