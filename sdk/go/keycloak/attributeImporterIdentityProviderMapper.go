// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # AttributeImporterIdentityProviderMapper
//
// Allows to create and manage identity provider mappers within Keycloak.
//
// ### Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := keycloak.NewAttributeImporterIdentityProviderMapper(ctx, "test_mapper", &keycloak.AttributeImporterIdentityProviderMapperArgs{
//				Realm:                 pulumi.String("my-realm"),
//				Name:                  pulumi.String("my-mapper"),
//				IdentityProviderAlias: pulumi.String("idp_alias"),
//				AttributeName:         pulumi.String("http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname"),
//				UserAttribute:         pulumi.String("lastName"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realm` - (Required) The name of the realm.
// - `name` - (Required) The name of the mapper.
// - `identityProviderAlias` - (Required) The alias of the associated identity provider.
// - `userAttribute` - (Required) The user attribute name to store SAML attribute.
// - `attributeName` - (Optional) The Name of attribute to search for in assertion. You can leave this blank and specify a friendly name instead.
// - `attributeFriendlyName` - (Optional) The friendly name of attribute to search for in assertion.  You can leave this blank and specify an attribute name instead.
// - `claimName` - (Optional) The claim name.
//
// ### Import
//
// Identity provider mapper can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idpAlias` is the identity provider alias, and `idpMapperId` is the unique ID that Keycloak
// assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID.
//
// Example:
type AttributeImporterIdentityProviderMapper struct {
	pulumi.CustomResourceState

	// Attribute Friendly Name
	AttributeFriendlyName pulumi.StringPtrOutput `pulumi:"attributeFriendlyName"`
	// Attribute Name
	AttributeName pulumi.StringPtrOutput `pulumi:"attributeName"`
	// Claim Name
	ClaimName   pulumi.StringPtrOutput `pulumi:"claimName"`
	ExtraConfig pulumi.StringMapOutput `pulumi:"extraConfig"`
	// IDP Alias
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Realm Name
	Realm pulumi.StringOutput `pulumi:"realm"`
	// User Attribute
	UserAttribute pulumi.StringOutput `pulumi:"userAttribute"`
}

// NewAttributeImporterIdentityProviderMapper registers a new resource with the given unique name, arguments, and options.
func NewAttributeImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, args *AttributeImporterIdentityProviderMapperArgs, opts ...pulumi.ResourceOption) (*AttributeImporterIdentityProviderMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityProviderAlias == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderAlias'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	if args.UserAttribute == nil {
		return nil, errors.New("invalid value for required argument 'UserAttribute'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AttributeImporterIdentityProviderMapper
	err := ctx.RegisterResource("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttributeImporterIdentityProviderMapper gets an existing AttributeImporterIdentityProviderMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttributeImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttributeImporterIdentityProviderMapperState, opts ...pulumi.ResourceOption) (*AttributeImporterIdentityProviderMapper, error) {
	var resource AttributeImporterIdentityProviderMapper
	err := ctx.ReadResource("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttributeImporterIdentityProviderMapper resources.
type attributeImporterIdentityProviderMapperState struct {
	// Attribute Friendly Name
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// Attribute Name
	AttributeName *string `pulumi:"attributeName"`
	// Claim Name
	ClaimName   *string           `pulumi:"claimName"`
	ExtraConfig map[string]string `pulumi:"extraConfig"`
	// IDP Alias
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm *string `pulumi:"realm"`
	// User Attribute
	UserAttribute *string `pulumi:"userAttribute"`
}

type AttributeImporterIdentityProviderMapperState struct {
	// Attribute Friendly Name
	AttributeFriendlyName pulumi.StringPtrInput
	// Attribute Name
	AttributeName pulumi.StringPtrInput
	// Claim Name
	ClaimName   pulumi.StringPtrInput
	ExtraConfig pulumi.StringMapInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringPtrInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringPtrInput
	// User Attribute
	UserAttribute pulumi.StringPtrInput
}

func (AttributeImporterIdentityProviderMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeImporterIdentityProviderMapperState)(nil)).Elem()
}

type attributeImporterIdentityProviderMapperArgs struct {
	// Attribute Friendly Name
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// Attribute Name
	AttributeName *string `pulumi:"attributeName"`
	// Claim Name
	ClaimName   *string           `pulumi:"claimName"`
	ExtraConfig map[string]string `pulumi:"extraConfig"`
	// IDP Alias
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm string `pulumi:"realm"`
	// User Attribute
	UserAttribute string `pulumi:"userAttribute"`
}

// The set of arguments for constructing a AttributeImporterIdentityProviderMapper resource.
type AttributeImporterIdentityProviderMapperArgs struct {
	// Attribute Friendly Name
	AttributeFriendlyName pulumi.StringPtrInput
	// Attribute Name
	AttributeName pulumi.StringPtrInput
	// Claim Name
	ClaimName   pulumi.StringPtrInput
	ExtraConfig pulumi.StringMapInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringInput
	// User Attribute
	UserAttribute pulumi.StringInput
}

func (AttributeImporterIdentityProviderMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeImporterIdentityProviderMapperArgs)(nil)).Elem()
}

