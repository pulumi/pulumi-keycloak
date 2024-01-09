// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows you to manage openid Client Authorization Permissions.
//
// ## Import
//
// Client authorization permissions can be imported using the format`{{realmId}}/{{resourceServerId}}/{{permissionId}}`. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission test my-realm/3bd4a686-1062-4b59-97b8-e4e3f10b99da/63b3cde8-987d-4cd9-9306-1955579281d9
//
// ```
type ClientAuthorizationPermission struct {
	pulumi.CustomResourceState

	DecisionStrategy pulumi.StringPtrOutput   `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	Policies         pulumi.StringArrayOutput `pulumi:"policies"`
	RealmId          pulumi.StringOutput      `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput      `pulumi:"resourceServerId"`
	ResourceType     pulumi.StringPtrOutput   `pulumi:"resourceType"`
	Resources        pulumi.StringArrayOutput `pulumi:"resources"`
	Scopes           pulumi.StringArrayOutput `pulumi:"scopes"`
	Type             pulumi.StringPtrOutput   `pulumi:"type"`
}

// NewClientAuthorizationPermission registers a new resource with the given unique name, arguments, and options.
func NewClientAuthorizationPermission(ctx *pulumi.Context,
	name string, args *ClientAuthorizationPermissionArgs, opts ...pulumi.ResourceOption) (*ClientAuthorizationPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.ResourceServerId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceServerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClientAuthorizationPermission
	err := ctx.RegisterResource("keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientAuthorizationPermission gets an existing ClientAuthorizationPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientAuthorizationPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientAuthorizationPermissionState, opts ...pulumi.ResourceOption) (*ClientAuthorizationPermission, error) {
	var resource ClientAuthorizationPermission
	err := ctx.ReadResource("keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientAuthorizationPermission resources.
type clientAuthorizationPermissionState struct {
	DecisionStrategy *string  `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Name             *string  `pulumi:"name"`
	Policies         []string `pulumi:"policies"`
	RealmId          *string  `pulumi:"realmId"`
	ResourceServerId *string  `pulumi:"resourceServerId"`
	ResourceType     *string  `pulumi:"resourceType"`
	Resources        []string `pulumi:"resources"`
	Scopes           []string `pulumi:"scopes"`
	Type             *string  `pulumi:"type"`
}

type ClientAuthorizationPermissionState struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Policies         pulumi.StringArrayInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
	ResourceType     pulumi.StringPtrInput
	Resources        pulumi.StringArrayInput
	Scopes           pulumi.StringArrayInput
	Type             pulumi.StringPtrInput
}

func (ClientAuthorizationPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationPermissionState)(nil)).Elem()
}

type clientAuthorizationPermissionArgs struct {
	DecisionStrategy *string  `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Name             *string  `pulumi:"name"`
	Policies         []string `pulumi:"policies"`
	RealmId          string   `pulumi:"realmId"`
	ResourceServerId string   `pulumi:"resourceServerId"`
	ResourceType     *string  `pulumi:"resourceType"`
	Resources        []string `pulumi:"resources"`
	Scopes           []string `pulumi:"scopes"`
	Type             *string  `pulumi:"type"`
}

// The set of arguments for constructing a ClientAuthorizationPermission resource.
type ClientAuthorizationPermissionArgs struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Policies         pulumi.StringArrayInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
	ResourceType     pulumi.StringPtrInput
	Resources        pulumi.StringArrayInput
	Scopes           pulumi.StringArrayInput
	Type             pulumi.StringPtrInput
}

func (ClientAuthorizationPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationPermissionArgs)(nil)).Elem()
}

type ClientAuthorizationPermissionInput interface {
	pulumi.Input

	ToClientAuthorizationPermissionOutput() ClientAuthorizationPermissionOutput
	ToClientAuthorizationPermissionOutputWithContext(ctx context.Context) ClientAuthorizationPermissionOutput
}

func (*ClientAuthorizationPermission) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorizationPermission)(nil)).Elem()
}

func (i *ClientAuthorizationPermission) ToClientAuthorizationPermissionOutput() ClientAuthorizationPermissionOutput {
	return i.ToClientAuthorizationPermissionOutputWithContext(context.Background())
}

func (i *ClientAuthorizationPermission) ToClientAuthorizationPermissionOutputWithContext(ctx context.Context) ClientAuthorizationPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationPermissionOutput)
}

// ClientAuthorizationPermissionArrayInput is an input type that accepts ClientAuthorizationPermissionArray and ClientAuthorizationPermissionArrayOutput values.
// You can construct a concrete instance of `ClientAuthorizationPermissionArrayInput` via:
//
//	ClientAuthorizationPermissionArray{ ClientAuthorizationPermissionArgs{...} }
type ClientAuthorizationPermissionArrayInput interface {
	pulumi.Input

	ToClientAuthorizationPermissionArrayOutput() ClientAuthorizationPermissionArrayOutput
	ToClientAuthorizationPermissionArrayOutputWithContext(context.Context) ClientAuthorizationPermissionArrayOutput
}

