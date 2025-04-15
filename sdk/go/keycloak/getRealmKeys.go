// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
//
// Remarks:
//
// - A key must meet all filter criteria
// - This data source may return more than one value.
// - If no key matches the filter criteria, then an error will be returned.
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
//			realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
//				Realm:   pulumi.String("my-realm"),
//				Enabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			realmKeys := keycloak.GetRealmKeysOutput(ctx, keycloak.GetRealmKeysOutputArgs{
//				RealmId: realm.ID(),
//				Algorithms: pulumi.StringArray{
//					pulumi.String("AES"),
//					pulumi.String("RS256"),
//				},
//				Statuses: pulumi.StringArray{
//					pulumi.String("ACTIVE"),
//					pulumi.String("PASSIVE"),
//				},
//			}, nil)
//			ctx.Export("certificate", realmKeys.ApplyT(func(realmKeys keycloak.GetRealmKeysResult) (*string, error) {
//				return &realmKeys.Keys[0].Certificate, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
func GetRealmKeys(ctx *pulumi.Context, args *GetRealmKeysArgs, opts ...pulumi.InvokeOption) (*GetRealmKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRealmKeysResult
	err := ctx.Invoke("keycloak:index/getRealmKeys:getRealmKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRealmKeys.
type GetRealmKeysArgs struct {
	// When specified, keys will be filtered by algorithm. The algorithms can be any of `HS256`, `RS256`,`AES`, etc.
	Algorithms []string `pulumi:"algorithms"`
	// The realm from which the keys will be retrieved.
	RealmId string `pulumi:"realmId"`
	// When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
	Statuses []string `pulumi:"statuses"`
}

// A collection of values returned by getRealmKeys.
type GetRealmKeysResult struct {
	Algorithms []string `pulumi:"algorithms"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) A list of keys that match the filter criteria. Each key has the following attributes:
	Keys    []GetRealmKeysKey `pulumi:"keys"`
	RealmId string            `pulumi:"realmId"`
	// Key status (string)
	Statuses []string `pulumi:"statuses"`
}

func GetRealmKeysOutput(ctx *pulumi.Context, args GetRealmKeysOutputArgs, opts ...pulumi.InvokeOption) GetRealmKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetRealmKeysResultOutput, error) {
			args := v.(GetRealmKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("keycloak:index/getRealmKeys:getRealmKeys", args, GetRealmKeysResultOutput{}, options).(GetRealmKeysResultOutput), nil
		}).(GetRealmKeysResultOutput)
}

// A collection of arguments for invoking getRealmKeys.
type GetRealmKeysOutputArgs struct {
	// When specified, keys will be filtered by algorithm. The algorithms can be any of `HS256`, `RS256`,`AES`, etc.
	Algorithms pulumi.StringArrayInput `pulumi:"algorithms"`
	// The realm from which the keys will be retrieved.
	RealmId pulumi.StringInput `pulumi:"realmId"`
	// When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
	Statuses pulumi.StringArrayInput `pulumi:"statuses"`
}

func (GetRealmKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRealmKeysArgs)(nil)).Elem()
}

// A collection of values returned by getRealmKeys.
type GetRealmKeysResultOutput struct{ *pulumi.OutputState }

func (GetRealmKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRealmKeysResult)(nil)).Elem()
}

func (o GetRealmKeysResultOutput) ToGetRealmKeysResultOutput() GetRealmKeysResultOutput {
	return o
}

func (o GetRealmKeysResultOutput) ToGetRealmKeysResultOutputWithContext(ctx context.Context) GetRealmKeysResultOutput {
	return o
}

func (o GetRealmKeysResultOutput) Algorithms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRealmKeysResult) []string { return v.Algorithms }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRealmKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRealmKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Computed) A list of keys that match the filter criteria. Each key has the following attributes:
func (o GetRealmKeysResultOutput) Keys() GetRealmKeysKeyArrayOutput {
	return o.ApplyT(func(v GetRealmKeysResult) []GetRealmKeysKey { return v.Keys }).(GetRealmKeysKeyArrayOutput)
}

func (o GetRealmKeysResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRealmKeysResult) string { return v.RealmId }).(pulumi.StringOutput)
}

// Key status (string)
func (o GetRealmKeysResultOutput) Statuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRealmKeysResult) []string { return v.Statuses }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRealmKeysResultOutput{})
}
