// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # openid.Client data source
//
// This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
//
// ### Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak/openid"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realmManagement, err := openid.LookupClient(ctx, &openid.LookupClientArgs{
// 			RealmId:  "my-realm",
// 			ClientId: "realm-management",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := realmManagement.Id
// 		_, err = keycloak.LookupRole(ctx, &keycloak.LookupRoleArgs{
// 			RealmId:  "my-realm",
// 			ClientId: &opt0,
// 			Name:     "realm-admin",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm id.
// - `clientId` - (Required) The client id.
//
// ### Attributes Reference
//
// See the docs for the `openid.Client` resource for details on the exported attributes.
func LookupClient(ctx *pulumi.Context, args *LookupClientArgs, opts ...pulumi.InvokeOption) (*LookupClientResult, error) {
	var rv LookupClientResult
	err := ctx.Invoke("keycloak:openid/getClient:getClient", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClient.
type LookupClientArgs struct {
	ClientId string `pulumi:"clientId"`
	RealmId  string `pulumi:"realmId"`
}

// A collection of values returned by getClient.
type LookupClientResult struct {
	AccessType                         string                                      `pulumi:"accessType"`
	AuthenticationFlowBindingOverrides GetClientAuthenticationFlowBindingOverrides `pulumi:"authenticationFlowBindingOverrides"`
	Authorization                      GetClientAuthorization                      `pulumi:"authorization"`
	ClientId                           string                                      `pulumi:"clientId"`
	ClientSecret                       string                                      `pulumi:"clientSecret"`
	ConsentRequired                    bool                                        `pulumi:"consentRequired"`
	Description                        string                                      `pulumi:"description"`
	DirectAccessGrantsEnabled          bool                                        `pulumi:"directAccessGrantsEnabled"`
	Enabled                            bool                                        `pulumi:"enabled"`
	FullScopeAllowed                   bool                                        `pulumi:"fullScopeAllowed"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string   `pulumi:"id"`
	ImplicitFlowEnabled    bool     `pulumi:"implicitFlowEnabled"`
	LoginTheme             string   `pulumi:"loginTheme"`
	Name                   string   `pulumi:"name"`
	RealmId                string   `pulumi:"realmId"`
	ResourceServerId       string   `pulumi:"resourceServerId"`
	RootUrl                string   `pulumi:"rootUrl"`
	ServiceAccountUserId   string   `pulumi:"serviceAccountUserId"`
	ServiceAccountsEnabled bool     `pulumi:"serviceAccountsEnabled"`
	StandardFlowEnabled    bool     `pulumi:"standardFlowEnabled"`
	ValidRedirectUris      []string `pulumi:"validRedirectUris"`
	WebOrigins             []string `pulumi:"webOrigins"`
}
