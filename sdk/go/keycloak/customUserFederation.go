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
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("test"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewCustomUserFederation(ctx, "customUserFederation", &keycloak.CustomUserFederationArgs{
// 			RealmId:    realm.ID(),
// 			ProviderId: pulumi.String("custom"),
// 			Enabled:    pulumi.Bool(true),
// 			Config: pulumi.AnyMap{
// 				"dummyString": pulumi.Any("foobar"),
// 				"dummyBool":   pulumi.Any(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Custom user federation providers can be imported using the format `{{realm_id}}/{{custom_user_federation_id}}`. The ID of the custom user federation provider can be found within the Keycloak GUI and is typically a GUIDbash
//
// ```sh
//  $ pulumi import keycloak:index/customUserFederation:CustomUserFederation custom_user_federation my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860
// ```
type CustomUserFederation struct {
	pulumi.CustomResourceState

	// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	CachePolicy pulumi.StringPtrOutput `pulumi:"cachePolicy"`
	// How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod pulumi.IntPtrOutput `pulumi:"changedSyncPeriod"`
	// The provider configuration handed over to your custom user federation provider.
	Config pulumi.MapOutput `pulumi:"config"`
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod pulumi.IntPtrOutput `pulumi:"fullSyncPeriod"`
	// Display name of the provider when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
	ProviderId pulumi.StringOutput `pulumi:"providerId"`
	// The realm that this provider will provide user federation for.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewCustomUserFederation registers a new resource with the given unique name, arguments, and options.
func NewCustomUserFederation(ctx *pulumi.Context,
	name string, args *CustomUserFederationArgs, opts ...pulumi.ResourceOption) (*CustomUserFederation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderId == nil {
		return nil, errors.New("invalid value for required argument 'ProviderId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource CustomUserFederation
	err := ctx.RegisterResource("keycloak:index/customUserFederation:CustomUserFederation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomUserFederation gets an existing CustomUserFederation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomUserFederation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomUserFederationState, opts ...pulumi.ResourceOption) (*CustomUserFederation, error) {
	var resource CustomUserFederation
	err := ctx.ReadResource("keycloak:index/customUserFederation:CustomUserFederation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomUserFederation resources.
type customUserFederationState struct {
	// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	CachePolicy *string `pulumi:"cachePolicy"`
	// How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod *int `pulumi:"changedSyncPeriod"`
	// The provider configuration handed over to your custom user federation provider.
	Config map[string]interface{} `pulumi:"config"`
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod *int `pulumi:"fullSyncPeriod"`
	// Display name of the provider when displayed in the console.
	Name *string `pulumi:"name"`
	// Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
	ParentId *string `pulumi:"parentId"`
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority *int `pulumi:"priority"`
	// The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
	ProviderId *string `pulumi:"providerId"`
	// The realm that this provider will provide user federation for.
	RealmId *string `pulumi:"realmId"`
}

type CustomUserFederationState struct {
	// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	CachePolicy pulumi.StringPtrInput
	// How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod pulumi.IntPtrInput
	// The provider configuration handed over to your custom user federation provider.
	Config pulumi.MapInput
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod pulumi.IntPtrInput
	// Display name of the provider when displayed in the console.
	Name pulumi.StringPtrInput
	// Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
	ParentId pulumi.StringPtrInput
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority pulumi.IntPtrInput
	// The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
	ProviderId pulumi.StringPtrInput
	// The realm that this provider will provide user federation for.
	RealmId pulumi.StringPtrInput
}

func (CustomUserFederationState) ElementType() reflect.Type {
	return reflect.TypeOf((*customUserFederationState)(nil)).Elem()
}

type customUserFederationArgs struct {
	// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	CachePolicy *string `pulumi:"cachePolicy"`
	// How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod *int `pulumi:"changedSyncPeriod"`
	// The provider configuration handed over to your custom user federation provider.
	Config map[string]interface{} `pulumi:"config"`
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod *int `pulumi:"fullSyncPeriod"`
	// Display name of the provider when displayed in the console.
	Name *string `pulumi:"name"`
	// Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
	ParentId *string `pulumi:"parentId"`
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority *int `pulumi:"priority"`
	// The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
	ProviderId string `pulumi:"providerId"`
	// The realm that this provider will provide user federation for.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a CustomUserFederation resource.
type CustomUserFederationArgs struct {
	// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	CachePolicy pulumi.StringPtrInput
	// How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod pulumi.IntPtrInput
	// The provider configuration handed over to your custom user federation provider.
	Config pulumi.MapInput
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod pulumi.IntPtrInput
	// Display name of the provider when displayed in the console.
	Name pulumi.StringPtrInput
	// Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
	ParentId pulumi.StringPtrInput
	// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
	Priority pulumi.IntPtrInput
	// The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
	ProviderId pulumi.StringInput
	// The realm that this provider will provide user federation for.
	RealmId pulumi.StringInput
}

func (CustomUserFederationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customUserFederationArgs)(nil)).Elem()
}

type CustomUserFederationInput interface {
	pulumi.Input

	ToCustomUserFederationOutput() CustomUserFederationOutput
	ToCustomUserFederationOutputWithContext(ctx context.Context) CustomUserFederationOutput
}

func (*CustomUserFederation) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomUserFederation)(nil)).Elem()
}

func (i *CustomUserFederation) ToCustomUserFederationOutput() CustomUserFederationOutput {
	return i.ToCustomUserFederationOutputWithContext(context.Background())
}

func (i *CustomUserFederation) ToCustomUserFederationOutputWithContext(ctx context.Context) CustomUserFederationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomUserFederationOutput)
}

// CustomUserFederationArrayInput is an input type that accepts CustomUserFederationArray and CustomUserFederationArrayOutput values.
// You can construct a concrete instance of `CustomUserFederationArrayInput` via:
//
//          CustomUserFederationArray{ CustomUserFederationArgs{...} }
type CustomUserFederationArrayInput interface {
	pulumi.Input

	ToCustomUserFederationArrayOutput() CustomUserFederationArrayOutput
	ToCustomUserFederationArrayOutputWithContext(context.Context) CustomUserFederationArrayOutput
}

type CustomUserFederationArray []CustomUserFederationInput

func (CustomUserFederationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomUserFederation)(nil)).Elem()
}

func (i CustomUserFederationArray) ToCustomUserFederationArrayOutput() CustomUserFederationArrayOutput {
	return i.ToCustomUserFederationArrayOutputWithContext(context.Background())
}

func (i CustomUserFederationArray) ToCustomUserFederationArrayOutputWithContext(ctx context.Context) CustomUserFederationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomUserFederationArrayOutput)
}

