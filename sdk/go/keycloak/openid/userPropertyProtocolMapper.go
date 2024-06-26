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

// ## # openid.UserPropertyProtocolMapper
//
// Allows for creating and managing user property protocol mappers within
// Keycloak.
//
// User property protocol mappers allow you to map built in properties defined
// on the Keycloak user interface to a claim in a token. Protocol mappers can be
// defined for a single client, or they can be defined for a client scope which
// can be shared between multiple different clients.
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
//			_, err = openid.NewUserPropertyProtocolMapper(ctx, "user_property_mapper", &openid.UserPropertyProtocolMapperArgs{
//				RealmId:      realm.ID(),
//				ClientId:     openidClient.ID(),
//				Name:         pulumi.String("test-mapper"),
//				UserProperty: pulumi.String("email"),
//				ClaimName:    pulumi.String("email"),
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
//			_, err = openid.NewUserPropertyProtocolMapper(ctx, "user_property_mapper", &openid.UserPropertyProtocolMapperArgs{
//				RealmId:       realm.ID(),
//				ClientScopeId: clientScope.ID(),
//				Name:          pulumi.String("test-mapper"),
//				UserProperty:  pulumi.String("email"),
//				ClaimName:     pulumi.String("email"),
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
// - `userProperty` - (Required) The built in user property (such as email) to map a claim for.
// - `claimName` - (Required) The name of the claim to insert into a token.
// - `claimValueType` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
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
type UserPropertyProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the property should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the property should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// Indicates if the property should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	ClaimName     pulumi.StringOutput  `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrOutput `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId      pulumi.StringOutput `pulumi:"realmId"`
	UserProperty pulumi.StringOutput `pulumi:"userProperty"`
}

// NewUserPropertyProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserPropertyProtocolMapper(ctx *pulumi.Context,
	name string, args *UserPropertyProtocolMapperArgs, opts ...pulumi.ResourceOption) (*UserPropertyProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClaimName == nil {
		return nil, errors.New("invalid value for required argument 'ClaimName'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.UserProperty == nil {
		return nil, errors.New("invalid value for required argument 'UserProperty'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserPropertyProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPropertyProtocolMapper gets an existing UserPropertyProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPropertyProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPropertyProtocolMapperState, opts ...pulumi.ResourceOption) (*UserPropertyProtocolMapper, error) {
	var resource UserPropertyProtocolMapper
	err := ctx.ReadResource("keycloak:openid/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPropertyProtocolMapper resources.
type userPropertyProtocolMapperState struct {
	// Indicates if the property should be a claim in the access token.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be a claim in the id token.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should appear in the userinfo response body.
	AddToUserinfo *bool   `pulumi:"addToUserinfo"`
	ClaimName     *string `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId      *string `pulumi:"realmId"`
	UserProperty *string `pulumi:"userProperty"`
}

type UserPropertyProtocolMapperState struct {
	// Indicates if the property should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrInput
	ClaimName     pulumi.StringPtrInput
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId      pulumi.StringPtrInput
	UserProperty pulumi.StringPtrInput
}

func (UserPropertyProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPropertyProtocolMapperState)(nil)).Elem()
}

type userPropertyProtocolMapperArgs struct {
	// Indicates if the property should be a claim in the access token.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be a claim in the id token.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should appear in the userinfo response body.
	AddToUserinfo *bool  `pulumi:"addToUserinfo"`
	ClaimName     string `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId      string `pulumi:"realmId"`
	UserProperty string `pulumi:"userProperty"`
}

// The set of arguments for constructing a UserPropertyProtocolMapper resource.
type UserPropertyProtocolMapperArgs struct {
	// Indicates if the property should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrInput
	ClaimName     pulumi.StringInput
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId      pulumi.StringInput
	UserProperty pulumi.StringInput
}

func (UserPropertyProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPropertyProtocolMapperArgs)(nil)).Elem()
}

type UserPropertyProtocolMapperInput interface {
	pulumi.Input

	ToUserPropertyProtocolMapperOutput() UserPropertyProtocolMapperOutput
	ToUserPropertyProtocolMapperOutputWithContext(ctx context.Context) UserPropertyProtocolMapperOutput
}

func (*UserPropertyProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPropertyProtocolMapper)(nil)).Elem()
}

func (i *UserPropertyProtocolMapper) ToUserPropertyProtocolMapperOutput() UserPropertyProtocolMapperOutput {
	return i.ToUserPropertyProtocolMapperOutputWithContext(context.Background())
}

func (i *UserPropertyProtocolMapper) ToUserPropertyProtocolMapperOutputWithContext(ctx context.Context) UserPropertyProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertyProtocolMapperOutput)
}

// UserPropertyProtocolMapperArrayInput is an input type that accepts UserPropertyProtocolMapperArray and UserPropertyProtocolMapperArrayOutput values.
// You can construct a concrete instance of `UserPropertyProtocolMapperArrayInput` via:
//
//	UserPropertyProtocolMapperArray{ UserPropertyProtocolMapperArgs{...} }
type UserPropertyProtocolMapperArrayInput interface {
	pulumi.Input

	ToUserPropertyProtocolMapperArrayOutput() UserPropertyProtocolMapperArrayOutput
	ToUserPropertyProtocolMapperArrayOutputWithContext(context.Context) UserPropertyProtocolMapperArrayOutput
}

type UserPropertyProtocolMapperArray []UserPropertyProtocolMapperInput

func (UserPropertyProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPropertyProtocolMapper)(nil)).Elem()
}

