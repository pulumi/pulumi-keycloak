// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type ClientPermissions struct {
	pulumi.CustomResourceState

	// Resource server id representing the realm management client on which this permission is managed
	AuthorizationResourceServerId pulumi.StringOutput                                `pulumi:"authorizationResourceServerId"`
	ClientId                      pulumi.StringOutput                                `pulumi:"clientId"`
	ConfigureScope                ClientPermissionsConfigureScopePtrOutput           `pulumi:"configureScope"`
	Enabled                       pulumi.BoolOutput                                  `pulumi:"enabled"`
	ManageScope                   ClientPermissionsManageScopePtrOutput              `pulumi:"manageScope"`
	MapRolesClientScopeScope      ClientPermissionsMapRolesClientScopeScopePtrOutput `pulumi:"mapRolesClientScopeScope"`
	MapRolesCompositeScope        ClientPermissionsMapRolesCompositeScopePtrOutput   `pulumi:"mapRolesCompositeScope"`
	MapRolesScope                 ClientPermissionsMapRolesScopePtrOutput            `pulumi:"mapRolesScope"`
	RealmId                       pulumi.StringOutput                                `pulumi:"realmId"`
	TokenExchangeScope            ClientPermissionsTokenExchangeScopePtrOutput       `pulumi:"tokenExchangeScope"`
	ViewScope                     ClientPermissionsViewScopePtrOutput                `pulumi:"viewScope"`
}

// NewClientPermissions registers a new resource with the given unique name, arguments, and options.
func NewClientPermissions(ctx *pulumi.Context,
	name string, args *ClientPermissionsArgs, opts ...pulumi.ResourceOption) (*ClientPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClientPermissions
	err := ctx.RegisterResource("keycloak:openid/clientPermissions:ClientPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientPermissions gets an existing ClientPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientPermissionsState, opts ...pulumi.ResourceOption) (*ClientPermissions, error) {
	var resource ClientPermissions
	err := ctx.ReadResource("keycloak:openid/clientPermissions:ClientPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientPermissions resources.
type clientPermissionsState struct {
	// Resource server id representing the realm management client on which this permission is managed
	AuthorizationResourceServerId *string                                    `pulumi:"authorizationResourceServerId"`
	ClientId                      *string                                    `pulumi:"clientId"`
	ConfigureScope                *ClientPermissionsConfigureScope           `pulumi:"configureScope"`
	Enabled                       *bool                                      `pulumi:"enabled"`
	ManageScope                   *ClientPermissionsManageScope              `pulumi:"manageScope"`
	MapRolesClientScopeScope      *ClientPermissionsMapRolesClientScopeScope `pulumi:"mapRolesClientScopeScope"`
	MapRolesCompositeScope        *ClientPermissionsMapRolesCompositeScope   `pulumi:"mapRolesCompositeScope"`
	MapRolesScope                 *ClientPermissionsMapRolesScope            `pulumi:"mapRolesScope"`
	RealmId                       *string                                    `pulumi:"realmId"`
	TokenExchangeScope            *ClientPermissionsTokenExchangeScope       `pulumi:"tokenExchangeScope"`
	ViewScope                     *ClientPermissionsViewScope                `pulumi:"viewScope"`
}

type ClientPermissionsState struct {
	// Resource server id representing the realm management client on which this permission is managed
	AuthorizationResourceServerId pulumi.StringPtrInput
	ClientId                      pulumi.StringPtrInput
	ConfigureScope                ClientPermissionsConfigureScopePtrInput
	Enabled                       pulumi.BoolPtrInput
	ManageScope                   ClientPermissionsManageScopePtrInput
	MapRolesClientScopeScope      ClientPermissionsMapRolesClientScopeScopePtrInput
	MapRolesCompositeScope        ClientPermissionsMapRolesCompositeScopePtrInput
	MapRolesScope                 ClientPermissionsMapRolesScopePtrInput
	RealmId                       pulumi.StringPtrInput
	TokenExchangeScope            ClientPermissionsTokenExchangeScopePtrInput
	ViewScope                     ClientPermissionsViewScopePtrInput
}

func (ClientPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientPermissionsState)(nil)).Elem()
}

type clientPermissionsArgs struct {
	ClientId                 string                                     `pulumi:"clientId"`
	ConfigureScope           *ClientPermissionsConfigureScope           `pulumi:"configureScope"`
	ManageScope              *ClientPermissionsManageScope              `pulumi:"manageScope"`
	MapRolesClientScopeScope *ClientPermissionsMapRolesClientScopeScope `pulumi:"mapRolesClientScopeScope"`
	MapRolesCompositeScope   *ClientPermissionsMapRolesCompositeScope   `pulumi:"mapRolesCompositeScope"`
	MapRolesScope            *ClientPermissionsMapRolesScope            `pulumi:"mapRolesScope"`
	RealmId                  string                                     `pulumi:"realmId"`
	TokenExchangeScope       *ClientPermissionsTokenExchangeScope       `pulumi:"tokenExchangeScope"`
	ViewScope                *ClientPermissionsViewScope                `pulumi:"viewScope"`
}

// The set of arguments for constructing a ClientPermissions resource.
type ClientPermissionsArgs struct {
	ClientId                 pulumi.StringInput
	ConfigureScope           ClientPermissionsConfigureScopePtrInput
	ManageScope              ClientPermissionsManageScopePtrInput
	MapRolesClientScopeScope ClientPermissionsMapRolesClientScopeScopePtrInput
	MapRolesCompositeScope   ClientPermissionsMapRolesCompositeScopePtrInput
	MapRolesScope            ClientPermissionsMapRolesScopePtrInput
	RealmId                  pulumi.StringInput
	TokenExchangeScope       ClientPermissionsTokenExchangeScopePtrInput
	ViewScope                ClientPermissionsViewScopePtrInput
}

func (ClientPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientPermissionsArgs)(nil)).Elem()
}

