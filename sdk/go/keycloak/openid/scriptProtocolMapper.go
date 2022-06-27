// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing script protocol mappers within Keycloak.
//
// Script protocol mappers evaluate a JavaScript function to produce a token claim based on context information.
//
// Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
// multiple different clients.
//
// > Support for this protocol mapper was removed in Keycloak 18.
//
// ## Example Usage
// ### Client)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/openid"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("my-realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		openidClient, err := openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
// 			RealmId:    realm.ID(),
// 			ClientId:   pulumi.String("client"),
// 			Enabled:    pulumi.Bool(true),
// 			AccessType: pulumi.String("CONFIDENTIAL"),
// 			ValidRedirectUris: pulumi.StringArray{
// 				pulumi.String("http://localhost:8080/openid-callback"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewScriptProtocolMapper(ctx, "scriptMapper", &openid.ScriptProtocolMapperArgs{
// 			RealmId:   realm.ID(),
// 			ClientId:  openidClient.ID(),
// 			ClaimName: pulumi.String("foo"),
// 			Script:    pulumi.String("exports = 'foo';"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Client Scope)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/openid"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("my-realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		clientScope, err := openid.NewClientScope(ctx, "clientScope", &openid.ClientScopeArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewScriptProtocolMapper(ctx, "scriptMapper", &openid.ScriptProtocolMapperArgs{
// 			RealmId:       realm.ID(),
// 			ClientScopeId: clientScope.ID(),
// 			ClaimName:     pulumi.String("foo"),
// 			Script:        pulumi.String("exports = 'foo';"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
//
// ```sh
//  $ pulumi import keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper script_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
//
// ```sh
//  $ pulumi import keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper script_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
type ScriptProtocolMapper struct {
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
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued pulumi.BoolPtrOutput `pulumi:"multivalued"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// JavaScript code to compute the claim value.
	Script pulumi.StringOutput `pulumi:"script"`
}

// NewScriptProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewScriptProtocolMapper(ctx *pulumi.Context,
	name string, args *ScriptProtocolMapperArgs, opts ...pulumi.ResourceOption) (*ScriptProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClaimName == nil {
		return nil, errors.New("invalid value for required argument 'ClaimName'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	var resource ScriptProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScriptProtocolMapper gets an existing ScriptProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScriptProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScriptProtocolMapperState, opts ...pulumi.ResourceOption) (*ScriptProtocolMapper, error) {
	var resource ScriptProtocolMapper
	err := ctx.ReadResource("keycloak:openid/scriptProtocolMapper:ScriptProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScriptProtocolMapper resources.
type scriptProtocolMapperState struct {
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
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued *bool `pulumi:"multivalued"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
	// JavaScript code to compute the claim value.
	Script *string `pulumi:"script"`
}

type ScriptProtocolMapperState struct {
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
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued pulumi.BoolPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
	// JavaScript code to compute the claim value.
	Script pulumi.StringPtrInput
}

func (ScriptProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptProtocolMapperState)(nil)).Elem()
}

