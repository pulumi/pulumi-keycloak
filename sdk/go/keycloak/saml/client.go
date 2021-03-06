// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing Keycloak clients that use the SAML protocol.
//
// Clients are entities that can use Keycloak for user authentication. Typically, clients are applications that redirect users
// to Keycloak for authentication in order to take advantage of Keycloak's user sessions for SSO.
//
// ## Import
//
// Clients can be imported using the format `{{realm_id}}/{{client_keycloak_id}}`, where `client_keycloak_id` is the unique ID that Keycloak assigns to the client upon creation. This value can be found in the URI when editing this client in the GUI, and is typically a GUID. Examplebash
//
// ```sh
//  $ pulumi import keycloak:saml/client:Client saml_client my-realm/dcbc4c73-e478-4928-ae2e-d5e420223352
// ```
type Client struct {
	pulumi.CustomResourceState

	// SAML POST Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerPostUrl pulumi.StringPtrOutput `pulumi:"assertionConsumerPostUrl"`
	// SAML Redirect Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerRedirectUrl pulumi.StringPtrOutput `pulumi:"assertionConsumerRedirectUrl"`
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides ClientAuthenticationFlowBindingOverridesPtrOutput `pulumi:"authenticationFlowBindingOverrides"`
	// When specified, this URL will be used whenever Keycloak needs to link to this client.
	BaseUrl pulumi.StringPtrOutput `pulumi:"baseUrl"`
	// The unique ID of this client, referenced in the URI during authentication and in issued tokens.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signingCertificate` and `signingPrivateKey`.
	ClientSignatureRequired pulumi.BoolOutput `pulumi:"clientSignatureRequired"`
	// The description of this client in the GUI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key.
	EncryptAssertions pulumi.BoolOutput `pulumi:"encryptAssertions"`
	// If assertions for the client are encrypted, this certificate will be used for encryption.
	EncryptionCertificate pulumi.StringPtrOutput `pulumi:"encryptionCertificate"`
	// Ignore requested NameID subject format and use the one defined in `nameIdFormat` instead.
	ForceNameIdFormat pulumi.BoolOutput `pulumi:"forceNameIdFormat"`
	// When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding.
	ForcePostBinding pulumi.BoolOutput `pulumi:"forcePostBinding"`
	// When `true`, this client will require a browser redirect in order to perform a logout.
	FrontChannelLogout pulumi.BoolOutput `pulumi:"frontChannelLogout"`
	// - Allow to include all roles mappings in the access token
	FullScopeAllowed pulumi.BoolPtrOutput `pulumi:"fullScopeAllowed"`
	// Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
	IdpInitiatedSsoRelayState pulumi.StringPtrOutput `pulumi:"idpInitiatedSsoRelayState"`
	// URL fragment name to reference client when you want to do IDP Initiated SSO.
	IdpInitiatedSsoUrlName pulumi.StringPtrOutput `pulumi:"idpInitiatedSsoUrlName"`
	// When `true`, an `AuthnStatement` will be included in the SAML response.
	IncludeAuthnStatement pulumi.BoolOutput `pulumi:"includeAuthnStatement"`
	// SAML POST Binding URL for the client's single logout service.
	LogoutServicePostBindingUrl pulumi.StringPtrOutput `pulumi:"logoutServicePostBindingUrl"`
	// SAML Redirect Binding URL for the client's single logout service.
	LogoutServiceRedirectBindingUrl pulumi.StringPtrOutput `pulumi:"logoutServiceRedirectBindingUrl"`
	// When specified, this URL will be used for all SAML requests.
	MasterSamlProcessingUrl pulumi.StringPtrOutput `pulumi:"masterSamlProcessingUrl"`
	// The display name of this client in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sets the Name ID format for the subject.
	NameIdFormat pulumi.StringOutput `pulumi:"nameIdFormat"`
	// The realm this client is attached to.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// When specified, this value is prepended to all relative URLs.
	RootUrl pulumi.StringPtrOutput `pulumi:"rootUrl"`
	// When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response.
	SignAssertions pulumi.BoolOutput `pulumi:"signAssertions"`
	// When `true`, the SAML document will be signed by Keycloak using the realm's private key.
	SignDocuments pulumi.BoolOutput `pulumi:"signDocuments"`
	// The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA512", or "DSA_SHA1".
	SignatureAlgorithm pulumi.StringPtrOutput `pulumi:"signatureAlgorithm"`
	// If documents or assertions from the client are signed, this certificate will be used to verify the signature.
	SigningCertificate pulumi.StringPtrOutput `pulumi:"signingCertificate"`
	// If documents or assertions from the client are signed, this private key will be used to verify the signature.
	SigningPrivateKey pulumi.StringPtrOutput `pulumi:"signingPrivateKey"`
	// When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
	ValidRedirectUris pulumi.StringArrayOutput `pulumi:"validRedirectUris"`
}

