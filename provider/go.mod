module github.com/pulumi/pulumi-keycloak/provider/v2

go 1.14

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/mrparkers/terraform-provider-keycloak => github.com/pulumi/terraform-provider-keycloak v0.0.0-20200605161740-baa6f31d03ac
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-00010101000000-000000000000
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.4.0
	github.com/pulumi/pulumi/pkg/v2 v2.3.0 // indirect
	github.com/pulumi/pulumi/sdk/v2 v2.3.0
)
