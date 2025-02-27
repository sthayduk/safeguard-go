module github.com/sthayduk/safeguard-go/tests

go 1.24.0

replace github.com/sthayduk/safeguard-go => ../pkg

require github.com/sthayduk/safeguard-go v0.0.0-00010101000000-000000000000

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
