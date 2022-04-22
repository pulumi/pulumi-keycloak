// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing `java-keystore` Realm keystores within Keycloak.
//
// A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
//
// ## Import
//
// Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash
//
// ```sh
//  $ pulumi import keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated java_keystore my-realm/my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
// ```
type RealmKeystoreJavaGenerated struct {
	pulumi.CustomResourceState

	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm pulumi.StringPtrOutput `pulumi:"algorithm"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Alias for the private key
	KeyAlias pulumi.StringOutput `pulumi:"keyAlias"`
	// Password for the private key
	KeyPassword pulumi.StringOutput `pulumi:"keyPassword"`
	// Path to keys file on keycloak instance.
	Keystore pulumi.StringOutput `pulumi:"keystore"`
	// Password for the private key.
	KeystorePassword pulumi.StringOutput `pulumi:"keystorePassword"`
	// Display name of provider when linked in admin console.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewRealmKeystoreJavaGenerated registers a new resource with the given unique name, arguments, and options.
func NewRealmKeystoreJavaGenerated(ctx *pulumi.Context,
	name string, args *RealmKeystoreJavaGeneratedArgs, opts ...pulumi.ResourceOption) (*RealmKeystoreJavaGenerated, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyAlias == nil {
		return nil, errors.New("invalid value for required argument 'KeyAlias'")
	}
	if args.KeyPassword == nil {
		return nil, errors.New("invalid value for required argument 'KeyPassword'")
	}
	if args.Keystore == nil {
		return nil, errors.New("invalid value for required argument 'Keystore'")
	}
	if args.KeystorePassword == nil {
		return nil, errors.New("invalid value for required argument 'KeystorePassword'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	var resource RealmKeystoreJavaGenerated
	err := ctx.RegisterResource("keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealmKeystoreJavaGenerated gets an existing RealmKeystoreJavaGenerated resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealmKeystoreJavaGenerated(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmKeystoreJavaGeneratedState, opts ...pulumi.ResourceOption) (*RealmKeystoreJavaGenerated, error) {
	var resource RealmKeystoreJavaGenerated
	err := ctx.ReadResource("keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealmKeystoreJavaGenerated resources.
type realmKeystoreJavaGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm *string `pulumi:"algorithm"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Alias for the private key
	KeyAlias *string `pulumi:"keyAlias"`
	// Password for the private key
	KeyPassword *string `pulumi:"keyPassword"`
	// Path to keys file on keycloak instance.
	Keystore *string `pulumi:"keystore"`
	// Password for the private key.
	KeystorePassword *string `pulumi:"keystorePassword"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId *string `pulumi:"realmId"`
}

type RealmKeystoreJavaGeneratedState struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm pulumi.StringPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Alias for the private key
	KeyAlias pulumi.StringPtrInput
	// Password for the private key
	KeyPassword pulumi.StringPtrInput
	// Path to keys file on keycloak instance.
	Keystore pulumi.StringPtrInput
	// Password for the private key.
	KeystorePassword pulumi.StringPtrInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringPtrInput
}

func (RealmKeystoreJavaGeneratedState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreJavaGeneratedState)(nil)).Elem()
}

type realmKeystoreJavaGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm *string `pulumi:"algorithm"`
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Alias for the private key
	KeyAlias string `pulumi:"keyAlias"`
	// Password for the private key
	KeyPassword string `pulumi:"keyPassword"`
	// Path to keys file on keycloak instance.
	Keystore string `pulumi:"keystore"`
	// Password for the private key.
	KeystorePassword string `pulumi:"keystorePassword"`
	// Display name of provider when linked in admin console.
	Name *string `pulumi:"name"`
	// Priority for the provider. Defaults to `0`
	Priority *int `pulumi:"priority"`
	// The realm this keystore exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a RealmKeystoreJavaGenerated resource.
type RealmKeystoreJavaGeneratedArgs struct {
	// When `false`, key in not used for signing. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// Intended algorithm for the key. Defaults to `RS256`
	Algorithm pulumi.StringPtrInput
	// When `false`, key is not accessible in this realm. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Alias for the private key
	KeyAlias pulumi.StringInput
	// Password for the private key
	KeyPassword pulumi.StringInput
	// Path to keys file on keycloak instance.
	Keystore pulumi.StringInput
	// Password for the private key.
	KeystorePassword pulumi.StringInput
	// Display name of provider when linked in admin console.
	Name pulumi.StringPtrInput
	// Priority for the provider. Defaults to `0`
	Priority pulumi.IntPtrInput
	// The realm this keystore exists in.
	RealmId pulumi.StringInput
}

func (RealmKeystoreJavaGeneratedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmKeystoreJavaGeneratedArgs)(nil)).Elem()
}

type RealmKeystoreJavaGeneratedInput interface {
	pulumi.Input

	ToRealmKeystoreJavaGeneratedOutput() RealmKeystoreJavaGeneratedOutput
	ToRealmKeystoreJavaGeneratedOutputWithContext(ctx context.Context) RealmKeystoreJavaGeneratedOutput
}

func (*RealmKeystoreJavaGenerated) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreJavaGenerated)(nil)).Elem()
}

func (i *RealmKeystoreJavaGenerated) ToRealmKeystoreJavaGeneratedOutput() RealmKeystoreJavaGeneratedOutput {
	return i.ToRealmKeystoreJavaGeneratedOutputWithContext(context.Background())
}

func (i *RealmKeystoreJavaGenerated) ToRealmKeystoreJavaGeneratedOutputWithContext(ctx context.Context) RealmKeystoreJavaGeneratedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreJavaGeneratedOutput)
}

// RealmKeystoreJavaGeneratedArrayInput is an input type that accepts RealmKeystoreJavaGeneratedArray and RealmKeystoreJavaGeneratedArrayOutput values.
// You can construct a concrete instance of `RealmKeystoreJavaGeneratedArrayInput` via:
//
//          RealmKeystoreJavaGeneratedArray{ RealmKeystoreJavaGeneratedArgs{...} }
type RealmKeystoreJavaGeneratedArrayInput interface {
	pulumi.Input

	ToRealmKeystoreJavaGeneratedArrayOutput() RealmKeystoreJavaGeneratedArrayOutput
	ToRealmKeystoreJavaGeneratedArrayOutputWithContext(context.Context) RealmKeystoreJavaGeneratedArrayOutput
}

type RealmKeystoreJavaGeneratedArray []RealmKeystoreJavaGeneratedInput

func (RealmKeystoreJavaGeneratedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreJavaGenerated)(nil)).Elem()
}

func (i RealmKeystoreJavaGeneratedArray) ToRealmKeystoreJavaGeneratedArrayOutput() RealmKeystoreJavaGeneratedArrayOutput {
	return i.ToRealmKeystoreJavaGeneratedArrayOutputWithContext(context.Background())
}

func (i RealmKeystoreJavaGeneratedArray) ToRealmKeystoreJavaGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreJavaGeneratedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreJavaGeneratedArrayOutput)
}

// RealmKeystoreJavaGeneratedMapInput is an input type that accepts RealmKeystoreJavaGeneratedMap and RealmKeystoreJavaGeneratedMapOutput values.
// You can construct a concrete instance of `RealmKeystoreJavaGeneratedMapInput` via:
//
//          RealmKeystoreJavaGeneratedMap{ "key": RealmKeystoreJavaGeneratedArgs{...} }
type RealmKeystoreJavaGeneratedMapInput interface {
	pulumi.Input

	ToRealmKeystoreJavaGeneratedMapOutput() RealmKeystoreJavaGeneratedMapOutput
	ToRealmKeystoreJavaGeneratedMapOutputWithContext(context.Context) RealmKeystoreJavaGeneratedMapOutput
}

type RealmKeystoreJavaGeneratedMap map[string]RealmKeystoreJavaGeneratedInput

func (RealmKeystoreJavaGeneratedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreJavaGenerated)(nil)).Elem()
}

func (i RealmKeystoreJavaGeneratedMap) ToRealmKeystoreJavaGeneratedMapOutput() RealmKeystoreJavaGeneratedMapOutput {
	return i.ToRealmKeystoreJavaGeneratedMapOutputWithContext(context.Background())
}

func (i RealmKeystoreJavaGeneratedMap) ToRealmKeystoreJavaGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreJavaGeneratedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmKeystoreJavaGeneratedMapOutput)
}

type RealmKeystoreJavaGeneratedOutput struct{ *pulumi.OutputState }

func (RealmKeystoreJavaGeneratedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmKeystoreJavaGenerated)(nil)).Elem()
}

func (o RealmKeystoreJavaGeneratedOutput) ToRealmKeystoreJavaGeneratedOutput() RealmKeystoreJavaGeneratedOutput {
	return o
}

func (o RealmKeystoreJavaGeneratedOutput) ToRealmKeystoreJavaGeneratedOutputWithContext(ctx context.Context) RealmKeystoreJavaGeneratedOutput {
	return o
}

type RealmKeystoreJavaGeneratedArrayOutput struct{ *pulumi.OutputState }

func (RealmKeystoreJavaGeneratedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmKeystoreJavaGenerated)(nil)).Elem()
}

func (o RealmKeystoreJavaGeneratedArrayOutput) ToRealmKeystoreJavaGeneratedArrayOutput() RealmKeystoreJavaGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreJavaGeneratedArrayOutput) ToRealmKeystoreJavaGeneratedArrayOutputWithContext(ctx context.Context) RealmKeystoreJavaGeneratedArrayOutput {
	return o
}

func (o RealmKeystoreJavaGeneratedArrayOutput) Index(i pulumi.IntInput) RealmKeystoreJavaGeneratedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealmKeystoreJavaGenerated {
		return vs[0].([]*RealmKeystoreJavaGenerated)[vs[1].(int)]
	}).(RealmKeystoreJavaGeneratedOutput)
}

type RealmKeystoreJavaGeneratedMapOutput struct{ *pulumi.OutputState }

func (RealmKeystoreJavaGeneratedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmKeystoreJavaGenerated)(nil)).Elem()
}

func (o RealmKeystoreJavaGeneratedMapOutput) ToRealmKeystoreJavaGeneratedMapOutput() RealmKeystoreJavaGeneratedMapOutput {
	return o
}

func (o RealmKeystoreJavaGeneratedMapOutput) ToRealmKeystoreJavaGeneratedMapOutputWithContext(ctx context.Context) RealmKeystoreJavaGeneratedMapOutput {
	return o
}

func (o RealmKeystoreJavaGeneratedMapOutput) MapIndex(k pulumi.StringInput) RealmKeystoreJavaGeneratedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealmKeystoreJavaGenerated {
		return vs[0].(map[string]*RealmKeystoreJavaGenerated)[vs[1].(string)]
	}).(RealmKeystoreJavaGeneratedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreJavaGeneratedInput)(nil)).Elem(), &RealmKeystoreJavaGenerated{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreJavaGeneratedArrayInput)(nil)).Elem(), RealmKeystoreJavaGeneratedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmKeystoreJavaGeneratedMapInput)(nil)).Elem(), RealmKeystoreJavaGeneratedMap{})
	pulumi.RegisterOutputType(RealmKeystoreJavaGeneratedOutput{})
	pulumi.RegisterOutputType(RealmKeystoreJavaGeneratedArrayOutput{})
	pulumi.RegisterOutputType(RealmKeystoreJavaGeneratedMapOutput{})
}
