// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing full name protocol mappers within Keycloak.
//
// Full name protocol mappers allow you to map a user's first and last name to the OpenID Connect `name` claim in a token.
//
// Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
// multiple different clients.
//
// ## Example Usage
// ### Client)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/openid"
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
//			_, err = openid.NewFullNameProtocolMapper(ctx, "fullNameMapper", &openid.FullNameProtocolMapperArgs{
//				RealmId:  realm.ID(),
//				ClientId: openidClient.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Client Scope)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/openid"
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
//			_, err = openid.NewFullNameProtocolMapper(ctx, "fullNameMapper", &openid.FullNameProtocolMapperArgs{
//				RealmId:       realm.ID(),
//				ClientScopeId: clientScope.ID(),
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
// Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:openid/fullNameProtocolMapper:FullNameProtocolMapper full_name_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
//
// ```sh
//
//	$ pulumi import keycloak:openid/fullNameProtocolMapper:FullNameProtocolMapper full_name_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
type FullNameProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the user's full name should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the user's full name should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// Indicates if the user's full name should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewFullNameProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewFullNameProtocolMapper(ctx *pulumi.Context,
	name string, args *FullNameProtocolMapperArgs, opts ...pulumi.ResourceOption) (*FullNameProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource FullNameProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/fullNameProtocolMapper:FullNameProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFullNameProtocolMapper gets an existing FullNameProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFullNameProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FullNameProtocolMapperState, opts ...pulumi.ResourceOption) (*FullNameProtocolMapper, error) {
	var resource FullNameProtocolMapper
	err := ctx.ReadResource("keycloak:openid/fullNameProtocolMapper:FullNameProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FullNameProtocolMapper resources.
type fullNameProtocolMapperState struct {
	// Indicates if the user's full name should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the user's full name should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the user's full name should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
}

type FullNameProtocolMapperState struct {
	// Indicates if the user's full name should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the user's full name should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the user's full name should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
}

func (FullNameProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*fullNameProtocolMapperState)(nil)).Elem()
}

type fullNameProtocolMapperArgs struct {
	// Indicates if the user's full name should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the user's full name should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the user's full name should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a FullNameProtocolMapper resource.
type FullNameProtocolMapperArgs struct {
	// Indicates if the user's full name should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the user's full name should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the user's full name should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
}

func (FullNameProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fullNameProtocolMapperArgs)(nil)).Elem()
}

type FullNameProtocolMapperInput interface {
	pulumi.Input

	ToFullNameProtocolMapperOutput() FullNameProtocolMapperOutput
	ToFullNameProtocolMapperOutputWithContext(ctx context.Context) FullNameProtocolMapperOutput
}

func (*FullNameProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**FullNameProtocolMapper)(nil)).Elem()
}

func (i *FullNameProtocolMapper) ToFullNameProtocolMapperOutput() FullNameProtocolMapperOutput {
	return i.ToFullNameProtocolMapperOutputWithContext(context.Background())
}

func (i *FullNameProtocolMapper) ToFullNameProtocolMapperOutputWithContext(ctx context.Context) FullNameProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FullNameProtocolMapperOutput)
}

// FullNameProtocolMapperArrayInput is an input type that accepts FullNameProtocolMapperArray and FullNameProtocolMapperArrayOutput values.
// You can construct a concrete instance of `FullNameProtocolMapperArrayInput` via:
//
//	FullNameProtocolMapperArray{ FullNameProtocolMapperArgs{...} }
type FullNameProtocolMapperArrayInput interface {
	pulumi.Input

	ToFullNameProtocolMapperArrayOutput() FullNameProtocolMapperArrayOutput
	ToFullNameProtocolMapperArrayOutputWithContext(context.Context) FullNameProtocolMapperArrayOutput
}

type FullNameProtocolMapperArray []FullNameProtocolMapperInput

func (FullNameProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FullNameProtocolMapper)(nil)).Elem()
}

func (i FullNameProtocolMapperArray) ToFullNameProtocolMapperArrayOutput() FullNameProtocolMapperArrayOutput {
	return i.ToFullNameProtocolMapperArrayOutputWithContext(context.Background())
}

func (i FullNameProtocolMapperArray) ToFullNameProtocolMapperArrayOutputWithContext(ctx context.Context) FullNameProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FullNameProtocolMapperArrayOutput)
}