// NewClient registers a new resource with the given unique name, arguments, and options.
func NewClient(ctx *pulumi.Context,
	name string, args *ClientArgs, opts ...pulumi.ResourceOption) (*Client, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource Client
	err := ctx.RegisterResource("keycloak:saml/client:Client", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClient gets an existing Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientState, opts ...pulumi.ResourceOption) (*Client, error) {
	var resource Client
	err := ctx.ReadResource("keycloak:saml/client:Client", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Client resources.
type clientState struct {
	// SAML POST Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerPostUrl *string `pulumi:"assertionConsumerPostUrl"`
	// SAML Redirect Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerRedirectUrl *string `pulumi:"assertionConsumerRedirectUrl"`
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides *ClientAuthenticationFlowBindingOverrides `pulumi:"authenticationFlowBindingOverrides"`
	// When specified, this URL will be used whenever Keycloak needs to link to this client.
	BaseUrl *string `pulumi:"baseUrl"`
	// The unique ID of this client, referenced in the URI during authentication and in issued tokens.
	ClientId *string `pulumi:"clientId"`
	// When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signingCertificate` and `signingPrivateKey`.
	ClientSignatureRequired *bool `pulumi:"clientSignatureRequired"`
	// The description of this client in the GUI.
	Description *string `pulumi:"description"`
	// When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key.
	EncryptAssertions *bool `pulumi:"encryptAssertions"`
	// If assertions for the client are encrypted, this certificate will be used for encryption.
	EncryptionCertificate *string `pulumi:"encryptionCertificate"`
	// Ignore requested NameID subject format and use the one defined in `nameIdFormat` instead.
	ForceNameIdFormat *bool `pulumi:"forceNameIdFormat"`
	// When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding.
	ForcePostBinding *bool `pulumi:"forcePostBinding"`
	// When `true`, this client will require a browser redirect in order to perform a logout.
	FrontChannelLogout *bool `pulumi:"frontChannelLogout"`
	// - Allow to include all roles mappings in the access token
	FullScopeAllowed *bool `pulumi:"fullScopeAllowed"`
	// Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
	IdpInitiatedSsoRelayState *string `pulumi:"idpInitiatedSsoRelayState"`
	// URL fragment name to reference client when you want to do IDP Initiated SSO.
	IdpInitiatedSsoUrlName *string `pulumi:"idpInitiatedSsoUrlName"`
	// When `true`, an `AuthnStatement` will be included in the SAML response.
	IncludeAuthnStatement *bool `pulumi:"includeAuthnStatement"`
	// SAML POST Binding URL for the client's single logout service.
	LogoutServicePostBindingUrl *string `pulumi:"logoutServicePostBindingUrl"`
	// SAML Redirect Binding URL for the client's single logout service.
	LogoutServiceRedirectBindingUrl *string `pulumi:"logoutServiceRedirectBindingUrl"`
	// When specified, this URL will be used for all SAML requests.
	MasterSamlProcessingUrl *string `pulumi:"masterSamlProcessingUrl"`
	// The display name of this client in the GUI.
	Name *string `pulumi:"name"`
	// Sets the Name ID format for the subject.
	NameIdFormat *string `pulumi:"nameIdFormat"`
	// The realm this client is attached to.
	RealmId *string `pulumi:"realmId"`
	// When specified, this value is prepended to all relative URLs.
	RootUrl *string `pulumi:"rootUrl"`
	// When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response.
	SignAssertions *bool `pulumi:"signAssertions"`
	// When `true`, the SAML document will be signed by Keycloak using the realm's private key.
	SignDocuments *bool `pulumi:"signDocuments"`
	// The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA512", or "DSA_SHA1".
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// If documents or assertions from the client are signed, this certificate will be used to verify the signature.
	SigningCertificate *string `pulumi:"signingCertificate"`
	// If documents or assertions from the client are signed, this private key will be used to verify the signature.
	SigningPrivateKey *string `pulumi:"signingPrivateKey"`
	// When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
	ValidRedirectUris []string `pulumi:"validRedirectUris"`
}

type ClientState struct {
	// SAML POST Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerPostUrl pulumi.StringPtrInput
	// SAML Redirect Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerRedirectUrl pulumi.StringPtrInput
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides ClientAuthenticationFlowBindingOverridesPtrInput
	// When specified, this URL will be used whenever Keycloak needs to link to this client.
	BaseUrl pulumi.StringPtrInput
	// The unique ID of this client, referenced in the URI during authentication and in issued tokens.
	ClientId pulumi.StringPtrInput
	// When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signingCertificate` and `signingPrivateKey`.
	ClientSignatureRequired pulumi.BoolPtrInput
	// The description of this client in the GUI.
	Description pulumi.StringPtrInput
	// When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key.
	EncryptAssertions pulumi.BoolPtrInput
	// If assertions for the client are encrypted, this certificate will be used for encryption.
	EncryptionCertificate pulumi.StringPtrInput
	// Ignore requested NameID subject format and use the one defined in `nameIdFormat` instead.
	ForceNameIdFormat pulumi.BoolPtrInput
	// When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding.
	ForcePostBinding pulumi.BoolPtrInput
	// When `true`, this client will require a browser redirect in order to perform a logout.
	FrontChannelLogout pulumi.BoolPtrInput
	// - Allow to include all roles mappings in the access token
	FullScopeAllowed pulumi.BoolPtrInput
	// Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
	IdpInitiatedSsoRelayState pulumi.StringPtrInput
	// URL fragment name to reference client when you want to do IDP Initiated SSO.
	IdpInitiatedSsoUrlName pulumi.StringPtrInput
	// When `true`, an `AuthnStatement` will be included in the SAML response.
	IncludeAuthnStatement pulumi.BoolPtrInput
	// SAML POST Binding URL for the client's single logout service.
	LogoutServicePostBindingUrl pulumi.StringPtrInput
	// SAML Redirect Binding URL for the client's single logout service.
	LogoutServiceRedirectBindingUrl pulumi.StringPtrInput
	// When specified, this URL will be used for all SAML requests.
	MasterSamlProcessingUrl pulumi.StringPtrInput
	// The display name of this client in the GUI.
	Name pulumi.StringPtrInput
	// Sets the Name ID format for the subject.
	NameIdFormat pulumi.StringPtrInput
	// The realm this client is attached to.
	RealmId pulumi.StringPtrInput
	// When specified, this value is prepended to all relative URLs.
	RootUrl pulumi.StringPtrInput
	// When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response.
	SignAssertions pulumi.BoolPtrInput
	// When `true`, the SAML document will be signed by Keycloak using the realm's private key.
	SignDocuments pulumi.BoolPtrInput
	// The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA512", or "DSA_SHA1".
	SignatureAlgorithm pulumi.StringPtrInput
	// If documents or assertions from the client are signed, this certificate will be used to verify the signature.
	SigningCertificate pulumi.StringPtrInput
	// If documents or assertions from the client are signed, this private key will be used to verify the signature.
	SigningPrivateKey pulumi.StringPtrInput
	// When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
	ValidRedirectUris pulumi.StringArrayInput
}

func (ClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientState)(nil)).Elem()
}

type clientArgs struct {
	// SAML POST Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerPostUrl *string `pulumi:"assertionConsumerPostUrl"`
	// SAML Redirect Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerRedirectUrl *string `pulumi:"assertionConsumerRedirectUrl"`
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides *ClientAuthenticationFlowBindingOverrides `pulumi:"authenticationFlowBindingOverrides"`
	// When specified, this URL will be used whenever Keycloak needs to link to this client.
	BaseUrl *string `pulumi:"baseUrl"`
	// The unique ID of this client, referenced in the URI during authentication and in issued tokens.
	ClientId string `pulumi:"clientId"`
	// When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signingCertificate` and `signingPrivateKey`.
	ClientSignatureRequired *bool `pulumi:"clientSignatureRequired"`
	// The description of this client in the GUI.
	Description *string `pulumi:"description"`
	// When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key.
	EncryptAssertions *bool `pulumi:"encryptAssertions"`
	// If assertions for the client are encrypted, this certificate will be used for encryption.
	EncryptionCertificate *string `pulumi:"encryptionCertificate"`
	// Ignore requested NameID subject format and use the one defined in `nameIdFormat` instead.
	ForceNameIdFormat *bool `pulumi:"forceNameIdFormat"`
	// When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding.
	ForcePostBinding *bool `pulumi:"forcePostBinding"`
	// When `true`, this client will require a browser redirect in order to perform a logout.
	FrontChannelLogout *bool `pulumi:"frontChannelLogout"`
	// - Allow to include all roles mappings in the access token
	FullScopeAllowed *bool `pulumi:"fullScopeAllowed"`
	// Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
	IdpInitiatedSsoRelayState *string `pulumi:"idpInitiatedSsoRelayState"`
	// URL fragment name to reference client when you want to do IDP Initiated SSO.
	IdpInitiatedSsoUrlName *string `pulumi:"idpInitiatedSsoUrlName"`
	// When `true`, an `AuthnStatement` will be included in the SAML response.
	IncludeAuthnStatement *bool `pulumi:"includeAuthnStatement"`
	// SAML POST Binding URL for the client's single logout service.
	LogoutServicePostBindingUrl *string `pulumi:"logoutServicePostBindingUrl"`
	// SAML Redirect Binding URL for the client's single logout service.
	LogoutServiceRedirectBindingUrl *string `pulumi:"logoutServiceRedirectBindingUrl"`
	// When specified, this URL will be used for all SAML requests.
	MasterSamlProcessingUrl *string `pulumi:"masterSamlProcessingUrl"`
	// The display name of this client in the GUI.
	Name *string `pulumi:"name"`
	// Sets the Name ID format for the subject.
	NameIdFormat *string `pulumi:"nameIdFormat"`
	// The realm this client is attached to.
	RealmId string `pulumi:"realmId"`
	// When specified, this value is prepended to all relative URLs.
	RootUrl *string `pulumi:"rootUrl"`
	// When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response.
	SignAssertions *bool `pulumi:"signAssertions"`
	// When `true`, the SAML document will be signed by Keycloak using the realm's private key.
	SignDocuments *bool `pulumi:"signDocuments"`
	// The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA512", or "DSA_SHA1".
	SignatureAlgorithm *string `pulumi:"signatureAlgorithm"`
	// If documents or assertions from the client are signed, this certificate will be used to verify the signature.
	SigningCertificate *string `pulumi:"signingCertificate"`
	// If documents or assertions from the client are signed, this private key will be used to verify the signature.
	SigningPrivateKey *string `pulumi:"signingPrivateKey"`
	// When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
	ValidRedirectUris []string `pulumi:"validRedirectUris"`
}

// The set of arguments for constructing a Client resource.
type ClientArgs struct {
	// SAML POST Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerPostUrl pulumi.StringPtrInput
	// SAML Redirect Binding URL for the client's assertion consumer service (login responses).
	AssertionConsumerRedirectUrl pulumi.StringPtrInput
	// Override realm authentication flow bindings
	AuthenticationFlowBindingOverrides ClientAuthenticationFlowBindingOverridesPtrInput
	// When specified, this URL will be used whenever Keycloak needs to link to this client.
	BaseUrl pulumi.StringPtrInput
	// The unique ID of this client, referenced in the URI during authentication and in issued tokens.
	ClientId pulumi.StringInput
	// When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signingCertificate` and `signingPrivateKey`.
	ClientSignatureRequired pulumi.BoolPtrInput
	// The description of this client in the GUI.
	Description pulumi.StringPtrInput
	// When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key.
	EncryptAssertions pulumi.BoolPtrInput
	// If assertions for the client are encrypted, this certificate will be used for encryption.
	EncryptionCertificate pulumi.StringPtrInput
	// Ignore requested NameID subject format and use the one defined in `nameIdFormat` instead.
	ForceNameIdFormat pulumi.BoolPtrInput
	// When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding.
	ForcePostBinding pulumi.BoolPtrInput
	// When `true`, this client will require a browser redirect in order to perform a logout.
	FrontChannelLogout pulumi.BoolPtrInput
	// - Allow to include all roles mappings in the access token
	FullScopeAllowed pulumi.BoolPtrInput
	// Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
	IdpInitiatedSsoRelayState pulumi.StringPtrInput
	// URL fragment name to reference client when you want to do IDP Initiated SSO.
	IdpInitiatedSsoUrlName pulumi.StringPtrInput
	// When `true`, an `AuthnStatement` will be included in the SAML response.
	IncludeAuthnStatement pulumi.BoolPtrInput
	// SAML POST Binding URL for the client's single logout service.
	LogoutServicePostBindingUrl pulumi.StringPtrInput
	// SAML Redirect Binding URL for the client's single logout service.
	LogoutServiceRedirectBindingUrl pulumi.StringPtrInput
	// When specified, this URL will be used for all SAML requests.
	MasterSamlProcessingUrl pulumi.StringPtrInput
	// The display name of this client in the GUI.
	Name pulumi.StringPtrInput
	// Sets the Name ID format for the subject.
	NameIdFormat pulumi.StringPtrInput
	// The realm this client is attached to.
	RealmId pulumi.StringInput
	// When specified, this value is prepended to all relative URLs.
	RootUrl pulumi.StringPtrInput
	// When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response.
	SignAssertions pulumi.BoolPtrInput
	// When `true`, the SAML document will be signed by Keycloak using the realm's private key.
	SignDocuments pulumi.BoolPtrInput
	// The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA512", or "DSA_SHA1".
	SignatureAlgorithm pulumi.StringPtrInput
	// If documents or assertions from the client are signed, this certificate will be used to verify the signature.
	SigningCertificate pulumi.StringPtrInput
	// If documents or assertions from the client are signed, this private key will be used to verify the signature.
	SigningPrivateKey pulumi.StringPtrInput
	// When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
	ValidRedirectUris pulumi.StringArrayInput
}

func (ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientArgs)(nil)).Elem()
}

