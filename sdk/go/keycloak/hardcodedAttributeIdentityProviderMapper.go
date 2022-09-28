// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HardcodedAttributeIdentityProviderMapper struct {
	pulumi.CustomResourceState

	// OIDC Claim
	AttributeName pulumi.StringPtrOutput `pulumi:"attributeName"`
	// User Attribute
	AttributeValue pulumi.StringPtrOutput `pulumi:"attributeValue"`
	ExtraConfig    pulumi.MapOutput       `pulumi:"extraConfig"`
	// IDP Alias
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Realm Name
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Is Attribute Related To a User Session
	UserSession pulumi.BoolOutput `pulumi:"userSession"`
}

// NewHardcodedAttributeIdentityProviderMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedAttributeIdentityProviderMapper(ctx *pulumi.Context,
	name string, args *HardcodedAttributeIdentityProviderMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedAttributeIdentityProviderMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityProviderAlias == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderAlias'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	if args.UserSession == nil {
		return nil, errors.New("invalid value for required argument 'UserSession'")
	}
	var resource HardcodedAttributeIdentityProviderMapper
	err := ctx.RegisterResource("keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedAttributeIdentityProviderMapper gets an existing HardcodedAttributeIdentityProviderMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedAttributeIdentityProviderMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedAttributeIdentityProviderMapperState, opts ...pulumi.ResourceOption) (*HardcodedAttributeIdentityProviderMapper, error) {
	var resource HardcodedAttributeIdentityProviderMapper
	err := ctx.ReadResource("keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedAttributeIdentityProviderMapper resources.
type hardcodedAttributeIdentityProviderMapperState struct {
	// OIDC Claim
	AttributeName *string `pulumi:"attributeName"`
	// User Attribute
	AttributeValue *string                `pulumi:"attributeValue"`
	ExtraConfig    map[string]interface{} `pulumi:"extraConfig"`
	// IDP Alias
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm *string `pulumi:"realm"`
	// Is Attribute Related To a User Session
	UserSession *bool `pulumi:"userSession"`
}

type HardcodedAttributeIdentityProviderMapperState struct {
	// OIDC Claim
	AttributeName pulumi.StringPtrInput
	// User Attribute
	AttributeValue pulumi.StringPtrInput
	ExtraConfig    pulumi.MapInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringPtrInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringPtrInput
	// Is Attribute Related To a User Session
	UserSession pulumi.BoolPtrInput
}

func (HardcodedAttributeIdentityProviderMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedAttributeIdentityProviderMapperState)(nil)).Elem()
}

type hardcodedAttributeIdentityProviderMapperArgs struct {
	// OIDC Claim
	AttributeName *string `pulumi:"attributeName"`
	// User Attribute
	AttributeValue *string                `pulumi:"attributeValue"`
	ExtraConfig    map[string]interface{} `pulumi:"extraConfig"`
	// IDP Alias
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm string `pulumi:"realm"`
	// Is Attribute Related To a User Session
	UserSession bool `pulumi:"userSession"`
}

// The set of arguments for constructing a HardcodedAttributeIdentityProviderMapper resource.
type HardcodedAttributeIdentityProviderMapperArgs struct {
	// OIDC Claim
	AttributeName pulumi.StringPtrInput
	// User Attribute
	AttributeValue pulumi.StringPtrInput
	ExtraConfig    pulumi.MapInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringInput
	// Is Attribute Related To a User Session
	UserSession pulumi.BoolInput
}

func (HardcodedAttributeIdentityProviderMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedAttributeIdentityProviderMapperArgs)(nil)).Elem()
}

type HardcodedAttributeIdentityProviderMapperInput interface {
	pulumi.Input

	ToHardcodedAttributeIdentityProviderMapperOutput() HardcodedAttributeIdentityProviderMapperOutput
	ToHardcodedAttributeIdentityProviderMapperOutputWithContext(ctx context.Context) HardcodedAttributeIdentityProviderMapperOutput
}

func (*HardcodedAttributeIdentityProviderMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedAttributeIdentityProviderMapper)(nil)).Elem()
}

func (i *HardcodedAttributeIdentityProviderMapper) ToHardcodedAttributeIdentityProviderMapperOutput() HardcodedAttributeIdentityProviderMapperOutput {
	return i.ToHardcodedAttributeIdentityProviderMapperOutputWithContext(context.Background())
}

