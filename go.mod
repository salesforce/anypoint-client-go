module github.com/mulesoft-consulting/cloudhub-client-go

go 1.13

require (
	github.com/mulesoft-consulting/cloudhub-client-go/vpc v0.0.1
	github.com/mulesoft-consulting/cloudhub-client-go/authorization v0.0.1
	github.com/stretchr/testify v1.7.0 // indirect
)

replace (
	github.com/mulesoft-consulting/cloudhub-client-go/vpc v0.0.1 => ./vpc
	github.com/mulesoft-consulting/cloudhub-client-go/authorization v0.0.1 => ./authorization
)
