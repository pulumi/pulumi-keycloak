// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing role mappers for Keycloak users federated via LDAP.
//
// The LDAP group mapper can be used to map an LDAP user's roles from some DN to Keycloak roles.
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
//				ConnectionUrl:  pulumi.String("ldap://openldap"),
//				UsersDn:        pulumi.String("dc=example,dc=org"),
//				BindDn:         pulumi.String("cn=admin,dc=example,dc=org"),
//				BindCredential: pulumi.String("admin"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ldap.NewRoleMapper(ctx, "ldap_role_mapper", &ldap.RoleMapperArgs{
//				RealmId:               realm.ID(),
//				LdapUserFederationId:  ldapUserFederation.ID(),
//				Name:                  pulumi.String("role-mapper"),
//				LdapRolesDn:           pulumi.String("dc=example,dc=org"),
//				RoleNameLdapAttribute: pulumi.String("cn"),
//				RoleObjectClasses: pulumi.StringArray{
//					pulumi.String("groupOfNames"),
//				},
//				MembershipAttributeType:     pulumi.String("DN"),
//				MembershipLdapAttribute:     pulumi.String("member"),
//				MembershipUserLdapAttribute: pulumi.String("cn"),
//				UserRolesRetrieveStrategy:   pulumi.String("GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE"),
//				MemberofLdapAttribute:       pulumi.String("memberOf"),
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
// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.
//
// The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs.
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:ldap/roleMapper:RoleMapper ldap_role_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
// ```
type RoleMapper struct {
	pulumi.CustomResourceState

	// When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `useRealmRolesMapping` is `false`.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The LDAP DN where roles can be found.
	LdapRolesDn pulumi.StringOutput `pulumi:"ldapRolesDn"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
	MemberofLdapAttribute pulumi.StringPtrOutput `pulumi:"memberofLdapAttribute"`
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType pulumi.StringPtrOutput `pulumi:"membershipAttributeType"`
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute pulumi.StringOutput `pulumi:"membershipLdapAttribute"`
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute pulumi.StringOutput `pulumi:"membershipUserLdapAttribute"`
	// Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
	RoleNameLdapAttribute pulumi.StringOutput `pulumi:"roleNameLdapAttribute"`
	// List of strings representing the object classes for the role. Must contain at least one.
	RoleObjectClasses pulumi.StringArrayOutput `pulumi:"roleObjectClasses"`
	// When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
	RolesLdapFilter pulumi.StringPtrOutput `pulumi:"rolesLdapFilter"`
	// When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
	UseRealmRolesMapping pulumi.BoolPtrOutput `pulumi:"useRealmRolesMapping"`
	// Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy pulumi.StringPtrOutput `pulumi:"userRolesRetrieveStrategy"`
}

// NewRoleMapper registers a new resource with the given unique name, arguments, and options.
func NewRoleMapper(ctx *pulumi.Context,
	name string, args *RoleMapperArgs, opts ...pulumi.ResourceOption) (*RoleMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LdapRolesDn == nil {
		return nil, errors.New("invalid value for required argument 'LdapRolesDn'")
	}
	if args.LdapUserFederationId == nil {
		return nil, errors.New("invalid value for required argument 'LdapUserFederationId'")
	}
	if args.MembershipLdapAttribute == nil {
		return nil, errors.New("invalid value for required argument 'MembershipLdapAttribute'")
	}
	if args.MembershipUserLdapAttribute == nil {
		return nil, errors.New("invalid value for required argument 'MembershipUserLdapAttribute'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.RoleNameLdapAttribute == nil {
		return nil, errors.New("invalid value for required argument 'RoleNameLdapAttribute'")
	}
	if args.RoleObjectClasses == nil {
		return nil, errors.New("invalid value for required argument 'RoleObjectClasses'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RoleMapper
	err := ctx.RegisterResource("keycloak:ldap/roleMapper:RoleMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleMapper gets an existing RoleMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleMapperState, opts ...pulumi.ResourceOption) (*RoleMapper, error) {
	var resource RoleMapper
	err := ctx.ReadResource("keycloak:ldap/roleMapper:RoleMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleMapper resources.
type roleMapperState struct {
	// When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `useRealmRolesMapping` is `false`.
	ClientId *string `pulumi:"clientId"`
	// The LDAP DN where roles can be found.
	LdapRolesDn *string `pulumi:"ldapRolesDn"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
	MemberofLdapAttribute *string `pulumi:"memberofLdapAttribute"`
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType *string `pulumi:"membershipAttributeType"`
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute *string `pulumi:"membershipLdapAttribute"`
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute *string `pulumi:"membershipUserLdapAttribute"`
	// Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
	Mode *string `pulumi:"mode"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId *string `pulumi:"realmId"`
	// The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
	RoleNameLdapAttribute *string `pulumi:"roleNameLdapAttribute"`
	// List of strings representing the object classes for the role. Must contain at least one.
	RoleObjectClasses []string `pulumi:"roleObjectClasses"`
	// When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
	RolesLdapFilter *string `pulumi:"rolesLdapFilter"`
	// When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
	UseRealmRolesMapping *bool `pulumi:"useRealmRolesMapping"`
	// Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy *string `pulumi:"userRolesRetrieveStrategy"`
}

type RoleMapperState struct {
	// When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `useRealmRolesMapping` is `false`.
	ClientId pulumi.StringPtrInput
	// The LDAP DN where roles can be found.
	LdapRolesDn pulumi.StringPtrInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
	MemberofLdapAttribute pulumi.StringPtrInput
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType pulumi.StringPtrInput
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute pulumi.StringPtrInput
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute pulumi.StringPtrInput
	// Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
	Mode pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringPtrInput
	// The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
	RoleNameLdapAttribute pulumi.StringPtrInput
	// List of strings representing the object classes for the role. Must contain at least one.
	RoleObjectClasses pulumi.StringArrayInput
	// When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
	RolesLdapFilter pulumi.StringPtrInput
	// When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
	UseRealmRolesMapping pulumi.BoolPtrInput
	// Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy pulumi.StringPtrInput
}

func (RoleMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleMapperState)(nil)).Elem()
}

type roleMapperArgs struct {
	// When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `useRealmRolesMapping` is `false`.
	ClientId *string `pulumi:"clientId"`
	// The LDAP DN where roles can be found.
	LdapRolesDn string `pulumi:"ldapRolesDn"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
	MemberofLdapAttribute *string `pulumi:"memberofLdapAttribute"`
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType *string `pulumi:"membershipAttributeType"`
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute string `pulumi:"membershipLdapAttribute"`
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute string `pulumi:"membershipUserLdapAttribute"`
	// Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
	Mode *string `pulumi:"mode"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId string `pulumi:"realmId"`
	// The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
	RoleNameLdapAttribute string `pulumi:"roleNameLdapAttribute"`
	// List of strings representing the object classes for the role. Must contain at least one.
	RoleObjectClasses []string `pulumi:"roleObjectClasses"`
	// When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
	RolesLdapFilter *string `pulumi:"rolesLdapFilter"`
	// When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
	UseRealmRolesMapping *bool `pulumi:"useRealmRolesMapping"`
	// Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy *string `pulumi:"userRolesRetrieveStrategy"`
}

// The set of arguments for constructing a RoleMapper resource.
type RoleMapperArgs struct {
	// When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `useRealmRolesMapping` is `false`.
	ClientId pulumi.StringPtrInput
	// The LDAP DN where roles can be found.
	LdapRolesDn pulumi.StringInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
	MemberofLdapAttribute pulumi.StringPtrInput
	// Can be one of `DN` or `UID`. Defaults to `DN`.
	MembershipAttributeType pulumi.StringPtrInput
	// The name of the LDAP attribute that is used for membership mappings.
	MembershipLdapAttribute pulumi.StringInput
	// The name of the LDAP attribute on a user that is used for membership mappings.
	MembershipUserLdapAttribute pulumi.StringInput
	// Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
	Mode pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringInput
	// The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
	RoleNameLdapAttribute pulumi.StringInput
	// List of strings representing the object classes for the role. Must contain at least one.
	RoleObjectClasses pulumi.StringArrayInput
	// When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
	RolesLdapFilter pulumi.StringPtrInput
	// When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
	UseRealmRolesMapping pulumi.BoolPtrInput
	// Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
	UserRolesRetrieveStrategy pulumi.StringPtrInput
}

func (RoleMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleMapperArgs)(nil)).Elem()
}

type RoleMapperInput interface {
	pulumi.Input

	ToRoleMapperOutput() RoleMapperOutput
	ToRoleMapperOutputWithContext(ctx context.Context) RoleMapperOutput
}

func (*RoleMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleMapper)(nil)).Elem()
}

func (i *RoleMapper) ToRoleMapperOutput() RoleMapperOutput {
	return i.ToRoleMapperOutputWithContext(context.Background())
}

func (i *RoleMapper) ToRoleMapperOutputWithContext(ctx context.Context) RoleMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapperOutput)
}

// RoleMapperArrayInput is an input type that accepts RoleMapperArray and RoleMapperArrayOutput values.
// You can construct a concrete instance of `RoleMapperArrayInput` via:
//
//	RoleMapperArray{ RoleMapperArgs{...} }
type RoleMapperArrayInput interface {
	pulumi.Input

	ToRoleMapperArrayOutput() RoleMapperArrayOutput
	ToRoleMapperArrayOutputWithContext(context.Context) RoleMapperArrayOutput
}

type RoleMapperArray []RoleMapperInput

func (RoleMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleMapper)(nil)).Elem()
}

func (i RoleMapperArray) ToRoleMapperArrayOutput() RoleMapperArrayOutput {
	return i.ToRoleMapperArrayOutputWithContext(context.Background())
}

func (i RoleMapperArray) ToRoleMapperArrayOutputWithContext(ctx context.Context) RoleMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapperArrayOutput)
}

// RoleMapperMapInput is an input type that accepts RoleMapperMap and RoleMapperMapOutput values.
// You can construct a concrete instance of `RoleMapperMapInput` via:
//
//	RoleMapperMap{ "key": RoleMapperArgs{...} }
type RoleMapperMapInput interface {
	pulumi.Input

	ToRoleMapperMapOutput() RoleMapperMapOutput
	ToRoleMapperMapOutputWithContext(context.Context) RoleMapperMapOutput
}

type RoleMapperMap map[string]RoleMapperInput

func (RoleMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleMapper)(nil)).Elem()
}

func (i RoleMapperMap) ToRoleMapperMapOutput() RoleMapperMapOutput {
	return i.ToRoleMapperMapOutputWithContext(context.Background())
}

func (i RoleMapperMap) ToRoleMapperMapOutputWithContext(ctx context.Context) RoleMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapperMapOutput)
}

type RoleMapperOutput struct{ *pulumi.OutputState }

func (RoleMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleMapper)(nil)).Elem()
}

func (o RoleMapperOutput) ToRoleMapperOutput() RoleMapperOutput {
	return o
}

func (o RoleMapperOutput) ToRoleMapperOutputWithContext(ctx context.Context) RoleMapperOutput {
	return o
}

// When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `useRealmRolesMapping` is `false`.
func (o RoleMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The LDAP DN where roles can be found.
func (o RoleMapperOutput) LdapRolesDn() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringOutput { return v.LdapRolesDn }).(pulumi.StringOutput)
}

// The ID of the LDAP user federation provider to attach this mapper to.
func (o RoleMapperOutput) LdapUserFederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringOutput { return v.LdapUserFederationId }).(pulumi.StringOutput)
}

// Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
func (o RoleMapperOutput) MemberofLdapAttribute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringPtrOutput { return v.MemberofLdapAttribute }).(pulumi.StringPtrOutput)
}

// Can be one of `DN` or `UID`. Defaults to `DN`.
func (o RoleMapperOutput) MembershipAttributeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringPtrOutput { return v.MembershipAttributeType }).(pulumi.StringPtrOutput)
}

// The name of the LDAP attribute that is used for membership mappings.
func (o RoleMapperOutput) MembershipLdapAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringOutput { return v.MembershipLdapAttribute }).(pulumi.StringOutput)
}

// The name of the LDAP attribute on a user that is used for membership mappings.
func (o RoleMapperOutput) MembershipUserLdapAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringOutput { return v.MembershipUserLdapAttribute }).(pulumi.StringOutput)
}

// Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
func (o RoleMapperOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// Display name of this mapper when displayed in the console.
func (o RoleMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm that this LDAP mapper will exist in.
func (o RoleMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
func (o RoleMapperOutput) RoleNameLdapAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringOutput { return v.RoleNameLdapAttribute }).(pulumi.StringOutput)
}

// List of strings representing the object classes for the role. Must contain at least one.
func (o RoleMapperOutput) RoleObjectClasses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringArrayOutput { return v.RoleObjectClasses }).(pulumi.StringArrayOutput)
}

// When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
func (o RoleMapperOutput) RolesLdapFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringPtrOutput { return v.RolesLdapFilter }).(pulumi.StringPtrOutput)
}

// When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
func (o RoleMapperOutput) UseRealmRolesMapping() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.BoolPtrOutput { return v.UseRealmRolesMapping }).(pulumi.BoolPtrOutput)
}

// Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
func (o RoleMapperOutput) UserRolesRetrieveStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleMapper) pulumi.StringPtrOutput { return v.UserRolesRetrieveStrategy }).(pulumi.StringPtrOutput)
}

type RoleMapperArrayOutput struct{ *pulumi.OutputState }

func (RoleMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleMapper)(nil)).Elem()
}

func (o RoleMapperArrayOutput) ToRoleMapperArrayOutput() RoleMapperArrayOutput {
	return o
}

func (o RoleMapperArrayOutput) ToRoleMapperArrayOutputWithContext(ctx context.Context) RoleMapperArrayOutput {
	return o
}

func (o RoleMapperArrayOutput) Index(i pulumi.IntInput) RoleMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoleMapper {
		return vs[0].([]*RoleMapper)[vs[1].(int)]
	}).(RoleMapperOutput)
}

type RoleMapperMapOutput struct{ *pulumi.OutputState }

func (RoleMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleMapper)(nil)).Elem()
}

func (o RoleMapperMapOutput) ToRoleMapperMapOutput() RoleMapperMapOutput {
	return o
}

func (o RoleMapperMapOutput) ToRoleMapperMapOutputWithContext(ctx context.Context) RoleMapperMapOutput {
	return o
}

func (o RoleMapperMapOutput) MapIndex(k pulumi.StringInput) RoleMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoleMapper {
		return vs[0].(map[string]*RoleMapper)[vs[1].(string)]
	}).(RoleMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapperInput)(nil)).Elem(), &RoleMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapperArrayInput)(nil)).Elem(), RoleMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapperMapInput)(nil)).Elem(), RoleMapperMap{})
	pulumi.RegisterOutputType(RoleMapperOutput{})
	pulumi.RegisterOutputType(RoleMapperArrayOutput{})
	pulumi.RegisterOutputType(RoleMapperMapOutput{})
}
