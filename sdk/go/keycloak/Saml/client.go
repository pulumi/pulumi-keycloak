// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package Saml

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Client struct {
	s *pulumi.ResourceState
}

// NewClient registers a new resource with the given unique name, arguments, and options.
func NewClient(ctx *pulumi.Context,
	name string, args *ClientArgs, opts ...pulumi.ResourceOpt) (*Client, error) {
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["assertionConsumerPostUrl"] = nil
		inputs["assertionConsumerRedirectUrl"] = nil
		inputs["baseUrl"] = nil
		inputs["clientId"] = nil
		inputs["clientSignatureRequired"] = nil
		inputs["description"] = nil
		inputs["enabled"] = nil
		inputs["forcePostBinding"] = nil
		inputs["frontChannelLogout"] = nil
		inputs["fullScopeAllowed"] = nil
		inputs["idpInitiatedSsoRelayState"] = nil
		inputs["idpInitiatedSsoUrlName"] = nil
		inputs["includeAuthnStatement"] = nil
		inputs["logoutServicePostBindingUrl"] = nil
		inputs["logoutServiceRedirectBindingUrl"] = nil
		inputs["masterSamlProcessingUrl"] = nil
		inputs["name"] = nil
		inputs["nameIdFormat"] = nil
		inputs["realmId"] = nil
		inputs["rootUrl"] = nil
		inputs["signAssertions"] = nil
		inputs["signDocuments"] = nil
		inputs["signingCertificate"] = nil
		inputs["signingPrivateKey"] = nil
		inputs["validRedirectUris"] = nil
	} else {
		inputs["assertionConsumerPostUrl"] = args.AssertionConsumerPostUrl
		inputs["assertionConsumerRedirectUrl"] = args.AssertionConsumerRedirectUrl
		inputs["baseUrl"] = args.BaseUrl
		inputs["clientId"] = args.ClientId
		inputs["clientSignatureRequired"] = args.ClientSignatureRequired
		inputs["description"] = args.Description
		inputs["enabled"] = args.Enabled
		inputs["forcePostBinding"] = args.ForcePostBinding
		inputs["frontChannelLogout"] = args.FrontChannelLogout
		inputs["fullScopeAllowed"] = args.FullScopeAllowed
		inputs["idpInitiatedSsoRelayState"] = args.IdpInitiatedSsoRelayState
		inputs["idpInitiatedSsoUrlName"] = args.IdpInitiatedSsoUrlName
		inputs["includeAuthnStatement"] = args.IncludeAuthnStatement
		inputs["logoutServicePostBindingUrl"] = args.LogoutServicePostBindingUrl
		inputs["logoutServiceRedirectBindingUrl"] = args.LogoutServiceRedirectBindingUrl
		inputs["masterSamlProcessingUrl"] = args.MasterSamlProcessingUrl
		inputs["name"] = args.Name
		inputs["nameIdFormat"] = args.NameIdFormat
		inputs["realmId"] = args.RealmId
		inputs["rootUrl"] = args.RootUrl
		inputs["signAssertions"] = args.SignAssertions
		inputs["signDocuments"] = args.SignDocuments
		inputs["signingCertificate"] = args.SigningCertificate
		inputs["signingPrivateKey"] = args.SigningPrivateKey
		inputs["validRedirectUris"] = args.ValidRedirectUris
	}
	s, err := ctx.RegisterResource("keycloak:Saml/client:Client", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{s: s}, nil
}

// GetClient gets an existing Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClient(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClientState, opts ...pulumi.ResourceOpt) (*Client, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["assertionConsumerPostUrl"] = state.AssertionConsumerPostUrl
		inputs["assertionConsumerRedirectUrl"] = state.AssertionConsumerRedirectUrl
		inputs["baseUrl"] = state.BaseUrl
		inputs["clientId"] = state.ClientId
		inputs["clientSignatureRequired"] = state.ClientSignatureRequired
		inputs["description"] = state.Description
		inputs["enabled"] = state.Enabled
		inputs["forcePostBinding"] = state.ForcePostBinding
		inputs["frontChannelLogout"] = state.FrontChannelLogout
		inputs["fullScopeAllowed"] = state.FullScopeAllowed
		inputs["idpInitiatedSsoRelayState"] = state.IdpInitiatedSsoRelayState
		inputs["idpInitiatedSsoUrlName"] = state.IdpInitiatedSsoUrlName
		inputs["includeAuthnStatement"] = state.IncludeAuthnStatement
		inputs["logoutServicePostBindingUrl"] = state.LogoutServicePostBindingUrl
		inputs["logoutServiceRedirectBindingUrl"] = state.LogoutServiceRedirectBindingUrl
		inputs["masterSamlProcessingUrl"] = state.MasterSamlProcessingUrl
		inputs["name"] = state.Name
		inputs["nameIdFormat"] = state.NameIdFormat
		inputs["realmId"] = state.RealmId
		inputs["rootUrl"] = state.RootUrl
		inputs["signAssertions"] = state.SignAssertions
		inputs["signDocuments"] = state.SignDocuments
		inputs["signingCertificate"] = state.SigningCertificate
		inputs["signingPrivateKey"] = state.SigningPrivateKey
		inputs["validRedirectUris"] = state.ValidRedirectUris
	}
	s, err := ctx.ReadResource("keycloak:Saml/client:Client", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Client) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Client) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *Client) AssertionConsumerPostUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["assertionConsumerPostUrl"])
}

func (r *Client) AssertionConsumerRedirectUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["assertionConsumerRedirectUrl"])
}

func (r *Client) BaseUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["baseUrl"])
}

func (r *Client) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

func (r *Client) ClientSignatureRequired() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["clientSignatureRequired"])
}

func (r *Client) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

func (r *Client) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

func (r *Client) ForcePostBinding() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["forcePostBinding"])
}

func (r *Client) FrontChannelLogout() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["frontChannelLogout"])
}

func (r *Client) FullScopeAllowed() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["fullScopeAllowed"])
}

func (r *Client) IdpInitiatedSsoRelayState() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["idpInitiatedSsoRelayState"])
}

func (r *Client) IdpInitiatedSsoUrlName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["idpInitiatedSsoUrlName"])
}

func (r *Client) IncludeAuthnStatement() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["includeAuthnStatement"])
}

func (r *Client) LogoutServicePostBindingUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["logoutServicePostBindingUrl"])
}

func (r *Client) LogoutServiceRedirectBindingUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["logoutServiceRedirectBindingUrl"])
}

func (r *Client) MasterSamlProcessingUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["masterSamlProcessingUrl"])
}

func (r *Client) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *Client) NameIdFormat() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["nameIdFormat"])
}

func (r *Client) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

func (r *Client) RootUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["rootUrl"])
}

func (r *Client) SignAssertions() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["signAssertions"])
}

func (r *Client) SignDocuments() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["signDocuments"])
}

func (r *Client) SigningCertificate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["signingCertificate"])
}

func (r *Client) SigningPrivateKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["signingPrivateKey"])
}

func (r *Client) ValidRedirectUris() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["validRedirectUris"])
}

// Input properties used for looking up and filtering Client resources.
type ClientState struct {
	AssertionConsumerPostUrl interface{}
	AssertionConsumerRedirectUrl interface{}
	BaseUrl interface{}
	ClientId interface{}
	ClientSignatureRequired interface{}
	Description interface{}
	Enabled interface{}
	ForcePostBinding interface{}
	FrontChannelLogout interface{}
	FullScopeAllowed interface{}
	IdpInitiatedSsoRelayState interface{}
	IdpInitiatedSsoUrlName interface{}
	IncludeAuthnStatement interface{}
	LogoutServicePostBindingUrl interface{}
	LogoutServiceRedirectBindingUrl interface{}
	MasterSamlProcessingUrl interface{}
	Name interface{}
	NameIdFormat interface{}
	RealmId interface{}
	RootUrl interface{}
	SignAssertions interface{}
	SignDocuments interface{}
	SigningCertificate interface{}
	SigningPrivateKey interface{}
	ValidRedirectUris interface{}
}

// The set of arguments for constructing a Client resource.
type ClientArgs struct {
	AssertionConsumerPostUrl interface{}
	AssertionConsumerRedirectUrl interface{}
	BaseUrl interface{}
	ClientId interface{}
	ClientSignatureRequired interface{}
	Description interface{}
	Enabled interface{}
	ForcePostBinding interface{}
	FrontChannelLogout interface{}
	FullScopeAllowed interface{}
	IdpInitiatedSsoRelayState interface{}
	IdpInitiatedSsoUrlName interface{}
	IncludeAuthnStatement interface{}
	LogoutServicePostBindingUrl interface{}
	LogoutServiceRedirectBindingUrl interface{}
	MasterSamlProcessingUrl interface{}
	Name interface{}
	NameIdFormat interface{}
	RealmId interface{}
	RootUrl interface{}
	SignAssertions interface{}
	SignDocuments interface{}
	SigningCertificate interface{}
	SigningPrivateKey interface{}
	ValidRedirectUris interface{}
}
