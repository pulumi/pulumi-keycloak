// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DefaultGroups struct {
	s *pulumi.ResourceState
}

// NewDefaultGroups registers a new resource with the given unique name, arguments, and options.
func NewDefaultGroups(ctx *pulumi.Context,
	name string, args *DefaultGroupsArgs, opts ...pulumi.ResourceOpt) (*DefaultGroups, error) {
	if args == nil || args.GroupIds == nil {
		return nil, errors.New("missing required argument 'GroupIds'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["groupIds"] = nil
		inputs["realmId"] = nil
	} else {
		inputs["groupIds"] = args.GroupIds
		inputs["realmId"] = args.RealmId
	}
	s, err := ctx.RegisterResource("keycloak:index/defaultGroups:DefaultGroups", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultGroups{s: s}, nil
}

// GetDefaultGroups gets an existing DefaultGroups resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultGroups(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DefaultGroupsState, opts ...pulumi.ResourceOpt) (*DefaultGroups, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["groupIds"] = state.GroupIds
		inputs["realmId"] = state.RealmId
	}
	s, err := ctx.ReadResource("keycloak:index/defaultGroups:DefaultGroups", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultGroups{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DefaultGroups) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DefaultGroups) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *DefaultGroups) GroupIds() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["groupIds"])
}

func (r *DefaultGroups) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

// Input properties used for looking up and filtering DefaultGroups resources.
type DefaultGroupsState struct {
	GroupIds interface{}
	RealmId interface{}
}

// The set of arguments for constructing a DefaultGroups resource.
type DefaultGroupsArgs struct {
	GroupIds interface{}
	RealmId interface{}
}
