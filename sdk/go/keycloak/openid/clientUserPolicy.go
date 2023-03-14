// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClientUserPolicy struct {
	pulumi.CustomResourceState

	DecisionStrategy pulumi.StringOutput      `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Logic            pulumi.StringPtrOutput   `pulumi:"logic"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	RealmId          pulumi.StringOutput      `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput      `pulumi:"resourceServerId"`
	Users            pulumi.StringArrayOutput `pulumi:"users"`
}

// NewClientUserPolicy registers a new resource with the given unique name, arguments, and options.
func NewClientUserPolicy(ctx *pulumi.Context,
	name string, args *ClientUserPolicyArgs, opts ...pulumi.ResourceOption) (*ClientUserPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DecisionStrategy == nil {
		return nil, errors.New("invalid value for required argument 'DecisionStrategy'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.ResourceServerId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceServerId'")
	}
	if args.Users == nil {
		return nil, errors.New("invalid value for required argument 'Users'")
	}
	var resource ClientUserPolicy
	err := ctx.RegisterResource("keycloak:openid/clientUserPolicy:ClientUserPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientUserPolicy gets an existing ClientUserPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientUserPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientUserPolicyState, opts ...pulumi.ResourceOption) (*ClientUserPolicy, error) {
	var resource ClientUserPolicy
	err := ctx.ReadResource("keycloak:openid/clientUserPolicy:ClientUserPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientUserPolicy resources.
type clientUserPolicyState struct {
	DecisionStrategy *string  `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Logic            *string  `pulumi:"logic"`
	Name             *string  `pulumi:"name"`
	RealmId          *string  `pulumi:"realmId"`
	ResourceServerId *string  `pulumi:"resourceServerId"`
	Users            []string `pulumi:"users"`
}

type ClientUserPolicyState struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
	Users            pulumi.StringArrayInput
}

func (ClientUserPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientUserPolicyState)(nil)).Elem()
}

