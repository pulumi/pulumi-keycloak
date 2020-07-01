// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # openid.FullNameProtocolMapper
//
// Allows for creating and managing full name protocol mappers within
// Keycloak.
//
// Full name protocol mappers allow you to map a user's first and last name
// to the OpenID Connect `name` claim in a token. Protocol mappers can be defined
// for a single client, or they can be defined for a client scope which can
// be shared between multiple different clients.
//
// ### Example Usage (Client)
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
// 		openidClient, err := openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
// 			AccessType: pulumi.String("CONFIDENTIAL"),
// 			ClientId:   pulumi.String("test-client"),
// 			Enabled:    pulumi.Bool(true),
// 			RealmId:    realm.ID(),
// 			ValidRedirectUris: pulumi.StringArray{
// 				pulumi.String("http://localhost:8080/openid-callback"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewFullNameProtocolMapper(ctx, "fullNameMapper", &openid.FullNameProtocolMapperArgs{
// 			ClientId: openidClient.ID(),
// 			RealmId:  realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Example Usage (Client Scope)
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
// 		clientScope, err := openid.NewClientScope(ctx, "clientScope", &openid.ClientScopeArgs{
// 			RealmId: realm.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = openid.NewFullNameProtocolMapper(ctx, "fullNameMapper", &openid.FullNameProtocolMapperArgs{
// 			ClientScopeId: clientScope.ID(),
// 			RealmId:       realm.ID(),
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
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `addToIdToken` - (Optional) Indicates if the user's full name should be added as a claim to the id token. Defaults to `true`.
// - `addToAccessToken` - (Optional) Indicates if the user's full name should be added as a claim to the access token. Defaults to `true`.
// - `addToUserinfo` - (Optional) Indicates if the user's full name should be added as a claim to the UserInfo response body. Defaults to `true`.
type FullNameProtocolMapper struct {
	pulumi.CustomResourceState

	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	AddToIdToken     pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	AddToUserinfo    pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewFullNameProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewFullNameProtocolMapper(ctx *pulumi.Context,
	name string, args *FullNameProtocolMapperArgs, opts ...pulumi.ResourceOption) (*FullNameProtocolMapper, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &FullNameProtocolMapperArgs{}
	}
	var resource FullNameProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/fullNameProtocolMapper:FullNameProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFullNameProtocolMapper gets an existing FullNameProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFullNameProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FullNameProtocolMapperState, opts ...pulumi.ResourceOption) (*FullNameProtocolMapper, error) {
	var resource FullNameProtocolMapper
	err := ctx.ReadResource("keycloak:openid/fullNameProtocolMapper:FullNameProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FullNameProtocolMapper resources.
type fullNameProtocolMapperState struct {
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	AddToIdToken     *bool `pulumi:"addToIdToken"`
	AddToUserinfo    *bool `pulumi:"addToUserinfo"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId *string `pulumi:"realmId"`
}

type FullNameProtocolMapperState struct {
	AddToAccessToken pulumi.BoolPtrInput
	AddToIdToken     pulumi.BoolPtrInput
	AddToUserinfo    pulumi.BoolPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringPtrInput
}

func (FullNameProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*fullNameProtocolMapperState)(nil)).Elem()
}

type fullNameProtocolMapperArgs struct {
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	AddToIdToken     *bool `pulumi:"addToIdToken"`
	AddToUserinfo    *bool `pulumi:"addToUserinfo"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a FullNameProtocolMapper resource.
type FullNameProtocolMapperArgs struct {
	AddToAccessToken pulumi.BoolPtrInput
	AddToIdToken     pulumi.BoolPtrInput
	AddToUserinfo    pulumi.BoolPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringInput
}

func (FullNameProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fullNameProtocolMapperArgs)(nil)).Elem()
}
