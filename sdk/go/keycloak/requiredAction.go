// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RequiredAction struct {
	s *pulumi.ResourceState
}

// NewRequiredAction registers a new resource with the given unique name, arguments, and options.
func NewRequiredAction(ctx *pulumi.Context,
	name string, args *RequiredActionArgs, opts ...pulumi.ResourceOpt) (*RequiredAction, error) {
	if args == nil || args.Alias == nil {
		return nil, errors.New("missing required argument 'Alias'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["alias"] = nil
		inputs["defaultAction"] = nil
		inputs["enabled"] = nil
		inputs["name"] = nil
		inputs["priority"] = nil
		inputs["realmId"] = nil
	} else {
		inputs["alias"] = args.Alias
		inputs["defaultAction"] = args.DefaultAction
		inputs["enabled"] = args.Enabled
		inputs["name"] = args.Name
		inputs["priority"] = args.Priority
		inputs["realmId"] = args.RealmId
	}
	s, err := ctx.RegisterResource("keycloak:index/requiredAction:RequiredAction", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RequiredAction{s: s}, nil
}

// GetRequiredAction gets an existing RequiredAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRequiredAction(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RequiredActionState, opts ...pulumi.ResourceOpt) (*RequiredAction, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["alias"] = state.Alias
		inputs["defaultAction"] = state.DefaultAction
		inputs["enabled"] = state.Enabled
		inputs["name"] = state.Name
		inputs["priority"] = state.Priority
		inputs["realmId"] = state.RealmId
	}
	s, err := ctx.ReadResource("keycloak:index/requiredAction:RequiredAction", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RequiredAction{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RequiredAction) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RequiredAction) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *RequiredAction) Alias() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["alias"])
}

func (r *RequiredAction) DefaultAction() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["defaultAction"])
}

func (r *RequiredAction) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

func (r *RequiredAction) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *RequiredAction) Priority() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["priority"])
}

func (r *RequiredAction) RealmId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["realmId"])
}

// Input properties used for looking up and filtering RequiredAction resources.
type RequiredActionState struct {
	Alias interface{}
	DefaultAction interface{}
	Enabled interface{}
	Name interface{}
	Priority interface{}
	RealmId interface{}
}

// The set of arguments for constructing a RequiredAction resource.
type RequiredActionArgs struct {
	Alias interface{}
	DefaultAction interface{}
	Enabled interface{}
	Name interface{}
	Priority interface{}
	RealmId interface{}
}
