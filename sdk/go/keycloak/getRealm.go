// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch properties of a Keycloak realm for
// usage with other resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			realm, err := keycloak.LookupRealm(ctx, &keycloak.LookupRealmArgs{
//				Realm: "my-realm",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// use the data source
//			_, err = keycloak.NewRole(ctx, "group", &keycloak.RoleArgs{
//				RealmId: pulumi.String(realm.Id),
//				Name:    pulumi.String("group"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRealm(ctx *pulumi.Context, args *LookupRealmArgs, opts ...pulumi.InvokeOption) (*LookupRealmResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRealmResult
	err := ctx.Invoke("keycloak:index/getRealm:getRealm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRealm.
type LookupRealmArgs struct {
	Attributes                  map[string]string              `pulumi:"attributes"`
	DefaultDefaultClientScopes  []string                       `pulumi:"defaultDefaultClientScopes"`
	DefaultOptionalClientScopes []string                       `pulumi:"defaultOptionalClientScopes"`
	DisplayNameHtml             *string                        `pulumi:"displayNameHtml"`
	Internationalizations       []GetRealmInternationalization `pulumi:"internationalizations"`
	OtpPolicy                   *GetRealmOtpPolicy             `pulumi:"otpPolicy"`
	// The realm name.
	Realm                      string                              `pulumi:"realm"`
	SecurityDefenses           []GetRealmSecurityDefense           `pulumi:"securityDefenses"`
	SmtpServers                []GetRealmSmtpServer                `pulumi:"smtpServers"`
	WebAuthnPasswordlessPolicy *GetRealmWebAuthnPasswordlessPolicy `pulumi:"webAuthnPasswordlessPolicy"`
	WebAuthnPolicy             *GetRealmWebAuthnPolicy             `pulumi:"webAuthnPolicy"`
}

// A collection of values returned by getRealm.
type LookupRealmResult struct {
	AccessCodeLifespan                  string            `pulumi:"accessCodeLifespan"`
	AccessCodeLifespanLogin             string            `pulumi:"accessCodeLifespanLogin"`
	AccessCodeLifespanUserAction        string            `pulumi:"accessCodeLifespanUserAction"`
	AccessTokenLifespan                 string            `pulumi:"accessTokenLifespan"`
	AccessTokenLifespanForImplicitFlow  string            `pulumi:"accessTokenLifespanForImplicitFlow"`
	AccountTheme                        string            `pulumi:"accountTheme"`
	ActionTokenGeneratedByAdminLifespan string            `pulumi:"actionTokenGeneratedByAdminLifespan"`
	ActionTokenGeneratedByUserLifespan  string            `pulumi:"actionTokenGeneratedByUserLifespan"`
	AdminTheme                          string            `pulumi:"adminTheme"`
	Attributes                          map[string]string `pulumi:"attributes"`
	BrowserFlow                         string            `pulumi:"browserFlow"`
	ClientAuthenticationFlow            string            `pulumi:"clientAuthenticationFlow"`
	ClientSessionIdleTimeout            string            `pulumi:"clientSessionIdleTimeout"`
	ClientSessionMaxLifespan            string            `pulumi:"clientSessionMaxLifespan"`
	DefaultDefaultClientScopes          []string          `pulumi:"defaultDefaultClientScopes"`
	DefaultOptionalClientScopes         []string          `pulumi:"defaultOptionalClientScopes"`
	DefaultSignatureAlgorithm           string            `pulumi:"defaultSignatureAlgorithm"`
	DirectGrantFlow                     string            `pulumi:"directGrantFlow"`
	DisplayName                         string            `pulumi:"displayName"`
	DisplayNameHtml                     *string           `pulumi:"displayNameHtml"`
	DockerAuthenticationFlow            string            `pulumi:"dockerAuthenticationFlow"`
	DuplicateEmailsAllowed              bool              `pulumi:"duplicateEmailsAllowed"`
	EditUsernameAllowed                 bool              `pulumi:"editUsernameAllowed"`
	EmailTheme                          string            `pulumi:"emailTheme"`
	Enabled                             bool              `pulumi:"enabled"`
	FirstBrokerLoginFlow                string            `pulumi:"firstBrokerLoginFlow"`
	// The provider-assigned unique ID for this managed resource.
	Id                               string                             `pulumi:"id"`
	InternalId                       string                             `pulumi:"internalId"`
	Internationalizations            []GetRealmInternationalization     `pulumi:"internationalizations"`
	LoginTheme                       string                             `pulumi:"loginTheme"`
	LoginWithEmailAllowed            bool                               `pulumi:"loginWithEmailAllowed"`
	Oauth2DeviceCodeLifespan         string                             `pulumi:"oauth2DeviceCodeLifespan"`
	Oauth2DevicePollingInterval      int                                `pulumi:"oauth2DevicePollingInterval"`
	OfflineSessionIdleTimeout        string                             `pulumi:"offlineSessionIdleTimeout"`
	OfflineSessionMaxLifespan        string                             `pulumi:"offlineSessionMaxLifespan"`
	OfflineSessionMaxLifespanEnabled bool                               `pulumi:"offlineSessionMaxLifespanEnabled"`
	OrganizationsEnabled             bool                               `pulumi:"organizationsEnabled"`
	OtpPolicy                        GetRealmOtpPolicy                  `pulumi:"otpPolicy"`
	PasswordPolicy                   string                             `pulumi:"passwordPolicy"`
	Realm                            string                             `pulumi:"realm"`
	RefreshTokenMaxReuse             int                                `pulumi:"refreshTokenMaxReuse"`
	RegistrationAllowed              bool                               `pulumi:"registrationAllowed"`
	RegistrationEmailAsUsername      bool                               `pulumi:"registrationEmailAsUsername"`
	RegistrationFlow                 string                             `pulumi:"registrationFlow"`
	RememberMe                       bool                               `pulumi:"rememberMe"`
	ResetCredentialsFlow             string                             `pulumi:"resetCredentialsFlow"`
	ResetPasswordAllowed             bool                               `pulumi:"resetPasswordAllowed"`
	RevokeRefreshToken               bool                               `pulumi:"revokeRefreshToken"`
	SecurityDefenses                 []GetRealmSecurityDefense          `pulumi:"securityDefenses"`
	SmtpServers                      []GetRealmSmtpServer               `pulumi:"smtpServers"`
	SslRequired                      string                             `pulumi:"sslRequired"`
	SsoSessionIdleTimeout            string                             `pulumi:"ssoSessionIdleTimeout"`
	SsoSessionIdleTimeoutRememberMe  string                             `pulumi:"ssoSessionIdleTimeoutRememberMe"`
	SsoSessionMaxLifespan            string                             `pulumi:"ssoSessionMaxLifespan"`
	SsoSessionMaxLifespanRememberMe  string                             `pulumi:"ssoSessionMaxLifespanRememberMe"`
	UserManagedAccess                bool                               `pulumi:"userManagedAccess"`
	VerifyEmail                      bool                               `pulumi:"verifyEmail"`
	WebAuthnPasswordlessPolicy       GetRealmWebAuthnPasswordlessPolicy `pulumi:"webAuthnPasswordlessPolicy"`
	WebAuthnPolicy                   GetRealmWebAuthnPolicy             `pulumi:"webAuthnPolicy"`
}

func LookupRealmOutput(ctx *pulumi.Context, args LookupRealmOutputArgs, opts ...pulumi.InvokeOption) LookupRealmResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRealmResultOutput, error) {
			args := v.(LookupRealmArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("keycloak:index/getRealm:getRealm", args, LookupRealmResultOutput{}, options).(LookupRealmResultOutput), nil
		}).(LookupRealmResultOutput)
}

