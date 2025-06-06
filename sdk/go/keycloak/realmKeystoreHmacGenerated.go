// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing `hmac-generated` Realm keystores within Keycloak.
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
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak"
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
//			_, err = keycloak.NewRealmKeystoreHmacGenerated(ctx, "keystore_hmac_generated", &keycloak.RealmKeystoreHmacGeneratedArgs{
//				Name:       pulumi.String("my-hmac-generated-key"),
//				RealmId:    realm.ID(),
//				Enabled:    pulumi.Bool(true),
//				Active:     pulumi.Bool(true),
//				Priority:   pulumi.Int(100),
//				Algorithm:  pulumi.String("HS256"),
//				SecretSize: pulumi.Int(64),
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
// Realm keys can be imported using realm name and keystore id, you can find it in web UI.
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:index/realmKeystoreHmacGenerated:RealmKeystoreHmacGenerated keystore_hmac_generated my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
// ```
type RealmKeystoreHmacGenerated struct {
	pulumi.CustomResourceState

	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Intended algorithm for the key. Defaults to `HS256`
	Algorithm pulumi.StringPtrOutput `pulumi:"algorithm"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Display name of provider when linked in admin console.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// Size in bytes for the generated secret. Defaults to `64`.
	SecretSize pulumi.IntPtrOutput `pulumi:"secretSize"`
}

// NewRealmKeystoreHmacGenerated registers a new resource with the given unique name, arguments, and options.
func NewRealmKeystoreHmacGenerated(ctx *pulumi.Context,
	name string, args *RealmKeystoreHmacGeneratedArgs, opts ...pulumi.ResourceOption) (*RealmKeystoreHmacGenerated, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RealmKeystoreHmacGenerated
	err := ctx.RegisterResource("keycloak:index/realmKeystoreHmacGenerated:RealmKeystoreHmacGenerated", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealmKeystoreHmacGenerated gets an existing RealmKeystoreHmacGenerated resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealmKeystoreHmacGenerated(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmKeystoreHmacGeneratedState, opts ...pulumi.ResourceOption) (*RealmKeystoreHmacGenerated, error) {
	var resource RealmKeystoreHmacGenerated
	err := ctx.ReadResource("keycloak:index/realmKeystoreHmacGenerated:RealmKeystoreHmacGenerated", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealmKeystoreHmacGenerated resources.
type realmKeystoreHmacGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Intended algorithm for the key. Defaults to `HS256`
	Algorithm *string `pulumi:"algorithm"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId *string `pulumi:"realmId"`
	// Size in bytes for the generated secret. Defaults to `64`.
	SecretSize *int `pulumi:"secretSize"`
}

type RealmKeystoreHmacGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Intended algorithm for the key. Defaults to `HS256`
	Algorithm pulumi.StringPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringPtrInput
	// Size in bytes for the generated secret. Defaults to `64`.
	SecretSize pulumi.IntPtrInput
}

func (RealmKeystoreHmacGeneratedState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreHmacGeneratedState)(nil)).Elem()
}

type realmKeystoreHmacGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Intended algorithm for the key. Defaults to `HS256`
	Algorithm *string `pulumi:"algorithm"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId string `pulumi:"realmId"`
	// Size in bytes for the generated secret. Defaults to `64`.
	SecretSize *int `pulumi:"secretSize"`
}

// The set of arguments for constructing a RealmKeystoreHmacGenerated resource.
type RealmKeystoreHmacGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Intended algorithm for the key. Defaults to `HS256`
	Algorithm pulumi.StringPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringInput
	// Size in bytes for the generated secret. Defaults to `64`.
	SecretSize pulumi.IntPtrInput
}

func (RealmKeystoreHmacGeneratedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreHmacGeneratedArgs)(nil)).Elem()
}

type RealmKeystoreHmacGeneratedInput interface {
	pulumi.Input

	ToRealmKeystoreHmacGeneratedOutput() RealmKeystoreHmacGeneratedOutput
	ToRealmKeystoreHmacGeneratedOutputWithContext(ctx context.Context) RealmKeystoreHmacGeneratedOutput
}

func (*RealmKeystoreHmacGenerated) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreHmacGenerated)(nil)).Elem()
}

func (i *RealmKeystoreHmacGenerated) ToRealmKeystoreHmacGeneratedOutput() RealmKeystoreHmacGeneratedOutput {
	return i.ToRealmKeystoreHmacGeneratedOutputWithContext(context.Background())
}

func (i *RealmKeystoreHmacGenerated) ToRealmKeystoreHmacGeneratedOutputWithContext(ctx context.Context) RealmKeystoreHmacGeneratedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreHmacGeneratedOutput)
}

// RealmKeystoreHmacGeneratedArrayInput is an input type that accepts RealmKeystoreHmacGeneratedArray and RealmKeystoreHmacGeneratedArrayOutput values.
// You can construct a concrete instance of `RealmKeystoreHmacGeneratedArrayInput` via:
//
//	RealmKeystoreHmacGeneratedArray{ RealmKeystoreHmacGeneratedArgs{...} }
type RealmKeystoreHmacGeneratedArrayInput interface {
	pulumi.Input

	ToRealmKeystoreHmacGeneratedArrayOutput() RealmKeystoreHmacGeneratedArrayOutput
	ToRealmKeystoreHmacGeneratedArrayOutputWithContext(context.Context) RealmKeystoreHmacGeneratedArrayOutput
}

type RealmKeystoreHmacGeneratedArray []RealmKeystoreHmacGeneratedInput

func (RealmKeystoreHmacGeneratedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreHmacGenerated)(nil)).Elem()
}

func (i RealmKeystoreHmacGeneratedArray) ToRealmKeystoreHmacGeneratedArrayOutput() RealmKeystoreHmacGeneratedArrayOutput {
	return i.ToRealmKeystoreHmacGeneratedArrayOutputWithContext(context.Background())
}

func (i RealmKeystoreHmacGeneratedArray) ToRealmKeystoreHmacGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreHmacGeneratedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreHmacGeneratedArrayOutput)
}

// RealmKeystoreHmacGeneratedMapInput is an input type that accepts RealmKeystoreHmacGeneratedMap and RealmKeystoreHmacGeneratedMapOutput values.
// You can construct a concrete instance of `RealmKeystoreHmacGeneratedMapInput` via:
//
//	RealmKeystoreHmacGeneratedMap{ "key": RealmKeystoreHmacGeneratedArgs{...} }
type RealmKeystoreHmacGeneratedMapInput interface {
	pulumi.Input

	ToRealmKeystoreHmacGeneratedMapOutput() RealmKeystoreHmacGeneratedMapOutput
	ToRealmKeystoreHmacGeneratedMapOutputWithContext(context.Context) RealmKeystoreHmacGeneratedMapOutput
}

type RealmKeystoreHmacGeneratedMap map[string]RealmKeystoreHmacGeneratedInput

func (RealmKeystoreHmacGeneratedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreHmacGenerated)(nil)).Elem()
}

func (i RealmKeystoreHmacGeneratedMap) ToRealmKeystoreHmacGeneratedMapOutput() RealmKeystoreHmacGeneratedMapOutput {
	return i.ToRealmKeystoreHmacGeneratedMapOutputWithContext(context.Background())
}

func (i RealmKeystoreHmacGeneratedMap) ToRealmKeystoreHmacGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreHmacGeneratedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreHmacGeneratedMapOutput)
}

type RealmKeystoreHmacGeneratedOutput struct{ *pulumi.OutputState }

func (RealmKeystoreHmacGeneratedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreHmacGenerated)(nil)).Elem()
}

func (o RealmKeystoreHmacGeneratedOutput) ToRealmKeystoreHmacGeneratedOutput() RealmKeystoreHmacGeneratedOutput {
	return o
}

func (o RealmKeystoreHmacGeneratedOutput) ToRealmKeystoreHmacGeneratedOutputWithContext(ctx context.Context) RealmKeystoreHmacGeneratedOutput {
	return o
}

// When `false`, key in not used for signing. Defaults to `true`.
func (o RealmKeystoreHmacGeneratedOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreHmacGenerated) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// Intended algorithm for the key. Defaults to `HS256`
func (o RealmKeystoreHmacGeneratedOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreHmacGenerated) pulumi.StringPtrOutput { return v.Algorithm }).(pulumi.StringPtrOutput)
}

// When `false`, key is not accessible in this realm. Defaults to `true`.
func (o RealmKeystoreHmacGeneratedOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreHmacGenerated) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Display name of provider when linked in admin console.
func (o RealmKeystoreHmacGeneratedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmKeystoreHmacGenerated) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Priority for the provider. Defaults to `0`
func (o RealmKeystoreHmacGeneratedOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreHmacGenerated) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The realm this keystore exists in.
func (o RealmKeystoreHmacGeneratedOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmKeystoreHmacGenerated) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// Size in bytes for the generated secret. Defaults to `64`.
func (o RealmKeystoreHmacGeneratedOutput) SecretSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreHmacGenerated) pulumi.IntPtrOutput { return v.SecretSize }).(pulumi.IntPtrOutput)
}

type RealmKeystoreHmacGeneratedArrayOutput struct{ *pulumi.OutputState }

func (RealmKeystoreHmacGeneratedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreHmacGenerated)(nil)).Elem()
}

func (o RealmKeystoreHmacGeneratedArrayOutput) ToRealmKeystoreHmacGeneratedArrayOutput() RealmKeystoreHmacGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreHmacGeneratedArrayOutput) ToRealmKeystoreHmacGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreHmacGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreHmacGeneratedArrayOutput) Index(i pulumi.IntInput) RealmKeystoreHmacGeneratedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealmKeystoreHmacGenerated {
		return vs[0].([]*RealmKeystoreHmacGenerated)[vs[1].(int)]
	}).(RealmKeystoreHmacGeneratedOutput)
}

type RealmKeystoreHmacGeneratedMapOutput struct{ *pulumi.OutputState }

func (RealmKeystoreHmacGeneratedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreHmacGenerated)(nil)).Elem()
}

func (o RealmKeystoreHmacGeneratedMapOutput) ToRealmKeystoreHmacGeneratedMapOutput() RealmKeystoreHmacGeneratedMapOutput {
	return o
}

func (o RealmKeystoreHmacGeneratedMapOutput) ToRealmKeystoreHmacGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreHmacGeneratedMapOutput {
	return o
}

func (o RealmKeystoreHmacGeneratedMapOutput) MapIndex(k pulumi.StringInput) RealmKeystoreHmacGeneratedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealmKeystoreHmacGenerated {
		return vs[0].(map[string]*RealmKeystoreHmacGenerated)[vs[1].(string)]
	}).(RealmKeystoreHmacGeneratedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreHmacGeneratedInput)(nil)).Elem(), &RealmKeystoreHmacGenerated{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreHmacGeneratedArrayInput)(nil)).Elem(), RealmKeystoreHmacGeneratedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreHmacGeneratedMapInput)(nil)).Elem(), RealmKeystoreHmacGeneratedMap{})
	pulumi.RegisterOutputType(RealmKeystoreHmacGeneratedOutput{})
	pulumi.RegisterOutputType(RealmKeystoreHmacGeneratedArrayOutput{})
	pulumi.RegisterOutputType(RealmKeystoreHmacGeneratedMapOutput{})
}
