// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to retrieve Installation Provider of a SAML Client.
//
// ## Example Usage
//
// In the example below, we extract the SAML metadata IDPSSODescriptor to pass it to the AWS IAM SAML Provider.
//
// ```go
// package main
//
// import (
//
//	"io/ioutil"
//
//	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/saml"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
//				Realm:   pulumi.String("my-realm"),
//				Enabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			samlClient, err := saml.NewClient(ctx, "samlClient", &saml.ClientArgs{
//				RealmId:               realm.ID(),
//				ClientId:              pulumi.String("test-saml-client"),
//				SignDocuments:         pulumi.Bool(false),
//				SignAssertions:        pulumi.Bool(true),
//				IncludeAuthnStatement: pulumi.Bool(true),
//				SigningCertificate:    readFileOrPanic("saml-cert.pem"),
//				SigningPrivateKey:     readFileOrPanic("saml-key.pem"),
//			})
//			if err != nil {
//				return err
//			}
//			samlIdpDescriptor := saml.GetClientInstallationProviderOutput(ctx, saml.GetClientInstallationProviderOutputArgs{
//				RealmId:    realm.ID(),
//				ClientId:   samlClient.ID(),
//				ProviderId: pulumi.String("saml-idp-descriptor"),
//			}, nil)
//			_, err = iam.NewSamlProvider(ctx, "default", &iam.SamlProviderArgs{
//				SamlMetadataDocument: samlIdpDescriptor.ApplyT(func(samlIdpDescriptor saml.GetClientInstallationProviderResult) (*string, error) {
//					return &samlIdpDescriptor.Value, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetClientInstallationProvider(ctx *pulumi.Context, args *GetClientInstallationProviderArgs, opts ...pulumi.InvokeOption) (*GetClientInstallationProviderResult, error) {
	var rv GetClientInstallationProviderResult
	err := ctx.Invoke("keycloak:saml/getClientInstallationProvider:getClientInstallationProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClientInstallationProvider.
type GetClientInstallationProviderArgs struct {
	// The ID of the SAML client. The `id` attribute of a `keycloakClient` resource should be used here.
	ClientId string `pulumi:"clientId"`
	// The ID of the SAML installation provider. Could be one of `saml-idp-descriptor`, `keycloak-saml`, `saml-sp-descriptor`, `keycloak-saml-subsystem`, `mod-auth-mellon`, etc.
	ProviderId string `pulumi:"providerId"`
	// The realm that the SAML client exists within.
	RealmId string `pulumi:"realmId"`
}

// A collection of values returned by getClientInstallationProvider.
type GetClientInstallationProviderResult struct {
	ClientId string `pulumi:"clientId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	ProviderId string `pulumi:"providerId"`
	RealmId    string `pulumi:"realmId"`
	// (Computed) The returned document needed for SAML installation.
	Value string `pulumi:"value"`
}

func GetClientInstallationProviderOutput(ctx *pulumi.Context, args GetClientInstallationProviderOutputArgs, opts ...pulumi.InvokeOption) GetClientInstallationProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClientInstallationProviderResult, error) {
			args := v.(GetClientInstallationProviderArgs)
			r, err := GetClientInstallationProvider(ctx, &args, opts...)
			var s GetClientInstallationProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClientInstallationProviderResultOutput)
}

// A collection of arguments for invoking getClientInstallationProvider.
type GetClientInstallationProviderOutputArgs struct {
	// The ID of the SAML client. The `id` attribute of a `keycloakClient` resource should be used here.
	ClientId pulumi.StringInput `pulumi:"clientId"`
	// The ID of the SAML installation provider. Could be one of `saml-idp-descriptor`, `keycloak-saml`, `saml-sp-descriptor`, `keycloak-saml-subsystem`, `mod-auth-mellon`, etc.
	ProviderId pulumi.StringInput `pulumi:"providerId"`
	// The realm that the SAML client exists within.
	RealmId pulumi.StringInput `pulumi:"realmId"`
}

func (GetClientInstallationProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientInstallationProviderArgs)(nil)).Elem()
}

// A collection of values returned by getClientInstallationProvider.
type GetClientInstallationProviderResultOutput struct{ *pulumi.OutputState }

func (GetClientInstallationProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientInstallationProviderResult)(nil)).Elem()
}

func (o GetClientInstallationProviderResultOutput) ToGetClientInstallationProviderResultOutput() GetClientInstallationProviderResultOutput {
	return o
}

func (o GetClientInstallationProviderResultOutput) ToGetClientInstallationProviderResultOutputWithContext(ctx context.Context) GetClientInstallationProviderResultOutput {
	return o
}

func (o GetClientInstallationProviderResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientInstallationProviderResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClientInstallationProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientInstallationProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetClientInstallationProviderResultOutput) ProviderId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientInstallationProviderResult) string { return v.ProviderId }).(pulumi.StringOutput)
}

func (o GetClientInstallationProviderResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientInstallationProviderResult) string { return v.RealmId }).(pulumi.StringOutput)
}

// (Computed) The returned document needed for SAML installation.
func (o GetClientInstallationProviderResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientInstallationProviderResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClientInstallationProviderResultOutput{})
}
