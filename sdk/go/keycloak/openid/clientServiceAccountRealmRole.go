// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for assigning realm roles to the service account of an openid client.
// You need to set `serviceAccountsEnabled` to `true` for the openid client that should be assigned the role.
//
// If you'd like to attach client roles to a service account, please use the `openid.ClientServiceAccountRole`
// resource.
//
// ## Example Usage
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
//			realmRole, err := keycloak.NewRole(ctx, "realm_role", &keycloak.RoleArgs{
//				RealmId: realm.ID(),
//				Name:    pulumi.String("my-realm-role"),
//			})
//			if err != nil {
//				return err
//			}
//			client, err := openid.NewClient(ctx, "client", &openid.ClientArgs{
//				RealmId:                realm.ID(),
//				Name:                   pulumi.String("client"),
//				ServiceAccountsEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewClientServiceAccountRealmRole(ctx, "client_service_account_role", &openid.ClientServiceAccountRealmRoleArgs{
//				RealmId:              realm.ID(),
//				ServiceAccountUserId: client.ServiceAccountUserId,
//				Role:                 realmRole.Name,
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
// This resource can be imported using the format `{{realmId}}/{{serviceAccountUserId}}/{{roleId}}`.
//
// Example:
//
// bash
//
// ```sh
// $ pulumi import keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole client_service_account_role my-realm/489ba513-1ceb-49ba-ae0b-1ab1f5099ebf/c7230ab7-8e4e-4135-995d-e81b50696ad8
// ```
type ClientServiceAccountRealmRole struct {
	pulumi.CustomResourceState

	// The realm that the client and role belong to.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The name of the role that is assigned.
	Role pulumi.StringOutput `pulumi:"role"`
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId pulumi.StringOutput `pulumi:"serviceAccountUserId"`
}

// NewClientServiceAccountRealmRole registers a new resource with the given unique name, arguments, and options.
func NewClientServiceAccountRealmRole(ctx *pulumi.Context,
	name string, args *ClientServiceAccountRealmRoleArgs, opts ...pulumi.ResourceOption) (*ClientServiceAccountRealmRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.ServiceAccountUserId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountUserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClientServiceAccountRealmRole
	err := ctx.RegisterResource("keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientServiceAccountRealmRole gets an existing ClientServiceAccountRealmRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientServiceAccountRealmRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientServiceAccountRealmRoleState, opts ...pulumi.ResourceOption) (*ClientServiceAccountRealmRole, error) {
	var resource ClientServiceAccountRealmRole
	err := ctx.ReadResource("keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientServiceAccountRealmRole resources.
type clientServiceAccountRealmRoleState struct {
	// The realm that the client and role belong to.
	RealmId *string `pulumi:"realmId"`
	// The name of the role that is assigned.
	Role *string `pulumi:"role"`
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId *string `pulumi:"serviceAccountUserId"`
}

type ClientServiceAccountRealmRoleState struct {
	// The realm that the client and role belong to.
	RealmId pulumi.StringPtrInput
	// The name of the role that is assigned.
	Role pulumi.StringPtrInput
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId pulumi.StringPtrInput
}

func (ClientServiceAccountRealmRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientServiceAccountRealmRoleState)(nil)).Elem()
}

type clientServiceAccountRealmRoleArgs struct {
	// The realm that the client and role belong to.
	RealmId string `pulumi:"realmId"`
	// The name of the role that is assigned.
	Role string `pulumi:"role"`
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId string `pulumi:"serviceAccountUserId"`
}

// The set of arguments for constructing a ClientServiceAccountRealmRole resource.
type ClientServiceAccountRealmRoleArgs struct {
	// The realm that the client and role belong to.
	RealmId pulumi.StringInput
	// The name of the role that is assigned.
	Role pulumi.StringInput
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId pulumi.StringInput
}

func (ClientServiceAccountRealmRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientServiceAccountRealmRoleArgs)(nil)).Elem()
}

type ClientServiceAccountRealmRoleInput interface {
	pulumi.Input

	ToClientServiceAccountRealmRoleOutput() ClientServiceAccountRealmRoleOutput
	ToClientServiceAccountRealmRoleOutputWithContext(ctx context.Context) ClientServiceAccountRealmRoleOutput
}

func (*ClientServiceAccountRealmRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientServiceAccountRealmRole)(nil)).Elem()
}

func (i *ClientServiceAccountRealmRole) ToClientServiceAccountRealmRoleOutput() ClientServiceAccountRealmRoleOutput {
	return i.ToClientServiceAccountRealmRoleOutputWithContext(context.Background())
}

func (i *ClientServiceAccountRealmRole) ToClientServiceAccountRealmRoleOutputWithContext(ctx context.Context) ClientServiceAccountRealmRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientServiceAccountRealmRoleOutput)
}

// ClientServiceAccountRealmRoleArrayInput is an input type that accepts ClientServiceAccountRealmRoleArray and ClientServiceAccountRealmRoleArrayOutput values.
// You can construct a concrete instance of `ClientServiceAccountRealmRoleArrayInput` via:
//
//	ClientServiceAccountRealmRoleArray{ ClientServiceAccountRealmRoleArgs{...} }
type ClientServiceAccountRealmRoleArrayInput interface {
	pulumi.Input

	ToClientServiceAccountRealmRoleArrayOutput() ClientServiceAccountRealmRoleArrayOutput
	ToClientServiceAccountRealmRoleArrayOutputWithContext(context.Context) ClientServiceAccountRealmRoleArrayOutput
}

