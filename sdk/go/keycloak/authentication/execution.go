// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentication

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing an authentication execution within Keycloak.
//
// An authentication execution is an action that the user or service may or may not take when authenticating through an authentication
// flow.
//
// > Following limitation affects Keycloak < 25:  Due to limitations in the Keycloak API, the ordering of authentication executions within a flow must be specified using `dependsOn`. Authentication executions that are created first will appear first within the flow.
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
//			// first execution
//			_, err = authentication.NewExecution(ctx, "execution_one", &authentication.ExecutionArgs{
//				RealmId:         realm.ID(),
//				ParentFlowAlias: flow.Alias,
//				Authenticator:   pulumi.String("auth-cookie"),
//				Requirement:     pulumi.String("ALTERNATIVE"),
//				Priority:        pulumi.Int(10),
//			})
//			if err != nil {
//				return err
//			}
//			// second execution
//			_, err = authentication.NewExecution(ctx, "execution_two", &authentication.ExecutionArgs{
//				RealmId:         realm.ID(),
//				ParentFlowAlias: flow.Alias,
//				Authenticator:   pulumi.String("identity-provider-redirector"),
//				Requirement:     pulumi.String("ALTERNATIVE"),
//				Priority:        pulumi.Int(20),
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
// Authentication executions can be imported using the formats: `{{realmId}}/{{parentFlowAlias}}/{{authenticationExecutionId}}`.
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:authentication/execution:Execution execution_one my-realm/my-flow-alias/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
// ```
type Execution struct {
	pulumi.CustomResourceState

	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator pulumi.StringOutput `pulumi:"authenticator"`
	// The alias of the flow this execution is attached to.
	ParentFlowAlias pulumi.StringOutput `pulumi:"parentFlowAlias"`
	// The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak >= 25).
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The realm the authentication execution exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement pulumi.StringPtrOutput `pulumi:"requirement"`
}

// NewExecution registers a new resource with the given unique name, arguments, and options.
func NewExecution(ctx *pulumi.Context,
	name string, args *ExecutionArgs, opts ...pulumi.ResourceOption) (*Execution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authenticator == nil {
		return nil, errors.New("invalid value for required argument 'Authenticator'")
	}
	if args.ParentFlowAlias == nil {
		return nil, errors.New("invalid value for required argument 'ParentFlowAlias'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Execution
	err := ctx.RegisterResource("keycloak:authentication/execution:Execution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExecution gets an existing Execution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExecution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExecutionState, opts ...pulumi.ResourceOption) (*Execution, error) {
	var resource Execution
	err := ctx.ReadResource("keycloak:authentication/execution:Execution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Execution resources.
type executionState struct {
	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator *string `pulumi:"authenticator"`
	// The alias of the flow this execution is attached to.
	ParentFlowAlias *string `pulumi:"parentFlowAlias"`
	// The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak >= 25).
	Priority *int `pulumi:"priority"`
	// The realm the authentication execution exists in.
	RealmId *string `pulumi:"realmId"`
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement *string `pulumi:"requirement"`
}

type ExecutionState struct {
	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator pulumi.StringPtrInput
	// The alias of the flow this execution is attached to.
	ParentFlowAlias pulumi.StringPtrInput
	// The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak >= 25).
	Priority pulumi.IntPtrInput
	// The realm the authentication execution exists in.
	RealmId pulumi.StringPtrInput
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement pulumi.StringPtrInput
}

func (ExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*executionState)(nil)).Elem()
}

type executionArgs struct {
	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator string `pulumi:"authenticator"`
	// The alias of the flow this execution is attached to.
	ParentFlowAlias string `pulumi:"parentFlowAlias"`
	// The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak >= 25).
	Priority *int `pulumi:"priority"`
	// The realm the authentication execution exists in.
	RealmId string `pulumi:"realmId"`
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement *string `pulumi:"requirement"`
}

// The set of arguments for constructing a Execution resource.
type ExecutionArgs struct {
	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator pulumi.StringInput
	// The alias of the flow this execution is attached to.
	ParentFlowAlias pulumi.StringInput
	// The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak >= 25).
	Priority pulumi.IntPtrInput
	// The realm the authentication execution exists in.
	RealmId pulumi.StringInput
	// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
	Requirement pulumi.StringPtrInput
}

func (ExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*executionArgs)(nil)).Elem()
}

type ExecutionInput interface {
	pulumi.Input

	ToExecutionOutput() ExecutionOutput
	ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput
}

func (*Execution) ElementType() reflect.Type {
	return reflect.TypeOf((**Execution)(nil)).Elem()
}

func (i *Execution) ToExecutionOutput() ExecutionOutput {
	return i.ToExecutionOutputWithContext(context.Background())
}

func (i *Execution) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionOutput)
}

