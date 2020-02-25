module github.com/pulumi/pulumi-keycloak

go 1.13

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/mrparkers/terraform-provider-keycloak => github.com/pulumi/terraform-provider-keycloak v0.0.0-20200223181932-550ccf3acf66
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi v1.9.1
	github.com/pulumi/pulumi-terraform-bridge v1.6.4
)
