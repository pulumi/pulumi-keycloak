// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows you to manage openid Client Authorization Client Scope type Policies.
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
//			test, err := openid.NewClient(ctx, "test", &openid.ClientArgs{
//				ClientId:               pulumi.String("client_id"),
//				RealmId:                realm.ID(),
//				AccessType:             pulumi.String("CONFIDENTIAL"),
//				ServiceAccountsEnabled: pulumi.Bool(true),
//				Authorization: &openid.ClientAuthorizationArgs{
//					PolicyEnforcementMode: pulumi.String("ENFORCING"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			test1, err := openid.NewClientScope(ctx, "test1", &openid.ClientScopeArgs{
//				RealmId:     realm.ID(),
//				Name:        pulumi.String("test1"),
//				Description: pulumi.String("test1"),
//			})
//			if err != nil {
//				return err
//			}
//			test2, err := openid.NewClientScope(ctx, "test2", &openid.ClientScopeArgs{
//				RealmId:     realm.ID(),
//				Name:        pulumi.String("test2"),
//				Description: pulumi.String("test2"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewClientAuthorizationClientScopePolicy(ctx, "test", &openid.ClientAuthorizationClientScopePolicyArgs{
//				ResourceServerId: test.ResourceServerId,
//				RealmId:          realm.ID(),
//				Name:             pulumi.String("test_policy_single"),
//				Description:      pulumi.String("test"),
//				DecisionStrategy: pulumi.String("AFFIRMATIVE"),
//				Logic:            pulumi.String("POSITIVE"),
//				Scopes: openid.ClientAuthorizationClientScopePolicyScopeArray{
//					&openid.ClientAuthorizationClientScopePolicyScopeArgs{
//						Id:       test1.ID(),
//						Required: pulumi.Bool(false),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewClientAuthorizationClientScopePolicy(ctx, "test_multiple", &openid.ClientAuthorizationClientScopePolicyArgs{
//				ResourceServerId: test.ResourceServerId,
//				RealmId:          realm.ID(),
//				Name:             pulumi.String("test_policy_multiple"),
//				Description:      pulumi.String("test"),
//				DecisionStrategy: pulumi.String("AFFIRMATIVE"),
//				Logic:            pulumi.String("POSITIVE"),
//				Scopes: openid.ClientAuthorizationClientScopePolicyScopeArray{
//					&openid.ClientAuthorizationClientScopePolicyScopeArgs{
//						Id:       test1.ID(),
//						Required: pulumi.Bool(false),
//					},
//					&openid.ClientAuthorizationClientScopePolicyScopeArgs{
//						Id:       test2.ID(),
//						Required: pulumi.Bool(true),
//					},
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
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm this group exists in.
// - `resourceServerId` - (Required) The ID of the resource server.
// - `name` - (Required) The name of the policy.
// - `description` - (Optional) A description for the authorization policy.
// - `decisionStrategy` - (Optional) The decision strategy, can be one of `UNANIMOUS`, `AFFIRMATIVE`, or `CONSENSUS`. Defaults to `UNANIMOUS`.
// - `logic` - (Optional) The logic, can be one of `POSITIVE` or `NEGATIVE`. Defaults to `POSITIVE`.
// - `scope` - An client scope to add client scope. At least one should be defined.
//
// ### Scope Arguments
//
// - `id` - (Required) Id of client scope.
// - `required` - (Optional) When `true`, then this client scope will be set as required. Defaults to `false`.
//
// ### Attributes Reference
//
// In addition to the arguments listed above, the following computed attributes are exported:
//
// - `id` - Policy ID representing the policy.
//
// ## Import
//
// Client authorization policies can be imported using the format: `{{realmId}}/{{resourceServerId}}/{{policyId}}`.
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:openid/clientAuthorizationClientScopePolicy:ClientAuthorizationClientScopePolicy test my-realm/3bd4a686-1062-4b59-97b8-e4e3f10b99da/63b3cde8-987d-4cd9-9306-1955579281d9
// ```
type ClientAuthorizationClientScopePolicy struct {
	pulumi.CustomResourceState

	DecisionStrategy pulumi.StringPtrOutput                               `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput                               `pulumi:"description"`
	Logic            pulumi.StringPtrOutput                               `pulumi:"logic"`
	Name             pulumi.StringOutput                                  `pulumi:"name"`
	RealmId          pulumi.StringOutput                                  `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput                                  `pulumi:"resourceServerId"`
	Scopes           ClientAuthorizationClientScopePolicyScopeArrayOutput `pulumi:"scopes"`
}

