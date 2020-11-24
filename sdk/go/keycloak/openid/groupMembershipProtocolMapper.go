// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows for creating and managing group membership protocol mappers within Keycloak.
//
// Group membership protocol mappers allow you to map a user's group memberships to a claim in a token.
//
// Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
// multiple different clients.
//
// ## Example Usage
// ### Client)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/openid"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("my-realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		openidClient, err := openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
// 			RealmId:    realm.ID(),
// 			ClientId:   pulumi.String("client"),
// 			Enabled:    pulumi.Bool(true),
// 			AccessType: pulumi.String("CONFIDENTIAL"),
// 			ValidRedirectUris: pulumi.StringArray{
// 				pulumi.String("http://localhost:8080/openid-callback"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewGroupMembershipProtocolMapper(ctx, "groupMembershipMapper", &openid.GroupMembershipProtocolMapperArgs{
// 			RealmId:   realm.ID(),
// 			ClientId:  openidClient.ID(),
// 			ClaimName: pulumi.String("groups"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Client Scope)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak"
// 	"github.com/pulumi/pulumi-keycloak/sdk/v3/go/keycloak/openid"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		realm, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
// 			Realm:   pulumi.String("my-realm"),
// 			Enabled: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		clientScope, err := openid.NewClientScope(ctx, "clientScope", &openid.ClientScopeArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewGroupMembershipProtocolMapper(ctx, "groupMembershipMapper", &openid.GroupMembershipProtocolMapperArgs{
// 			RealmId:       realm.ID(),
// 			ClientScopeId: clientScope.ID(),
// 			ClaimName:     pulumi.String("groups"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
//
// ```sh
//  $ pulumi import keycloak:openid/groupMembershipProtocolMapper:GroupMembershipProtocolMapper group_membership_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
//
// ```sh
//  $ pulumi import keycloak:openid/groupMembershipProtocolMapper:GroupMembershipProtocolMapper group_membership_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
// ```
type GroupMembershipProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringOutput `pulumi:"claimName"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
	FullPath pulumi.BoolPtrOutput `pulumi:"fullPath"`
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewGroupMembershipProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewGroupMembershipProtocolMapper(ctx *pulumi.Context,
	name string, args *GroupMembershipProtocolMapperArgs, opts ...pulumi.ResourceOption) (*GroupMembershipProtocolMapper, error) {
	if args == nil || args.ClaimName == nil {
		return nil, errors.New("missing required argument 'ClaimName'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &GroupMembershipProtocolMapperArgs{}
	}
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
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName *string `pulumi:"claimName"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
	FullPath *bool `pulumi:"fullPath"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
}

type GroupMembershipProtocolMapperState struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringPtrInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
	FullPath pulumi.BoolPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
}

func (GroupMembershipProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipProtocolMapperState)(nil)).Elem()
}

type groupMembershipProtocolMapperArgs struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo *bool `pulumi:"addToUserinfo"`
	// The name of the claim to insert into a token.
	ClaimName string `pulumi:"claimName"`
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
	FullPath *bool `pulumi:"fullPath"`
	// The display name of this protocol mapper in the GUI.
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a GroupMembershipProtocolMapper resource.
type GroupMembershipProtocolMapperArgs struct {
	// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
	AddToUserinfo pulumi.BoolPtrInput
	// The name of the claim to insert into a token.
	ClaimName pulumi.StringInput
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// Indicates whether the full path of the group including its parents will be used. Defaults to `true`.
	FullPath pulumi.BoolPtrInput
	// The display name of this protocol mapper in the GUI.
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
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

func (GroupMembershipProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembershipProtocolMapper)(nil)).Elem()
}

func (i GroupMembershipProtocolMapper) ToGroupMembershipProtocolMapperOutput() GroupMembershipProtocolMapperOutput {
	return i.ToGroupMembershipProtocolMapperOutputWithContext(context.Background())
}

func (i GroupMembershipProtocolMapper) ToGroupMembershipProtocolMapperOutputWithContext(ctx context.Context) GroupMembershipProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipProtocolMapperOutput)
}

type GroupMembershipProtocolMapperOutput struct {
	*pulumi.OutputState
}

func (GroupMembershipProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembershipProtocolMapperOutput)(nil)).Elem()
}

func (o GroupMembershipProtocolMapperOutput) ToGroupMembershipProtocolMapperOutput() GroupMembershipProtocolMapperOutput {
	return o
}

func (o GroupMembershipProtocolMapperOutput) ToGroupMembershipProtocolMapperOutputWithContext(ctx context.Context) GroupMembershipProtocolMapperOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupMembershipProtocolMapperOutput{})
}
