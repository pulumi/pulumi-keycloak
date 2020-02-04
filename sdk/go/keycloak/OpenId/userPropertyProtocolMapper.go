// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package OpenId

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type UserPropertyProtocolMapper struct {
	s *pulumi.ResourceState
}

// NewUserPropertyProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserPropertyProtocolMapper(ctx *pulumi.Context,
	name string, args *UserPropertyProtocolMapperArgs, opts ...pulumi.ResourceOpt) (*UserPropertyProtocolMapper, error) {
	if args == nil || args.ClaimName == nil {
		return nil, errors.New("missing required argument 'ClaimName'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.UserProperty == nil {
		return nil, errors.New("missing required argument 'UserProperty'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addToAccessToken"] = nil
		inputs["addToIdToken"] = nil
		inputs["addToUserinfo"] = nil
		inputs["claimName"] = nil
		inputs["claimValueType"] = nil
		inputs["clientId"] = nil
		inputs["clientScopeId"] = nil
		inputs["name"] = nil
		inputs["realmId"] = nil
		inputs["userProperty"] = nil
	} else {
		inputs["addToAccessToken"] = args.AddToAccessToken
		inputs["addToIdToken"] = args.AddToIdToken
		inputs["addToUserinfo"] = args.AddToUserinfo
		inputs["claimName"] = args.ClaimName
		inputs["claimValueType"] = args.ClaimValueType
		inputs["clientId"] = args.ClientId
		inputs["clientScopeId"] = args.ClientScopeId
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
		inputs["userProperty"] = args.UserProperty
	}
	s, err := ctx.RegisterResource("keycloak:OpenId/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserPropertyProtocolMapper{s: s}, nil
}

// GetUserPropertyProtocolMapper gets an existing UserPropertyProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPropertyProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserPropertyProtocolMapperState, opts ...pulumi.ResourceOpt) (*UserPropertyProtocolMapper, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addToAccessToken"] = state.AddToAccessToken
		inputs["addToIdToken"] = state.AddToIdToken
		inputs["addToUserinfo"] = state.AddToUserinfo
		inputs["claimName"] = state.ClaimName
		inputs["claimValueType"] = state.ClaimValueType
		inputs["clientId"] = state.ClientId
		inputs["clientScopeId"] = state.ClientScopeId
		inputs["name"] = state.Name
		inputs["realmId"] = state.RealmId
		inputs["userProperty"] = state.UserProperty
	}
	s, err := ctx.ReadResource("keycloak:OpenId/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserPropertyProtocolMapper{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *UserPropertyProtocolMapper) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *UserPropertyProtocolMapper) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Indicates if the property should be a claim in the access token.
func (r *UserPropertyProtocolMapper) AddToAccessToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToAccessToken"])
}

// Indicates if the property should be a claim in the id token.
func (r *UserPropertyProtocolMapper) AddToIdToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToIdToken"])
}

// Indicates if the property should appear in the userinfo response body.
func (r *UserPropertyProtocolMapper) AddToUserinfo() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToUserinfo"])
}

func (r *UserPropertyProtocolMapper) ClaimName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["claimName"])
}

// Claim type used when serializing tokens.
func (r *UserPropertyProtocolMapper) ClaimValueType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["claimValueType"])
}

// The mapper's associated client. Cannot be used at the same time as client_scope_id.
func (r *UserPropertyProtocolMapper) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

// The mapper's associated client scope. Cannot be used at the same time as client_id.
func (r *UserPropertyProtocolMapper) ClientScopeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientScopeId"])
}

// A human-friendly name that will appear in the Keycloak console.
func (r *UserPropertyProtocolMapper) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The realm id where the associated client or client scope exists.
func (r *UserPropertyProtocolMapper) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

func (r *UserPropertyProtocolMapper) UserProperty() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userProperty"])
}

// Input properties used for looking up and filtering UserPropertyProtocolMapper resources.
type UserPropertyProtocolMapperState struct {
	// Indicates if the property should be a claim in the access token.
	AddToAccessToken interface{}
	// Indicates if the property should be a claim in the id token.
	AddToIdToken interface{}
	// Indicates if the property should appear in the userinfo response body.
	AddToUserinfo interface{}
	ClaimName interface{}
	// Claim type used when serializing tokens.
	ClaimValueType interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
	UserProperty interface{}
}

// The set of arguments for constructing a UserPropertyProtocolMapper resource.
type UserPropertyProtocolMapperArgs struct {
	// Indicates if the property should be a claim in the access token.
	AddToAccessToken interface{}
	// Indicates if the property should be a claim in the id token.
	AddToIdToken interface{}
	// Indicates if the property should appear in the userinfo response body.
	AddToUserinfo interface{}
	ClaimName interface{}
	// Claim type used when serializing tokens.
	ClaimValueType interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
	UserProperty interface{}
}
