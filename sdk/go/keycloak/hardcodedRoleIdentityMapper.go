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

// Allows for creating and managing hardcoded role mappers for Keycloak identity provider.
//
// The identity provider hardcoded role mapper grants a specified Keycloak role to each Keycloak user from the LDAP provider.
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
//			oidcIdentityProvider, err := oidc.NewIdentityProvider(ctx, "oidcIdentityProvider", &oidc.IdentityProviderArgs{
//				Realm:            realm.ID(),
//				Alias:            pulumi.String("my-idp"),
//				AuthorizationUrl: pulumi.String("https://authorizationurl.com"),
//				ClientId:         pulumi.String("clientID"),
//				ClientSecret:     pulumi.String("clientSecret"),
//				TokenUrl:         pulumi.String("https://tokenurl.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "realmRole", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				Description: pulumi.String("My Realm Role"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewHardcodedRoleIdentityMapper(ctx, "oidcHardcodedRoleIdentityMapper", &keycloak.HardcodedRoleIdentityMapperArgs{
//				Realm:                 realm.ID(),
//				IdentityProviderAlias: oidcIdentityProvider.Alias,
//				Role:                  pulumi.String("my-realm-role"),
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
type HardcodedRoleIdentityMapper struct {
	pulumi.CustomResourceState

	ExtraConfig pulumi.MapOutput `pulumi:"extraConfig"`
	// The IDP alias of the attribute to set.
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm ID that this mapper will exist in.
	Realm pulumi.StringOutput `pulumi:"realm"`
	// The name of the role which should be assigned to the users.
	Role pulumi.StringPtrOutput `pulumi:"role"`
}

// NewHardcodedRoleIdentityMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedRoleIdentityMapper(ctx *pulumi.Context,
	name string, args *HardcodedRoleIdentityMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedRoleIdentityMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityProviderAlias == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderAlias'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HardcodedRoleIdentityMapper
	err := ctx.RegisterResource("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedRoleIdentityMapper gets an existing HardcodedRoleIdentityMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedRoleIdentityMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedRoleIdentityMapperState, opts ...pulumi.ResourceOption) (*HardcodedRoleIdentityMapper, error) {
	var resource HardcodedRoleIdentityMapper
	err := ctx.ReadResource("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedRoleIdentityMapper resources.
type hardcodedRoleIdentityMapperState struct {
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The IDP alias of the attribute to set.
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm ID that this mapper will exist in.
	Realm *string `pulumi:"realm"`
	// The name of the role which should be assigned to the users.
	Role *string `pulumi:"role"`
}

type HardcodedRoleIdentityMapperState struct {
	ExtraConfig pulumi.MapInput
	// The IDP alias of the attribute to set.
	IdentityProviderAlias pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm ID that this mapper will exist in.
	Realm pulumi.StringPtrInput
	// The name of the role which should be assigned to the users.
	Role pulumi.StringPtrInput
}

func (HardcodedRoleIdentityMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleIdentityMapperState)(nil)).Elem()
}

type hardcodedRoleIdentityMapperArgs struct {
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The IDP alias of the attribute to set.
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm ID that this mapper will exist in.
	Realm string `pulumi:"realm"`
	// The name of the role which should be assigned to the users.
	Role *string `pulumi:"role"`
}

// The set of arguments for constructing a HardcodedRoleIdentityMapper resource.
type HardcodedRoleIdentityMapperArgs struct {
	ExtraConfig pulumi.MapInput
	// The IDP alias of the attribute to set.
	IdentityProviderAlias pulumi.StringInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm ID that this mapper will exist in.
	Realm pulumi.StringInput
	// The name of the role which should be assigned to the users.
	Role pulumi.StringPtrInput
}

func (HardcodedRoleIdentityMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleIdentityMapperArgs)(nil)).Elem()
}

type HardcodedRoleIdentityMapperInput interface {
	pulumi.Input

	ToHardcodedRoleIdentityMapperOutput() HardcodedRoleIdentityMapperOutput
	ToHardcodedRoleIdentityMapperOutputWithContext(ctx context.Context) HardcodedRoleIdentityMapperOutput
}

func (*HardcodedRoleIdentityMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedRoleIdentityMapper)(nil)).Elem()
}

func (i *HardcodedRoleIdentityMapper) ToHardcodedRoleIdentityMapperOutput() HardcodedRoleIdentityMapperOutput {
	return i.ToHardcodedRoleIdentityMapperOutputWithContext(context.Background())
}

func (i *HardcodedRoleIdentityMapper) ToHardcodedRoleIdentityMapperOutputWithContext(ctx context.Context) HardcodedRoleIdentityMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleIdentityMapperOutput)
}

// HardcodedRoleIdentityMapperArrayInput is an input type that accepts HardcodedRoleIdentityMapperArray and HardcodedRoleIdentityMapperArrayOutput values.
// You can construct a concrete instance of `HardcodedRoleIdentityMapperArrayInput` via:
//
//	HardcodedRoleIdentityMapperArray{ HardcodedRoleIdentityMapperArgs{...} }
type HardcodedRoleIdentityMapperArrayInput interface {
	pulumi.Input

	ToHardcodedRoleIdentityMapperArrayOutput() HardcodedRoleIdentityMapperArrayOutput
	ToHardcodedRoleIdentityMapperArrayOutputWithContext(context.Context) HardcodedRoleIdentityMapperArrayOutput
}

type HardcodedRoleIdentityMapperArray []HardcodedRoleIdentityMapperInput

func (HardcodedRoleIdentityMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedRoleIdentityMapper)(nil)).Elem()
}

func (i HardcodedRoleIdentityMapperArray) ToHardcodedRoleIdentityMapperArrayOutput() HardcodedRoleIdentityMapperArrayOutput {
	return i.ToHardcodedRoleIdentityMapperArrayOutputWithContext(context.Background())
}

