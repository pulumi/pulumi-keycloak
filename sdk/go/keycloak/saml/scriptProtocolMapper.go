// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing script protocol mappers for SAML clients within Keycloak.
//
// Script protocol mappers evaluate a JavaScript function to produce an attribute value based on context information.
//
// Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
// multiple different clients.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/saml"
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
//			samlClient, err := saml.NewClient(ctx, "samlClient", &saml.ClientArgs{
//				RealmId:  realm.ID(),
//				ClientId: pulumi.String("saml-client"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = saml.NewScriptProtocolMapper(ctx, "samlScriptMapper", &saml.ScriptProtocolMapperArgs{
//				RealmId:                 realm.ID(),
//				ClientId:                samlClient.ID(),
//				Script:                  pulumi.String("exports = 'foo';"),
//				SamlAttributeName:       pulumi.String("displayName"),
//				SamlAttributeNameFormat: pulumi.String("Unspecified"),
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
//	$ pulumi import keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper saml_script_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
//
// ```sh
//
//	$ pulumi import keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper saml_script_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
type ScriptProtocolMapper struct {
	pulumi.CustomResourceState

	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// An optional human-friendly name for this attribute.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The name of the SAML attribute.
	SamlAttributeName pulumi.StringOutput `pulumi:"samlAttributeName"`
	// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
	SamlAttributeNameFormat pulumi.StringOutput `pulumi:"samlAttributeNameFormat"`
	// JavaScript code to compute the attribute value.
	Script pulumi.StringOutput `pulumi:"script"`
	// When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
	SingleValueAttribute pulumi.BoolPtrOutput `pulumi:"singleValueAttribute"`
}

// NewScriptProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewScriptProtocolMapper(ctx *pulumi.Context,
	name string, args *ScriptProtocolMapperArgs, opts ...pulumi.ResourceOption) (*ScriptProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.SamlAttributeName == nil {
		return nil, errors.New("invalid value for required argument 'SamlAttributeName'")
	}
	if args.SamlAttributeNameFormat == nil {
		return nil, errors.New("invalid value for required argument 'SamlAttributeNameFormat'")
	}
	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	var resource ScriptProtocolMapper
	err := ctx.RegisterResource("keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper", name, args, &resource, opts...)
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
	err := ctx.ReadResource("keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScriptProtocolMapper resources.
type scriptProtocolMapperState struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// An optional human-friendly name for this attribute.
	FriendlyName *string `pulumi:"friendlyName"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
	// The name of the SAML attribute.
	SamlAttributeName *string `pulumi:"samlAttributeName"`
	// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
	SamlAttributeNameFormat *string `pulumi:"samlAttributeNameFormat"`
	// JavaScript code to compute the attribute value.
	Script *string `pulumi:"script"`
	// When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
	SingleValueAttribute *bool `pulumi:"singleValueAttribute"`
}

type ScriptProtocolMapperState struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// An optional human-friendly name for this attribute.
	FriendlyName pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
	// The name of the SAML attribute.
	SamlAttributeName pulumi.StringPtrInput
	// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
	SamlAttributeNameFormat pulumi.StringPtrInput
	// JavaScript code to compute the attribute value.
	Script pulumi.StringPtrInput
	// When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
	SingleValueAttribute pulumi.BoolPtrInput
}

func (ScriptProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptProtocolMapperState)(nil)).Elem()
}

type scriptProtocolMapperArgs struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// An optional human-friendly name for this attribute.
	FriendlyName *string `pulumi:"friendlyName"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
	// The name of the SAML attribute.
	SamlAttributeName string `pulumi:"samlAttributeName"`
	// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
	SamlAttributeNameFormat string `pulumi:"samlAttributeNameFormat"`
	// JavaScript code to compute the attribute value.
	Script string `pulumi:"script"`
	// When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
	SingleValueAttribute *bool `pulumi:"singleValueAttribute"`
}

// The set of arguments for constructing a ScriptProtocolMapper resource.
type ScriptProtocolMapperArgs struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// An optional human-friendly name for this attribute.
	FriendlyName pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
	// The name of the SAML attribute.
	SamlAttributeName pulumi.StringInput
	// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
	SamlAttributeNameFormat pulumi.StringInput
	// JavaScript code to compute the attribute value.
	Script pulumi.StringInput
	// When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
	SingleValueAttribute pulumi.BoolPtrInput
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
//	ScriptProtocolMapperArray{ ScriptProtocolMapperArgs{...} }
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
//	ScriptProtocolMapperMap{ "key": ScriptProtocolMapperArgs{...} }
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

// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
func (o ScriptProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
func (o ScriptProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// An optional human-friendly name for this attribute.
func (o ScriptProtocolMapperOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptProtocolMapper) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// The display name of this protocol mapper in the GUI.
func (o ScriptProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm this protocol mapper exists within.
func (o ScriptProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The name of the SAML attribute.
func (o ScriptProtocolMapperOutput) SamlAttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptProtocolMapper) pulumi.StringOutput { return v.SamlAttributeName }).(pulumi.StringOutput)
}

// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
func (o ScriptProtocolMapperOutput) SamlAttributeNameFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptProtocolMapper) pulumi.StringOutput { return v.SamlAttributeNameFormat }).(pulumi.StringOutput)
}

// JavaScript code to compute the attribute value.
func (o ScriptProtocolMapperOutput) Script() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptProtocolMapper) pulumi.StringOutput { return v.Script }).(pulumi.StringOutput)
}

// When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
func (o ScriptProtocolMapperOutput) SingleValueAttribute() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScriptProtocolMapper) pulumi.BoolPtrOutput { return v.SingleValueAttribute }).(pulumi.BoolPtrOutput)
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
