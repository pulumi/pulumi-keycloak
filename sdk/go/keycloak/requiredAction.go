// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing required actions within Keycloak.
//
// [Required actions](https://www.keycloak.org/docs/latest/server_admin/#con-required-actions_server_administration_guide) specify actions required before the first login of all new users.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
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
//			_, err = keycloak.NewRequiredAction(ctx, "requiredAction", &keycloak.RequiredActionArgs{
//				RealmId: realm.Realm,
//				Alias:   pulumi.String("webauthn-register"),
//				Enabled: pulumi.Bool(true),
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
// Authentication executions can be imported using the formats`{{realm}}/{{alias}}`. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:index/requiredAction:RequiredAction required_action my-realm/my-default-action-alias
//
// ```
type RequiredAction struct {
	pulumi.CustomResourceState

	// The alias of the action to attach as a required action.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// When `true`, the required action is set as the default action for new users. Defaults to `false`.
	DefaultAction pulumi.BoolPtrOutput `pulumi:"defaultAction"`
	// When `false`, the required action is not enabled for new users. Defaults to `false`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The name of the required action.
	Name pulumi.StringOutput `pulumi:"name"`
	// The priority of the required action.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The realm the required action exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewRequiredAction registers a new resource with the given unique name, arguments, and options.
func NewRequiredAction(ctx *pulumi.Context,
	name string, args *RequiredActionArgs, opts ...pulumi.ResourceOption) (*RequiredAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource RequiredAction
	err := ctx.RegisterResource("keycloak:index/requiredAction:RequiredAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRequiredAction gets an existing RequiredAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRequiredAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RequiredActionState, opts ...pulumi.ResourceOption) (*RequiredAction, error) {
	var resource RequiredAction
	err := ctx.ReadResource("keycloak:index/requiredAction:RequiredAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RequiredAction resources.
type requiredActionState struct {
	// The alias of the action to attach as a required action.
	Alias *string `pulumi:"alias"`
	// When `true`, the required action is set as the default action for new users. Defaults to `false`.
	DefaultAction *bool `pulumi:"defaultAction"`
	// When `false`, the required action is not enabled for new users. Defaults to `false`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the required action.
	Name *string `pulumi:"name"`
	// The priority of the required action.
	Priority *int `pulumi:"priority"`
	// The realm the required action exists in.
	RealmId *string `pulumi:"realmId"`
}

type RequiredActionState struct {
	// The alias of the action to attach as a required action.
	Alias pulumi.StringPtrInput
	// When `true`, the required action is set as the default action for new users. Defaults to `false`.
	DefaultAction pulumi.BoolPtrInput
	// When `false`, the required action is not enabled for new users. Defaults to `false`.
	Enabled pulumi.BoolPtrInput
	// The name of the required action.
	Name pulumi.StringPtrInput
	// The priority of the required action.
	Priority pulumi.IntPtrInput
	// The realm the required action exists in.
	RealmId pulumi.StringPtrInput
}

func (RequiredActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*requiredActionState)(nil)).Elem()
}

type requiredActionArgs struct {
	// The alias of the action to attach as a required action.
	Alias string `pulumi:"alias"`
	// When `true`, the required action is set as the default action for new users. Defaults to `false`.
	DefaultAction *bool `pulumi:"defaultAction"`
	// When `false`, the required action is not enabled for new users. Defaults to `false`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the required action.
	Name *string `pulumi:"name"`
	// The priority of the required action.
	Priority *int `pulumi:"priority"`
	// The realm the required action exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a RequiredAction resource.
type RequiredActionArgs struct {
	// The alias of the action to attach as a required action.
	Alias pulumi.StringInput
	// When `true`, the required action is set as the default action for new users. Defaults to `false`.
	DefaultAction pulumi.BoolPtrInput
	// When `false`, the required action is not enabled for new users. Defaults to `false`.
	Enabled pulumi.BoolPtrInput
	// The name of the required action.
	Name pulumi.StringPtrInput
	// The priority of the required action.
	Priority pulumi.IntPtrInput
	// The realm the required action exists in.
	RealmId pulumi.StringInput
}

func (RequiredActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*requiredActionArgs)(nil)).Elem()
}

type RequiredActionInput interface {
	pulumi.Input

	ToRequiredActionOutput() RequiredActionOutput
	ToRequiredActionOutputWithContext(ctx context.Context) RequiredActionOutput
}

func (*RequiredAction) ElementType() reflect.Type {
	return reflect.TypeOf((**RequiredAction)(nil)).Elem()
}

func (i *RequiredAction) ToRequiredActionOutput() RequiredActionOutput {
	return i.ToRequiredActionOutputWithContext(context.Background())
}

func (i *RequiredAction) ToRequiredActionOutputWithContext(ctx context.Context) RequiredActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredActionOutput)
}

