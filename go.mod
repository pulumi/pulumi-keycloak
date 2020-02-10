module github.com/pulumi/pulumi-keycloak

go 1.13

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/mrparkers/terraform-provider-keycloak => github.com/pulumi/terraform-provider-keycloak v0.0.0-20200210232349-e800c4f2e07e
)

require (
	github.com/hashicorp/terraform v0.12.9
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20191218161228-a467c7185cbc
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20191202134852-87cfb4dc8ae1
)
