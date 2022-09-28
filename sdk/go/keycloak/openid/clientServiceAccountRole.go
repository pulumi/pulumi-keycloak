// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for assigning client roles to the service account of an openid client.
// You need to set `serviceAccountsEnabled` to `true` for the openid client that should be assigned the role.
//
// If you'd like to attach realm roles to a service account, please use the `openid.ClientServiceAccountRealmRole`
// resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak/openid"
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
//			client1, err := openid.NewClient(ctx, "client1", &openid.ClientArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			client1Role, err := keycloak.NewRole(ctx, "client1Role", &keycloak.RoleArgs{
//				RealmId:     realm.ID(),
//				ClientId:    client1.ID(),
//				Description: pulumi.String("A role that client1 provides"),
//			})
//			if err != nil {
//				return err
//			}
//			client2, err := openid.NewClient(ctx, "client2", &openid.ClientArgs{
//				RealmId:                realm.ID(),
//				ServiceAccountsEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewClientServiceAccountRole(ctx, "client2ServiceAccountRole", &openid.ClientServiceAccountRoleArgs{
//				RealmId:              realm.ID(),
//				ServiceAccountUserId: client2.ServiceAccountUserId,
//				ClientId:             client1.ID(),
//				Role:                 client1Role.Name,
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
// This resource can be imported using the format `{{realmId}}/{{serviceAccountUserId}}/{{clientId}}/{{roleId}}`. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole client2_service_account_role my-realm/489ba513-1ceb-49ba-ae0b-1ab1f5099ebf/baf01820-0f8b-4494-9be2-fb3bc8a397a4/c7230ab7-8e4e-4135-995d-e81b50696ad8
//
// ```
type ClientServiceAccountRole struct {
	pulumi.CustomResourceState

	// The id of the client that provides the role.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The realm the clients and roles belong to.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The name of the role that is assigned.
	Role pulumi.StringOutput `pulumi:"role"`
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId pulumi.StringOutput `pulumi:"serviceAccountUserId"`
}

// NewClientServiceAccountRole registers a new resource with the given unique name, arguments, and options.
func NewClientServiceAccountRole(ctx *pulumi.Context,
	name string, args *ClientServiceAccountRoleArgs, opts ...pulumi.ResourceOption) (*ClientServiceAccountRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
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
	var resource ClientServiceAccountRole
	err := ctx.RegisterResource("keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientServiceAccountRole gets an existing ClientServiceAccountRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientServiceAccountRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientServiceAccountRoleState, opts ...pulumi.ResourceOption) (*ClientServiceAccountRole, error) {
	var resource ClientServiceAccountRole
	err := ctx.ReadResource("keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientServiceAccountRole resources.
type clientServiceAccountRoleState struct {
	// The id of the client that provides the role.
	ClientId *string `pulumi:"clientId"`
	// The realm the clients and roles belong to.
	RealmId *string `pulumi:"realmId"`
	// The name of the role that is assigned.
	Role *string `pulumi:"role"`
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId *string `pulumi:"serviceAccountUserId"`
}

type ClientServiceAccountRoleState struct {
	// The id of the client that provides the role.
	ClientId pulumi.StringPtrInput
	// The realm the clients and roles belong to.
	RealmId pulumi.StringPtrInput
	// The name of the role that is assigned.
	Role pulumi.StringPtrInput
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId pulumi.StringPtrInput
}

func (ClientServiceAccountRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientServiceAccountRoleState)(nil)).Elem()
}

type clientServiceAccountRoleArgs struct {
	// The id of the client that provides the role.
	ClientId string `pulumi:"clientId"`
	// The realm the clients and roles belong to.
	RealmId string `pulumi:"realmId"`
	// The name of the role that is assigned.
	Role string `pulumi:"role"`
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId string `pulumi:"serviceAccountUserId"`
}

// The set of arguments for constructing a ClientServiceAccountRole resource.
type ClientServiceAccountRoleArgs struct {
	// The id of the client that provides the role.
	ClientId pulumi.StringInput
	// The realm the clients and roles belong to.
	RealmId pulumi.StringInput
	// The name of the role that is assigned.
	Role pulumi.StringInput
	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserId pulumi.StringInput
}

func (ClientServiceAccountRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientServiceAccountRoleArgs)(nil)).Elem()
}

type ClientServiceAccountRoleInput interface {
	pulumi.Input

	ToClientServiceAccountRoleOutput() ClientServiceAccountRoleOutput
	ToClientServiceAccountRoleOutputWithContext(ctx context.Context) ClientServiceAccountRoleOutput
}

func (*ClientServiceAccountRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientServiceAccountRole)(nil)).Elem()
}

func (i *ClientServiceAccountRole) ToClientServiceAccountRoleOutput() ClientServiceAccountRoleOutput {
	return i.ToClientServiceAccountRoleOutputWithContext(context.Background())
}

func (i *ClientServiceAccountRole) ToClientServiceAccountRoleOutputWithContext(ctx context.Context) ClientServiceAccountRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientServiceAccountRoleOutput)
}

// ClientServiceAccountRoleArrayInput is an input type that accepts ClientServiceAccountRoleArray and ClientServiceAccountRoleArrayOutput values.
// You can construct a concrete instance of `ClientServiceAccountRoleArrayInput` via:
//
//	ClientServiceAccountRoleArray{ ClientServiceAccountRoleArgs{...} }
type ClientServiceAccountRoleArrayInput interface {
	pulumi.Input

	ToClientServiceAccountRoleArrayOutput() ClientServiceAccountRoleArrayOutput
	ToClientServiceAccountRoleArrayOutputWithContext(context.Context) ClientServiceAccountRoleArrayOutput
}