// RequiredActionArrayInput is an input type that accepts RequiredActionArray and RequiredActionArrayOutput values.
// You can construct a concrete instance of `RequiredActionArrayInput` via:
//
//	RequiredActionArray{ RequiredActionArgs{...} }
type RequiredActionArrayInput interface {
	pulumi.Input

	ToRequiredActionArrayOutput() RequiredActionArrayOutput
	ToRequiredActionArrayOutputWithContext(context.Context) RequiredActionArrayOutput
}

type RequiredActionArray []RequiredActionInput

func (RequiredActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RequiredAction)(nil)).Elem()
}

func (i RequiredActionArray) ToRequiredActionArrayOutput() RequiredActionArrayOutput {
	return i.ToRequiredActionArrayOutputWithContext(context.Background())
}

func (i RequiredActionArray) ToRequiredActionArrayOutputWithContext(ctx context.Context) RequiredActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredActionArrayOutput)
}

// RequiredActionMapInput is an input type that accepts RequiredActionMap and RequiredActionMapOutput values.
// You can construct a concrete instance of `RequiredActionMapInput` via:
//
//	RequiredActionMap{ "key": RequiredActionArgs{...} }
type RequiredActionMapInput interface {
	pulumi.Input

	ToRequiredActionMapOutput() RequiredActionMapOutput
	ToRequiredActionMapOutputWithContext(context.Context) RequiredActionMapOutput
}

type RequiredActionMap map[string]RequiredActionInput

func (RequiredActionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RequiredAction)(nil)).Elem()
}

func (i RequiredActionMap) ToRequiredActionMapOutput() RequiredActionMapOutput {
	return i.ToRequiredActionMapOutputWithContext(context.Background())
}

func (i RequiredActionMap) ToRequiredActionMapOutputWithContext(ctx context.Context) RequiredActionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredActionMapOutput)
}

type RequiredActionOutput struct{ *pulumi.OutputState }

func (RequiredActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequiredAction)(nil)).Elem()
}

func (o RequiredActionOutput) ToRequiredActionOutput() RequiredActionOutput {
	return o
}

func (o RequiredActionOutput) ToRequiredActionOutputWithContext(ctx context.Context) RequiredActionOutput {
	return o
}

// The alias of the action to attach as a required action.
func (o RequiredActionOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *RequiredAction) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// When `true`, the required action is set as the default action for new users. Defaults to `false`.
func (o RequiredActionOutput) DefaultAction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredAction) pulumi.BoolPtrOutput { return v.DefaultAction }).(pulumi.BoolPtrOutput)
}

// When `false`, the required action is not enabled for new users. Defaults to `false`.
func (o RequiredActionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredAction) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The name of the required action.
func (o RequiredActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RequiredAction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The priority of the required action.
func (o RequiredActionOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *RequiredAction) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// The realm the required action exists in.
func (o RequiredActionOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *RequiredAction) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type RequiredActionArrayOutput struct{ *pulumi.OutputState }

func (RequiredActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RequiredAction)(nil)).Elem()
}

func (o RequiredActionArrayOutput) ToRequiredActionArrayOutput() RequiredActionArrayOutput {
	return o
}

func (o RequiredActionArrayOutput) ToRequiredActionArrayOutputWithContext(ctx context.Context) RequiredActionArrayOutput {
	return o
}

func (o RequiredActionArrayOutput) Index(i pulumi.IntInput) RequiredActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RequiredAction {
		return vs[0].([]*RequiredAction)[vs[1].(int)]
	}).(RequiredActionOutput)
}

type RequiredActionMapOutput struct{ *pulumi.OutputState }

func (RequiredActionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RequiredAction)(nil)).Elem()
}

func (o RequiredActionMapOutput) ToRequiredActionMapOutput() RequiredActionMapOutput {
	return o
}

func (o RequiredActionMapOutput) ToRequiredActionMapOutputWithContext(ctx context.Context) RequiredActionMapOutput {
	return o
}

func (o RequiredActionMapOutput) MapIndex(k pulumi.StringInput) RequiredActionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RequiredAction {
		return vs[0].(map[string]*RequiredAction)[vs[1].(string)]
	}).(RequiredActionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RequiredActionInput)(nil)).Elem(), &RequiredAction{})
	pulumi.RegisterInputType(reflect.TypeOf((*RequiredActionArrayInput)(nil)).Elem(), RequiredActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RequiredActionMapInput)(nil)).Elem(), RequiredActionMap{})
	pulumi.RegisterOutputType(RequiredActionOutput{})
	pulumi.RegisterOutputType(RequiredActionArrayOutput{})
	pulumi.RegisterOutputType(RequiredActionMapOutput{})
}
