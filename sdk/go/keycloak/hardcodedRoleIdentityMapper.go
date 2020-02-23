// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type HardcodedRoleIdentityMapper struct {
	pulumi.CustomResourceState

	// IDP Alias
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Realm Name
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Role Name
	Role pulumi.StringPtrOutput `pulumi:"role"`
}

// NewHardcodedRoleIdentityMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedRoleIdentityMapper(ctx *pulumi.Context,
	name string, args *HardcodedRoleIdentityMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedRoleIdentityMapper, error) {
	if args == nil || args.IdentityProviderAlias == nil {
		return nil, errors.New("missing required argument 'IdentityProviderAlias'")
	}
	if args == nil || args.Realm == nil {
		return nil, errors.New("missing required argument 'Realm'")
	}
	if args == nil {
		args = &HardcodedRoleIdentityMapperArgs{}
	}
	var resource HardcodedRoleIdentityMapper
	err := ctx.RegisterResource("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedRoleIdentityMapper gets an existing HardcodedRoleIdentityMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedRoleIdentityMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedRoleIdentityMapperState, opts ...pulumi.ResourceOption) (*HardcodedRoleIdentityMapper, error) {
	var resource HardcodedRoleIdentityMapper
	err := ctx.ReadResource("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedRoleIdentityMapper resources.
type hardcodedRoleIdentityMapperState struct {
	// IDP Alias
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm *string `pulumi:"realm"`
	// Role Name
	Role *string `pulumi:"role"`
}

type HardcodedRoleIdentityMapperState struct {
	// IDP Alias
	IdentityProviderAlias pulumi.StringPtrInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringPtrInput
	// Role Name
	Role pulumi.StringPtrInput
}

func (HardcodedRoleIdentityMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleIdentityMapperState)(nil)).Elem()
}

type hardcodedRoleIdentityMapperArgs struct {
	// IDP Alias
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm string `pulumi:"realm"`
	// Role Name
	Role *string `pulumi:"role"`
}

// The set of arguments for constructing a HardcodedRoleIdentityMapper resource.
type HardcodedRoleIdentityMapperArgs struct {
	// IDP Alias
	IdentityProviderAlias pulumi.StringInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringInput
	// Role Name
	Role pulumi.StringPtrInput
}

func (HardcodedRoleIdentityMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedRoleIdentityMapperArgs)(nil)).Elem()
}

