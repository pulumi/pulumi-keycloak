// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RealmLocalization struct {
	pulumi.CustomResourceState

	// The locale for the localization texts.
	Locale pulumi.StringOutput `pulumi:"locale"`
	// The realm in which the texts exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The mapping of localization texts keys to values.
	Texts pulumi.StringMapOutput `pulumi:"texts"`
}

// NewRealmLocalization registers a new resource with the given unique name, arguments, and options.
func NewRealmLocalization(ctx *pulumi.Context,
	name string, args *RealmLocalizationArgs, opts ...pulumi.ResourceOption) (*RealmLocalization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Locale == nil {
		return nil, errors.New("invalid value for required argument 'Locale'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RealmLocalization
	err := ctx.RegisterResource("keycloak:index/realmLocalization:RealmLocalization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealmLocalization gets an existing RealmLocalization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealmLocalization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmLocalizationState, opts ...pulumi.ResourceOption) (*RealmLocalization, error) {
	var resource RealmLocalization
	err := ctx.ReadResource("keycloak:index/realmLocalization:RealmLocalization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealmLocalization resources.
type realmLocalizationState struct {
	// The locale for the localization texts.
	Locale *string `pulumi:"locale"`
	// The realm in which the texts exists.
	RealmId *string `pulumi:"realmId"`
	// The mapping of localization texts keys to values.
	Texts map[string]string `pulumi:"texts"`
}

type RealmLocalizationState struct {
	// The locale for the localization texts.
	Locale pulumi.StringPtrInput
	// The realm in which the texts exists.
	RealmId pulumi.StringPtrInput
	// The mapping of localization texts keys to values.
	Texts pulumi.StringMapInput
}

func (RealmLocalizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmLocalizationState)(nil)).Elem()
}

type realmLocalizationArgs struct {
	// The locale for the localization texts.
	Locale string `pulumi:"locale"`
	// The realm in which the texts exists.
	RealmId string `pulumi:"realmId"`
	// The mapping of localization texts keys to values.
	Texts map[string]string `pulumi:"texts"`
}

// The set of arguments for constructing a RealmLocalization resource.
type RealmLocalizationArgs struct {
	// The locale for the localization texts.
	Locale pulumi.StringInput
	// The realm in which the texts exists.
	RealmId pulumi.StringInput
	// The mapping of localization texts keys to values.
	Texts pulumi.StringMapInput
}

func (RealmLocalizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmLocalizationArgs)(nil)).Elem()
}

type RealmLocalizationInput interface {
	pulumi.Input

	ToRealmLocalizationOutput() RealmLocalizationOutput
	ToRealmLocalizationOutputWithContext(ctx context.Context) RealmLocalizationOutput
}

func (*RealmLocalization) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmLocalization)(nil)).Elem()
}

func (i *RealmLocalization) ToRealmLocalizationOutput() RealmLocalizationOutput {
	return i.ToRealmLocalizationOutputWithContext(context.Background())
}

func (i *RealmLocalization) ToRealmLocalizationOutputWithContext(ctx context.Context) RealmLocalizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmLocalizationOutput)
}

// RealmLocalizationArrayInput is an input type that accepts RealmLocalizationArray and RealmLocalizationArrayOutput values.
// You can construct a concrete instance of `RealmLocalizationArrayInput` via:
//
//	RealmLocalizationArray{ RealmLocalizationArgs{...} }
type RealmLocalizationArrayInput interface {
	pulumi.Input

	ToRealmLocalizationArrayOutput() RealmLocalizationArrayOutput
	ToRealmLocalizationArrayOutputWithContext(context.Context) RealmLocalizationArrayOutput
}

type RealmLocalizationArray []RealmLocalizationInput

func (RealmLocalizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmLocalization)(nil)).Elem()
}

func (i RealmLocalizationArray) ToRealmLocalizationArrayOutput() RealmLocalizationArrayOutput {
	return i.ToRealmLocalizationArrayOutputWithContext(context.Background())
}