type ClientPermissionsInput interface {
	pulumi.Input

	ToClientPermissionsOutput() ClientPermissionsOutput
	ToClientPermissionsOutputWithContext(ctx context.Context) ClientPermissionsOutput
}

func (*ClientPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientPermissions)(nil)).Elem()
}

func (i *ClientPermissions) ToClientPermissionsOutput() ClientPermissionsOutput {
	return i.ToClientPermissionsOutputWithContext(context.Background())
}

func (i *ClientPermissions) ToClientPermissionsOutputWithContext(ctx context.Context) ClientPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientPermissionsOutput)
}

func (i *ClientPermissions) ToOutput(ctx context.Context) pulumix.Output[*ClientPermissions] {
	return pulumix.Output[*ClientPermissions]{
		OutputState: i.ToClientPermissionsOutputWithContext(ctx).OutputState,
	}
}

// ClientPermissionsArrayInput is an input type that accepts ClientPermissionsArray and ClientPermissionsArrayOutput values.
// You can construct a concrete instance of `ClientPermissionsArrayInput` via:
//
//	ClientPermissionsArray{ ClientPermissionsArgs{...} }
type ClientPermissionsArrayInput interface {
	pulumi.Input

	ToClientPermissionsArrayOutput() ClientPermissionsArrayOutput
	ToClientPermissionsArrayOutputWithContext(context.Context) ClientPermissionsArrayOutput
}

type ClientPermissionsArray []ClientPermissionsInput

func (ClientPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientPermissions)(nil)).Elem()
}

func (i ClientPermissionsArray) ToClientPermissionsArrayOutput() ClientPermissionsArrayOutput {
	return i.ToClientPermissionsArrayOutputWithContext(context.Background())
}

func (i ClientPermissionsArray) ToClientPermissionsArrayOutputWithContext(ctx context.Context) ClientPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientPermissionsArrayOutput)
}

func (i ClientPermissionsArray) ToOutput(ctx context.Context) pulumix.Output[[]*ClientPermissions] {
	return pulumix.Output[[]*ClientPermissions]{
		OutputState: i.ToClientPermissionsArrayOutputWithContext(ctx).OutputState,
	}
}

// ClientPermissionsMapInput is an input type that accepts ClientPermissionsMap and ClientPermissionsMapOutput values.
// You can construct a concrete instance of `ClientPermissionsMapInput` via:
//
//	ClientPermissionsMap{ "key": ClientPermissionsArgs{...} }
type ClientPermissionsMapInput interface {
	pulumi.Input

	ToClientPermissionsMapOutput() ClientPermissionsMapOutput
	ToClientPermissionsMapOutputWithContext(context.Context) ClientPermissionsMapOutput
}

type ClientPermissionsMap map[string]ClientPermissionsInput

func (ClientPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientPermissions)(nil)).Elem()
}

func (i ClientPermissionsMap) ToClientPermissionsMapOutput() ClientPermissionsMapOutput {
	return i.ToClientPermissionsMapOutputWithContext(context.Background())
}

func (i ClientPermissionsMap) ToClientPermissionsMapOutputWithContext(ctx context.Context) ClientPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientPermissionsMapOutput)
}

func (i ClientPermissionsMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClientPermissions] {
	return pulumix.Output[map[string]*ClientPermissions]{
		OutputState: i.ToClientPermissionsMapOutputWithContext(ctx).OutputState,
	}
}

type ClientPermissionsOutput struct{ *pulumi.OutputState }

func (ClientPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientPermissions)(nil)).Elem()
}

func (o ClientPermissionsOutput) ToClientPermissionsOutput() ClientPermissionsOutput {
	return o
}

func (o ClientPermissionsOutput) ToClientPermissionsOutputWithContext(ctx context.Context) ClientPermissionsOutput {
	return o
}

func (o ClientPermissionsOutput) ToOutput(ctx context.Context) pulumix.Output[*ClientPermissions] {
	return pulumix.Output[*ClientPermissions]{
		OutputState: o.OutputState,
	}
}

