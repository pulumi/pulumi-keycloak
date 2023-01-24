// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch properties of a user within Keycloak.
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
//			masterRealm, err := keycloak.LookupRealm(ctx, &keycloak.LookupRealmArgs{
//				Realm: "master",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultAdminUser, err := keycloak.LookupUser(ctx, &keycloak.LookupUserArgs{
//				RealmId:  masterRealm.Id,
//				Username: "keycloak",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("keycloakUserId", defaultAdminUser.Id)
//			return nil
//		})
//	}
//
// ```
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("keycloak:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// The realm this user belongs to.
	RealmId string `pulumi:"realmId"`
	// The unique username of this user.
	Username string `pulumi:"username"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// (Computed) A map representing attributes for the user
	Attributes map[string]interface{} `pulumi:"attributes"`
	// (Computed) The user's email.
	Email string `pulumi:"email"`
	// (Computed) Whether the email address was validated or not. Default to `false`.
	EmailVerified bool `pulumi:"emailVerified"`
	// (Computed) When false, this user cannot log in. Defaults to `true`.
	Enabled bool `pulumi:"enabled"`
	// (Computed) The user's federated identities, if applicable. This block has the following schema:
	FederatedIdentities []string `pulumi:"federatedIdentities"`
	// (Computed) The user's first name.
	FirstName string `pulumi:"firstName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) The user's last name.
	LastName string `pulumi:"lastName"`
	RealmId  string `pulumi:"realmId"`
	Username string `pulumi:"username"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// The realm this user belongs to.
	RealmId pulumi.StringInput `pulumi:"realmId"`
	// The unique username of this user.
	Username pulumi.StringInput `pulumi:"username"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// (Computed) A map representing attributes for the user
func (o LookupUserResultOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v LookupUserResult) map[string]interface{} { return v.Attributes }).(pulumi.MapOutput)
}

// (Computed) The user's email.
func (o LookupUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// (Computed) Whether the email address was validated or not. Default to `false`.
func (o LookupUserResultOutput) EmailVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUserResult) bool { return v.EmailVerified }).(pulumi.BoolOutput)
}

// (Computed) When false, this user cannot log in. Defaults to `true`.
func (o LookupUserResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUserResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// (Computed) The user's federated identities, if applicable. This block has the following schema:
func (o LookupUserResultOutput) FederatedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.FederatedIdentities }).(pulumi.StringArrayOutput)
}

// (Computed) The user's first name.
func (o LookupUserResultOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.FirstName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Computed) The user's last name.
func (o LookupUserResultOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.LastName }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.RealmId }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
