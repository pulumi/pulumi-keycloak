// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Role struct {
	s *pulumi.ResourceState
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOpt) (*Role, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["clientId"] = nil
		inputs["compositeRoles"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["realmId"] = nil
	} else {
		inputs["clientId"] = args.ClientId
		inputs["compositeRoles"] = args.CompositeRoles
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["realmId"] = args.RealmId
	}
	s, err := ctx.RegisterResource("keycloak:index/role:Role", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Role{s: s}, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RoleState, opts ...pulumi.ResourceOpt) (*Role, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["clientId"] = state.ClientId
		inputs["compositeRoles"] = state.CompositeRoles
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["realmId"] = state.RealmId
	}
	s, err := ctx.ReadResource("keycloak:index/role:Role", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Role{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Role) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Role) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *Role) ClientId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientId"])
}

func (r *Role) CompositeRoles() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["compositeRoles"])
}

func (r *Role) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

func (r *Role) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *Role) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

// Input properties used for looking up and filtering Role resources.
type RoleState struct {
	ClientId interface{}
	CompositeRoles interface{}
	Description interface{}
	Name interface{}
	RealmId interface{}
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	ClientId interface{}
	CompositeRoles interface{}
	Description interface{}
	Name interface{}
	RealmId interface{}
}