type AttributeImporterIdentityProviderMapperInput interface {
	pulumi.Input

	ToAttributeImporterIdentityProviderMapperOutput() AttributeImporterIdentityProviderMapperOutput
	ToAttributeImporterIdentityProviderMapperOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperOutput
}

func (*AttributeImporterIdentityProviderMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeImporterIdentityProviderMapper)(nil)).Elem()
}

func (i *AttributeImporterIdentityProviderMapper) ToAttributeImporterIdentityProviderMapperOutput() AttributeImporterIdentityProviderMapperOutput {
	return i.ToAttributeImporterIdentityProviderMapperOutputWithContext(context.Background())
}

func (i *AttributeImporterIdentityProviderMapper) ToAttributeImporterIdentityProviderMapperOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeImporterIdentityProviderMapperOutput)
}

// AttributeImporterIdentityProviderMapperArrayInput is an input type that accepts AttributeImporterIdentityProviderMapperArray and AttributeImporterIdentityProviderMapperArrayOutput values.
// You can construct a concrete instance of `AttributeImporterIdentityProviderMapperArrayInput` via:
//
//	AttributeImporterIdentityProviderMapperArray{ AttributeImporterIdentityProviderMapperArgs{...} }
type AttributeImporterIdentityProviderMapperArrayInput interface {
	pulumi.Input

	ToAttributeImporterIdentityProviderMapperArrayOutput() AttributeImporterIdentityProviderMapperArrayOutput
	ToAttributeImporterIdentityProviderMapperArrayOutputWithContext(context.Context) AttributeImporterIdentityProviderMapperArrayOutput
}

type AttributeImporterIdentityProviderMapperArray []AttributeImporterIdentityProviderMapperInput

func (AttributeImporterIdentityProviderMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttributeImporterIdentityProviderMapper)(nil)).Elem()
}

func (i AttributeImporterIdentityProviderMapperArray) ToAttributeImporterIdentityProviderMapperArrayOutput() AttributeImporterIdentityProviderMapperArrayOutput {
	return i.ToAttributeImporterIdentityProviderMapperArrayOutputWithContext(context.Background())
}

func (i AttributeImporterIdentityProviderMapperArray) ToAttributeImporterIdentityProviderMapperArrayOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeImporterIdentityProviderMapperArrayOutput)
}

// AttributeImporterIdentityProviderMapperMapInput is an input type that accepts AttributeImporterIdentityProviderMapperMap and AttributeImporterIdentityProviderMapperMapOutput values.
// You can construct a concrete instance of `AttributeImporterIdentityProviderMapperMapInput` via:
//
//	AttributeImporterIdentityProviderMapperMap{ "key": AttributeImporterIdentityProviderMapperArgs{...} }
type AttributeImporterIdentityProviderMapperMapInput interface {
	pulumi.Input

	ToAttributeImporterIdentityProviderMapperMapOutput() AttributeImporterIdentityProviderMapperMapOutput
	ToAttributeImporterIdentityProviderMapperMapOutputWithContext(context.Context) AttributeImporterIdentityProviderMapperMapOutput
}

type AttributeImporterIdentityProviderMapperMap map[string]AttributeImporterIdentityProviderMapperInput

func (AttributeImporterIdentityProviderMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttributeImporterIdentityProviderMapper)(nil)).Elem()
}

func (i AttributeImporterIdentityProviderMapperMap) ToAttributeImporterIdentityProviderMapperMapOutput() AttributeImporterIdentityProviderMapperMapOutput {
	return i.ToAttributeImporterIdentityProviderMapperMapOutputWithContext(context.Background())
}

func (i AttributeImporterIdentityProviderMapperMap) ToAttributeImporterIdentityProviderMapperMapOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeImporterIdentityProviderMapperMapOutput)
}

type AttributeImporterIdentityProviderMapperOutput struct{ *pulumi.OutputState }

func (AttributeImporterIdentityProviderMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeImporterIdentityProviderMapper)(nil)).Elem()
}

func (o AttributeImporterIdentityProviderMapperOutput) ToAttributeImporterIdentityProviderMapperOutput() AttributeImporterIdentityProviderMapperOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperOutput) ToAttributeImporterIdentityProviderMapperOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperOutput {
	return o
}

// Attribute Friendly Name
func (o AttributeImporterIdentityProviderMapperOutput) AttributeFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttributeImporterIdentityProviderMapper) pulumi.StringPtrOutput {
		return v.AttributeFriendlyName
	}).(pulumi.StringPtrOutput)
}

// Attribute Name
func (o AttributeImporterIdentityProviderMapperOutput) AttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttributeImporterIdentityProviderMapper) pulumi.StringPtrOutput { return v.AttributeName }).(pulumi.StringPtrOutput)
}

// Claim Name
func (o AttributeImporterIdentityProviderMapperOutput) ClaimName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttributeImporterIdentityProviderMapper) pulumi.StringPtrOutput { return v.ClaimName }).(pulumi.StringPtrOutput)
}

func (o AttributeImporterIdentityProviderMapperOutput) ExtraConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AttributeImporterIdentityProviderMapper) pulumi.StringMapOutput { return v.ExtraConfig }).(pulumi.StringMapOutput)
}

// IDP Alias
func (o AttributeImporterIdentityProviderMapperOutput) IdentityProviderAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *AttributeImporterIdentityProviderMapper) pulumi.StringOutput { return v.IdentityProviderAlias }).(pulumi.StringOutput)
}

// IDP Mapper Name
func (o AttributeImporterIdentityProviderMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttributeImporterIdentityProviderMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Realm Name
func (o AttributeImporterIdentityProviderMapperOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v *AttributeImporterIdentityProviderMapper) pulumi.StringOutput { return v.Realm }).(pulumi.StringOutput)
}

// User Attribute
func (o AttributeImporterIdentityProviderMapperOutput) UserAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *AttributeImporterIdentityProviderMapper) pulumi.StringOutput { return v.UserAttribute }).(pulumi.StringOutput)
}

type AttributeImporterIdentityProviderMapperArrayOutput struct{ *pulumi.OutputState }

func (AttributeImporterIdentityProviderMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttributeImporterIdentityProviderMapper)(nil)).Elem()
}

func (o AttributeImporterIdentityProviderMapperArrayOutput) ToAttributeImporterIdentityProviderMapperArrayOutput() AttributeImporterIdentityProviderMapperArrayOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperArrayOutput) ToAttributeImporterIdentityProviderMapperArrayOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperArrayOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperArrayOutput) Index(i pulumi.IntInput) AttributeImporterIdentityProviderMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AttributeImporterIdentityProviderMapper {
		return vs[0].([]*AttributeImporterIdentityProviderMapper)[vs[1].(int)]
	}).(AttributeImporterIdentityProviderMapperOutput)
}

type AttributeImporterIdentityProviderMapperMapOutput struct{ *pulumi.OutputState }

func (AttributeImporterIdentityProviderMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttributeImporterIdentityProviderMapper)(nil)).Elem()
}

func (o AttributeImporterIdentityProviderMapperMapOutput) ToAttributeImporterIdentityProviderMapperMapOutput() AttributeImporterIdentityProviderMapperMapOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperMapOutput) ToAttributeImporterIdentityProviderMapperMapOutputWithContext(ctx context.Context) AttributeImporterIdentityProviderMapperMapOutput {
	return o
}

func (o AttributeImporterIdentityProviderMapperMapOutput) MapIndex(k pulumi.StringInput) AttributeImporterIdentityProviderMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AttributeImporterIdentityProviderMapper {
		return vs[0].(map[string]*AttributeImporterIdentityProviderMapper)[vs[1].(string)]
	}).(AttributeImporterIdentityProviderMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeImporterIdentityProviderMapperInput)(nil)).Elem(), &AttributeImporterIdentityProviderMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeImporterIdentityProviderMapperArrayInput)(nil)).Elem(), AttributeImporterIdentityProviderMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeImporterIdentityProviderMapperMapInput)(nil)).Elem(), AttributeImporterIdentityProviderMapperMap{})
	pulumi.RegisterOutputType(AttributeImporterIdentityProviderMapperOutput{})
	pulumi.RegisterOutputType(AttributeImporterIdentityProviderMapperArrayOutput{})
	pulumi.RegisterOutputType(AttributeImporterIdentityProviderMapperMapOutput{})
}