// A collection of arguments for invoking getRealm.
type LookupRealmOutputArgs struct {
	Attributes                  pulumi.StringMapInput                  `pulumi:"attributes"`
	DefaultDefaultClientScopes  pulumi.StringArrayInput                `pulumi:"defaultDefaultClientScopes"`
	DefaultOptionalClientScopes pulumi.StringArrayInput                `pulumi:"defaultOptionalClientScopes"`
	DisplayNameHtml             pulumi.StringPtrInput                  `pulumi:"displayNameHtml"`
	Internationalizations       GetRealmInternationalizationArrayInput `pulumi:"internationalizations"`
	OtpPolicy                   GetRealmOtpPolicyPtrInput              `pulumi:"otpPolicy"`
	// The realm name.
	Realm                      pulumi.StringInput                         `pulumi:"realm"`
	SecurityDefenses           GetRealmSecurityDefenseArrayInput          `pulumi:"securityDefenses"`
	SmtpServers                GetRealmSmtpServerArrayInput               `pulumi:"smtpServers"`
	WebAuthnPasswordlessPolicy GetRealmWebAuthnPasswordlessPolicyPtrInput `pulumi:"webAuthnPasswordlessPolicy"`
	WebAuthnPolicy             GetRealmWebAuthnPolicyPtrInput             `pulumi:"webAuthnPolicy"`
}

