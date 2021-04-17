// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server.
type ClientDefaultScopes struct {
	pulumi.CustomResourceState

	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// An array of client scope names to attach to this client.
	DefaultScopes pulumi.StringArrayOutput `pulumi:"defaultScopes"`
	// The realm this client and scopes exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewClientDefaultScopes registers a new resource with the given unique name, arguments, and options.
func NewClientDefaultScopes(ctx *pulumi.Context,
	name string, args *ClientDefaultScopesArgs, opts ...pulumi.ResourceOption) (*ClientDefaultScopes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.DefaultScopes == nil {
		return nil, errors.New("invalid value for required argument 'DefaultScopes'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource ClientDefaultScopes
	err := ctx.RegisterResource("keycloak:openid/clientDefaultScopes:ClientDefaultScopes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientDefaultScopes gets an existing ClientDefaultScopes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientDefaultScopes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientDefaultScopesState, opts ...pulumi.ResourceOption) (*ClientDefaultScopes, error) {
	var resource ClientDefaultScopes
	err := ctx.ReadResource("keycloak:openid/clientDefaultScopes:ClientDefaultScopes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientDefaultScopes resources.
type clientDefaultScopesState struct {
	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId *string `pulumi:"clientId"`
	// An array of client scope names to attach to this client.
	DefaultScopes []string `pulumi:"defaultScopes"`
	// The realm this client and scopes exists in.
	RealmId *string `pulumi:"realmId"`
}

type ClientDefaultScopesState struct {
	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId pulumi.StringPtrInput
	// An array of client scope names to attach to this client.
	DefaultScopes pulumi.StringArrayInput
	// The realm this client and scopes exists in.
	RealmId pulumi.StringPtrInput
}

func (ClientDefaultScopesState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientDefaultScopesState)(nil)).Elem()
}

type clientDefaultScopesArgs struct {
	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId string `pulumi:"clientId"`
	// An array of client scope names to attach to this client.
	DefaultScopes []string `pulumi:"defaultScopes"`
	// The realm this client and scopes exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a ClientDefaultScopes resource.
type ClientDefaultScopesArgs struct {
	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId pulumi.StringInput
	// An array of client scope names to attach to this client.
	DefaultScopes pulumi.StringArrayInput
	// The realm this client and scopes exists in.
	RealmId pulumi.StringInput
}

func (ClientDefaultScopesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientDefaultScopesArgs)(nil)).Elem()
}

type ClientDefaultScopesInput interface {
	pulumi.Input

	ToClientDefaultScopesOutput() ClientDefaultScopesOutput
	ToClientDefaultScopesOutputWithContext(ctx context.Context) ClientDefaultScopesOutput
}

func (*ClientDefaultScopes) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientDefaultScopes)(nil))
}

func (i *ClientDefaultScopes) ToClientDefaultScopesOutput() ClientDefaultScopesOutput {
	return i.ToClientDefaultScopesOutputWithContext(context.Background())
}

func (i *ClientDefaultScopes) ToClientDefaultScopesOutputWithContext(ctx context.Context) ClientDefaultScopesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientDefaultScopesOutput)
}

func (i *ClientDefaultScopes) ToClientDefaultScopesPtrOutput() ClientDefaultScopesPtrOutput {
	return i.ToClientDefaultScopesPtrOutputWithContext(context.Background())
}

func (i *ClientDefaultScopes) ToClientDefaultScopesPtrOutputWithContext(ctx context.Context) ClientDefaultScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientDefaultScopesPtrOutput)
}

type ClientDefaultScopesPtrInput interface {
	pulumi.Input

	ToClientDefaultScopesPtrOutput() ClientDefaultScopesPtrOutput
	ToClientDefaultScopesPtrOutputWithContext(ctx context.Context) ClientDefaultScopesPtrOutput
}

type clientDefaultScopesPtrType ClientDefaultScopesArgs

func (*clientDefaultScopesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientDefaultScopes)(nil))
}

func (i *clientDefaultScopesPtrType) ToClientDefaultScopesPtrOutput() ClientDefaultScopesPtrOutput {
	return i.ToClientDefaultScopesPtrOutputWithContext(context.Background())
}

func (i *clientDefaultScopesPtrType) ToClientDefaultScopesPtrOutputWithContext(ctx context.Context) ClientDefaultScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientDefaultScopesPtrOutput)
}

// ClientDefaultScopesArrayInput is an input type that accepts ClientDefaultScopesArray and ClientDefaultScopesArrayOutput values.
// You can construct a concrete instance of `ClientDefaultScopesArrayInput` via:
//
//          ClientDefaultScopesArray{ ClientDefaultScopesArgs{...} }
type ClientDefaultScopesArrayInput interface {
	pulumi.Input

	ToClientDefaultScopesArrayOutput() ClientDefaultScopesArrayOutput
	ToClientDefaultScopesArrayOutputWithContext(context.Context) ClientDefaultScopesArrayOutput
}

