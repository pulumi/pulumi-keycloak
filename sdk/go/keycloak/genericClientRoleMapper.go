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

// !> **WARNING:** This resource is deprecated and will be removed in the next major version. Please use `GenericRoleMapper` instead.
//
// Allow for creating and managing a client's scope mappings within Keycloak.
//
// By default, all the user role mappings of the user are added as claims within the token (OIDC) or assertion (SAML). When
// `fullScopeAllowed` is set to `false` for a client, role scope mapping allows you to limit the roles that get declared
// inside an access token for a client.
//
// ## Example Usage
//
// ### Realm Role To Client)
//
// <!--Start PulumiCodeChooser -->
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
//			client, err := openid.NewClient(ctx, "client", &openid.ClientArgs{
//				RealmId:    realm.ID(),
//				ClientId:   pulumi.String("client"),
//				Enabled:    pulumi.Bool(true),
//				AccessType: pulumi.String("BEARER-ONLY"),
//			})
//			if err != nil {
//				return err
//			}
//			realmRole, err := keycloak.NewRole(ctx, "realmRole", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				Description: pulumi.String("My Realm Role"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewGenericClientRoleMapper(ctx, "clientRoleMapper", &keycloak.GenericClientRoleMapperArgs{
//				RealmId:  realm.ID(),
//				ClientId: client.ID(),
//				RoleId:   realmRole.ID(),
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
// ### Client Role To Client)
//
// <!--Start PulumiCodeChooser -->
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
//			clientA, err := openid.NewClient(ctx, "clientA", &openid.ClientArgs{
//				RealmId:          realm.ID(),
//				ClientId:         pulumi.String("client-a"),
//				Enabled:          pulumi.Bool(true),
//				AccessType:       pulumi.String("BEARER-ONLY"),
//				FullScopeAllowed: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			clientRoleA, err := keycloak.NewRole(ctx, "clientRoleA", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				ClientId:    clientA.ID(),
//				Description: pulumi.String("My Client Role"),
//			})
//			if err != nil {
//				return err
//			}
//			clientB, err := openid.NewClient(ctx, "clientB", &openid.ClientArgs{
//				RealmId:    realm.ID(),
//				ClientId:   pulumi.String("client-b"),
//				Enabled:    pulumi.Bool(true),
//				AccessType: pulumi.String("BEARER-ONLY"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "clientRoleB", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				ClientId:    clientB.ID(),
//				Description: pulumi.String("My Client Role"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewGenericClientRoleMapper(ctx, "clientBRoleMapper", &keycloak.GenericClientRoleMapperArgs{
//				RealmId:  realm.ID(),
//				ClientId: clientB.ID(),
//				RoleId:   clientRoleA.ID(),
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
// ### Realm Role To Client Scope)
//
// <!--Start PulumiCodeChooser -->
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
//			clientScope, err := openid.NewClientScope(ctx, "clientScope", &openid.ClientScopeArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			realmRole, err := keycloak.NewRole(ctx, "realmRole", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				Description: pulumi.String("My Realm Role"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewGenericClientRoleMapper(ctx, "clientRoleMapper", &keycloak.GenericClientRoleMapperArgs{
//				RealmId:       realm.ID(),
//				ClientScopeId: clientScope.ID(),
//				RoleId:        realmRole.ID(),
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
// ### Client Role To Client Scope)
//
// <!--Start PulumiCodeChooser -->
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
//			client, err := openid.NewClient(ctx, "client", &openid.ClientArgs{
//				RealmId:    realm.ID(),
//				ClientId:   pulumi.String("client"),
//				Enabled:    pulumi.Bool(true),
//				AccessType: pulumi.String("BEARER-ONLY"),
//			})
//			if err != nil {
//				return err
//			}
//			clientRole, err := keycloak.NewRole(ctx, "clientRole", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				ClientId:    client.ID(),
//				Description: pulumi.String("My Client Role"),
//			})
//			if err != nil {
//				return err
//			}
//			clientScope, err := openid.NewClientScope(ctx, "clientScope", &openid.ClientScopeArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewGenericClientRoleMapper(ctx, "clientBRoleMapper", &keycloak.GenericClientRoleMapperArgs{
//				RealmId:       realm.ID(),
//				ClientScopeId: clientScope.ID(),
//				RoleId:        clientRole.ID(),
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
// Generic client role mappers can be imported using one of the following two formats:
//
// - When mapping a role to a client, use the format `{{realmId}}/client/{{clientId}}/scope-mappings/{{roleClientId}}/{{roleId}}`
//
// - When mapping a role to a client scope, use the format `{{realmId}}/client-scope/{{clientScopeId}}/scope-mappings/{{roleClientId}}/{{roleId}}`
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:index/genericClientRoleMapper:GenericClientRoleMapper client_role_mapper my-realm/client/23888550-5dcd-41f6-85ba-554233021e9c/scope-mappings/ce51f004-bdfb-4dd5-a963-c4487d2dec5b/ff3aa49f-bc07-4030-8783-41918c3614a3
// ```
type GenericClientRoleMapper struct {
	pulumi.CustomResourceState

	// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// The realm this role mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The ID of the role to be added to this role mapper.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
}

