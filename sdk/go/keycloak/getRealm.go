// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.LookupRealm(ctx, &keycloak.LookupRealmArgs{
// 			Realm: "my-realm",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewRole(ctx, "group", &keycloak.RoleArgs{
// 			RealmId: pulumi.String(realm.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRealm(ctx *pulumi.Context, args *LookupRealmArgs, opts ...pulumi.InvokeOption) (*LookupRealmResult, error) {
	var rv LookupRealmResult
	err := ctx.Invoke("keycloak:index/getRealm:getRealm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRealm.
type LookupRealmArgs struct {
	Attributes                  map[string]interface{}         `pulumi:"attributes"`
	DefaultDefaultClientScopes  []string                       `pulumi:"defaultDefaultClientScopes"`
	DefaultOptionalClientScopes []string                       `pulumi:"defaultOptionalClientScopes"`
	DisplayNameHtml             *string                        `pulumi:"displayNameHtml"`
	Internationalizations       []GetRealmInternationalization `pulumi:"internationalizations"`
	// The realm name.
	Realm                      string                              `pulumi:"realm"`
	SecurityDefenses           []GetRealmSecurityDefense           `pulumi:"securityDefenses"`
	SmtpServers                []GetRealmSmtpServer                `pulumi:"smtpServers"`
	WebAuthnPasswordlessPolicy *GetRealmWebAuthnPasswordlessPolicy `pulumi:"webAuthnPasswordlessPolicy"`
	WebAuthnPolicy             *GetRealmWebAuthnPolicy             `pulumi:"webAuthnPolicy"`
}

// A collection of values returned by getRealm.
type LookupRealmResult struct {
	AccessCodeLifespan                  string                 `pulumi:"accessCodeLifespan"`
	AccessCodeLifespanLogin             string                 `pulumi:"accessCodeLifespanLogin"`
	AccessCodeLifespanUserAction        string                 `pulumi:"accessCodeLifespanUserAction"`
	AccessTokenLifespan                 string                 `pulumi:"accessTokenLifespan"`
	AccessTokenLifespanForImplicitFlow  string                 `pulumi:"accessTokenLifespanForImplicitFlow"`
	AccountTheme                        string                 `pulumi:"accountTheme"`
	ActionTokenGeneratedByAdminLifespan string                 `pulumi:"actionTokenGeneratedByAdminLifespan"`
	ActionTokenGeneratedByUserLifespan  string                 `pulumi:"actionTokenGeneratedByUserLifespan"`
	AdminTheme                          string                 `pulumi:"adminTheme"`
	Attributes                          map[string]interface{} `pulumi:"attributes"`
	BrowserFlow                         string                 `pulumi:"browserFlow"`
	ClientAuthenticationFlow            string                 `pulumi:"clientAuthenticationFlow"`
	DefaultDefaultClientScopes          []string               `pulumi:"defaultDefaultClientScopes"`
	DefaultOptionalClientScopes         []string               `pulumi:"defaultOptionalClientScopes"`
	DefaultSignatureAlgorithm           string                 `pulumi:"defaultSignatureAlgorithm"`
	DirectGrantFlow                     string                 `pulumi:"directGrantFlow"`
	DisplayName                         string                 `pulumi:"displayName"`
	DisplayNameHtml                     *string                `pulumi:"displayNameHtml"`
	DockerAuthenticationFlow            string                 `pulumi:"dockerAuthenticationFlow"`
	DuplicateEmailsAllowed              bool                   `pulumi:"duplicateEmailsAllowed"`
	EditUsernameAllowed                 bool                   `pulumi:"editUsernameAllowed"`
	EmailTheme                          string                 `pulumi:"emailTheme"`
	Enabled                             bool                   `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id                               string                             `pulumi:"id"`
	InternalId                       string                             `pulumi:"internalId"`
	Internationalizations            []GetRealmInternationalization     `pulumi:"internationalizations"`
	LoginTheme                       string                             `pulumi:"loginTheme"`
	LoginWithEmailAllowed            bool                               `pulumi:"loginWithEmailAllowed"`
	OfflineSessionIdleTimeout        string                             `pulumi:"offlineSessionIdleTimeout"`
	OfflineSessionMaxLifespan        string                             `pulumi:"offlineSessionMaxLifespan"`
	OfflineSessionMaxLifespanEnabled bool                               `pulumi:"offlineSessionMaxLifespanEnabled"`
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