// Resource server id representing the realm management client on which this permission is managed
func (o ClientPermissionsOutput) AuthorizationResourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientPermissions) pulumi.StringOutput { return v.AuthorizationResourceServerId }).(pulumi.StringOutput)
}

func (o ClientPermissionsOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientPermissions) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

func (o ClientPermissionsOutput) ConfigureScope() ClientPermissionsConfigureScopePtrOutput {
	return o.ApplyT(func(v *ClientPermissions) ClientPermissionsConfigureScopePtrOutput { return v.ConfigureScope }).(ClientPermissionsConfigureScopePtrOutput)
}

func (o ClientPermissionsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClientPermissions) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ClientPermissionsOutput) ManageScope() ClientPermissionsManageScopePtrOutput {
	return o.ApplyT(func(v *ClientPermissions) ClientPermissionsManageScopePtrOutput { return v.ManageScope }).(ClientPermissionsManageScopePtrOutput)
}

func (o ClientPermissionsOutput) MapRolesClientScopeScope() ClientPermissionsMapRolesClientScopeScopePtrOutput {
	return o.ApplyT(func(v *ClientPermissions) ClientPermissionsMapRolesClientScopeScopePtrOutput {
		return v.MapRolesClientScopeScope
	}).(ClientPermissionsMapRolesClientScopeScopePtrOutput)
}

func (o ClientPermissionsOutput) MapRolesCompositeScope() ClientPermissionsMapRolesCompositeScopePtrOutput {
	return o.ApplyT(func(v *ClientPermissions) ClientPermissionsMapRolesCompositeScopePtrOutput {
		return v.MapRolesCompositeScope
	}).(ClientPermissionsMapRolesCompositeScopePtrOutput)
}

func (o ClientPermissionsOutput) MapRolesScope() ClientPermissionsMapRolesScopePtrOutput {
	return o.ApplyT(func(v *ClientPermissions) ClientPermissionsMapRolesScopePtrOutput { return v.MapRolesScope }).(ClientPermissionsMapRolesScopePtrOutput)
}

func (o ClientPermissionsOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientPermissions) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

func (o ClientPermissionsOutput) TokenExchangeScope() ClientPermissionsTokenExchangeScopePtrOutput {
	return o.ApplyT(func(v *ClientPermissions) ClientPermissionsTokenExchangeScopePtrOutput { return v.TokenExchangeScope }).(ClientPermissionsTokenExchangeScopePtrOutput)
}

func (o ClientPermissionsOutput) ViewScope() ClientPermissionsViewScopePtrOutput {
	return o.ApplyT(func(v *ClientPermissions) ClientPermissionsViewScopePtrOutput { return v.ViewScope }).(ClientPermissionsViewScopePtrOutput)
}

type ClientPermissionsArrayOutput struct{ *pulumi.OutputState }

func (ClientPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientPermissions)(nil)).Elem()
}

func (o ClientPermissionsArrayOutput) ToClientPermissionsArrayOutput() ClientPermissionsArrayOutput {
	return o
}

func (o ClientPermissionsArrayOutput) ToClientPermissionsArrayOutputWithContext(ctx context.Context) ClientPermissionsArrayOutput {
	return o
}

func (o ClientPermissionsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ClientPermissions] {
	return pulumix.Output[[]*ClientPermissions]{
		OutputState: o.OutputState,
	}
}

func (o ClientPermissionsArrayOutput) Index(i pulumi.IntInput) ClientPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientPermissions {
		return vs[0].([]*ClientPermissions)[vs[1].(int)]
	}).(ClientPermissionsOutput)
}

type ClientPermissionsMapOutput struct{ *pulumi.OutputState }

func (ClientPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientPermissions)(nil)).Elem()
}

func (o ClientPermissionsMapOutput) ToClientPermissionsMapOutput() ClientPermissionsMapOutput {
	return o
}

func (o ClientPermissionsMapOutput) ToClientPermissionsMapOutputWithContext(ctx context.Context) ClientPermissionsMapOutput {
	return o
}

func (o ClientPermissionsMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClientPermissions] {
	return pulumix.Output[map[string]*ClientPermissions]{
		OutputState: o.OutputState,
	}
}

func (o ClientPermissionsMapOutput) MapIndex(k pulumi.StringInput) ClientPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientPermissions {
		return vs[0].(map[string]*ClientPermissions)[vs[1].(string)]
	}).(ClientPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientPermissionsInput)(nil)).Elem(), &ClientPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientPermissionsArrayInput)(nil)).Elem(), ClientPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientPermissionsMapInput)(nil)).Elem(), ClientPermissionsMap{})
	pulumi.RegisterOutputType(ClientPermissionsOutput{})
	pulumi.RegisterOutputType(ClientPermissionsArrayOutput{})
	pulumi.RegisterOutputType(ClientPermissionsMapOutput{})
}