// NewGenericClientRoleMapper registers a new resource with the given unique name, arguments, and options.
func NewGenericClientRoleMapper(ctx *pulumi.Context,
	name string, args *GenericClientRoleMapperArgs, opts ...pulumi.ResourceOption) (*GenericClientRoleMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GenericClientRoleMapper
	err := ctx.RegisterResource("keycloak:index/genericClientRoleMapper:GenericClientRoleMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGenericClientRoleMapper gets an existing GenericClientRoleMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGenericClientRoleMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GenericClientRoleMapperState, opts ...pulumi.ResourceOption) (*GenericClientRoleMapper, error) {
	var resource GenericClientRoleMapper
	err := ctx.ReadResource("keycloak:index/genericClientRoleMapper:GenericClientRoleMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GenericClientRoleMapper resources.
type genericClientRoleMapperState struct {
	// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId *string `pulumi:"clientId"`
	// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The realm this role mapper exists within.
	RealmId *string `pulumi:"realmId"`
	// The ID of the role to be added to this role mapper.
	RoleId *string `pulumi:"roleId"`
}

type GenericClientRoleMapperState struct {
	// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId pulumi.StringPtrInput
	// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId pulumi.StringPtrInput
	// The realm this role mapper exists within.
	RealmId pulumi.StringPtrInput
	// The ID of the role to be added to this role mapper.
	RoleId pulumi.StringPtrInput
}

func (GenericClientRoleMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*genericClientRoleMapperState)(nil)).Elem()
}

type genericClientRoleMapperArgs struct {
	// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId *string `pulumi:"clientId"`
	// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The realm this role mapper exists within.
	RealmId string `pulumi:"realmId"`
	// The ID of the role to be added to this role mapper.
	RoleId string `pulumi:"roleId"`
}

// The set of arguments for constructing a GenericClientRoleMapper resource.
type GenericClientRoleMapperArgs struct {
	// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId pulumi.StringPtrInput
	// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId pulumi.StringPtrInput
	// The realm this role mapper exists within.
	RealmId pulumi.StringInput
	// The ID of the role to be added to this role mapper.
	RoleId pulumi.StringInput
}

func (GenericClientRoleMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*genericClientRoleMapperArgs)(nil)).Elem()
}

type GenericClientRoleMapperInput interface {
	pulumi.Input

	ToGenericClientRoleMapperOutput() GenericClientRoleMapperOutput
	ToGenericClientRoleMapperOutputWithContext(ctx context.Context) GenericClientRoleMapperOutput
}

func (*GenericClientRoleMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericClientRoleMapper)(nil)).Elem()
}

func (i *GenericClientRoleMapper) ToGenericClientRoleMapperOutput() GenericClientRoleMapperOutput {
	return i.ToGenericClientRoleMapperOutputWithContext(context.Background())
}

func (i *GenericClientRoleMapper) ToGenericClientRoleMapperOutputWithContext(ctx context.Context) GenericClientRoleMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericClientRoleMapperOutput)
}

// GenericClientRoleMapperArrayInput is an input type that accepts GenericClientRoleMapperArray and GenericClientRoleMapperArrayOutput values.
// You can construct a concrete instance of `GenericClientRoleMapperArrayInput` via:
//
//	GenericClientRoleMapperArray{ GenericClientRoleMapperArgs{...} }
type GenericClientRoleMapperArrayInput interface {
	pulumi.Input

	ToGenericClientRoleMapperArrayOutput() GenericClientRoleMapperArrayOutput
	ToGenericClientRoleMapperArrayOutputWithContext(context.Context) GenericClientRoleMapperArrayOutput
}

type GenericClientRoleMapperArray []GenericClientRoleMapperInput

func (GenericClientRoleMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GenericClientRoleMapper)(nil)).Elem()
}

