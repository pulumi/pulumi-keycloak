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

// Allows managing default realm roles within Keycloak.
//
// Note: This feature was added in Keycloak v13, so this resource will not work on older versions of Keycloak.
//
// ## Example Usage
//
// ### Realm Role)
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
//				Realm:   pulumi.String("my-realm"),
//				Enabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewDefaultRoles(ctx, "default_roles", &keycloak.DefaultRolesArgs{
//				RealmId: realm.ID(),
//				DefaultRoles: pulumi.StringArray{
//					pulumi.String("uma_authorization"),
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
// Default roles can be imported using the format `{{realm_id}}/{{default_role_id}}`, where `default_role_id` is the unique ID of the composite
//
// role that Keycloak uses to control default realm level roles. The ID is not easy to find in the GUI, but it appears in the dev tools when editing
//
// the default roles.
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:index/defaultRoles:DefaultRoles default_roles my-realm/a04c35c2-e95a-4dc5-bd32-e83a21be9e7d
// ```
type DefaultRoles struct {
	pulumi.CustomResourceState

	// Realm level roles assigned to new users by default.
	DefaultRoles pulumi.StringArrayOutput `pulumi:"defaultRoles"`
	// The realm this role exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewDefaultRoles registers a new resource with the given unique name, arguments, and options.
func NewDefaultRoles(ctx *pulumi.Context,
	name string, args *DefaultRolesArgs, opts ...pulumi.ResourceOption) (*DefaultRoles, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultRoles == nil {
		return nil, errors.New("invalid value for required argument 'DefaultRoles'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultRoles
	err := ctx.RegisterResource("keycloak:index/defaultRoles:DefaultRoles", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultRoles gets an existing DefaultRoles resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultRoles(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultRolesState, opts ...pulumi.ResourceOption) (*DefaultRoles, error) {
	var resource DefaultRoles
	err := ctx.ReadResource("keycloak:index/defaultRoles:DefaultRoles", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultRoles resources.
type defaultRolesState struct {
	// Realm level roles assigned to new users by default.
	DefaultRoles []string `pulumi:"defaultRoles"`
	// The realm this role exists within.
	RealmId *string `pulumi:"realmId"`
}

type DefaultRolesState struct {
	// Realm level roles assigned to new users by default.
	DefaultRoles pulumi.StringArrayInput
	// The realm this role exists within.
	RealmId pulumi.StringPtrInput
}

func (DefaultRolesState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultRolesState)(nil)).Elem()
}

type defaultRolesArgs struct {
	// Realm level roles assigned to new users by default.
	DefaultRoles []string `pulumi:"defaultRoles"`
	// The realm this role exists within.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a DefaultRoles resource.
type DefaultRolesArgs struct {
	// Realm level roles assigned to new users by default.
	DefaultRoles pulumi.StringArrayInput
	// The realm this role exists within.
	RealmId pulumi.StringInput
}

func (DefaultRolesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultRolesArgs)(nil)).Elem()
}

type DefaultRolesInput interface {
	pulumi.Input

	ToDefaultRolesOutput() DefaultRolesOutput
	ToDefaultRolesOutputWithContext(ctx context.Context) DefaultRolesOutput
}

func (*DefaultRoles) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultRoles)(nil)).Elem()
}

func (i *DefaultRoles) ToDefaultRolesOutput() DefaultRolesOutput {
	return i.ToDefaultRolesOutputWithContext(context.Background())
}

func (i *DefaultRoles) ToDefaultRolesOutputWithContext(ctx context.Context) DefaultRolesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultRolesOutput)
}

// DefaultRolesArrayInput is an input type that accepts DefaultRolesArray and DefaultRolesArrayOutput values.
// You can construct a concrete instance of `DefaultRolesArrayInput` via:
//
//	DefaultRolesArray{ DefaultRolesArgs{...} }
type DefaultRolesArrayInput interface {
	pulumi.Input

	ToDefaultRolesArrayOutput() DefaultRolesArrayOutput
	ToDefaultRolesArrayOutputWithContext(context.Context) DefaultRolesArrayOutput
}

