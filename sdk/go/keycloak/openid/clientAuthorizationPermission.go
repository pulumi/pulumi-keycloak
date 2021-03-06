// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClientAuthorizationPermission struct {
	pulumi.CustomResourceState

	DecisionStrategy pulumi.StringPtrOutput   `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	Policies         pulumi.StringArrayOutput `pulumi:"policies"`
	RealmId          pulumi.StringOutput      `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput      `pulumi:"resourceServerId"`
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
	return reflect.TypeOf((*ClientAuthorizationPermission)(nil))
}

func (i *ClientAuthorizationPermission) ToClientAuthorizationPermissionOutput() ClientAuthorizationPermissionOutput {
	return i.ToClientAuthorizationPermissionOutputWithContext(context.Background())
}

func (i *ClientAuthorizationPermission) ToClientAuthorizationPermissionOutputWithContext(ctx context.Context) ClientAuthorizationPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationPermissionOutput)
}

func (i *ClientAuthorizationPermission) ToClientAuthorizationPermissionPtrOutput() ClientAuthorizationPermissionPtrOutput {
	return i.ToClientAuthorizationPermissionPtrOutputWithContext(context.Background())
}

func (i *ClientAuthorizationPermission) ToClientAuthorizationPermissionPtrOutputWithContext(ctx context.Context) ClientAuthorizationPermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationPermissionPtrOutput)
}

type ClientAuthorizationPermissionPtrInput interface {
	pulumi.Input

	ToClientAuthorizationPermissionPtrOutput() ClientAuthorizationPermissionPtrOutput
	ToClientAuthorizationPermissionPtrOutputWithContext(ctx context.Context) ClientAuthorizationPermissionPtrOutput
}

type clientAuthorizationPermissionPtrType ClientAuthorizationPermissionArgs

func (*clientAuthorizationPermissionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorizationPermission)(nil))
}

func (i *clientAuthorizationPermissionPtrType) ToClientAuthorizationPermissionPtrOutput() ClientAuthorizationPermissionPtrOutput {
	return i.ToClientAuthorizationPermissionPtrOutputWithContext(context.Background())
}

func (i *clientAuthorizationPermissionPtrType) ToClientAuthorizationPermissionPtrOutputWithContext(ctx context.Context) ClientAuthorizationPermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationPermissionPtrOutput)
}

// ClientAuthorizationPermissionArrayInput is an input type that accepts ClientAuthorizationPermissionArray and ClientAuthorizationPermissionArrayOutput values.
// You can construct a concrete instance of `ClientAuthorizationPermissionArrayInput` via:
//
//          ClientAuthorizationPermissionArray{ ClientAuthorizationPermissionArgs{...} }
type ClientAuthorizationPermissionArrayInput interface {
	pulumi.Input

	ToClientAuthorizationPermissionArrayOutput() ClientAuthorizationPermissionArrayOutput
	ToClientAuthorizationPermissionArrayOutputWithContext(context.Context) ClientAuthorizationPermissionArrayOutput
}

type ClientAuthorizationPermissionArray []ClientAuthorizationPermissionInput

func (ClientAuthorizationPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ClientAuthorizationPermission)(nil))
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
//          ClientAuthorizationPermissionMap{ "key": ClientAuthorizationPermissionArgs{...} }
type ClientAuthorizationPermissionMapInput interface {
	pulumi.Input

	ToClientAuthorizationPermissionMapOutput() ClientAuthorizationPermissionMapOutput
	ToClientAuthorizationPermissionMapOutputWithContext(context.Context) ClientAuthorizationPermissionMapOutput
}

type ClientAuthorizationPermissionMap map[string]ClientAuthorizationPermissionInput

func (ClientAuthorizationPermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ClientAuthorizationPermission)(nil))
}

func (i ClientAuthorizationPermissionMap) ToClientAuthorizationPermissionMapOutput() ClientAuthorizationPermissionMapOutput {
	return i.ToClientAuthorizationPermissionMapOutputWithContext(context.Background())
}

func (i ClientAuthorizationPermissionMap) ToClientAuthorizationPermissionMapOutputWithContext(ctx context.Context) ClientAuthorizationPermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationPermissionMapOutput)
}

type ClientAuthorizationPermissionOutput struct {
	*pulumi.OutputState
}

func (ClientAuthorizationPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthorizationPermission)(nil))
}

func (o ClientAuthorizationPermissionOutput) ToClientAuthorizationPermissionOutput() ClientAuthorizationPermissionOutput {
	return o
}

func (o ClientAuthorizationPermissionOutput) ToClientAuthorizationPermissionOutputWithContext(ctx context.Context) ClientAuthorizationPermissionOutput {
	return o
}

func (o ClientAuthorizationPermissionOutput) ToClientAuthorizationPermissionPtrOutput() ClientAuthorizationPermissionPtrOutput {
	return o.ToClientAuthorizationPermissionPtrOutputWithContext(context.Background())
}

func (o ClientAuthorizationPermissionOutput) ToClientAuthorizationPermissionPtrOutputWithContext(ctx context.Context) ClientAuthorizationPermissionPtrOutput {
	return o.ApplyT(func(v ClientAuthorizationPermission) *ClientAuthorizationPermission {
		return &v
	}).(ClientAuthorizationPermissionPtrOutput)
}

type ClientAuthorizationPermissionPtrOutput struct {
	*pulumi.OutputState
}

func (ClientAuthorizationPermissionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorizationPermission)(nil))
}

func (o ClientAuthorizationPermissionPtrOutput) ToClientAuthorizationPermissionPtrOutput() ClientAuthorizationPermissionPtrOutput {
	return o
}

func (o ClientAuthorizationPermissionPtrOutput) ToClientAuthorizationPermissionPtrOutputWithContext(ctx context.Context) ClientAuthorizationPermissionPtrOutput {
	return o
}

type ClientAuthorizationPermissionArrayOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientAuthorizationPermission)(nil))
}

func (o ClientAuthorizationPermissionArrayOutput) ToClientAuthorizationPermissionArrayOutput() ClientAuthorizationPermissionArrayOutput {
	return o
}

func (o ClientAuthorizationPermissionArrayOutput) ToClientAuthorizationPermissionArrayOutputWithContext(ctx context.Context) ClientAuthorizationPermissionArrayOutput {
	return o
}

func (o ClientAuthorizationPermissionArrayOutput) Index(i pulumi.IntInput) ClientAuthorizationPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientAuthorizationPermission {
		return vs[0].([]ClientAuthorizationPermission)[vs[1].(int)]
	}).(ClientAuthorizationPermissionOutput)
}

type ClientAuthorizationPermissionMapOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationPermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClientAuthorizationPermission)(nil))
}

func (o ClientAuthorizationPermissionMapOutput) ToClientAuthorizationPermissionMapOutput() ClientAuthorizationPermissionMapOutput {
	return o
}

func (o ClientAuthorizationPermissionMapOutput) ToClientAuthorizationPermissionMapOutputWithContext(ctx context.Context) ClientAuthorizationPermissionMapOutput {
	return o
}

func (o ClientAuthorizationPermissionMapOutput) MapIndex(k pulumi.StringInput) ClientAuthorizationPermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClientAuthorizationPermission {
		return vs[0].(map[string]ClientAuthorizationPermission)[vs[1].(string)]
	}).(ClientAuthorizationPermissionOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientAuthorizationPermissionOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationPermissionPtrOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationPermissionArrayOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationPermissionMapOutput{})
}
