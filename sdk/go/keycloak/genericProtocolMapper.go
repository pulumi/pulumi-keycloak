// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing protocol mappers for both types of clients (openid-connect and saml) within Keycloak.
//
// There are two uses cases for using this resource:
// * If you implemented a custom protocol mapper, this resource can be used to configure it
// * If the provider doesn't support a particular protocol mapper, this resource can be used instead.
//
// Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors.
// Therefore, if possible, a specific mapper should be used instead.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/saml"
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
//			samlClient, err := saml.NewClient(ctx, "saml_client", &saml.ClientArgs{
//				RealmId:  realm.ID(),
//				ClientId: pulumi.String("test-client"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewGenericProtocolMapper(ctx, "saml_hardcode_attribute_mapper", &keycloak.GenericProtocolMapperArgs{
//				RealmId:        realm.ID(),
//				ClientId:       samlClient.ID(),
//				Name:           pulumi.String("test-mapper"),
//				Protocol:       pulumi.String("saml"),
//				ProtocolMapper: pulumi.String("saml-hardcode-attribute-mapper"),
//				Config: pulumi.StringMap{
//					"attribute.name":       pulumi.String("name"),
//					"attribute.nameformat": pulumi.String("Basic"),
//					"attribute.value":      pulumi.String("value"),
//					"friendly.name":        pulumi.String("display name"),
//				},
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
// Protocol mappers can be imported using the following format: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:index/genericProtocolMapper:GenericProtocolMapper saml_hardcode_attribute_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
type GenericProtocolMapper struct {
	pulumi.CustomResourceState

	// The ID of the client this protocol mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The ID of the client scope this protocol mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
	Config pulumi.StringMapOutput `pulumi:"config"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
	ProtocolMapper pulumi.StringOutput `pulumi:"protocolMapper"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewGenericProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewGenericProtocolMapper(ctx *pulumi.Context,
	name string, args *GenericProtocolMapperArgs, opts ...pulumi.ResourceOption) (*GenericProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ProtocolMapper == nil {
		return nil, errors.New("invalid value for required argument 'ProtocolMapper'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GenericProtocolMapper
	err := ctx.RegisterResource("keycloak:index/genericProtocolMapper:GenericProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGenericProtocolMapper gets an existing GenericProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGenericProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GenericProtocolMapperState, opts ...pulumi.ResourceOption) (*GenericProtocolMapper, error) {
	var resource GenericProtocolMapper
	err := ctx.ReadResource("keycloak:index/genericProtocolMapper:GenericProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GenericProtocolMapper resources.
type genericProtocolMapperState struct {
	// The ID of the client this protocol mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId *string `pulumi:"clientId"`
	// The ID of the client scope this protocol mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
	Config map[string]string `pulumi:"config"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
	Protocol *string `pulumi:"protocol"`
	// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
	ProtocolMapper *string `pulumi:"protocolMapper"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
}

type GenericProtocolMapperState struct {
	// The ID of the client this protocol mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId pulumi.StringPtrInput
	// The ID of the client scope this protocol mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId pulumi.StringPtrInput
	// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
	Config pulumi.StringMapInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
	Protocol pulumi.StringPtrInput
	// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
	ProtocolMapper pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
}

func (GenericProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*genericProtocolMapperState)(nil)).Elem()
}

type genericProtocolMapperArgs struct {
	// The ID of the client this protocol mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId *string `pulumi:"clientId"`
	// The ID of the client scope this protocol mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
	Config map[string]string `pulumi:"config"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
	Protocol string `pulumi:"protocol"`
	// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
	ProtocolMapper string `pulumi:"protocolMapper"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a GenericProtocolMapper resource.
type GenericProtocolMapperArgs struct {
	// The ID of the client this protocol mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
	ClientId pulumi.StringPtrInput
	// The ID of the client scope this protocol mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
	ClientScopeId pulumi.StringPtrInput
	// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
	Config pulumi.StringMapInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
	Protocol pulumi.StringInput
	// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
	ProtocolMapper pulumi.StringInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
}

func (GenericProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*genericProtocolMapperArgs)(nil)).Elem()
}

type GenericProtocolMapperInput interface {
	pulumi.Input

	ToGenericProtocolMapperOutput() GenericProtocolMapperOutput
	ToGenericProtocolMapperOutputWithContext(ctx context.Context) GenericProtocolMapperOutput
}

func (*GenericProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericProtocolMapper)(nil)).Elem()
}

func (i *GenericProtocolMapper) ToGenericProtocolMapperOutput() GenericProtocolMapperOutput {
	return i.ToGenericProtocolMapperOutputWithContext(context.Background())
}

func (i *GenericProtocolMapper) ToGenericProtocolMapperOutputWithContext(ctx context.Context) GenericProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericProtocolMapperOutput)
}

