// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch properties of a Keycloak client that uses the SAML protocol.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/saml"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			realmManagement, err := saml.LookupClient(ctx, &saml.LookupClientArgs{
//				RealmId:  "my-realm",
//				ClientId: "realm-management",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// use the data source
//			_, err = keycloak.LookupRole(ctx, &keycloak.LookupRoleArgs{
//				RealmId:  "my-realm",
//				ClientId: pulumi.StringRef(realmManagement.Id),
//				Name:     "realm-admin",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupClient(ctx *pulumi.Context, args *LookupClientArgs, opts ...pulumi.InvokeOption) (*LookupClientResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClientResult
	err := ctx.Invoke("keycloak:saml/getClient:getClient", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClient.
type LookupClientArgs struct {
	// The client id (not its unique ID).
	ClientId string `pulumi:"clientId"`
	// The realm id.
	RealmId string `pulumi:"realmId"`
}

// A collection of values returned by getClient.
type LookupClientResult struct {
	AssertionConsumerPostUrl           string                                       `pulumi:"assertionConsumerPostUrl"`
	AssertionConsumerRedirectUrl       string                                       `pulumi:"assertionConsumerRedirectUrl"`
	AuthenticationFlowBindingOverrides []GetClientAuthenticationFlowBindingOverride `pulumi:"authenticationFlowBindingOverrides"`
	BaseUrl                            string                                       `pulumi:"baseUrl"`
	CanonicalizationMethod             string                                       `pulumi:"canonicalizationMethod"`
	ClientId                           string                                       `pulumi:"clientId"`
	ClientSignatureRequired            bool                                         `pulumi:"clientSignatureRequired"`
	Description                        string                                       `pulumi:"description"`
	Enabled                            bool                                         `pulumi:"enabled"`
	EncryptAssertions                  bool                                         `pulumi:"encryptAssertions"`
	EncryptionCertificate              string                                       `pulumi:"encryptionCertificate"`
	EncryptionCertificateSha1          string                                       `pulumi:"encryptionCertificateSha1"`
	ExtraConfig                        map[string]string                            `pulumi:"extraConfig"`
	ForceNameIdFormat                  bool                                         `pulumi:"forceNameIdFormat"`
	ForcePostBinding                   bool                                         `pulumi:"forcePostBinding"`
	FrontChannelLogout                 bool                                         `pulumi:"frontChannelLogout"`
	FullScopeAllowed                   bool                                         `pulumi:"fullScopeAllowed"`
	// The provider-assigned unique ID for this managed resource.
	Id                              string   `pulumi:"id"`
	IdpInitiatedSsoRelayState       string   `pulumi:"idpInitiatedSsoRelayState"`
	IdpInitiatedSsoUrlName          string   `pulumi:"idpInitiatedSsoUrlName"`
	IncludeAuthnStatement           bool     `pulumi:"includeAuthnStatement"`
	LoginTheme                      string   `pulumi:"loginTheme"`
	LogoutServicePostBindingUrl     string   `pulumi:"logoutServicePostBindingUrl"`
	LogoutServiceRedirectBindingUrl string   `pulumi:"logoutServiceRedirectBindingUrl"`
	MasterSamlProcessingUrl         string   `pulumi:"masterSamlProcessingUrl"`
	Name                            string   `pulumi:"name"`
	NameIdFormat                    string   `pulumi:"nameIdFormat"`
	RealmId                         string   `pulumi:"realmId"`
	RootUrl                         string   `pulumi:"rootUrl"`
	SamlSignatureKeyName            string   `pulumi:"samlSignatureKeyName"`
	SignAssertions                  bool     `pulumi:"signAssertions"`
	SignDocuments                   bool     `pulumi:"signDocuments"`
	SignatureAlgorithm              string   `pulumi:"signatureAlgorithm"`
	SignatureKeyName                string   `pulumi:"signatureKeyName"`
	SigningCertificate              string   `pulumi:"signingCertificate"`
	SigningCertificateSha1          string   `pulumi:"signingCertificateSha1"`
	SigningPrivateKey               string   `pulumi:"signingPrivateKey"`
	SigningPrivateKeySha1           string   `pulumi:"signingPrivateKeySha1"`
	ValidRedirectUris               []string `pulumi:"validRedirectUris"`
}

func LookupClientOutput(ctx *pulumi.Context, args LookupClientOutputArgs, opts ...pulumi.InvokeOption) LookupClientResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClientResultOutput, error) {
			args := v.(LookupClientArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("keycloak:saml/getClient:getClient", args, LookupClientResultOutput{}, options).(LookupClientResultOutput), nil
		}).(LookupClientResultOutput)
}

