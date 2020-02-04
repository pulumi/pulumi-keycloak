module github.com/pulumi/pulumi-keycloak

go 1.13

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

require (
	github.com/aws/aws-sdk-go v1.25.3 // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/hashicorp/terraform v0.12.9
	github.com/hashicorp/terraform-config-inspect v0.0.0-20191115094559-17f92b0546e8 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20191218161228-a467c7185cbc
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20191202134852-87cfb4dc8ae1
	github.com/spf13/afero v1.2.2 // indirect
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586 // indirect
	golang.org/x/net v0.0.0-20191009170851-d66e71096ffb // indirect
	google.golang.org/grpc v1.23.0 // indirect
)
