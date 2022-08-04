// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing MSAD-LDS user account control mappers for Keycloak
// users federated via LDAP.
//
// The MSAD-LDS (Microsoft Active Directory Lightweight Directory Service) user account control mapper is specific
// to LDAP user federation providers that are pulling from AD-LDS, and it can propagate
// AD-LDS user state to Keycloak in order to enforce settings like expired passwords
// or disabled accounts.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/ldap"
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
// 		ldapUserFederation, err := ldap.NewUserFederation(ctx, "ldapUserFederation", &ldap.UserFederationArgs{
// 			RealmId:               realm.ID(),
// 			UsernameLdapAttribute: pulumi.String("cn"),
// 			RdnLdapAttribute:      pulumi.String("cn"),
// 			UuidLdapAttribute:     pulumi.String("objectGUID"),
// 			UserObjectClasses: pulumi.StringArray{
// 				pulumi.String("person"),
// 				pulumi.String("organizationalPerson"),
// 				pulumi.String("user"),
// 			},
// 			ConnectionUrl:  pulumi.String("ldap://my-ad-server"),
// 			UsersDn:        pulumi.String("dc=example,dc=org"),
// 			BindDn:         pulumi.String("cn=admin,dc=example,dc=org"),
// 			BindCredential: pulumi.String("admin"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ldap.NewMsadLdsUserAccountControlMapper(ctx, "msadLdsUserAccountControlMapper", &ldap.MsadLdsUserAccountControlMapperArgs{
// 			RealmId:              realm.ID(),
// 			LdapUserFederationId: ldapUserFederation.ID(),
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
// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
//
// ```sh
//  $ pulumi import keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper msad_lds_user_account_control_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
// ```
type MsadLdsUserAccountControlMapper struct {
	pulumi.CustomResourceState

	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewMsadLdsUserAccountControlMapper registers a new resource with the given unique name, arguments, and options.
func NewMsadLdsUserAccountControlMapper(ctx *pulumi.Context,
	name string, args *MsadLdsUserAccountControlMapperArgs, opts ...pulumi.ResourceOption) (*MsadLdsUserAccountControlMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LdapUserFederationId == nil {
		return nil, errors.New("invalid value for required argument 'LdapUserFederationId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource MsadLdsUserAccountControlMapper
	err := ctx.RegisterResource("keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMsadLdsUserAccountControlMapper gets an existing MsadLdsUserAccountControlMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMsadLdsUserAccountControlMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MsadLdsUserAccountControlMapperState, opts ...pulumi.ResourceOption) (*MsadLdsUserAccountControlMapper, error) {
	var resource MsadLdsUserAccountControlMapper
	err := ctx.ReadResource("keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MsadLdsUserAccountControlMapper resources.
type msadLdsUserAccountControlMapperState struct {
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId *string `pulumi:"realmId"`
}

type MsadLdsUserAccountControlMapperState struct {
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringPtrInput
}

func (MsadLdsUserAccountControlMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*msadLdsUserAccountControlMapperState)(nil)).Elem()
}

type msadLdsUserAccountControlMapperArgs struct {
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a MsadLdsUserAccountControlMapper resource.
type MsadLdsUserAccountControlMapperArgs struct {
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringInput
}

func (MsadLdsUserAccountControlMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*msadLdsUserAccountControlMapperArgs)(nil)).Elem()
}

type MsadLdsUserAccountControlMapperInput interface {
	pulumi.Input

	ToMsadLdsUserAccountControlMapperOutput() MsadLdsUserAccountControlMapperOutput
	ToMsadLdsUserAccountControlMapperOutputWithContext(ctx context.Context) MsadLdsUserAccountControlMapperOutput
}

func (*MsadLdsUserAccountControlMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**MsadLdsUserAccountControlMapper)(nil)).Elem()
}

func (i *MsadLdsUserAccountControlMapper) ToMsadLdsUserAccountControlMapperOutput() MsadLdsUserAccountControlMapperOutput {
	return i.ToMsadLdsUserAccountControlMapperOutputWithContext(context.Background())
}

func (i *MsadLdsUserAccountControlMapper) ToMsadLdsUserAccountControlMapperOutputWithContext(ctx context.Context) MsadLdsUserAccountControlMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsadLdsUserAccountControlMapperOutput)
}

// MsadLdsUserAccountControlMapperArrayInput is an input type that accepts MsadLdsUserAccountControlMapperArray and MsadLdsUserAccountControlMapperArrayOutput values.
// You can construct a concrete instance of `MsadLdsUserAccountControlMapperArrayInput` via:
//
//          MsadLdsUserAccountControlMapperArray{ MsadLdsUserAccountControlMapperArgs{...} }
type MsadLdsUserAccountControlMapperArrayInput interface {
	pulumi.Input

	ToMsadLdsUserAccountControlMapperArrayOutput() MsadLdsUserAccountControlMapperArrayOutput
	ToMsadLdsUserAccountControlMapperArrayOutputWithContext(context.Context) MsadLdsUserAccountControlMapperArrayOutput
}

type MsadLdsUserAccountControlMapperArray []MsadLdsUserAccountControlMapperInput

func (MsadLdsUserAccountControlMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MsadLdsUserAccountControlMapper)(nil)).Elem()
}

func (i MsadLdsUserAccountControlMapperArray) ToMsadLdsUserAccountControlMapperArrayOutput() MsadLdsUserAccountControlMapperArrayOutput {
	return i.ToMsadLdsUserAccountControlMapperArrayOutputWithContext(context.Background())
}