func (LookupRealmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRealmArgs)(nil)).Elem()
}

// A collection of values returned by getRealm.
type LookupRealmResultOutput struct{ *pulumi.OutputState }

func (LookupRealmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRealmResult)(nil)).Elem()
}

func (o LookupRealmResultOutput) ToLookupRealmResultOutput() LookupRealmResultOutput {
	return o
}

func (o LookupRealmResultOutput) ToLookupRealmResultOutputWithContext(ctx context.Context) LookupRealmResultOutput {
	return o
}

func (o LookupRealmResultOutput) AccessCodeLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.AccessCodeLifespan }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) AccessCodeLifespanLogin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.AccessCodeLifespanLogin }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) AccessCodeLifespanUserAction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.AccessCodeLifespanUserAction }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) AccessTokenLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.AccessTokenLifespan }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) AccessTokenLifespanForImplicitFlow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.AccessTokenLifespanForImplicitFlow }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) AccountTheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.AccountTheme }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) ActionTokenGeneratedByAdminLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.ActionTokenGeneratedByAdminLifespan }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) ActionTokenGeneratedByUserLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.ActionTokenGeneratedByUserLifespan }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) AdminTheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.AdminTheme }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRealmResult) map[string]string { return v.Attributes }).(pulumi.StringMapOutput)
}

func (o LookupRealmResultOutput) BrowserFlow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.BrowserFlow }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) ClientAuthenticationFlow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.ClientAuthenticationFlow }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) ClientSessionIdleTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.ClientSessionIdleTimeout }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) ClientSessionMaxLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.ClientSessionMaxLifespan }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) DefaultDefaultClientScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRealmResult) []string { return v.DefaultDefaultClientScopes }).(pulumi.StringArrayOutput)
}

func (o LookupRealmResultOutput) DefaultOptionalClientScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRealmResult) []string { return v.DefaultOptionalClientScopes }).(pulumi.StringArrayOutput)
}

func (o LookupRealmResultOutput) DefaultSignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.DefaultSignatureAlgorithm }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) DirectGrantFlow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.DirectGrantFlow }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) DisplayNameHtml() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRealmResult) *string { return v.DisplayNameHtml }).(pulumi.StringPtrOutput)
}

