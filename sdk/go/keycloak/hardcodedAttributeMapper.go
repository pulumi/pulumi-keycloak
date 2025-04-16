// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing hardcoded attribute mappers for Keycloak users federated via LDAP.
//
// The user model hardcoded attribute mapper will set the specified value to the attribute.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/ldap"
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
//			ldapUserFederation, err := ldap.NewUserFederation(ctx, "ldap_user_federation", &ldap.UserFederationArgs{
//				Name:                  pulumi.String("openldap"),
//				RealmId:               realm.ID(),
//				UsernameLdapAttribute: pulumi.String("cn"),
//				RdnLdapAttribute:      pulumi.String("cn"),
//				UuidLdapAttribute:     pulumi.String("entryDN"),
//				UserObjectClasses: pulumi.StringArray{
//					pulumi.String("simpleSecurityObject"),
//					pulumi.String("organizationalRole"),
//				},
//				ConnectionUrl:     pulumi.String("ldap://openldap"),
//				UsersDn:           pulumi.String("dc=example,dc=org"),
//				BindDn:            pulumi.String("cn=admin,dc=example,dc=org"),
//				BindCredential:    pulumi.String("admin"),
//				SyncRegistrations: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewHardcodedAttributeMapper(ctx, "email_verified", &keycloak.HardcodedAttributeMapperArgs{
//				RealmId:              realm.ID(),
//				LdapUserFederationId: ldapUserFederation.ID(),
//				Name:                 pulumi.String("email_verified"),
//				AttributeName:        pulumi.String("email_verified"),
//				AttributeValue:       pulumi.String("true"),
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
// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{attribute__mapper_id}}`.
//
// The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs.
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:index/hardcodedAttributeMapper:HardcodedAttributeMapper email_verified my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
// ```
type HardcodedAttributeMapper struct {
	pulumi.CustomResourceState

	// The name of the user model attribute to set.
	AttributeName pulumi.StringOutput `pulumi:"attributeName"`
	// The value to set to model attribute. You can hardcode any value like 'foo'.
	AttributeValue pulumi.StringOutput `pulumi:"attributeValue"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewHardcodedAttributeMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedAttributeMapper(ctx *pulumi.Context,
	name string, args *HardcodedAttributeMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedAttributeMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AttributeName == nil {
		return nil, errors.New("invalid value for required argument 'AttributeName'")
	}
	if args.AttributeValue == nil {
		return nil, errors.New("invalid value for required argument 'AttributeValue'")
	}
	if args.LdapUserFederationId == nil {
		return nil, errors.New("invalid value for required argument 'LdapUserFederationId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HardcodedAttributeMapper
	err := ctx.RegisterResource("keycloak:index/hardcodedAttributeMapper:HardcodedAttributeMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedAttributeMapper gets an existing HardcodedAttributeMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedAttributeMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedAttributeMapperState, opts ...pulumi.ResourceOption) (*HardcodedAttributeMapper, error) {
	var resource HardcodedAttributeMapper
	err := ctx.ReadResource("keycloak:index/hardcodedAttributeMapper:HardcodedAttributeMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedAttributeMapper resources.
type hardcodedAttributeMapperState struct {
	// The name of the user model attribute to set.
	AttributeName *string `pulumi:"attributeName"`
	// The value to set to model attribute. You can hardcode any value like 'foo'.
	AttributeValue *string `pulumi:"attributeValue"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId *string `pulumi:"realmId"`
}

type HardcodedAttributeMapperState struct {
	// The name of the user model attribute to set.
	AttributeName pulumi.StringPtrInput
	// The value to set to model attribute. You can hardcode any value like 'foo'.
	AttributeValue pulumi.StringPtrInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringPtrInput
}

func (HardcodedAttributeMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedAttributeMapperState)(nil)).Elem()
}

