// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	// The provider configuration handed over to your custom user federation provider.
	Config pulumi.MapOutput `pulumi:"config"`
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Display name of the provider when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// Must be set to the realms' `internalId`  when it differs from the realm. This can happen when existing resources are imported into the state.
	ParentId pulumi.StringPtrOutput `pulumi:"parentId"`
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
	// The provider configuration handed over to your custom user federation provider.
	Config map[string]interface{} `pulumi:"config"`
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
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
	// The provider configuration handed over to your custom user federation provider.
	Config pulumi.MapInput
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
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
	// The provider configuration handed over to your custom user federation provider.
	Config map[string]interface{} `pulumi:"config"`
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
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
	// The provider configuration handed over to your custom user federation provider.
	Config pulumi.MapInput
	// When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
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
	return reflect.TypeOf((*CustomUserFederation)(nil))
}

func (i *CustomUserFederation) ToCustomUserFederationOutput() CustomUserFederationOutput {
	return i.ToCustomUserFederationOutputWithContext(context.Background())
}

func (i *CustomUserFederation) ToCustomUserFederationOutputWithContext(ctx context.Context) CustomUserFederationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomUserFederationOutput)
}

func (i *CustomUserFederation) ToCustomUserFederationPtrOutput() CustomUserFederationPtrOutput {
	return i.ToCustomUserFederationPtrOutputWithContext(context.Background())
}

func (i *CustomUserFederation) ToCustomUserFederationPtrOutputWithContext(ctx context.Context) CustomUserFederationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomUserFederationPtrOutput)
}

type CustomUserFederationPtrInput interface {
	pulumi.Input

	ToCustomUserFederationPtrOutput() CustomUserFederationPtrOutput
	ToCustomUserFederationPtrOutputWithContext(ctx context.Context) CustomUserFederationPtrOutput
}

type customUserFederationPtrType CustomUserFederationArgs

func (*customUserFederationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomUserFederation)(nil))
}

func (i *customUserFederationPtrType) ToCustomUserFederationPtrOutput() CustomUserFederationPtrOutput {
	return i.ToCustomUserFederationPtrOutputWithContext(context.Background())
}

func (i *customUserFederationPtrType) ToCustomUserFederationPtrOutputWithContext(ctx context.Context) CustomUserFederationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomUserFederationPtrOutput)
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
	return reflect.TypeOf(([]*CustomUserFederation)(nil))
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
	return reflect.TypeOf((map[string]*CustomUserFederation)(nil))
}

func (i CustomUserFederationMap) ToCustomUserFederationMapOutput() CustomUserFederationMapOutput {
	return i.ToCustomUserFederationMapOutputWithContext(context.Background())
}

func (i CustomUserFederationMap) ToCustomUserFederationMapOutputWithContext(ctx context.Context) CustomUserFederationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomUserFederationMapOutput)
}

type CustomUserFederationOutput struct {
	*pulumi.OutputState
}

func (CustomUserFederationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomUserFederation)(nil))
}

func (o CustomUserFederationOutput) ToCustomUserFederationOutput() CustomUserFederationOutput {
	return o
}

func (o CustomUserFederationOutput) ToCustomUserFederationOutputWithContext(ctx context.Context) CustomUserFederationOutput {
	return o
}

func (o CustomUserFederationOutput) ToCustomUserFederationPtrOutput() CustomUserFederationPtrOutput {
	return o.ToCustomUserFederationPtrOutputWithContext(context.Background())
}

func (o CustomUserFederationOutput) ToCustomUserFederationPtrOutputWithContext(ctx context.Context) CustomUserFederationPtrOutput {
	return o.ApplyT(func(v CustomUserFederation) *CustomUserFederation {
		return &v
	}).(CustomUserFederationPtrOutput)
}

type CustomUserFederationPtrOutput struct {
	*pulumi.OutputState
}

func (CustomUserFederationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomUserFederation)(nil))
}

func (o CustomUserFederationPtrOutput) ToCustomUserFederationPtrOutput() CustomUserFederationPtrOutput {
	return o
}

func (o CustomUserFederationPtrOutput) ToCustomUserFederationPtrOutputWithContext(ctx context.Context) CustomUserFederationPtrOutput {
	return o
}

type CustomUserFederationArrayOutput struct{ *pulumi.OutputState }

func (CustomUserFederationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomUserFederation)(nil))
}

func (o CustomUserFederationArrayOutput) ToCustomUserFederationArrayOutput() CustomUserFederationArrayOutput {
	return o
}

func (o CustomUserFederationArrayOutput) ToCustomUserFederationArrayOutputWithContext(ctx context.Context) CustomUserFederationArrayOutput {
	return o
}

func (o CustomUserFederationArrayOutput) Index(i pulumi.IntInput) CustomUserFederationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomUserFederation {
		return vs[0].([]CustomUserFederation)[vs[1].(int)]
	}).(CustomUserFederationOutput)
}

type CustomUserFederationMapOutput struct{ *pulumi.OutputState }

func (CustomUserFederationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomUserFederation)(nil))
}

func (o CustomUserFederationMapOutput) ToCustomUserFederationMapOutput() CustomUserFederationMapOutput {
	return o
}

func (o CustomUserFederationMapOutput) ToCustomUserFederationMapOutputWithContext(ctx context.Context) CustomUserFederationMapOutput {
	return o
}

func (o CustomUserFederationMapOutput) MapIndex(k pulumi.StringInput) CustomUserFederationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomUserFederation {
		return vs[0].(map[string]CustomUserFederation)[vs[1].(string)]
	}).(CustomUserFederationOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomUserFederationOutput{})
	pulumi.RegisterOutputType(CustomUserFederationPtrOutput{})
	pulumi.RegisterOutputType(CustomUserFederationArrayOutput{})
	pulumi.RegisterOutputType(CustomUserFederationMapOutput{})
}
