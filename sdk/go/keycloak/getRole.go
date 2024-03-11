// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Role data source
//
// This data source can be used to fetch properties of a Keycloak role for
// usage with other resources, such as `GroupRoles`.
func LookupRole(ctx *pulumi.Context, args *LookupRoleArgs, opts ...pulumi.InvokeOption) (*LookupRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRoleResult
	err := ctx.Invoke("keycloak:index/getRole:getRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRole.
type LookupRoleArgs struct {
	ClientId *string `pulumi:"clientId"`
	Name     string  `pulumi:"name"`
	RealmId  string  `pulumi:"realmId"`
}

// A collection of values returned by getRole.
type LookupRoleResult struct {
	Attributes     map[string]interface{} `pulumi:"attributes"`
	ClientId       *string                `pulumi:"clientId"`
	CompositeRoles []string               `pulumi:"compositeRoles"`
	Description    string                 `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	RealmId string `pulumi:"realmId"`
}

func LookupRoleOutput(ctx *pulumi.Context, args LookupRoleOutputArgs, opts ...pulumi.InvokeOption) LookupRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleResult, error) {
			args := v.(LookupRoleArgs)
			r, err := LookupRole(ctx, &args, opts...)
			var s LookupRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleResultOutput)
}

// A collection of arguments for invoking getRole.
type LookupRoleOutputArgs struct {
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
	Name     pulumi.StringInput    `pulumi:"name"`
	RealmId  pulumi.StringInput    `pulumi:"realmId"`
}

func (LookupRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleArgs)(nil)).Elem()
}

// A collection of values returned by getRole.
type LookupRoleResultOutput struct{ *pulumi.OutputState }

func (LookupRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleResult)(nil)).Elem()
}

func (o LookupRoleResultOutput) ToLookupRoleResultOutput() LookupRoleResultOutput {
	return o
}

func (o LookupRoleResultOutput) ToLookupRoleResultOutputWithContext(ctx context.Context) LookupRoleResultOutput {
	return o
}

func (o LookupRoleResultOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v LookupRoleResult) map[string]interface{} { return v.Attributes }).(pulumi.MapOutput)
}

func (o LookupRoleResultOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleResult) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o LookupRoleResultOutput) CompositeRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRoleResult) []string { return v.CompositeRoles }).(pulumi.StringArrayOutput)
}

func (o LookupRoleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleResult) string { return v.RealmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleResultOutput{})
}