// A collection of arguments for invoking getClient.
type LookupClientOutputArgs struct {
	// The client id (not its unique ID).
	ClientId pulumi.StringInput `pulumi:"clientId"`
	// The realm id.
	RealmId pulumi.StringInput `pulumi:"realmId"`
}

func (LookupClientOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClientArgs)(nil)).Elem()
}

// A collection of values returned by getClient.
type LookupClientResultOutput struct{ *pulumi.OutputState }

func (LookupClientResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClientResult)(nil)).Elem()
}

func (o LookupClientResultOutput) ToLookupClientResultOutput() LookupClientResultOutput {
	return o
}

func (o LookupClientResultOutput) ToLookupClientResultOutputWithContext(ctx context.Context) LookupClientResultOutput {
	return o
}

func (o LookupClientResultOutput) AssertionConsumerPostUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.AssertionConsumerPostUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) AssertionConsumerRedirectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.AssertionConsumerRedirectUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) AuthenticationFlowBindingOverrides() GetClientAuthenticationFlowBindingOverrideArrayOutput {
	return o.ApplyT(func(v LookupClientResult) []GetClientAuthenticationFlowBindingOverride {
		return v.AuthenticationFlowBindingOverrides
	}).(GetClientAuthenticationFlowBindingOverrideArrayOutput)
}

func (o LookupClientResultOutput) BaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.BaseUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) CanonicalizationMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.CanonicalizationMethod }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientSignatureRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.ClientSignatureRequired }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) EncryptAssertions() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.EncryptAssertions }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) EncryptionCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.EncryptionCertificate }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) EncryptionCertificateSha1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.EncryptionCertificateSha1 }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ExtraConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClientResult) map[string]string { return v.ExtraConfig }).(pulumi.StringMapOutput)
}

func (o LookupClientResultOutput) ForceNameIdFormat() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.ForceNameIdFormat }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) ForcePostBinding() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.ForcePostBinding }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) FrontChannelLogout() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.FrontChannelLogout }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) FullScopeAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.FullScopeAllowed }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClientResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) IdpInitiatedSsoRelayState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.IdpInitiatedSsoRelayState }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) IdpInitiatedSsoUrlName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.IdpInitiatedSsoUrlName }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) IncludeAuthnStatement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.IncludeAuthnStatement }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) LoginTheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.LoginTheme }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) LogoutServicePostBindingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.LogoutServicePostBindingUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) LogoutServiceRedirectBindingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.LogoutServiceRedirectBindingUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) MasterSamlProcessingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.MasterSamlProcessingUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) NameIdFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.NameIdFormat }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.RealmId }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) RootUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.RootUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) SamlSignatureKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.SamlSignatureKeyName }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) SignAssertions() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.SignAssertions }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) SignDocuments() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.SignDocuments }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) SignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.SignatureAlgorithm }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) SignatureKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.SignatureKeyName }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) SigningCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.SigningCertificate }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) SigningCertificateSha1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.SigningCertificateSha1 }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) SigningPrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.SigningPrivateKey }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) SigningPrivateKeySha1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.SigningPrivateKeySha1 }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ValidRedirectUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClientResult) []string { return v.ValidRedirectUris }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClientResultOutput{})
}
