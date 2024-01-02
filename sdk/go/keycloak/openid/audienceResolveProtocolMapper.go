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

// Allows for creating the "Audience Resolve" OIDC protocol mapper within Keycloak.
//
// This protocol mapper is useful to avoid manual management of audiences, instead relying on the presence of client roles
// to imply which audiences are appropriate for the token. See the
// [Keycloak docs](https://www.keycloak.org/docs/latest/server_admin/#_audience_resolve) for more details.
//
// ## Example Usage
// ### Client)
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
//			openidClient, err := openid.NewClient(ctx, "openidClient", &openid.ClientArgs{
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
//			_, err = openid.NewAudienceResolveProtocolMapper(ctx, "audienceMapper", &openid.AudienceResolveProtocolMapperArgs{
//				RealmId:  realm.ID(),
//				ClientId: openidClient.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Client Scope)
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
//			clientScope, err := openid.NewClientScope(ctx, "clientScope", &openid.ClientScopeArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = openid.NewAudienceProtocolMapper(ctx, "audienceMapper", &openid.AudienceProtocolMapperArgs{
//				RealmId:       realm.ID(),
//				ClientScopeId: clientScope.ID(),
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
// Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
//
// ```sh
//
//	$ pulumi import keycloak:openid/audienceResolveProtocolMapper:AudienceResolveProtocolMapper audience_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
//
// ```sh
//
//	$ pulumi import keycloak:openid/audienceResolveProtocolMapper:AudienceResolveProtocolMapper audience_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
//
// ```
type AudienceResolveProtocolMapper struct {
	pulumi.CustomResourceState

	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewAudienceResolveProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewAudienceResolveProtocolMapper(ctx *pulumi.Context,
	name string, args *AudienceResolveProtocolMapperArgs, opts ...pulumi.ResourceOption) (*AudienceResolveProtocolMapper, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AudienceResolveProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/audienceResolveProtocolMapper:AudienceResolveProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAudienceResolveProtocolMapper gets an existing AudienceResolveProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAudienceResolveProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AudienceResolveProtocolMapperState, opts ...pulumi.ResourceOption) (*AudienceResolveProtocolMapper, error) {
	var resource AudienceResolveProtocolMapper
	err := ctx.ReadResource("keycloak:openid/audienceResolveProtocolMapper:AudienceResolveProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AudienceResolveProtocolMapper resources.
type audienceResolveProtocolMapperState struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId *string `pulumi:"realmId"`
}

type AudienceResolveProtocolMapperState struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringPtrInput
}

func (AudienceResolveProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*audienceResolveProtocolMapperState)(nil)).Elem()
}

type audienceResolveProtocolMapperArgs struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId *string `pulumi:"clientId"`
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
	Name *string `pulumi:"name"`
	// The realm this protocol mapper exists within.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a AudienceResolveProtocolMapper resource.
type AudienceResolveProtocolMapperArgs struct {
	// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
	ClientId pulumi.StringPtrInput
	// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
	ClientScopeId pulumi.StringPtrInput
	// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
	Name pulumi.StringPtrInput
	// The realm this protocol mapper exists within.
	RealmId pulumi.StringInput
}

func (AudienceResolveProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*audienceResolveProtocolMapperArgs)(nil)).Elem()
}

type AudienceResolveProtocolMapperInput interface {
	pulumi.Input

	ToAudienceResolveProtocolMapperOutput() AudienceResolveProtocolMapperOutput
	ToAudienceResolveProtocolMapperOutputWithContext(ctx context.Context) AudienceResolveProtocolMapperOutput
}

func (*AudienceResolveProtocolMapper) ElementType() reflect.Type {
	return reflect.TypeOf((**AudienceResolveProtocolMapper)(nil)).Elem()
}

func (i *AudienceResolveProtocolMapper) ToAudienceResolveProtocolMapperOutput() AudienceResolveProtocolMapperOutput {
	return i.ToAudienceResolveProtocolMapperOutputWithContext(context.Background())
}

func (i *AudienceResolveProtocolMapper) ToAudienceResolveProtocolMapperOutputWithContext(ctx context.Context) AudienceResolveProtocolMapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudienceResolveProtocolMapperOutput)
}

// AudienceResolveProtocolMapperArrayInput is an input type that accepts AudienceResolveProtocolMapperArray and AudienceResolveProtocolMapperArrayOutput values.
// You can construct a concrete instance of `AudienceResolveProtocolMapperArrayInput` via:
//
//	AudienceResolveProtocolMapperArray{ AudienceResolveProtocolMapperArgs{...} }
type AudienceResolveProtocolMapperArrayInput interface {
	pulumi.Input

	ToAudienceResolveProtocolMapperArrayOutput() AudienceResolveProtocolMapperArrayOutput
	ToAudienceResolveProtocolMapperArrayOutputWithContext(context.Context) AudienceResolveProtocolMapperArrayOutput
}

type AudienceResolveProtocolMapperArray []AudienceResolveProtocolMapperInput

func (AudienceResolveProtocolMapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AudienceResolveProtocolMapper)(nil)).Elem()
}

func (i AudienceResolveProtocolMapperArray) ToAudienceResolveProtocolMapperArrayOutput() AudienceResolveProtocolMapperArrayOutput {
	return i.ToAudienceResolveProtocolMapperArrayOutputWithContext(context.Background())
}

