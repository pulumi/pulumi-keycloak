// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClientJsPolicy struct {
	pulumi.CustomResourceState

	Code             pulumi.StringOutput    `pulumi:"code"`
	DecisionStrategy pulumi.StringOutput    `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	Logic            pulumi.StringPtrOutput `pulumi:"logic"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	RealmId          pulumi.StringOutput    `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput    `pulumi:"resourceServerId"`
	Type             pulumi.StringPtrOutput `pulumi:"type"`
}

// NewClientJsPolicy registers a new resource with the given unique name, arguments, and options.
func NewClientJsPolicy(ctx *pulumi.Context,
	name string, args *ClientJsPolicyArgs, opts ...pulumi.ResourceOption) (*ClientJsPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Code == nil {
		return nil, errors.New("invalid value for required argument 'Code'")
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClientJsPolicy
	err := ctx.RegisterResource("keycloak:openid/clientJsPolicy:ClientJsPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientJsPolicy gets an existing ClientJsPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientJsPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientJsPolicyState, opts ...pulumi.ResourceOption) (*ClientJsPolicy, error) {
	var resource ClientJsPolicy
	err := ctx.ReadResource("keycloak:openid/clientJsPolicy:ClientJsPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientJsPolicy resources.
type clientJsPolicyState struct {
	Code             *string `pulumi:"code"`
	DecisionStrategy *string `pulumi:"decisionStrategy"`
	Description      *string `pulumi:"description"`
	Logic            *string `pulumi:"logic"`
	Name             *string `pulumi:"name"`
	RealmId          *string `pulumi:"realmId"`
	ResourceServerId *string `pulumi:"resourceServerId"`
	Type             *string `pulumi:"type"`
}

type ClientJsPolicyState struct {
	Code             pulumi.StringPtrInput
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
	Type             pulumi.StringPtrInput
}

func (ClientJsPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientJsPolicyState)(nil)).Elem()
}

type clientJsPolicyArgs struct {
	Code             string  `pulumi:"code"`
	DecisionStrategy string  `pulumi:"decisionStrategy"`
	Description      *string `pulumi:"description"`
	Logic            *string `pulumi:"logic"`
	Name             *string `pulumi:"name"`
	RealmId          string  `pulumi:"realmId"`
	ResourceServerId string  `pulumi:"resourceServerId"`
	Type             *string `pulumi:"type"`
}

// The set of arguments for constructing a ClientJsPolicy resource.
type ClientJsPolicyArgs struct {
	Code             pulumi.StringInput
	DecisionStrategy pulumi.StringInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
	Type             pulumi.StringPtrInput
}

func (ClientJsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientJsPolicyArgs)(nil)).Elem()
}

type ClientJsPolicyInput interface {
	pulumi.Input

	ToClientJsPolicyOutput() ClientJsPolicyOutput
	ToClientJsPolicyOutputWithContext(ctx context.Context) ClientJsPolicyOutput
}

func (*ClientJsPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientJsPolicy)(nil)).Elem()
}

func (i *ClientJsPolicy) ToClientJsPolicyOutput() ClientJsPolicyOutput {
	return i.ToClientJsPolicyOutputWithContext(context.Background())
}

func (i *ClientJsPolicy) ToClientJsPolicyOutputWithContext(ctx context.Context) ClientJsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientJsPolicyOutput)
}

// ClientJsPolicyArrayInput is an input type that accepts ClientJsPolicyArray and ClientJsPolicyArrayOutput values.
// You can construct a concrete instance of `ClientJsPolicyArrayInput` via:
//
//	ClientJsPolicyArray{ ClientJsPolicyArgs{...} }
type ClientJsPolicyArrayInput interface {
	pulumi.Input

	ToClientJsPolicyArrayOutput() ClientJsPolicyArrayOutput
	ToClientJsPolicyArrayOutputWithContext(context.Context) ClientJsPolicyArrayOutput
}

type ClientJsPolicyArray []ClientJsPolicyInput

func (ClientJsPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientJsPolicy)(nil)).Elem()
}

func (i ClientJsPolicyArray) ToClientJsPolicyArrayOutput() ClientJsPolicyArrayOutput {
	return i.ToClientJsPolicyArrayOutputWithContext(context.Background())
}

func (i ClientJsPolicyArray) ToClientJsPolicyArrayOutputWithContext(ctx context.Context) ClientJsPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientJsPolicyArrayOutput)
}

// ClientJsPolicyMapInput is an input type that accepts ClientJsPolicyMap and ClientJsPolicyMapOutput values.
// You can construct a concrete instance of `ClientJsPolicyMapInput` via:
//
//	ClientJsPolicyMap{ "key": ClientJsPolicyArgs{...} }
type ClientJsPolicyMapInput interface {
	pulumi.Input

	ToClientJsPolicyMapOutput() ClientJsPolicyMapOutput
	ToClientJsPolicyMapOutputWithContext(context.Context) ClientJsPolicyMapOutput
}

type ClientJsPolicyMap map[string]ClientJsPolicyInput

func (ClientJsPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientJsPolicy)(nil)).Elem()
}

func (i ClientJsPolicyMap) ToClientJsPolicyMapOutput() ClientJsPolicyMapOutput {
	return i.ToClientJsPolicyMapOutputWithContext(context.Background())
}

func (i ClientJsPolicyMap) ToClientJsPolicyMapOutputWithContext(ctx context.Context) ClientJsPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientJsPolicyMapOutput)
}

type ClientJsPolicyOutput struct{ *pulumi.OutputState }

func (ClientJsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientJsPolicy)(nil)).Elem()
}

func (o ClientJsPolicyOutput) ToClientJsPolicyOutput() ClientJsPolicyOutput {
	return o
}

func (o ClientJsPolicyOutput) ToClientJsPolicyOutputWithContext(ctx context.Context) ClientJsPolicyOutput {
	return o
}

func (o ClientJsPolicyOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientJsPolicy) pulumi.StringOutput { return v.Code }).(pulumi.StringOutput)
}

func (o ClientJsPolicyOutput) DecisionStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientJsPolicy) pulumi.StringOutput { return v.DecisionStrategy }).(pulumi.StringOutput)
}

func (o ClientJsPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientJsPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ClientJsPolicyOutput) Logic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientJsPolicy) pulumi.StringPtrOutput { return v.Logic }).(pulumi.StringPtrOutput)
}

func (o ClientJsPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientJsPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClientJsPolicyOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientJsPolicy) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

func (o ClientJsPolicyOutput) ResourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientJsPolicy) pulumi.StringOutput { return v.ResourceServerId }).(pulumi.StringOutput)
}

func (o ClientJsPolicyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientJsPolicy) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type ClientJsPolicyArrayOutput struct{ *pulumi.OutputState }

func (ClientJsPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientJsPolicy)(nil)).Elem()
}

func (o ClientJsPolicyArrayOutput) ToClientJsPolicyArrayOutput() ClientJsPolicyArrayOutput {
	return o
}

func (o ClientJsPolicyArrayOutput) ToClientJsPolicyArrayOutputWithContext(ctx context.Context) ClientJsPolicyArrayOutput {
	return o
}

func (o ClientJsPolicyArrayOutput) Index(i pulumi.IntInput) ClientJsPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientJsPolicy {
		return vs[0].([]*ClientJsPolicy)[vs[1].(int)]
	}).(ClientJsPolicyOutput)
}

type ClientJsPolicyMapOutput struct{ *pulumi.OutputState }

func (ClientJsPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientJsPolicy)(nil)).Elem()
}

func (o ClientJsPolicyMapOutput) ToClientJsPolicyMapOutput() ClientJsPolicyMapOutput {
	return o
}

func (o ClientJsPolicyMapOutput) ToClientJsPolicyMapOutputWithContext(ctx context.Context) ClientJsPolicyMapOutput {
	return o
}

func (o ClientJsPolicyMapOutput) MapIndex(k pulumi.StringInput) ClientJsPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientJsPolicy {
		return vs[0].(map[string]*ClientJsPolicy)[vs[1].(string)]
	}).(ClientJsPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientJsPolicyInput)(nil)).Elem(), &ClientJsPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientJsPolicyArrayInput)(nil)).Elem(), ClientJsPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientJsPolicyMapInput)(nil)).Elem(), ClientJsPolicyMap{})
	pulumi.RegisterOutputType(ClientJsPolicyOutput{})
	pulumi.RegisterOutputType(ClientJsPolicyArrayOutput{})
	pulumi.RegisterOutputType(ClientJsPolicyMapOutput{})
}
