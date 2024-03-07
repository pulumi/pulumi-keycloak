// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing `aes-generated` Realm keystores within Keycloak.
//
// A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err = keycloak.NewRealmKeystoreAesGenerated(ctx, "keystoreAesGenerated", &keycloak.RealmKeystoreAesGeneratedArgs{
//				RealmId:    realm.ID(),
//				Enabled:    pulumi.Bool(true),
//				Active:     pulumi.Bool(true),
//				Priority:   pulumi.Int(100),
//				SecretSize: pulumi.Int(16),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
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
// $ pulumi import keycloak:index/realmKeystoreAesGenerated:RealmKeystoreAesGenerated keystore_aes_generated my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
// ```
type RealmKeystoreAesGenerated struct {
	pulumi.CustomResourceState

	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Display name of provider when linked in admin console.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
	SecretSize pulumi.IntPtrOutput `pulumi:"secretSize"`
}

// NewRealmKeystoreAesGenerated registers a new resource with the given unique name, arguments, and options.
func NewRealmKeystoreAesGenerated(ctx *pulumi.Context,
	name string, args *RealmKeystoreAesGeneratedArgs, opts ...pulumi.ResourceOption) (*RealmKeystoreAesGenerated, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RealmKeystoreAesGenerated
	err := ctx.RegisterResource("keycloak:index/realmKeystoreAesGenerated:RealmKeystoreAesGenerated", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealmKeystoreAesGenerated gets an existing RealmKeystoreAesGenerated resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealmKeystoreAesGenerated(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmKeystoreAesGeneratedState, opts ...pulumi.ResourceOption) (*RealmKeystoreAesGenerated, error) {
	var resource RealmKeystoreAesGenerated
	err := ctx.ReadResource("keycloak:index/realmKeystoreAesGenerated:RealmKeystoreAesGenerated", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealmKeystoreAesGenerated resources.
type realmKeystoreAesGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId *string `pulumi:"realmId"`
	// Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
	SecretSize *int `pulumi:"secretSize"`
}

type RealmKeystoreAesGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringPtrInput
	// Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
	SecretSize pulumi.IntPtrInput
}

func (RealmKeystoreAesGeneratedState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreAesGeneratedState)(nil)).Elem()
}

type realmKeystoreAesGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId string `pulumi:"realmId"`
	// Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
	SecretSize *int `pulumi:"secretSize"`
}

// The set of arguments for constructing a RealmKeystoreAesGenerated resource.
type RealmKeystoreAesGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringInput
	// Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
	SecretSize pulumi.IntPtrInput
}

func (RealmKeystoreAesGeneratedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreAesGeneratedArgs)(nil)).Elem()
}

type RealmKeystoreAesGeneratedInput interface {
	pulumi.Input

	ToRealmKeystoreAesGeneratedOutput() RealmKeystoreAesGeneratedOutput
	ToRealmKeystoreAesGeneratedOutputWithContext(ctx context.Context) RealmKeystoreAesGeneratedOutput
}

func (*RealmKeystoreAesGenerated) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreAesGenerated)(nil)).Elem()
}

func (i *RealmKeystoreAesGenerated) ToRealmKeystoreAesGeneratedOutput() RealmKeystoreAesGeneratedOutput {
	return i.ToRealmKeystoreAesGeneratedOutputWithContext(context.Background())
}

func (i *RealmKeystoreAesGenerated) ToRealmKeystoreAesGeneratedOutputWithContext(ctx context.Context) RealmKeystoreAesGeneratedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreAesGeneratedOutput)
}

// RealmKeystoreAesGeneratedArrayInput is an input type that accepts RealmKeystoreAesGeneratedArray and RealmKeystoreAesGeneratedArrayOutput values.
// You can construct a concrete instance of `RealmKeystoreAesGeneratedArrayInput` via:
//
//	RealmKeystoreAesGeneratedArray{ RealmKeystoreAesGeneratedArgs{...} }
type RealmKeystoreAesGeneratedArrayInput interface {
	pulumi.Input

	ToRealmKeystoreAesGeneratedArrayOutput() RealmKeystoreAesGeneratedArrayOutput
	ToRealmKeystoreAesGeneratedArrayOutputWithContext(context.Context) RealmKeystoreAesGeneratedArrayOutput
}

type RealmKeystoreAesGeneratedArray []RealmKeystoreAesGeneratedInput

func (RealmKeystoreAesGeneratedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreAesGenerated)(nil)).Elem()
}

func (i RealmKeystoreAesGeneratedArray) ToRealmKeystoreAesGeneratedArrayOutput() RealmKeystoreAesGeneratedArrayOutput {
	return i.ToRealmKeystoreAesGeneratedArrayOutputWithContext(context.Background())
}

func (i RealmKeystoreAesGeneratedArray) ToRealmKeystoreAesGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreAesGeneratedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreAesGeneratedArrayOutput)
}

