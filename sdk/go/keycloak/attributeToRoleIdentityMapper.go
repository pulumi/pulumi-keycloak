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

// Allows for creating and managing an attribute to role identity provider mapper within Keycloak.
//
// > If you are using Keycloak 10 or higher, you will need to specify the `extraConfig` argument in order to define a `syncMode` for the mapper.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/oidc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
//				Realm:   pulumi.String("my-realm"),
//				Enabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			oidc, err := oidc.NewIdentityProvider(ctx, "oidc", &oidc.IdentityProviderArgs{
//				Realm:            realm.ID(),
//				Alias:            pulumi.String("oidc"),
//				AuthorizationUrl: pulumi.String("https://example.com/auth"),
//				TokenUrl:         pulumi.String("https://example.com/token"),
//				ClientId:         pulumi.String("example_id"),
//				ClientSecret:     pulumi.String("example_token"),
//				DefaultScopes:    pulumi.String("openid random profile"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "realm_role", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				Name:        pulumi.String("my-realm-role"),
//				Description: pulumi.String("My Realm Role"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewAttributeToRoleIdentityMapper(ctx, "oidc", &keycloak.AttributeToRoleIdentityMapperArgs{
//				Realm:                 realm.ID(),
//				Name:                  pulumi.String("role-attribute"),
//				IdentityProviderAlias: oidc.Alias,
//				Role:                  pulumi.String("my-realm-role"),
//				ClaimName:             pulumi.String("my-claim"),
//				ClaimValue:            pulumi.String("my-value"),
//				ExtraConfig: pulumi.Map{
//					"syncMode": pulumi.Any("INHERIT"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak
//
// assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID.
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper test_mapper my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
// ```
type AttributeToRoleIdentityMapper struct {
	pulumi.CustomResourceState

	// Attribute Friendly Name. Conflicts with `attributeName`.
	AttributeFriendlyName pulumi.StringPtrOutput `pulumi:"attributeFriendlyName"`
	// Attribute Name.
	AttributeName pulumi.StringPtrOutput `pulumi:"attributeName"`
	// Attribute Value.
	AttributeValue pulumi.StringPtrOutput `pulumi:"attributeValue"`
	// OIDC Claim Name
	ClaimName pulumi.StringPtrOutput `pulumi:"claimName"`
	// OIDC Claim Value
	ClaimValue pulumi.StringPtrOutput `pulumi:"claimValue"`
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapOutput `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// The name of the mapper.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the realm.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Role Name.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewAttributeToRoleIdentityMapper registers a new resource with the given unique name, arguments, and options.
func NewAttributeToRoleIdentityMapper(ctx *pulumi.Context,
	name string, args *AttributeToRoleIdentityMapperArgs, opts ...pulumi.ResourceOption) (*AttributeToRoleIdentityMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityProviderAlias == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderAlias'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AttributeToRoleIdentityMapper
	err := ctx.RegisterResource("keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttributeToRoleIdentityMapper gets an existing AttributeToRoleIdentityMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttributeToRoleIdentityMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttributeToRoleIdentityMapperState, opts ...pulumi.ResourceOption) (*AttributeToRoleIdentityMapper, error) {
	var resource AttributeToRoleIdentityMapper
	err := ctx.ReadResource("keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttributeToRoleIdentityMapper resources.
type attributeToRoleIdentityMapperState struct {
	// Attribute Friendly Name. Conflicts with `attributeName`.
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// Attribute Name.
	AttributeName *string `pulumi:"attributeName"`
	// Attribute Value.
	AttributeValue *string `pulumi:"attributeValue"`
	// OIDC Claim Name
	ClaimName *string `pulumi:"claimName"`
	// OIDC Claim Value
	ClaimValue *string `pulumi:"claimValue"`
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// The name of the mapper.
	Name *string `pulumi:"name"`
	// The name of the realm.
	Realm *string `pulumi:"realm"`
	// Role Name.
	Role *string `pulumi:"role"`
}

type AttributeToRoleIdentityMapperState struct {
	// Attribute Friendly Name. Conflicts with `attributeName`.
	AttributeFriendlyName pulumi.StringPtrInput
	// Attribute Name.
	AttributeName pulumi.StringPtrInput
	// Attribute Value.
	AttributeValue pulumi.StringPtrInput
	// OIDC Claim Name
	ClaimName pulumi.StringPtrInput
	// OIDC Claim Value
	ClaimValue pulumi.StringPtrInput
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapInput
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringPtrInput
	// The name of the mapper.
	Name pulumi.StringPtrInput
	// The name of the realm.
	Realm pulumi.StringPtrInput
	// Role Name.
	Role pulumi.StringPtrInput
}

func (AttributeToRoleIdentityMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeToRoleIdentityMapperState)(nil)).Elem()
}

type attributeToRoleIdentityMapperArgs struct {
	// Attribute Friendly Name. Conflicts with `attributeName`.
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// Attribute Name.
	AttributeName *string `pulumi:"attributeName"`
	// Attribute Value.
	AttributeValue *string `pulumi:"attributeValue"`
	// OIDC Claim Name
	ClaimName *string `pulumi:"claimName"`
	// OIDC Claim Value
	ClaimValue *string `pulumi:"claimValue"`
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// The name of the mapper.
	Name *string `pulumi:"name"`
	// The name of the realm.
	Realm string `pulumi:"realm"`
	// Role Name.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a AttributeToRoleIdentityMapper resource.
type AttributeToRoleIdentityMapperArgs struct {
	// Attribute Friendly Name. Conflicts with `attributeName`.
	AttributeFriendlyName pulumi.StringPtrInput
	// Attribute Name.
	AttributeName pulumi.StringPtrInput
	// Attribute Value.
	AttributeValue pulumi.StringPtrInput
	// OIDC Claim Name
	ClaimName pulumi.StringPtrInput
	// OIDC Claim Value
	ClaimValue pulumi.StringPtrInput
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapInput
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringInput
	// The name of the mapper.
	Name pulumi.StringPtrInput
	// The name of the realm.
	Realm pulumi.StringInput
	// Role Name.
	Role pulumi.StringInput
}

func (AttributeToRoleIdentityMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeToRoleIdentityMapperArgs)(nil)).Elem()
}

type AttributeToRoleIdentityMapperInput interface {
	pulumi.Input

	ToAttributeToRoleIdentityMapperOutput() AttributeToRoleIdentityMapperOutput
	ToAttributeToRoleIdentityMapperOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperOutput
}

func (*AttributeToRoleIdentityMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (i *AttributeToRoleIdentityMapper) ToAttributeToRoleIdentityMapperOutput() AttributeToRoleIdentityMapperOutput {
	return i.ToAttributeToRoleIdentityMapperOutputWithContext(context.Background())
}

func (i *AttributeToRoleIdentityMapper) ToAttributeToRoleIdentityMapperOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeToRoleIdentityMapperOutput)
}

// AttributeToRoleIdentityMapperArrayInput is an input type that accepts AttributeToRoleIdentityMapperArray and AttributeToRoleIdentityMapperArrayOutput values.
// You can construct a concrete instance of `AttributeToRoleIdentityMapperArrayInput` via:
//
//	AttributeToRoleIdentityMapperArray{ AttributeToRoleIdentityMapperArgs{...} }
type AttributeToRoleIdentityMapperArrayInput interface {
	pulumi.Input

	ToAttributeToRoleIdentityMapperArrayOutput() AttributeToRoleIdentityMapperArrayOutput
	ToAttributeToRoleIdentityMapperArrayOutputWithContext(context.Context) AttributeToRoleIdentityMapperArrayOutput
}

type AttributeToRoleIdentityMapperArray []AttributeToRoleIdentityMapperInput

func (AttributeToRoleIdentityMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (i AttributeToRoleIdentityMapperArray) ToAttributeToRoleIdentityMapperArrayOutput() AttributeToRoleIdentityMapperArrayOutput {
	return i.ToAttributeToRoleIdentityMapperArrayOutputWithContext(context.Background())
}

func (i AttributeToRoleIdentityMapperArray) ToAttributeToRoleIdentityMapperArrayOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeToRoleIdentityMapperArrayOutput)
}

// AttributeToRoleIdentityMapperMapInput is an input type that accepts AttributeToRoleIdentityMapperMap and AttributeToRoleIdentityMapperMapOutput values.
// You can construct a concrete instance of `AttributeToRoleIdentityMapperMapInput` via:
//
//	AttributeToRoleIdentityMapperMap{ "key": AttributeToRoleIdentityMapperArgs{...} }
type AttributeToRoleIdentityMapperMapInput interface {
	pulumi.Input

	ToAttributeToRoleIdentityMapperMapOutput() AttributeToRoleIdentityMapperMapOutput
	ToAttributeToRoleIdentityMapperMapOutputWithContext(context.Context) AttributeToRoleIdentityMapperMapOutput
}

type AttributeToRoleIdentityMapperMap map[string]AttributeToRoleIdentityMapperInput

func (AttributeToRoleIdentityMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (i AttributeToRoleIdentityMapperMap) ToAttributeToRoleIdentityMapperMapOutput() AttributeToRoleIdentityMapperMapOutput {
	return i.ToAttributeToRoleIdentityMapperMapOutputWithContext(context.Background())
}

func (i AttributeToRoleIdentityMapperMap) ToAttributeToRoleIdentityMapperMapOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttributeToRoleIdentityMapperMapOutput)
}

type AttributeToRoleIdentityMapperOutput struct{ *pulumi.OutputState }

func (AttributeToRoleIdentityMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (o AttributeToRoleIdentityMapperOutput) ToAttributeToRoleIdentityMapperOutput() AttributeToRoleIdentityMapperOutput {
	return o
}

func (o AttributeToRoleIdentityMapperOutput) ToAttributeToRoleIdentityMapperOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperOutput {
	return o
}

// Attribute Friendly Name. Conflicts with `attributeName`.
func (o AttributeToRoleIdentityMapperOutput) AttributeFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.StringPtrOutput { return v.AttributeFriendlyName }).(pulumi.StringPtrOutput)
}

// Attribute Name.
func (o AttributeToRoleIdentityMapperOutput) AttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.StringPtrOutput { return v.AttributeName }).(pulumi.StringPtrOutput)
}

// Attribute Value.
func (o AttributeToRoleIdentityMapperOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.StringPtrOutput { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

// OIDC Claim Name
func (o AttributeToRoleIdentityMapperOutput) ClaimName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.StringPtrOutput { return v.ClaimName }).(pulumi.StringPtrOutput)
}

// OIDC Claim Value
func (o AttributeToRoleIdentityMapperOutput) ClaimValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.StringPtrOutput { return v.ClaimValue }).(pulumi.StringPtrOutput)
}

// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
func (o AttributeToRoleIdentityMapperOutput) ExtraConfig() pulumi.MapOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.MapOutput { return v.ExtraConfig }).(pulumi.MapOutput)
}

// The alias of the associated identity provider.
func (o AttributeToRoleIdentityMapperOutput) IdentityProviderAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.StringOutput { return v.IdentityProviderAlias }).(pulumi.StringOutput)
}

// The name of the mapper.
func (o AttributeToRoleIdentityMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the realm.
func (o AttributeToRoleIdentityMapperOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.StringOutput { return v.Realm }).(pulumi.StringOutput)
}

// Role Name.
func (o AttributeToRoleIdentityMapperOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *AttributeToRoleIdentityMapper) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type AttributeToRoleIdentityMapperArrayOutput struct{ *pulumi.OutputState }

func (AttributeToRoleIdentityMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (o AttributeToRoleIdentityMapperArrayOutput) ToAttributeToRoleIdentityMapperArrayOutput() AttributeToRoleIdentityMapperArrayOutput {
	return o
}

func (o AttributeToRoleIdentityMapperArrayOutput) ToAttributeToRoleIdentityMapperArrayOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperArrayOutput {
	return o
}

func (o AttributeToRoleIdentityMapperArrayOutput) Index(i pulumi.IntInput) AttributeToRoleIdentityMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AttributeToRoleIdentityMapper {
		return vs[0].([]*AttributeToRoleIdentityMapper)[vs[1].(int)]
	}).(AttributeToRoleIdentityMapperOutput)
}

type AttributeToRoleIdentityMapperMapOutput struct{ *pulumi.OutputState }

func (AttributeToRoleIdentityMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttributeToRoleIdentityMapper)(nil)).Elem()
}

func (o AttributeToRoleIdentityMapperMapOutput) ToAttributeToRoleIdentityMapperMapOutput() AttributeToRoleIdentityMapperMapOutput {
	return o
}

func (o AttributeToRoleIdentityMapperMapOutput) ToAttributeToRoleIdentityMapperMapOutputWithContext(ctx context.Context) AttributeToRoleIdentityMapperMapOutput {
	return o
}

func (o AttributeToRoleIdentityMapperMapOutput) MapIndex(k pulumi.StringInput) AttributeToRoleIdentityMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AttributeToRoleIdentityMapper {
		return vs[0].(map[string]*AttributeToRoleIdentityMapper)[vs[1].(string)]
	}).(AttributeToRoleIdentityMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeToRoleIdentityMapperInput)(nil)).Elem(), &AttributeToRoleIdentityMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeToRoleIdentityMapperArrayInput)(nil)).Elem(), AttributeToRoleIdentityMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttributeToRoleIdentityMapperMapInput)(nil)).Elem(), AttributeToRoleIdentityMapperMap{})
	pulumi.RegisterOutputType(AttributeToRoleIdentityMapperOutput{})
	pulumi.RegisterOutputType(AttributeToRoleIdentityMapperArrayOutput{})
	pulumi.RegisterOutputType(AttributeToRoleIdentityMapperMapOutput{})
}
