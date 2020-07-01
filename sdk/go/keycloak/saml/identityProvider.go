// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # saml.IdentityProvider
//
// Allows to create and manage SAML Identity Providers within Keycloak.
//
// SAML (Security Assertion Markup Language) identity providers allows to authenticate through a third-party system, using SAML standard.
//
// ### Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak/saml"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := saml.NewIdentityProvider(ctx, "realmIdentityProvider", &saml.IdentityProviderArgs{
// 			Alias:                   pulumi.String("my-idp"),
// 			BackchannelSupported:    pulumi.Bool(true),
// 			ForceAuthn:              pulumi.Bool(true),
// 			PostBindingAuthnRequest: pulumi.Bool(true),
// 			PostBindingLogout:       pulumi.Bool(true),
// 			PostBindingResponse:     pulumi.Bool(true),
// 			Realm:                   pulumi.String("my-realm"),
// 			SingleLogoutServiceUrl:  pulumi.String("https://domain.com/adfs/ls/?wa=wsignout1.0"),
// 			SingleSignOnServiceUrl:  pulumi.String("https://domain.com/adfs/ls/"),
// 			StoreToken:              pulumi.Bool(false),
// 			TrustEmail:              pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realm` - (Required) The name of the realm. This is unique across Keycloak.
// - `alias` - (Optional) The uniq name of identity provider.
// - `enabled` - (Optional) When false, users and clients will not be able to access this realm. Defaults to `true`.
// - `displayName` - (Optional) The display name for the realm that is shown when logging in to the admin console.
// - `storeToken` - (Optional) Enable/disable if tokens must be stored after authenticating users. Defaults to `true`.
// - `addReadTokenRoleOnCreate` - (Optional) Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role. Defaults to `false`.
// - `trustEmail` - (Optional) If enabled then email provided by this provider is not verified even if verification is enabled for the realm. Defaults to `false`.
// - `linkOnly` - (Optional) If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't want to allow login from the provider, but want to integrate with a provider. Defaults to `false`.
// - `hideOnLoginPage` - (Optional) If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
// - `firstBrokerLoginFlowAlias` - (Optional) Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
// - `postBrokerLoginFlowAlias` - (Optional) Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
// - `authenticateByDefault` - (Optional) Authenticate users by default. Defaults to `false`.
//
// #### SAML Configuration
//
// - `singleSignOnServiceUrl` - (Optional) The Url that must be used to send authentication requests (SAML AuthnRequest).
// - `singleLogoutServiceUrl` - (Optional) The Url that must be used to send logout requests.
// - `backchannelSupported` - (Optional) Does the external IDP support back-channel logout ?.
// - `nameIdPolicyFormat` - (Optional) Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
// - `postBindingResponse` - (Optional) Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
// - `postBindingAuthnRequest` - (Optional) Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
// - `postBindingLogout` - (Optional) Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
// - `wantAssertionsSigned` - (Optional) Indicates whether this service provider expects a signed Assertion.
// - `wantAssertionsEncrypted` - (Optional) Indicates whether this service provider expects an encrypted Assertion.
// - `forceAuthn` - (Optional) Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
// - `validateSignature` - (Optional) Enable/disable signature validation of SAML responses.
// - `signingCertificate` - (Optional) Signing Certificate.
// - `signatureAlgorithm` - (Optional) Signing Algorithm. Defaults to empty.
// - `xmlSignKeyInfoKeyNameTransformer` - (Optional) Sign Key Transformer. Defaults to empty.
type IdentityProvider struct {
	pulumi.CustomResourceState

	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	AddReadTokenRoleOnCreate pulumi.BoolPtrOutput `pulumi:"addReadTokenRoleOnCreate"`
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Enable/disable authenticate users by default.
	AuthenticateByDefault pulumi.BoolPtrOutput `pulumi:"authenticateByDefault"`
	// Does the external IDP support backchannel logout?
	BackchannelSupported pulumi.BoolPtrOutput `pulumi:"backchannelSupported"`
	// Friendly name for Identity Providers.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Enable/disable this identity provider.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
	// that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	FirstBrokerLoginFlowAlias pulumi.StringPtrOutput `pulumi:"firstBrokerLoginFlowAlias"`
	// Require Force Authn.
	ForceAuthn pulumi.BoolPtrOutput `pulumi:"forceAuthn"`
	// Hide On Login Page.
	HideOnLoginPage pulumi.BoolPtrOutput `pulumi:"hideOnLoginPage"`
	// Internal Identity Provider Id
	InternalId pulumi.StringOutput `pulumi:"internalId"`
	// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
	// want to allow login from the provider, but want to integrate with a provider
	LinkOnly pulumi.BoolPtrOutput `pulumi:"linkOnly"`
	// Name ID Policy Format.
	NameIdPolicyFormat pulumi.StringPtrOutput `pulumi:"nameIdPolicyFormat"`
	// Post Binding Authn Request.
	PostBindingAuthnRequest pulumi.BoolPtrOutput `pulumi:"postBindingAuthnRequest"`
	// Post Binding Logout.
	PostBindingLogout pulumi.BoolPtrOutput `pulumi:"postBindingLogout"`
	// Post Binding Response.
	PostBindingResponse pulumi.BoolPtrOutput `pulumi:"postBindingResponse"`
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
	// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
	// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
	// authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
	PostBrokerLoginFlowAlias pulumi.StringPtrOutput `pulumi:"postBrokerLoginFlowAlias"`
	// Realm Name
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Signing Algorithm.
	SignatureAlgorithm pulumi.StringPtrOutput `pulumi:"signatureAlgorithm"`
	// Signing Certificate.
	SigningCertificate pulumi.StringPtrOutput `pulumi:"signingCertificate"`
	// Logout URL.
	SingleLogoutServiceUrl pulumi.StringPtrOutput `pulumi:"singleLogoutServiceUrl"`
	// SSO Logout URL.
	SingleSignOnServiceUrl pulumi.StringOutput `pulumi:"singleSignOnServiceUrl"`
	// Enable/disable if tokens must be stored after authenticating users.
	StoreToken pulumi.BoolPtrOutput `pulumi:"storeToken"`
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail pulumi.BoolPtrOutput `pulumi:"trustEmail"`
	// Enable/disable signature validation of SAML responses.
	ValidateSignature pulumi.BoolPtrOutput `pulumi:"validateSignature"`
	// Want Assertions Encrypted.
	WantAssertionsEncrypted pulumi.BoolPtrOutput `pulumi:"wantAssertionsEncrypted"`
	// Want Assertions Signed.
	WantAssertionsSigned pulumi.BoolPtrOutput `pulumi:"wantAssertionsSigned"`
	// Sign Key Transformer.
	XmlSignKeyInfoKeyNameTransformer pulumi.StringPtrOutput `pulumi:"xmlSignKeyInfoKeyNameTransformer"`
}

// NewIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewIdentityProvider(ctx *pulumi.Context,
	name string, args *IdentityProviderArgs, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	if args == nil || args.Alias == nil {
		return nil, errors.New("missing required argument 'Alias'")
	}
	if args == nil || args.Realm == nil {
		return nil, errors.New("missing required argument 'Realm'")
	}
	if args == nil || args.SingleSignOnServiceUrl == nil {
		return nil, errors.New("missing required argument 'SingleSignOnServiceUrl'")
	}
	if args == nil {
		args = &IdentityProviderArgs{}
	}
	var resource IdentityProvider
	err := ctx.RegisterResource("keycloak:saml/identityProvider:IdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProvider gets an existing IdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderState, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	var resource IdentityProvider
	err := ctx.ReadResource("keycloak:saml/identityProvider:IdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProvider resources.
type identityProviderState struct {
	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	AddReadTokenRoleOnCreate *bool `pulumi:"addReadTokenRoleOnCreate"`
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias *string `pulumi:"alias"`
	// Enable/disable authenticate users by default.
	AuthenticateByDefault *bool `pulumi:"authenticateByDefault"`
	// Does the external IDP support backchannel logout?
	BackchannelSupported *bool `pulumi:"backchannelSupported"`
	// Friendly name for Identity Providers.
	DisplayName *string `pulumi:"displayName"`
	// Enable/disable this identity provider.
	Enabled *bool `pulumi:"enabled"`
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
	// that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	FirstBrokerLoginFlowAlias *string `pulumi:"firstBrokerLoginFlowAlias"`
	// Require Force Authn.
	ForceAuthn *bool `pulumi:"forceAuthn"`
	// Hide On Login Page.
	HideOnLoginPage *bool `pulumi:"hideOnLoginPage"`
	// Internal Identity Provider Id
	InternalId *string `pulumi:"internalId"`
	// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
	// want to allow login from the provider, but want to integrate with a provider
	LinkOnly *bool `pulumi:"linkOnly"`
	// Name ID Policy Format.
	NameIdPolicyFormat *string `pulumi:"nameIdPolicyFormat"`
	// Post Binding Authn Request.
	PostBindingAuthnRequest *bool `pulumi:"postBindingAuthnRequest"`
	// Post Binding Logout.
	PostBindingLogout *bool `pulumi:"postBindingLogout"`
	// Post Binding Response.
	PostBindingResponse *bool `pulumi:"postBindingResponse"`
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
	// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
	// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
	// authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
	PostBrokerLoginFlowAlias *string `pulumi:"postBrokerLoginFlowAlias"`
	// Realm Name
	Realm *string `pulumi:"realm"`
	// Signing Algorithm.
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// Signing Certificate.
	SigningCertificate *string `pulumi:"signingCertificate"`
	// Logout URL.
	SingleLogoutServiceUrl *string `pulumi:"singleLogoutServiceUrl"`
	// SSO Logout URL.
	SingleSignOnServiceUrl *string `pulumi:"singleSignOnServiceUrl"`
	// Enable/disable if tokens must be stored after authenticating users.
	StoreToken *bool `pulumi:"storeToken"`
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail *bool `pulumi:"trustEmail"`
	// Enable/disable signature validation of SAML responses.
	ValidateSignature *bool `pulumi:"validateSignature"`
	// Want Assertions Encrypted.
	WantAssertionsEncrypted *bool `pulumi:"wantAssertionsEncrypted"`
	// Want Assertions Signed.
	WantAssertionsSigned *bool `pulumi:"wantAssertionsSigned"`
	// Sign Key Transformer.
	XmlSignKeyInfoKeyNameTransformer *string `pulumi:"xmlSignKeyInfoKeyNameTransformer"`
}

type IdentityProviderState struct {
	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	AddReadTokenRoleOnCreate pulumi.BoolPtrInput
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias pulumi.StringPtrInput
	// Enable/disable authenticate users by default.
	AuthenticateByDefault pulumi.BoolPtrInput
	// Does the external IDP support backchannel logout?
	BackchannelSupported pulumi.BoolPtrInput
	// Friendly name for Identity Providers.
	DisplayName pulumi.StringPtrInput
	// Enable/disable this identity provider.
	Enabled pulumi.BoolPtrInput
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
	// that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	FirstBrokerLoginFlowAlias pulumi.StringPtrInput
	// Require Force Authn.
	ForceAuthn pulumi.BoolPtrInput
	// Hide On Login Page.
	HideOnLoginPage pulumi.BoolPtrInput
	// Internal Identity Provider Id
	InternalId pulumi.StringPtrInput
	// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
	// want to allow login from the provider, but want to integrate with a provider
	LinkOnly pulumi.BoolPtrInput
	// Name ID Policy Format.
	NameIdPolicyFormat pulumi.StringPtrInput
	// Post Binding Authn Request.
	PostBindingAuthnRequest pulumi.BoolPtrInput
	// Post Binding Logout.
	PostBindingLogout pulumi.BoolPtrInput
	// Post Binding Response.
	PostBindingResponse pulumi.BoolPtrInput
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
	// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
	// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
	// authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
	PostBrokerLoginFlowAlias pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringPtrInput
	// Signing Algorithm.
	SignatureAlgorithm pulumi.StringPtrInput
	// Signing Certificate.
	SigningCertificate pulumi.StringPtrInput
	// Logout URL.
	SingleLogoutServiceUrl pulumi.StringPtrInput
	// SSO Logout URL.
	SingleSignOnServiceUrl pulumi.StringPtrInput
	// Enable/disable if tokens must be stored after authenticating users.
	StoreToken pulumi.BoolPtrInput
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail pulumi.BoolPtrInput
	// Enable/disable signature validation of SAML responses.
	ValidateSignature pulumi.BoolPtrInput
	// Want Assertions Encrypted.
	WantAssertionsEncrypted pulumi.BoolPtrInput
	// Want Assertions Signed.
	WantAssertionsSigned pulumi.BoolPtrInput
	// Sign Key Transformer.
	XmlSignKeyInfoKeyNameTransformer pulumi.StringPtrInput
}

func (IdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderState)(nil)).Elem()
}