// CustomUserFederationMapInput is an input type that accepts CustomUserFederationMap and CustomUserFederationMapOutput values.
// You can construct a concrete instance of `CustomUserFederationMapInput` via:
//
//          CustomUserFederationMap{ "key": CustomUserFederationArgs{...} }
type CustomUserFederationMapInput interface {
	pulumi.Input

	ToCustomUserFederationMapOutput() CustomUserFederationMapOutput
	ToCustomUserFederationMapOutputWithContext(context.Context) CustomUserFederationMapOutput
}

type CustomUserFederationMap map[string]CustomUserFederationInput

func (CustomUserFederationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomUserFederation)(nil)).Elem()
}

func (i CustomUserFederationMap) ToCustomUserFederationMapOutput() CustomUserFederationMapOutput {
	return i.ToCustomUserFederationMapOutputWithContext(context.Background())
}

func (i CustomUserFederationMap) ToCustomUserFederationMapOutputWithContext(ctx context.Context) CustomUserFederationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomUserFederationMapOutput)
}

type CustomUserFederationOutput struct{ *pulumi.OutputState }

func (CustomUserFederationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomUserFederation)(nil)).Elem()
}

func (o CustomUserFederationOutput) ToCustomUserFederationOutput() CustomUserFederationOutput {
	return o
}

func (o CustomUserFederationOutput) ToCustomUserFederationOutputWithContext(ctx context.Context) CustomUserFederationOutput {
	return o
}

// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
func (o CustomUserFederationOutput) CachePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.StringPtrOutput { return v.CachePolicy }).(pulumi.StringPtrOutput)
}

// How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users sync.
func (o CustomUserFederationOutput) ChangedSyncPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.IntPtrOutput { return v.ChangedSyncPeriod }).(pulumi.IntPtrOutput)
}

// The provider configuration handed over to your custom user federation provider.
func (o CustomUserFederationOutput) Config() pulumi.MapOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.MapOutput { return v.Config }).(pulumi.MapOutput)
}

// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
func (o CustomUserFederationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
func (o CustomUserFederationOutput) FullSyncPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.IntPtrOutput { return v.FullSyncPeriod }).(pulumi.IntPtrOutput)
}

// Display name of the provider when displayed in the console.
func (o CustomUserFederationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
func (o CustomUserFederationOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

// Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
func (o CustomUserFederationOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
func (o CustomUserFederationOutput) ProviderId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.StringOutput { return v.ProviderId }).(pulumi.StringOutput)
}

// The realm that this provider will provide user federation for.
func (o CustomUserFederationOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomUserFederation) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type CustomUserFederationArrayOutput struct{ *pulumi.OutputState }

func (CustomUserFederationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomUserFederation)(nil)).Elem()
}

func (o CustomUserFederationArrayOutput) ToCustomUserFederationArrayOutput() CustomUserFederationArrayOutput {
	return o
}

func (o CustomUserFederationArrayOutput) ToCustomUserFederationArrayOutputWithContext(ctx context.Context) CustomUserFederationArrayOutput {
	return o
}

func (o CustomUserFederationArrayOutput) Index(i pulumi.IntInput) CustomUserFederationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomUserFederation {
		return vs[0].([]*CustomUserFederation)[vs[1].(int)]
	}).(CustomUserFederationOutput)
}

type CustomUserFederationMapOutput struct{ *pulumi.OutputState }

func (CustomUserFederationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomUserFederation)(nil)).Elem()
}

func (o CustomUserFederationMapOutput) ToCustomUserFederationMapOutput() CustomUserFederationMapOutput {
	return o
}

func (o CustomUserFederationMapOutput) ToCustomUserFederationMapOutputWithContext(ctx context.Context) CustomUserFederationMapOutput {
	return o
}

func (o CustomUserFederationMapOutput) MapIndex(k pulumi.StringInput) CustomUserFederationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomUserFederation {
		return vs[0].(map[string]*CustomUserFederation)[vs[1].(string)]
	}).(CustomUserFederationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomUserFederationInput)(nil)).Elem(), &CustomUserFederation{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomUserFederationArrayInput)(nil)).Elem(), CustomUserFederationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomUserFederationMapInput)(nil)).Elem(), CustomUserFederationMap{})
	pulumi.RegisterOutputType(CustomUserFederationOutput{})
	pulumi.RegisterOutputType(CustomUserFederationArrayOutput{})
	pulumi.RegisterOutputType(CustomUserFederationMapOutput{})
}
