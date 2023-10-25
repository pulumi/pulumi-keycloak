// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Allows for creating and managing full name mappers for Keycloak users federated via LDAP.
//
// The LDAP full name mapper can map a user's full name from an LDAP attribute to the first and last name attributes of a
// Keycloak user.
//
// ## Import
//
// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:ldap/fullNameMapper:FullNameMapper ldap_full_name_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
//
// ```
type FullNameMapper struct {
	pulumi.CustomResourceState

	// The name of the LDAP attribute containing the user's full name.
	LdapFullNameAttribute pulumi.StringOutput `pulumi:"ldapFullNameAttribute"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
	ReadOnly pulumi.BoolPtrOutput `pulumi:"readOnly"`
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
	WriteOnly pulumi.BoolPtrOutput `pulumi:"writeOnly"`
}

// NewFullNameMapper registers a new resource with the given unique name, arguments, and options.
func NewFullNameMapper(ctx *pulumi.Context,
	name string, args *FullNameMapperArgs, opts ...pulumi.ResourceOption) (*FullNameMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LdapFullNameAttribute == nil {
		return nil, errors.New("invalid value for required argument 'LdapFullNameAttribute'")
	}
	if args.LdapUserFederationId == nil {
		return nil, errors.New("invalid value for required argument 'LdapUserFederationId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FullNameMapper
	err := ctx.RegisterResource("keycloak:ldap/fullNameMapper:FullNameMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFullNameMapper gets an existing FullNameMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFullNameMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FullNameMapperState, opts ...pulumi.ResourceOption) (*FullNameMapper, error) {
	var resource FullNameMapper
	err := ctx.ReadResource("keycloak:ldap/fullNameMapper:FullNameMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FullNameMapper resources.
type fullNameMapperState struct {
	// The name of the LDAP attribute containing the user's full name.
	LdapFullNameAttribute *string `pulumi:"ldapFullNameAttribute"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
	ReadOnly *bool `pulumi:"readOnly"`
	// The realm that this LDAP mapper will exist in.
	RealmId *string `pulumi:"realmId"`
	// When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
	WriteOnly *bool `pulumi:"writeOnly"`
}

type FullNameMapperState struct {
	// The name of the LDAP attribute containing the user's full name.
	LdapFullNameAttribute pulumi.StringPtrInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
	ReadOnly pulumi.BoolPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringPtrInput
	// When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
	WriteOnly pulumi.BoolPtrInput
}

func (FullNameMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*fullNameMapperState)(nil)).Elem()
}

type fullNameMapperArgs struct {
	// The name of the LDAP attribute containing the user's full name.
	LdapFullNameAttribute string `pulumi:"ldapFullNameAttribute"`
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Display name of this mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
	ReadOnly *bool `pulumi:"readOnly"`
	// The realm that this LDAP mapper will exist in.
	RealmId string `pulumi:"realmId"`
	// When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
	WriteOnly *bool `pulumi:"writeOnly"`
}

// The set of arguments for constructing a FullNameMapper resource.
type FullNameMapperArgs struct {
	// The name of the LDAP attribute containing the user's full name.
	LdapFullNameAttribute pulumi.StringInput
	// The ID of the LDAP user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Display name of this mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
	ReadOnly pulumi.BoolPtrInput
	// The realm that this LDAP mapper will exist in.
	RealmId pulumi.StringInput
	// When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
	WriteOnly pulumi.BoolPtrInput
}

func (FullNameMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fullNameMapperArgs)(nil)).Elem()
}

type FullNameMapperInput interface {
	pulumi.Input

	ToFullNameMapperOutput() FullNameMapperOutput
	ToFullNameMapperOutputWithContext(ctx context.Context) FullNameMapperOutput
}

func (*FullNameMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**FullNameMapper)(nil)).Elem()
}

func (i *FullNameMapper) ToFullNameMapperOutput() FullNameMapperOutput {
	return i.ToFullNameMapperOutputWithContext(context.Background())
}

func (i *FullNameMapper) ToFullNameMapperOutputWithContext(ctx context.Context) FullNameMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FullNameMapperOutput)
}

func (i *FullNameMapper) ToOutput(ctx context.Context) pulumix.Output[*FullNameMapper] {
	return pulumix.Output[*FullNameMapper]{
		OutputState: i.ToFullNameMapperOutputWithContext(ctx).OutputState,
	}
}

// FullNameMapperArrayInput is an input type that accepts FullNameMapperArray and FullNameMapperArrayOutput values.
// You can construct a concrete instance of `FullNameMapperArrayInput` via:
//
//	FullNameMapperArray{ FullNameMapperArgs{...} }
type FullNameMapperArrayInput interface {
	pulumi.Input

	ToFullNameMapperArrayOutput() FullNameMapperArrayOutput
	ToFullNameMapperArrayOutputWithContext(context.Context) FullNameMapperArrayOutput
}

type FullNameMapperArray []FullNameMapperInput

func (FullNameMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FullNameMapper)(nil)).Elem()
}

func (i FullNameMapperArray) ToFullNameMapperArrayOutput() FullNameMapperArrayOutput {
	return i.ToFullNameMapperArrayOutputWithContext(context.Background())
}

func (i FullNameMapperArray) ToFullNameMapperArrayOutputWithContext(ctx context.Context) FullNameMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FullNameMapperArrayOutput)
}