type ClientDefaultScopesArray []ClientDefaultScopesInput

func (ClientDefaultScopesArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ClientDefaultScopes)(nil))
}

func (i ClientDefaultScopesArray) ToClientDefaultScopesArrayOutput() ClientDefaultScopesArrayOutput {
	return i.ToClientDefaultScopesArrayOutputWithContext(context.Background())
}

func (i ClientDefaultScopesArray) ToClientDefaultScopesArrayOutputWithContext(ctx context.Context) ClientDefaultScopesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientDefaultScopesArrayOutput)
}

// ClientDefaultScopesMapInput is an input type that accepts ClientDefaultScopesMap and ClientDefaultScopesMapOutput values.
// You can construct a concrete instance of `ClientDefaultScopesMapInput` via:
//
//          ClientDefaultScopesMap{ "key": ClientDefaultScopesArgs{...} }
type ClientDefaultScopesMapInput interface {
	pulumi.Input

	ToClientDefaultScopesMapOutput() ClientDefaultScopesMapOutput
	ToClientDefaultScopesMapOutputWithContext(context.Context) ClientDefaultScopesMapOutput
}

type ClientDefaultScopesMap map[string]ClientDefaultScopesInput

func (ClientDefaultScopesMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ClientDefaultScopes)(nil))
}

func (i ClientDefaultScopesMap) ToClientDefaultScopesMapOutput() ClientDefaultScopesMapOutput {
	return i.ToClientDefaultScopesMapOutputWithContext(context.Background())
}

func (i ClientDefaultScopesMap) ToClientDefaultScopesMapOutputWithContext(ctx context.Context) ClientDefaultScopesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientDefaultScopesMapOutput)
}

type ClientDefaultScopesOutput struct {
	*pulumi.OutputState
}

func (ClientDefaultScopesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientDefaultScopes)(nil))
}

func (o ClientDefaultScopesOutput) ToClientDefaultScopesOutput() ClientDefaultScopesOutput {
	return o
}

func (o ClientDefaultScopesOutput) ToClientDefaultScopesOutputWithContext(ctx context.Context) ClientDefaultScopesOutput {
	return o
}

func (o ClientDefaultScopesOutput) ToClientDefaultScopesPtrOutput() ClientDefaultScopesPtrOutput {
	return o.ToClientDefaultScopesPtrOutputWithContext(context.Background())
}

func (o ClientDefaultScopesOutput) ToClientDefaultScopesPtrOutputWithContext(ctx context.Context) ClientDefaultScopesPtrOutput {
	return o.ApplyT(func(v ClientDefaultScopes) *ClientDefaultScopes {
		return &v
	}).(ClientDefaultScopesPtrOutput)
}

type ClientDefaultScopesPtrOutput struct {
	*pulumi.OutputState
}

func (ClientDefaultScopesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientDefaultScopes)(nil))
}

func (o ClientDefaultScopesPtrOutput) ToClientDefaultScopesPtrOutput() ClientDefaultScopesPtrOutput {
	return o
}

func (o ClientDefaultScopesPtrOutput) ToClientDefaultScopesPtrOutputWithContext(ctx context.Context) ClientDefaultScopesPtrOutput {
	return o
}

type ClientDefaultScopesArrayOutput struct{ *pulumi.OutputState }

func (ClientDefaultScopesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientDefaultScopes)(nil))
}

func (o ClientDefaultScopesArrayOutput) ToClientDefaultScopesArrayOutput() ClientDefaultScopesArrayOutput {
	return o
}

func (o ClientDefaultScopesArrayOutput) ToClientDefaultScopesArrayOutputWithContext(ctx context.Context) ClientDefaultScopesArrayOutput {
	return o
}

func (o ClientDefaultScopesArrayOutput) Index(i pulumi.IntInput) ClientDefaultScopesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientDefaultScopes {
		return vs[0].([]ClientDefaultScopes)[vs[1].(int)]
	}).(ClientDefaultScopesOutput)
}

type ClientDefaultScopesMapOutput struct{ *pulumi.OutputState }

func (ClientDefaultScopesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClientDefaultScopes)(nil))
}

func (o ClientDefaultScopesMapOutput) ToClientDefaultScopesMapOutput() ClientDefaultScopesMapOutput {
	return o
}

func (o ClientDefaultScopesMapOutput) ToClientDefaultScopesMapOutputWithContext(ctx context.Context) ClientDefaultScopesMapOutput {
	return o
}

func (o ClientDefaultScopesMapOutput) MapIndex(k pulumi.StringInput) ClientDefaultScopesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClientDefaultScopes {
		return vs[0].(map[string]ClientDefaultScopes)[vs[1].(string)]
	}).(ClientDefaultScopesOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientDefaultScopesOutput{})
	pulumi.RegisterOutputType(ClientDefaultScopesPtrOutput{})
	pulumi.RegisterOutputType(ClientDefaultScopesArrayOutput{})
	pulumi.RegisterOutputType(ClientDefaultScopesMapOutput{})
}
