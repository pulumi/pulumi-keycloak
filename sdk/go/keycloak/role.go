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

// ## # Role
//
// Allows for creating and managing roles within Keycloak.
//
// Roles allow you define privileges within Keycloak and map them to users
// and groups.
//
// ### Example Usage (Realm role)
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
//				Enabled: pulumi.Bool(true),
//				Realm:   pulumi.String("my-realm"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "realmRole", &keycloak.RoleArgs{
//				Description: pulumi.String("My Realm Role"),
//				RealmId:     realm.ID(),
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
// ### Example Usage (Client role)
//
// <!--Start PulumiCodeChooser -->
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
//				Enabled: pulumi.Bool(true),
//				Realm:   pulumi.String("my-realm"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewClient(ctx, "client", &openid.ClientArgs{
//				AccessType: pulumi.String("BEARER-ONLY"),
//				ClientId:   pulumi.String("client"),
//				Enabled:    pulumi.Bool(true),
//				RealmId:    realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "clientRole", &keycloak.RoleArgs{
//				ClientId:    pulumi.Any(keycloak_client.Client.Id),
//				Description: pulumi.String("My Client Role"),
//				RealmId:     realm.ID(),
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
// ### Example Usage (Composite role)
//
// <!--Start PulumiCodeChooser -->
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
//				Enabled: pulumi.Bool(true),
//				Realm:   pulumi.String("my-realm"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "createRole", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "readRole", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "updateRole", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "deleteRole", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewClient(ctx, "client", &openid.ClientArgs{
//				AccessType: pulumi.String("BEARER-ONLY"),
//				ClientId:   pulumi.String("client"),
//				Enabled:    pulumi.Bool(true),
//				RealmId:    realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "clientRole", &keycloak.RoleArgs{
//				ClientId:    pulumi.Any(keycloak_client.Client.Id),
//				Description: pulumi.String("My Client Role"),
//				RealmId:     realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewRole(ctx, "adminRole", &keycloak.RoleArgs{
//				CompositeRoles: pulumi.StringArray{
//					pulumi.String("{keycloak_role.create_role.id}"),
//					pulumi.String("{keycloak_role.read_role.id}"),
//					pulumi.String("{keycloak_role.update_role.id}"),
//					pulumi.String("{keycloak_role.delete_role.id}"),
//					pulumi.String("{keycloak_role.client_role.id}"),
//				},
//				RealmId: realm.ID(),
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
// ### Argument Reference
//
// The following arguments are supported:
//
//   - `realmId` - (Required) The realm this role exists within.
//   - `clientId` - (Optional) When specified, this role will be created as
//     a client role attached to the client with the provided ID
//   - `name` - (Required) The name of the role
//   - `description` - (Optional) The description of the role
//   - `compositeRoles` - (Optional) When specified, this role will be a
//     composite role, composed of all roles that have an ID present within
//     this list.
//
// ### Import
//
// Roles can be imported using the format `{{realm_id}}/{{role_id}}`, where
// `roleId` is the unique ID that Keycloak assigns to the role. The ID is
// not easy to find in the GUI, but it appears in the URL when editing the
// role.
//
// Example:
type Role struct {
	pulumi.CustomResourceState

	Attributes     pulumi.MapOutput         `pulumi:"attributes"`
	ClientId       pulumi.StringPtrOutput   `pulumi:"clientId"`
	CompositeRoles pulumi.StringArrayOutput `pulumi:"compositeRoles"`
	Description    pulumi.StringPtrOutput   `pulumi:"description"`
	Name           pulumi.StringOutput      `pulumi:"name"`
	RealmId        pulumi.StringOutput      `pulumi:"realmId"`
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
	Attributes     map[string]interface{} `pulumi:"attributes"`
	ClientId       *string                `pulumi:"clientId"`
	CompositeRoles []string               `pulumi:"compositeRoles"`
	Description    *string                `pulumi:"description"`
	Name           *string                `pulumi:"name"`
	RealmId        *string                `pulumi:"realmId"`
}

type RoleState struct {
	Attributes     pulumi.MapInput
	ClientId       pulumi.StringPtrInput
	CompositeRoles pulumi.StringArrayInput
	Description    pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	RealmId        pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	Attributes     map[string]interface{} `pulumi:"attributes"`
	ClientId       *string                `pulumi:"clientId"`
	CompositeRoles []string               `pulumi:"compositeRoles"`
	Description    *string                `pulumi:"description"`
	Name           *string                `pulumi:"name"`
	RealmId        string                 `pulumi:"realmId"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	Attributes     pulumi.MapInput
	ClientId       pulumi.StringPtrInput
	CompositeRoles pulumi.StringArrayInput
	Description    pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	RealmId        pulumi.StringInput
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

func (o RoleOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v *Role) pulumi.MapOutput { return v.Attributes }).(pulumi.MapOutput)
}

func (o RoleOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o RoleOutput) CompositeRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.CompositeRoles }).(pulumi.StringArrayOutput)
}

func (o RoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

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
