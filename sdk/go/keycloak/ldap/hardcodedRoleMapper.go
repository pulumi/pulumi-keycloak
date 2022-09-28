// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing hardcoded role mappers for Keycloak users federated via LDAP.
//
// The LDAP hardcoded role mapper will grant a specified Keycloak role to each Keycloak user linked with LDAP.
//
// ## Example Usage
// ### Realm Role)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/ldap"
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
//			ldapUserFederation, err := ldap.NewUserFederation(ctx, "ldapUserFederation", &ldap.UserFederationArgs{
//				RealmId:               realm.ID(),
//				UsernameLdapAttribute: pulumi.String("cn"),
//				RdnLdapAttribute:      pulumi.String("cn"),
//				UuidLdapAttribute:     pulumi.String("entryDN"),
//				UserObjectClasses: pulumi.StringArray{
//					pulumi.String("simpleSecurityObject"),
//					pulumi.String("organizationalRole"),
//				},
//				ConnectionUrl:  pulumi.String("ldap://openldap"),
//				UsersDn:        pulumi.String("dc=example,dc=org"),
//				BindDn:         pulumi.String("cn=admin,dc=example,dc=org"),
//				BindCredential: pulumi.String("admin"),
//			})
//			if err != nil {
//				return err
//			}
//			realmAdminRole, err := keycloak.NewRole(ctx, "realmAdminRole", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				Description: pulumi.String("My Realm Role"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ldap.NewHardcodedRoleMapper(ctx, "assignAdminRoleToAllUsers", &ldap.HardcodedRoleMapperArgs{
//				RealmId:              realm.ID(),
//				LdapUserFederationId: ldapUserFederation.ID(),
//				Role:                 realmAdminRole.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Client Role)
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/ldap"
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
//			ldapUserFederation, err := ldap.NewUserFederation(ctx, "ldapUserFederation", &ldap.UserFederationArgs{
//				RealmId:               realm.ID(),
//				UsernameLdapAttribute: pulumi.String("cn"),
//				RdnLdapAttribute:      pulumi.String("cn"),
//				UuidLdapAttribute:     pulumi.String("entryDN"),
//				UserObjectClasses: pulumi.StringArray{
//					pulumi.String("simpleSecurityObject"),
//					pulumi.String("organizationalRole"),
//				},
//				ConnectionUrl:  pulumi.String("ldap://openldap"),
//				UsersDn:        pulumi.String("dc=example,dc=org"),
//				BindDn:         pulumi.String("cn=admin,dc=example,dc=org"),
//				BindCredential: pulumi.String("admin"),
//			})
//			if err != nil {
//				return err
//			}
//			realmManagement := openid.LookupClientOutput(ctx, openid.GetClientOutputArgs{
//				RealmId:  realm.ID(),
//				ClientId: pulumi.String("realm-management"),
//			}, nil)
//			_, err = ldap.NewHardcodedRoleMapper(ctx, "assignAdminRoleToAllUsers", &ldap.HardcodedRoleMapperArgs{
//				RealmId:              realm.ID(),
//				LdapUserFederationId: ldapUserFederation.ID(),
//				Role: pulumi.All(realmManagement, createClient).ApplyT(func(_args []interface{}) (string, error) {
//					realmManagement := _args[0].(openid.GetClientResult)
//					createClient := _args[1].(GetRoleResult)
//					return fmt.Sprintf("%v.%v", realmManagement.ClientId, createClient.Name), nil
//				}).(pulumi.StringOutput),
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
// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper assign_admin_role_to_all_users my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
//
// ```
type HardcodedRoleMapper struct {
	pulumi.CustomResourceState

	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewHardcodedRoleMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedRoleMapper(ctx *pulumi.Context,
	name string, args *HardcodedRoleMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedRoleMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LdapUserFederationId == nil {
		return nil, errors.New("invalid value for required argument 'LdapUserFederationId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource HardcodedRoleMapper
	err := ctx.RegisterResource("keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedRoleMapper gets an existing HardcodedRoleMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedRoleMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedRoleMapperState, opts ...pulumi.ResourceOption) (*HardcodedRoleMapper, error) {
	var resource HardcodedRoleMapper
	err := ctx.ReadResource("keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedRoleMapper resources.
type hardcodedRoleMapperState struct {
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId *string `pulumi:"realmId"`
	// The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
	Role *string `pulumi:"role"`
}

type HardcodedRoleMapperState struct {
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringPtrInput
	// The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
	Role pulumi.StringPtrInput
}

func (HardcodedRoleMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleMapperState)(nil)).Elem()
}

type hardcodedRoleMapperArgs struct {
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId string `pulumi:"realmId"`
	// The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a HardcodedRoleMapper resource.
type HardcodedRoleMapperArgs struct {
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringInput
	// The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
	Role pulumi.StringInput
}

func (HardcodedRoleMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleMapperArgs)(nil)).Elem()
}

type HardcodedRoleMapperInput interface {
	pulumi.Input

	ToHardcodedRoleMapperOutput() HardcodedRoleMapperOutput
	ToHardcodedRoleMapperOutputWithContext(ctx context.Context) HardcodedRoleMapperOutput
}

func (*HardcodedRoleMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedRoleMapper)(nil)).Elem()
}

func (i *HardcodedRoleMapper) ToHardcodedRoleMapperOutput() HardcodedRoleMapperOutput {
	return i.ToHardcodedRoleMapperOutputWithContext(context.Background())
}

func (i *HardcodedRoleMapper) ToHardcodedRoleMapperOutputWithContext(ctx context.Context) HardcodedRoleMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleMapperOutput)
}