func (i AudienceResolveProtocolMapperArray) ToAudienceResolveProtocolMapperArrayOutputWithContext(ctx context.Context) AudienceResolveProtocolMapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudienceResolveProtocolMapperArrayOutput)
}

// AudienceResolveProtocolMapperMapInput is an input type that accepts AudienceResolveProtocolMapperMap and AudienceResolveProtocolMapperMapOutput values.
// You can construct a concrete instance of `AudienceResolveProtocolMapperMapInput` via:
//
//	AudienceResolveProtocolMapperMap{ "key": AudienceResolveProtocolMapperArgs{...} }
type AudienceResolveProtocolMapperMapInput interface {
	pulumi.Input

	ToAudienceResolveProtocolMapperMapOutput() AudienceResolveProtocolMapperMapOutput
	ToAudienceResolveProtocolMapperMapOutputWithContext(context.Context) AudienceResolveProtocolMapperMapOutput
}

type AudienceResolveProtocolMapperMap map[string]AudienceResolveProtocolMapperInput

func (AudienceResolveProtocolMapperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AudienceResolveProtocolMapper)(nil)).Elem()
}

func (i AudienceResolveProtocolMapperMap) ToAudienceResolveProtocolMapperMapOutput() AudienceResolveProtocolMapperMapOutput {
	return i.ToAudienceResolveProtocolMapperMapOutputWithContext(context.Background())
}

func (i AudienceResolveProtocolMapperMap) ToAudienceResolveProtocolMapperMapOutputWithContext(ctx context.Context) AudienceResolveProtocolMapperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AudienceResolveProtocolMapperMapOutput)
}

type AudienceResolveProtocolMapperOutput struct{ *pulumi.OutputState }

func (AudienceResolveProtocolMapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AudienceResolveProtocolMapper)(nil)).Elem()
}

func (o AudienceResolveProtocolMapperOutput) ToAudienceResolveProtocolMapperOutput() AudienceResolveProtocolMapperOutput {
	return o
}

func (o AudienceResolveProtocolMapperOutput) ToAudienceResolveProtocolMapperOutputWithContext(ctx context.Context) AudienceResolveProtocolMapperOutput {
	return o
}

// The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
func (o AudienceResolveProtocolMapperOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AudienceResolveProtocolMapper) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
func (o AudienceResolveProtocolMapperOutput) ClientScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AudienceResolveProtocolMapper) pulumi.StringPtrOutput { return v.ClientScopeId }).(pulumi.StringPtrOutput)
}

// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
func (o AudienceResolveProtocolMapperOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AudienceResolveProtocolMapper) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The realm this protocol mapper exists within.
func (o AudienceResolveProtocolMapperOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *AudienceResolveProtocolMapper) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type AudienceResolveProtocolMapperArrayOutput struct{ *pulumi.OutputState }

func (AudienceResolveProtocolMapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AudienceResolveProtocolMapper)(nil)).Elem()
}

func (o AudienceResolveProtocolMapperArrayOutput) ToAudienceResolveProtocolMapperArrayOutput() AudienceResolveProtocolMapperArrayOutput {
	return o
}

func (o AudienceResolveProtocolMapperArrayOutput) ToAudienceResolveProtocolMapperArrayOutputWithContext(ctx context.Context) AudienceResolveProtocolMapperArrayOutput {
	return o
}

func (o AudienceResolveProtocolMapperArrayOutput) Index(i pulumi.IntInput) AudienceResolveProtocolMapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AudienceResolveProtocolMapper {
		return vs[0].([]*AudienceResolveProtocolMapper)[vs[1].(int)]
	}).(AudienceResolveProtocolMapperOutput)
}

type AudienceResolveProtocolMapperMapOutput struct{ *pulumi.OutputState }

func (AudienceResolveProtocolMapperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AudienceResolveProtocolMapper)(nil)).Elem()
}

func (o AudienceResolveProtocolMapperMapOutput) ToAudienceResolveProtocolMapperMapOutput() AudienceResolveProtocolMapperMapOutput {
	return o
}

func (o AudienceResolveProtocolMapperMapOutput) ToAudienceResolveProtocolMapperMapOutputWithContext(ctx context.Context) AudienceResolveProtocolMapperMapOutput {
	return o
}

func (o AudienceResolveProtocolMapperMapOutput) MapIndex(k pulumi.StringInput) AudienceResolveProtocolMapperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AudienceResolveProtocolMapper {
		return vs[0].(map[string]*AudienceResolveProtocolMapper)[vs[1].(string)]
	}).(AudienceResolveProtocolMapperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AudienceResolveProtocolMapperInput)(nil)).Elem(), &AudienceResolveProtocolMapper{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudienceResolveProtocolMapperArrayInput)(nil)).Elem(), AudienceResolveProtocolMapperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AudienceResolveProtocolMapperMapInput)(nil)).Elem(), AudienceResolveProtocolMapperMap{})
	pulumi.RegisterOutputType(AudienceResolveProtocolMapperOutput{})
	pulumi.RegisterOutputType(AudienceResolveProtocolMapperArrayOutput{})
	pulumi.RegisterOutputType(AudienceResolveProtocolMapperMapOutput{})
}
