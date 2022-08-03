// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClientTimePolicy struct {
	pulumi.CustomResourceState

	DayMonth         pulumi.StringPtrOutput `pulumi:"dayMonth"`
	DayMonthEnd      pulumi.StringPtrOutput `pulumi:"dayMonthEnd"`
	DecisionStrategy pulumi.StringOutput    `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	Hour             pulumi.StringPtrOutput `pulumi:"hour"`
	HourEnd          pulumi.StringPtrOutput `pulumi:"hourEnd"`
	Logic            pulumi.StringPtrOutput `pulumi:"logic"`
	Minute           pulumi.StringPtrOutput `pulumi:"minute"`
	MinuteEnd        pulumi.StringPtrOutput `pulumi:"minuteEnd"`
	Month            pulumi.StringPtrOutput `pulumi:"month"`
	MonthEnd         pulumi.StringPtrOutput `pulumi:"monthEnd"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	NotBefore        pulumi.StringPtrOutput `pulumi:"notBefore"`
	NotOnOrAfter     pulumi.StringPtrOutput `pulumi:"notOnOrAfter"`
	RealmId          pulumi.StringOutput    `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput    `pulumi:"resourceServerId"`
	Year             pulumi.StringPtrOutput `pulumi:"year"`
	YearEnd          pulumi.StringPtrOutput `pulumi:"yearEnd"`
}

// NewClientTimePolicy registers a new resource with the given unique name, arguments, and options.
func NewClientTimePolicy(ctx *pulumi.Context,
	name string, args *ClientTimePolicyArgs, opts ...pulumi.ResourceOption) (*ClientTimePolicy, error) {
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
	var resource ClientTimePolicy
	err := ctx.RegisterResource("keycloak:openid/clientTimePolicy:ClientTimePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientTimePolicy gets an existing ClientTimePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientTimePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientTimePolicyState, opts ...pulumi.ResourceOption) (*ClientTimePolicy, error) {
	var resource ClientTimePolicy
	err := ctx.ReadResource("keycloak:openid/clientTimePolicy:ClientTimePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientTimePolicy resources.
type clientTimePolicyState struct {
	DayMonth         *string `pulumi:"dayMonth"`
	DayMonthEnd      *string `pulumi:"dayMonthEnd"`
	DecisionStrategy *string `pulumi:"decisionStrategy"`
	Description      *string `pulumi:"description"`
	Hour             *string `pulumi:"hour"`
	HourEnd          *string `pulumi:"hourEnd"`
	Logic            *string `pulumi:"logic"`
	Minute           *string `pulumi:"minute"`
	MinuteEnd        *string `pulumi:"minuteEnd"`
	Month            *string `pulumi:"month"`
	MonthEnd         *string `pulumi:"monthEnd"`
	Name             *string `pulumi:"name"`
	NotBefore        *string `pulumi:"notBefore"`
	NotOnOrAfter     *string `pulumi:"notOnOrAfter"`
	RealmId          *string `pulumi:"realmId"`
	ResourceServerId *string `pulumi:"resourceServerId"`
	Year             *string `pulumi:"year"`
	YearEnd          *string `pulumi:"yearEnd"`
}

type ClientTimePolicyState struct {
	DayMonth         pulumi.StringPtrInput
	DayMonthEnd      pulumi.StringPtrInput
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Hour             pulumi.StringPtrInput
	HourEnd          pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Minute           pulumi.StringPtrInput
	MinuteEnd        pulumi.StringPtrInput
	Month            pulumi.StringPtrInput
	MonthEnd         pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	NotBefore        pulumi.StringPtrInput
	NotOnOrAfter     pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
	Year             pulumi.StringPtrInput
	YearEnd          pulumi.StringPtrInput
}

func (ClientTimePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientTimePolicyState)(nil)).Elem()
}

type clientTimePolicyArgs struct {
	DayMonth         *string `pulumi:"dayMonth"`
	DayMonthEnd      *string `pulumi:"dayMonthEnd"`
	DecisionStrategy string  `pulumi:"decisionStrategy"`
	Description      *string `pulumi:"description"`
	Hour             *string `pulumi:"hour"`
	HourEnd          *string `pulumi:"hourEnd"`
	Logic            *string `pulumi:"logic"`
	Minute           *string `pulumi:"minute"`
	MinuteEnd        *string `pulumi:"minuteEnd"`
	Month            *string `pulumi:"month"`
	MonthEnd         *string `pulumi:"monthEnd"`
	Name             *string `pulumi:"name"`
	NotBefore        *string `pulumi:"notBefore"`
	NotOnOrAfter     *string `pulumi:"notOnOrAfter"`
	RealmId          string  `pulumi:"realmId"`
	ResourceServerId string  `pulumi:"resourceServerId"`
	Year             *string `pulumi:"year"`
	YearEnd          *string `pulumi:"yearEnd"`
}

// The set of arguments for constructing a ClientTimePolicy resource.
type ClientTimePolicyArgs struct {
	DayMonth         pulumi.StringPtrInput
	DayMonthEnd      pulumi.StringPtrInput
	DecisionStrategy pulumi.StringInput
	Description      pulumi.StringPtrInput
	Hour             pulumi.StringPtrInput
	HourEnd          pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Minute           pulumi.StringPtrInput
	MinuteEnd        pulumi.StringPtrInput
	Month            pulumi.StringPtrInput
	MonthEnd         pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	NotBefore        pulumi.StringPtrInput
	NotOnOrAfter     pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
	Year             pulumi.StringPtrInput
	YearEnd          pulumi.StringPtrInput
}

func (ClientTimePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientTimePolicyArgs)(nil)).Elem()
}

type ClientTimePolicyInput interface {
	pulumi.Input

	ToClientTimePolicyOutput() ClientTimePolicyOutput
	ToClientTimePolicyOutputWithContext(ctx context.Context) ClientTimePolicyOutput
}

func (*ClientTimePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientTimePolicy)(nil)).Elem()
}

func (i *ClientTimePolicy) ToClientTimePolicyOutput() ClientTimePolicyOutput {
	return i.ToClientTimePolicyOutputWithContext(context.Background())
}

func (i *ClientTimePolicy) ToClientTimePolicyOutputWithContext(ctx context.Context) ClientTimePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientTimePolicyOutput)
}

// ClientTimePolicyArrayInput is an input type that accepts ClientTimePolicyArray and ClientTimePolicyArrayOutput values.
// You can construct a concrete instance of `ClientTimePolicyArrayInput` via:
//
//          ClientTimePolicyArray{ ClientTimePolicyArgs{...} }
type ClientTimePolicyArrayInput interface {
	pulumi.Input

	ToClientTimePolicyArrayOutput() ClientTimePolicyArrayOutput
	ToClientTimePolicyArrayOutputWithContext(context.Context) ClientTimePolicyArrayOutput
}

type ClientTimePolicyArray []ClientTimePolicyInput

func (ClientTimePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientTimePolicy)(nil)).Elem()
}

func (i ClientTimePolicyArray) ToClientTimePolicyArrayOutput() ClientTimePolicyArrayOutput {
	return i.ToClientTimePolicyArrayOutputWithContext(context.Background())
}

func (i ClientTimePolicyArray) ToClientTimePolicyArrayOutputWithContext(ctx context.Context) ClientTimePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientTimePolicyArrayOutput)
}

// ClientTimePolicyMapInput is an input type that accepts ClientTimePolicyMap and ClientTimePolicyMapOutput values.
// You can construct a concrete instance of `ClientTimePolicyMapInput` via:
//
//          ClientTimePolicyMap{ "key": ClientTimePolicyArgs{...} }
type ClientTimePolicyMapInput interface {
	pulumi.Input

	ToClientTimePolicyMapOutput() ClientTimePolicyMapOutput
	ToClientTimePolicyMapOutputWithContext(context.Context) ClientTimePolicyMapOutput
}

type ClientTimePolicyMap map[string]ClientTimePolicyInput

func (ClientTimePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientTimePolicy)(nil)).Elem()
}

func (i ClientTimePolicyMap) ToClientTimePolicyMapOutput() ClientTimePolicyMapOutput {
	return i.ToClientTimePolicyMapOutputWithContext(context.Background())
}

func (i ClientTimePolicyMap) ToClientTimePolicyMapOutputWithContext(ctx context.Context) ClientTimePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientTimePolicyMapOutput)
}

type ClientTimePolicyOutput struct{ *pulumi.OutputState }

func (ClientTimePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientTimePolicy)(nil)).Elem()
}

func (o ClientTimePolicyOutput) ToClientTimePolicyOutput() ClientTimePolicyOutput {
	return o
}

func (o ClientTimePolicyOutput) ToClientTimePolicyOutputWithContext(ctx context.Context) ClientTimePolicyOutput {
	return o
}

func (o ClientTimePolicyOutput) DayMonth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.DayMonth }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) DayMonthEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.DayMonthEnd }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) DecisionStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringOutput { return v.DecisionStrategy }).(pulumi.StringOutput)
}

func (o ClientTimePolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) Hour() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.Hour }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) HourEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.HourEnd }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) Logic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.Logic }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) Minute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.Minute }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) MinuteEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.MinuteEnd }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) Month() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.Month }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) MonthEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.MonthEnd }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClientTimePolicyOutput) NotBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.NotBefore }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) NotOnOrAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.NotOnOrAfter }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

func (o ClientTimePolicyOutput) ResourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringOutput { return v.ResourceServerId }).(pulumi.StringOutput)
}

func (o ClientTimePolicyOutput) Year() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.Year }).(pulumi.StringPtrOutput)
}

func (o ClientTimePolicyOutput) YearEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientTimePolicy) pulumi.StringPtrOutput { return v.YearEnd }).(pulumi.StringPtrOutput)
}

type ClientTimePolicyArrayOutput struct{ *pulumi.OutputState }

func (ClientTimePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientTimePolicy)(nil)).Elem()
}

func (o ClientTimePolicyArrayOutput) ToClientTimePolicyArrayOutput() ClientTimePolicyArrayOutput {
	return o
}

func (o ClientTimePolicyArrayOutput) ToClientTimePolicyArrayOutputWithContext(ctx context.Context) ClientTimePolicyArrayOutput {
	return o
}

func (o ClientTimePolicyArrayOutput) Index(i pulumi.IntInput) ClientTimePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientTimePolicy {
		return vs[0].([]*ClientTimePolicy)[vs[1].(int)]
	}).(ClientTimePolicyOutput)
}

type ClientTimePolicyMapOutput struct{ *pulumi.OutputState }

func (ClientTimePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientTimePolicy)(nil)).Elem()
}

func (o ClientTimePolicyMapOutput) ToClientTimePolicyMapOutput() ClientTimePolicyMapOutput {
	return o
}

func (o ClientTimePolicyMapOutput) ToClientTimePolicyMapOutputWithContext(ctx context.Context) ClientTimePolicyMapOutput {
	return o
}

func (o ClientTimePolicyMapOutput) MapIndex(k pulumi.StringInput) ClientTimePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientTimePolicy {
		return vs[0].(map[string]*ClientTimePolicy)[vs[1].(string)]
	}).(ClientTimePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientTimePolicyInput)(nil)).Elem(), &ClientTimePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientTimePolicyArrayInput)(nil)).Elem(), ClientTimePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientTimePolicyMapInput)(nil)).Elem(), ClientTimePolicyMap{})
	pulumi.RegisterOutputType(ClientTimePolicyOutput{})
	pulumi.RegisterOutputType(ClientTimePolicyArrayOutput{})
	pulumi.RegisterOutputType(ClientTimePolicyMapOutput{})
}