func (i *HardcodedAttributeIdentityProviderMapper) ToHardcodedAttributeIdentityProviderMapperOutputWithContext(ctx context.Context) HardcodedAttributeIdentityProviderMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedAttributeIdentityProviderMapperOutput)
}

// HardcodedAttributeIdentityProviderMapperArrayInput is an input type that accepts HardcodedAttributeIdentityProviderMapperArray and HardcodedAttributeIdentityProviderMapperArrayOutput values.
// You can construct a concrete instance of `HardcodedAttributeIdentityProviderMapperArrayInput` via:
//
//	HardcodedAttributeIdentityProviderMapperArray{ HardcodedAttributeIdentityProviderMapperArgs{...} }
type HardcodedAttributeIdentityProviderMapperArrayInput interface {
	pulumi.Input

	ToHardcodedAttributeIdentityProviderMapperArrayOutput() HardcodedAttributeIdentityProviderMapperArrayOutput
	ToHardcodedAttributeIdentityProviderMapperArrayOutputWithContext(context.Context) HardcodedAttributeIdentityProviderMapperArrayOutput
}

type HardcodedAttributeIdentityProviderMapperArray []HardcodedAttributeIdentityProviderMapperInput

func (HardcodedAttributeIdentityProviderMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedAttributeIdentityProviderMapper)(nil)).Elem()
}

func (i HardcodedAttributeIdentityProviderMapperArray) ToHardcodedAttributeIdentityProviderMapperArrayOutput() HardcodedAttributeIdentityProviderMapperArrayOutput {
	return i.ToHardcodedAttributeIdentityProviderMapperArrayOutputWithContext(context.Background())
}

func (i HardcodedAttributeIdentityProviderMapperArray) ToHardcodedAttributeIdentityProviderMapperArrayOutputWithContext(ctx context.Context) HardcodedAttributeIdentityProviderMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedAttributeIdentityProviderMapperArrayOutput)
}

// HardcodedAttributeIdentityProviderMapperMapInput is an input type that accepts HardcodedAttributeIdentityProviderMapperMap and HardcodedAttributeIdentityProviderMapperMapOutput values.
// You can construct a concrete instance of `HardcodedAttributeIdentityProviderMapperMapInput` via:
//
//	HardcodedAttributeIdentityProviderMapperMap{ "key": HardcodedAttributeIdentityProviderMapperArgs{...} }
type HardcodedAttributeIdentityProviderMapperMapInput interface {
	pulumi.Input

	ToHardcodedAttributeIdentityProviderMapperMapOutput() HardcodedAttributeIdentityProviderMapperMapOutput
	ToHardcodedAttributeIdentityProviderMapperMapOutputWithContext(context.Context) HardcodedAttributeIdentityProviderMapperMapOutput
}

type HardcodedAttributeIdentityProviderMapperMap map[string]HardcodedAttributeIdentityProviderMapperInput

func (HardcodedAttributeIdentityProviderMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedAttributeIdentityProviderMapper)(nil)).Elem()
}

func (i HardcodedAttributeIdentityProviderMapperMap) ToHardcodedAttributeIdentityProviderMapperMapOutput() HardcodedAttributeIdentityProviderMapperMapOutput {
	return i.ToHardcodedAttributeIdentityProviderMapperMapOutputWithContext(context.Background())
}

func (i HardcodedAttributeIdentityProviderMapperMap) ToHardcodedAttributeIdentityProviderMapperMapOutputWithContext(ctx context.Context) HardcodedAttributeIdentityProviderMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedAttributeIdentityProviderMapperMapOutput)
}

type HardcodedAttributeIdentityProviderMapperOutput struct{ *pulumi.OutputState }

func (HardcodedAttributeIdentityProviderMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedAttributeIdentityProviderMapper)(nil)).Elem()
}

func (o HardcodedAttributeIdentityProviderMapperOutput) ToHardcodedAttributeIdentityProviderMapperOutput() HardcodedAttributeIdentityProviderMapperOutput {
	return o
}

func (o HardcodedAttributeIdentityProviderMapperOutput) ToHardcodedAttributeIdentityProviderMapperOutputWithContext(ctx context.Context) HardcodedAttributeIdentityProviderMapperOutput {
	return o
}

