// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing Groups within Keycloak.
//
// Groups provide a logical wrapping for users within Keycloak. Users within a group can share attributes and roles, and
// group membership can be mapped to a claim.
//
// Attributes can also be defined on Groups.
//
// Groups can also be federated from external data sources, such as LDAP or Active Directory. This resource **should not**
// be used to manage groups that were created this way.
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
//			parentGroup, err := keycloak.NewGroup(ctx, "parentGroup", &keycloak.GroupArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewGroup(ctx, "childGroup", &keycloak.GroupArgs{
//				RealmId:  realm.ID(),
//				ParentId: parentGroup.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewGroup(ctx, "childGroupWithOptionalAttributes", &keycloak.GroupArgs{
//				RealmId:  realm.ID(),
//				ParentId: parentGroup.ID(),
//				Attributes: pulumi.AnyMap{
//					"foo":        pulumi.Any("bar"),
//					"multivalue": pulumi.Any("value1##value2"),
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
// Groups can be imported using the format `{{realm_id}}/{{group_id}}`, where `group_id` is the unique ID that Keycloak assigns to the group upon creation. This value can be found in the URI when editing this group in the GUI, and is typically a GUID. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:index/group:Group child_group my-realm/934a4a4e-28bd-4703-a0fa-332df153aabd
//
// ```
type Group struct {
	pulumi.CustomResourceState

	// A map representing attributes for the group. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes pulumi.MapOutput `pulumi:"attributes"`
	// The name of the group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of this group's parent. If omitted, this group will be defined at the root level.
	ParentId pulumi.StringPtrOutput `pulumi:"parentId"`
	// (Computed) The complete path of the group. For example, the child group's path in the example configuration would be `/parent-group/child-group`.
	Path pulumi.StringOutput `pulumi:"path"`
	// The realm this group exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource Group
	err := ctx.RegisterResource("keycloak:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("keycloak:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// A map representing attributes for the group. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes map[string]interface{} `pulumi:"attributes"`
	// The name of the group.
	Name *string `pulumi:"name"`
	// The ID of this group's parent. If omitted, this group will be defined at the root level.
	ParentId *string `pulumi:"parentId"`
	// (Computed) The complete path of the group. For example, the child group's path in the example configuration would be `/parent-group/child-group`.
	Path *string `pulumi:"path"`
	// The realm this group exists in.
	RealmId *string `pulumi:"realmId"`
}

type GroupState struct {
	// A map representing attributes for the group. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes pulumi.MapInput
	// The name of the group.
	Name pulumi.StringPtrInput
	// The ID of this group's parent. If omitted, this group will be defined at the root level.
	ParentId pulumi.StringPtrInput
	// (Computed) The complete path of the group. For example, the child group's path in the example configuration would be `/parent-group/child-group`.
	Path pulumi.StringPtrInput
	// The realm this group exists in.
	RealmId pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// A map representing attributes for the group. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes map[string]interface{} `pulumi:"attributes"`
	// The name of the group.
	Name *string `pulumi:"name"`
	// The ID of this group's parent. If omitted, this group will be defined at the root level.
	ParentId *string `pulumi:"parentId"`
	// The realm this group exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// A map representing attributes for the group. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes pulumi.MapInput
	// The name of the group.
	Name pulumi.StringPtrInput
	// The ID of this group's parent. If omitted, this group will be defined at the root level.
	ParentId pulumi.StringPtrInput
	// The realm this group exists in.
	RealmId pulumi.StringInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// A map representing attributes for the group. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
func (o GroupOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v *Group) pulumi.MapOutput { return v.Attributes }).(pulumi.MapOutput)
}

// The name of the group.
func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of this group's parent. If omitted, this group will be defined at the root level.
func (o GroupOutput) ParentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.ParentId }).(pulumi.StringPtrOutput)
}

// (Computed) The complete path of the group. For example, the child group's path in the example configuration would be `/parent-group/child-group`.
func (o GroupOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The realm this group exists in.
func (o GroupOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
