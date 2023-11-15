// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing user attribute protocol mappers for SAML clients within Keycloak.
//
// SAML user attribute protocol mappers allow you to map custom attributes defined for a user within Keycloak to an attribute
// in a SAML assertion.
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
//			_, err = saml.NewUserAttributeProtocolMapper(ctx, "samlUserAttributeMapper", &saml.UserAttributeProtocolMapperArgs{
//				RealmId:                 realm.ID(),
//				ClientId:                samlClient.ID(),
//				UserAttribute:           pulumi.String("displayName"),
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
//	$ pulumi import keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper saml_user_attribute_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
//
// ```sh
//
//	$ pulumi import keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper saml_user_attribute_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
type UserAttributeProtocolMapper struct {
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
	// The custom user attribute to map.
	UserAttribute pulumi.StringOutput `pulumi:"userAttribute"`
}

// NewUserAttributeProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserAttributeProtocolMapper(ctx *pulumi.Context,
	name string, args *UserAttributeProtocolMapperArgs, opts ...pulumi.ResourceOption) (*UserAttributeProtocolMapper, error) {
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
	if args.UserAttribute == nil {
		return nil, errors.New("invalid value for required argument 'UserAttribute'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserAttributeProtocolMapper
	err := ctx.RegisterResource("keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserAttributeProtocolMapper gets an existing UserAttributeProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAttributeProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserAttributeProtocolMapperState, opts ...pulumi.ResourceOption) (*UserAttributeProtocolMapper, error) {
	var resource UserAttributeProtocolMapper
	err := ctx.ReadResource("keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserAttributeProtocolMapper resources.
type userAttributeProtocolMapperState struct {
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
	// The custom user attribute to map.
	UserAttribute *string `pulumi:"userAttribute"`
}

type UserAttributeProtocolMapperState struct {
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
	// The custom user attribute to map.
	UserAttribute pulumi.StringPtrInput
}

func (UserAttributeProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userAttributeProtocolMapperState)(nil)).Elem()
}

type userAttributeProtocolMapperArgs struct {
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
	// The custom user attribute to map.
	UserAttribute string `pulumi:"userAttribute"`
}

// The set of arguments for constructing a UserAttributeProtocolMapper resource.
type UserAttributeProtocolMapperArgs struct {
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
	// The custom user attribute to map.
	UserAttribute pulumi.StringInput
}

func (UserAttributeProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userAttributeProtocolMapperArgs)(nil)).Elem()
}

type UserAttributeProtocolMapperInput interface {
	pulumi.Input

	ToUserAttributeProtocolMapperOutput() UserAttributeProtocolMapperOutput
	ToUserAttributeProtocolMapperOutputWithContext(ctx context.Context) UserAttributeProtocolMapperOutput
}

func (*UserAttributeProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAttributeProtocolMapper)(nil)).Elem()
}

func (i *UserAttributeProtocolMapper) ToUserAttributeProtocolMapperOutput() UserAttributeProtocolMapperOutput {
	return i.ToUserAttributeProtocolMapperOutputWithContext(context.Background())
}

func (i *UserAttributeProtocolMapper) ToUserAttributeProtocolMapperOutputWithContext(ctx context.Context) UserAttributeProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAttributeProtocolMapperOutput)
}

// UserAttributeProtocolMapperArrayInput is an input type that accepts UserAttributeProtocolMapperArray and UserAttributeProtocolMapperArrayOutput values.
// You can construct a concrete instance of `UserAttributeProtocolMapperArrayInput` via:
//
//	UserAttributeProtocolMapperArray{ UserAttributeProtocolMapperArgs{...} }
type UserAttributeProtocolMapperArrayInput interface {
	pulumi.Input

	ToUserAttributeProtocolMapperArrayOutput() UserAttributeProtocolMapperArrayOutput
	ToUserAttributeProtocolMapperArrayOutputWithContext(context.Context) UserAttributeProtocolMapperArrayOutput
}

type UserAttributeProtocolMapperArray []UserAttributeProtocolMapperInput

func (UserAttributeProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserAttributeProtocolMapper)(nil)).Elem()
}

func (i UserAttributeProtocolMapperArray) ToUserAttributeProtocolMapperArrayOutput() UserAttributeProtocolMapperArrayOutput {
	return i.ToUserAttributeProtocolMapperArrayOutputWithContext(context.Background())
}

func (i UserAttributeProtocolMapperArray) ToUserAttributeProtocolMapperArrayOutputWithContext(ctx context.Context) UserAttributeProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAttributeProtocolMapperArrayOutput)
}

