module recaptcha/plugin

go 1.23.4

require (
	google.golang.org/protobuf v1.36.6
	recaptcha v0.0.0-00010101000000-000000000000
)

require github.com/golang/protobuf v1.5.4 // indirect

replace recaptcha => ./recaptcha