type ClientServiceAccountRoleArray []ClientServiceAccountRoleInput

func (ClientServiceAccountRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientServiceAccountRole)(nil)).Elem()
}

func (i ClientServiceAccountRoleArray) ToClientServiceAccountRoleArrayOutput() ClientServiceAccountRoleArrayOutput {
	return i.ToClientServiceAccountRoleArrayOutputWithContext(context.Background())
}

func (i ClientServiceAccountRoleArray) ToClientServiceAccountRoleArrayOutputWithContext(ctx context.Context) ClientServiceAccountRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientServiceAccountRoleArrayOutput)
}

// ClientServiceAccountRoleMapInput is an input type that accepts ClientServiceAccountRoleMap and ClientServiceAccountRoleMapOutput values.
// You can construct a concrete instance of `ClientServiceAccountRoleMapInput` via:
//
//	ClientServiceAccountRoleMap{ "key": ClientServiceAccountRoleArgs{...} }
type ClientServiceAccountRoleMapInput interface {
	pulumi.Input

	ToClientServiceAccountRoleMapOutput() ClientServiceAccountRoleMapOutput
	ToClientServiceAccountRoleMapOutputWithContext(context.Context) ClientServiceAccountRoleMapOutput
}

type ClientServiceAccountRoleMap map[string]ClientServiceAccountRoleInput

func (ClientServiceAccountRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientServiceAccountRole)(nil)).Elem()
}

func (i ClientServiceAccountRoleMap) ToClientServiceAccountRoleMapOutput() ClientServiceAccountRoleMapOutput {
	return i.ToClientServiceAccountRoleMapOutputWithContext(context.Background())
}

func (i ClientServiceAccountRoleMap) ToClientServiceAccountRoleMapOutputWithContext(ctx context.Context) ClientServiceAccountRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientServiceAccountRoleMapOutput)
}

type ClientServiceAccountRoleOutput struct{ *pulumi.OutputState }

func (ClientServiceAccountRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientServiceAccountRole)(nil)).Elem()
}

func (o ClientServiceAccountRoleOutput) ToClientServiceAccountRoleOutput() ClientServiceAccountRoleOutput {
	return o
}

func (o ClientServiceAccountRoleOutput) ToClientServiceAccountRoleOutputWithContext(ctx context.Context) ClientServiceAccountRoleOutput {
	return o
}

// The id of the client that provides the role.
func (o ClientServiceAccountRoleOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientServiceAccountRole) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The realm the clients and roles belong to.
func (o ClientServiceAccountRoleOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientServiceAccountRole) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The name of the role that is assigned.
func (o ClientServiceAccountRoleOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientServiceAccountRole) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
func (o ClientServiceAccountRoleOutput) ServiceAccountUserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientServiceAccountRole) pulumi.StringOutput { return v.ServiceAccountUserId }).(pulumi.StringOutput)
}

type ClientServiceAccountRoleArrayOutput struct{ *pulumi.OutputState }

func (ClientServiceAccountRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientServiceAccountRole)(nil)).Elem()
}

func (o ClientServiceAccountRoleArrayOutput) ToClientServiceAccountRoleArrayOutput() ClientServiceAccountRoleArrayOutput {
	return o
}

func (o ClientServiceAccountRoleArrayOutput) ToClientServiceAccountRoleArrayOutputWithContext(ctx context.Context) ClientServiceAccountRoleArrayOutput {
	return o
}

func (o ClientServiceAccountRoleArrayOutput) Index(i pulumi.IntInput) ClientServiceAccountRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientServiceAccountRole {
		return vs[0].([]*ClientServiceAccountRole)[vs[1].(int)]
	}).(ClientServiceAccountRoleOutput)
}

type ClientServiceAccountRoleMapOutput struct{ *pulumi.OutputState }

func (ClientServiceAccountRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientServiceAccountRole)(nil)).Elem()
}

func (o ClientServiceAccountRoleMapOutput) ToClientServiceAccountRoleMapOutput() ClientServiceAccountRoleMapOutput {
	return o
}

func (o ClientServiceAccountRoleMapOutput) ToClientServiceAccountRoleMapOutputWithContext(ctx context.Context) ClientServiceAccountRoleMapOutput {
	return o
}

func (o ClientServiceAccountRoleMapOutput) MapIndex(k pulumi.StringInput) ClientServiceAccountRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientServiceAccountRole {
		return vs[0].(map[string]*ClientServiceAccountRole)[vs[1].(string)]
	}).(ClientServiceAccountRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientServiceAccountRoleInput)(nil)).Elem(), &ClientServiceAccountRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientServiceAccountRoleArrayInput)(nil)).Elem(), ClientServiceAccountRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientServiceAccountRoleMapInput)(nil)).Elem(), ClientServiceAccountRoleMap{})
	pulumi.RegisterOutputType(ClientServiceAccountRoleOutput{})
	pulumi.RegisterOutputType(ClientServiceAccountRoleArrayOutput{})
	pulumi.RegisterOutputType(ClientServiceAccountRoleMapOutput{})
}