func (i UserPropertyProtocolMapperArray) ToUserPropertyProtocolMapperArrayOutput() UserPropertyProtocolMapperArrayOutput {
	return i.ToUserPropertyProtocolMapperArrayOutputWithContext(context.Background())
}

func (i UserPropertyProtocolMapperArray) ToUserPropertyProtocolMapperArrayOutputWithContext(ctx context.Context) UserPropertyProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertyProtocolMapperArrayOutput)
}

// UserPropertyProtocolMapperMapInput is an input type that accepts UserPropertyProtocolMapperMap and UserPropertyProtocolMapperMapOutput values.
// You can construct a concrete instance of `UserPropertyProtocolMapperMapInput` via:
//
//	UserPropertyProtocolMapperMap{ "key": UserPropertyProtocolMapperArgs{...} }
type UserPropertyProtocolMapperMapInput interface {
	pulumi.Input

	ToUserPropertyProtocolMapperMapOutput() UserPropertyProtocolMapperMapOutput
	ToUserPropertyProtocolMapperMapOutputWithContext(context.Context) UserPropertyProtocolMapperMapOutput
}

type UserPropertyProtocolMapperMap map[string]UserPropertyProtocolMapperInput

func (UserPropertyProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPropertyProtocolMapper)(nil)).Elem()
}

func (i UserPropertyProtocolMapperMap) ToUserPropertyProtocolMapperMapOutput() UserPropertyProtocolMapperMapOutput {
	return i.ToUserPropertyProtocolMapperMapOutputWithContext(context.Background())
}

func (i UserPropertyProtocolMapperMap) ToUserPropertyProtocolMapperMapOutputWithContext(ctx context.Context) UserPropertyProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertyProtocolMapperMapOutput)
}

type UserPropertyProtocolMapperOutput struct{ *pulumi.OutputState }

func (UserPropertyProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPropertyProtocolMapper)(nil)).Elem()
}

func (o UserPropertyProtocolMapperOutput) ToUserPropertyProtocolMapperOutput() UserPropertyProtocolMapperOutput {
	return o
}

func (o UserPropertyProtocolMapperOutput) ToUserPropertyProtocolMapperOutputWithContext(ctx context.Context) UserPropertyProtocolMapperOutput {
	return o
}

// Indicates if the property should be a claim in the access token.
func (o UserPropertyProtocolMapperOutput) AddToAccessToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.BoolPtrOutput { return v.AddToAccessToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the property should be a claim in the id token.
func (o UserPropertyProtocolMapperOutput) AddToIdToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.BoolPtrOutput { return v.AddToIdToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the property should appear in the userinfo response body.
func (o UserPropertyProtocolMapperOutput) AddToUserinfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.BoolPtrOutput { return v.AddToUserinfo }).(pulumi.BoolPtrOutput)
}

func (o UserPropertyProtocolMapperOutput) ClaimName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringOutput { return v.ClaimName }).(pulumi.StringOutput)
}

// Claim type used when serializing tokens.
func (o UserPropertyProtocolMapperOutput) ClaimValueType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringPtrOutput { return v.ClaimValueType }).(pulumi.StringPtrOutput)
}

// The mapper's associated client. Cannot be used at the same time as client_scope_id.
func (o UserPropertyProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The mapper's associated client scope. Cannot be used at the same time as client_id.
func (o UserPropertyProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// A human-friendly name that will appear in the Keycloak console.
func (o UserPropertyProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm id where the associated client or client scope exists.
func (o UserPropertyProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

func (o UserPropertyProtocolMapperOutput) UserProperty() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringOutput { return v.UserProperty }).(pulumi.StringOutput)
}

type UserPropertyProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (UserPropertyProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPropertyProtocolMapper)(nil)).Elem()
}

func (o UserPropertyProtocolMapperArrayOutput) ToUserPropertyProtocolMapperArrayOutput() UserPropertyProtocolMapperArrayOutput {
	return o
}

func (o UserPropertyProtocolMapperArrayOutput) ToUserPropertyProtocolMapperArrayOutputWithContext(ctx context.Context) UserPropertyProtocolMapperArrayOutput {
	return o
}

func (o UserPropertyProtocolMapperArrayOutput) Index(i pulumi.IntInput) UserPropertyProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPropertyProtocolMapper {
		return vs[0].([]*UserPropertyProtocolMapper)[vs[1].(int)]
	}).(UserPropertyProtocolMapperOutput)
}

type UserPropertyProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (UserPropertyProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPropertyProtocolMapper)(nil)).Elem()
}

func (o UserPropertyProtocolMapperMapOutput) ToUserPropertyProtocolMapperMapOutput() UserPropertyProtocolMapperMapOutput {
	return o
}

func (o UserPropertyProtocolMapperMapOutput) ToUserPropertyProtocolMapperMapOutputWithContext(ctx context.Context) UserPropertyProtocolMapperMapOutput {
	return o
}

func (o UserPropertyProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) UserPropertyProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPropertyProtocolMapper {
		return vs[0].(map[string]*UserPropertyProtocolMapper)[vs[1].(string)]
	}).(UserPropertyProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertyProtocolMapperInput)(nil)).Elem(), &UserPropertyProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertyProtocolMapperArrayInput)(nil)).Elem(), UserPropertyProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertyProtocolMapperMapInput)(nil)).Elem(), UserPropertyProtocolMapperMap{})
	pulumi.RegisterOutputType(UserPropertyProtocolMapperOutput{})
	pulumi.RegisterOutputType(UserPropertyProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(UserPropertyProtocolMapperMapOutput{})
}