type ClientInput interface {
	pulumi.Input

	ToClientOutput() ClientOutput
	ToClientOutputWithContext(ctx context.Context) ClientOutput
}

func (*Client) ElementType() reflect.Type {
	return reflect.TypeOf((*Client)(nil))
}

func (i *Client) ToClientOutput() ClientOutput {
	return i.ToClientOutputWithContext(context.Background())
}

func (i *Client) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOutput)
}

func (i *Client) ToClientPtrOutput() ClientPtrOutput {
	return i.ToClientPtrOutputWithContext(context.Background())
}

func (i *Client) ToClientPtrOutputWithContext(ctx context.Context) ClientPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientPtrOutput)
}

type ClientPtrInput interface {
	pulumi.Input

	ToClientPtrOutput() ClientPtrOutput
	ToClientPtrOutputWithContext(ctx context.Context) ClientPtrOutput
}

type clientPtrType ClientArgs

func (*clientPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Client)(nil))
}

func (i *clientPtrType) ToClientPtrOutput() ClientPtrOutput {
	return i.ToClientPtrOutputWithContext(context.Background())
}

func (i *clientPtrType) ToClientPtrOutputWithContext(ctx context.Context) ClientPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientPtrOutput)
}

// ClientArrayInput is an input type that accepts ClientArray and ClientArrayOutput values.
// You can construct a concrete instance of `ClientArrayInput` via:
//
//          ClientArray{ ClientArgs{...} }
type ClientArrayInput interface {
	pulumi.Input

	ToClientArrayOutput() ClientArrayOutput
	ToClientArrayOutputWithContext(context.Context) ClientArrayOutput
}

