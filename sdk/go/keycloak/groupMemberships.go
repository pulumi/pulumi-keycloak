// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for managing a Keycloak group's members.
//
// Note that this resource attempts to be an **authoritative** source over group members. When this resource takes control
// over a group's members, users that are manually added to the group will be removed, and users that are manually removed
// from the group will be added upon the next run of `pulumi up`.
//
// Also note that you should not use `GroupMemberships` with a group has been assigned as a default group via
// `DefaultGroups`.
//
// This resource **should not** be used to control membership of a group that has its members federated from an external
// source via group mapping.
//
// To non-exclusively manage the group's of a user, see the [`UserGroups` resource][1]
//
// This resource paginates its data loading on refresh by 50 items.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
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
//			group, err := keycloak.NewGroup(ctx, "group", &keycloak.GroupArgs{
//				RealmId: realm.ID(),
//				Name:    pulumi.String("my-group"),
//			})
//			if err != nil {
//				return err
//			}
//			user, err := keycloak.NewUser(ctx, "user", &keycloak.UserArgs{
//				RealmId:  realm.ID(),
//				Username: pulumi.String("my-user"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewGroupMemberships(ctx, "group_members", &keycloak.GroupMembershipsArgs{
//				RealmId: realm.ID(),
//				GroupId: group.ID(),
//				Members: pulumi.StringArray{
//					user.Username,
//				},
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
// This resource does not support import. Instead of importing, feel free to create this resource
//
// as if it did not already exist on the server.
//
// [1]: providers/keycloak/keycloak/latest/docs/resources/group_memberships
type GroupMemberships struct {
	pulumi.CustomResourceState

	// The ID of the group this resource should manage memberships for.
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// A list of usernames that belong to this group.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The realm this group exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewGroupMemberships registers a new resource with the given unique name, arguments, and options.
func NewGroupMemberships(ctx *pulumi.Context,
	name string, args *GroupMembershipsArgs, opts ...pulumi.ResourceOption) (*GroupMemberships, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupMemberships
	err := ctx.RegisterResource("keycloak:index/groupMemberships:GroupMemberships", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMemberships gets an existing GroupMemberships resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMemberships(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipsState, opts ...pulumi.ResourceOption) (*GroupMemberships, error) {
	var resource GroupMemberships
	err := ctx.ReadResource("keycloak:index/groupMemberships:GroupMemberships", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMemberships resources.
type groupMembershipsState struct {
	// The ID of the group this resource should manage memberships for.
	GroupId *string `pulumi:"groupId"`
	// A list of usernames that belong to this group.
	Members []string `pulumi:"members"`
	// The realm this group exists in.
	RealmId *string `pulumi:"realmId"`
}

type GroupMembershipsState struct {
	// The ID of the group this resource should manage memberships for.
	GroupId pulumi.StringPtrInput
	// A list of usernames that belong to this group.
	Members pulumi.StringArrayInput
	// The realm this group exists in.
	RealmId pulumi.StringPtrInput
}

func (GroupMembershipsState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipsState)(nil)).Elem()
}

type groupMembershipsArgs struct {
	// The ID of the group this resource should manage memberships for.
	GroupId *string `pulumi:"groupId"`
	// A list of usernames that belong to this group.
	Members []string `pulumi:"members"`
	// The realm this group exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a GroupMemberships resource.
type GroupMembershipsArgs struct {
	// The ID of the group this resource should manage memberships for.
	GroupId pulumi.StringPtrInput
	// A list of usernames that belong to this group.
	Members pulumi.StringArrayInput
	// The realm this group exists in.
	RealmId pulumi.StringInput
}

func (GroupMembershipsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipsArgs)(nil)).Elem()
}

type GroupMembershipsInput interface {
	pulumi.Input

	ToGroupMembershipsOutput() GroupMembershipsOutput
	ToGroupMembershipsOutputWithContext(ctx context.Context) GroupMembershipsOutput
}

func (*GroupMemberships) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMemberships)(nil)).Elem()
}

func (i *GroupMemberships) ToGroupMembershipsOutput() GroupMembershipsOutput {
	return i.ToGroupMembershipsOutputWithContext(context.Background())
}

func (i *GroupMemberships) ToGroupMembershipsOutputWithContext(ctx context.Context) GroupMembershipsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipsOutput)
}

