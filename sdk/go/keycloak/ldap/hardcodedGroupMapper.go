// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing hardcoded group mappers for Keycloak users federated via LDAP.
//
// The LDAP hardcoded group mapper will grant a specified Keycloak group to each Keycloak user linked with LDAP.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/ldap"
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
//			realmGroup, err := keycloak.NewGroup(ctx, "realm_group", &keycloak.GroupArgs{
//				RealmId: realm.ID(),
//				Name:    pulumi.String("my-group"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ldap.NewHardcodedGroupMapper(ctx, "assign_group_to_users", &ldap.HardcodedGroupMapperArgs{
//				RealmId:              realm.ID(),
//				LdapUserFederationId: ldapUserFederation.ID(),
//				Name:                 pulumi.String("assign-group-to-users"),
//				Group:                realmGroup.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
// $ pulumi import keycloak:ldap/hardcodedGroupMapper:HardcodedGroupMapper assign_group_to_users my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
// ```
type HardcodedGroupMapper struct {
	pulumi.CustomResourceState

	// The name of the group which should be assigned to the users.
	Group pulumi.StringOutput `pulumi:"group"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewHardcodedGroupMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedGroupMapper(ctx *pulumi.Context,
	name string, args *HardcodedGroupMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedGroupMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.LdapUserFederationId == nil {
		return nil, errors.New("invalid value for required argument 'LdapUserFederationId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HardcodedGroupMapper
	err := ctx.RegisterResource("keycloak:ldap/hardcodedGroupMapper:HardcodedGroupMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedGroupMapper gets an existing HardcodedGroupMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedGroupMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedGroupMapperState, opts ...pulumi.ResourceOption) (*HardcodedGroupMapper, error) {
	var resource HardcodedGroupMapper
	err := ctx.ReadResource("keycloak:ldap/hardcodedGroupMapper:HardcodedGroupMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedGroupMapper resources.
type hardcodedGroupMapperState struct {
	// The name of the group which should be assigned to the users.
	Group *string `pulumi:"group"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId *string `pulumi:"realmId"`
}

type HardcodedGroupMapperState struct {
	// The name of the group which should be assigned to the users.
	Group pulumi.StringPtrInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringPtrInput
}

func (HardcodedGroupMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedGroupMapperState)(nil)).Elem()
}