// FullNameProtocolMapperMapInput is an input type that accepts FullNameProtocolMapperMap and FullNameProtocolMapperMapOutput values.
// You can construct a concrete instance of `FullNameProtocolMapperMapInput` via:
//
//	FullNameProtocolMapperMap{ "key": FullNameProtocolMapperArgs{...} }
type FullNameProtocolMapperMapInput interface {
	pulumi.Input

	ToFullNameProtocolMapperMapOutput() FullNameProtocolMapperMapOutput
	ToFullNameProtocolMapperMapOutputWithContext(context.Context) FullNameProtocolMapperMapOutput
}

type FullNameProtocolMapperMap map[string]FullNameProtocolMapperInput

func (FullNameProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FullNameProtocolMapper)(nil)).Elem()
}

func (i FullNameProtocolMapperMap) ToFullNameProtocolMapperMapOutput() FullNameProtocolMapperMapOutput {
	return i.ToFullNameProtocolMapperMapOutputWithContext(context.Background())
}

func (i FullNameProtocolMapperMap) ToFullNameProtocolMapperMapOutputWithContext(ctx context.Context) FullNameProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FullNameProtocolMapperMapOutput)
}

type FullNameProtocolMapperOutput struct{ *pulumi.OutputState }

func (FullNameProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FullNameProtocolMapper)(nil)).Elem()
}

func (o FullNameProtocolMapperOutput) ToFullNameProtocolMapperOutput() FullNameProtocolMapperOutput {
	return o
}

func (o FullNameProtocolMapperOutput) ToFullNameProtocolMapperOutputWithContext(ctx context.Context) FullNameProtocolMapperOutput {
	return o
}

// Indicates if the user's full name should be added as a claim to the access token. Defaults to `true`.
func (o FullNameProtocolMapperOutput) AddToAccessToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FullNameProtocolMapper) pulumi.BoolPtrOutput { return v.AddToAccessToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the user's full name should be added as a claim to the id token. Defaults to `true`.
func (o FullNameProtocolMapperOutput) AddToIdToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FullNameProtocolMapper) pulumi.BoolPtrOutput { return v.AddToIdToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the user's full name should be added as a claim to the UserInfo response body. Defaults to `true`.
func (o FullNameProtocolMapperOutput) AddToUserinfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FullNameProtocolMapper) pulumi.BoolPtrOutput { return v.AddToUserinfo }).(pulumi.BoolPtrOutput)
}

// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
func (o FullNameProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FullNameProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
func (o FullNameProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FullNameProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// The display name of this protocol mapper in the GUI.
func (o FullNameProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FullNameProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm this protocol mapper exists within.
func (o FullNameProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *FullNameProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type FullNameProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (FullNameProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FullNameProtocolMapper)(nil)).Elem()
}

func (o FullNameProtocolMapperArrayOutput) ToFullNameProtocolMapperArrayOutput() FullNameProtocolMapperArrayOutput {
	return o
}

func (o FullNameProtocolMapperArrayOutput) ToFullNameProtocolMapperArrayOutputWithContext(ctx context.Context) FullNameProtocolMapperArrayOutput {
	return o
}

func (o FullNameProtocolMapperArrayOutput) Index(i pulumi.IntInput) FullNameProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FullNameProtocolMapper {
		return vs[0].([]*FullNameProtocolMapper)[vs[1].(int)]
	}).(FullNameProtocolMapperOutput)
}

type FullNameProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (FullNameProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FullNameProtocolMapper)(nil)).Elem()
}

func (o FullNameProtocolMapperMapOutput) ToFullNameProtocolMapperMapOutput() FullNameProtocolMapperMapOutput {
	return o
}

func (o FullNameProtocolMapperMapOutput) ToFullNameProtocolMapperMapOutputWithContext(ctx context.Context) FullNameProtocolMapperMapOutput {
	return o
}

func (o FullNameProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) FullNameProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FullNameProtocolMapper {
		return vs[0].(map[string]*FullNameProtocolMapper)[vs[1].(string)]
	}).(FullNameProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FullNameProtocolMapperInput)(nil)).Elem(), &FullNameProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*FullNameProtocolMapperArrayInput)(nil)).Elem(), FullNameProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FullNameProtocolMapperMapInput)(nil)).Elem(), FullNameProtocolMapperMap{})
	pulumi.RegisterOutputType(FullNameProtocolMapperOutput{})
	pulumi.RegisterOutputType(FullNameProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(FullNameProtocolMapperMapOutput{})
}
