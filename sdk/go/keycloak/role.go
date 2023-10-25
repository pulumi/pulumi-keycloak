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

// Allows for creating and managing roles within Keycloak.
//
// Roles allow you define privileges within Keycloak and map them to users and groups.
//
// ## Example Usage
// ### Realm Role)
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
//			_, err = keycloak.NewRole(ctx, "realmRole", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				Description: pulumi.String("My Realm Role"),
//				Attributes: pulumi.Map{
//					"key":        pulumi.Any("value"),
//					"multivalue": pulumi.Any("value1##value2"),
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
// ### Client Role)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/openid"
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
//			_, err = openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
//				RealmId:    realm.ID(),
//				ClientId:   pulumi.String("client"),
//				Enabled:    pulumi.Bool(true),
//				AccessType: pulumi.String("CONFIDENTIAL"),
//				ValidRedirectUris: pulumi.StringArray{
//					pulumi.String("http://localhost:8080/openid-callback"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "clientRole", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				ClientId:    pulumi.Any(keycloak_client.Openid_client.Id),
//				Description: pulumi.String("My Client Role"),
//				Attributes: pulumi.Map{
//					"key": pulumi.Any("value"),
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
// ### Composite Role)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/openid"
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
//			createRole, err := keycloak.NewRole(ctx, "createRole", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//				Attributes: pulumi.Map{
//					"key": pulumi.Any("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			readRole, err := keycloak.NewRole(ctx, "readRole", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//				Attributes: pulumi.Map{
//					"key": pulumi.Any("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			updateRole, err := keycloak.NewRole(ctx, "updateRole", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//				Attributes: pulumi.Map{
//					"key": pulumi.Any("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			deleteRole, err := keycloak.NewRole(ctx, "deleteRole", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//				Attributes: pulumi.Map{
//					"key": pulumi.Any("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
//				RealmId:    realm.ID(),
//				ClientId:   pulumi.String("client"),
//				Enabled:    pulumi.Bool(true),
//				AccessType: pulumi.String("CONFIDENTIAL"),
//				ValidRedirectUris: pulumi.StringArray{
//					pulumi.String("http://localhost:8080/openid-callback"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			clientRole, err := keycloak.NewRole(ctx, "clientRole", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				ClientId:    pulumi.Any(keycloak_client.Openid_client.Id),
//				Description: pulumi.String("My Client Role"),
//				Attributes: pulumi.Map{
//					"key": pulumi.Any("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "adminRole", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//				CompositeRoles: pulumi.StringArray{
//					createRole.ID(),
//					readRole.ID(),
//					updateRole.ID(),
//					deleteRole.ID(),
//					clientRole.ID(),
//				},
//				Attributes: pulumi.Map{
//					"key": pulumi.Any("value"),
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
// Roles can be imported using the format `{{realm_id}}/{{role_id}}`, where `role_id` is the unique ID that Keycloak assigns to the role. The ID is not easy to find in the GUI, but it appears in the URL when editing the role. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:index/role:Role role my-realm/7e8cf32a-8acb-4d34-89c4-04fb1d10ccad
//
// ```
type Role struct {
	pulumi.CustomResourceState

	// A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes pulumi.MapOutput `pulumi:"attributes"`
	// When specified, this role will be created as a client role attached to the client with the provided ID
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
	CompositeRoles pulumi.StringArrayOutput `pulumi:"compositeRoles"`
	// The description of the role
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the role
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this role exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Role
	err := ctx.RegisterResource("keycloak:index/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("keycloak:index/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes map[string]interface{} `pulumi:"attributes"`
	// When specified, this role will be created as a client role attached to the client with the provided ID
	ClientId *string `pulumi:"clientId"`
	// When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
	CompositeRoles []string `pulumi:"compositeRoles"`
	// The description of the role
	Description *string `pulumi:"description"`
	// The name of the role
	Name *string `pulumi:"name"`
	// The realm this role exists within.
	RealmId *string `pulumi:"realmId"`
}

type RoleState struct {
	// A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes pulumi.MapInput
	// When specified, this role will be created as a client role attached to the client with the provided ID
	ClientId pulumi.StringPtrInput
	// When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
	CompositeRoles pulumi.StringArrayInput
	// The description of the role
	Description pulumi.StringPtrInput
	// The name of the role
	Name pulumi.StringPtrInput
	// The realm this role exists within.
	RealmId pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes map[string]interface{} `pulumi:"attributes"`
	// When specified, this role will be created as a client role attached to the client with the provided ID
	ClientId *string `pulumi:"clientId"`
	// When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
	CompositeRoles []string `pulumi:"compositeRoles"`
	// The description of the role
	Description *string `pulumi:"description"`
	// The name of the role
	Name *string `pulumi:"name"`
	// The realm this role exists within.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes pulumi.MapInput
	// When specified, this role will be created as a client role attached to the client with the provided ID
	ClientId pulumi.StringPtrInput
	// When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
	CompositeRoles pulumi.StringArrayInput
	// The description of the role
	Description pulumi.StringPtrInput
	// The name of the role
	Name pulumi.StringPtrInput
	// The realm this role exists within.
	RealmId pulumi.StringInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

func (i *Role) ToOutput(ctx context.Context) pulumix.Output[*Role] {
	return pulumix.Output[*Role]{
		OutputState: i.ToRoleOutputWithContext(ctx).OutputState,
	}
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//	RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

func (i RoleArray) ToOutput(ctx context.Context) pulumix.Output[[]*Role] {
	return pulumix.Output[[]*Role]{
		OutputState: i.ToRoleArrayOutputWithContext(ctx).OutputState,
	}
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//	RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

func (i RoleMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Role] {
	return pulumix.Output[map[string]*Role]{
		OutputState: i.ToRoleMapOutputWithContext(ctx).OutputState,
	}
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

func (o RoleOutput) ToOutput(ctx context.Context) pulumix.Output[*Role] {
	return pulumix.Output[*Role]{
		OutputState: o.OutputState,
	}
}

// A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
func (o RoleOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v *Role) pulumi.MapOutput { return v.Attributes }).(pulumi.MapOutput)
}

// When specified, this role will be created as a client role attached to the client with the provided ID
func (o RoleOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
func (o RoleOutput) CompositeRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.CompositeRoles }).(pulumi.StringArrayOutput)
}

// The description of the role
func (o RoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the role
func (o RoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm this role exists within.
func (o RoleOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Role] {
	return pulumix.Output[[]*Role]{
		OutputState: o.OutputState,
	}
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Role] {
	return pulumix.Output[map[string]*Role]{
		OutputState: o.OutputState,
	}
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}
