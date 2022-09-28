// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing `rsa-generated` Realm keystores within Keycloak.
//
// A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
//
// ## Import
//
// Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:index/realmKeystoreRsaGenerated:RealmKeystoreRsaGenerated keystore_rsa_generated my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
//
// ```
type RealmKeystoreRsaGenerated struct {
	pulumi.CustomResourceState

	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm pulumi.StringPtrOutput `pulumi:"algorithm"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Size for the generated keys
	KeySize pulumi.IntPtrOutput `pulumi:"keySize"`
	// Display name of provider when linked in admin console.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewRealmKeystoreRsaGenerated registers a new resource with the given unique name, arguments, and options.
func NewRealmKeystoreRsaGenerated(ctx *pulumi.Context,
	name string, args *RealmKeystoreRsaGeneratedArgs, opts ...pulumi.ResourceOption) (*RealmKeystoreRsaGenerated, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource RealmKeystoreRsaGenerated
	err := ctx.RegisterResource("keycloak:index/realmKeystoreRsaGenerated:RealmKeystoreRsaGenerated", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealmKeystoreRsaGenerated gets an existing RealmKeystoreRsaGenerated resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealmKeystoreRsaGenerated(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmKeystoreRsaGeneratedState, opts ...pulumi.ResourceOption) (*RealmKeystoreRsaGenerated, error) {
	var resource RealmKeystoreRsaGenerated
	err := ctx.ReadResource("keycloak:index/realmKeystoreRsaGenerated:RealmKeystoreRsaGenerated", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealmKeystoreRsaGenerated resources.
type realmKeystoreRsaGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm *string `pulumi:"algorithm"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Size for the generated keys
	KeySize *int `pulumi:"keySize"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId *string `pulumi:"realmId"`
}

type RealmKeystoreRsaGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm pulumi.StringPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Size for the generated keys
	KeySize pulumi.IntPtrInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringPtrInput
}

func (RealmKeystoreRsaGeneratedState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreRsaGeneratedState)(nil)).Elem()
}

type realmKeystoreRsaGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm *string `pulumi:"algorithm"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Size for the generated keys
	KeySize *int `pulumi:"keySize"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a RealmKeystoreRsaGenerated resource.
type RealmKeystoreRsaGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm pulumi.StringPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Size for the generated keys
	KeySize pulumi.IntPtrInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringInput
}

func (RealmKeystoreRsaGeneratedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreRsaGeneratedArgs)(nil)).Elem()
}

type RealmKeystoreRsaGeneratedInput interface {
	pulumi.Input

	ToRealmKeystoreRsaGeneratedOutput() RealmKeystoreRsaGeneratedOutput
	ToRealmKeystoreRsaGeneratedOutputWithContext(ctx context.Context) RealmKeystoreRsaGeneratedOutput
}

func (*RealmKeystoreRsaGenerated) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreRsaGenerated)(nil)).Elem()
}

func (i *RealmKeystoreRsaGenerated) ToRealmKeystoreRsaGeneratedOutput() RealmKeystoreRsaGeneratedOutput {
	return i.ToRealmKeystoreRsaGeneratedOutputWithContext(context.Background())
}

func (i *RealmKeystoreRsaGenerated) ToRealmKeystoreRsaGeneratedOutputWithContext(ctx context.Context) RealmKeystoreRsaGeneratedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreRsaGeneratedOutput)
}

// RealmKeystoreRsaGeneratedArrayInput is an input type that accepts RealmKeystoreRsaGeneratedArray and RealmKeystoreRsaGeneratedArrayOutput values.
// You can construct a concrete instance of `RealmKeystoreRsaGeneratedArrayInput` via:
//
//	RealmKeystoreRsaGeneratedArray{ RealmKeystoreRsaGeneratedArgs{...} }
type RealmKeystoreRsaGeneratedArrayInput interface {
	pulumi.Input

	ToRealmKeystoreRsaGeneratedArrayOutput() RealmKeystoreRsaGeneratedArrayOutput
	ToRealmKeystoreRsaGeneratedArrayOutputWithContext(context.Context) RealmKeystoreRsaGeneratedArrayOutput
}

type RealmKeystoreRsaGeneratedArray []RealmKeystoreRsaGeneratedInput

func (RealmKeystoreRsaGeneratedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreRsaGenerated)(nil)).Elem()
}

func (i RealmKeystoreRsaGeneratedArray) ToRealmKeystoreRsaGeneratedArrayOutput() RealmKeystoreRsaGeneratedArrayOutput {
	return i.ToRealmKeystoreRsaGeneratedArrayOutputWithContext(context.Background())
}

func (i RealmKeystoreRsaGeneratedArray) ToRealmKeystoreRsaGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreRsaGeneratedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreRsaGeneratedArrayOutput)
}

// RealmKeystoreRsaGeneratedMapInput is an input type that accepts RealmKeystoreRsaGeneratedMap and RealmKeystoreRsaGeneratedMapOutput values.
// You can construct a concrete instance of `RealmKeystoreRsaGeneratedMapInput` via:
//
//	RealmKeystoreRsaGeneratedMap{ "key": RealmKeystoreRsaGeneratedArgs{...} }
type RealmKeystoreRsaGeneratedMapInput interface {
	pulumi.Input

	ToRealmKeystoreRsaGeneratedMapOutput() RealmKeystoreRsaGeneratedMapOutput
	ToRealmKeystoreRsaGeneratedMapOutputWithContext(context.Context) RealmKeystoreRsaGeneratedMapOutput
}

type RealmKeystoreRsaGeneratedMap map[string]RealmKeystoreRsaGeneratedInput

func (RealmKeystoreRsaGeneratedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreRsaGenerated)(nil)).Elem()
}

func (i RealmKeystoreRsaGeneratedMap) ToRealmKeystoreRsaGeneratedMapOutput() RealmKeystoreRsaGeneratedMapOutput {
	return i.ToRealmKeystoreRsaGeneratedMapOutputWithContext(context.Background())
}

func (i RealmKeystoreRsaGeneratedMap) ToRealmKeystoreRsaGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreRsaGeneratedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreRsaGeneratedMapOutput)
}

type RealmKeystoreRsaGeneratedOutput struct{ *pulumi.OutputState }

func (RealmKeystoreRsaGeneratedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreRsaGenerated)(nil)).Elem()
}

func (o RealmKeystoreRsaGeneratedOutput) ToRealmKeystoreRsaGeneratedOutput() RealmKeystoreRsaGeneratedOutput {
	return o
}

func (o RealmKeystoreRsaGeneratedOutput) ToRealmKeystoreRsaGeneratedOutputWithContext(ctx context.Context) RealmKeystoreRsaGeneratedOutput {
	return o
}

// When `false`, key in not used for signing. Defaults to `true`.
func (o RealmKeystoreRsaGeneratedOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreRsaGenerated) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// Intended algorithm for the key. Defaults to `RS256`
func (o RealmKeystoreRsaGeneratedOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreRsaGenerated) pulumi.StringPtrOutput { return v.Algorithm }).(pulumi.StringPtrOutput)
}

// When `false`, key is not accessible in this realm. Defaults to `true`.
func (o RealmKeystoreRsaGeneratedOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreRsaGenerated) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Size for the generated keys
func (o RealmKeystoreRsaGeneratedOutput) KeySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreRsaGenerated) pulumi.IntPtrOutput { return v.KeySize }).(pulumi.IntPtrOutput)
}

// Display name of provider when linked in admin console.
func (o RealmKeystoreRsaGeneratedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmKeystoreRsaGenerated) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Priority for the provider. Defaults to `0`
func (o RealmKeystoreRsaGeneratedOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreRsaGenerated) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The realm this keystore exists in.
func (o RealmKeystoreRsaGeneratedOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmKeystoreRsaGenerated) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type RealmKeystoreRsaGeneratedArrayOutput struct{ *pulumi.OutputState }

func (RealmKeystoreRsaGeneratedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreRsaGenerated)(nil)).Elem()
}

func (o RealmKeystoreRsaGeneratedArrayOutput) ToRealmKeystoreRsaGeneratedArrayOutput() RealmKeystoreRsaGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreRsaGeneratedArrayOutput) ToRealmKeystoreRsaGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreRsaGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreRsaGeneratedArrayOutput) Index(i pulumi.IntInput) RealmKeystoreRsaGeneratedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealmKeystoreRsaGenerated {
		return vs[0].([]*RealmKeystoreRsaGenerated)[vs[1].(int)]
	}).(RealmKeystoreRsaGeneratedOutput)
}

type RealmKeystoreRsaGeneratedMapOutput struct{ *pulumi.OutputState }

func (RealmKeystoreRsaGeneratedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreRsaGenerated)(nil)).Elem()
}

func (o RealmKeystoreRsaGeneratedMapOutput) ToRealmKeystoreRsaGeneratedMapOutput() RealmKeystoreRsaGeneratedMapOutput {
	return o
}

func (o RealmKeystoreRsaGeneratedMapOutput) ToRealmKeystoreRsaGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreRsaGeneratedMapOutput {
	return o
}

func (o RealmKeystoreRsaGeneratedMapOutput) MapIndex(k pulumi.StringInput) RealmKeystoreRsaGeneratedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealmKeystoreRsaGenerated {
		return vs[0].(map[string]*RealmKeystoreRsaGenerated)[vs[1].(string)]
	}).(RealmKeystoreRsaGeneratedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreRsaGeneratedInput)(nil)).Elem(), &RealmKeystoreRsaGenerated{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreRsaGeneratedArrayInput)(nil)).Elem(), RealmKeystoreRsaGeneratedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreRsaGeneratedMapInput)(nil)).Elem(), RealmKeystoreRsaGeneratedMap{})
	pulumi.RegisterOutputType(RealmKeystoreRsaGeneratedOutput{})
	pulumi.RegisterOutputType(RealmKeystoreRsaGeneratedArrayOutput{})
	pulumi.RegisterOutputType(RealmKeystoreRsaGeneratedMapOutput{})
}