// NewClientAuthorizationClientScopePolicy registers a new resource with the given unique name, arguments, and options.
func NewClientAuthorizationClientScopePolicy(ctx *pulumi.Context,
	name string, args *ClientAuthorizationClientScopePolicyArgs, opts ...pulumi.ResourceOption) (*ClientAuthorizationClientScopePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.ResourceServerId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceServerId'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClientAuthorizationClientScopePolicy
	err := ctx.RegisterResource("keycloak:openid/clientAuthorizationClientScopePolicy:ClientAuthorizationClientScopePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientAuthorizationClientScopePolicy gets an existing ClientAuthorizationClientScopePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientAuthorizationClientScopePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientAuthorizationClientScopePolicyState, opts ...pulumi.ResourceOption) (*ClientAuthorizationClientScopePolicy, error) {
	var resource ClientAuthorizationClientScopePolicy
	err := ctx.ReadResource("keycloak:openid/clientAuthorizationClientScopePolicy:ClientAuthorizationClientScopePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientAuthorizationClientScopePolicy resources.
type clientAuthorizationClientScopePolicyState struct {
	DecisionStrategy *string                                     `pulumi:"decisionStrategy"`
	Description      *string                                     `pulumi:"description"`
	Logic            *string                                     `pulumi:"logic"`
	Name             *string                                     `pulumi:"name"`
	RealmId          *string                                     `pulumi:"realmId"`
	ResourceServerId *string                                     `pulumi:"resourceServerId"`
	Scopes           []ClientAuthorizationClientScopePolicyScope `pulumi:"scopes"`
}

type ClientAuthorizationClientScopePolicyState struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
	Scopes           ClientAuthorizationClientScopePolicyScopeArrayInput
}

func (ClientAuthorizationClientScopePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationClientScopePolicyState)(nil)).Elem()
}

type clientAuthorizationClientScopePolicyArgs struct {
	DecisionStrategy *string                                     `pulumi:"decisionStrategy"`
	Description      *string                                     `pulumi:"description"`
	Logic            *string                                     `pulumi:"logic"`
	Name             *string                                     `pulumi:"name"`
	RealmId          string                                      `pulumi:"realmId"`
	ResourceServerId string                                      `pulumi:"resourceServerId"`
	Scopes           []ClientAuthorizationClientScopePolicyScope `pulumi:"scopes"`
}

// The set of arguments for constructing a ClientAuthorizationClientScopePolicy resource.
type ClientAuthorizationClientScopePolicyArgs struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
	Scopes           ClientAuthorizationClientScopePolicyScopeArrayInput
}

func (ClientAuthorizationClientScopePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationClientScopePolicyArgs)(nil)).Elem()
}

type ClientAuthorizationClientScopePolicyInput interface {
	pulumi.Input

	ToClientAuthorizationClientScopePolicyOutput() ClientAuthorizationClientScopePolicyOutput
	ToClientAuthorizationClientScopePolicyOutputWithContext(ctx context.Context) ClientAuthorizationClientScopePolicyOutput
}

func (*ClientAuthorizationClientScopePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorizationClientScopePolicy)(nil)).Elem()
}

func (i *ClientAuthorizationClientScopePolicy) ToClientAuthorizationClientScopePolicyOutput() ClientAuthorizationClientScopePolicyOutput {
	return i.ToClientAuthorizationClientScopePolicyOutputWithContext(context.Background())
}

func (i *ClientAuthorizationClientScopePolicy) ToClientAuthorizationClientScopePolicyOutputWithContext(ctx context.Context) ClientAuthorizationClientScopePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationClientScopePolicyOutput)
}

// ClientAuthorizationClientScopePolicyArrayInput is an input type that accepts ClientAuthorizationClientScopePolicyArray and ClientAuthorizationClientScopePolicyArrayOutput values.
// You can construct a concrete instance of `ClientAuthorizationClientScopePolicyArrayInput` via:
//
//	ClientAuthorizationClientScopePolicyArray{ ClientAuthorizationClientScopePolicyArgs{...} }
type ClientAuthorizationClientScopePolicyArrayInput interface {
	pulumi.Input

	ToClientAuthorizationClientScopePolicyArrayOutput() ClientAuthorizationClientScopePolicyArrayOutput
	ToClientAuthorizationClientScopePolicyArrayOutputWithContext(context.Context) ClientAuthorizationClientScopePolicyArrayOutput
}

type ClientAuthorizationClientScopePolicyArray []ClientAuthorizationClientScopePolicyInput

func (ClientAuthorizationClientScopePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientAuthorizationClientScopePolicy)(nil)).Elem()
}

func (i ClientAuthorizationClientScopePolicyArray) ToClientAuthorizationClientScopePolicyArrayOutput() ClientAuthorizationClientScopePolicyArrayOutput {
	return i.ToClientAuthorizationClientScopePolicyArrayOutputWithContext(context.Background())
}

func (i ClientAuthorizationClientScopePolicyArray) ToClientAuthorizationClientScopePolicyArrayOutputWithContext(ctx context.Context) ClientAuthorizationClientScopePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationClientScopePolicyArrayOutput)
}

// ClientAuthorizationClientScopePolicyMapInput is an input type that accepts ClientAuthorizationClientScopePolicyMap and ClientAuthorizationClientScopePolicyMapOutput values.
// You can construct a concrete instance of `ClientAuthorizationClientScopePolicyMapInput` via:
//
//	ClientAuthorizationClientScopePolicyMap{ "key": ClientAuthorizationClientScopePolicyArgs{...} }
type ClientAuthorizationClientScopePolicyMapInput interface {
	pulumi.Input

	ToClientAuthorizationClientScopePolicyMapOutput() ClientAuthorizationClientScopePolicyMapOutput
	ToClientAuthorizationClientScopePolicyMapOutputWithContext(context.Context) ClientAuthorizationClientScopePolicyMapOutput
}

