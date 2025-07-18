// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing user realm role protocol mappers within Keycloak.
//
// User realm role protocol mappers allow you to define a claim containing the list of the realm roles.
//
// Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
// multiple different clients.
//
// ## Example Usage
//
// ### Client)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/openid"
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
//				ClientId:   pulumi.String("client"),
//				Name:       pulumi.String("client"),
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
// ### Client Scope)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/openid"
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
// ## Import
//
// Protocol mappers can be imported using one of the following formats:
//
// - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
//
// - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:openid/userRealmRoleProtocolMapper:UserRealmRoleProtocolMapper user_realm_role_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
//
// ```sh
// $ pulumi import keycloak:openid/userRealmRoleProtocolMapper:UserRealmRoleProtocolMapper user_realm_role_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
type UserRealmRoleProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the Token Introspection response body. Defaults to `true`.
	AddToTokenIntrospection pulumi.BoolPtrOutput `pulumi:"addToTokenIntrospection"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringOutput `pulumi:"claimName"`
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType pulumi.StringPtrOutput `pulumi:"claimValueType"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued pulumi.BoolPtrOutput `pulumi:"multivalued"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// A prefix for each Realm Role.
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
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the Token Introspection response body. Defaults to `true`.
	AddToTokenIntrospection *bool `pulumi:"addToTokenIntrospection"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName *string `pulumi:"claimName"`
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued *bool `pulumi:"multivalued"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
	// A prefix for each Realm Role.
	RealmRolePrefix *string `pulumi:"realmRolePrefix"`
}

type UserRealmRoleProtocolMapperState struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the Token Introspection response body. Defaults to `true`.
	AddToTokenIntrospection pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringPtrInput
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType pulumi.StringPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued pulumi.BoolPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
	// A prefix for each Realm Role.
	RealmRolePrefix pulumi.StringPtrInput
}

func (UserRealmRoleProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRealmRoleProtocolMapperState)(nil)).Elem()
}

type userRealmRoleProtocolMapperArgs struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the Token Introspection response body. Defaults to `true`.
	AddToTokenIntrospection *bool `pulumi:"addToTokenIntrospection"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName string `pulumi:"claimName"`
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued *bool `pulumi:"multivalued"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
	// A prefix for each Realm Role.
	RealmRolePrefix *string `pulumi:"realmRolePrefix"`
}

// The set of arguments for constructing a UserRealmRoleProtocolMapper resource.
type UserRealmRoleProtocolMapperArgs struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the Token Introspection response body. Defaults to `true`.
	AddToTokenIntrospection pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringInput
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType pulumi.StringPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued pulumi.BoolPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
	// A prefix for each Realm Role.
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

// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
func (o UserRealmRoleProtocolMapperOutput) AddToAccessToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToAccessToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
func (o UserRealmRoleProtocolMapperOutput) AddToIdToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToIdToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the property should be added as a claim to the Token Introspection response body. Defaults to `true`.
func (o UserRealmRoleProtocolMapperOutput) AddToTokenIntrospection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToTokenIntrospection }).(pulumi.BoolPtrOutput)
}

// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
func (o UserRealmRoleProtocolMapperOutput) AddToUserinfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToUserinfo }).(pulumi.BoolPtrOutput)
}

// The name of the claim to insert into a token.
func (o UserRealmRoleProtocolMapperOutput) ClaimName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringOutput { return v.ClaimName }).(pulumi.StringOutput)
}

// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
func (o UserRealmRoleProtocolMapperOutput) ClaimValueType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClaimValueType }).(pulumi.StringPtrOutput)
}

// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
func (o UserRealmRoleProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
func (o UserRealmRoleProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
func (o UserRealmRoleProtocolMapperOutput) Multivalued() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.BoolPtrOutput { return v.Multivalued }).(pulumi.BoolPtrOutput)
}

// The display name of this protocol mapper in the GUI.
func (o UserRealmRoleProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm this protocol mapper exists within.
func (o UserRealmRoleProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRealmRoleProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// A prefix for each Realm Role.
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
