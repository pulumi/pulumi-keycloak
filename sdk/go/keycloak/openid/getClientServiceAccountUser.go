// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch information about the service account user that is associated with an OpenID client
// that has service accounts enabled.
//
// ## Example Usage
//
// In this example, we'll create an OpenID client with service accounts enabled. This causes Keycloak to create a special user
// that represents the service account. We'll use this data source to grab this user's ID in order to assign some roles to this
// user, using the `UserRoles` resource.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/openid"
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
//			client, err := openid.NewClient(ctx, "client", &openid.ClientArgs{
//				RealmId:                realm.ID(),
//				ClientId:               pulumi.String("client"),
//				AccessType:             pulumi.String("CONFIDENTIAL"),
//				ServiceAccountsEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			serviceAccountUser := openid.GetClientServiceAccountUserOutput(ctx, openid.GetClientServiceAccountUserOutputArgs{
//				RealmId:  realm.ID(),
//				ClientId: client.ID(),
//			}, nil)
//			offlineAccess := keycloak.LookupRoleOutput(ctx, keycloak.GetRoleOutputArgs{
//				RealmId: realm.ID(),
//				Name:    pulumi.String("offline_access"),
//			}, nil)
//			_, err = keycloak.NewUserRoles(ctx, "serviceAccountUserRoles", &keycloak.UserRolesArgs{
//				RealmId: realm.ID(),
//				UserId: serviceAccountUser.ApplyT(func(serviceAccountUser openid.GetClientServiceAccountUserResult) (*string, error) {
//					return &serviceAccountUser.Id, nil
//				}).(pulumi.StringPtrOutput),
//				RoleIds: pulumi.StringArray{
//					offlineAccess.ApplyT(func(offlineAccess keycloak.GetRoleResult) (*string, error) {
//						return &offlineAccess.Id, nil
//					}).(pulumi.StringPtrOutput),
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
func GetClientServiceAccountUser(ctx *pulumi.Context, args *GetClientServiceAccountUserArgs, opts ...pulumi.InvokeOption) (*GetClientServiceAccountUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetClientServiceAccountUserResult
	err := ctx.Invoke("keycloak:openid/getClientServiceAccountUser:getClientServiceAccountUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClientServiceAccountUser.
type GetClientServiceAccountUserArgs struct {
	// The ID of the OpenID client with service accounts enabled.
	ClientId string `pulumi:"clientId"`
	// The realm that the OpenID client exists within.
	RealmId string `pulumi:"realmId"`
}

// A collection of values returned by getClientServiceAccountUser.
type GetClientServiceAccountUserResult struct {
	Attributes          map[string]interface{}                         `pulumi:"attributes"`
	ClientId            string                                         `pulumi:"clientId"`
	Email               string                                         `pulumi:"email"`
	EmailVerified       bool                                           `pulumi:"emailVerified"`
	Enabled             bool                                           `pulumi:"enabled"`
	FederatedIdentities []GetClientServiceAccountUserFederatedIdentity `pulumi:"federatedIdentities"`
	FirstName           string                                         `pulumi:"firstName"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	LastName string `pulumi:"lastName"`
	RealmId  string `pulumi:"realmId"`
	Username string `pulumi:"username"`
}

func GetClientServiceAccountUserOutput(ctx *pulumi.Context, args GetClientServiceAccountUserOutputArgs, opts ...pulumi.InvokeOption) GetClientServiceAccountUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClientServiceAccountUserResult, error) {
			args := v.(GetClientServiceAccountUserArgs)
			r, err := GetClientServiceAccountUser(ctx, &args, opts...)
			var s GetClientServiceAccountUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClientServiceAccountUserResultOutput)
}

// A collection of arguments for invoking getClientServiceAccountUser.
type GetClientServiceAccountUserOutputArgs struct {
	// The ID of the OpenID client with service accounts enabled.
	ClientId pulumi.StringInput `pulumi:"clientId"`
	// The realm that the OpenID client exists within.
	RealmId pulumi.StringInput `pulumi:"realmId"`
}

func (GetClientServiceAccountUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientServiceAccountUserArgs)(nil)).Elem()
}

// A collection of values returned by getClientServiceAccountUser.
type GetClientServiceAccountUserResultOutput struct{ *pulumi.OutputState }

func (GetClientServiceAccountUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientServiceAccountUserResult)(nil)).Elem()
}

func (o GetClientServiceAccountUserResultOutput) ToGetClientServiceAccountUserResultOutput() GetClientServiceAccountUserResultOutput {
	return o
}

func (o GetClientServiceAccountUserResultOutput) ToGetClientServiceAccountUserResultOutputWithContext(ctx context.Context) GetClientServiceAccountUserResultOutput {
	return o
}

func (o GetClientServiceAccountUserResultOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) map[string]interface{} { return v.Attributes }).(pulumi.MapOutput)
}

func (o GetClientServiceAccountUserResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o GetClientServiceAccountUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetClientServiceAccountUserResultOutput) EmailVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) bool { return v.EmailVerified }).(pulumi.BoolOutput)
}

func (o GetClientServiceAccountUserResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetClientServiceAccountUserResultOutput) FederatedIdentities() GetClientServiceAccountUserFederatedIdentityArrayOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) []GetClientServiceAccountUserFederatedIdentity {
		return v.FederatedIdentities
	}).(GetClientServiceAccountUserFederatedIdentityArrayOutput)
}

func (o GetClientServiceAccountUserResultOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) string { return v.FirstName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClientServiceAccountUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetClientServiceAccountUserResultOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) string { return v.LastName }).(pulumi.StringOutput)
}

func (o GetClientServiceAccountUserResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) string { return v.RealmId }).(pulumi.StringOutput)
}

func (o GetClientServiceAccountUserResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClientServiceAccountUserResultOutput{})
}