// UserAttributeProtocolMapperMapInput is an input type that accepts UserAttributeProtocolMapperMap and UserAttributeProtocolMapperMapOutput values.
// You can construct a concrete instance of `UserAttributeProtocolMapperMapInput` via:
//
//	UserAttributeProtocolMapperMap{ "key": UserAttributeProtocolMapperArgs{...} }
type UserAttributeProtocolMapperMapInput interface {
	pulumi.Input

	ToUserAttributeProtocolMapperMapOutput() UserAttributeProtocolMapperMapOutput
	ToUserAttributeProtocolMapperMapOutputWithContext(context.Context) UserAttributeProtocolMapperMapOutput
}

type UserAttributeProtocolMapperMap map[string]UserAttributeProtocolMapperInput

func (UserAttributeProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserAttributeProtocolMapper)(nil)).Elem()
}

func (i UserAttributeProtocolMapperMap) ToUserAttributeProtocolMapperMapOutput() UserAttributeProtocolMapperMapOutput {
	return i.ToUserAttributeProtocolMapperMapOutputWithContext(context.Background())
}

func (i UserAttributeProtocolMapperMap) ToUserAttributeProtocolMapperMapOutputWithContext(ctx context.Context) UserAttributeProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAttributeProtocolMapperMapOutput)
}

type UserAttributeProtocolMapperOutput struct{ *pulumi.OutputState }

func (UserAttributeProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAttributeProtocolMapper)(nil)).Elem()
}

func (o UserAttributeProtocolMapperOutput) ToUserAttributeProtocolMapperOutput() UserAttributeProtocolMapperOutput {
	return o
}

func (o UserAttributeProtocolMapperOutput) ToUserAttributeProtocolMapperOutputWithContext(ctx context.Context) UserAttributeProtocolMapperOutput {
	return o
}

// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
func (o UserAttributeProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAttributeProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
func (o UserAttributeProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAttributeProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// An optional human-friendly name for this attribute.
func (o UserAttributeProtocolMapperOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAttributeProtocolMapper) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// The display name of this protocol mapper in the GUI.
func (o UserAttributeProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAttributeProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm this protocol mapper exists within.
func (o UserAttributeProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAttributeProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The name of the SAML attribute.
func (o UserAttributeProtocolMapperOutput) SamlAttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAttributeProtocolMapper) pulumi.StringOutput { return v.SamlAttributeName }).(pulumi.StringOutput)
}

// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
func (o UserAttributeProtocolMapperOutput) SamlAttributeNameFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAttributeProtocolMapper) pulumi.StringOutput { return v.SamlAttributeNameFormat }).(pulumi.StringOutput)
}

// The custom user attribute to map.
func (o UserAttributeProtocolMapperOutput) UserAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *UserAttributeProtocolMapper) pulumi.StringOutput { return v.UserAttribute }).(pulumi.StringOutput)
}

type UserAttributeProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (UserAttributeProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserAttributeProtocolMapper)(nil)).Elem()
}

func (o UserAttributeProtocolMapperArrayOutput) ToUserAttributeProtocolMapperArrayOutput() UserAttributeProtocolMapperArrayOutput {
	return o
}

func (o UserAttributeProtocolMapperArrayOutput) ToUserAttributeProtocolMapperArrayOutputWithContext(ctx context.Context) UserAttributeProtocolMapperArrayOutput {
	return o
}

func (o UserAttributeProtocolMapperArrayOutput) Index(i pulumi.IntInput) UserAttributeProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserAttributeProtocolMapper {
		return vs[0].([]*UserAttributeProtocolMapper)[vs[1].(int)]
	}).(UserAttributeProtocolMapperOutput)
}

type UserAttributeProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (UserAttributeProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserAttributeProtocolMapper)(nil)).Elem()
}

func (o UserAttributeProtocolMapperMapOutput) ToUserAttributeProtocolMapperMapOutput() UserAttributeProtocolMapperMapOutput {
	return o
}

func (o UserAttributeProtocolMapperMapOutput) ToUserAttributeProtocolMapperMapOutputWithContext(ctx context.Context) UserAttributeProtocolMapperMapOutput {
	return o
}

func (o UserAttributeProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) UserAttributeProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserAttributeProtocolMapper {
		return vs[0].(map[string]*UserAttributeProtocolMapper)[vs[1].(string)]
	}).(UserAttributeProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserAttributeProtocolMapperInput)(nil)).Elem(), &UserAttributeProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAttributeProtocolMapperArrayInput)(nil)).Elem(), UserAttributeProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAttributeProtocolMapperMapInput)(nil)).Elem(), UserAttributeProtocolMapperMap{})
	pulumi.RegisterOutputType(UserAttributeProtocolMapperOutput{})
	pulumi.RegisterOutputType(UserAttributeProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(UserAttributeProtocolMapperMapOutput{})
}
