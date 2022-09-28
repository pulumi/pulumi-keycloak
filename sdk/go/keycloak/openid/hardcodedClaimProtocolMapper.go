// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing hardcoded claim protocol mappers within Keycloak.
//
// Hardcoded claim protocol mappers allow you to define a claim with a hardcoded value.
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
//			_, err = openid.NewHardcodedClaimProtocolMapper(ctx, "hardcodedClaimMapper", &openid.HardcodedClaimProtocolMapperArgs{
//				RealmId:    realm.ID(),
//				ClientId:   openidClient.ID(),
//				ClaimName:  pulumi.String("foo"),
//				ClaimValue: pulumi.String("bar"),
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
//			_, err = openid.NewHardcodedClaimProtocolMapper(ctx, "hardcodedClaimMapper", &openid.HardcodedClaimProtocolMapperArgs{
//				RealmId:       realm.ID(),
//				ClientScopeId: clientScope.ID(),
//				ClaimName:     pulumi.String("foo"),
//				ClaimValue:    pulumi.String("bar"),
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
//	$ pulumi import keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper hardcoded_claim_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
//
// ```sh
//
//	$ pulumi import keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper hardcoded_claim_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
type HardcodedClaimProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringOutput `pulumi:"claimName"`
	// The hardcoded value of the claim.
	ClaimValue pulumi.StringOutput `pulumi:"claimValue"`
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType pulumi.StringPtrOutput `pulumi:"claimValueType"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewHardcodedClaimProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedClaimProtocolMapper(ctx *pulumi.Context,
	name string, args *HardcodedClaimProtocolMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedClaimProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClaimName == nil {
		return nil, errors.New("invalid value for required argument 'ClaimName'")
	}
	if args.ClaimValue == nil {
		return nil, errors.New("invalid value for required argument 'ClaimValue'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource HardcodedClaimProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedClaimProtocolMapper gets an existing HardcodedClaimProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedClaimProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedClaimProtocolMapperState, opts ...pulumi.ResourceOption) (*HardcodedClaimProtocolMapper, error) {
	var resource HardcodedClaimProtocolMapper
	err := ctx.ReadResource("keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedClaimProtocolMapper resources.
type hardcodedClaimProtocolMapperState struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName *string `pulumi:"claimName"`
	// The hardcoded value of the claim.
	ClaimValue *string `pulumi:"claimValue"`
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
}

type HardcodedClaimProtocolMapperState struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringPtrInput
	// The hardcoded value of the claim.
	ClaimValue pulumi.StringPtrInput
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType pulumi.StringPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
}

func (HardcodedClaimProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedClaimProtocolMapperState)(nil)).Elem()
}

type hardcodedClaimProtocolMapperArgs struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName string `pulumi:"claimName"`
	// The hardcoded value of the claim.
	ClaimValue string `pulumi:"claimValue"`
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a HardcodedClaimProtocolMapper resource.
type HardcodedClaimProtocolMapperArgs struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringInput
	// The hardcoded value of the claim.
	ClaimValue pulumi.StringInput
	// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
	ClaimValueType pulumi.StringPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
}

func (HardcodedClaimProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedClaimProtocolMapperArgs)(nil)).Elem()
}

type HardcodedClaimProtocolMapperInput interface {
	pulumi.Input

	ToHardcodedClaimProtocolMapperOutput() HardcodedClaimProtocolMapperOutput
	ToHardcodedClaimProtocolMapperOutputWithContext(ctx context.Context) HardcodedClaimProtocolMapperOutput
}

func (*HardcodedClaimProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedClaimProtocolMapper)(nil)).Elem()
}

func (i *HardcodedClaimProtocolMapper) ToHardcodedClaimProtocolMapperOutput() HardcodedClaimProtocolMapperOutput {
	return i.ToHardcodedClaimProtocolMapperOutputWithContext(context.Background())
}

func (i *HardcodedClaimProtocolMapper) ToHardcodedClaimProtocolMapperOutputWithContext(ctx context.Context) HardcodedClaimProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedClaimProtocolMapperOutput)
}

// HardcodedClaimProtocolMapperArrayInput is an input type that accepts HardcodedClaimProtocolMapperArray and HardcodedClaimProtocolMapperArrayOutput values.
// You can construct a concrete instance of `HardcodedClaimProtocolMapperArrayInput` via:
//
//	HardcodedClaimProtocolMapperArray{ HardcodedClaimProtocolMapperArgs{...} }
type HardcodedClaimProtocolMapperArrayInput interface {
	pulumi.Input

	ToHardcodedClaimProtocolMapperArrayOutput() HardcodedClaimProtocolMapperArrayOutput
	ToHardcodedClaimProtocolMapperArrayOutputWithContext(context.Context) HardcodedClaimProtocolMapperArrayOutput
}

type HardcodedClaimProtocolMapperArray []HardcodedClaimProtocolMapperInput

func (HardcodedClaimProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedClaimProtocolMapper)(nil)).Elem()
}

func (i HardcodedClaimProtocolMapperArray) ToHardcodedClaimProtocolMapperArrayOutput() HardcodedClaimProtocolMapperArrayOutput {
	return i.ToHardcodedClaimProtocolMapperArrayOutputWithContext(context.Background())
}

func (i HardcodedClaimProtocolMapperArray) ToHardcodedClaimProtocolMapperArrayOutputWithContext(ctx context.Context) HardcodedClaimProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedClaimProtocolMapperArrayOutput)
}

// HardcodedClaimProtocolMapperMapInput is an input type that accepts HardcodedClaimProtocolMapperMap and HardcodedClaimProtocolMapperMapOutput values.
// You can construct a concrete instance of `HardcodedClaimProtocolMapperMapInput` via:
//
//	HardcodedClaimProtocolMapperMap{ "key": HardcodedClaimProtocolMapperArgs{...} }
type HardcodedClaimProtocolMapperMapInput interface {
	pulumi.Input

	ToHardcodedClaimProtocolMapperMapOutput() HardcodedClaimProtocolMapperMapOutput
	ToHardcodedClaimProtocolMapperMapOutputWithContext(context.Context) HardcodedClaimProtocolMapperMapOutput
}

type HardcodedClaimProtocolMapperMap map[string]HardcodedClaimProtocolMapperInput

func (HardcodedClaimProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedClaimProtocolMapper)(nil)).Elem()
}

func (i HardcodedClaimProtocolMapperMap) ToHardcodedClaimProtocolMapperMapOutput() HardcodedClaimProtocolMapperMapOutput {
	return i.ToHardcodedClaimProtocolMapperMapOutputWithContext(context.Background())
}

func (i HardcodedClaimProtocolMapperMap) ToHardcodedClaimProtocolMapperMapOutputWithContext(ctx context.Context) HardcodedClaimProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedClaimProtocolMapperMapOutput)
}

type HardcodedClaimProtocolMapperOutput struct{ *pulumi.OutputState }

func (HardcodedClaimProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedClaimProtocolMapper)(nil)).Elem()
}

func (o HardcodedClaimProtocolMapperOutput) ToHardcodedClaimProtocolMapperOutput() HardcodedClaimProtocolMapperOutput {
	return o
}

func (o HardcodedClaimProtocolMapperOutput) ToHardcodedClaimProtocolMapperOutputWithContext(ctx context.Context) HardcodedClaimProtocolMapperOutput {
	return o
}

// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
func (o HardcodedClaimProtocolMapperOutput) AddToAccessToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.BoolPtrOutput { return v.AddToAccessToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
func (o HardcodedClaimProtocolMapperOutput) AddToIdToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.BoolPtrOutput { return v.AddToIdToken }).(pulumi.BoolPtrOutput)
}

// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
func (o HardcodedClaimProtocolMapperOutput) AddToUserinfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.BoolPtrOutput { return v.AddToUserinfo }).(pulumi.BoolPtrOutput)
}

// The name of the claim to insert into a token.
func (o HardcodedClaimProtocolMapperOutput) ClaimName() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.StringOutput { return v.ClaimName }).(pulumi.StringOutput)
}

// The hardcoded value of the claim.
func (o HardcodedClaimProtocolMapperOutput) ClaimValue() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.StringOutput { return v.ClaimValue }).(pulumi.StringOutput)
}

// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
func (o HardcodedClaimProtocolMapperOutput) ClaimValueType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.StringPtrOutput { return v.ClaimValueType }).(pulumi.StringPtrOutput)
}

// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
func (o HardcodedClaimProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
func (o HardcodedClaimProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// The display name of this protocol mapper in the GUI.
func (o HardcodedClaimProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm this protocol mapper exists within.
func (o HardcodedClaimProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedClaimProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type HardcodedClaimProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (HardcodedClaimProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedClaimProtocolMapper)(nil)).Elem()
}

func (o HardcodedClaimProtocolMapperArrayOutput) ToHardcodedClaimProtocolMapperArrayOutput() HardcodedClaimProtocolMapperArrayOutput {
	return o
}

func (o HardcodedClaimProtocolMapperArrayOutput) ToHardcodedClaimProtocolMapperArrayOutputWithContext(ctx context.Context) HardcodedClaimProtocolMapperArrayOutput {
	return o
}

func (o HardcodedClaimProtocolMapperArrayOutput) Index(i pulumi.IntInput) HardcodedClaimProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HardcodedClaimProtocolMapper {
		return vs[0].([]*HardcodedClaimProtocolMapper)[vs[1].(int)]
	}).(HardcodedClaimProtocolMapperOutput)
}

type HardcodedClaimProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (HardcodedClaimProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedClaimProtocolMapper)(nil)).Elem()
}

func (o HardcodedClaimProtocolMapperMapOutput) ToHardcodedClaimProtocolMapperMapOutput() HardcodedClaimProtocolMapperMapOutput {
	return o
}

func (o HardcodedClaimProtocolMapperMapOutput) ToHardcodedClaimProtocolMapperMapOutputWithContext(ctx context.Context) HardcodedClaimProtocolMapperMapOutput {
	return o
}

func (o HardcodedClaimProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) HardcodedClaimProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HardcodedClaimProtocolMapper {
		return vs[0].(map[string]*HardcodedClaimProtocolMapper)[vs[1].(string)]
	}).(HardcodedClaimProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedClaimProtocolMapperInput)(nil)).Elem(), &HardcodedClaimProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedClaimProtocolMapperArrayInput)(nil)).Elem(), HardcodedClaimProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedClaimProtocolMapperMapInput)(nil)).Elem(), HardcodedClaimProtocolMapperMap{})
	pulumi.RegisterOutputType(HardcodedClaimProtocolMapperOutput{})
	pulumi.RegisterOutputType(HardcodedClaimProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(HardcodedClaimProtocolMapperMapOutput{})
}
