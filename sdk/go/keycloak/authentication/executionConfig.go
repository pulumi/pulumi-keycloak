// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentication

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for managing an authentication execution's configuration. If a particular authentication execution supports additional
// configuration (such as with the `identity-provider-redirector` execution), this can be managed with this resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/authentication"
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
//			flow, err := authentication.NewFlow(ctx, "flow", &authentication.FlowArgs{
//				RealmId: realm.ID(),
//				Alias:   pulumi.String("my-flow-alias"),
//			})
//			if err != nil {
//				return err
//			}
//			execution, err := authentication.NewExecution(ctx, "execution", &authentication.ExecutionArgs{
//				RealmId:         realm.ID(),
//				ParentFlowAlias: flow.Alias,
//				Authenticator:   pulumi.String("identity-provider-redirector"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = authentication.NewExecutionConfig(ctx, "config", &authentication.ExecutionConfigArgs{
//				RealmId:     realm.ID(),
//				ExecutionId: execution.ID(),
//				Alias:       pulumi.String("my-config-alias"),
//				Config: pulumi.StringMap{
//					"defaultProvider": pulumi.String("my-config-default-idp"),
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
// Configurations can be imported using the format `{{realm}}/{{authenticationExecutionId}}/{{authenticationExecutionConfigId}}`. If the `authenticationExecutionId` is incorrect, the import will still be successful. A subsequent apply will change the `authenticationExecutionId` to the correct one, which causes the configuration to be replaced. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:authentication/executionConfig:ExecutionConfig config my-realm/be081463-ddbf-4b42-9eff-9c97886f24ff/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
//
// ```
type ExecutionConfig struct {
	pulumi.CustomResourceState

	// The name of the configuration.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
	Config pulumi.StringMapOutput `pulumi:"config"`
	// The authentication execution this configuration is attached to.
	ExecutionId pulumi.StringOutput `pulumi:"executionId"`
	// The realm the authentication execution exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewExecutionConfig registers a new resource with the given unique name, arguments, and options.
func NewExecutionConfig(ctx *pulumi.Context,
	name string, args *ExecutionConfigArgs, opts ...pulumi.ResourceOption) (*ExecutionConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.ExecutionId == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionId'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource ExecutionConfig
	err := ctx.RegisterResource("keycloak:authentication/executionConfig:ExecutionConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExecutionConfig gets an existing ExecutionConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExecutionConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExecutionConfigState, opts ...pulumi.ResourceOption) (*ExecutionConfig, error) {
	var resource ExecutionConfig
	err := ctx.ReadResource("keycloak:authentication/executionConfig:ExecutionConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExecutionConfig resources.
type executionConfigState struct {
	// The name of the configuration.
	Alias *string `pulumi:"alias"`
	// The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
	Config map[string]string `pulumi:"config"`
	// The authentication execution this configuration is attached to.
	ExecutionId *string `pulumi:"executionId"`
	// The realm the authentication execution exists in.
	RealmId *string `pulumi:"realmId"`
}

type ExecutionConfigState struct {
	// The name of the configuration.
	Alias pulumi.StringPtrInput
	// The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
	Config pulumi.StringMapInput
	// The authentication execution this configuration is attached to.
	ExecutionId pulumi.StringPtrInput
	// The realm the authentication execution exists in.
	RealmId pulumi.StringPtrInput
}

func (ExecutionConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*executionConfigState)(nil)).Elem()
}

type executionConfigArgs struct {
	// The name of the configuration.
	Alias string `pulumi:"alias"`
	// The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
	Config map[string]string `pulumi:"config"`
	// The authentication execution this configuration is attached to.
	ExecutionId string `pulumi:"executionId"`
	// The realm the authentication execution exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a ExecutionConfig resource.
type ExecutionConfigArgs struct {
	// The name of the configuration.
	Alias pulumi.StringInput
	// The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
	Config pulumi.StringMapInput
	// The authentication execution this configuration is attached to.
	ExecutionId pulumi.StringInput
	// The realm the authentication execution exists in.
	RealmId pulumi.StringInput
}

func (ExecutionConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*executionConfigArgs)(nil)).Elem()
}

type ExecutionConfigInput interface {
	pulumi.Input

	ToExecutionConfigOutput() ExecutionConfigOutput
	ToExecutionConfigOutputWithContext(ctx context.Context) ExecutionConfigOutput
}

func (*ExecutionConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ExecutionConfig)(nil)).Elem()
}

func (i *ExecutionConfig) ToExecutionConfigOutput() ExecutionConfigOutput {
	return i.ToExecutionConfigOutputWithContext(context.Background())
}

func (i *ExecutionConfig) ToExecutionConfigOutputWithContext(ctx context.Context) ExecutionConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionConfigOutput)
}