type ClientAuthorizationClientScopePolicyMap map[string]ClientAuthorizationClientScopePolicyInput

func (ClientAuthorizationClientScopePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientAuthorizationClientScopePolicy)(nil)).Elem()
}

func (i ClientAuthorizationClientScopePolicyMap) ToClientAuthorizationClientScopePolicyMapOutput() ClientAuthorizationClientScopePolicyMapOutput {
	return i.ToClientAuthorizationClientScopePolicyMapOutputWithContext(context.Background())
}

func (i ClientAuthorizationClientScopePolicyMap) ToClientAuthorizationClientScopePolicyMapOutputWithContext(ctx context.Context) ClientAuthorizationClientScopePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationClientScopePolicyMapOutput)
}

type ClientAuthorizationClientScopePolicyOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationClientScopePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorizationClientScopePolicy)(nil)).Elem()
}

func (o ClientAuthorizationClientScopePolicyOutput) ToClientAuthorizationClientScopePolicyOutput() ClientAuthorizationClientScopePolicyOutput {
	return o
}

func (o ClientAuthorizationClientScopePolicyOutput) ToClientAuthorizationClientScopePolicyOutputWithContext(ctx context.Context) ClientAuthorizationClientScopePolicyOutput {
	return o
}

func (o ClientAuthorizationClientScopePolicyOutput) DecisionStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthorizationClientScopePolicy) pulumi.StringPtrOutput { return v.DecisionStrategy }).(pulumi.StringPtrOutput)
}

func (o ClientAuthorizationClientScopePolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthorizationClientScopePolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ClientAuthorizationClientScopePolicyOutput) Logic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthorizationClientScopePolicy) pulumi.StringPtrOutput { return v.Logic }).(pulumi.StringPtrOutput)
}

func (o ClientAuthorizationClientScopePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientAuthorizationClientScopePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClientAuthorizationClientScopePolicyOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientAuthorizationClientScopePolicy) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

func (o ClientAuthorizationClientScopePolicyOutput) ResourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientAuthorizationClientScopePolicy) pulumi.StringOutput { return v.ResourceServerId }).(pulumi.StringOutput)
}

func (o ClientAuthorizationClientScopePolicyOutput) Scopes() ClientAuthorizationClientScopePolicyScopeArrayOutput {
	return o.ApplyT(func(v *ClientAuthorizationClientScopePolicy) ClientAuthorizationClientScopePolicyScopeArrayOutput {
		return v.Scopes
	}).(ClientAuthorizationClientScopePolicyScopeArrayOutput)
}

type ClientAuthorizationClientScopePolicyArrayOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationClientScopePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientAuthorizationClientScopePolicy)(nil)).Elem()
}

func (o ClientAuthorizationClientScopePolicyArrayOutput) ToClientAuthorizationClientScopePolicyArrayOutput() ClientAuthorizationClientScopePolicyArrayOutput {
	return o
}

func (o ClientAuthorizationClientScopePolicyArrayOutput) ToClientAuthorizationClientScopePolicyArrayOutputWithContext(ctx context.Context) ClientAuthorizationClientScopePolicyArrayOutput {
	return o
}

func (o ClientAuthorizationClientScopePolicyArrayOutput) Index(i pulumi.IntInput) ClientAuthorizationClientScopePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientAuthorizationClientScopePolicy {
		return vs[0].([]*ClientAuthorizationClientScopePolicy)[vs[1].(int)]
	}).(ClientAuthorizationClientScopePolicyOutput)
}

type ClientAuthorizationClientScopePolicyMapOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationClientScopePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientAuthorizationClientScopePolicy)(nil)).Elem()
}

func (o ClientAuthorizationClientScopePolicyMapOutput) ToClientAuthorizationClientScopePolicyMapOutput() ClientAuthorizationClientScopePolicyMapOutput {
	return o
}

func (o ClientAuthorizationClientScopePolicyMapOutput) ToClientAuthorizationClientScopePolicyMapOutputWithContext(ctx context.Context) ClientAuthorizationClientScopePolicyMapOutput {
	return o
}

func (o ClientAuthorizationClientScopePolicyMapOutput) MapIndex(k pulumi.StringInput) ClientAuthorizationClientScopePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientAuthorizationClientScopePolicy {
		return vs[0].(map[string]*ClientAuthorizationClientScopePolicy)[vs[1].(string)]
	}).(ClientAuthorizationClientScopePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientAuthorizationClientScopePolicyInput)(nil)).Elem(), &ClientAuthorizationClientScopePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientAuthorizationClientScopePolicyArrayInput)(nil)).Elem(), ClientAuthorizationClientScopePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientAuthorizationClientScopePolicyMapInput)(nil)).Elem(), ClientAuthorizationClientScopePolicyMap{})
	pulumi.RegisterOutputType(ClientAuthorizationClientScopePolicyOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationClientScopePolicyArrayOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationClientScopePolicyMapOutput{})
}