type hardcodedGroupMapperArgs struct {
	// The name of the group which should be assigned to the users.
	Group string `pulumi:"group"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm that this LDAP mapper will exist in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a HardcodedGroupMapper resource.
type HardcodedGroupMapperArgs struct {
	// The name of the group which should be assigned to the users.
	Group pulumi.StringInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringInput
}

func (HardcodedGroupMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedGroupMapperArgs)(nil)).Elem()
}

type HardcodedGroupMapperInput interface {
	pulumi.Input

	ToHardcodedGroupMapperOutput() HardcodedGroupMapperOutput
	ToHardcodedGroupMapperOutputWithContext(ctx context.Context) HardcodedGroupMapperOutput
}

func (*HardcodedGroupMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedGroupMapper)(nil)).Elem()
}

func (i *HardcodedGroupMapper) ToHardcodedGroupMapperOutput() HardcodedGroupMapperOutput {
	return i.ToHardcodedGroupMapperOutputWithContext(context.Background())
}

func (i *HardcodedGroupMapper) ToHardcodedGroupMapperOutputWithContext(ctx context.Context) HardcodedGroupMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedGroupMapperOutput)
}

// HardcodedGroupMapperArrayInput is an input type that accepts HardcodedGroupMapperArray and HardcodedGroupMapperArrayOutput values.
// You can construct a concrete instance of `HardcodedGroupMapperArrayInput` via:
//
//	HardcodedGroupMapperArray{ HardcodedGroupMapperArgs{...} }
type HardcodedGroupMapperArrayInput interface {
	pulumi.Input

	ToHardcodedGroupMapperArrayOutput() HardcodedGroupMapperArrayOutput
	ToHardcodedGroupMapperArrayOutputWithContext(context.Context) HardcodedGroupMapperArrayOutput
}

type HardcodedGroupMapperArray []HardcodedGroupMapperInput

func (HardcodedGroupMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedGroupMapper)(nil)).Elem()
}

func (i HardcodedGroupMapperArray) ToHardcodedGroupMapperArrayOutput() HardcodedGroupMapperArrayOutput {
	return i.ToHardcodedGroupMapperArrayOutputWithContext(context.Background())
}

func (i HardcodedGroupMapperArray) ToHardcodedGroupMapperArrayOutputWithContext(ctx context.Context) HardcodedGroupMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedGroupMapperArrayOutput)
}

// HardcodedGroupMapperMapInput is an input type that accepts HardcodedGroupMapperMap and HardcodedGroupMapperMapOutput values.
// You can construct a concrete instance of `HardcodedGroupMapperMapInput` via:
//
//	HardcodedGroupMapperMap{ "key": HardcodedGroupMapperArgs{...} }
type HardcodedGroupMapperMapInput interface {
	pulumi.Input

	ToHardcodedGroupMapperMapOutput() HardcodedGroupMapperMapOutput
	ToHardcodedGroupMapperMapOutputWithContext(context.Context) HardcodedGroupMapperMapOutput
}

type HardcodedGroupMapperMap map[string]HardcodedGroupMapperInput

func (HardcodedGroupMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedGroupMapper)(nil)).Elem()
}

func (i HardcodedGroupMapperMap) ToHardcodedGroupMapperMapOutput() HardcodedGroupMapperMapOutput {
	return i.ToHardcodedGroupMapperMapOutputWithContext(context.Background())
}

func (i HardcodedGroupMapperMap) ToHardcodedGroupMapperMapOutputWithContext(ctx context.Context) HardcodedGroupMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardcodedGroupMapperMapOutput)
}

type HardcodedGroupMapperOutput struct{ *pulumi.OutputState }

func (HardcodedGroupMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardcodedGroupMapper)(nil)).Elem()
}

func (o HardcodedGroupMapperOutput) ToHardcodedGroupMapperOutput() HardcodedGroupMapperOutput {
	return o
}

func (o HardcodedGroupMapperOutput) ToHardcodedGroupMapperOutputWithContext(ctx context.Context) HardcodedGroupMapperOutput {
	return o
}

// The name of the group which should be assigned to the users.
func (o HardcodedGroupMapperOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedGroupMapper) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The ID of the LDAP user federation provider to attach this mapper to.
func (o HardcodedGroupMapperOutput) LdapUserFederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedGroupMapper) pulumi.StringOutput { return v.LdapUserFederationId }).(pulumi.StringOutput)
}

// Display name of this mapper when displayed in the console.
func (o HardcodedGroupMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedGroupMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm that this LDAP mapper will exist in.
func (o HardcodedGroupMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *HardcodedGroupMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type HardcodedGroupMapperArrayOutput struct{ *pulumi.OutputState }

func (HardcodedGroupMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HardcodedGroupMapper)(nil)).Elem()
}

func (o HardcodedGroupMapperArrayOutput) ToHardcodedGroupMapperArrayOutput() HardcodedGroupMapperArrayOutput {
	return o
}

func (o HardcodedGroupMapperArrayOutput) ToHardcodedGroupMapperArrayOutputWithContext(ctx context.Context) HardcodedGroupMapperArrayOutput {
	return o
}

func (o HardcodedGroupMapperArrayOutput) Index(i pulumi.IntInput) HardcodedGroupMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HardcodedGroupMapper {
		return vs[0].([]*HardcodedGroupMapper)[vs[1].(int)]
	}).(HardcodedGroupMapperOutput)
}

type HardcodedGroupMapperMapOutput struct{ *pulumi.OutputState }

func (HardcodedGroupMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HardcodedGroupMapper)(nil)).Elem()
}

func (o HardcodedGroupMapperMapOutput) ToHardcodedGroupMapperMapOutput() HardcodedGroupMapperMapOutput {
	return o
}

func (o HardcodedGroupMapperMapOutput) ToHardcodedGroupMapperMapOutputWithContext(ctx context.Context) HardcodedGroupMapperMapOutput {
	return o
}

func (o HardcodedGroupMapperMapOutput) MapIndex(k pulumi.StringInput) HardcodedGroupMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HardcodedGroupMapper {
		return vs[0].(map[string]*HardcodedGroupMapper)[vs[1].(string)]
	}).(HardcodedGroupMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedGroupMapperInput)(nil)).Elem(), &HardcodedGroupMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedGroupMapperArrayInput)(nil)).Elem(), HardcodedGroupMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HardcodedGroupMapperMapInput)(nil)).Elem(), HardcodedGroupMapperMap{})
	pulumi.RegisterOutputType(HardcodedGroupMapperOutput{})
	pulumi.RegisterOutputType(HardcodedGroupMapperArrayOutput{})
	pulumi.RegisterOutputType(HardcodedGroupMapperMapOutput{})
}
