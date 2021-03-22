module github.com/mulesoft-consulting/cloudhub-client-go

go 1.13

require (
	github.com/mulesoft-consulting/cloudhub-client-go/vpc v1.0.0
	github.com/mulesoft-consulting/cloudhub-client-go/authorization v1.0.0
	github.com/stretchr/testify v1.7.0 // indirect
)

replace (
	github.com/mulesoft-consulting/cloudhub-client-go/vpc v1.0.0 => ./vpc
	github.com/mulesoft-consulting/cloudhub-client-go/authorization v1.0.0 => ./authorization
)