type DefaultRolesArray []DefaultRolesInput

func (DefaultRolesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultRoles)(nil)).Elem()
}

func (i DefaultRolesArray) ToDefaultRolesArrayOutput() DefaultRolesArrayOutput {
	return i.ToDefaultRolesArrayOutputWithContext(context.Background())
}

func (i DefaultRolesArray) ToDefaultRolesArrayOutputWithContext(ctx context.Context) DefaultRolesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultRolesArrayOutput)
}

// DefaultRolesMapInput is an input type that accepts DefaultRolesMap and DefaultRolesMapOutput values.
// You can construct a concrete instance of `DefaultRolesMapInput` via:
//
//	DefaultRolesMap{ "key": DefaultRolesArgs{...} }
type DefaultRolesMapInput interface {
	pulumi.Input

	ToDefaultRolesMapOutput() DefaultRolesMapOutput
	ToDefaultRolesMapOutputWithContext(context.Context) DefaultRolesMapOutput
}

type DefaultRolesMap map[string]DefaultRolesInput

func (DefaultRolesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultRoles)(nil)).Elem()
}

func (i DefaultRolesMap) ToDefaultRolesMapOutput() DefaultRolesMapOutput {
	return i.ToDefaultRolesMapOutputWithContext(context.Background())
}

func (i DefaultRolesMap) ToDefaultRolesMapOutputWithContext(ctx context.Context) DefaultRolesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultRolesMapOutput)
}

type DefaultRolesOutput struct{ *pulumi.OutputState }

func (DefaultRolesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultRoles)(nil)).Elem()
}

func (o DefaultRolesOutput) ToDefaultRolesOutput() DefaultRolesOutput {
	return o
}

func (o DefaultRolesOutput) ToDefaultRolesOutputWithContext(ctx context.Context) DefaultRolesOutput {
	return o
}

// Realm level roles assigned to new users by default.
func (o DefaultRolesOutput) DefaultRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultRoles) pulumi.StringArrayOutput { return v.DefaultRoles }).(pulumi.StringArrayOutput)
}

// The realm this role exists within.
func (o DefaultRolesOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultRoles) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type DefaultRolesArrayOutput struct{ *pulumi.OutputState }

func (DefaultRolesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultRoles)(nil)).Elem()
}

func (o DefaultRolesArrayOutput) ToDefaultRolesArrayOutput() DefaultRolesArrayOutput {
	return o
}

func (o DefaultRolesArrayOutput) ToDefaultRolesArrayOutputWithContext(ctx context.Context) DefaultRolesArrayOutput {
	return o
}

func (o DefaultRolesArrayOutput) Index(i pulumi.IntInput) DefaultRolesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultRoles {
		return vs[0].([]*DefaultRoles)[vs[1].(int)]
	}).(DefaultRolesOutput)
}

type DefaultRolesMapOutput struct{ *pulumi.OutputState }

func (DefaultRolesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultRoles)(nil)).Elem()
}

func (o DefaultRolesMapOutput) ToDefaultRolesMapOutput() DefaultRolesMapOutput {
	return o
}

func (o DefaultRolesMapOutput) ToDefaultRolesMapOutputWithContext(ctx context.Context) DefaultRolesMapOutput {
	return o
}

func (o DefaultRolesMapOutput) MapIndex(k pulumi.StringInput) DefaultRolesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultRoles {
		return vs[0].(map[string]*DefaultRoles)[vs[1].(string)]
	}).(DefaultRolesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultRolesInput)(nil)).Elem(), &DefaultRoles{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultRolesArrayInput)(nil)).Elem(), DefaultRolesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultRolesMapInput)(nil)).Elem(), DefaultRolesMap{})
	pulumi.RegisterOutputType(DefaultRolesOutput{})
	pulumi.RegisterOutputType(DefaultRolesArrayOutput{})
	pulumi.RegisterOutputType(DefaultRolesMapOutput{})
}
