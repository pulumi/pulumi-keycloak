// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
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
// 	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("my-realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewGroupRoles(ctx, "groupRoles", &keycloak.GroupRolesArgs{
// 			RealmId: realm.ID(),
// 			GroupId: group.ApplyT(func(group keycloak.LookupGroupResult) (string, error) {
// 				return group.Id, nil
// 			}).(pulumi.StringOutput),
// 			RoleIds: pulumi.StringArray{
// 				offlineAccess.ApplyT(func(offlineAccess keycloak.LookupRoleResult) (string, error) {
// 					return offlineAccess.Id, nil
// 				}).(pulumi.StringOutput),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