func (i RealmLocalizationArray) ToRealmLocalizationArrayOutputWithContext(ctx context.Context) RealmLocalizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmLocalizationArrayOutput)
}

// RealmLocalizationMapInput is an input type that accepts RealmLocalizationMap and RealmLocalizationMapOutput values.
// You can construct a concrete instance of `RealmLocalizationMapInput` via:
//
//	RealmLocalizationMap{ "key": RealmLocalizationArgs{...} }
type RealmLocalizationMapInput interface {
	pulumi.Input

	ToRealmLocalizationMapOutput() RealmLocalizationMapOutput
	ToRealmLocalizationMapOutputWithContext(context.Context) RealmLocalizationMapOutput
}

type RealmLocalizationMap map[string]RealmLocalizationInput

func (RealmLocalizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmLocalization)(nil)).Elem()
}

func (i RealmLocalizationMap) ToRealmLocalizationMapOutput() RealmLocalizationMapOutput {
	return i.ToRealmLocalizationMapOutputWithContext(context.Background())
}

func (i RealmLocalizationMap) ToRealmLocalizationMapOutputWithContext(ctx context.Context) RealmLocalizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmLocalizationMapOutput)
}

type RealmLocalizationOutput struct{ *pulumi.OutputState }

func (RealmLocalizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmLocalization)(nil)).Elem()
}

func (o RealmLocalizationOutput) ToRealmLocalizationOutput() RealmLocalizationOutput {
	return o
}

func (o RealmLocalizationOutput) ToRealmLocalizationOutputWithContext(ctx context.Context) RealmLocalizationOutput {
	return o
}

// The locale for the localization texts.
func (o RealmLocalizationOutput) Locale() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmLocalization) pulumi.StringOutput { return v.Locale }).(pulumi.StringOutput)
}

// The realm in which the texts exists.
func (o RealmLocalizationOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmLocalization) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The mapping of localization texts keys to values.
func (o RealmLocalizationOutput) Texts() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RealmLocalization) pulumi.StringMapOutput { return v.Texts }).(pulumi.StringMapOutput)
}

type RealmLocalizationArrayOutput struct{ *pulumi.OutputState }

func (RealmLocalizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmLocalization)(nil)).Elem()
}

func (o RealmLocalizationArrayOutput) ToRealmLocalizationArrayOutput() RealmLocalizationArrayOutput {
	return o
}

func (o RealmLocalizationArrayOutput) ToRealmLocalizationArrayOutputWithContext(ctx context.Context) RealmLocalizationArrayOutput {
	return o
}

func (o RealmLocalizationArrayOutput) Index(i pulumi.IntInput) RealmLocalizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealmLocalization {
		return vs[0].([]*RealmLocalization)[vs[1].(int)]
	}).(RealmLocalizationOutput)
}

type RealmLocalizationMapOutput struct{ *pulumi.OutputState }

func (RealmLocalizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmLocalization)(nil)).Elem()
}

func (o RealmLocalizationMapOutput) ToRealmLocalizationMapOutput() RealmLocalizationMapOutput {
	return o
}

func (o RealmLocalizationMapOutput) ToRealmLocalizationMapOutputWithContext(ctx context.Context) RealmLocalizationMapOutput {
	return o
}

func (o RealmLocalizationMapOutput) MapIndex(k pulumi.StringInput) RealmLocalizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealmLocalization {
		return vs[0].(map[string]*RealmLocalization)[vs[1].(string)]
	}).(RealmLocalizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealmLocalizationInput)(nil)).Elem(), &RealmLocalization{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmLocalizationArrayInput)(nil)).Elem(), RealmLocalizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmLocalizationMapInput)(nil)).Elem(), RealmLocalizationMap{})
	pulumi.RegisterOutputType(RealmLocalizationOutput{})
	pulumi.RegisterOutputType(RealmLocalizationArrayOutput{})
	pulumi.RegisterOutputType(RealmLocalizationMapOutput{})
}
