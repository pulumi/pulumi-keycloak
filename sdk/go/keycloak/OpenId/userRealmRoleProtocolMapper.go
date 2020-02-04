// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package OpenId

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type UserRealmRoleProtocolMapper struct {
	s *pulumi.ResourceState
}

// NewUserRealmRoleProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserRealmRoleProtocolMapper(ctx *pulumi.Context,
	name string, args *UserRealmRoleProtocolMapperArgs, opts ...pulumi.ResourceOpt) (*UserRealmRoleProtocolMapper, error) {
	if args == nil || args.ClaimName == nil {
		return nil, errors.New("missing required argument 'ClaimName'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
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
		inputs["multivalued"] = nil
		inputs["name"] = nil
		inputs["realmId"] = nil
		inputs["realmRolePrefix"] = nil
	} else {
		inputs["addToAccessToken"] = args.AddToAccessToken
		inputs["addToIdToken"] = args.AddToIdToken
		inputs["addToUserinfo"] = args.AddToUserinfo
		inputs["claimName"] = args.ClaimName
		inputs["claimValueType"] = args.ClaimValueType
		inputs["clientId"] = args.ClientId
		inputs["clientScopeId"] = args.ClientScopeId
		inputs["multivalued"] = args.Multivalued
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
		inputs["realmRolePrefix"] = args.RealmRolePrefix
	}
	s, err := ctx.RegisterResource("keycloak:OpenId/userRealmRoleProtocolMapper:UserRealmRoleProtocolMapper", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserRealmRoleProtocolMapper{s: s}, nil
}

// GetUserRealmRoleProtocolMapper gets an existing UserRealmRoleProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRealmRoleProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserRealmRoleProtocolMapperState, opts ...pulumi.ResourceOpt) (*UserRealmRoleProtocolMapper, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addToAccessToken"] = state.AddToAccessToken
		inputs["addToIdToken"] = state.AddToIdToken
		inputs["addToUserinfo"] = state.AddToUserinfo
		inputs["claimName"] = state.ClaimName
		inputs["claimValueType"] = state.ClaimValueType
		inputs["clientId"] = state.ClientId
		inputs["clientScopeId"] = state.ClientScopeId
		inputs["multivalued"] = state.Multivalued
		inputs["name"] = state.Name
		inputs["realmId"] = state.RealmId
		inputs["realmRolePrefix"] = state.RealmRolePrefix
	}
	s, err := ctx.ReadResource("keycloak:OpenId/userRealmRoleProtocolMapper:UserRealmRoleProtocolMapper", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &UserRealmRoleProtocolMapper{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *UserRealmRoleProtocolMapper) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *UserRealmRoleProtocolMapper) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Indicates if the attribute should be a claim in the access token.
func (r *UserRealmRoleProtocolMapper) AddToAccessToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToAccessToken"])
}

// Indicates if the attribute should be a claim in the id token.
func (r *UserRealmRoleProtocolMapper) AddToIdToken() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToIdToken"])
}

// Indicates if the attribute should appear in the userinfo response body.
func (r *UserRealmRoleProtocolMapper) AddToUserinfo() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["addToUserinfo"])
}

func (r *UserRealmRoleProtocolMapper) ClaimName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["claimName"])
}

// Claim type used when serializing tokens.
func (r *UserRealmRoleProtocolMapper) ClaimValueType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["claimValueType"])
}

// The mapper's associated client. Cannot be used at the same time as client_scope_id.
func (r *UserRealmRoleProtocolMapper) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

// The mapper's associated client scope. Cannot be used at the same time as client_id.
func (r *UserRealmRoleProtocolMapper) ClientScopeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientScopeId"])
}

// Indicates whether this attribute is a single value or an array of values.
func (r *UserRealmRoleProtocolMapper) Multivalued() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["multivalued"])
}

// A human-friendly name that will appear in the Keycloak console.
func (r *UserRealmRoleProtocolMapper) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The realm id where the associated client or client scope exists.
func (r *UserRealmRoleProtocolMapper) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

// Prefix that will be added to each realm role.
func (r *UserRealmRoleProtocolMapper) RealmRolePrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmRolePrefix"])
}

// Input properties used for looking up and filtering UserRealmRoleProtocolMapper resources.
type UserRealmRoleProtocolMapperState struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken interface{}
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken interface{}
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo interface{}
	ClaimName interface{}
	// Claim type used when serializing tokens.
	ClaimValueType interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
	// Prefix that will be added to each realm role.
	RealmRolePrefix interface{}
}

// The set of arguments for constructing a UserRealmRoleProtocolMapper resource.
type UserRealmRoleProtocolMapperArgs struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken interface{}
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken interface{}
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo interface{}
	ClaimName interface{}
	// Claim type used when serializing tokens.
	ClaimValueType interface{}
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId interface{}
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId interface{}
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued interface{}
	// A human-friendly name that will appear in the Keycloak console.
	Name interface{}
	// The realm id where the associated client or client scope exists.
	RealmId interface{}
	// Prefix that will be added to each realm role.
	RealmRolePrefix interface{}
}