type scriptProtocolMapperArgs struct {
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
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued *bool `pulumi:"multivalued"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
	// JavaScript code to compute the claim value.
	Script string `pulumi:"script"`
}

// The set of arguments for constructing a ScriptProtocolMapper resource.
type ScriptProtocolMapperArgs struct {
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
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `false`.
	Multivalued pulumi.BoolPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
	// JavaScript code to compute the claim value.
	Script pulumi.StringInput
}

func (ScriptProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptProtocolMapperArgs)(nil)).Elem()
}

type ScriptProtocolMapperInput interface {
	pulumi.Input

	ToScriptProtocolMapperOutput() ScriptProtocolMapperOutput
	ToScriptProtocolMapperOutputWithContext(ctx context.Context) ScriptProtocolMapperOutput
}

func (*ScriptProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptProtocolMapper)(nil)).Elem()
}

func (i *ScriptProtocolMapper) ToScriptProtocolMapperOutput() ScriptProtocolMapperOutput {
	return i.ToScriptProtocolMapperOutputWithContext(context.Background())
}

func (i *ScriptProtocolMapper) ToScriptProtocolMapperOutputWithContext(ctx context.Context) ScriptProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptProtocolMapperOutput)
}

// ScriptProtocolMapperArrayInput is an input type that accepts ScriptProtocolMapperArray and ScriptProtocolMapperArrayOutput values.
// You can construct a concrete instance of `ScriptProtocolMapperArrayInput` via:
//
//          ScriptProtocolMapperArray{ ScriptProtocolMapperArgs{...} }
type ScriptProtocolMapperArrayInput interface {
	pulumi.Input

	ToScriptProtocolMapperArrayOutput() ScriptProtocolMapperArrayOutput
	ToScriptProtocolMapperArrayOutputWithContext(context.Context) ScriptProtocolMapperArrayOutput
}

type ScriptProtocolMapperArray []ScriptProtocolMapperInput

func (ScriptProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScriptProtocolMapper)(nil)).Elem()
}

func (i ScriptProtocolMapperArray) ToScriptProtocolMapperArrayOutput() ScriptProtocolMapperArrayOutput {
	return i.ToScriptProtocolMapperArrayOutputWithContext(context.Background())
}

func (i ScriptProtocolMapperArray) ToScriptProtocolMapperArrayOutputWithContext(ctx context.Context) ScriptProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptProtocolMapperArrayOutput)
}

// ScriptProtocolMapperMapInput is an input type that accepts ScriptProtocolMapperMap and ScriptProtocolMapperMapOutput values.
// You can construct a concrete instance of `ScriptProtocolMapperMapInput` via:
//
//          ScriptProtocolMapperMap{ "key": ScriptProtocolMapperArgs{...} }
type ScriptProtocolMapperMapInput interface {
	pulumi.Input

	ToScriptProtocolMapperMapOutput() ScriptProtocolMapperMapOutput
	ToScriptProtocolMapperMapOutputWithContext(context.Context) ScriptProtocolMapperMapOutput
}

type ScriptProtocolMapperMap map[string]ScriptProtocolMapperInput

func (ScriptProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScriptProtocolMapper)(nil)).Elem()
}

func (i ScriptProtocolMapperMap) ToScriptProtocolMapperMapOutput() ScriptProtocolMapperMapOutput {
	return i.ToScriptProtocolMapperMapOutputWithContext(context.Background())
}

func (i ScriptProtocolMapperMap) ToScriptProtocolMapperMapOutputWithContext(ctx context.Context) ScriptProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptProtocolMapperMapOutput)
}

type ScriptProtocolMapperOutput struct{ *pulumi.OutputState }

func (ScriptProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptProtocolMapper)(nil)).Elem()
}

func (o ScriptProtocolMapperOutput) ToScriptProtocolMapperOutput() ScriptProtocolMapperOutput {
	return o
}

func (o ScriptProtocolMapperOutput) ToScriptProtocolMapperOutputWithContext(ctx context.Context) ScriptProtocolMapperOutput {
	return o
}

type ScriptProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (ScriptProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScriptProtocolMapper)(nil)).Elem()
}

func (o ScriptProtocolMapperArrayOutput) ToScriptProtocolMapperArrayOutput() ScriptProtocolMapperArrayOutput {
	return o
}

func (o ScriptProtocolMapperArrayOutput) ToScriptProtocolMapperArrayOutputWithContext(ctx context.Context) ScriptProtocolMapperArrayOutput {
	return o
}

func (o ScriptProtocolMapperArrayOutput) Index(i pulumi.IntInput) ScriptProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScriptProtocolMapper {
		return vs[0].([]*ScriptProtocolMapper)[vs[1].(int)]
	}).(ScriptProtocolMapperOutput)
}

type ScriptProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (ScriptProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScriptProtocolMapper)(nil)).Elem()
}

func (o ScriptProtocolMapperMapOutput) ToScriptProtocolMapperMapOutput() ScriptProtocolMapperMapOutput {
	return o
}

func (o ScriptProtocolMapperMapOutput) ToScriptProtocolMapperMapOutputWithContext(ctx context.Context) ScriptProtocolMapperMapOutput {
	return o
}

func (o ScriptProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) ScriptProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScriptProtocolMapper {
		return vs[0].(map[string]*ScriptProtocolMapper)[vs[1].(string)]
	}).(ScriptProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptProtocolMapperInput)(nil)).Elem(), &ScriptProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptProtocolMapperArrayInput)(nil)).Elem(), ScriptProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptProtocolMapperMapInput)(nil)).Elem(), ScriptProtocolMapperMap{})
	pulumi.RegisterOutputType(ScriptProtocolMapperOutput{})
	pulumi.RegisterOutputType(ScriptProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(ScriptProtocolMapperMapOutput{})
}
