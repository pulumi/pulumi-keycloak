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

// ## # openid.UserRealmRoleProtocolMapper
//
// Allows for creating and managing user realm role protocol mappers within
// Keycloak.
//
// User realm role protocol mappers allow you to define a claim containing the list of the realm roles.
// Protocol mappers can be defined for a single client, or they can
// be defined for a client scope which can be shared between multiple different
// clients.
//
// ### Example Usage (Client)
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
//			openidClient, err := openid.NewClient(ctx, "openid_client", &openid.ClientArgs{
//				RealmId:    realm.ID(),
//				ClientId:   pulumi.String("test-client"),
//				Name:       pulumi.String("test client"),
//				Enabled:    pulumi.Bool(true),
//				AccessType: pulumi.String("CONFIDENTIAL"),
//				ValidRedirectUris: pulumi.StringArray{
//					pulumi.String("http://localhost:8080/openid-callback"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewUserRealmRoleProtocolMapper(ctx, "user_realm_role_mapper", &openid.UserRealmRoleProtocolMapperArgs{
//				RealmId:   realm.ID(),
//				ClientId:  openidClient.ID(),
//				Name:      pulumi.String("user-realm-role-mapper"),
//				ClaimName: pulumi.String("foo"),
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
// ### Example Usage (Client Scope)
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
//			clientScope, err := openid.NewClientScope(ctx, "client_scope", &openid.ClientScopeArgs{
//				RealmId: realm.ID(),
//				Name:    pulumi.String("test-client-scope"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewUserRealmRoleProtocolMapper(ctx, "user_realm_role_mapper", &openid.UserRealmRoleProtocolMapperArgs{
//				RealmId:       realm.ID(),
//				ClientScopeId: clientScope.ID(),
//				Name:          pulumi.String("user-realm-role-mapper"),
//				ClaimName:     pulumi.String("foo"),
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
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `claimName` - (Required) The name of the claim to insert into a token.
// - `claimValueType` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
// - `multivalued` - (Optional) Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `true`.
// - `realmRolePrefix` - (Optional) A prefix for each Realm Role.
// - `addToIdToken` - (Optional) Indicates if the property should be added as a claim to the id token. Defaults to `true`.
// - `addToAccessToken` - (Optional) Indicates if the property should be added as a claim to the access token. Defaults to `true`.
// - `addToUserinfo` - (Optional) Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
//
// ### Import
//
// Protocol mappers can be imported using one of the following formats:
// - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
// - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`
//
// Example:
type UserRealmRoleProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	ClaimName     pulumi.StringOutput  `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrOutput `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued pulumi.BoolPtrOutput `pulumi:"multivalued"`
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// Prefix that will be added to each realm role.
	RealmRolePrefix pulumi.StringPtrOutput `pulumi:"realmRolePrefix"`
}

// NewUserRealmRoleProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserRealmRoleProtocolMapper(ctx *pulumi.Context,
	name string, args *UserRealmRoleProtocolMapperArgs, opts ...pulumi.ResourceOption) (*UserRealmRoleProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClaimName == nil {
		return nil, errors.New("invalid value for required argument 'ClaimName'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserRealmRoleProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/userRealmRoleProtocolMapper:UserRealmRoleProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserRealmRoleProtocolMapper gets an existing UserRealmRoleProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRealmRoleProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRealmRoleProtocolMapperState, opts ...pulumi.ResourceOption) (*UserRealmRoleProtocolMapper, error) {
	var resource UserRealmRoleProtocolMapper
	err := ctx.ReadResource("keycloak:openid/userRealmRoleProtocolMapper:UserRealmRoleProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserRealmRoleProtocolMapper resources.
type userRealmRoleProtocolMapperState struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo *bool   `pulumi:"addToUserinfo"`
	ClaimName     *string `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued *bool `pulumi:"multivalued"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId *string `pulumi:"realmId"`
	// Prefix that will be added to each realm role.
	RealmRolePrefix *string `pulumi:"realmRolePrefix"`
}

type UserRealmRoleProtocolMapperState struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrInput
	ClaimName     pulumi.StringPtrInput
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringPtrInput
	// Prefix that will be added to each realm role.
	RealmRolePrefix pulumi.StringPtrInput
}

func (UserRealmRoleProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRealmRoleProtocolMapperState)(nil)).Elem()
}

