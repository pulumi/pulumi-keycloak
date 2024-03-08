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

// Allows for creating and managing user client role protocol mappers within Keycloak.
//
// User client role protocol mappers allow you to define a claim containing the list of a client roles.
//
// Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
// multiple different clients.
//
// ## Example Usage
//
// ### Client)
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
//			openidClient, err := openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
//				RealmId:    realm.ID(),
//				ClientId:   pulumi.String("client"),
//				Enabled:    pulumi.Bool(true),
//				AccessType: pulumi.String("CONFIDENTIAL"),
//				ValidRedirectUris: pulumi.StringArray{
//					pulumi.String("http://localhost:8080/openid-callback"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewUserClientRoleProtocolMapper(ctx, "userClientRoleMapper", &openid.UserClientRoleProtocolMapperArgs{
//				RealmId:   realm.ID(),
//				ClientId:  openidClient.ID(),
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
// <!--End PulumiCodeChooser -->
//
// ### Client Scope)
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
//			_, err = openid.NewUserClientRoleProtocolMapper(ctx, "userClientRoleMapper", &openid.UserClientRoleProtocolMapperArgs{
//				RealmId:       realm.ID(),
//				ClientScopeId: clientScope.ID(),
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
// <!--End PulumiCodeChooser -->
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
// $ pulumi import keycloak:openid/userClientRoleProtocolMapper:UserClientRoleProtocolMapper user_client_role_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
//
// ```sh
// $ pulumi import keycloak:openid/userClientRoleProtocolMapper:UserClientRoleProtocolMapper user_client_role_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
type UserClientRoleProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringOutput `pulumi:"claimName"`
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType pulumi.StringPtrOutput `pulumi:"claimValueType"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The Client ID for role mappings. Just client roles of this client will be added to the token. If this is unset, client roles of all clients will be added to the token.
	ClientIdForRoleMappings pulumi.StringPtrOutput `pulumi:"clientIdForRoleMappings"`
	// A prefix for each Client Role.
	ClientRolePrefix pulumi.StringPtrOutput `pulumi:"clientRolePrefix"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued pulumi.BoolPtrOutput `pulumi:"multivalued"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewUserClientRoleProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserClientRoleProtocolMapper(ctx *pulumi.Context,
	name string, args *UserClientRoleProtocolMapperArgs, opts ...pulumi.ResourceOption) (*UserClientRoleProtocolMapper, error) {
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
	var resource UserClientRoleProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/userClientRoleProtocolMapper:UserClientRoleProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserClientRoleProtocolMapper gets an existing UserClientRoleProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserClientRoleProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserClientRoleProtocolMapperState, opts ...pulumi.ResourceOption) (*UserClientRoleProtocolMapper, error) {
	var resource UserClientRoleProtocolMapper
	err := ctx.ReadResource("keycloak:openid/userClientRoleProtocolMapper:UserClientRoleProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserClientRoleProtocolMapper resources.
type userClientRoleProtocolMapperState struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName *string `pulumi:"claimName"`
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The Client ID for role mappings. Just client roles of this client will be added to the token. If this is unset, client roles of all clients will be added to the token.
	ClientIdForRoleMappings *string `pulumi:"clientIdForRoleMappings"`
	// A prefix for each Client Role.
	ClientRolePrefix *string `pulumi:"clientRolePrefix"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued *bool `pulumi:"multivalued"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
}

type UserClientRoleProtocolMapperState struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringPtrInput
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType pulumi.StringPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The Client ID for role mappings. Just client roles of this client will be added to the token. If this is unset, client roles of all clients will be added to the token.
	ClientIdForRoleMappings pulumi.StringPtrInput
	// A prefix for each Client Role.
	ClientRolePrefix pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued pulumi.BoolPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
}

func (UserClientRoleProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userClientRoleProtocolMapperState)(nil)).Elem()
}

type userClientRoleProtocolMapperArgs struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName string `pulumi:"claimName"`
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The Client ID for role mappings. Just client roles of this client will be added to the token. If this is unset, client roles of all clients will be added to the token.
	ClientIdForRoleMappings *string `pulumi:"clientIdForRoleMappings"`
	// A prefix for each Client Role.
	ClientRolePrefix *string `pulumi:"clientRolePrefix"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued *bool `pulumi:"multivalued"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a UserClientRoleProtocolMapper resource.
type UserClientRoleProtocolMapperArgs struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringInput
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType pulumi.StringPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The Client ID for role mappings. Just client roles of this client will be added to the token. If this is unset, client roles of all clients will be added to the token.
	ClientIdForRoleMappings pulumi.StringPtrInput
	// A prefix for each Client Role.
	ClientRolePrefix pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued pulumi.BoolPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
}

func (UserClientRoleProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userClientRoleProtocolMapperArgs)(nil)).Elem()
}

type UserClientRoleProtocolMapperInput interface {
	pulumi.Input

	ToUserClientRoleProtocolMapperOutput() UserClientRoleProtocolMapperOutput
	ToUserClientRoleProtocolMapperOutputWithContext(ctx context.Context) UserClientRoleProtocolMapperOutput
}

func (*UserClientRoleProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**UserClientRoleProtocolMapper)(nil)).Elem()
}

func (i *UserClientRoleProtocolMapper) ToUserClientRoleProtocolMapperOutput() UserClientRoleProtocolMapperOutput {
	return i.ToUserClientRoleProtocolMapperOutputWithContext(context.Background())
}

func (i *UserClientRoleProtocolMapper) ToUserClientRoleProtocolMapperOutputWithContext(ctx context.Context) UserClientRoleProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserClientRoleProtocolMapperOutput)
}

// UserClientRoleProtocolMapperArrayInput is an input type that accepts UserClientRoleProtocolMapperArray and UserClientRoleProtocolMapperArrayOutput values.
// You can construct a concrete instance of `UserClientRoleProtocolMapperArrayInput` via:
//
//	UserClientRoleProtocolMapperArray{ UserClientRoleProtocolMapperArgs{...} }
type UserClientRoleProtocolMapperArrayInput interface {
	pulumi.Input

	ToUserClientRoleProtocolMapperArrayOutput() UserClientRoleProtocolMapperArrayOutput
	ToUserClientRoleProtocolMapperArrayOutputWithContext(context.Context) UserClientRoleProtocolMapperArrayOutput
}

type UserClientRoleProtocolMapperArray []UserClientRoleProtocolMapperInput

func (UserClientRoleProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserClientRoleProtocolMapper)(nil)).Elem()
}

func (i UserClientRoleProtocolMapperArray) ToUserClientRoleProtocolMapperArrayOutput() UserClientRoleProtocolMapperArrayOutput {
	return i.ToUserClientRoleProtocolMapperArrayOutputWithContext(context.Background())
}

func (i UserClientRoleProtocolMapperArray) ToUserClientRoleProtocolMapperArrayOutputWithContext(ctx context.Context) UserClientRoleProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserClientRoleProtocolMapperArrayOutput)
}

// UserClientRoleProtocolMapperMapInput is an input type that accepts UserClientRoleProtocolMapperMap and UserClientRoleProtocolMapperMapOutput values.
// You can construct a concrete instance of `UserClientRoleProtocolMapperMapInput` via:
//
//	UserClientRoleProtocolMapperMap{ "key": UserClientRoleProtocolMapperArgs{...} }
type UserClientRoleProtocolMapperMapInput interface {
	pulumi.Input

	ToUserClientRoleProtocolMapperMapOutput() UserClientRoleProtocolMapperMapOutput
	ToUserClientRoleProtocolMapperMapOutputWithContext(context.Context) UserClientRoleProtocolMapperMapOutput
}

type UserClientRoleProtocolMapperMap map[string]UserClientRoleProtocolMapperInput

func (UserClientRoleProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserClientRoleProtocolMapper)(nil)).Elem()
}

func (i UserClientRoleProtocolMapperMap) ToUserClientRoleProtocolMapperMapOutput() UserClientRoleProtocolMapperMapOutput {
	return i.ToUserClientRoleProtocolMapperMapOutputWithContext(context.Background())
}

func (i UserClientRoleProtocolMapperMap) ToUserClientRoleProtocolMapperMapOutputWithContext(ctx context.Context) UserClientRoleProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserClientRoleProtocolMapperMapOutput)
}

type UserClientRoleProtocolMapperOutput struct{ *pulumi.OutputState }

func (UserClientRoleProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserClientRoleProtocolMapper)(nil)).Elem()
}

func (o UserClientRoleProtocolMapperOutput) ToUserClientRoleProtocolMapperOutput() UserClientRoleProtocolMapperOutput {
	return o
}

func (o UserClientRoleProtocolMapperOutput) ToUserClientRoleProtocolMapperOutputWithContext(ctx context.Context) UserClientRoleProtocolMapperOutput {
	return o
}

// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
func (o UserClientRoleProtocolMapperOutput) AddToAccessToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToAccessToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
func (o UserClientRoleProtocolMapperOutput) AddToIdToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToIdToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
func (o UserClientRoleProtocolMapperOutput) AddToUserinfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.BoolPtrOutput { return v.AddToUserinfo }).(pulumi.BoolPtrOutput)
}

// The name of the claim to insert into a token.
func (o UserClientRoleProtocolMapperOutput) ClaimName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.StringOutput { return v.ClaimName }).(pulumi.StringOutput)
}

// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
func (o UserClientRoleProtocolMapperOutput) ClaimValueType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClaimValueType }).(pulumi.StringPtrOutput)
}

// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
func (o UserClientRoleProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The Client ID for role mappings. Just client roles of this client will be added to the token. If this is unset, client roles of all clients will be added to the token.
func (o UserClientRoleProtocolMapperOutput) ClientIdForRoleMappings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClientIdForRoleMappings }).(pulumi.StringPtrOutput)
}

// A prefix for each Client Role.
func (o UserClientRoleProtocolMapperOutput) ClientRolePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClientRolePrefix }).(pulumi.StringPtrOutput)
}

// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
func (o UserClientRoleProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
func (o UserClientRoleProtocolMapperOutput) Multivalued() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.BoolPtrOutput { return v.Multivalued }).(pulumi.BoolPtrOutput)
}

// The display name of this protocol mapper in the GUI.
func (o UserClientRoleProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm this protocol mapper exists within.
func (o UserClientRoleProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserClientRoleProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type UserClientRoleProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (UserClientRoleProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserClientRoleProtocolMapper)(nil)).Elem()
}

func (o UserClientRoleProtocolMapperArrayOutput) ToUserClientRoleProtocolMapperArrayOutput() UserClientRoleProtocolMapperArrayOutput {
	return o
}

func (o UserClientRoleProtocolMapperArrayOutput) ToUserClientRoleProtocolMapperArrayOutputWithContext(ctx context.Context) UserClientRoleProtocolMapperArrayOutput {
	return o
}

func (o UserClientRoleProtocolMapperArrayOutput) Index(i pulumi.IntInput) UserClientRoleProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserClientRoleProtocolMapper {
		return vs[0].([]*UserClientRoleProtocolMapper)[vs[1].(int)]
	}).(UserClientRoleProtocolMapperOutput)
}

type UserClientRoleProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (UserClientRoleProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserClientRoleProtocolMapper)(nil)).Elem()
}

func (o UserClientRoleProtocolMapperMapOutput) ToUserClientRoleProtocolMapperMapOutput() UserClientRoleProtocolMapperMapOutput {
	return o
}

func (o UserClientRoleProtocolMapperMapOutput) ToUserClientRoleProtocolMapperMapOutputWithContext(ctx context.Context) UserClientRoleProtocolMapperMapOutput {
	return o
}

func (o UserClientRoleProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) UserClientRoleProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserClientRoleProtocolMapper {
		return vs[0].(map[string]*UserClientRoleProtocolMapper)[vs[1].(string)]
	}).(UserClientRoleProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserClientRoleProtocolMapperInput)(nil)).Elem(), &UserClientRoleProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserClientRoleProtocolMapperArrayInput)(nil)).Elem(), UserClientRoleProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserClientRoleProtocolMapperMapInput)(nil)).Elem(), UserClientRoleProtocolMapperMap{})
	pulumi.RegisterOutputType(UserClientRoleProtocolMapperOutput{})
	pulumi.RegisterOutputType(UserClientRoleProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(UserClientRoleProtocolMapperMapOutput{})
}
