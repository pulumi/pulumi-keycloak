// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing user property protocol mappers for SAML clients within Keycloak.
//
// SAML user property protocol mappers allow you to map properties of the Keycloak
// user model to an attribute in a SAML assertion.
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
//				ClientId: pulumi.String("saml-client"),
//				Name:     pulumi.String("saml-client"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = saml.NewUserPropertyProtocolMapper(ctx, "saml_user_property_mapper", &saml.UserPropertyProtocolMapperArgs{
//				RealmId:                 realm.ID(),
//				ClientId:                samlClient.ID(),
//				Name:                    pulumi.String("email-user-property-mapper"),
//				UserProperty:            pulumi.String("email"),
//				SamlAttributeName:       pulumi.String("email"),
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
// $ pulumi import keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper saml_user_property_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
//
// ```sh
// $ pulumi import keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper saml_user_property_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
type UserPropertyProtocolMapper struct {
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
	// The property of the Keycloak user model to map.
	UserProperty pulumi.StringOutput `pulumi:"userProperty"`
}

// NewUserPropertyProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserPropertyProtocolMapper(ctx *pulumi.Context,
	name string, args *UserPropertyProtocolMapperArgs, opts ...pulumi.ResourceOption) (*UserPropertyProtocolMapper, error) {
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
	if args.UserProperty == nil {
		return nil, errors.New("invalid value for required argument 'UserProperty'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserPropertyProtocolMapper
	err := ctx.RegisterResource("keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, args, &resource, opts...)
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
	err := ctx.ReadResource("keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPropertyProtocolMapper resources.
type userPropertyProtocolMapperState struct {
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
	// The property of the Keycloak user model to map.
	UserProperty *string `pulumi:"userProperty"`
}

type UserPropertyProtocolMapperState struct {
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
	// The property of the Keycloak user model to map.
	UserProperty pulumi.StringPtrInput
}

func (UserPropertyProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPropertyProtocolMapperState)(nil)).Elem()
}

type userPropertyProtocolMapperArgs struct {
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
	// The property of the Keycloak user model to map.
	UserProperty string `pulumi:"userProperty"`
}

// The set of arguments for constructing a UserPropertyProtocolMapper resource.
type UserPropertyProtocolMapperArgs struct {
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
	// The property of the Keycloak user model to map.
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

// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
func (o UserPropertyProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
func (o UserPropertyProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// An optional human-friendly name for this attribute.
func (o UserPropertyProtocolMapperOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// The display name of this protocol mapper in the GUI.
func (o UserPropertyProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm this protocol mapper exists within.
func (o UserPropertyProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The name of the SAML attribute.
func (o UserPropertyProtocolMapperOutput) SamlAttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringOutput { return v.SamlAttributeName }).(pulumi.StringOutput)
}

// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
func (o UserPropertyProtocolMapperOutput) SamlAttributeNameFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPropertyProtocolMapper) pulumi.StringOutput { return v.SamlAttributeNameFormat }).(pulumi.StringOutput)
}

// The property of the Keycloak user model to map.
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