type ClientArray []ClientInput

func (ClientArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Client)(nil))
}

func (i ClientArray) ToClientArrayOutput() ClientArrayOutput {
	return i.ToClientArrayOutputWithContext(context.Background())
}

func (i ClientArray) ToClientArrayOutputWithContext(ctx context.Context) ClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientArrayOutput)
}

// ClientMapInput is an input type that accepts ClientMap and ClientMapOutput values.
// You can construct a concrete instance of `ClientMapInput` via:
//
//          ClientMap{ "key": ClientArgs{...} }
type ClientMapInput interface {
	pulumi.Input

	ToClientMapOutput() ClientMapOutput
	ToClientMapOutputWithContext(context.Context) ClientMapOutput
}

type ClientMap map[string]ClientInput

func (ClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Client)(nil))
}

func (i ClientMap) ToClientMapOutput() ClientMapOutput {
	return i.ToClientMapOutputWithContext(context.Background())
}

func (i ClientMap) ToClientMapOutputWithContext(ctx context.Context) ClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientMapOutput)
}

type ClientOutput struct {
	*pulumi.OutputState
}

func (ClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Client)(nil))
}

func (o ClientOutput) ToClientOutput() ClientOutput {
	return o
}

func (o ClientOutput) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return o
}

