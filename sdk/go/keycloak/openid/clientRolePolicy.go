// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientRolePolicy struct {
	pulumi.CustomResourceState

	DecisionStrategy pulumi.StringPtrOutput          `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput          `pulumi:"description"`
	Logic            pulumi.StringPtrOutput          `pulumi:"logic"`
	Name             pulumi.StringOutput             `pulumi:"name"`
	RealmId          pulumi.StringOutput             `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput             `pulumi:"resourceServerId"`
	Roles            ClientRolePolicyRoleArrayOutput `pulumi:"roles"`
	Type             pulumi.StringOutput             `pulumi:"type"`
}

// NewClientRolePolicy registers a new resource with the given unique name, arguments, and options.
func NewClientRolePolicy(ctx *pulumi.Context,
	name string, args *ClientRolePolicyArgs, opts ...pulumi.ResourceOption) (*ClientRolePolicy, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.ResourceServerId == nil {
		return nil, errors.New("missing required argument 'ResourceServerId'")
	}
	if args == nil || args.Roles == nil {
		return nil, errors.New("missing required argument 'Roles'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &ClientRolePolicyArgs{}
	}
	var resource ClientRolePolicy
	err := ctx.RegisterResource("keycloak:openid/clientRolePolicy:ClientRolePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientRolePolicy gets an existing ClientRolePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientRolePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientRolePolicyState, opts ...pulumi.ResourceOption) (*ClientRolePolicy, error) {
	var resource ClientRolePolicy
	err := ctx.ReadResource("keycloak:openid/clientRolePolicy:ClientRolePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientRolePolicy resources.
type clientRolePolicyState struct {
	DecisionStrategy *string                `pulumi:"decisionStrategy"`
	Description      *string                `pulumi:"description"`
	Logic            *string                `pulumi:"logic"`
	Name             *string                `pulumi:"name"`
	RealmId          *string                `pulumi:"realmId"`
	ResourceServerId *string                `pulumi:"resourceServerId"`
	Roles            []ClientRolePolicyRole `pulumi:"roles"`
	Type             *string                `pulumi:"type"`
}

type ClientRolePolicyState struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
	Roles            ClientRolePolicyRoleArrayInput
	Type             pulumi.StringPtrInput
}

func (ClientRolePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientRolePolicyState)(nil)).Elem()
}

type clientRolePolicyArgs struct {
	DecisionStrategy *string                `pulumi:"decisionStrategy"`
	Description      *string                `pulumi:"description"`
	Logic            *string                `pulumi:"logic"`
	Name             *string                `pulumi:"name"`
	RealmId          string                 `pulumi:"realmId"`
	ResourceServerId string                 `pulumi:"resourceServerId"`
	Roles            []ClientRolePolicyRole `pulumi:"roles"`
	Type             string                 `pulumi:"type"`
}

// The set of arguments for constructing a ClientRolePolicy resource.
type ClientRolePolicyArgs struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
	Roles            ClientRolePolicyRoleArrayInput
	Type             pulumi.StringInput
}

func (ClientRolePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientRolePolicyArgs)(nil)).Elem()
}
