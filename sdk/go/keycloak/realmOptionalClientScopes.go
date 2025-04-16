// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows you to manage the set of optional client scopes for a Keycloak realm, which are used when new clients are created.
//
// Note that this resource attempts to be an **authoritative** source over the optional client scopes for a Keycloak realm,
// so any Keycloak defaults and manual adjustments will be overwritten.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/openid"
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
//			clientScope, err := openid.NewClientScope(ctx, "client_scope", &openid.ClientScopeArgs{
//				RealmId: realm.ID(),
//				Name:    pulumi.String("test-client-scope"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRealmOptionalClientScopes(ctx, "optional_scopes", &keycloak.RealmOptionalClientScopesArgs{
//				RealmId: realm.ID(),
//				OptionalScopes: pulumi.StringArray{
//					pulumi.String("address"),
//					pulumi.String("phone"),
//					pulumi.String("offline_access"),
//					pulumi.String("microprofile-jwt"),
//					clientScope.Name,
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
// This resource does not support import. Instead of importing, feel free to create this resource
//
// as if it did not already exist on the server.
type RealmOptionalClientScopes struct {
	pulumi.CustomResourceState

	// An array of optional client scope names that should be used when creating new Keycloak clients.
	OptionalScopes pulumi.StringArrayOutput `pulumi:"optionalScopes"`
	// The realm this client and scopes exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewRealmOptionalClientScopes registers a new resource with the given unique name, arguments, and options.
func NewRealmOptionalClientScopes(ctx *pulumi.Context,
	name string, args *RealmOptionalClientScopesArgs, opts ...pulumi.ResourceOption) (*RealmOptionalClientScopes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OptionalScopes == nil {
		return nil, errors.New("invalid value for required argument 'OptionalScopes'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RealmOptionalClientScopes
	err := ctx.RegisterResource("keycloak:index/realmOptionalClientScopes:RealmOptionalClientScopes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealmOptionalClientScopes gets an existing RealmOptionalClientScopes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealmOptionalClientScopes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmOptionalClientScopesState, opts ...pulumi.ResourceOption) (*RealmOptionalClientScopes, error) {
	var resource RealmOptionalClientScopes
	err := ctx.ReadResource("keycloak:index/realmOptionalClientScopes:RealmOptionalClientScopes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealmOptionalClientScopes resources.
type realmOptionalClientScopesState struct {
	// An array of optional client scope names that should be used when creating new Keycloak clients.
	OptionalScopes []string `pulumi:"optionalScopes"`
	// The realm this client and scopes exists in.
	RealmId *string `pulumi:"realmId"`
}

type RealmOptionalClientScopesState struct {
	// An array of optional client scope names that should be used when creating new Keycloak clients.
	OptionalScopes pulumi.StringArrayInput
	// The realm this client and scopes exists in.
	RealmId pulumi.StringPtrInput
}

func (RealmOptionalClientScopesState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmOptionalClientScopesState)(nil)).Elem()
}

type realmOptionalClientScopesArgs struct {
	// An array of optional client scope names that should be used when creating new Keycloak clients.
	OptionalScopes []string `pulumi:"optionalScopes"`
	// The realm this client and scopes exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a RealmOptionalClientScopes resource.
type RealmOptionalClientScopesArgs struct {
	// An array of optional client scope names that should be used when creating new Keycloak clients.
	OptionalScopes pulumi.StringArrayInput
	// The realm this client and scopes exists in.
	RealmId pulumi.StringInput
}

func (RealmOptionalClientScopesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmOptionalClientScopesArgs)(nil)).Elem()
}

type RealmOptionalClientScopesInput interface {
	pulumi.Input

	ToRealmOptionalClientScopesOutput() RealmOptionalClientScopesOutput
	ToRealmOptionalClientScopesOutputWithContext(ctx context.Context) RealmOptionalClientScopesOutput
}

func (*RealmOptionalClientScopes) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmOptionalClientScopes)(nil)).Elem()
}

func (i *RealmOptionalClientScopes) ToRealmOptionalClientScopesOutput() RealmOptionalClientScopesOutput {
	return i.ToRealmOptionalClientScopesOutputWithContext(context.Background())
}

func (i *RealmOptionalClientScopes) ToRealmOptionalClientScopesOutputWithContext(ctx context.Context) RealmOptionalClientScopesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmOptionalClientScopesOutput)
}

// RealmOptionalClientScopesArrayInput is an input type that accepts RealmOptionalClientScopesArray and RealmOptionalClientScopesArrayOutput values.
// You can construct a concrete instance of `RealmOptionalClientScopesArrayInput` via:
//
//	RealmOptionalClientScopesArray{ RealmOptionalClientScopesArgs{...} }
type RealmOptionalClientScopesArrayInput interface {
	pulumi.Input

	ToRealmOptionalClientScopesArrayOutput() RealmOptionalClientScopesArrayOutput
	ToRealmOptionalClientScopesArrayOutputWithContext(context.Context) RealmOptionalClientScopesArrayOutput
}