func (o ClientOutput) ToClientPtrOutput() ClientPtrOutput {
	return o.ToClientPtrOutputWithContext(context.Background())
}

func (o ClientOutput) ToClientPtrOutputWithContext(ctx context.Context) ClientPtrOutput {
	return o.ApplyT(func(v Client) *Client {
		return &v
	}).(ClientPtrOutput)
}

type ClientPtrOutput struct {
	*pulumi.OutputState
}

func (ClientPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Client)(nil))
}

func (o ClientPtrOutput) ToClientPtrOutput() ClientPtrOutput {
	return o
}

func (o ClientPtrOutput) ToClientPtrOutputWithContext(ctx context.Context) ClientPtrOutput {
	return o
}

type ClientArrayOutput struct{ *pulumi.OutputState }

func (ClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Client)(nil))
}

func (o ClientArrayOutput) ToClientArrayOutput() ClientArrayOutput {
	return o
}

func (o ClientArrayOutput) ToClientArrayOutputWithContext(ctx context.Context) ClientArrayOutput {
	return o
}

func (o ClientArrayOutput) Index(i pulumi.IntInput) ClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Client {
		return vs[0].([]Client)[vs[1].(int)]
	}).(ClientOutput)
}

type ClientMapOutput struct{ *pulumi.OutputState }

func (ClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Client)(nil))
}

func (o ClientMapOutput) ToClientMapOutput() ClientMapOutput {
	return o
}

func (o ClientMapOutput) ToClientMapOutputWithContext(ctx context.Context) ClientMapOutput {
	return o
}

func (o ClientMapOutput) MapIndex(k pulumi.StringInput) ClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Client {
		return vs[0].(map[string]Client)[vs[1].(string)]
	}).(ClientOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientOutput{})
	pulumi.RegisterOutputType(ClientPtrOutput{})
	pulumi.RegisterOutputType(ClientArrayOutput{})
	pulumi.RegisterOutputType(ClientMapOutput{})
}
