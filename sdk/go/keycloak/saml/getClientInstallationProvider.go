// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to retrieve Installation Provider of a SAML Client.
func GetClientInstallationProvider(ctx *pulumi.Context, args *GetClientInstallationProviderArgs, opts ...pulumi.InvokeOption) (*GetClientInstallationProviderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
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
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetClientInstallationProviderResultOutput, error) {
			args := v.(GetClientInstallationProviderArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("keycloak:saml/getClientInstallationProvider:getClientInstallationProvider", args, GetClientInstallationProviderResultOutput{}, options).(GetClientInstallationProviderResultOutput), nil
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