func (i MsadLdsUserAccountControlMapperArray) ToMsadLdsUserAccountControlMapperArrayOutputWithContext(ctx context.Context) MsadLdsUserAccountControlMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsadLdsUserAccountControlMapperArrayOutput)
}

// MsadLdsUserAccountControlMapperMapInput is an input type that accepts MsadLdsUserAccountControlMapperMap and MsadLdsUserAccountControlMapperMapOutput values.
// You can construct a concrete instance of `MsadLdsUserAccountControlMapperMapInput` via:
//
//          MsadLdsUserAccountControlMapperMap{ "key": MsadLdsUserAccountControlMapperArgs{...} }
type MsadLdsUserAccountControlMapperMapInput interface {
	pulumi.Input

	ToMsadLdsUserAccountControlMapperMapOutput() MsadLdsUserAccountControlMapperMapOutput
	ToMsadLdsUserAccountControlMapperMapOutputWithContext(context.Context) MsadLdsUserAccountControlMapperMapOutput
}

type MsadLdsUserAccountControlMapperMap map[string]MsadLdsUserAccountControlMapperInput

func (MsadLdsUserAccountControlMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MsadLdsUserAccountControlMapper)(nil)).Elem()
}

func (i MsadLdsUserAccountControlMapperMap) ToMsadLdsUserAccountControlMapperMapOutput() MsadLdsUserAccountControlMapperMapOutput {
	return i.ToMsadLdsUserAccountControlMapperMapOutputWithContext(context.Background())
}

func (i MsadLdsUserAccountControlMapperMap) ToMsadLdsUserAccountControlMapperMapOutputWithContext(ctx context.Context) MsadLdsUserAccountControlMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsadLdsUserAccountControlMapperMapOutput)
}

type MsadLdsUserAccountControlMapperOutput struct{ *pulumi.OutputState }

func (MsadLdsUserAccountControlMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MsadLdsUserAccountControlMapper)(nil)).Elem()
}

func (o MsadLdsUserAccountControlMapperOutput) ToMsadLdsUserAccountControlMapperOutput() MsadLdsUserAccountControlMapperOutput {
	return o
}

func (o MsadLdsUserAccountControlMapperOutput) ToMsadLdsUserAccountControlMapperOutputWithContext(ctx context.Context) MsadLdsUserAccountControlMapperOutput {
	return o
}

// The ID of the LDAP user federation provider to attach this mapper to.
func (o MsadLdsUserAccountControlMapperOutput) LdapUserFederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *MsadLdsUserAccountControlMapper) pulumi.StringOutput { return v.LdapUserFederationId }).(pulumi.StringOutput)
}

// Display name of this mapper when displayed in the console.
func (o MsadLdsUserAccountControlMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MsadLdsUserAccountControlMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm that this LDAP mapper will exist in.
func (o MsadLdsUserAccountControlMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *MsadLdsUserAccountControlMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type MsadLdsUserAccountControlMapperArrayOutput struct{ *pulumi.OutputState }

func (MsadLdsUserAccountControlMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MsadLdsUserAccountControlMapper)(nil)).Elem()
}

func (o MsadLdsUserAccountControlMapperArrayOutput) ToMsadLdsUserAccountControlMapperArrayOutput() MsadLdsUserAccountControlMapperArrayOutput {
	return o
}

func (o MsadLdsUserAccountControlMapperArrayOutput) ToMsadLdsUserAccountControlMapperArrayOutputWithContext(ctx context.Context) MsadLdsUserAccountControlMapperArrayOutput {
	return o
}

func (o MsadLdsUserAccountControlMapperArrayOutput) Index(i pulumi.IntInput) MsadLdsUserAccountControlMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MsadLdsUserAccountControlMapper {
		return vs[0].([]*MsadLdsUserAccountControlMapper)[vs[1].(int)]
	}).(MsadLdsUserAccountControlMapperOutput)
}

type MsadLdsUserAccountControlMapperMapOutput struct{ *pulumi.OutputState }

func (MsadLdsUserAccountControlMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MsadLdsUserAccountControlMapper)(nil)).Elem()
}

func (o MsadLdsUserAccountControlMapperMapOutput) ToMsadLdsUserAccountControlMapperMapOutput() MsadLdsUserAccountControlMapperMapOutput {
	return o
}

func (o MsadLdsUserAccountControlMapperMapOutput) ToMsadLdsUserAccountControlMapperMapOutputWithContext(ctx context.Context) MsadLdsUserAccountControlMapperMapOutput {
	return o
}

func (o MsadLdsUserAccountControlMapperMapOutput) MapIndex(k pulumi.StringInput) MsadLdsUserAccountControlMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MsadLdsUserAccountControlMapper {
		return vs[0].(map[string]*MsadLdsUserAccountControlMapper)[vs[1].(string)]
	}).(MsadLdsUserAccountControlMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MsadLdsUserAccountControlMapperInput)(nil)).Elem(), &MsadLdsUserAccountControlMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsadLdsUserAccountControlMapperArrayInput)(nil)).Elem(), MsadLdsUserAccountControlMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsadLdsUserAccountControlMapperMapInput)(nil)).Elem(), MsadLdsUserAccountControlMapperMap{})
	pulumi.RegisterOutputType(MsadLdsUserAccountControlMapperOutput{})
	pulumi.RegisterOutputType(MsadLdsUserAccountControlMapperArrayOutput{})
	pulumi.RegisterOutputType(MsadLdsUserAccountControlMapperMapOutput{})
}