// GroupMembershipsArrayInput is an input type that accepts GroupMembershipsArray and GroupMembershipsArrayOutput values.
// You can construct a concrete instance of `GroupMembershipsArrayInput` via:
//
//	GroupMembershipsArray{ GroupMembershipsArgs{...} }
type GroupMembershipsArrayInput interface {
	pulumi.Input

	ToGroupMembershipsArrayOutput() GroupMembershipsArrayOutput
	ToGroupMembershipsArrayOutputWithContext(context.Context) GroupMembershipsArrayOutput
}

type GroupMembershipsArray []GroupMembershipsInput

func (GroupMembershipsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMemberships)(nil)).Elem()
}

func (i GroupMembershipsArray) ToGroupMembershipsArrayOutput() GroupMembershipsArrayOutput {
	return i.ToGroupMembershipsArrayOutputWithContext(context.Background())
}

func (i GroupMembershipsArray) ToGroupMembershipsArrayOutputWithContext(ctx context.Context) GroupMembershipsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipsArrayOutput)
}

// GroupMembershipsMapInput is an input type that accepts GroupMembershipsMap and GroupMembershipsMapOutput values.
// You can construct a concrete instance of `GroupMembershipsMapInput` via:
//
//	GroupMembershipsMap{ "key": GroupMembershipsArgs{...} }
type GroupMembershipsMapInput interface {
	pulumi.Input

	ToGroupMembershipsMapOutput() GroupMembershipsMapOutput
	ToGroupMembershipsMapOutputWithContext(context.Context) GroupMembershipsMapOutput
}

type GroupMembershipsMap map[string]GroupMembershipsInput

func (GroupMembershipsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMemberships)(nil)).Elem()
}

func (i GroupMembershipsMap) ToGroupMembershipsMapOutput() GroupMembershipsMapOutput {
	return i.ToGroupMembershipsMapOutputWithContext(context.Background())
}

func (i GroupMembershipsMap) ToGroupMembershipsMapOutputWithContext(ctx context.Context) GroupMembershipsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipsMapOutput)
}

type GroupMembershipsOutput struct{ *pulumi.OutputState }

func (GroupMembershipsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMemberships)(nil)).Elem()
}

func (o GroupMembershipsOutput) ToGroupMembershipsOutput() GroupMembershipsOutput {
	return o
}

func (o GroupMembershipsOutput) ToGroupMembershipsOutputWithContext(ctx context.Context) GroupMembershipsOutput {
	return o
}

// The ID of the group this resource should manage memberships for.
func (o GroupMembershipsOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMemberships) pulumi.StringPtrOutput { return v.GroupId }).(pulumi.StringPtrOutput)
}

// A list of usernames that belong to this group.
func (o GroupMembershipsOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupMemberships) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The realm this group exists in.
func (o GroupMembershipsOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMemberships) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type GroupMembershipsArrayOutput struct{ *pulumi.OutputState }

func (GroupMembershipsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMemberships)(nil)).Elem()
}

func (o GroupMembershipsArrayOutput) ToGroupMembershipsArrayOutput() GroupMembershipsArrayOutput {
	return o
}

func (o GroupMembershipsArrayOutput) ToGroupMembershipsArrayOutputWithContext(ctx context.Context) GroupMembershipsArrayOutput {
	return o
}

func (o GroupMembershipsArrayOutput) Index(i pulumi.IntInput) GroupMembershipsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMemberships {
		return vs[0].([]*GroupMemberships)[vs[1].(int)]
	}).(GroupMembershipsOutput)
}

type GroupMembershipsMapOutput struct{ *pulumi.OutputState }

func (GroupMembershipsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMemberships)(nil)).Elem()
}

func (o GroupMembershipsMapOutput) ToGroupMembershipsMapOutput() GroupMembershipsMapOutput {
	return o
}

func (o GroupMembershipsMapOutput) ToGroupMembershipsMapOutputWithContext(ctx context.Context) GroupMembershipsMapOutput {
	return o
}

func (o GroupMembershipsMapOutput) MapIndex(k pulumi.StringInput) GroupMembershipsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMemberships {
		return vs[0].(map[string]*GroupMemberships)[vs[1].(string)]
	}).(GroupMembershipsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipsInput)(nil)).Elem(), &GroupMemberships{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipsArrayInput)(nil)).Elem(), GroupMembershipsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipsMapInput)(nil)).Elem(), GroupMembershipsMap{})
	pulumi.RegisterOutputType(GroupMembershipsOutput{})
	pulumi.RegisterOutputType(GroupMembershipsArrayOutput{})
	pulumi.RegisterOutputType(GroupMembershipsMapOutput{})
}