func (i FullNameMapperArray) ToOutput(ctx context.Context) pulumix.Output[[]*FullNameMapper] {
	return pulumix.Output[[]*FullNameMapper]{
		OutputState: i.ToFullNameMapperArrayOutputWithContext(ctx).OutputState,
	}
}

// FullNameMapperMapInput is an input type that accepts FullNameMapperMap and FullNameMapperMapOutput values.
// You can construct a concrete instance of `FullNameMapperMapInput` via:
//
//	FullNameMapperMap{ "key": FullNameMapperArgs{...} }
type FullNameMapperMapInput interface {
	pulumi.Input

	ToFullNameMapperMapOutput() FullNameMapperMapOutput
	ToFullNameMapperMapOutputWithContext(context.Context) FullNameMapperMapOutput
}

type FullNameMapperMap map[string]FullNameMapperInput

func (FullNameMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FullNameMapper)(nil)).Elem()
}

func (i FullNameMapperMap) ToFullNameMapperMapOutput() FullNameMapperMapOutput {
	return i.ToFullNameMapperMapOutputWithContext(context.Background())
}

func (i FullNameMapperMap) ToFullNameMapperMapOutputWithContext(ctx context.Context) FullNameMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FullNameMapperMapOutput)
}

func (i FullNameMapperMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FullNameMapper] {
	return pulumix.Output[map[string]*FullNameMapper]{
		OutputState: i.ToFullNameMapperMapOutputWithContext(ctx).OutputState,
	}
}

type FullNameMapperOutput struct{ *pulumi.OutputState }

func (FullNameMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FullNameMapper)(nil)).Elem()
}

func (o FullNameMapperOutput) ToFullNameMapperOutput() FullNameMapperOutput {
	return o
}

func (o FullNameMapperOutput) ToFullNameMapperOutputWithContext(ctx context.Context) FullNameMapperOutput {
	return o
}

func (o FullNameMapperOutput) ToOutput(ctx context.Context) pulumix.Output[*FullNameMapper] {
	return pulumix.Output[*FullNameMapper]{
		OutputState: o.OutputState,
	}
}

// The name of the LDAP attribute containing the user's full name.
func (o FullNameMapperOutput) LdapFullNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *FullNameMapper) pulumi.StringOutput { return v.LdapFullNameAttribute }).(pulumi.StringOutput)
}

// The ID of the LDAP user federation provider to attach this mapper to.
func (o FullNameMapperOutput) LdapUserFederationId() pulumi.StringOutput {
	return o.ApplyT(func(v *FullNameMapper) pulumi.StringOutput { return v.LdapUserFederationId }).(pulumi.StringOutput)
}

// Display name of this mapper when displayed in the console.
func (o FullNameMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FullNameMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
func (o FullNameMapperOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FullNameMapper) pulumi.BoolPtrOutput { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

// The realm that this LDAP mapper will exist in.
func (o FullNameMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *FullNameMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
func (o FullNameMapperOutput) WriteOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FullNameMapper) pulumi.BoolPtrOutput { return v.WriteOnly }).(pulumi.BoolPtrOutput)
}

type FullNameMapperArrayOutput struct{ *pulumi.OutputState }

func (FullNameMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FullNameMapper)(nil)).Elem()
}

func (o FullNameMapperArrayOutput) ToFullNameMapperArrayOutput() FullNameMapperArrayOutput {
	return o
}

func (o FullNameMapperArrayOutput) ToFullNameMapperArrayOutputWithContext(ctx context.Context) FullNameMapperArrayOutput {
	return o
}

func (o FullNameMapperArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FullNameMapper] {
	return pulumix.Output[[]*FullNameMapper]{
		OutputState: o.OutputState,
	}
}

func (o FullNameMapperArrayOutput) Index(i pulumi.IntInput) FullNameMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FullNameMapper {
		return vs[0].([]*FullNameMapper)[vs[1].(int)]
	}).(FullNameMapperOutput)
}

type FullNameMapperMapOutput struct{ *pulumi.OutputState }

func (FullNameMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FullNameMapper)(nil)).Elem()
}

func (o FullNameMapperMapOutput) ToFullNameMapperMapOutput() FullNameMapperMapOutput {
	return o
}

func (o FullNameMapperMapOutput) ToFullNameMapperMapOutputWithContext(ctx context.Context) FullNameMapperMapOutput {
	return o
}

func (o FullNameMapperMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FullNameMapper] {
	return pulumix.Output[map[string]*FullNameMapper]{
		OutputState: o.OutputState,
	}
}

func (o FullNameMapperMapOutput) MapIndex(k pulumi.StringInput) FullNameMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FullNameMapper {
		return vs[0].(map[string]*FullNameMapper)[vs[1].(string)]
	}).(FullNameMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FullNameMapperInput)(nil)).Elem(), &FullNameMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*FullNameMapperArrayInput)(nil)).Elem(), FullNameMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FullNameMapperMapInput)(nil)).Elem(), FullNameMapperMap{})
	pulumi.RegisterOutputType(FullNameMapperOutput{})
	pulumi.RegisterOutputType(FullNameMapperArrayOutput{})
	pulumi.RegisterOutputType(FullNameMapperMapOutput{})
}
