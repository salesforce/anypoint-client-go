# ANYPOINT CLIENT GO

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

## Disclaimer 
**This is an [open source software, please review the considerations](LICENSE.md).** 
This is an open source project, it does not form part of the official MuleSoft product stack, and is therefore not included in MuleSoft support SLAs. Issues should be directed to the community, who will try to assist on a best endeavours basis. This application is distributed **as is**. 