type hardcodedAttributeMapperArgs struct {
	// The name of the user model attribute to set.
	AttributeName string `pulumi:"attributeName"`
	// The value to set to model attribute. You can hardcode any value like 'foo'.
	AttributeValue string `pulumi:"attributeValue"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a HardcodedAttributeMapper resource.
type HardcodedAttributeMapperArgs struct {
	// The name of the user model attribute to set.
	AttributeName pulumi.StringInput
	// The value to set to model attribute. You can hardcode any value like 'foo'.
	AttributeValue pulumi.StringInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringInput
}

func (HardcodedAttributeMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedAttributeMapperArgs)(nil)).Elem()
}

type HardcodedAttributeMapperInput interface {
	pulumi.Input

	ToHardcodedAttributeMapperOutput() HardcodedAttributeMapperOutput
	ToHardcodedAttributeMapperOutputWithContext(ctx context.Context) HardcodedAttributeMapperOutput
}

func (*HardcodedAttributeMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedAttributeMapper)(nil)).Elem()
}

func (i *HardcodedAttributeMapper) ToHardcodedAttributeMapperOutput() HardcodedAttributeMapperOutput {
	return i.ToHardcodedAttributeMapperOutputWithContext(context.Background())
}

func (i *HardcodedAttributeMapper) ToHardcodedAttributeMapperOutputWithContext(ctx context.Context) HardcodedAttributeMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedAttributeMapperOutput)
}

// HardcodedAttributeMapperArrayInput is an input type that accepts HardcodedAttributeMapperArray and HardcodedAttributeMapperArrayOutput values.
// You can construct a concrete instance of `HardcodedAttributeMapperArrayInput` via:
//
//	HardcodedAttributeMapperArray{ HardcodedAttributeMapperArgs{...} }
type HardcodedAttributeMapperArrayInput interface {
	pulumi.Input

	ToHardcodedAttributeMapperArrayOutput() HardcodedAttributeMapperArrayOutput
	ToHardcodedAttributeMapperArrayOutputWithContext(context.Context) HardcodedAttributeMapperArrayOutput
}

type HardcodedAttributeMapperArray []HardcodedAttributeMapperInput

func (HardcodedAttributeMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedAttributeMapper)(nil)).Elem()
}

func (i HardcodedAttributeMapperArray) ToHardcodedAttributeMapperArrayOutput() HardcodedAttributeMapperArrayOutput {
	return i.ToHardcodedAttributeMapperArrayOutputWithContext(context.Background())
}

func (i HardcodedAttributeMapperArray) ToHardcodedAttributeMapperArrayOutputWithContext(ctx context.Context) HardcodedAttributeMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedAttributeMapperArrayOutput)
}

// HardcodedAttributeMapperMapInput is an input type that accepts HardcodedAttributeMapperMap and HardcodedAttributeMapperMapOutput values.
// You can construct a concrete instance of `HardcodedAttributeMapperMapInput` via:
//
//	HardcodedAttributeMapperMap{ "key": HardcodedAttributeMapperArgs{...} }
type HardcodedAttributeMapperMapInput interface {
	pulumi.Input

	ToHardcodedAttributeMapperMapOutput() HardcodedAttributeMapperMapOutput
	ToHardcodedAttributeMapperMapOutputWithContext(context.Context) HardcodedAttributeMapperMapOutput
}

type HardcodedAttributeMapperMap map[string]HardcodedAttributeMapperInput

func (HardcodedAttributeMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedAttributeMapper)(nil)).Elem()
}

func (i HardcodedAttributeMapperMap) ToHardcodedAttributeMapperMapOutput() HardcodedAttributeMapperMapOutput {
	return i.ToHardcodedAttributeMapperMapOutputWithContext(context.Background())
}

func (i HardcodedAttributeMapperMap) ToHardcodedAttributeMapperMapOutputWithContext(ctx context.Context) HardcodedAttributeMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedAttributeMapperMapOutput)
}

type HardcodedAttributeMapperOutput struct{ *pulumi.OutputState }

func (HardcodedAttributeMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedAttributeMapper)(nil)).Elem()
}

func (o HardcodedAttributeMapperOutput) ToHardcodedAttributeMapperOutput() HardcodedAttributeMapperOutput {
	return o
}

func (o HardcodedAttributeMapperOutput) ToHardcodedAttributeMapperOutputWithContext(ctx context.Context) HardcodedAttributeMapperOutput {
	return o
}

// The name of the user model attribute to set.
func (o HardcodedAttributeMapperOutput) AttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedAttributeMapper) pulumi.StringOutput { return v.AttributeName }).(pulumi.StringOutput)
}

// The value to set to model attribute. You can hardcode any value like 'foo'.
func (o HardcodedAttributeMapperOutput) AttributeValue() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedAttributeMapper) pulumi.StringOutput { return v.AttributeValue }).(pulumi.StringOutput)
}

// The ID of the LDAP user federation provider to attach this mapper to.
func (o HardcodedAttributeMapperOutput) LdapUserFederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedAttributeMapper) pulumi.StringOutput { return v.LdapUserFederationId }).(pulumi.StringOutput)
}

// Display name of this mapper when displayed in the console.
func (o HardcodedAttributeMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedAttributeMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm that this LDAP mapper will exist in.
func (o HardcodedAttributeMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedAttributeMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type HardcodedAttributeMapperArrayOutput struct{ *pulumi.OutputState }

func (HardcodedAttributeMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedAttributeMapper)(nil)).Elem()
}

func (o HardcodedAttributeMapperArrayOutput) ToHardcodedAttributeMapperArrayOutput() HardcodedAttributeMapperArrayOutput {
	return o
}

func (o HardcodedAttributeMapperArrayOutput) ToHardcodedAttributeMapperArrayOutputWithContext(ctx context.Context) HardcodedAttributeMapperArrayOutput {
	return o
}

func (o HardcodedAttributeMapperArrayOutput) Index(i pulumi.IntInput) HardcodedAttributeMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HardcodedAttributeMapper {
		return vs[0].([]*HardcodedAttributeMapper)[vs[1].(int)]
	}).(HardcodedAttributeMapperOutput)
}

type HardcodedAttributeMapperMapOutput struct{ *pulumi.OutputState }

func (HardcodedAttributeMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedAttributeMapper)(nil)).Elem()
}

func (o HardcodedAttributeMapperMapOutput) ToHardcodedAttributeMapperMapOutput() HardcodedAttributeMapperMapOutput {
	return o
}

func (o HardcodedAttributeMapperMapOutput) ToHardcodedAttributeMapperMapOutputWithContext(ctx context.Context) HardcodedAttributeMapperMapOutput {
	return o
}

func (o HardcodedAttributeMapperMapOutput) MapIndex(k pulumi.StringInput) HardcodedAttributeMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HardcodedAttributeMapper {
		return vs[0].(map[string]*HardcodedAttributeMapper)[vs[1].(string)]
	}).(HardcodedAttributeMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedAttributeMapperInput)(nil)).Elem(), &HardcodedAttributeMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedAttributeMapperArrayInput)(nil)).Elem(), HardcodedAttributeMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedAttributeMapperMapInput)(nil)).Elem(), HardcodedAttributeMapperMap{})
	pulumi.RegisterOutputType(HardcodedAttributeMapperOutput{})
	pulumi.RegisterOutputType(HardcodedAttributeMapperArrayOutput{})
	pulumi.RegisterOutputType(HardcodedAttributeMapperMapOutput{})
}
