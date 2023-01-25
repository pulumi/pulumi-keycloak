// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for creating and managing Users within Keycloak.
//
// This resource was created primarily to enable the acceptance tests for the `Group` resource. Creating users within
// Keycloak is not recommended. Instead, users should be federated from external sources by configuring user federation providers
// or identity providers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v4/go/keycloak"
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
//			_, err = keycloak.NewUser(ctx, "user", &keycloak.UserArgs{
//				RealmId:   realm.ID(),
//				Username:  pulumi.String("bob"),
//				Enabled:   pulumi.Bool(true),
//				Email:     pulumi.String("bob@domain.com"),
//				FirstName: pulumi.String("Bob"),
//				LastName:  pulumi.String("Bobson"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = keycloak.NewUser(ctx, "userWithInitialPassword", &keycloak.UserArgs{
//				RealmId:   realm.ID(),
//				Username:  pulumi.String("alice"),
//				Enabled:   pulumi.Bool(true),
//				Email:     pulumi.String("alice@domain.com"),
//				FirstName: pulumi.String("Alice"),
//				LastName:  pulumi.String("Aliceberg"),
//				Attributes: pulumi.AnyMap{
//					"foo":        pulumi.Any("bar"),
//					"multivalue": pulumi.Any("value1##value2"),
//				},
//				InitialPassword: &keycloak.UserInitialPasswordArgs{
//					Value:     pulumi.String("some password"),
//					Temporary: pulumi.Bool(true),
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
// Users can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak assigns to the user upon creation. This value can be found in the GUI when editing the user. Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:index/user:User user my-realm/60c3f971-b1d3-4b3a-9035-d16d7540a5e4
//
// ```
type User struct {
	pulumi.CustomResourceState

	// A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes pulumi.MapOutput `pulumi:"attributes"`
	// The user's email.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// Whether the email address was validated or not. Default to `false`.
	EmailVerified pulumi.BoolPtrOutput `pulumi:"emailVerified"`
	// When false, this user cannot log in. Defaults to `true`.
	Enabled             pulumi.BoolPtrOutput             `pulumi:"enabled"`
	FederatedIdentities UserFederatedIdentityArrayOutput `pulumi:"federatedIdentities"`
	// The user's first name.
	FirstName pulumi.StringPtrOutput `pulumi:"firstName"`
	// When given, the user's initial password will be set. This attribute is only respected during initial user creation.
	InitialPassword UserInitialPasswordPtrOutput `pulumi:"initialPassword"`
	// The user's last name.
	LastName pulumi.StringPtrOutput `pulumi:"lastName"`
	// The realm this user belongs to.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// The unique username of this user.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource User
	err := ctx.RegisterResource("keycloak:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("keycloak:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes map[string]interface{} `pulumi:"attributes"`
	// The user's email.
	Email *string `pulumi:"email"`
	// Whether the email address was validated or not. Default to `false`.
	EmailVerified *bool `pulumi:"emailVerified"`
	// When false, this user cannot log in. Defaults to `true`.
	Enabled             *bool                   `pulumi:"enabled"`
	FederatedIdentities []UserFederatedIdentity `pulumi:"federatedIdentities"`
	// The user's first name.
	FirstName *string `pulumi:"firstName"`
	// When given, the user's initial password will be set. This attribute is only respected during initial user creation.
	InitialPassword *UserInitialPassword `pulumi:"initialPassword"`
	// The user's last name.
	LastName *string `pulumi:"lastName"`
	// The realm this user belongs to.
	RealmId *string `pulumi:"realmId"`
	// The unique username of this user.
	Username *string `pulumi:"username"`
}

type UserState struct {
	// A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes pulumi.MapInput
	// The user's email.
	Email pulumi.StringPtrInput
	// Whether the email address was validated or not. Default to `false`.
	EmailVerified pulumi.BoolPtrInput
	// When false, this user cannot log in. Defaults to `true`.
	Enabled             pulumi.BoolPtrInput
	FederatedIdentities UserFederatedIdentityArrayInput
	// The user's first name.
	FirstName pulumi.StringPtrInput
	// When given, the user's initial password will be set. This attribute is only respected during initial user creation.
	InitialPassword UserInitialPasswordPtrInput
	// The user's last name.
	LastName pulumi.StringPtrInput
	// The realm this user belongs to.
	RealmId pulumi.StringPtrInput
	// The unique username of this user.
	Username pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes map[string]interface{} `pulumi:"attributes"`
	// The user's email.
	Email *string `pulumi:"email"`
	// Whether the email address was validated or not. Default to `false`.
	EmailVerified *bool `pulumi:"emailVerified"`
	// When false, this user cannot log in. Defaults to `true`.
	Enabled             *bool                   `pulumi:"enabled"`
	FederatedIdentities []UserFederatedIdentity `pulumi:"federatedIdentities"`
	// The user's first name.
	FirstName *string `pulumi:"firstName"`
	// When given, the user's initial password will be set. This attribute is only respected during initial user creation.
	InitialPassword *UserInitialPassword `pulumi:"initialPassword"`
	// The user's last name.
	LastName *string `pulumi:"lastName"`
	// The realm this user belongs to.
	RealmId string `pulumi:"realmId"`
	// The unique username of this user.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
	Attributes pulumi.MapInput
	// The user's email.
	Email pulumi.StringPtrInput
	// Whether the email address was validated or not. Default to `false`.
	EmailVerified pulumi.BoolPtrInput
	// When false, this user cannot log in. Defaults to `true`.
	Enabled             pulumi.BoolPtrInput
	FederatedIdentities UserFederatedIdentityArrayInput
	// The user's first name.
	FirstName pulumi.StringPtrInput
	// When given, the user's initial password will be set. This attribute is only respected during initial user creation.
	InitialPassword UserInitialPasswordPtrInput
	// The user's last name.
	LastName pulumi.StringPtrInput
	// The realm this user belongs to.
	RealmId pulumi.StringInput
	// The unique username of this user.
	Username pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
func (o UserOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v *User) pulumi.MapOutput { return v.Attributes }).(pulumi.MapOutput)
}

// The user's email.
func (o UserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// Whether the email address was validated or not. Default to `false`.
func (o UserOutput) EmailVerified() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.EmailVerified }).(pulumi.BoolPtrOutput)
}

// When false, this user cannot log in. Defaults to `true`.
func (o UserOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o UserOutput) FederatedIdentities() UserFederatedIdentityArrayOutput {
	return o.ApplyT(func(v *User) UserFederatedIdentityArrayOutput { return v.FederatedIdentities }).(UserFederatedIdentityArrayOutput)
}

// The user's first name.
func (o UserOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.FirstName }).(pulumi.StringPtrOutput)
}

// When given, the user's initial password will be set. This attribute is only respected during initial user creation.
func (o UserOutput) InitialPassword() UserInitialPasswordPtrOutput {
	return o.ApplyT(func(v *User) UserInitialPasswordPtrOutput { return v.InitialPassword }).(UserInitialPasswordPtrOutput)
}

// The user's last name.
func (o UserOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.LastName }).(pulumi.StringPtrOutput)
}

// The realm this user belongs to.
func (o UserOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// The unique username of this user.
func (o UserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