func (o LookupRealmResultOutput) DockerAuthenticationFlow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.DockerAuthenticationFlow }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) DuplicateEmailsAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.DuplicateEmailsAllowed }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) EditUsernameAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.EditUsernameAllowed }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) EmailTheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.EmailTheme }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) FirstBrokerLoginFlow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.FirstBrokerLoginFlow }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRealmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.InternalId }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) Internationalizations() GetRealmInternationalizationArrayOutput {
	return o.ApplyT(func(v LookupRealmResult) []GetRealmInternationalization { return v.Internationalizations }).(GetRealmInternationalizationArrayOutput)
}

func (o LookupRealmResultOutput) LoginTheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.LoginTheme }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) LoginWithEmailAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.LoginWithEmailAllowed }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) Oauth2DeviceCodeLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.Oauth2DeviceCodeLifespan }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) Oauth2DevicePollingInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRealmResult) int { return v.Oauth2DevicePollingInterval }).(pulumi.IntOutput)
}

func (o LookupRealmResultOutput) OfflineSessionIdleTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.OfflineSessionIdleTimeout }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) OfflineSessionMaxLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.OfflineSessionMaxLifespan }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) OfflineSessionMaxLifespanEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.OfflineSessionMaxLifespanEnabled }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) OrganizationsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.OrganizationsEnabled }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) OtpPolicy() GetRealmOtpPolicyOutput {
	return o.ApplyT(func(v LookupRealmResult) GetRealmOtpPolicy { return v.OtpPolicy }).(GetRealmOtpPolicyOutput)
}

func (o LookupRealmResultOutput) PasswordPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.PasswordPolicy }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.Realm }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) RefreshTokenMaxReuse() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRealmResult) int { return v.RefreshTokenMaxReuse }).(pulumi.IntOutput)
}

func (o LookupRealmResultOutput) RegistrationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.RegistrationAllowed }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) RegistrationEmailAsUsername() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.RegistrationEmailAsUsername }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) RegistrationFlow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.RegistrationFlow }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) RememberMe() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.RememberMe }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) ResetCredentialsFlow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.ResetCredentialsFlow }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) ResetPasswordAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.ResetPasswordAllowed }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) RevokeRefreshToken() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.RevokeRefreshToken }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) SecurityDefenses() GetRealmSecurityDefenseArrayOutput {
	return o.ApplyT(func(v LookupRealmResult) []GetRealmSecurityDefense { return v.SecurityDefenses }).(GetRealmSecurityDefenseArrayOutput)
}

func (o LookupRealmResultOutput) SmtpServers() GetRealmSmtpServerArrayOutput {
	return o.ApplyT(func(v LookupRealmResult) []GetRealmSmtpServer { return v.SmtpServers }).(GetRealmSmtpServerArrayOutput)
}

func (o LookupRealmResultOutput) SslRequired() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.SslRequired }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) SsoSessionIdleTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.SsoSessionIdleTimeout }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) SsoSessionIdleTimeoutRememberMe() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.SsoSessionIdleTimeoutRememberMe }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) SsoSessionMaxLifespan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.SsoSessionMaxLifespan }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) SsoSessionMaxLifespanRememberMe() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRealmResult) string { return v.SsoSessionMaxLifespanRememberMe }).(pulumi.StringOutput)
}

func (o LookupRealmResultOutput) UserManagedAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.UserManagedAccess }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) VerifyEmail() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRealmResult) bool { return v.VerifyEmail }).(pulumi.BoolOutput)
}

func (o LookupRealmResultOutput) WebAuthnPasswordlessPolicy() GetRealmWebAuthnPasswordlessPolicyOutput {
	return o.ApplyT(func(v LookupRealmResult) GetRealmWebAuthnPasswordlessPolicy { return v.WebAuthnPasswordlessPolicy }).(GetRealmWebAuthnPasswordlessPolicyOutput)
}

func (o LookupRealmResultOutput) WebAuthnPolicy() GetRealmWebAuthnPolicyOutput {
	return o.ApplyT(func(v LookupRealmResult) GetRealmWebAuthnPolicy { return v.WebAuthnPolicy }).(GetRealmWebAuthnPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRealmResultOutput{})
}
