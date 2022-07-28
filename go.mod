module github.com/ory/hydra-client-go

go 1.17

require (
	github.com/go-openapi/errors v0.20.2
	github.com/go-openapi/runtime v0.19.31
	github.com/go-openapi/strfmt v0.21.3
	github.com/go-openapi/swag v0.21.1
	github.com/go-openapi/validate v0.21.0
)

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/loads v0.21.1 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	go.mongodb.org/mongo-driver v1.10.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/go-openapi/loads => github.com/goshippo/loads v0.21.2

replace github.com/go-openapi/validate => github.com/goshippo/validate v0.22.1

replace github.com/go-openapi/runtime => github.com/goshippo/runtime v0.24.2
