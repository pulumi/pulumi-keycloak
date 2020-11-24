// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientPolicy struct {
	pulumi.CustomResourceState

	Clients          pulumi.StringArrayOutput `pulumi:"clients"`
	DecisionStrategy pulumi.StringPtrOutput   `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Logic            pulumi.StringPtrOutput   `pulumi:"logic"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	RealmId          pulumi.StringOutput      `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput      `pulumi:"resourceServerId"`
}

// NewClientPolicy registers a new resource with the given unique name, arguments, and options.
func NewClientPolicy(ctx *pulumi.Context,
	name string, args *ClientPolicyArgs, opts ...pulumi.ResourceOption) (*ClientPolicy, error) {
	if args == nil || args.Clients == nil {
		return nil, errors.New("missing required argument 'Clients'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.ResourceServerId == nil {
		return nil, errors.New("missing required argument 'ResourceServerId'")
	}
	if args == nil {
		args = &ClientPolicyArgs{}
	}
	var resource ClientPolicy
	err := ctx.RegisterResource("keycloak:openid/clientPolicy:ClientPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientPolicy gets an existing ClientPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientPolicyState, opts ...pulumi.ResourceOption) (*ClientPolicy, error) {
	var resource ClientPolicy
	err := ctx.ReadResource("keycloak:openid/clientPolicy:ClientPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientPolicy resources.
type clientPolicyState struct {
	Clients          []string `pulumi:"clients"`
	DecisionStrategy *string  `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Logic            *string  `pulumi:"logic"`
	Name             *string  `pulumi:"name"`
	RealmId          *string  `pulumi:"realmId"`
	ResourceServerId *string  `pulumi:"resourceServerId"`
}

type ClientPolicyState struct {
	Clients          pulumi.StringArrayInput
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
}

func (ClientPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientPolicyState)(nil)).Elem()
}

type clientPolicyArgs struct {
	Clients          []string `pulumi:"clients"`
	DecisionStrategy *string  `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Logic            *string  `pulumi:"logic"`
	Name             *string  `pulumi:"name"`
	RealmId          string   `pulumi:"realmId"`
	ResourceServerId string   `pulumi:"resourceServerId"`
}

// The set of arguments for constructing a ClientPolicy resource.
type ClientPolicyArgs struct {
	Clients          pulumi.StringArrayInput
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
}

func (ClientPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientPolicyArgs)(nil)).Elem()
}

type ClientPolicyInput interface {
	pulumi.Input

	ToClientPolicyOutput() ClientPolicyOutput
	ToClientPolicyOutputWithContext(ctx context.Context) ClientPolicyOutput
}

func (ClientPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientPolicy)(nil)).Elem()
}

func (i ClientPolicy) ToClientPolicyOutput() ClientPolicyOutput {
	return i.ToClientPolicyOutputWithContext(context.Background())
}

func (i ClientPolicy) ToClientPolicyOutputWithContext(ctx context.Context) ClientPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientPolicyOutput)
}

type ClientPolicyOutput struct {
	*pulumi.OutputState
}

func (ClientPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientPolicyOutput)(nil)).Elem()
}

func (o ClientPolicyOutput) ToClientPolicyOutput() ClientPolicyOutput {
	return o
}

func (o ClientPolicyOutput) ToClientPolicyOutputWithContext(ctx context.Context) ClientPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClientPolicyOutput{})
}
