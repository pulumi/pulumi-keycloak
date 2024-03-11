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

// Allows for managing a Keycloak user's groups.
//
// If `exhaustive` is true, this resource attempts to be an **authoritative** source over user groups: groups that are manually added to the user will be removed, and groups that are manually removed from the user group will be added upon the next run of `pulumi up`.
// If `exhaustive` is false, this resource is a partial assignation of groups to a user. As a result, you can get multiple `UserGroups` for the same `userId`.
//
// ## Example Usage
//
// ### Exhaustive Groups)
// <!--Start PulumiCodeChooser -->
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
//			_, err = keycloak.NewUserGroups(ctx, "userGroups", &keycloak.UserGroupsArgs{
//				RealmId: realm.ID(),
//				UserId:  user.ID(),
//				GroupIds: pulumi.StringArray{
//					group.ID(),
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
// <!--End PulumiCodeChooser -->
//
// ### Non Exhaustive Groups)
// <!--Start PulumiCodeChooser -->
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
//			groupFoo, err := keycloak.NewGroup(ctx, "groupFoo", &keycloak.GroupArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			groupBar, err := keycloak.NewGroup(ctx, "groupBar", &keycloak.GroupArgs{
//				RealmId: realm.ID(),
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
//			_, err = keycloak.NewUserGroups(ctx, "userGroupsAssociation1UserGroups", &keycloak.UserGroupsArgs{
//				RealmId:    realm.ID(),
//				UserId:     user.ID(),
//				Exhaustive: pulumi.Bool(false),
//				GroupIds: pulumi.StringArray{
//					groupFoo.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewUserGroups(ctx, "userGroupsAssociation1Index/userGroupsUserGroups", &keycloak.UserGroupsArgs{
//				RealmId:    realm.ID(),
//				UserId:     user.ID(),
//				Exhaustive: pulumi.Bool(false),
//				GroupIds: pulumi.StringArray{
//					groupBar.ID(),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// This resource does not support import. Instead of importing, feel free to create this resource
//
// as if it did not already exist on the server.
type UserGroups struct {
	pulumi.CustomResourceState

	// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive pulumi.BoolPtrOutput `pulumi:"exhaustive"`
	// A list of group IDs that the user is member of.
	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	// The realm this group exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The ID of the user this resource should manage groups for.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserGroups registers a new resource with the given unique name, arguments, and options.
func NewUserGroups(ctx *pulumi.Context,
	name string, args *UserGroupsArgs, opts ...pulumi.ResourceOption) (*UserGroups, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupIds == nil {
		return nil, errors.New("invalid value for required argument 'GroupIds'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserGroups
	err := ctx.RegisterResource("keycloak:index/userGroups:UserGroups", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserGroups gets an existing UserGroups resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserGroups(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserGroupsState, opts ...pulumi.ResourceOption) (*UserGroups, error) {
	var resource UserGroups
	err := ctx.ReadResource("keycloak:index/userGroups:UserGroups", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserGroups resources.
type userGroupsState struct {
	// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive *bool `pulumi:"exhaustive"`
	// A list of group IDs that the user is member of.
	GroupIds []string `pulumi:"groupIds"`
	// The realm this group exists in.
	RealmId *string `pulumi:"realmId"`
	// The ID of the user this resource should manage groups for.
	UserId *string `pulumi:"userId"`
}

type UserGroupsState struct {
	// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive pulumi.BoolPtrInput
	// A list of group IDs that the user is member of.
	GroupIds pulumi.StringArrayInput
	// The realm this group exists in.
	RealmId pulumi.StringPtrInput
	// The ID of the user this resource should manage groups for.
	UserId pulumi.StringPtrInput
}

func (UserGroupsState) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupsState)(nil)).Elem()
}

type userGroupsArgs struct {
	// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive *bool `pulumi:"exhaustive"`
	// A list of group IDs that the user is member of.
	GroupIds []string `pulumi:"groupIds"`
	// The realm this group exists in.
	RealmId string `pulumi:"realmId"`
	// The ID of the user this resource should manage groups for.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a UserGroups resource.
type UserGroupsArgs struct {
	// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
	Exhaustive pulumi.BoolPtrInput
	// A list of group IDs that the user is member of.
	GroupIds pulumi.StringArrayInput
	// The realm this group exists in.
	RealmId pulumi.StringInput
	// The ID of the user this resource should manage groups for.
	UserId pulumi.StringInput
}

func (UserGroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupsArgs)(nil)).Elem()
}

type UserGroupsInput interface {
	pulumi.Input

	ToUserGroupsOutput() UserGroupsOutput
	ToUserGroupsOutputWithContext(ctx context.Context) UserGroupsOutput
}

func (*UserGroups) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGroups)(nil)).Elem()
}