func (i HardcodedRoleIdentityMapperArray) ToHardcodedRoleIdentityMapperArrayOutputWithContext(ctx context.Context) HardcodedRoleIdentityMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleIdentityMapperArrayOutput)
}

// HardcodedRoleIdentityMapperMapInput is an input type that accepts HardcodedRoleIdentityMapperMap and HardcodedRoleIdentityMapperMapOutput values.
// You can construct a concrete instance of `HardcodedRoleIdentityMapperMapInput` via:
//
//	HardcodedRoleIdentityMapperMap{ "key": HardcodedRoleIdentityMapperArgs{...} }
type HardcodedRoleIdentityMapperMapInput interface {
	pulumi.Input

	ToHardcodedRoleIdentityMapperMapOutput() HardcodedRoleIdentityMapperMapOutput
	ToHardcodedRoleIdentityMapperMapOutputWithContext(context.Context) HardcodedRoleIdentityMapperMapOutput
}

type HardcodedRoleIdentityMapperMap map[string]HardcodedRoleIdentityMapperInput

func (HardcodedRoleIdentityMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedRoleIdentityMapper)(nil)).Elem()
}

func (i HardcodedRoleIdentityMapperMap) ToHardcodedRoleIdentityMapperMapOutput() HardcodedRoleIdentityMapperMapOutput {
	return i.ToHardcodedRoleIdentityMapperMapOutputWithContext(context.Background())
}

func (i HardcodedRoleIdentityMapperMap) ToHardcodedRoleIdentityMapperMapOutputWithContext(ctx context.Context) HardcodedRoleIdentityMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleIdentityMapperMapOutput)
}

type HardcodedRoleIdentityMapperOutput struct{ *pulumi.OutputState }

func (HardcodedRoleIdentityMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedRoleIdentityMapper)(nil)).Elem()
}

func (o HardcodedRoleIdentityMapperOutput) ToHardcodedRoleIdentityMapperOutput() HardcodedRoleIdentityMapperOutput {
	return o
}

func (o HardcodedRoleIdentityMapperOutput) ToHardcodedRoleIdentityMapperOutputWithContext(ctx context.Context) HardcodedRoleIdentityMapperOutput {
	return o
}

func (o HardcodedRoleIdentityMapperOutput) ExtraConfig() pulumi.MapOutput {
	return o.ApplyT(func(v *HardcodedRoleIdentityMapper) pulumi.MapOutput { return v.ExtraConfig }).(pulumi.MapOutput)
}

// The IDP alias of the attribute to set.
func (o HardcodedRoleIdentityMapperOutput) IdentityProviderAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedRoleIdentityMapper) pulumi.StringOutput { return v.IdentityProviderAlias }).(pulumi.StringOutput)
}

// Display name of this mapper when displayed in the console.
func (o HardcodedRoleIdentityMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedRoleIdentityMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm ID that this mapper will exist in.
func (o HardcodedRoleIdentityMapperOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedRoleIdentityMapper) pulumi.StringOutput { return v.Realm }).(pulumi.StringOutput)
}

// The name of the role which should be assigned to the users.
func (o HardcodedRoleIdentityMapperOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardcodedRoleIdentityMapper) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

type HardcodedRoleIdentityMapperArrayOutput struct{ *pulumi.OutputState }

func (HardcodedRoleIdentityMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedRoleIdentityMapper)(nil)).Elem()
}

func (o HardcodedRoleIdentityMapperArrayOutput) ToHardcodedRoleIdentityMapperArrayOutput() HardcodedRoleIdentityMapperArrayOutput {
	return o
}

func (o HardcodedRoleIdentityMapperArrayOutput) ToHardcodedRoleIdentityMapperArrayOutputWithContext(ctx context.Context) HardcodedRoleIdentityMapperArrayOutput {
	return o
}

func (o HardcodedRoleIdentityMapperArrayOutput) Index(i pulumi.IntInput) HardcodedRoleIdentityMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HardcodedRoleIdentityMapper {
		return vs[0].([]*HardcodedRoleIdentityMapper)[vs[1].(int)]
	}).(HardcodedRoleIdentityMapperOutput)
}

type HardcodedRoleIdentityMapperMapOutput struct{ *pulumi.OutputState }

func (HardcodedRoleIdentityMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedRoleIdentityMapper)(nil)).Elem()
}

func (o HardcodedRoleIdentityMapperMapOutput) ToHardcodedRoleIdentityMapperMapOutput() HardcodedRoleIdentityMapperMapOutput {
	return o
}

func (o HardcodedRoleIdentityMapperMapOutput) ToHardcodedRoleIdentityMapperMapOutputWithContext(ctx context.Context) HardcodedRoleIdentityMapperMapOutput {
	return o
}

func (o HardcodedRoleIdentityMapperMapOutput) MapIndex(k pulumi.StringInput) HardcodedRoleIdentityMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HardcodedRoleIdentityMapper {
		return vs[0].(map[string]*HardcodedRoleIdentityMapper)[vs[1].(string)]
	}).(HardcodedRoleIdentityMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleIdentityMapperInput)(nil)).Elem(), &HardcodedRoleIdentityMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleIdentityMapperArrayInput)(nil)).Elem(), HardcodedRoleIdentityMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleIdentityMapperMapInput)(nil)).Elem(), HardcodedRoleIdentityMapperMap{})
	pulumi.RegisterOutputType(HardcodedRoleIdentityMapperOutput{})
	pulumi.RegisterOutputType(HardcodedRoleIdentityMapperArrayOutput{})
	pulumi.RegisterOutputType(HardcodedRoleIdentityMapperMapOutput{})
}
