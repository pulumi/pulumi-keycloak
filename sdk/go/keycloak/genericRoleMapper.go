// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allow for creating and managing a client's or client scope's role mappings within Keycloak.
//
// By default, all the user role mappings of the user are added as claims within the token (OIDC) or assertion (SAML). When
// `fullScopeAllowed` is set to `false` for a client, role scope mapping allows you to limit the roles that get declared
// inside an access token for a client.
//
// ## Example Usage
// ### Realm Role To Client)
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
//			_, err = keycloak.NewGenericRoleMapper(ctx, "clientRoleMapper", &keycloak.GenericRoleMapperArgs{
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
// ### Client Role To Client)
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
//			_, err = keycloak.NewGenericRoleMapper(ctx, "clientBRoleMapper", &keycloak.GenericRoleMapperArgs{
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
// ### Realm Role To Client Scope)
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
//			_, err = keycloak.NewGenericRoleMapper(ctx, "clientRoleMapper", &keycloak.GenericRoleMapperArgs{
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
// ### Client Role To Client Scope)
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
//			_, err = keycloak.NewGenericRoleMapper(ctx, "clientBRoleMapper", &keycloak.GenericRoleMapperArgs{
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
//
// ## Import
//
// Generic client role mappers can be imported using one of the following two formats- When mapping a role to a client, use the format `{{realmId}}/client/{{clientId}}/scope-mappings/{{roleClientId}}/{{roleId}}` - When mapping a role to a client scope, use the format `{{realmId}}/client-scope/{{clientScopeId}}/scope-mappings/{{roleClientId}}/{{roleId}}` Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:index/genericRoleMapper:GenericRoleMapper client_role_mapper my-realm/client/23888550-5dcd-41f6-85ba-554233021e9c/scope-mappings/ce51f004-bdfb-4dd5-a963-c4487d2dec5b/ff3aa49f-bc07-4030-8783-41918c3614a3
//
// ```
type GenericRoleMapper struct {
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

// NewGenericRoleMapper registers a new resource with the given unique name, arguments, and options.
func NewGenericRoleMapper(ctx *pulumi.Context,
	name string, args *GenericRoleMapperArgs, opts ...pulumi.ResourceOption) (*GenericRoleMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	var resource GenericRoleMapper
	err := ctx.RegisterResource("keycloak:index/genericRoleMapper:GenericRoleMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGenericRoleMapper gets an existing GenericRoleMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGenericRoleMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GenericRoleMapperState, opts ...pulumi.ResourceOption) (*GenericRoleMapper, error) {
	var resource GenericRoleMapper
	err := ctx.ReadResource("keycloak:index/genericRoleMapper:GenericRoleMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GenericRoleMapper resources.
type genericRoleMapperState struct {
	// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId *string `pulumi:"clientId"`
	// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The realm this role mapper exists within.
	RealmId *string `pulumi:"realmId"`
	// The ID of the role to be added to this role mapper.
	RoleId *string `pulumi:"roleId"`
}

type GenericRoleMapperState struct {
	// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId pulumi.StringPtrInput
	// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId pulumi.StringPtrInput
	// The realm this role mapper exists within.
	RealmId pulumi.StringPtrInput
	// The ID of the role to be added to this role mapper.
	RoleId pulumi.StringPtrInput
}

func (GenericRoleMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*genericRoleMapperState)(nil)).Elem()
}

type genericRoleMapperArgs struct {
	// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId *string `pulumi:"clientId"`
	// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The realm this role mapper exists within.
	RealmId string `pulumi:"realmId"`
	// The ID of the role to be added to this role mapper.
	RoleId string `pulumi:"roleId"`
}

// The set of arguments for constructing a GenericRoleMapper resource.
type GenericRoleMapperArgs struct {
	// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId pulumi.StringPtrInput
	// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId pulumi.StringPtrInput
	// The realm this role mapper exists within.
	RealmId pulumi.StringInput
	// The ID of the role to be added to this role mapper.
	RoleId pulumi.StringInput
}

func (GenericRoleMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*genericRoleMapperArgs)(nil)).Elem()
}

type GenericRoleMapperInput interface {
	pulumi.Input

	ToGenericRoleMapperOutput() GenericRoleMapperOutput
	ToGenericRoleMapperOutputWithContext(ctx context.Context) GenericRoleMapperOutput
}

func (*GenericRoleMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericRoleMapper)(nil)).Elem()
}

func (i *GenericRoleMapper) ToGenericRoleMapperOutput() GenericRoleMapperOutput {
	return i.ToGenericRoleMapperOutputWithContext(context.Background())
}

func (i *GenericRoleMapper) ToGenericRoleMapperOutputWithContext(ctx context.Context) GenericRoleMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericRoleMapperOutput)
}