type userRealmRoleProtocolMapperArgs struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo *bool  `pulumi:"addToUserinfo"`
	ClaimName     string `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued *bool `pulumi:"multivalued"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId string `pulumi:"realmId"`
	// Prefix that will be added to each realm role.
	RealmRolePrefix *string `pulumi:"realmRolePrefix"`
}

// The set of arguments for constructing a UserRealmRoleProtocolMapper resource.
type UserRealmRoleProtocolMapperArgs struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrInput
	ClaimName     pulumi.StringInput
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringInput
	// Prefix that will be added to each realm role.
	RealmRolePrefix pulumi.StringPtrInput
}

func (UserRealmRoleProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRealmRoleProtocolMapperArgs)(nil)).Elem()
}

type UserRealmRoleProtocolMapperInput interface {
	pulumi.Input

	ToUserRealmRoleProtocolMapperOutput() UserRealmRoleProtocolMapperOutput
	ToUserRealmRoleProtocolMapperOutputWithContext(ctx context.Context) UserRealmRoleProtocolMapperOutput
}

func (*UserRealmRoleProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRealmRoleProtocolMapper)(nil)).Elem()
}

func (i *UserRealmRoleProtocolMapper) ToUserRealmRoleProtocolMapperOutput() UserRealmRoleProtocolMapperOutput {
	return i.ToUserRealmRoleProtocolMapperOutputWithContext(context.Background())
}

func (i *UserRealmRoleProtocolMapper) ToUserRealmRoleProtocolMapperOutputWithContext(ctx context.Context) UserRealmRoleProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRealmRoleProtocolMapperOutput)
}

// UserRealmRoleProtocolMapperArrayInput is an input type that accepts UserRealmRoleProtocolMapperArray and UserRealmRoleProtocolMapperArrayOutput values.
// You can construct a concrete instance of `UserRealmRoleProtocolMapperArrayInput` via:
//
//	UserRealmRoleProtocolMapperArray{ UserRealmRoleProtocolMapperArgs{...} }
type UserRealmRoleProtocolMapperArrayInput interface {
	pulumi.Input

	ToUserRealmRoleProtocolMapperArrayOutput() UserRealmRoleProtocolMapperArrayOutput
	ToUserRealmRoleProtocolMapperArrayOutputWithContext(context.Context) UserRealmRoleProtocolMapperArrayOutput
}

type UserRealmRoleProtocolMapperArray []UserRealmRoleProtocolMapperInput

func (UserRealmRoleProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserRealmRoleProtocolMapper)(nil)).Elem()
}

func (i UserRealmRoleProtocolMapperArray) ToUserRealmRoleProtocolMapperArrayOutput() UserRealmRoleProtocolMapperArrayOutput {
	return i.ToUserRealmRoleProtocolMapperArrayOutputWithContext(context.Background())
}

func (i UserRealmRoleProtocolMapperArray) ToUserRealmRoleProtocolMapperArrayOutputWithContext(ctx context.Context) UserRealmRoleProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRealmRoleProtocolMapperArrayOutput)
}

// UserRealmRoleProtocolMapperMapInput is an input type that accepts UserRealmRoleProtocolMapperMap and UserRealmRoleProtocolMapperMapOutput values.
// You can construct a concrete instance of `UserRealmRoleProtocolMapperMapInput` via:
//
//	UserRealmRoleProtocolMapperMap{ "key": UserRealmRoleProtocolMapperArgs{...} }
type UserRealmRoleProtocolMapperMapInput interface {
	pulumi.Input

	ToUserRealmRoleProtocolMapperMapOutput() UserRealmRoleProtocolMapperMapOutput
	ToUserRealmRoleProtocolMapperMapOutputWithContext(context.Context) UserRealmRoleProtocolMapperMapOutput
}

type UserRealmRoleProtocolMapperMap map[string]UserRealmRoleProtocolMapperInput

func (UserRealmRoleProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserRealmRoleProtocolMapper)(nil)).Elem()
}

func (i UserRealmRoleProtocolMapperMap) ToUserRealmRoleProtocolMapperMapOutput() UserRealmRoleProtocolMapperMapOutput {
	return i.ToUserRealmRoleProtocolMapperMapOutputWithContext(context.Background())
}

func (i UserRealmRoleProtocolMapperMap) ToUserRealmRoleProtocolMapperMapOutputWithContext(ctx context.Context) UserRealmRoleProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRealmRoleProtocolMapperMapOutput)
}