type clientUserPolicyArgs struct {
	DecisionStrategy string   `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Logic            *string  `pulumi:"logic"`
	Name             *string  `pulumi:"name"`
	RealmId          string   `pulumi:"realmId"`
	ResourceServerId string   `pulumi:"resourceServerId"`
	Users            []string `pulumi:"users"`
}

// The set of arguments for constructing a ClientUserPolicy resource.
type ClientUserPolicyArgs struct {
	DecisionStrategy pulumi.StringInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
	Users            pulumi.StringArrayInput
}

func (ClientUserPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientUserPolicyArgs)(nil)).Elem()
}

type ClientUserPolicyInput interface {
	pulumi.Input

	ToClientUserPolicyOutput() ClientUserPolicyOutput
	ToClientUserPolicyOutputWithContext(ctx context.Context) ClientUserPolicyOutput
}

func (*ClientUserPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientUserPolicy)(nil)).Elem()
}

func (i *ClientUserPolicy) ToClientUserPolicyOutput() ClientUserPolicyOutput {
	return i.ToClientUserPolicyOutputWithContext(context.Background())
}

func (i *ClientUserPolicy) ToClientUserPolicyOutputWithContext(ctx context.Context) ClientUserPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientUserPolicyOutput)
}

// ClientUserPolicyArrayInput is an input type that accepts ClientUserPolicyArray and ClientUserPolicyArrayOutput values.
// You can construct a concrete instance of `ClientUserPolicyArrayInput` via:
//
//	ClientUserPolicyArray{ ClientUserPolicyArgs{...} }
type ClientUserPolicyArrayInput interface {
	pulumi.Input

	ToClientUserPolicyArrayOutput() ClientUserPolicyArrayOutput
	ToClientUserPolicyArrayOutputWithContext(context.Context) ClientUserPolicyArrayOutput
}

type ClientUserPolicyArray []ClientUserPolicyInput

func (ClientUserPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientUserPolicy)(nil)).Elem()
}

func (i ClientUserPolicyArray) ToClientUserPolicyArrayOutput() ClientUserPolicyArrayOutput {
	return i.ToClientUserPolicyArrayOutputWithContext(context.Background())
}

func (i ClientUserPolicyArray) ToClientUserPolicyArrayOutputWithContext(ctx context.Context) ClientUserPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientUserPolicyArrayOutput)
}

// ClientUserPolicyMapInput is an input type that accepts ClientUserPolicyMap and ClientUserPolicyMapOutput values.
// You can construct a concrete instance of `ClientUserPolicyMapInput` via:
//
//	ClientUserPolicyMap{ "key": ClientUserPolicyArgs{...} }
type ClientUserPolicyMapInput interface {
	pulumi.Input

	ToClientUserPolicyMapOutput() ClientUserPolicyMapOutput
	ToClientUserPolicyMapOutputWithContext(context.Context) ClientUserPolicyMapOutput
}

type ClientUserPolicyMap map[string]ClientUserPolicyInput

func (ClientUserPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientUserPolicy)(nil)).Elem()
}

func (i ClientUserPolicyMap) ToClientUserPolicyMapOutput() ClientUserPolicyMapOutput {
	return i.ToClientUserPolicyMapOutputWithContext(context.Background())
}

func (i ClientUserPolicyMap) ToClientUserPolicyMapOutputWithContext(ctx context.Context) ClientUserPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientUserPolicyMapOutput)
}

type ClientUserPolicyOutput struct{ *pulumi.OutputState }

func (ClientUserPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientUserPolicy)(nil)).Elem()
}

func (o ClientUserPolicyOutput) ToClientUserPolicyOutput() ClientUserPolicyOutput {
	return o
}

func (o ClientUserPolicyOutput) ToClientUserPolicyOutputWithContext(ctx context.Context) ClientUserPolicyOutput {
	return o
}

func (o ClientUserPolicyOutput) DecisionStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientUserPolicy) pulumi.StringOutput { return v.DecisionStrategy }).(pulumi.StringOutput)
}

func (o ClientUserPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientUserPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ClientUserPolicyOutput) Logic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientUserPolicy) pulumi.StringPtrOutput { return v.Logic }).(pulumi.StringPtrOutput)
}

func (o ClientUserPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientUserPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClientUserPolicyOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientUserPolicy) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

func (o ClientUserPolicyOutput) ResourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientUserPolicy) pulumi.StringOutput { return v.ResourceServerId }).(pulumi.StringOutput)
}

func (o ClientUserPolicyOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClientUserPolicy) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

type ClientUserPolicyArrayOutput struct{ *pulumi.OutputState }

func (ClientUserPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientUserPolicy)(nil)).Elem()
}

func (o ClientUserPolicyArrayOutput) ToClientUserPolicyArrayOutput() ClientUserPolicyArrayOutput {
	return o
}

func (o ClientUserPolicyArrayOutput) ToClientUserPolicyArrayOutputWithContext(ctx context.Context) ClientUserPolicyArrayOutput {
	return o
}

func (o ClientUserPolicyArrayOutput) Index(i pulumi.IntInput) ClientUserPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientUserPolicy {
		return vs[0].([]*ClientUserPolicy)[vs[1].(int)]
	}).(ClientUserPolicyOutput)
}

type ClientUserPolicyMapOutput struct{ *pulumi.OutputState }

func (ClientUserPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientUserPolicy)(nil)).Elem()
}

func (o ClientUserPolicyMapOutput) ToClientUserPolicyMapOutput() ClientUserPolicyMapOutput {
	return o
}

func (o ClientUserPolicyMapOutput) ToClientUserPolicyMapOutputWithContext(ctx context.Context) ClientUserPolicyMapOutput {
	return o
}

func (o ClientUserPolicyMapOutput) MapIndex(k pulumi.StringInput) ClientUserPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientUserPolicy {
		return vs[0].(map[string]*ClientUserPolicy)[vs[1].(string)]
	}).(ClientUserPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientUserPolicyInput)(nil)).Elem(), &ClientUserPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientUserPolicyArrayInput)(nil)).Elem(), ClientUserPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientUserPolicyMapInput)(nil)).Elem(), ClientUserPolicyMap{})
	pulumi.RegisterOutputType(ClientUserPolicyOutput{})
	pulumi.RegisterOutputType(ClientUserPolicyArrayOutput{})
	pulumi.RegisterOutputType(ClientUserPolicyMapOutput{})
}