// GenericRoleMapperArrayInput is an input type that accepts GenericRoleMapperArray and GenericRoleMapperArrayOutput values.
// You can construct a concrete instance of `GenericRoleMapperArrayInput` via:
//
//	GenericRoleMapperArray{ GenericRoleMapperArgs{...} }
type GenericRoleMapperArrayInput interface {
	pulumi.Input

	ToGenericRoleMapperArrayOutput() GenericRoleMapperArrayOutput
	ToGenericRoleMapperArrayOutputWithContext(context.Context) GenericRoleMapperArrayOutput
}

type GenericRoleMapperArray []GenericRoleMapperInput

func (GenericRoleMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GenericRoleMapper)(nil)).Elem()
}

func (i GenericRoleMapperArray) ToGenericRoleMapperArrayOutput() GenericRoleMapperArrayOutput {
	return i.ToGenericRoleMapperArrayOutputWithContext(context.Background())
}

func (i GenericRoleMapperArray) ToGenericRoleMapperArrayOutputWithContext(ctx context.Context) GenericRoleMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericRoleMapperArrayOutput)
}

// GenericRoleMapperMapInput is an input type that accepts GenericRoleMapperMap and GenericRoleMapperMapOutput values.
// You can construct a concrete instance of `GenericRoleMapperMapInput` via:
//
//	GenericRoleMapperMap{ "key": GenericRoleMapperArgs{...} }
type GenericRoleMapperMapInput interface {
	pulumi.Input

	ToGenericRoleMapperMapOutput() GenericRoleMapperMapOutput
	ToGenericRoleMapperMapOutputWithContext(context.Context) GenericRoleMapperMapOutput
}

type GenericRoleMapperMap map[string]GenericRoleMapperInput

func (GenericRoleMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GenericRoleMapper)(nil)).Elem()
}

func (i GenericRoleMapperMap) ToGenericRoleMapperMapOutput() GenericRoleMapperMapOutput {
	return i.ToGenericRoleMapperMapOutputWithContext(context.Background())
}

func (i GenericRoleMapperMap) ToGenericRoleMapperMapOutputWithContext(ctx context.Context) GenericRoleMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericRoleMapperMapOutput)
}

type GenericRoleMapperOutput struct{ *pulumi.OutputState }

func (GenericRoleMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericRoleMapper)(nil)).Elem()
}

func (o GenericRoleMapperOutput) ToGenericRoleMapperOutput() GenericRoleMapperOutput {
	return o
}

func (o GenericRoleMapperOutput) ToGenericRoleMapperOutputWithContext(ctx context.Context) GenericRoleMapperOutput {
	return o
}

// The ID of the client this role mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
func (o GenericRoleMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GenericRoleMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The ID of the client scope this role mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
func (o GenericRoleMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GenericRoleMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// The realm this role mapper exists within.
func (o GenericRoleMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *GenericRoleMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The ID of the role to be added to this role mapper.
func (o GenericRoleMapperOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *GenericRoleMapper) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

type GenericRoleMapperArrayOutput struct{ *pulumi.OutputState }

func (GenericRoleMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GenericRoleMapper)(nil)).Elem()
}

func (o GenericRoleMapperArrayOutput) ToGenericRoleMapperArrayOutput() GenericRoleMapperArrayOutput {
	return o
}

func (o GenericRoleMapperArrayOutput) ToGenericRoleMapperArrayOutputWithContext(ctx context.Context) GenericRoleMapperArrayOutput {
	return o
}

func (o GenericRoleMapperArrayOutput) Index(i pulumi.IntInput) GenericRoleMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GenericRoleMapper {
		return vs[0].([]*GenericRoleMapper)[vs[1].(int)]
	}).(GenericRoleMapperOutput)
}

type GenericRoleMapperMapOutput struct{ *pulumi.OutputState }

func (GenericRoleMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GenericRoleMapper)(nil)).Elem()
}

func (o GenericRoleMapperMapOutput) ToGenericRoleMapperMapOutput() GenericRoleMapperMapOutput {
	return o
}

func (o GenericRoleMapperMapOutput) ToGenericRoleMapperMapOutputWithContext(ctx context.Context) GenericRoleMapperMapOutput {
	return o
}

func (o GenericRoleMapperMapOutput) MapIndex(k pulumi.StringInput) GenericRoleMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GenericRoleMapper {
		return vs[0].(map[string]*GenericRoleMapper)[vs[1].(string)]
	}).(GenericRoleMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GenericRoleMapperInput)(nil)).Elem(), &GenericRoleMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenericRoleMapperArrayInput)(nil)).Elem(), GenericRoleMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenericRoleMapperMapInput)(nil)).Elem(), GenericRoleMapperMap{})
	pulumi.RegisterOutputType(GenericRoleMapperOutput{})
	pulumi.RegisterOutputType(GenericRoleMapperArrayOutput{})
	pulumi.RegisterOutputType(GenericRoleMapperMapOutput{})
}