// RealmKeystoreAesGeneratedMapInput is an input type that accepts RealmKeystoreAesGeneratedMap and RealmKeystoreAesGeneratedMapOutput values.
// You can construct a concrete instance of `RealmKeystoreAesGeneratedMapInput` via:
//
//	RealmKeystoreAesGeneratedMap{ "key": RealmKeystoreAesGeneratedArgs{...} }
type RealmKeystoreAesGeneratedMapInput interface {
	pulumi.Input

	ToRealmKeystoreAesGeneratedMapOutput() RealmKeystoreAesGeneratedMapOutput
	ToRealmKeystoreAesGeneratedMapOutputWithContext(context.Context) RealmKeystoreAesGeneratedMapOutput
}

type RealmKeystoreAesGeneratedMap map[string]RealmKeystoreAesGeneratedInput

func (RealmKeystoreAesGeneratedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreAesGenerated)(nil)).Elem()
}

func (i RealmKeystoreAesGeneratedMap) ToRealmKeystoreAesGeneratedMapOutput() RealmKeystoreAesGeneratedMapOutput {
	return i.ToRealmKeystoreAesGeneratedMapOutputWithContext(context.Background())
}

func (i RealmKeystoreAesGeneratedMap) ToRealmKeystoreAesGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreAesGeneratedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreAesGeneratedMapOutput)
}

type RealmKeystoreAesGeneratedOutput struct{ *pulumi.OutputState }

func (RealmKeystoreAesGeneratedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreAesGenerated)(nil)).Elem()
}

func (o RealmKeystoreAesGeneratedOutput) ToRealmKeystoreAesGeneratedOutput() RealmKeystoreAesGeneratedOutput {
	return o
}

func (o RealmKeystoreAesGeneratedOutput) ToRealmKeystoreAesGeneratedOutputWithContext(ctx context.Context) RealmKeystoreAesGeneratedOutput {
	return o
}

// When `false`, key in not used for signing. Defaults to `true`.
func (o RealmKeystoreAesGeneratedOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreAesGenerated) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// When `false`, key is not accessible in this realm. Defaults to `true`.
func (o RealmKeystoreAesGeneratedOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreAesGenerated) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Display name of provider when linked in admin console.
func (o RealmKeystoreAesGeneratedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmKeystoreAesGenerated) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Priority for the provider. Defaults to `0`
func (o RealmKeystoreAesGeneratedOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreAesGenerated) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The realm this keystore exists in.
func (o RealmKeystoreAesGeneratedOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmKeystoreAesGenerated) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
func (o RealmKeystoreAesGeneratedOutput) SecretSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RealmKeystoreAesGenerated) pulumi.IntPtrOutput { return v.SecretSize }).(pulumi.IntPtrOutput)
}

type RealmKeystoreAesGeneratedArrayOutput struct{ *pulumi.OutputState }

func (RealmKeystoreAesGeneratedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreAesGenerated)(nil)).Elem()
}

func (o RealmKeystoreAesGeneratedArrayOutput) ToRealmKeystoreAesGeneratedArrayOutput() RealmKeystoreAesGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreAesGeneratedArrayOutput) ToRealmKeystoreAesGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreAesGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreAesGeneratedArrayOutput) Index(i pulumi.IntInput) RealmKeystoreAesGeneratedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealmKeystoreAesGenerated {
		return vs[0].([]*RealmKeystoreAesGenerated)[vs[1].(int)]
	}).(RealmKeystoreAesGeneratedOutput)
}

type RealmKeystoreAesGeneratedMapOutput struct{ *pulumi.OutputState }

func (RealmKeystoreAesGeneratedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreAesGenerated)(nil)).Elem()
}

func (o RealmKeystoreAesGeneratedMapOutput) ToRealmKeystoreAesGeneratedMapOutput() RealmKeystoreAesGeneratedMapOutput {
	return o
}

func (o RealmKeystoreAesGeneratedMapOutput) ToRealmKeystoreAesGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreAesGeneratedMapOutput {
	return o
}

func (o RealmKeystoreAesGeneratedMapOutput) MapIndex(k pulumi.StringInput) RealmKeystoreAesGeneratedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealmKeystoreAesGenerated {
		return vs[0].(map[string]*RealmKeystoreAesGenerated)[vs[1].(string)]
	}).(RealmKeystoreAesGeneratedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreAesGeneratedInput)(nil)).Elem(), &RealmKeystoreAesGenerated{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreAesGeneratedArrayInput)(nil)).Elem(), RealmKeystoreAesGeneratedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreAesGeneratedMapInput)(nil)).Elem(), RealmKeystoreAesGeneratedMap{})
	pulumi.RegisterOutputType(RealmKeystoreAesGeneratedOutput{})
	pulumi.RegisterOutputType(RealmKeystoreAesGeneratedArrayOutput{})
	pulumi.RegisterOutputType(RealmKeystoreAesGeneratedMapOutput{})
}