type ClientServiceAccountRealmRoleArray []ClientServiceAccountRealmRoleInput

func (ClientServiceAccountRealmRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientServiceAccountRealmRole)(nil)).Elem()
}

func (i ClientServiceAccountRealmRoleArray) ToClientServiceAccountRealmRoleArrayOutput() ClientServiceAccountRealmRoleArrayOutput {
	return i.ToClientServiceAccountRealmRoleArrayOutputWithContext(context.Background())
}

func (i ClientServiceAccountRealmRoleArray) ToClientServiceAccountRealmRoleArrayOutputWithContext(ctx context.Context) ClientServiceAccountRealmRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientServiceAccountRealmRoleArrayOutput)
}

// ClientServiceAccountRealmRoleMapInput is an input type that accepts ClientServiceAccountRealmRoleMap and ClientServiceAccountRealmRoleMapOutput values.
// You can construct a concrete instance of `ClientServiceAccountRealmRoleMapInput` via:
//
//	ClientServiceAccountRealmRoleMap{ "key": ClientServiceAccountRealmRoleArgs{...} }
type ClientServiceAccountRealmRoleMapInput interface {
	pulumi.Input

	ToClientServiceAccountRealmRoleMapOutput() ClientServiceAccountRealmRoleMapOutput
	ToClientServiceAccountRealmRoleMapOutputWithContext(context.Context) ClientServiceAccountRealmRoleMapOutput
}

type ClientServiceAccountRealmRoleMap map[string]ClientServiceAccountRealmRoleInput

func (ClientServiceAccountRealmRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientServiceAccountRealmRole)(nil)).Elem()
}

func (i ClientServiceAccountRealmRoleMap) ToClientServiceAccountRealmRoleMapOutput() ClientServiceAccountRealmRoleMapOutput {
	return i.ToClientServiceAccountRealmRoleMapOutputWithContext(context.Background())
}

func (i ClientServiceAccountRealmRoleMap) ToClientServiceAccountRealmRoleMapOutputWithContext(ctx context.Context) ClientServiceAccountRealmRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientServiceAccountRealmRoleMapOutput)
}

type ClientServiceAccountRealmRoleOutput struct{ *pulumi.OutputState }

func (ClientServiceAccountRealmRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientServiceAccountRealmRole)(nil)).Elem()
}

func (o ClientServiceAccountRealmRoleOutput) ToClientServiceAccountRealmRoleOutput() ClientServiceAccountRealmRoleOutput {
	return o
}

func (o ClientServiceAccountRealmRoleOutput) ToClientServiceAccountRealmRoleOutputWithContext(ctx context.Context) ClientServiceAccountRealmRoleOutput {
	return o
}

// The realm that the client and role belong to.
func (o ClientServiceAccountRealmRoleOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientServiceAccountRealmRole) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The name of the role that is assigned.
func (o ClientServiceAccountRealmRoleOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientServiceAccountRealmRole) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
func (o ClientServiceAccountRealmRoleOutput) ServiceAccountUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientServiceAccountRealmRole) pulumi.StringOutput { return v.ServiceAccountUserId }).(pulumi.StringOutput)
}

type ClientServiceAccountRealmRoleArrayOutput struct{ *pulumi.OutputState }

func (ClientServiceAccountRealmRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientServiceAccountRealmRole)(nil)).Elem()
}

func (o ClientServiceAccountRealmRoleArrayOutput) ToClientServiceAccountRealmRoleArrayOutput() ClientServiceAccountRealmRoleArrayOutput {
	return o
}

func (o ClientServiceAccountRealmRoleArrayOutput) ToClientServiceAccountRealmRoleArrayOutputWithContext(ctx context.Context) ClientServiceAccountRealmRoleArrayOutput {
	return o
}

func (o ClientServiceAccountRealmRoleArrayOutput) Index(i pulumi.IntInput) ClientServiceAccountRealmRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientServiceAccountRealmRole {
		return vs[0].([]*ClientServiceAccountRealmRole)[vs[1].(int)]
	}).(ClientServiceAccountRealmRoleOutput)
}

type ClientServiceAccountRealmRoleMapOutput struct{ *pulumi.OutputState }

func (ClientServiceAccountRealmRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientServiceAccountRealmRole)(nil)).Elem()
}

func (o ClientServiceAccountRealmRoleMapOutput) ToClientServiceAccountRealmRoleMapOutput() ClientServiceAccountRealmRoleMapOutput {
	return o
}

func (o ClientServiceAccountRealmRoleMapOutput) ToClientServiceAccountRealmRoleMapOutputWithContext(ctx context.Context) ClientServiceAccountRealmRoleMapOutput {
	return o
}

func (o ClientServiceAccountRealmRoleMapOutput) MapIndex(k pulumi.StringInput) ClientServiceAccountRealmRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientServiceAccountRealmRole {
		return vs[0].(map[string]*ClientServiceAccountRealmRole)[vs[1].(string)]
	}).(ClientServiceAccountRealmRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientServiceAccountRealmRoleInput)(nil)).Elem(), &ClientServiceAccountRealmRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientServiceAccountRealmRoleArrayInput)(nil)).Elem(), ClientServiceAccountRealmRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientServiceAccountRealmRoleMapInput)(nil)).Elem(), ClientServiceAccountRealmRoleMap{})
	pulumi.RegisterOutputType(ClientServiceAccountRealmRoleOutput{})
	pulumi.RegisterOutputType(ClientServiceAccountRealmRoleArrayOutput{})
	pulumi.RegisterOutputType(ClientServiceAccountRealmRoleMapOutput{})
}
