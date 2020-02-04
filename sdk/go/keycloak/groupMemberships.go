// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GroupMemberships struct {
	s *pulumi.ResourceState
}

// NewGroupMemberships registers a new resource with the given unique name, arguments, and options.
func NewGroupMemberships(ctx *pulumi.Context,
	name string, args *GroupMembershipsArgs, opts ...pulumi.ResourceOpt) (*GroupMemberships, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["groupId"] = nil
		inputs["members"] = nil
		inputs["realmId"] = nil
	} else {
		inputs["groupId"] = args.GroupId
		inputs["members"] = args.Members
		inputs["realmId"] = args.RealmId
	}
	s, err := ctx.RegisterResource("keycloak:index/groupMemberships:GroupMemberships", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GroupMemberships{s: s}, nil
}

// GetGroupMemberships gets an existing GroupMemberships resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMemberships(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GroupMembershipsState, opts ...pulumi.ResourceOpt) (*GroupMemberships, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["groupId"] = state.GroupId
		inputs["members"] = state.Members
		inputs["realmId"] = state.RealmId
	}
	s, err := ctx.ReadResource("keycloak:index/groupMemberships:GroupMemberships", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GroupMemberships{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *GroupMemberships) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *GroupMemberships) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *GroupMemberships) GroupId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["groupId"])
}

func (r *GroupMemberships) Members() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["members"])
}

func (r *GroupMemberships) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

// Input properties used for looking up and filtering GroupMemberships resources.
type GroupMembershipsState struct {
	GroupId interface{}
	Members interface{}
	RealmId interface{}
}

// The set of arguments for constructing a GroupMemberships resource.
type GroupMembershipsArgs struct {
	GroupId interface{}
	Members interface{}
	RealmId interface{}
}
