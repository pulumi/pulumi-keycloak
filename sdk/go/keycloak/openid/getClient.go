// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # openid.Client data source
//
// This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
//
// ### Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/openid"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			realmManagement, err := openid.LookupClient(ctx, &openid.LookupClientArgs{
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
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm id.
// - `clientId` - (Required) The client id.
//
// ### Attributes Reference
//
// See the docs for the `openid.Client` resource for details on the exported attributes.
func LookupClient(ctx *pulumi.Context, args *LookupClientArgs, opts ...pulumi.InvokeOption) (*LookupClientResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupClientResult
	err := ctx.Invoke("keycloak:openid/getClient:getClient", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClient.
type LookupClientArgs struct {
	ClientId                              string            `pulumi:"clientId"`
	ConsentScreenText                     *string           `pulumi:"consentScreenText"`
	DisplayOnConsentScreen                *bool             `pulumi:"displayOnConsentScreen"`
	ExtraConfig                           map[string]string `pulumi:"extraConfig"`
	Oauth2DeviceAuthorizationGrantEnabled *bool             `pulumi:"oauth2DeviceAuthorizationGrantEnabled"`
	Oauth2DeviceCodeLifespan              *string           `pulumi:"oauth2DeviceCodeLifespan"`
	Oauth2DevicePollingInterval           *string           `pulumi:"oauth2DevicePollingInterval"`
	RealmId                               string            `pulumi:"realmId"`
}

// A collection of values returned by getClient.
type LookupClientResult struct {
	AccessTokenLifespan                    string                                       `pulumi:"accessTokenLifespan"`
	AccessType                             string                                       `pulumi:"accessType"`
	AdminUrl                               string                                       `pulumi:"adminUrl"`
	AuthenticationFlowBindingOverrides     []GetClientAuthenticationFlowBindingOverride `pulumi:"authenticationFlowBindingOverrides"`
	Authorizations                         []GetClientAuthorization                     `pulumi:"authorizations"`
	BackchannelLogoutRevokeOfflineSessions bool                                         `pulumi:"backchannelLogoutRevokeOfflineSessions"`
	BackchannelLogoutSessionRequired       bool                                         `pulumi:"backchannelLogoutSessionRequired"`
	BackchannelLogoutUrl                   string                                       `pulumi:"backchannelLogoutUrl"`
	BaseUrl                                string                                       `pulumi:"baseUrl"`
	ClientAuthenticatorType                string                                       `pulumi:"clientAuthenticatorType"`
	ClientId                               string                                       `pulumi:"clientId"`
	ClientOfflineSessionIdleTimeout        string                                       `pulumi:"clientOfflineSessionIdleTimeout"`
	ClientOfflineSessionMaxLifespan        string                                       `pulumi:"clientOfflineSessionMaxLifespan"`
	ClientSecret                           string                                       `pulumi:"clientSecret"`
	ClientSessionIdleTimeout               string                                       `pulumi:"clientSessionIdleTimeout"`
	ClientSessionMaxLifespan               string                                       `pulumi:"clientSessionMaxLifespan"`
	ConsentRequired                        bool                                         `pulumi:"consentRequired"`
	ConsentScreenText                      *string                                      `pulumi:"consentScreenText"`
	Description                            string                                       `pulumi:"description"`
	DirectAccessGrantsEnabled              bool                                         `pulumi:"directAccessGrantsEnabled"`
	DisplayOnConsentScreen                 *bool                                        `pulumi:"displayOnConsentScreen"`
	Enabled                                bool                                         `pulumi:"enabled"`
	ExcludeSessionStateFromAuthResponse    bool                                         `pulumi:"excludeSessionStateFromAuthResponse"`
	ExtraConfig                            map[string]string                            `pulumi:"extraConfig"`
	FrontchannelLogoutEnabled              bool                                         `pulumi:"frontchannelLogoutEnabled"`
	FrontchannelLogoutUrl                  string                                       `pulumi:"frontchannelLogoutUrl"`
	FullScopeAllowed                       bool                                         `pulumi:"fullScopeAllowed"`
	// The provider-assigned unique ID for this managed resource.
	Id                                    string   `pulumi:"id"`
	ImplicitFlowEnabled                   bool     `pulumi:"implicitFlowEnabled"`
	LoginTheme                            string   `pulumi:"loginTheme"`
	Name                                  string   `pulumi:"name"`
	Oauth2DeviceAuthorizationGrantEnabled *bool    `pulumi:"oauth2DeviceAuthorizationGrantEnabled"`
	Oauth2DeviceCodeLifespan              *string  `pulumi:"oauth2DeviceCodeLifespan"`
	Oauth2DevicePollingInterval           *string  `pulumi:"oauth2DevicePollingInterval"`
	PkceCodeChallengeMethod               string   `pulumi:"pkceCodeChallengeMethod"`
	RealmId                               string   `pulumi:"realmId"`
	ResourceServerId                      string   `pulumi:"resourceServerId"`
	RootUrl                               string   `pulumi:"rootUrl"`
	ServiceAccountUserId                  string   `pulumi:"serviceAccountUserId"`
	ServiceAccountsEnabled                bool     `pulumi:"serviceAccountsEnabled"`
	StandardFlowEnabled                   bool     `pulumi:"standardFlowEnabled"`
	UseRefreshTokens                      bool     `pulumi:"useRefreshTokens"`
	UseRefreshTokensClientCredentials     bool     `pulumi:"useRefreshTokensClientCredentials"`
	ValidPostLogoutRedirectUris           []string `pulumi:"validPostLogoutRedirectUris"`
	ValidRedirectUris                     []string `pulumi:"validRedirectUris"`
	WebOrigins                            []string `pulumi:"webOrigins"`
}

func LookupClientOutput(ctx *pulumi.Context, args LookupClientOutputArgs, opts ...pulumi.InvokeOption) LookupClientResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClientResultOutput, error) {
			args := v.(LookupClientArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupClientResult
			secret, err := ctx.InvokePackageRaw("keycloak:openid/getClient:getClient", args, &rv, "", opts...)
			if err != nil {
				return LookupClientResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupClientResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupClientResultOutput), nil
			}
			return output, nil
		}).(LookupClientResultOutput)
}

