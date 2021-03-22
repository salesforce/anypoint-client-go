# CLOUDHUB CLIENT GO

This repository hosts golang modules to query resources on anypoint platform. 

Each resource has its own module that was generated using [openapi-generator](https://openapi-generator.tech/).

## How to use 

Include each module separately from your go project using 
```
require github.com/mulesoft-consulting/cloudhub-client-go/[module_name] vx.x.x
```
As an example, if you wanted to import the version 1.0.1 of the `vpc` module you would need to use the following syntax:  
```
require github.com/mulesoft-consulting/cloudhub-client-go/vpc v1.0.1
```

## Modules Versioning 

Each module is released independently from the other, therefore each release is prefixed with the name of the module.

Examples : 
```
vpc/v1.0.0
authorization/v2.2.0
```