// ExecutionConfigArrayInput is an input type that accepts ExecutionConfigArray and ExecutionConfigArrayOutput values.
// You can construct a concrete instance of `ExecutionConfigArrayInput` via:
//
//	ExecutionConfigArray{ ExecutionConfigArgs{...} }
type ExecutionConfigArrayInput interface {
	pulumi.Input

	ToExecutionConfigArrayOutput() ExecutionConfigArrayOutput
	ToExecutionConfigArrayOutputWithContext(context.Context) ExecutionConfigArrayOutput
}

type ExecutionConfigArray []ExecutionConfigInput

func (ExecutionConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExecutionConfig)(nil)).Elem()
}

func (i ExecutionConfigArray) ToExecutionConfigArrayOutput() ExecutionConfigArrayOutput {
	return i.ToExecutionConfigArrayOutputWithContext(context.Background())
}

func (i ExecutionConfigArray) ToExecutionConfigArrayOutputWithContext(ctx context.Context) ExecutionConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionConfigArrayOutput)
}

// ExecutionConfigMapInput is an input type that accepts ExecutionConfigMap and ExecutionConfigMapOutput values.
// You can construct a concrete instance of `ExecutionConfigMapInput` via:
//
//	ExecutionConfigMap{ "key": ExecutionConfigArgs{...} }
type ExecutionConfigMapInput interface {
	pulumi.Input

	ToExecutionConfigMapOutput() ExecutionConfigMapOutput
	ToExecutionConfigMapOutputWithContext(context.Context) ExecutionConfigMapOutput
}

type ExecutionConfigMap map[string]ExecutionConfigInput

func (ExecutionConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExecutionConfig)(nil)).Elem()
}

func (i ExecutionConfigMap) ToExecutionConfigMapOutput() ExecutionConfigMapOutput {
	return i.ToExecutionConfigMapOutputWithContext(context.Background())
}

func (i ExecutionConfigMap) ToExecutionConfigMapOutputWithContext(ctx context.Context) ExecutionConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionConfigMapOutput)
}

type ExecutionConfigOutput struct{ *pulumi.OutputState }

func (ExecutionConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExecutionConfig)(nil)).Elem()
}

func (o ExecutionConfigOutput) ToExecutionConfigOutput() ExecutionConfigOutput {
	return o
}

func (o ExecutionConfigOutput) ToExecutionConfigOutputWithContext(ctx context.Context) ExecutionConfigOutput {
	return o
}

// The name of the configuration.
func (o ExecutionConfigOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *ExecutionConfig) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
func (o ExecutionConfigOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExecutionConfig) pulumi.StringMapOutput { return v.Config }).(pulumi.StringMapOutput)
}

// The authentication execution this configuration is attached to.
func (o ExecutionConfigOutput) ExecutionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExecutionConfig) pulumi.StringOutput { return v.ExecutionId }).(pulumi.StringOutput)
}

// The realm the authentication execution exists in.
func (o ExecutionConfigOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExecutionConfig) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type ExecutionConfigArrayOutput struct{ *pulumi.OutputState }

func (ExecutionConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExecutionConfig)(nil)).Elem()
}

func (o ExecutionConfigArrayOutput) ToExecutionConfigArrayOutput() ExecutionConfigArrayOutput {
	return o
}

func (o ExecutionConfigArrayOutput) ToExecutionConfigArrayOutputWithContext(ctx context.Context) ExecutionConfigArrayOutput {
	return o
}

func (o ExecutionConfigArrayOutput) Index(i pulumi.IntInput) ExecutionConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExecutionConfig {
		return vs[0].([]*ExecutionConfig)[vs[1].(int)]
	}).(ExecutionConfigOutput)
}

type ExecutionConfigMapOutput struct{ *pulumi.OutputState }

func (ExecutionConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExecutionConfig)(nil)).Elem()
}

func (o ExecutionConfigMapOutput) ToExecutionConfigMapOutput() ExecutionConfigMapOutput {
	return o
}

func (o ExecutionConfigMapOutput) ToExecutionConfigMapOutputWithContext(ctx context.Context) ExecutionConfigMapOutput {
	return o
}

func (o ExecutionConfigMapOutput) MapIndex(k pulumi.StringInput) ExecutionConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExecutionConfig {
		return vs[0].(map[string]*ExecutionConfig)[vs[1].(string)]
	}).(ExecutionConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionConfigInput)(nil)).Elem(), &ExecutionConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionConfigArrayInput)(nil)).Elem(), ExecutionConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionConfigMapInput)(nil)).Elem(), ExecutionConfigMap{})
	pulumi.RegisterOutputType(ExecutionConfigOutput{})
	pulumi.RegisterOutputType(ExecutionConfigArrayOutput{})
	pulumi.RegisterOutputType(ExecutionConfigMapOutput{})
}