// GenericProtocolMapperArrayInput is an input type that accepts GenericProtocolMapperArray and GenericProtocolMapperArrayOutput values.
// You can construct a concrete instance of `GenericProtocolMapperArrayInput` via:
//
//	GenericProtocolMapperArray{ GenericProtocolMapperArgs{...} }
type GenericProtocolMapperArrayInput interface {
	pulumi.Input

	ToGenericProtocolMapperArrayOutput() GenericProtocolMapperArrayOutput
	ToGenericProtocolMapperArrayOutputWithContext(context.Context) GenericProtocolMapperArrayOutput
}

type GenericProtocolMapperArray []GenericProtocolMapperInput

func (GenericProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GenericProtocolMapper)(nil)).Elem()
}

func (i GenericProtocolMapperArray) ToGenericProtocolMapperArrayOutput() GenericProtocolMapperArrayOutput {
	return i.ToGenericProtocolMapperArrayOutputWithContext(context.Background())
}

func (i GenericProtocolMapperArray) ToGenericProtocolMapperArrayOutputWithContext(ctx context.Context) GenericProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericProtocolMapperArrayOutput)
}

// GenericProtocolMapperMapInput is an input type that accepts GenericProtocolMapperMap and GenericProtocolMapperMapOutput values.
// You can construct a concrete instance of `GenericProtocolMapperMapInput` via:
//
//	GenericProtocolMapperMap{ "key": GenericProtocolMapperArgs{...} }
type GenericProtocolMapperMapInput interface {
	pulumi.Input

	ToGenericProtocolMapperMapOutput() GenericProtocolMapperMapOutput
	ToGenericProtocolMapperMapOutputWithContext(context.Context) GenericProtocolMapperMapOutput
}

type GenericProtocolMapperMap map[string]GenericProtocolMapperInput

func (GenericProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GenericProtocolMapper)(nil)).Elem()
}

func (i GenericProtocolMapperMap) ToGenericProtocolMapperMapOutput() GenericProtocolMapperMapOutput {
	return i.ToGenericProtocolMapperMapOutputWithContext(context.Background())
}

func (i GenericProtocolMapperMap) ToGenericProtocolMapperMapOutputWithContext(ctx context.Context) GenericProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GenericProtocolMapperMapOutput)
}

type GenericProtocolMapperOutput struct{ *pulumi.OutputState }

func (GenericProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GenericProtocolMapper)(nil)).Elem()
}

func (o GenericProtocolMapperOutput) ToGenericProtocolMapperOutput() GenericProtocolMapperOutput {
	return o
}

func (o GenericProtocolMapperOutput) ToGenericProtocolMapperOutputWithContext(ctx context.Context) GenericProtocolMapperOutput {
	return o
}

// The ID of the client this protocol mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
func (o GenericProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GenericProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The ID of the client scope this protocol mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
func (o GenericProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GenericProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
func (o GenericProtocolMapperOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GenericProtocolMapper) pulumi.StringMapOutput { return v.Config }).(pulumi.StringMapOutput)
}

// The display name of this protocol mapper in the GUI.
func (o GenericProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GenericProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
func (o GenericProtocolMapperOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *GenericProtocolMapper) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
func (o GenericProtocolMapperOutput) ProtocolMapper() pulumi.StringOutput {
	return o.ApplyT(func(v *GenericProtocolMapper) pulumi.StringOutput { return v.ProtocolMapper }).(pulumi.StringOutput)
}

// The realm this protocol mapper exists within.
func (o GenericProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *GenericProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type GenericProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (GenericProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GenericProtocolMapper)(nil)).Elem()
}

func (o GenericProtocolMapperArrayOutput) ToGenericProtocolMapperArrayOutput() GenericProtocolMapperArrayOutput {
	return o
}

func (o GenericProtocolMapperArrayOutput) ToGenericProtocolMapperArrayOutputWithContext(ctx context.Context) GenericProtocolMapperArrayOutput {
	return o
}

func (o GenericProtocolMapperArrayOutput) Index(i pulumi.IntInput) GenericProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GenericProtocolMapper {
		return vs[0].([]*GenericProtocolMapper)[vs[1].(int)]
	}).(GenericProtocolMapperOutput)
}

type GenericProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (GenericProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GenericProtocolMapper)(nil)).Elem()
}

func (o GenericProtocolMapperMapOutput) ToGenericProtocolMapperMapOutput() GenericProtocolMapperMapOutput {
	return o
}

func (o GenericProtocolMapperMapOutput) ToGenericProtocolMapperMapOutputWithContext(ctx context.Context) GenericProtocolMapperMapOutput {
	return o
}

func (o GenericProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) GenericProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GenericProtocolMapper {
		return vs[0].(map[string]*GenericProtocolMapper)[vs[1].(string)]
	}).(GenericProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GenericProtocolMapperInput)(nil)).Elem(), &GenericProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenericProtocolMapperArrayInput)(nil)).Elem(), GenericProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GenericProtocolMapperMapInput)(nil)).Elem(), GenericProtocolMapperMap{})
	pulumi.RegisterOutputType(GenericProtocolMapperOutput{})
	pulumi.RegisterOutputType(GenericProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(GenericProtocolMapperMapOutput{})
}