// ExecutionArrayInput is an input type that accepts ExecutionArray and ExecutionArrayOutput values.
// You can construct a concrete instance of `ExecutionArrayInput` via:
//
//	ExecutionArray{ ExecutionArgs{...} }
type ExecutionArrayInput interface {
	pulumi.Input

	ToExecutionArrayOutput() ExecutionArrayOutput
	ToExecutionArrayOutputWithContext(context.Context) ExecutionArrayOutput
}

type ExecutionArray []ExecutionInput

func (ExecutionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Execution)(nil)).Elem()
}

func (i ExecutionArray) ToExecutionArrayOutput() ExecutionArrayOutput {
	return i.ToExecutionArrayOutputWithContext(context.Background())
}

func (i ExecutionArray) ToExecutionArrayOutputWithContext(ctx context.Context) ExecutionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionArrayOutput)
}

// ExecutionMapInput is an input type that accepts ExecutionMap and ExecutionMapOutput values.
// You can construct a concrete instance of `ExecutionMapInput` via:
//
//	ExecutionMap{ "key": ExecutionArgs{...} }
type ExecutionMapInput interface {
	pulumi.Input

	ToExecutionMapOutput() ExecutionMapOutput
	ToExecutionMapOutputWithContext(context.Context) ExecutionMapOutput
}

type ExecutionMap map[string]ExecutionInput

func (ExecutionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Execution)(nil)).Elem()
}

func (i ExecutionMap) ToExecutionMapOutput() ExecutionMapOutput {
	return i.ToExecutionMapOutputWithContext(context.Background())
}

func (i ExecutionMap) ToExecutionMapOutputWithContext(ctx context.Context) ExecutionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExecutionMapOutput)
}

type ExecutionOutput struct{ *pulumi.OutputState }

func (ExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Execution)(nil)).Elem()
}

func (o ExecutionOutput) ToExecutionOutput() ExecutionOutput {
	return o
}

func (o ExecutionOutput) ToExecutionOutputWithContext(ctx context.Context) ExecutionOutput {
	return o
}

// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
func (o ExecutionOutput) Authenticator() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.Authenticator }).(pulumi.StringOutput)
}

// The alias of the flow this execution is attached to.
func (o ExecutionOutput) ParentFlowAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.ParentFlowAlias }).(pulumi.StringOutput)
}

// The authenticator priority. Lower values will be executed prior higher values (Only supported by Keycloak >= 25).
func (o ExecutionOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Execution) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The realm the authentication execution exists in.
func (o ExecutionOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`, or `DISABLED`. Defaults to `DISABLED`.
func (o ExecutionOutput) Requirement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Execution) pulumi.StringPtrOutput { return v.Requirement }).(pulumi.StringPtrOutput)
}

type ExecutionArrayOutput struct{ *pulumi.OutputState }

func (ExecutionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Execution)(nil)).Elem()
}

func (o ExecutionArrayOutput) ToExecutionArrayOutput() ExecutionArrayOutput {
	return o
}

func (o ExecutionArrayOutput) ToExecutionArrayOutputWithContext(ctx context.Context) ExecutionArrayOutput {
	return o
}

func (o ExecutionArrayOutput) Index(i pulumi.IntInput) ExecutionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Execution {
		return vs[0].([]*Execution)[vs[1].(int)]
	}).(ExecutionOutput)
}

type ExecutionMapOutput struct{ *pulumi.OutputState }

func (ExecutionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Execution)(nil)).Elem()
}

func (o ExecutionMapOutput) ToExecutionMapOutput() ExecutionMapOutput {
	return o
}

func (o ExecutionMapOutput) ToExecutionMapOutputWithContext(ctx context.Context) ExecutionMapOutput {
	return o
}

func (o ExecutionMapOutput) MapIndex(k pulumi.StringInput) ExecutionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Execution {
		return vs[0].(map[string]*Execution)[vs[1].(string)]
	}).(ExecutionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionInput)(nil)).Elem(), &Execution{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionArrayInput)(nil)).Elem(), ExecutionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExecutionMapInput)(nil)).Elem(), ExecutionMap{})
	pulumi.RegisterOutputType(ExecutionOutput{})
	pulumi.RegisterOutputType(ExecutionArrayOutput{})
	pulumi.RegisterOutputType(ExecutionMapOutput{})
}
