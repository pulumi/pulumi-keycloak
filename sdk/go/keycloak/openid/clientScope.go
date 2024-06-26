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

// ## # openid.ClientScope
//
// Allows for creating and managing Keycloak client scopes that can be attached to
// clients that use the OpenID Connect protocol.
//
// Client Scopes can be used to share common protocol and role mappings between multiple
// clients within a realm. They can also be used by clients to conditionally request
// claims or roles for a user based on the OAuth 2.0 `scope` parameter.
//
// ### Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/openid"
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
//			_, err = openid.NewClientScope(ctx, "openid_client_scope", &openid.ClientScopeArgs{
//				RealmId:     realm.ID(),
//				Name:        pulumi.String("groups"),
//				Description: pulumi.String("When requested, this scope will map a user's group memberships to a claim"),
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
//   - `realmId` - (Required) The realm this client scope belongs to.
//   - `name` - (Required) The display name of this client scope in the GUI.
//   - `description` - (Optional) The description of this client scope in the GUI.
//   - `consentScreenText` - (Optional) When set, a consent screen will be displayed to users
//     authenticating to clients with this scope attached. The consent screen will display the string
//     value of this attribute.
//
// ### Import
//
// Client scopes can be imported using the format `{{realm_id}}/{{client_scope_id}}`, where `clientScopeId` is the unique ID that Keycloak
// assigns to the client scope upon creation. This value can be found in the URI when editing this client scope in the GUI, and is typically a GUID.
//
// Example:
type ClientScope struct {
	pulumi.CustomResourceState

	ConsentScreenText   pulumi.StringPtrOutput `pulumi:"consentScreenText"`
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	GuiOrder            pulumi.IntPtrOutput    `pulumi:"guiOrder"`
	IncludeInTokenScope pulumi.BoolPtrOutput   `pulumi:"includeInTokenScope"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	RealmId             pulumi.StringOutput    `pulumi:"realmId"`
}

// NewClientScope registers a new resource with the given unique name, arguments, and options.
func NewClientScope(ctx *pulumi.Context,
	name string, args *ClientScopeArgs, opts ...pulumi.ResourceOption) (*ClientScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClientScope
	err := ctx.RegisterResource("keycloak:openid/clientScope:ClientScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientScope gets an existing ClientScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientScopeState, opts ...pulumi.ResourceOption) (*ClientScope, error) {
	var resource ClientScope
	err := ctx.ReadResource("keycloak:openid/clientScope:ClientScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientScope resources.
type clientScopeState struct {
	ConsentScreenText   *string `pulumi:"consentScreenText"`
	Description         *string `pulumi:"description"`
	GuiOrder            *int    `pulumi:"guiOrder"`
	IncludeInTokenScope *bool   `pulumi:"includeInTokenScope"`
	Name                *string `pulumi:"name"`
	RealmId             *string `pulumi:"realmId"`
}

type ClientScopeState struct {
	ConsentScreenText   pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	GuiOrder            pulumi.IntPtrInput
	IncludeInTokenScope pulumi.BoolPtrInput
	Name                pulumi.StringPtrInput
	RealmId             pulumi.StringPtrInput
}

func (ClientScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientScopeState)(nil)).Elem()
}

type clientScopeArgs struct {
	ConsentScreenText   *string `pulumi:"consentScreenText"`
	Description         *string `pulumi:"description"`
	GuiOrder            *int    `pulumi:"guiOrder"`
	IncludeInTokenScope *bool   `pulumi:"includeInTokenScope"`
	Name                *string `pulumi:"name"`
	RealmId             string  `pulumi:"realmId"`
}

// The set of arguments for constructing a ClientScope resource.
type ClientScopeArgs struct {
	ConsentScreenText   pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	GuiOrder            pulumi.IntPtrInput
	IncludeInTokenScope pulumi.BoolPtrInput
	Name                pulumi.StringPtrInput
	RealmId             pulumi.StringInput
}

func (ClientScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientScopeArgs)(nil)).Elem()
}

type ClientScopeInput interface {
	pulumi.Input

	ToClientScopeOutput() ClientScopeOutput
	ToClientScopeOutputWithContext(ctx context.Context) ClientScopeOutput
}

func (*ClientScope) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientScope)(nil)).Elem()
}

func (i *ClientScope) ToClientScopeOutput() ClientScopeOutput {
	return i.ToClientScopeOutputWithContext(context.Background())
}

func (i *ClientScope) ToClientScopeOutputWithContext(ctx context.Context) ClientScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientScopeOutput)
}

// ClientScopeArrayInput is an input type that accepts ClientScopeArray and ClientScopeArrayOutput values.
// You can construct a concrete instance of `ClientScopeArrayInput` via:
//
//	ClientScopeArray{ ClientScopeArgs{...} }
type ClientScopeArrayInput interface {
	pulumi.Input

	ToClientScopeArrayOutput() ClientScopeArrayOutput
	ToClientScopeArrayOutputWithContext(context.Context) ClientScopeArrayOutput
}

type ClientScopeArray []ClientScopeInput

func (ClientScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientScope)(nil)).Elem()
}

func (i ClientScopeArray) ToClientScopeArrayOutput() ClientScopeArrayOutput {
	return i.ToClientScopeArrayOutputWithContext(context.Background())
}

func (i ClientScopeArray) ToClientScopeArrayOutputWithContext(ctx context.Context) ClientScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientScopeArrayOutput)
}

// ClientScopeMapInput is an input type that accepts ClientScopeMap and ClientScopeMapOutput values.
// You can construct a concrete instance of `ClientScopeMapInput` via:
//
//	ClientScopeMap{ "key": ClientScopeArgs{...} }
type ClientScopeMapInput interface {
	pulumi.Input

	ToClientScopeMapOutput() ClientScopeMapOutput
	ToClientScopeMapOutputWithContext(context.Context) ClientScopeMapOutput
}

type ClientScopeMap map[string]ClientScopeInput

func (ClientScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientScope)(nil)).Elem()
}

func (i ClientScopeMap) ToClientScopeMapOutput() ClientScopeMapOutput {
	return i.ToClientScopeMapOutputWithContext(context.Background())
}

func (i ClientScopeMap) ToClientScopeMapOutputWithContext(ctx context.Context) ClientScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientScopeMapOutput)
}

type ClientScopeOutput struct{ *pulumi.OutputState }

func (ClientScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientScope)(nil)).Elem()
}

func (o ClientScopeOutput) ToClientScopeOutput() ClientScopeOutput {
	return o
}

func (o ClientScopeOutput) ToClientScopeOutputWithContext(ctx context.Context) ClientScopeOutput {
	return o
}

func (o ClientScopeOutput) ConsentScreenText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientScope) pulumi.StringPtrOutput { return v.ConsentScreenText }).(pulumi.StringPtrOutput)
}

func (o ClientScopeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientScope) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ClientScopeOutput) GuiOrder() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClientScope) pulumi.IntPtrOutput { return v.GuiOrder }).(pulumi.IntPtrOutput)
}

func (o ClientScopeOutput) IncludeInTokenScope() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClientScope) pulumi.BoolPtrOutput { return v.IncludeInTokenScope }).(pulumi.BoolPtrOutput)
}

func (o ClientScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClientScopeOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientScope) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type ClientScopeArrayOutput struct{ *pulumi.OutputState }

func (ClientScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientScope)(nil)).Elem()
}

func (o ClientScopeArrayOutput) ToClientScopeArrayOutput() ClientScopeArrayOutput {
	return o
}

func (o ClientScopeArrayOutput) ToClientScopeArrayOutputWithContext(ctx context.Context) ClientScopeArrayOutput {
	return o
}

func (o ClientScopeArrayOutput) Index(i pulumi.IntInput) ClientScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientScope {
		return vs[0].([]*ClientScope)[vs[1].(int)]
	}).(ClientScopeOutput)
}

type ClientScopeMapOutput struct{ *pulumi.OutputState }

func (ClientScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientScope)(nil)).Elem()
}

func (o ClientScopeMapOutput) ToClientScopeMapOutput() ClientScopeMapOutput {
	return o
}

func (o ClientScopeMapOutput) ToClientScopeMapOutputWithContext(ctx context.Context) ClientScopeMapOutput {
	return o
}

func (o ClientScopeMapOutput) MapIndex(k pulumi.StringInput) ClientScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientScope {
		return vs[0].(map[string]*ClientScope)[vs[1].(string)]
	}).(ClientScopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientScopeInput)(nil)).Elem(), &ClientScope{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientScopeArrayInput)(nil)).Elem(), ClientScopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientScopeMapInput)(nil)).Elem(), ClientScopeMap{})
	pulumi.RegisterOutputType(ClientScopeOutput{})
	pulumi.RegisterOutputType(ClientScopeArrayOutput{})
	pulumi.RegisterOutputType(ClientScopeMapOutput{})
}