func (i GenericClientRoleMapperArray) ToGenericClientRoleMapperArrayOutput() GenericClientRoleMapperArrayOutput {
	return i.ToGenericClientRoleMapperArrayOutputWithContext(context.Background())
}

func (i GenericClientRoleMapperArray) ToGenericClientRoleMapperArrayOutputWithContext(ctx context.Context) GenericClientRoleMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericClientRoleMapperArrayOutput)
}

// GenericClientRoleMapperMapInput is an input type that accepts GenericClientRoleMapperMap and GenericClientRoleMapperMapOutput values.
// You can construct a concrete instance of `GenericClientRoleMapperMapInput` via:
//
//	GenericClientRoleMapperMap{ "key": GenericClientRoleMapperArgs{...} }
type GenericClientRoleMapperMapInput interface {
	pulumi.Input

	ToGenericClientRoleMapperMapOutput() GenericClientRoleMapperMapOutput
	ToGenericClientRoleMapperMapOutputWithContext(context.Context) GenericClientRoleMapperMapOutput
}

type GenericClientRoleMapperMap map[string]GenericClientRoleMapperInput

func (GenericClientRoleMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GenericClientRoleMapper)(nil)).Elem()
}

func (i GenericClientRoleMapperMap) ToGenericClientRoleMapperMapOutput() GenericClientRoleMapperMapOutput {
	return i.ToGenericClientRoleMapperMapOutputWithContext(context.Background())
}

func (i GenericClientRoleMapperMap) ToGenericClientRoleMapperMapOutputWithContext(ctx context.Context) GenericClientRoleMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericClientRoleMapperMapOutput)
}

type GenericClientRoleMapperOutput struct{ *pulumi.OutputState }

func (GenericClientRoleMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericClientRoleMapper)(nil)).Elem()
}

func (o GenericClientRoleMapperOutput) ToGenericClientRoleMapperOutput() GenericClientRoleMapperOutput {
	return o
}

func (o GenericClientRoleMapperOutput) ToGenericClientRoleMapperOutputWithContext(ctx context.Context) GenericClientRoleMapperOutput {
	return o
}

// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
func (o GenericClientRoleMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GenericClientRoleMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
func (o GenericClientRoleMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GenericClientRoleMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// The realm this role mapper exists within.
func (o GenericClientRoleMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *GenericClientRoleMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The ID of the role to be added to this role mapper.
func (o GenericClientRoleMapperOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *GenericClientRoleMapper) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

type GenericClientRoleMapperArrayOutput struct{ *pulumi.OutputState }

func (GenericClientRoleMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GenericClientRoleMapper)(nil)).Elem()
}

func (o GenericClientRoleMapperArrayOutput) ToGenericClientRoleMapperArrayOutput() GenericClientRoleMapperArrayOutput {
	return o
}

func (o GenericClientRoleMapperArrayOutput) ToGenericClientRoleMapperArrayOutputWithContext(ctx context.Context) GenericClientRoleMapperArrayOutput {
	return o
}

func (o GenericClientRoleMapperArrayOutput) Index(i pulumi.IntInput) GenericClientRoleMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GenericClientRoleMapper {
		return vs[0].([]*GenericClientRoleMapper)[vs[1].(int)]
	}).(GenericClientRoleMapperOutput)
}

type GenericClientRoleMapperMapOutput struct{ *pulumi.OutputState }

func (GenericClientRoleMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GenericClientRoleMapper)(nil)).Elem()
}

func (o GenericClientRoleMapperMapOutput) ToGenericClientRoleMapperMapOutput() GenericClientRoleMapperMapOutput {
	return o
}

func (o GenericClientRoleMapperMapOutput) ToGenericClientRoleMapperMapOutputWithContext(ctx context.Context) GenericClientRoleMapperMapOutput {
	return o
}

func (o GenericClientRoleMapperMapOutput) MapIndex(k pulumi.StringInput) GenericClientRoleMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GenericClientRoleMapper {
		return vs[0].(map[string]*GenericClientRoleMapper)[vs[1].(string)]
	}).(GenericClientRoleMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GenericClientRoleMapperInput)(nil)).Elem(), &GenericClientRoleMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenericClientRoleMapperArrayInput)(nil)).Elem(), GenericClientRoleMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenericClientRoleMapperMapInput)(nil)).Elem(), GenericClientRoleMapperMap{})
	pulumi.RegisterOutputType(GenericClientRoleMapperOutput{})
	pulumi.RegisterOutputType(GenericClientRoleMapperArrayOutput{})
	pulumi.RegisterOutputType(GenericClientRoleMapperMapOutput{})
}