type ClientAuthorizationPermissionArray []ClientAuthorizationPermissionInput

func (ClientAuthorizationPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientAuthorizationPermission)(nil)).Elem()
}

func (i ClientAuthorizationPermissionArray) ToClientAuthorizationPermissionArrayOutput() ClientAuthorizationPermissionArrayOutput {
	return i.ToClientAuthorizationPermissionArrayOutputWithContext(context.Background())
}

func (i ClientAuthorizationPermissionArray) ToClientAuthorizationPermissionArrayOutputWithContext(ctx context.Context) ClientAuthorizationPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationPermissionArrayOutput)
}

// ClientAuthorizationPermissionMapInput is an input type that accepts ClientAuthorizationPermissionMap and ClientAuthorizationPermissionMapOutput values.
// You can construct a concrete instance of `ClientAuthorizationPermissionMapInput` via:
//
//	ClientAuthorizationPermissionMap{ "key": ClientAuthorizationPermissionArgs{...} }
type ClientAuthorizationPermissionMapInput interface {
	pulumi.Input

	ToClientAuthorizationPermissionMapOutput() ClientAuthorizationPermissionMapOutput
	ToClientAuthorizationPermissionMapOutputWithContext(context.Context) ClientAuthorizationPermissionMapOutput
}

type ClientAuthorizationPermissionMap map[string]ClientAuthorizationPermissionInput

func (ClientAuthorizationPermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientAuthorizationPermission)(nil)).Elem()
}

func (i ClientAuthorizationPermissionMap) ToClientAuthorizationPermissionMapOutput() ClientAuthorizationPermissionMapOutput {
	return i.ToClientAuthorizationPermissionMapOutputWithContext(context.Background())
}

func (i ClientAuthorizationPermissionMap) ToClientAuthorizationPermissionMapOutputWithContext(ctx context.Context) ClientAuthorizationPermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationPermissionMapOutput)
}

type ClientAuthorizationPermissionOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorizationPermission)(nil)).Elem()
}

func (o ClientAuthorizationPermissionOutput) ToClientAuthorizationPermissionOutput() ClientAuthorizationPermissionOutput {
	return o
}

func (o ClientAuthorizationPermissionOutput) ToClientAuthorizationPermissionOutputWithContext(ctx context.Context) ClientAuthorizationPermissionOutput {
	return o
}

func (o ClientAuthorizationPermissionOutput) DecisionStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringPtrOutput { return v.DecisionStrategy }).(pulumi.StringPtrOutput)
}

func (o ClientAuthorizationPermissionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ClientAuthorizationPermissionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClientAuthorizationPermissionOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

func (o ClientAuthorizationPermissionOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

func (o ClientAuthorizationPermissionOutput) ResourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringOutput { return v.ResourceServerId }).(pulumi.StringOutput)
}

func (o ClientAuthorizationPermissionOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringPtrOutput { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o ClientAuthorizationPermissionOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringArrayOutput { return v.Resources }).(pulumi.StringArrayOutput)
}

func (o ClientAuthorizationPermissionOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o ClientAuthorizationPermissionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthorizationPermission) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type ClientAuthorizationPermissionArrayOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientAuthorizationPermission)(nil)).Elem()
}

func (o ClientAuthorizationPermissionArrayOutput) ToClientAuthorizationPermissionArrayOutput() ClientAuthorizationPermissionArrayOutput {
	return o
}

func (o ClientAuthorizationPermissionArrayOutput) ToClientAuthorizationPermissionArrayOutputWithContext(ctx context.Context) ClientAuthorizationPermissionArrayOutput {
	return o
}

func (o ClientAuthorizationPermissionArrayOutput) Index(i pulumi.IntInput) ClientAuthorizationPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientAuthorizationPermission {
		return vs[0].([]*ClientAuthorizationPermission)[vs[1].(int)]
	}).(ClientAuthorizationPermissionOutput)
}

type ClientAuthorizationPermissionMapOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationPermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientAuthorizationPermission)(nil)).Elem()
}

func (o ClientAuthorizationPermissionMapOutput) ToClientAuthorizationPermissionMapOutput() ClientAuthorizationPermissionMapOutput {
	return o
}

func (o ClientAuthorizationPermissionMapOutput) ToClientAuthorizationPermissionMapOutputWithContext(ctx context.Context) ClientAuthorizationPermissionMapOutput {
	return o
}

func (o ClientAuthorizationPermissionMapOutput) MapIndex(k pulumi.StringInput) ClientAuthorizationPermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientAuthorizationPermission {
		return vs[0].(map[string]*ClientAuthorizationPermission)[vs[1].(string)]
	}).(ClientAuthorizationPermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientAuthorizationPermissionInput)(nil)).Elem(), &ClientAuthorizationPermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientAuthorizationPermissionArrayInput)(nil)).Elem(), ClientAuthorizationPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientAuthorizationPermissionMapInput)(nil)).Elem(), ClientAuthorizationPermissionMap{})
	pulumi.RegisterOutputType(ClientAuthorizationPermissionOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationPermissionArrayOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationPermissionMapOutput{})
}