type identityProviderArgs struct {
	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	AddReadTokenRoleOnCreate *bool `pulumi:"addReadTokenRoleOnCreate"`
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias string `pulumi:"alias"`
	// Enable/disable authenticate users by default.
	AuthenticateByDefault *bool `pulumi:"authenticateByDefault"`
	// Does the external IDP support backchannel logout?
	BackchannelSupported *bool `pulumi:"backchannelSupported"`
	// Friendly name for Identity Providers.
	DisplayName *string `pulumi:"displayName"`
	// Enable/disable this identity provider.
	Enabled *bool `pulumi:"enabled"`
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
	// that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	FirstBrokerLoginFlowAlias *string `pulumi:"firstBrokerLoginFlowAlias"`
	// Require Force Authn.
	ForceAuthn *bool `pulumi:"forceAuthn"`
	// Hide On Login Page.
	HideOnLoginPage *bool `pulumi:"hideOnLoginPage"`
	// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
	// want to allow login from the provider, but want to integrate with a provider
	LinkOnly *bool `pulumi:"linkOnly"`
	// Name ID Policy Format.
	NameIdPolicyFormat *string `pulumi:"nameIdPolicyFormat"`
	// Post Binding Authn Request.
	PostBindingAuthnRequest *bool `pulumi:"postBindingAuthnRequest"`
	// Post Binding Logout.
	PostBindingLogout *bool `pulumi:"postBindingLogout"`
	// Post Binding Response.
	PostBindingResponse *bool `pulumi:"postBindingResponse"`
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
	// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
	// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
	// authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
	PostBrokerLoginFlowAlias *string `pulumi:"postBrokerLoginFlowAlias"`
	// Realm Name
	Realm string `pulumi:"realm"`
	// Signing Algorithm.
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// Signing Certificate.
	SigningCertificate *string `pulumi:"signingCertificate"`
	// Logout URL.
	SingleLogoutServiceUrl *string `pulumi:"singleLogoutServiceUrl"`
	// SSO Logout URL.
	SingleSignOnServiceUrl string `pulumi:"singleSignOnServiceUrl"`
	// Enable/disable if tokens must be stored after authenticating users.
	StoreToken *bool `pulumi:"storeToken"`
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail *bool `pulumi:"trustEmail"`
	// Enable/disable signature validation of SAML responses.
	ValidateSignature *bool `pulumi:"validateSignature"`
	// Want Assertions Encrypted.
	WantAssertionsEncrypted *bool `pulumi:"wantAssertionsEncrypted"`
	// Want Assertions Signed.
	WantAssertionsSigned *bool `pulumi:"wantAssertionsSigned"`
	// Sign Key Transformer.
	XmlSignKeyInfoKeyNameTransformer *string `pulumi:"xmlSignKeyInfoKeyNameTransformer"`
}

