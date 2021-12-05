module github.com/mulesoft-consulting/anypoint-client-go

go 1.13

require (
	github.com/mulesoft-consulting/anypoint-client-go/authorization v0.2.0
	github.com/mulesoft-consulting/anypoint-client-go/org v0.3.0
	github.com/mulesoft-consulting/anypoint-client-go/vpc v0.5.0
	github.com/stretchr/testify v1.7.0 // indirect
)

replace (
	github.com/mulesoft-consulting/anypoint-client-go/authorization v0.1.0 => ./authorization
	github.com/mulesoft-consulting/anypoint-client-go/vpc v0.1.0 => ./vpc
)
