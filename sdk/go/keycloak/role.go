// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Enabled: pulumi.Bool(true),
// 			Realm:   pulumi.String("my-realm"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewRole(ctx, "realmRole", &keycloak.RoleArgs{
// 			Description: pulumi.String("My Realm Role"),
// 			RealmId:     realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Example Usage (Client role)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak/openid"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Enabled: pulumi.Bool(true),
// 			Realm:   pulumi.String("my-realm"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewClient(ctx, "client", &openid.ClientArgs{
// 			AccessType: pulumi.String("BEARER-ONLY"),
// 			ClientId:   pulumi.String("client"),
// 			Enabled:    pulumi.Bool(true),
// 			RealmId:    realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewRole(ctx, "clientRole", &keycloak.RoleArgs{
// 			ClientId:    pulumi.String(keycloak_client.Client.Id),
// 			Description: pulumi.String("My Client Role"),
// 			RealmId:     realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Example Usage (Composite role)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v2/go/keycloak/openid"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Enabled: pulumi.Bool(true),
// 			Realm:   pulumi.String("my-realm"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewRole(ctx, "createRole", &keycloak.RoleArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewRole(ctx, "readRole", &keycloak.RoleArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewRole(ctx, "updateRole", &keycloak.RoleArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewRole(ctx, "deleteRole", &keycloak.RoleArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewClient(ctx, "client", &openid.ClientArgs{
// 			AccessType: pulumi.String("BEARER-ONLY"),
// 			ClientId:   pulumi.String("client"),
// 			Enabled:    pulumi.Bool(true),
// 			RealmId:    realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewRole(ctx, "clientRole", &keycloak.RoleArgs{
// 			ClientId:    pulumi.String(keycloak_client.Client.Id),
// 			Description: pulumi.String("My Client Role"),
// 			RealmId:     realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keycloak.NewRole(ctx, "adminRole", &keycloak.RoleArgs{
// 			CompositeRoles: pulumi.StringArray{
// 				pulumi.String("{keycloak_role.create_role.id}"),
// 				pulumi.String("{keycloak_role.read_role.id}"),
// 				pulumi.String("{keycloak_role.update_role.id}"),
// 				pulumi.String("{keycloak_role.delete_role.id}"),
// 				pulumi.String("{keycloak_role.client_role.id}"),
// 			},
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm this role exists within.
// - `clientId` - (Optional) When specified, this role will be created as
//   a client role attached to the client with the provided ID
// - `name` - (Required) The name of the role
// - `description` - (Optional) The description of the role
// - `compositeRoles` - (Optional) When specified, this role will be a
//   composite role, composed of all roles that have an ID present within
//   this list.
type Role struct {
	pulumi.CustomResourceState

	ClientId       pulumi.StringPtrOutput   `pulumi:"clientId"`
	CompositeRoles pulumi.StringArrayOutput `pulumi:"compositeRoles"`
	Description    pulumi.StringPtrOutput   `pulumi:"description"`
	Name           pulumi.StringOutput      `pulumi:"name"`
	RealmId        pulumi.StringOutput      `pulumi:"realmId"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &RoleArgs{}
	}
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
	ClientId       *string  `pulumi:"clientId"`
	CompositeRoles []string `pulumi:"compositeRoles"`
	Description    *string  `pulumi:"description"`
	Name           *string  `pulumi:"name"`
	RealmId        *string  `pulumi:"realmId"`
}

type RoleState struct {
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
	ClientId       *string  `pulumi:"clientId"`
	CompositeRoles []string `pulumi:"compositeRoles"`
	Description    *string  `pulumi:"description"`
	Name           *string  `pulumi:"name"`
	RealmId        string   `pulumi:"realmId"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	ClientId       pulumi.StringPtrInput
	CompositeRoles pulumi.StringArrayInput
	Description    pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	RealmId        pulumi.StringInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}