type RealmOptionalClientScopesArray []RealmOptionalClientScopesInput

func (RealmOptionalClientScopesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmOptionalClientScopes)(nil)).Elem()
}

func (i RealmOptionalClientScopesArray) ToRealmOptionalClientScopesArrayOutput() RealmOptionalClientScopesArrayOutput {
	return i.ToRealmOptionalClientScopesArrayOutputWithContext(context.Background())
}

func (i RealmOptionalClientScopesArray) ToRealmOptionalClientScopesArrayOutputWithContext(ctx context.Context) RealmOptionalClientScopesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmOptionalClientScopesArrayOutput)
}

// RealmOptionalClientScopesMapInput is an input type that accepts RealmOptionalClientScopesMap and RealmOptionalClientScopesMapOutput values.
// You can construct a concrete instance of `RealmOptionalClientScopesMapInput` via:
//
//	RealmOptionalClientScopesMap{ "key": RealmOptionalClientScopesArgs{...} }
type RealmOptionalClientScopesMapInput interface {
	pulumi.Input

	ToRealmOptionalClientScopesMapOutput() RealmOptionalClientScopesMapOutput
	ToRealmOptionalClientScopesMapOutputWithContext(context.Context) RealmOptionalClientScopesMapOutput
}

type RealmOptionalClientScopesMap map[string]RealmOptionalClientScopesInput

func (RealmOptionalClientScopesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmOptionalClientScopes)(nil)).Elem()
}

func (i RealmOptionalClientScopesMap) ToRealmOptionalClientScopesMapOutput() RealmOptionalClientScopesMapOutput {
	return i.ToRealmOptionalClientScopesMapOutputWithContext(context.Background())
}

func (i RealmOptionalClientScopesMap) ToRealmOptionalClientScopesMapOutputWithContext(ctx context.Context) RealmOptionalClientScopesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmOptionalClientScopesMapOutput)
}

type RealmOptionalClientScopesOutput struct{ *pulumi.OutputState }

func (RealmOptionalClientScopesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmOptionalClientScopes)(nil)).Elem()
}

func (o RealmOptionalClientScopesOutput) ToRealmOptionalClientScopesOutput() RealmOptionalClientScopesOutput {
	return o
}

func (o RealmOptionalClientScopesOutput) ToRealmOptionalClientScopesOutputWithContext(ctx context.Context) RealmOptionalClientScopesOutput {
	return o
}

// An array of optional client scope names that should be used when creating new Keycloak clients.
func (o RealmOptionalClientScopesOutput) OptionalScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RealmOptionalClientScopes) pulumi.StringArrayOutput { return v.OptionalScopes }).(pulumi.StringArrayOutput)
}

// The realm this client and scopes exists in.
func (o RealmOptionalClientScopesOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmOptionalClientScopes) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type RealmOptionalClientScopesArrayOutput struct{ *pulumi.OutputState }

func (RealmOptionalClientScopesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmOptionalClientScopes)(nil)).Elem()
}

func (o RealmOptionalClientScopesArrayOutput) ToRealmOptionalClientScopesArrayOutput() RealmOptionalClientScopesArrayOutput {
	return o
}

func (o RealmOptionalClientScopesArrayOutput) ToRealmOptionalClientScopesArrayOutputWithContext(ctx context.Context) RealmOptionalClientScopesArrayOutput {
	return o
}

func (o RealmOptionalClientScopesArrayOutput) Index(i pulumi.IntInput) RealmOptionalClientScopesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealmOptionalClientScopes {
		return vs[0].([]*RealmOptionalClientScopes)[vs[1].(int)]
	}).(RealmOptionalClientScopesOutput)
}

type RealmOptionalClientScopesMapOutput struct{ *pulumi.OutputState }

func (RealmOptionalClientScopesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmOptionalClientScopes)(nil)).Elem()
}

func (o RealmOptionalClientScopesMapOutput) ToRealmOptionalClientScopesMapOutput() RealmOptionalClientScopesMapOutput {
	return o
}

func (o RealmOptionalClientScopesMapOutput) ToRealmOptionalClientScopesMapOutputWithContext(ctx context.Context) RealmOptionalClientScopesMapOutput {
	return o
}

func (o RealmOptionalClientScopesMapOutput) MapIndex(k pulumi.StringInput) RealmOptionalClientScopesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealmOptionalClientScopes {
		return vs[0].(map[string]*RealmOptionalClientScopes)[vs[1].(string)]
	}).(RealmOptionalClientScopesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealmOptionalClientScopesInput)(nil)).Elem(), &RealmOptionalClientScopes{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmOptionalClientScopesArrayInput)(nil)).Elem(), RealmOptionalClientScopesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmOptionalClientScopesMapInput)(nil)).Elem(), RealmOptionalClientScopesMap{})
	pulumi.RegisterOutputType(RealmOptionalClientScopesOutput{})
	pulumi.RegisterOutputType(RealmOptionalClientScopesArrayOutput{})
	pulumi.RegisterOutputType(RealmOptionalClientScopesMapOutput{})
}
