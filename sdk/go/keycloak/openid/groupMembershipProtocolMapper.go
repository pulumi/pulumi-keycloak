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

// # openid.GroupMembershipProtocolMapper
//
// Allows for creating and managing group membership protocol mappers within
// Keycloak.
//
// Group membership protocol mappers allow you to map a user's group memberships
// to a claim in a token. Protocol mappers can be defined for a single client,
// or they can be defined for a client scope which can be shared between multiple
// different clients.
//
// ### Example Usage (Client)
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
//			openidClient, err := openid.NewClient(ctx, "openid_client", &openid.ClientArgs{
//				RealmId:    realm.ID(),
//				ClientId:   pulumi.String("test-client"),
//				Name:       pulumi.String("test client"),
//				Enabled:    pulumi.Bool(true),
//				AccessType: pulumi.String("CONFIDENTIAL"),
//				ValidRedirectUris: pulumi.StringArray{
//					pulumi.String("http://localhost:8080/openid-callback"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewGroupMembershipProtocolMapper(ctx, "group_membership_mapper", &openid.GroupMembershipProtocolMapperArgs{
//				RealmId:   realm.ID(),
//				ClientId:  openidClient.ID(),
//				Name:      pulumi.String("group-membership-mapper"),
//				ClaimName: pulumi.String("groups"),
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
// ### Example Usage (Client Scope)
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
//			clientScope, err := openid.NewClientScope(ctx, "client_scope", &openid.ClientScopeArgs{
//				RealmId: realm.ID(),
//				Name:    pulumi.String("test-client-scope"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewGroupMembershipProtocolMapper(ctx, "group_membership_mapper", &openid.GroupMembershipProtocolMapperArgs{
//				RealmId:       realm.ID(),
//				ClientScopeId: clientScope.ID(),
//				Name:          pulumi.String("group-membership-mapper"),
//				ClaimName:     pulumi.String("groups"),
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
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `claimName` - (Required) The name of the claim to insert into a token.
// - `fullPath` - (Optional) Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
// - `addToIdToken` - (Optional) Indicates if the property should be added as a claim to the id token. Defaults to `true`.
// - `addToAccessToken` - (Optional) Indicates if the property should be added as a claim to the access token. Defaults to `true`.
// - `addToUserinfo` - (Optional) Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
//
// ### Import
//
// Protocol mappers can be imported using one of the following formats:
// - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
// - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`
//
// Example:
type GroupMembershipProtocolMapper struct {
	pulumi.CustomResourceState

	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	AddToIdToken     pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	AddToUserinfo    pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	ClaimName        pulumi.StringOutput  `pulumi:"claimName"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	FullPath      pulumi.BoolPtrOutput   `pulumi:"fullPath"`
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewGroupMembershipProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewGroupMembershipProtocolMapper(ctx *pulumi.Context,
	name string, args *GroupMembershipProtocolMapperArgs, opts ...pulumi.ResourceOption) (*GroupMembershipProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClaimName == nil {
		return nil, errors.New("invalid value for required argument 'ClaimName'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupMembershipProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/groupMembershipProtocolMapper:GroupMembershipProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMembershipProtocolMapper gets an existing GroupMembershipProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMembershipProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipProtocolMapperState, opts ...pulumi.ResourceOption) (*GroupMembershipProtocolMapper, error) {
	var resource GroupMembershipProtocolMapper
	err := ctx.ReadResource("keycloak:openid/groupMembershipProtocolMapper:GroupMembershipProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMembershipProtocolMapper resources.
type groupMembershipProtocolMapperState struct {
	AddToAccessToken *bool   `pulumi:"addToAccessToken"`
	AddToIdToken     *bool   `pulumi:"addToIdToken"`
	AddToUserinfo    *bool   `pulumi:"addToUserinfo"`
	ClaimName        *string `pulumi:"claimName"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	FullPath      *bool   `pulumi:"fullPath"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId *string `pulumi:"realmId"`
}

type GroupMembershipProtocolMapperState struct {
	AddToAccessToken pulumi.BoolPtrInput
	AddToIdToken     pulumi.BoolPtrInput
	AddToUserinfo    pulumi.BoolPtrInput
	ClaimName        pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	FullPath      pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringPtrInput
}

func (GroupMembershipProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipProtocolMapperState)(nil)).Elem()
}

type groupMembershipProtocolMapperArgs struct {
	AddToAccessToken *bool  `pulumi:"addToAccessToken"`
	AddToIdToken     *bool  `pulumi:"addToIdToken"`
	AddToUserinfo    *bool  `pulumi:"addToUserinfo"`
	ClaimName        string `pulumi:"claimName"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	FullPath      *bool   `pulumi:"fullPath"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a GroupMembershipProtocolMapper resource.
type GroupMembershipProtocolMapperArgs struct {
	AddToAccessToken pulumi.BoolPtrInput
	AddToIdToken     pulumi.BoolPtrInput
	AddToUserinfo    pulumi.BoolPtrInput
	ClaimName        pulumi.StringInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	FullPath      pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringInput
}

func (GroupMembershipProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipProtocolMapperArgs)(nil)).Elem()
}

type GroupMembershipProtocolMapperInput interface {
	pulumi.Input

	ToGroupMembershipProtocolMapperOutput() GroupMembershipProtocolMapperOutput
	ToGroupMembershipProtocolMapperOutputWithContext(ctx context.Context) GroupMembershipProtocolMapperOutput
}

func (*GroupMembershipProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembershipProtocolMapper)(nil)).Elem()
}

func (i *GroupMembershipProtocolMapper) ToGroupMembershipProtocolMapperOutput() GroupMembershipProtocolMapperOutput {
	return i.ToGroupMembershipProtocolMapperOutputWithContext(context.Background())
}

func (i *GroupMembershipProtocolMapper) ToGroupMembershipProtocolMapperOutputWithContext(ctx context.Context) GroupMembershipProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipProtocolMapperOutput)
}

// GroupMembershipProtocolMapperArrayInput is an input type that accepts GroupMembershipProtocolMapperArray and GroupMembershipProtocolMapperArrayOutput values.
// You can construct a concrete instance of `GroupMembershipProtocolMapperArrayInput` via:
//
//	GroupMembershipProtocolMapperArray{ GroupMembershipProtocolMapperArgs{...} }
type GroupMembershipProtocolMapperArrayInput interface {
	pulumi.Input

	ToGroupMembershipProtocolMapperArrayOutput() GroupMembershipProtocolMapperArrayOutput
	ToGroupMembershipProtocolMapperArrayOutputWithContext(context.Context) GroupMembershipProtocolMapperArrayOutput
}

type GroupMembershipProtocolMapperArray []GroupMembershipProtocolMapperInput

func (GroupMembershipProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembershipProtocolMapper)(nil)).Elem()
}

func (i GroupMembershipProtocolMapperArray) ToGroupMembershipProtocolMapperArrayOutput() GroupMembershipProtocolMapperArrayOutput {
	return i.ToGroupMembershipProtocolMapperArrayOutputWithContext(context.Background())
}

func (i GroupMembershipProtocolMapperArray) ToGroupMembershipProtocolMapperArrayOutputWithContext(ctx context.Context) GroupMembershipProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipProtocolMapperArrayOutput)
}

// GroupMembershipProtocolMapperMapInput is an input type that accepts GroupMembershipProtocolMapperMap and GroupMembershipProtocolMapperMapOutput values.
// You can construct a concrete instance of `GroupMembershipProtocolMapperMapInput` via:
//
//	GroupMembershipProtocolMapperMap{ "key": GroupMembershipProtocolMapperArgs{...} }
type GroupMembershipProtocolMapperMapInput interface {
	pulumi.Input

	ToGroupMembershipProtocolMapperMapOutput() GroupMembershipProtocolMapperMapOutput
	ToGroupMembershipProtocolMapperMapOutputWithContext(context.Context) GroupMembershipProtocolMapperMapOutput
}

type GroupMembershipProtocolMapperMap map[string]GroupMembershipProtocolMapperInput

func (GroupMembershipProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembershipProtocolMapper)(nil)).Elem()
}

func (i GroupMembershipProtocolMapperMap) ToGroupMembershipProtocolMapperMapOutput() GroupMembershipProtocolMapperMapOutput {
	return i.ToGroupMembershipProtocolMapperMapOutputWithContext(context.Background())
}

func (i GroupMembershipProtocolMapperMap) ToGroupMembershipProtocolMapperMapOutputWithContext(ctx context.Context) GroupMembershipProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipProtocolMapperMapOutput)
}

type GroupMembershipProtocolMapperOutput struct{ *pulumi.OutputState }

func (GroupMembershipProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembershipProtocolMapper)(nil)).Elem()
}

func (o GroupMembershipProtocolMapperOutput) ToGroupMembershipProtocolMapperOutput() GroupMembershipProtocolMapperOutput {
	return o
}

func (o GroupMembershipProtocolMapperOutput) ToGroupMembershipProtocolMapperOutputWithContext(ctx context.Context) GroupMembershipProtocolMapperOutput {
	return o
}

func (o GroupMembershipProtocolMapperOutput) AddToAccessToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupMembershipProtocolMapper) pulumi.BoolPtrOutput { return v.AddToAccessToken }).(pulumi.BoolPtrOutput)
}

func (o GroupMembershipProtocolMapperOutput) AddToIdToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupMembershipProtocolMapper) pulumi.BoolPtrOutput { return v.AddToIdToken }).(pulumi.BoolPtrOutput)
}

func (o GroupMembershipProtocolMapperOutput) AddToUserinfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupMembershipProtocolMapper) pulumi.BoolPtrOutput { return v.AddToUserinfo }).(pulumi.BoolPtrOutput)
}

func (o GroupMembershipProtocolMapperOutput) ClaimName() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMembershipProtocolMapper) pulumi.StringOutput { return v.ClaimName }).(pulumi.StringOutput)
}

// The mapper's associated client. Cannot be used at the same time as client_scope_id.
func (o GroupMembershipProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMembershipProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The mapper's associated client scope. Cannot be used at the same time as client_id.
func (o GroupMembershipProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupMembershipProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

func (o GroupMembershipProtocolMapperOutput) FullPath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupMembershipProtocolMapper) pulumi.BoolPtrOutput { return v.FullPath }).(pulumi.BoolPtrOutput)
}

// A human-friendly name that will appear in the Keycloak console.
func (o GroupMembershipProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMembershipProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm id where the associated client or client scope exists.
func (o GroupMembershipProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupMembershipProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type GroupMembershipProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (GroupMembershipProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupMembershipProtocolMapper)(nil)).Elem()
}

func (o GroupMembershipProtocolMapperArrayOutput) ToGroupMembershipProtocolMapperArrayOutput() GroupMembershipProtocolMapperArrayOutput {
	return o
}

func (o GroupMembershipProtocolMapperArrayOutput) ToGroupMembershipProtocolMapperArrayOutputWithContext(ctx context.Context) GroupMembershipProtocolMapperArrayOutput {
	return o
}

func (o GroupMembershipProtocolMapperArrayOutput) Index(i pulumi.IntInput) GroupMembershipProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupMembershipProtocolMapper {
		return vs[0].([]*GroupMembershipProtocolMapper)[vs[1].(int)]
	}).(GroupMembershipProtocolMapperOutput)
}

type GroupMembershipProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (GroupMembershipProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupMembershipProtocolMapper)(nil)).Elem()
}

func (o GroupMembershipProtocolMapperMapOutput) ToGroupMembershipProtocolMapperMapOutput() GroupMembershipProtocolMapperMapOutput {
	return o
}

func (o GroupMembershipProtocolMapperMapOutput) ToGroupMembershipProtocolMapperMapOutputWithContext(ctx context.Context) GroupMembershipProtocolMapperMapOutput {
	return o
}

func (o GroupMembershipProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) GroupMembershipProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupMembershipProtocolMapper {
		return vs[0].(map[string]*GroupMembershipProtocolMapper)[vs[1].(string)]
	}).(GroupMembershipProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipProtocolMapperInput)(nil)).Elem(), &GroupMembershipProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipProtocolMapperArrayInput)(nil)).Elem(), GroupMembershipProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMembershipProtocolMapperMapInput)(nil)).Elem(), GroupMembershipProtocolMapperMap{})
	pulumi.RegisterOutputType(GroupMembershipProtocolMapperOutput{})
	pulumi.RegisterOutputType(GroupMembershipProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(GroupMembershipProtocolMapperMapOutput{})
}
