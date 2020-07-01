// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # getRealmKeys data source
//
// Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
//
// Remarks:
//
// - A key must meet all filter criteria
// - This datasource may return more than one value.
// - If no key matches the filter criteria, then an error is returned.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm of which the keys are retrieved.
// - `algorithms` - (Optional) When specified, keys are filtered by algorithm (values for algorithm: `HS256`, `RS256`,`AES`, ...)
// - `status` - (Optional) When specified, keys are filtered by status (values for status: `ACTIVE`, `DISABLED` and `PASSIVE`)
func GetRealmKeys(ctx *pulumi.Context, args *GetRealmKeysArgs, opts ...pulumi.InvokeOption) (*GetRealmKeysResult, error) {
	var rv GetRealmKeysResult
	err := ctx.Invoke("keycloak:index/getRealmKeys:getRealmKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRealmKeys.
type GetRealmKeysArgs struct {
	Algorithms []string `pulumi:"algorithms"`
	RealmId    string   `pulumi:"realmId"`
	Statuses   []string `pulumi:"statuses"`
}

// A collection of values returned by getRealmKeys.
type GetRealmKeysResult struct {
	Algorithms []string `pulumi:"algorithms"`
	// The provider-assigned unique ID for this managed resource.
	Id       string            `pulumi:"id"`
	Keys     []GetRealmKeysKey `pulumi:"keys"`
	RealmId  string            `pulumi:"realmId"`
	Statuses []string          `pulumi:"statuses"`
}
