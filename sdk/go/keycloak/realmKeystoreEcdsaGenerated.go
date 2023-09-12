// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Allows for creating and managing `acdsaGenerated` Realm keystores within Keycloak.
//
// A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
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
//				Realm: pulumi.String("my-realm"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRealmKeystoreEcdsaGenerated(ctx, "keystoreEcdsaGenerated", &keycloak.RealmKeystoreEcdsaGeneratedArgs{
//				RealmId:          realm.ID(),
//				Enabled:          pulumi.Bool(true),
//				Active:           pulumi.Bool(true),
//				Priority:         pulumi.Int(100),
//				EllipticCurveKey: pulumi.String("P-256"),
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
// Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:index/realmKeystoreEcdsaGenerated:RealmKeystoreEcdsaGenerated keystore_ecdsa_generated my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
//
// ```
type RealmKeystoreEcdsaGenerated struct {
	pulumi.CustomResourceState

	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Elliptic Curve used in ECDSA. Defaults to `P-256`.
	EllipticCurveKey pulumi.StringPtrOutput `pulumi:"ellipticCurveKey"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Display name of provider when linked in admin console.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewRealmKeystoreEcdsaGenerated registers a new resource with the given unique name, arguments, and options.
func NewRealmKeystoreEcdsaGenerated(ctx *pulumi.Context,
	name string, args *RealmKeystoreEcdsaGeneratedArgs, opts ...pulumi.ResourceOption) (*RealmKeystoreEcdsaGenerated, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RealmKeystoreEcdsaGenerated
	err := ctx.RegisterResource("keycloak:index/realmKeystoreEcdsaGenerated:RealmKeystoreEcdsaGenerated", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealmKeystoreEcdsaGenerated gets an existing RealmKeystoreEcdsaGenerated resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealmKeystoreEcdsaGenerated(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmKeystoreEcdsaGeneratedState, opts ...pulumi.ResourceOption) (*RealmKeystoreEcdsaGenerated, error) {
	var resource RealmKeystoreEcdsaGenerated
	err := ctx.ReadResource("keycloak:index/realmKeystoreEcdsaGenerated:RealmKeystoreEcdsaGenerated", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealmKeystoreEcdsaGenerated resources.
type realmKeystoreEcdsaGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Elliptic Curve used in ECDSA. Defaults to `P-256`.
	EllipticCurveKey *string `pulumi:"ellipticCurveKey"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId *string `pulumi:"realmId"`
}

type RealmKeystoreEcdsaGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Elliptic Curve used in ECDSA. Defaults to `P-256`.
	EllipticCurveKey pulumi.StringPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringPtrInput
}

func (RealmKeystoreEcdsaGeneratedState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreEcdsaGeneratedState)(nil)).Elem()
}

type realmKeystoreEcdsaGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Elliptic Curve used in ECDSA. Defaults to `P-256`.
	EllipticCurveKey *string `pulumi:"ellipticCurveKey"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a RealmKeystoreEcdsaGenerated resource.
type RealmKeystoreEcdsaGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Elliptic Curve used in ECDSA. Defaults to `P-256`.
	EllipticCurveKey pulumi.StringPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringInput
}

func (RealmKeystoreEcdsaGeneratedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreEcdsaGeneratedArgs)(nil)).Elem()
}

type RealmKeystoreEcdsaGeneratedInput interface {
	pulumi.Input

	ToRealmKeystoreEcdsaGeneratedOutput() RealmKeystoreEcdsaGeneratedOutput
	ToRealmKeystoreEcdsaGeneratedOutputWithContext(ctx context.Context) RealmKeystoreEcdsaGeneratedOutput
}

func (*RealmKeystoreEcdsaGenerated) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreEcdsaGenerated)(nil)).Elem()
}

func (i *RealmKeystoreEcdsaGenerated) ToRealmKeystoreEcdsaGeneratedOutput() RealmKeystoreEcdsaGeneratedOutput {
	return i.ToRealmKeystoreEcdsaGeneratedOutputWithContext(context.Background())
}

func (i *RealmKeystoreEcdsaGenerated) ToRealmKeystoreEcdsaGeneratedOutputWithContext(ctx context.Context) RealmKeystoreEcdsaGeneratedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreEcdsaGeneratedOutput)
}

func (i *RealmKeystoreEcdsaGenerated) ToOutput(ctx context.Context) pulumix.Output[*RealmKeystoreEcdsaGenerated] {
	return pulumix.Output[*RealmKeystoreEcdsaGenerated]{
		OutputState: i.ToRealmKeystoreEcdsaGeneratedOutputWithContext(ctx).OutputState,
	}
}

// RealmKeystoreEcdsaGeneratedArrayInput is an input type that accepts RealmKeystoreEcdsaGeneratedArray and RealmKeystoreEcdsaGeneratedArrayOutput values.
// You can construct a concrete instance of `RealmKeystoreEcdsaGeneratedArrayInput` via:
//
//	RealmKeystoreEcdsaGeneratedArray{ RealmKeystoreEcdsaGeneratedArgs{...} }
type RealmKeystoreEcdsaGeneratedArrayInput interface {
	pulumi.Input

	ToRealmKeystoreEcdsaGeneratedArrayOutput() RealmKeystoreEcdsaGeneratedArrayOutput
	ToRealmKeystoreEcdsaGeneratedArrayOutputWithContext(context.Context) RealmKeystoreEcdsaGeneratedArrayOutput
}

type RealmKeystoreEcdsaGeneratedArray []RealmKeystoreEcdsaGeneratedInput

func (RealmKeystoreEcdsaGeneratedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreEcdsaGenerated)(nil)).Elem()
}

func (i RealmKeystoreEcdsaGeneratedArray) ToRealmKeystoreEcdsaGeneratedArrayOutput() RealmKeystoreEcdsaGeneratedArrayOutput {
	return i.ToRealmKeystoreEcdsaGeneratedArrayOutputWithContext(context.Background())
}

func (i RealmKeystoreEcdsaGeneratedArray) ToRealmKeystoreEcdsaGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreEcdsaGeneratedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreEcdsaGeneratedArrayOutput)
}

func (i RealmKeystoreEcdsaGeneratedArray) ToOutput(ctx context.Context) pulumix.Output[[]*RealmKeystoreEcdsaGenerated] {
	return pulumix.Output[[]*RealmKeystoreEcdsaGenerated]{
		OutputState: i.ToRealmKeystoreEcdsaGeneratedArrayOutputWithContext(ctx).OutputState,
	}
}

// RealmKeystoreEcdsaGeneratedMapInput is an input type that accepts RealmKeystoreEcdsaGeneratedMap and RealmKeystoreEcdsaGeneratedMapOutput values.
// You can construct a concrete instance of `RealmKeystoreEcdsaGeneratedMapInput` via:
//
//	RealmKeystoreEcdsaGeneratedMap{ "key": RealmKeystoreEcdsaGeneratedArgs{...} }
type RealmKeystoreEcdsaGeneratedMapInput interface {
	pulumi.Input

	ToRealmKeystoreEcdsaGeneratedMapOutput() RealmKeystoreEcdsaGeneratedMapOutput
	ToRealmKeystoreEcdsaGeneratedMapOutputWithContext(context.Context) RealmKeystoreEcdsaGeneratedMapOutput
}

type RealmKeystoreEcdsaGeneratedMap map[string]RealmKeystoreEcdsaGeneratedInput

func (RealmKeystoreEcdsaGeneratedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreEcdsaGenerated)(nil)).Elem()
}

func (i RealmKeystoreEcdsaGeneratedMap) ToRealmKeystoreEcdsaGeneratedMapOutput() RealmKeystoreEcdsaGeneratedMapOutput {
	return i.ToRealmKeystoreEcdsaGeneratedMapOutputWithContext(context.Background())
}

func (i RealmKeystoreEcdsaGeneratedMap) ToRealmKeystoreEcdsaGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreEcdsaGeneratedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreEcdsaGeneratedMapOutput)
}

func (i RealmKeystoreEcdsaGeneratedMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RealmKeystoreEcdsaGenerated] {
	return pulumix.Output[map[string]*RealmKeystoreEcdsaGenerated]{
		OutputState: i.ToRealmKeystoreEcdsaGeneratedMapOutputWithContext(ctx).OutputState,
	}
}

type RealmKeystoreEcdsaGeneratedOutput struct{ *pulumi.OutputState }

func (RealmKeystoreEcdsaGeneratedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreEcdsaGenerated)(nil)).Elem()
}

func (o RealmKeystoreEcdsaGeneratedOutput) ToRealmKeystoreEcdsaGeneratedOutput() RealmKeystoreEcdsaGeneratedOutput {
	return o
}

func (o RealmKeystoreEcdsaGeneratedOutput) ToRealmKeystoreEcdsaGeneratedOutputWithContext(ctx context.Context) RealmKeystoreEcdsaGeneratedOutput {
	return o
}

func (o RealmKeystoreEcdsaGeneratedOutput) ToOutput(ctx context.Context) pulumix.Output[*RealmKeystoreEcdsaGenerated] {
	return pulumix.Output[*RealmKeystoreEcdsaGenerated]{
		OutputState: o.OutputState,
	}
}

// When `false`, key in not used for signing. Defaults to `true`.
func (o RealmKeystoreEcdsaGeneratedOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreEcdsaGenerated) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// Elliptic Curve used in ECDSA. Defaults to `P-256`.
func (o RealmKeystoreEcdsaGeneratedOutput) EllipticCurveKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreEcdsaGenerated) pulumi.StringPtrOutput { return v.EllipticCurveKey }).(pulumi.StringPtrOutput)
}

// When `false`, key is not accessible in this realm. Defaults to `true`.
func (o RealmKeystoreEcdsaGeneratedOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreEcdsaGenerated) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Display name of provider when linked in admin console.
func (o RealmKeystoreEcdsaGeneratedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmKeystoreEcdsaGenerated) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Priority for the provider. Defaults to `0`
func (o RealmKeystoreEcdsaGeneratedOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreEcdsaGenerated) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The realm this keystore exists in.
func (o RealmKeystoreEcdsaGeneratedOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmKeystoreEcdsaGenerated) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type RealmKeystoreEcdsaGeneratedArrayOutput struct{ *pulumi.OutputState }

func (RealmKeystoreEcdsaGeneratedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreEcdsaGenerated)(nil)).Elem()
}

func (o RealmKeystoreEcdsaGeneratedArrayOutput) ToRealmKeystoreEcdsaGeneratedArrayOutput() RealmKeystoreEcdsaGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreEcdsaGeneratedArrayOutput) ToRealmKeystoreEcdsaGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreEcdsaGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreEcdsaGeneratedArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RealmKeystoreEcdsaGenerated] {
	return pulumix.Output[[]*RealmKeystoreEcdsaGenerated]{
		OutputState: o.OutputState,
	}
}

func (o RealmKeystoreEcdsaGeneratedArrayOutput) Index(i pulumi.IntInput) RealmKeystoreEcdsaGeneratedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealmKeystoreEcdsaGenerated {
		return vs[0].([]*RealmKeystoreEcdsaGenerated)[vs[1].(int)]
	}).(RealmKeystoreEcdsaGeneratedOutput)
}

type RealmKeystoreEcdsaGeneratedMapOutput struct{ *pulumi.OutputState }

func (RealmKeystoreEcdsaGeneratedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreEcdsaGenerated)(nil)).Elem()
}

func (o RealmKeystoreEcdsaGeneratedMapOutput) ToRealmKeystoreEcdsaGeneratedMapOutput() RealmKeystoreEcdsaGeneratedMapOutput {
	return o
}

func (o RealmKeystoreEcdsaGeneratedMapOutput) ToRealmKeystoreEcdsaGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreEcdsaGeneratedMapOutput {
	return o
}

func (o RealmKeystoreEcdsaGeneratedMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RealmKeystoreEcdsaGenerated] {
	return pulumix.Output[map[string]*RealmKeystoreEcdsaGenerated]{
		OutputState: o.OutputState,
	}
}

func (o RealmKeystoreEcdsaGeneratedMapOutput) MapIndex(k pulumi.StringInput) RealmKeystoreEcdsaGeneratedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealmKeystoreEcdsaGenerated {
		return vs[0].(map[string]*RealmKeystoreEcdsaGenerated)[vs[1].(string)]
	}).(RealmKeystoreEcdsaGeneratedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreEcdsaGeneratedInput)(nil)).Elem(), &RealmKeystoreEcdsaGenerated{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreEcdsaGeneratedArrayInput)(nil)).Elem(), RealmKeystoreEcdsaGeneratedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreEcdsaGeneratedMapInput)(nil)).Elem(), RealmKeystoreEcdsaGeneratedMap{})
	pulumi.RegisterOutputType(RealmKeystoreEcdsaGeneratedOutput{})
	pulumi.RegisterOutputType(RealmKeystoreEcdsaGeneratedArrayOutput{})
	pulumi.RegisterOutputType(RealmKeystoreEcdsaGeneratedMapOutput{})
}
