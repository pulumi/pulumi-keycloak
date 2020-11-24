// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// This resource can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak assigns to the user upon creation. This value can be found in the GUI when editing the user, and is typically a GUID. Examplebash
//
// ```sh
//  $ pulumi import keycloak:index/userRoles:UserRoles user_roles my-realm/b0ae6924-1bd5-4655-9e38-dae7c5e42924
// ```
type UserRoles struct {
	pulumi.CustomResourceState

	// The realm this user exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// A list of role IDs to map to the user
	RoleIds pulumi.StringArrayOutput `pulumi:"roleIds"`
	// The ID of the user this resource should manage roles for.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserRoles registers a new resource with the given unique name, arguments, and options.
func NewUserRoles(ctx *pulumi.Context,
	name string, args *UserRolesArgs, opts ...pulumi.ResourceOption) (*UserRoles, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.RoleIds == nil {
		return nil, errors.New("missing required argument 'RoleIds'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	if args == nil {
		args = &UserRolesArgs{}
	}
	var resource UserRoles
	err := ctx.RegisterResource("keycloak:index/userRoles:UserRoles", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserRoles gets an existing UserRoles resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRoles(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRolesState, opts ...pulumi.ResourceOption) (*UserRoles, error) {
	var resource UserRoles
	err := ctx.ReadResource("keycloak:index/userRoles:UserRoles", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserRoles resources.
type userRolesState struct {
	// The realm this user exists in.
	RealmId *string `pulumi:"realmId"`
	// A list of role IDs to map to the user
	RoleIds []string `pulumi:"roleIds"`
	// The ID of the user this resource should manage roles for.
	UserId *string `pulumi:"userId"`
}

type UserRolesState struct {
	// The realm this user exists in.
	RealmId pulumi.StringPtrInput
	// A list of role IDs to map to the user
	RoleIds pulumi.StringArrayInput
	// The ID of the user this resource should manage roles for.
	UserId pulumi.StringPtrInput
}

func (UserRolesState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRolesState)(nil)).Elem()
}

type userRolesArgs struct {
	// The realm this user exists in.
	RealmId string `pulumi:"realmId"`
	// A list of role IDs to map to the user
	RoleIds []string `pulumi:"roleIds"`
	// The ID of the user this resource should manage roles for.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a UserRoles resource.
type UserRolesArgs struct {
	// The realm this user exists in.
	RealmId pulumi.StringInput
	// A list of role IDs to map to the user
	RoleIds pulumi.StringArrayInput
	// The ID of the user this resource should manage roles for.
	UserId pulumi.StringInput
}

func (UserRolesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRolesArgs)(nil)).Elem()
}

type UserRolesInput interface {
	pulumi.Input

	ToUserRolesOutput() UserRolesOutput
	ToUserRolesOutputWithContext(ctx context.Context) UserRolesOutput
}

func (UserRoles) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRoles)(nil)).Elem()
}

func (i UserRoles) ToUserRolesOutput() UserRolesOutput {
	return i.ToUserRolesOutputWithContext(context.Background())
}

func (i UserRoles) ToUserRolesOutputWithContext(ctx context.Context) UserRolesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRolesOutput)
}

type UserRolesOutput struct {
	*pulumi.OutputState
}

func (UserRolesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRolesOutput)(nil)).Elem()
}

func (o UserRolesOutput) ToUserRolesOutput() UserRolesOutput {
	return o
}

func (o UserRolesOutput) ToUserRolesOutputWithContext(ctx context.Context) UserRolesOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserRolesOutput{})
}
