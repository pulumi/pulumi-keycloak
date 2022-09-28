// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch properties of a Keycloak group for
// usage with other resources, such as `GroupRoles`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
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
//			offlineAccess := keycloak.LookupRoleOutput(ctx, GetRoleOutputArgs{
//				RealmId: realm.ID(),
//				Name:    pulumi.String("offline_access"),
//			}, nil)
//			group := keycloak.LookupGroupOutput(ctx, GetGroupOutputArgs{
//				RealmId: realm.ID(),
//				Name:    pulumi.String("group"),
//			}, nil)
//			_, err = keycloak.NewGroupRoles(ctx, "groupRoles", &keycloak.GroupRolesArgs{
//				RealmId: realm.ID(),
//				GroupId: group.ApplyT(func(group GetGroupResult) (string, error) {
//					return group.Id, nil
//				}).(pulumi.StringOutput),
//				RoleIds: pulumi.StringArray{
//					offlineAccess.ApplyT(func(offlineAccess GetRoleResult) (string, error) {
//						return offlineAccess.Id, nil
//					}).(pulumi.StringOutput),
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
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("keycloak:index/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// The name of the group. If there are multiple groups match `name`, the first result will be returned.
	Name string `pulumi:"name"`
	// The realm this group exists within.
	RealmId string `pulumi:"realmId"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	Attributes map[string]interface{} `pulumi:"attributes"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	ParentId string `pulumi:"parentId"`
	Path     string `pulumi:"path"`
	RealmId  string `pulumi:"realmId"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupOutputArgs struct {
	// The name of the group. If there are multiple groups match `name`, the first result will be returned.
	Name pulumi.StringInput `pulumi:"name"`
	// The realm this group exists within.
	RealmId pulumi.StringInput `pulumi:"realmId"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v LookupGroupResult) map[string]interface{} { return v.Attributes }).(pulumi.MapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.ParentId }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.RealmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
