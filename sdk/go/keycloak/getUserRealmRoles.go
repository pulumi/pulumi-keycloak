// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch the realm roles of a user within Keycloak.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			masterRealm, err := keycloak.LookupRealm(ctx, &keycloak.LookupRealmArgs{
//				Realm: "master",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// use the keycloak_user data source to grab the admin user's ID
//			defaultAdminUser, err := keycloak.LookupUser(ctx, &keycloak.LookupUserArgs{
//				RealmId:  masterRealm.Id,
//				Username: "keycloak",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// use the keycloak_user_realm_roles data source to list role names
//			userRealmRoles, err := keycloak.GetUserRealmRoles(ctx, &keycloak.GetUserRealmRolesArgs{
//				RealmId: masterRealm.Id,
//				UserId:  defaultAdminUser.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("keycloakUserRoleNames", userRealmRoles.RoleNames)
//			return nil
//		})
//	}
//
// ```
func GetUserRealmRoles(ctx *pulumi.Context, args *GetUserRealmRolesArgs, opts ...pulumi.InvokeOption) (*GetUserRealmRolesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserRealmRolesResult
	err := ctx.Invoke("keycloak:index/getUserRealmRoles:getUserRealmRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserRealmRoles.
type GetUserRealmRolesArgs struct {
	// The realm this user belongs to.
	RealmId string `pulumi:"realmId"`
	// The ID of the user to query realm roles for.
	UserId string `pulumi:"userId"`
}

// A collection of values returned by getUserRealmRoles.
type GetUserRealmRolesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	RealmId string `pulumi:"realmId"`
	// (Computed) A list of realm roles that belong to this user.
	RoleNames []string `pulumi:"roleNames"`
	UserId    string   `pulumi:"userId"`
}

func GetUserRealmRolesOutput(ctx *pulumi.Context, args GetUserRealmRolesOutputArgs, opts ...pulumi.InvokeOption) GetUserRealmRolesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetUserRealmRolesResultOutput, error) {
			args := v.(GetUserRealmRolesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("keycloak:index/getUserRealmRoles:getUserRealmRoles", args, GetUserRealmRolesResultOutput{}, options).(GetUserRealmRolesResultOutput), nil
		}).(GetUserRealmRolesResultOutput)
}

// A collection of arguments for invoking getUserRealmRoles.
type GetUserRealmRolesOutputArgs struct {
	// The realm this user belongs to.
	RealmId pulumi.StringInput `pulumi:"realmId"`
	// The ID of the user to query realm roles for.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (GetUserRealmRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserRealmRolesArgs)(nil)).Elem()
}

// A collection of values returned by getUserRealmRoles.
type GetUserRealmRolesResultOutput struct{ *pulumi.OutputState }

func (GetUserRealmRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserRealmRolesResult)(nil)).Elem()
}

func (o GetUserRealmRolesResultOutput) ToGetUserRealmRolesResultOutput() GetUserRealmRolesResultOutput {
	return o
}

func (o GetUserRealmRolesResultOutput) ToGetUserRealmRolesResultOutputWithContext(ctx context.Context) GetUserRealmRolesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserRealmRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserRealmRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUserRealmRolesResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserRealmRolesResult) string { return v.RealmId }).(pulumi.StringOutput)
}

// (Computed) A list of realm roles that belong to this user.
func (o GetUserRealmRolesResultOutput) RoleNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUserRealmRolesResult) []string { return v.RoleNames }).(pulumi.StringArrayOutput)
}

func (o GetUserRealmRolesResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserRealmRolesResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserRealmRolesResultOutput{})
}