type UserRealmRoleProtocolMapperOutput struct{ *pulumi.OutputState }

func (UserRealmRoleProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRealmRoleProtocolMapper)(nil)).Elem()
}

func (o UserRealmRoleProtocolMapperOutput) ToUserRealmRoleProtocolMapperOutput() UserRealmRoleProtocolMapperOutput {
	return o
}

func (o UserRealmRoleProtocolMapperOutput) ToUserRealmRoleProtocolMapperOutputWithContext(ctx context.Context) UserRealmRoleProtocolMapperOutput {
	return o
}

// Indicates if the attribute should be a claim in the access token.
func (o UserRealmRoleProtocolMapperOutput) AddToAccessToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToAccessToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the attribute should be a claim in the id token.
func (o UserRealmRoleProtocolMapperOutput) AddToIdToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToIdToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the attribute should appear in the userinfo response body.
func (o UserRealmRoleProtocolMapperOutput) AddToUserinfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToUserinfo }).(pulumi.BoolPtrOutput)
}

func (o UserRealmRoleProtocolMapperOutput) ClaimName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringOutput { return v.ClaimName }).(pulumi.StringOutput)
}

// Claim type used when serializing tokens.
func (o UserRealmRoleProtocolMapperOutput) ClaimValueType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClaimValueType }).(pulumi.StringPtrOutput)
}

// The mapper's associated client. Cannot be used at the same time as client_scope_id.
func (o UserRealmRoleProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The mapper's associated client scope. Cannot be used at the same time as client_id.
func (o UserRealmRoleProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// Indicates whether this attribute is a single value or an array of values.
func (o UserRealmRoleProtocolMapperOutput) Multivalued() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.BoolPtrOutput { return v.Multivalued }).(pulumi.BoolPtrOutput)
}

// A human-friendly name that will appear in the Keycloak console.
func (o UserRealmRoleProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm id where the associated client or client scope exists.
func (o UserRealmRoleProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// Prefix that will be added to each realm role.
func (o UserRealmRoleProtocolMapperOutput) RealmRolePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringPtrOutput { return v.RealmRolePrefix }).(pulumi.StringPtrOutput)
}

type UserRealmRoleProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (UserRealmRoleProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserRealmRoleProtocolMapper)(nil)).Elem()
}

func (o UserRealmRoleProtocolMapperArrayOutput) ToUserRealmRoleProtocolMapperArrayOutput() UserRealmRoleProtocolMapperArrayOutput {
	return o
}

func (o UserRealmRoleProtocolMapperArrayOutput) ToUserRealmRoleProtocolMapperArrayOutputWithContext(ctx context.Context) UserRealmRoleProtocolMapperArrayOutput {
	return o
}

func (o UserRealmRoleProtocolMapperArrayOutput) Index(i pulumi.IntInput) UserRealmRoleProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserRealmRoleProtocolMapper {
		return vs[0].([]*UserRealmRoleProtocolMapper)[vs[1].(int)]
	}).(UserRealmRoleProtocolMapperOutput)
}

type UserRealmRoleProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (UserRealmRoleProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserRealmRoleProtocolMapper)(nil)).Elem()
}

func (o UserRealmRoleProtocolMapperMapOutput) ToUserRealmRoleProtocolMapperMapOutput() UserRealmRoleProtocolMapperMapOutput {
	return o
}

func (o UserRealmRoleProtocolMapperMapOutput) ToUserRealmRoleProtocolMapperMapOutputWithContext(ctx context.Context) UserRealmRoleProtocolMapperMapOutput {
	return o
}

func (o UserRealmRoleProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) UserRealmRoleProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserRealmRoleProtocolMapper {
		return vs[0].(map[string]*UserRealmRoleProtocolMapper)[vs[1].(string)]
	}).(UserRealmRoleProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserRealmRoleProtocolMapperInput)(nil)).Elem(), &UserRealmRoleProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRealmRoleProtocolMapperArrayInput)(nil)).Elem(), UserRealmRoleProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserRealmRoleProtocolMapperMapInput)(nil)).Elem(), UserRealmRoleProtocolMapperMap{})
	pulumi.RegisterOutputType(UserRealmRoleProtocolMapperOutput{})
	pulumi.RegisterOutputType(UserRealmRoleProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(UserRealmRoleProtocolMapperMapOutput{})
}