// HardcodedRoleMapperArrayInput is an input type that accepts HardcodedRoleMapperArray and HardcodedRoleMapperArrayOutput values.
// You can construct a concrete instance of `HardcodedRoleMapperArrayInput` via:
//
//	HardcodedRoleMapperArray{ HardcodedRoleMapperArgs{...} }
type HardcodedRoleMapperArrayInput interface {
	pulumi.Input

	ToHardcodedRoleMapperArrayOutput() HardcodedRoleMapperArrayOutput
	ToHardcodedRoleMapperArrayOutputWithContext(context.Context) HardcodedRoleMapperArrayOutput
}

type HardcodedRoleMapperArray []HardcodedRoleMapperInput

func (HardcodedRoleMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedRoleMapper)(nil)).Elem()
}

func (i HardcodedRoleMapperArray) ToHardcodedRoleMapperArrayOutput() HardcodedRoleMapperArrayOutput {
	return i.ToHardcodedRoleMapperArrayOutputWithContext(context.Background())
}

func (i HardcodedRoleMapperArray) ToHardcodedRoleMapperArrayOutputWithContext(ctx context.Context) HardcodedRoleMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleMapperArrayOutput)
}

// HardcodedRoleMapperMapInput is an input type that accepts HardcodedRoleMapperMap and HardcodedRoleMapperMapOutput values.
// You can construct a concrete instance of `HardcodedRoleMapperMapInput` via:
//
//	HardcodedRoleMapperMap{ "key": HardcodedRoleMapperArgs{...} }
type HardcodedRoleMapperMapInput interface {
	pulumi.Input

	ToHardcodedRoleMapperMapOutput() HardcodedRoleMapperMapOutput
	ToHardcodedRoleMapperMapOutputWithContext(context.Context) HardcodedRoleMapperMapOutput
}

type HardcodedRoleMapperMap map[string]HardcodedRoleMapperInput

func (HardcodedRoleMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedRoleMapper)(nil)).Elem()
}

func (i HardcodedRoleMapperMap) ToHardcodedRoleMapperMapOutput() HardcodedRoleMapperMapOutput {
	return i.ToHardcodedRoleMapperMapOutputWithContext(context.Background())
}

func (i HardcodedRoleMapperMap) ToHardcodedRoleMapperMapOutputWithContext(ctx context.Context) HardcodedRoleMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedRoleMapperMapOutput)
}

type HardcodedRoleMapperOutput struct{ *pulumi.OutputState }

func (HardcodedRoleMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedRoleMapper)(nil)).Elem()
}

func (o HardcodedRoleMapperOutput) ToHardcodedRoleMapperOutput() HardcodedRoleMapperOutput {
	return o
}

func (o HardcodedRoleMapperOutput) ToHardcodedRoleMapperOutputWithContext(ctx context.Context) HardcodedRoleMapperOutput {
	return o
}

// The ID of the LDAP user federation provider to attach this mapper to.
func (o HardcodedRoleMapperOutput) LdapUserFederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedRoleMapper) pulumi.StringOutput { return v.LdapUserFederationId }).(pulumi.StringOutput)
}

// Display name of this mapper when displayed in the console.
func (o HardcodedRoleMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedRoleMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm that this LDAP mapper will exist in.
func (o HardcodedRoleMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedRoleMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
func (o HardcodedRoleMapperOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedRoleMapper) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

type HardcodedRoleMapperArrayOutput struct{ *pulumi.OutputState }

func (HardcodedRoleMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedRoleMapper)(nil)).Elem()
}

func (o HardcodedRoleMapperArrayOutput) ToHardcodedRoleMapperArrayOutput() HardcodedRoleMapperArrayOutput {
	return o
}

func (o HardcodedRoleMapperArrayOutput) ToHardcodedRoleMapperArrayOutputWithContext(ctx context.Context) HardcodedRoleMapperArrayOutput {
	return o
}

func (o HardcodedRoleMapperArrayOutput) Index(i pulumi.IntInput) HardcodedRoleMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HardcodedRoleMapper {
		return vs[0].([]*HardcodedRoleMapper)[vs[1].(int)]
	}).(HardcodedRoleMapperOutput)
}

type HardcodedRoleMapperMapOutput struct{ *pulumi.OutputState }

func (HardcodedRoleMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedRoleMapper)(nil)).Elem()
}

func (o HardcodedRoleMapperMapOutput) ToHardcodedRoleMapperMapOutput() HardcodedRoleMapperMapOutput {
	return o
}

func (o HardcodedRoleMapperMapOutput) ToHardcodedRoleMapperMapOutputWithContext(ctx context.Context) HardcodedRoleMapperMapOutput {
	return o
}

func (o HardcodedRoleMapperMapOutput) MapIndex(k pulumi.StringInput) HardcodedRoleMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HardcodedRoleMapper {
		return vs[0].(map[string]*HardcodedRoleMapper)[vs[1].(string)]
	}).(HardcodedRoleMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleMapperInput)(nil)).Elem(), &HardcodedRoleMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleMapperArrayInput)(nil)).Elem(), HardcodedRoleMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedRoleMapperMapInput)(nil)).Elem(), HardcodedRoleMapperMap{})
	pulumi.RegisterOutputType(HardcodedRoleMapperOutput{})
	pulumi.RegisterOutputType(HardcodedRoleMapperArrayOutput{})
	pulumi.RegisterOutputType(HardcodedRoleMapperMapOutput{})
}
