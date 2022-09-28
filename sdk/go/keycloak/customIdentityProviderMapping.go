// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/oidc"
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
//			_, err = keycloak.NewCustomIdentityProviderMapping(ctx, "oidcCustomIdentityProviderMapping", &keycloak.CustomIdentityProviderMappingArgs{
//				Realm:                  realm.ID(),
//				IdentityProviderAlias:  oidcIdentityProvider.Alias,
//				IdentityProviderMapper: pulumi.String(fmt.Sprintf("%vs-user-attribute-idp-mapper", "%")),
//				ExtraConfig: pulumi.AnyMap{
//					"syncMode":      pulumi.Any("INHERIT"),
//					"Claim":         pulumi.Any("my-email-claim"),
//					"UserAttribute": pulumi.Any("email"),
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
//
// ## Import
//
// Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping test_mapper my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
//
// ```
type CustomIdentityProviderMapping struct {
	pulumi.CustomResourceState

	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapOutput `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
	IdentityProviderMapper pulumi.StringOutput `pulumi:"identityProviderMapper"`
	// The name of the mapper.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the realm.
	Realm pulumi.StringOutput `pulumi:"realm"`
}

// NewCustomIdentityProviderMapping registers a new resource with the given unique name, arguments, and options.
func NewCustomIdentityProviderMapping(ctx *pulumi.Context,
	name string, args *CustomIdentityProviderMappingArgs, opts ...pulumi.ResourceOption) (*CustomIdentityProviderMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityProviderAlias == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderAlias'")
	}
	if args.IdentityProviderMapper == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderMapper'")
	}
	if args.Realm == nil {
		return nil, errors.New("invalid value for required argument 'Realm'")
	}
	var resource CustomIdentityProviderMapping
	err := ctx.RegisterResource("keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomIdentityProviderMapping gets an existing CustomIdentityProviderMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomIdentityProviderMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomIdentityProviderMappingState, opts ...pulumi.ResourceOption) (*CustomIdentityProviderMapping, error) {
	var resource CustomIdentityProviderMapping
	err := ctx.ReadResource("keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomIdentityProviderMapping resources.
type customIdentityProviderMappingState struct {
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
	IdentityProviderMapper *string `pulumi:"identityProviderMapper"`
	// The name of the mapper.
	Name *string `pulumi:"name"`
	// The name of the realm.
	Realm *string `pulumi:"realm"`
}

type CustomIdentityProviderMappingState struct {
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapInput
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringPtrInput
	// The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
	IdentityProviderMapper pulumi.StringPtrInput
	// The name of the mapper.
	Name pulumi.StringPtrInput
	// The name of the realm.
	Realm pulumi.StringPtrInput
}

func (CustomIdentityProviderMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*customIdentityProviderMappingState)(nil)).Elem()
}

type customIdentityProviderMappingArgs struct {
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig map[string]interface{} `pulumi:"extraConfig"`
	// The alias of the associated identity provider.
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
	IdentityProviderMapper string `pulumi:"identityProviderMapper"`
	// The name of the mapper.
	Name *string `pulumi:"name"`
	// The name of the realm.
	Realm string `pulumi:"realm"`
}

// The set of arguments for constructing a CustomIdentityProviderMapping resource.
type CustomIdentityProviderMappingArgs struct {
	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	ExtraConfig pulumi.MapInput
	// The alias of the associated identity provider.
	IdentityProviderAlias pulumi.StringInput
	// The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
	IdentityProviderMapper pulumi.StringInput
	// The name of the mapper.
	Name pulumi.StringPtrInput
	// The name of the realm.
	Realm pulumi.StringInput
}

func (CustomIdentityProviderMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customIdentityProviderMappingArgs)(nil)).Elem()
}

type CustomIdentityProviderMappingInput interface {
	pulumi.Input

	ToCustomIdentityProviderMappingOutput() CustomIdentityProviderMappingOutput
	ToCustomIdentityProviderMappingOutputWithContext(ctx context.Context) CustomIdentityProviderMappingOutput
}

func (*CustomIdentityProviderMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomIdentityProviderMapping)(nil)).Elem()
}

func (i *CustomIdentityProviderMapping) ToCustomIdentityProviderMappingOutput() CustomIdentityProviderMappingOutput {
	return i.ToCustomIdentityProviderMappingOutputWithContext(context.Background())
}

func (i *CustomIdentityProviderMapping) ToCustomIdentityProviderMappingOutputWithContext(ctx context.Context) CustomIdentityProviderMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomIdentityProviderMappingOutput)
}

// CustomIdentityProviderMappingArrayInput is an input type that accepts CustomIdentityProviderMappingArray and CustomIdentityProviderMappingArrayOutput values.
// You can construct a concrete instance of `CustomIdentityProviderMappingArrayInput` via:
//
//	CustomIdentityProviderMappingArray{ CustomIdentityProviderMappingArgs{...} }
type CustomIdentityProviderMappingArrayInput interface {
	pulumi.Input

	ToCustomIdentityProviderMappingArrayOutput() CustomIdentityProviderMappingArrayOutput
	ToCustomIdentityProviderMappingArrayOutputWithContext(context.Context) CustomIdentityProviderMappingArrayOutput
}

type CustomIdentityProviderMappingArray []CustomIdentityProviderMappingInput

func (CustomIdentityProviderMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomIdentityProviderMapping)(nil)).Elem()
}

func (i CustomIdentityProviderMappingArray) ToCustomIdentityProviderMappingArrayOutput() CustomIdentityProviderMappingArrayOutput {
	return i.ToCustomIdentityProviderMappingArrayOutputWithContext(context.Background())
}

func (i CustomIdentityProviderMappingArray) ToCustomIdentityProviderMappingArrayOutputWithContext(ctx context.Context) CustomIdentityProviderMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomIdentityProviderMappingArrayOutput)
}

// CustomIdentityProviderMappingMapInput is an input type that accepts CustomIdentityProviderMappingMap and CustomIdentityProviderMappingMapOutput values.
// You can construct a concrete instance of `CustomIdentityProviderMappingMapInput` via:
//
//	CustomIdentityProviderMappingMap{ "key": CustomIdentityProviderMappingArgs{...} }
type CustomIdentityProviderMappingMapInput interface {
	pulumi.Input

	ToCustomIdentityProviderMappingMapOutput() CustomIdentityProviderMappingMapOutput
	ToCustomIdentityProviderMappingMapOutputWithContext(context.Context) CustomIdentityProviderMappingMapOutput
}

type CustomIdentityProviderMappingMap map[string]CustomIdentityProviderMappingInput

func (CustomIdentityProviderMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomIdentityProviderMapping)(nil)).Elem()
}

func (i CustomIdentityProviderMappingMap) ToCustomIdentityProviderMappingMapOutput() CustomIdentityProviderMappingMapOutput {
	return i.ToCustomIdentityProviderMappingMapOutputWithContext(context.Background())
}

func (i CustomIdentityProviderMappingMap) ToCustomIdentityProviderMappingMapOutputWithContext(ctx context.Context) CustomIdentityProviderMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomIdentityProviderMappingMapOutput)
}

type CustomIdentityProviderMappingOutput struct{ *pulumi.OutputState }

func (CustomIdentityProviderMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomIdentityProviderMapping)(nil)).Elem()
}

func (o CustomIdentityProviderMappingOutput) ToCustomIdentityProviderMappingOutput() CustomIdentityProviderMappingOutput {
	return o
}

func (o CustomIdentityProviderMappingOutput) ToCustomIdentityProviderMappingOutputWithContext(ctx context.Context) CustomIdentityProviderMappingOutput {
	return o
}

// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
func (o CustomIdentityProviderMappingOutput) ExtraConfig() pulumi.MapOutput {
	return o.ApplyT(func(v *CustomIdentityProviderMapping) pulumi.MapOutput { return v.ExtraConfig }).(pulumi.MapOutput)
}

// The alias of the associated identity provider.
func (o CustomIdentityProviderMappingOutput) IdentityProviderAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIdentityProviderMapping) pulumi.StringOutput { return v.IdentityProviderAlias }).(pulumi.StringOutput)
}

// The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
func (o CustomIdentityProviderMappingOutput) IdentityProviderMapper() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIdentityProviderMapping) pulumi.StringOutput { return v.IdentityProviderMapper }).(pulumi.StringOutput)
}

// The name of the mapper.
func (o CustomIdentityProviderMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIdentityProviderMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the realm.
func (o CustomIdentityProviderMappingOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomIdentityProviderMapping) pulumi.StringOutput { return v.Realm }).(pulumi.StringOutput)
}

type CustomIdentityProviderMappingArrayOutput struct{ *pulumi.OutputState }

func (CustomIdentityProviderMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomIdentityProviderMapping)(nil)).Elem()
}

func (o CustomIdentityProviderMappingArrayOutput) ToCustomIdentityProviderMappingArrayOutput() CustomIdentityProviderMappingArrayOutput {
	return o
}

func (o CustomIdentityProviderMappingArrayOutput) ToCustomIdentityProviderMappingArrayOutputWithContext(ctx context.Context) CustomIdentityProviderMappingArrayOutput {
	return o
}

func (o CustomIdentityProviderMappingArrayOutput) Index(i pulumi.IntInput) CustomIdentityProviderMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomIdentityProviderMapping {
		return vs[0].([]*CustomIdentityProviderMapping)[vs[1].(int)]
	}).(CustomIdentityProviderMappingOutput)
}

type CustomIdentityProviderMappingMapOutput struct{ *pulumi.OutputState }

func (CustomIdentityProviderMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomIdentityProviderMapping)(nil)).Elem()
}

func (o CustomIdentityProviderMappingMapOutput) ToCustomIdentityProviderMappingMapOutput() CustomIdentityProviderMappingMapOutput {
	return o
}

func (o CustomIdentityProviderMappingMapOutput) ToCustomIdentityProviderMappingMapOutputWithContext(ctx context.Context) CustomIdentityProviderMappingMapOutput {
	return o
}

func (o CustomIdentityProviderMappingMapOutput) MapIndex(k pulumi.StringInput) CustomIdentityProviderMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomIdentityProviderMapping {
		return vs[0].(map[string]*CustomIdentityProviderMapping)[vs[1].(string)]
	}).(CustomIdentityProviderMappingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomIdentityProviderMappingInput)(nil)).Elem(), &CustomIdentityProviderMapping{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomIdentityProviderMappingArrayInput)(nil)).Elem(), CustomIdentityProviderMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomIdentityProviderMappingMapInput)(nil)).Elem(), CustomIdentityProviderMappingMap{})
	pulumi.RegisterOutputType(CustomIdentityProviderMappingOutput{})
	pulumi.RegisterOutputType(CustomIdentityProviderMappingArrayOutput{})
	pulumi.RegisterOutputType(CustomIdentityProviderMappingMapOutput{})
}