// A collection of arguments for invoking getClient.
type LookupClientOutputArgs struct {
	ClientId                              pulumi.StringInput    `pulumi:"clientId"`
	ConsentScreenText                     pulumi.StringPtrInput `pulumi:"consentScreenText"`
	DisplayOnConsentScreen                pulumi.BoolPtrInput   `pulumi:"displayOnConsentScreen"`
	ExtraConfig                           pulumi.StringMapInput `pulumi:"extraConfig"`
	Oauth2DeviceAuthorizationGrantEnabled pulumi.BoolPtrInput   `pulumi:"oauth2DeviceAuthorizationGrantEnabled"`
	Oauth2DeviceCodeLifespan              pulumi.StringPtrInput `pulumi:"oauth2DeviceCodeLifespan"`
	Oauth2DevicePollingInterval           pulumi.StringPtrInput `pulumi:"oauth2DevicePollingInterval"`
	RealmId                               pulumi.StringInput    `pulumi:"realmId"`
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

func (o LookupClientResultOutput) AccessTokenLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.AccessTokenLifespan }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) AccessType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.AccessType }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) AdminUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.AdminUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) AuthenticationFlowBindingOverrides() GetClientAuthenticationFlowBindingOverrideArrayOutput {
	return o.ApplyT(func(v LookupClientResult) []GetClientAuthenticationFlowBindingOverride {
		return v.AuthenticationFlowBindingOverrides
	}).(GetClientAuthenticationFlowBindingOverrideArrayOutput)
}

func (o LookupClientResultOutput) Authorizations() GetClientAuthorizationArrayOutput {
	return o.ApplyT(func(v LookupClientResult) []GetClientAuthorization { return v.Authorizations }).(GetClientAuthorizationArrayOutput)
}

func (o LookupClientResultOutput) BackchannelLogoutRevokeOfflineSessions() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.BackchannelLogoutRevokeOfflineSessions }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) BackchannelLogoutSessionRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.BackchannelLogoutSessionRequired }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) BackchannelLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.BackchannelLogoutUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) BaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.BaseUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientAuthenticatorType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ClientAuthenticatorType }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientOfflineSessionIdleTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ClientOfflineSessionIdleTimeout }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientOfflineSessionMaxLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ClientOfflineSessionMaxLifespan }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientSessionIdleTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ClientSessionIdleTimeout }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientSessionMaxLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ClientSessionMaxLifespan }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ConsentRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.ConsentRequired }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) ConsentScreenText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClientResult) *string { return v.ConsentScreenText }).(pulumi.StringPtrOutput)
}

func (o LookupClientResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) DirectAccessGrantsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.DirectAccessGrantsEnabled }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) DisplayOnConsentScreen() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClientResult) *bool { return v.DisplayOnConsentScreen }).(pulumi.BoolPtrOutput)
}

func (o LookupClientResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) ExcludeSessionStateFromAuthResponse() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.ExcludeSessionStateFromAuthResponse }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) ExtraConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClientResult) map[string]string { return v.ExtraConfig }).(pulumi.StringMapOutput)
}

func (o LookupClientResultOutput) FrontchannelLogoutEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.FrontchannelLogoutEnabled }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) FrontchannelLogoutUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.FrontchannelLogoutUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) FullScopeAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.FullScopeAllowed }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClientResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ImplicitFlowEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.ImplicitFlowEnabled }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) LoginTheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.LoginTheme }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) Oauth2DeviceAuthorizationGrantEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupClientResult) *bool { return v.Oauth2DeviceAuthorizationGrantEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupClientResultOutput) Oauth2DeviceCodeLifespan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClientResult) *string { return v.Oauth2DeviceCodeLifespan }).(pulumi.StringPtrOutput)
}

func (o LookupClientResultOutput) Oauth2DevicePollingInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClientResult) *string { return v.Oauth2DevicePollingInterval }).(pulumi.StringPtrOutput)
}

func (o LookupClientResultOutput) PkceCodeChallengeMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.PkceCodeChallengeMethod }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.RealmId }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ResourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ResourceServerId }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) RootUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.RootUrl }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ServiceAccountUserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ServiceAccountUserId }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ServiceAccountsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.ServiceAccountsEnabled }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) StandardFlowEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.StandardFlowEnabled }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) UseRefreshTokens() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.UseRefreshTokens }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) UseRefreshTokensClientCredentials() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClientResult) bool { return v.UseRefreshTokensClientCredentials }).(pulumi.BoolOutput)
}

func (o LookupClientResultOutput) ValidPostLogoutRedirectUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClientResult) []string { return v.ValidPostLogoutRedirectUris }).(pulumi.StringArrayOutput)
}

func (o LookupClientResultOutput) ValidRedirectUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClientResult) []string { return v.ValidRedirectUris }).(pulumi.StringArrayOutput)
}

func (o LookupClientResultOutput) WebOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClientResult) []string { return v.WebOrigins }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClientResultOutput{})
}
