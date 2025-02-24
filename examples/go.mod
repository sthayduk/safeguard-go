module github.com/sthayduk/safeguard-go/examples

go 1.24

require github.com/sthayduk/safeguard-go/src v0.0.0

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace github.com/sthayduk/safeguard-go/src => ../src