func (i *UserGroups) ToUserGroupsOutput() UserGroupsOutput {
	return i.ToUserGroupsOutputWithContext(context.Background())
}

func (i *UserGroups) ToUserGroupsOutputWithContext(ctx context.Context) UserGroupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupsOutput)
}

// UserGroupsArrayInput is an input type that accepts UserGroupsArray and UserGroupsArrayOutput values.
// You can construct a concrete instance of `UserGroupsArrayInput` via:
//
//	UserGroupsArray{ UserGroupsArgs{...} }
type UserGroupsArrayInput interface {
	pulumi.Input

	ToUserGroupsArrayOutput() UserGroupsArrayOutput
	ToUserGroupsArrayOutputWithContext(context.Context) UserGroupsArrayOutput
}

type UserGroupsArray []UserGroupsInput

func (UserGroupsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGroups)(nil)).Elem()
}

func (i UserGroupsArray) ToUserGroupsArrayOutput() UserGroupsArrayOutput {
	return i.ToUserGroupsArrayOutputWithContext(context.Background())
}

func (i UserGroupsArray) ToUserGroupsArrayOutputWithContext(ctx context.Context) UserGroupsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupsArrayOutput)
}

// UserGroupsMapInput is an input type that accepts UserGroupsMap and UserGroupsMapOutput values.
// You can construct a concrete instance of `UserGroupsMapInput` via:
//
//	UserGroupsMap{ "key": UserGroupsArgs{...} }
type UserGroupsMapInput interface {
	pulumi.Input

	ToUserGroupsMapOutput() UserGroupsMapOutput
	ToUserGroupsMapOutputWithContext(context.Context) UserGroupsMapOutput
}

type UserGroupsMap map[string]UserGroupsInput

func (UserGroupsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGroups)(nil)).Elem()
}

func (i UserGroupsMap) ToUserGroupsMapOutput() UserGroupsMapOutput {
	return i.ToUserGroupsMapOutputWithContext(context.Background())
}

func (i UserGroupsMap) ToUserGroupsMapOutputWithContext(ctx context.Context) UserGroupsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupsMapOutput)
}

type UserGroupsOutput struct{ *pulumi.OutputState }

func (UserGroupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGroups)(nil)).Elem()
}

func (o UserGroupsOutput) ToUserGroupsOutput() UserGroupsOutput {
	return o
}

func (o UserGroupsOutput) ToUserGroupsOutputWithContext(ctx context.Context) UserGroupsOutput {
	return o
}

// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
func (o UserGroupsOutput) Exhaustive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserGroups) pulumi.BoolPtrOutput { return v.Exhaustive }).(pulumi.BoolPtrOutput)
}

// A list of group IDs that the user is member of.
func (o UserGroupsOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserGroups) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// The realm this group exists in.
func (o UserGroupsOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGroups) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The ID of the user this resource should manage groups for.
func (o UserGroupsOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGroups) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type UserGroupsArrayOutput struct{ *pulumi.OutputState }

func (UserGroupsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGroups)(nil)).Elem()
}

func (o UserGroupsArrayOutput) ToUserGroupsArrayOutput() UserGroupsArrayOutput {
	return o
}

func (o UserGroupsArrayOutput) ToUserGroupsArrayOutputWithContext(ctx context.Context) UserGroupsArrayOutput {
	return o
}

func (o UserGroupsArrayOutput) Index(i pulumi.IntInput) UserGroupsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserGroups {
		return vs[0].([]*UserGroups)[vs[1].(int)]
	}).(UserGroupsOutput)
}

type UserGroupsMapOutput struct{ *pulumi.OutputState }

func (UserGroupsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGroups)(nil)).Elem()
}

func (o UserGroupsMapOutput) ToUserGroupsMapOutput() UserGroupsMapOutput {
	return o
}

func (o UserGroupsMapOutput) ToUserGroupsMapOutputWithContext(ctx context.Context) UserGroupsMapOutput {
	return o
}

func (o UserGroupsMapOutput) MapIndex(k pulumi.StringInput) UserGroupsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserGroups {
		return vs[0].(map[string]*UserGroups)[vs[1].(string)]
	}).(UserGroupsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserGroupsInput)(nil)).Elem(), &UserGroups{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGroupsArrayInput)(nil)).Elem(), UserGroupsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGroupsMapInput)(nil)).Elem(), UserGroupsMap{})
	pulumi.RegisterOutputType(UserGroupsOutput{})
	pulumi.RegisterOutputType(UserGroupsArrayOutput{})
	pulumi.RegisterOutputType(UserGroupsMapOutput{})
}