// The set of arguments for constructing a IdentityProvider resource.
type IdentityProviderArgs struct {
	// Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
	AddReadTokenRoleOnCreate pulumi.BoolPtrInput
	// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
	Alias pulumi.StringInput
	// Enable/disable authenticate users by default.
	AuthenticateByDefault pulumi.BoolPtrInput
	// Does the external IDP support backchannel logout?
	BackchannelSupported pulumi.BoolPtrInput
	// Friendly name for Identity Providers.
	DisplayName pulumi.StringPtrInput
	// Enable/disable this identity provider.
	Enabled pulumi.BoolPtrInput
	// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
	// that there is not yet existing Keycloak account linked with the authenticated identity provider account.
	FirstBrokerLoginFlowAlias pulumi.StringPtrInput
	// Require Force Authn.
	ForceAuthn pulumi.BoolPtrInput
	// Hide On Login Page.
	HideOnLoginPage pulumi.BoolPtrInput
	// If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
	// want to allow login from the provider, but want to integrate with a provider
	LinkOnly pulumi.BoolPtrInput
	// Name ID Policy Format.
	NameIdPolicyFormat pulumi.StringPtrInput
	// Post Binding Authn Request.
	PostBindingAuthnRequest pulumi.BoolPtrInput
	// Post Binding Logout.
	PostBindingLogout pulumi.BoolPtrInput
	// Post Binding Response.
	PostBindingResponse pulumi.BoolPtrInput
	// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
	// additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
	// you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
	// authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
	PostBrokerLoginFlowAlias pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringInput
	// Signing Algorithm.
	SignatureAlgorithm pulumi.StringPtrInput
	// Signing Certificate.
	SigningCertificate pulumi.StringPtrInput
	// Logout URL.
	SingleLogoutServiceUrl pulumi.StringPtrInput
	// SSO Logout URL.
	SingleSignOnServiceUrl pulumi.StringInput
	// Enable/disable if tokens must be stored after authenticating users.
	StoreToken pulumi.BoolPtrInput
	// If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail pulumi.BoolPtrInput
	// Enable/disable signature validation of SAML responses.
	ValidateSignature pulumi.BoolPtrInput
	// Want Assertions Encrypted.
	WantAssertionsEncrypted pulumi.BoolPtrInput
	// Want Assertions Signed.
	WantAssertionsSigned pulumi.BoolPtrInput
	// Sign Key Transformer.
	XmlSignKeyInfoKeyNameTransformer pulumi.StringPtrInput
}

func (IdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderArgs)(nil)).Elem()
}