// OIDC Claim
func (o HardcodedAttributeIdentityProviderMapperOutput) AttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardcodedAttributeIdentityProviderMapper) pulumi.StringPtrOutput { return v.AttributeName }).(pulumi.StringPtrOutput)
}

// User Attribute
func (o HardcodedAttributeIdentityProviderMapperOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardcodedAttributeIdentityProviderMapper) pulumi.StringPtrOutput { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

func (o HardcodedAttributeIdentityProviderMapperOutput) ExtraConfig() pulumi.MapOutput {
	return o.ApplyT(func(v *HardcodedAttributeIdentityProviderMapper) pulumi.MapOutput { return v.ExtraConfig }).(pulumi.MapOutput)
}

// IDP Alias
func (o HardcodedAttributeIdentityProviderMapperOutput) IdentityProviderAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedAttributeIdentityProviderMapper) pulumi.StringOutput { return v.IdentityProviderAlias }).(pulumi.StringOutput)
}

// IDP Mapper Name
func (o HardcodedAttributeIdentityProviderMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedAttributeIdentityProviderMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Realm Name
func (o HardcodedAttributeIdentityProviderMapperOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedAttributeIdentityProviderMapper) pulumi.StringOutput { return v.Realm }).(pulumi.StringOutput)
}

// Is Attribute Related To a User Session
func (o HardcodedAttributeIdentityProviderMapperOutput) UserSession() pulumi.BoolOutput {
	return o.ApplyT(func(v *HardcodedAttributeIdentityProviderMapper) pulumi.BoolOutput { return v.UserSession }).(pulumi.BoolOutput)
}

type HardcodedAttributeIdentityProviderMapperArrayOutput struct{ *pulumi.OutputState }

func (HardcodedAttributeIdentityProviderMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedAttributeIdentityProviderMapper)(nil)).Elem()
}

func (o HardcodedAttributeIdentityProviderMapperArrayOutput) ToHardcodedAttributeIdentityProviderMapperArrayOutput() HardcodedAttributeIdentityProviderMapperArrayOutput {
	return o
}

func (o HardcodedAttributeIdentityProviderMapperArrayOutput) ToHardcodedAttributeIdentityProviderMapperArrayOutputWithContext(ctx context.Context) HardcodedAttributeIdentityProviderMapperArrayOutput {
	return o
}

func (o HardcodedAttributeIdentityProviderMapperArrayOutput) Index(i pulumi.IntInput) HardcodedAttributeIdentityProviderMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HardcodedAttributeIdentityProviderMapper {
		return vs[0].([]*HardcodedAttributeIdentityProviderMapper)[vs[1].(int)]
	}).(HardcodedAttributeIdentityProviderMapperOutput)
}

type HardcodedAttributeIdentityProviderMapperMapOutput struct{ *pulumi.OutputState }

func (HardcodedAttributeIdentityProviderMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedAttributeIdentityProviderMapper)(nil)).Elem()
}

func (o HardcodedAttributeIdentityProviderMapperMapOutput) ToHardcodedAttributeIdentityProviderMapperMapOutput() HardcodedAttributeIdentityProviderMapperMapOutput {
	return o
}

func (o HardcodedAttributeIdentityProviderMapperMapOutput) ToHardcodedAttributeIdentityProviderMapperMapOutputWithContext(ctx context.Context) HardcodedAttributeIdentityProviderMapperMapOutput {
	return o
}

func (o HardcodedAttributeIdentityProviderMapperMapOutput) MapIndex(k pulumi.StringInput) HardcodedAttributeIdentityProviderMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HardcodedAttributeIdentityProviderMapper {
		return vs[0].(map[string]*HardcodedAttributeIdentityProviderMapper)[vs[1].(string)]
	}).(HardcodedAttributeIdentityProviderMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedAttributeIdentityProviderMapperInput)(nil)).Elem(), &HardcodedAttributeIdentityProviderMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedAttributeIdentityProviderMapperArrayInput)(nil)).Elem(), HardcodedAttributeIdentityProviderMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedAttributeIdentityProviderMapperMapInput)(nil)).Elem(), HardcodedAttributeIdentityProviderMapperMap{})
	pulumi.RegisterOutputType(HardcodedAttributeIdentityProviderMapperOutput{})
	pulumi.RegisterOutputType(HardcodedAttributeIdentityProviderMapperArrayOutput{})
	pulumi.RegisterOutputType(HardcodedAttributeIdentityProviderMapperMapOutput{})
}
