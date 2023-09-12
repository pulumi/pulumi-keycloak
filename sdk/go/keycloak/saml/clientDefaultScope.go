// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package saml

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/saml"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
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
//			_, err = saml.NewClient(ctx, "samlClient", &saml.ClientArgs{
//				RealmId:               realm.ID(),
//				ClientId:              pulumi.String("saml-client"),
//				SignDocuments:         pulumi.Bool(false),
//				SignAssertions:        pulumi.Bool(true),
//				IncludeAuthnStatement: pulumi.Bool(true),
//				SigningCertificate:    readFileOrPanic("saml-cert.pem"),
//				SigningPrivateKey:     readFileOrPanic("saml-key.pem"),
//			})
//			if err != nil {
//				return err
//			}
//			clientScope, err := saml.NewClientScope(ctx, "clientScope", &saml.ClientScopeArgs{
//				RealmId: realm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = saml.NewClientDefaultScope(ctx, "clientDefaultScopes", &saml.ClientDefaultScopeArgs{
//				RealmId:  realm.ID(),
//				ClientId: pulumi.Any(keycloak_saml_client.Client.Id),
//				DefaultScopes: pulumi.StringArray{
//					pulumi.String("role_list"),
//					clientScope.Name,
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
// This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server.
type ClientDefaultScope struct {
	pulumi.CustomResourceState

	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// An array of client scope names to attach to this client.
	DefaultScopes pulumi.StringArrayOutput `pulumi:"defaultScopes"`
	// The realm this client and scopes exists in.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewClientDefaultScope registers a new resource with the given unique name, arguments, and options.
func NewClientDefaultScope(ctx *pulumi.Context,
	name string, args *ClientDefaultScopeArgs, opts ...pulumi.ResourceOption) (*ClientDefaultScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.DefaultScopes == nil {
		return nil, errors.New("invalid value for required argument 'DefaultScopes'")
	}
	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClientDefaultScope
	err := ctx.RegisterResource("keycloak:saml/clientDefaultScope:ClientDefaultScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientDefaultScope gets an existing ClientDefaultScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientDefaultScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientDefaultScopeState, opts ...pulumi.ResourceOption) (*ClientDefaultScope, error) {
	var resource ClientDefaultScope
	err := ctx.ReadResource("keycloak:saml/clientDefaultScope:ClientDefaultScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientDefaultScope resources.
type clientDefaultScopeState struct {
	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId *string `pulumi:"clientId"`
	// An array of client scope names to attach to this client.
	DefaultScopes []string `pulumi:"defaultScopes"`
	// The realm this client and scopes exists in.
	RealmId *string `pulumi:"realmId"`
}

type ClientDefaultScopeState struct {
	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId pulumi.StringPtrInput
	// An array of client scope names to attach to this client.
	DefaultScopes pulumi.StringArrayInput
	// The realm this client and scopes exists in.
	RealmId pulumi.StringPtrInput
}

func (ClientDefaultScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientDefaultScopeState)(nil)).Elem()
}

type clientDefaultScopeArgs struct {
	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId string `pulumi:"clientId"`
	// An array of client scope names to attach to this client.
	DefaultScopes []string `pulumi:"defaultScopes"`
	// The realm this client and scopes exists in.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a ClientDefaultScope resource.
type ClientDefaultScopeArgs struct {
	// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
	ClientId pulumi.StringInput
	// An array of client scope names to attach to this client.
	DefaultScopes pulumi.StringArrayInput
	// The realm this client and scopes exists in.
	RealmId pulumi.StringInput
}

func (ClientDefaultScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientDefaultScopeArgs)(nil)).Elem()
}

type ClientDefaultScopeInput interface {
	pulumi.Input

	ToClientDefaultScopeOutput() ClientDefaultScopeOutput
	ToClientDefaultScopeOutputWithContext(ctx context.Context) ClientDefaultScopeOutput
}

func (*ClientDefaultScope) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientDefaultScope)(nil)).Elem()
}

func (i *ClientDefaultScope) ToClientDefaultScopeOutput() ClientDefaultScopeOutput {
	return i.ToClientDefaultScopeOutputWithContext(context.Background())
}

func (i *ClientDefaultScope) ToClientDefaultScopeOutputWithContext(ctx context.Context) ClientDefaultScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientDefaultScopeOutput)
}

func (i *ClientDefaultScope) ToOutput(ctx context.Context) pulumix.Output[*ClientDefaultScope] {
	return pulumix.Output[*ClientDefaultScope]{
		OutputState: i.ToClientDefaultScopeOutputWithContext(ctx).OutputState,
	}
}

// ClientDefaultScopeArrayInput is an input type that accepts ClientDefaultScopeArray and ClientDefaultScopeArrayOutput values.
// You can construct a concrete instance of `ClientDefaultScopeArrayInput` via:
//
//	ClientDefaultScopeArray{ ClientDefaultScopeArgs{...} }
type ClientDefaultScopeArrayInput interface {
	pulumi.Input

	ToClientDefaultScopeArrayOutput() ClientDefaultScopeArrayOutput
	ToClientDefaultScopeArrayOutputWithContext(context.Context) ClientDefaultScopeArrayOutput
}

type ClientDefaultScopeArray []ClientDefaultScopeInput

func (ClientDefaultScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientDefaultScope)(nil)).Elem()
}

func (i ClientDefaultScopeArray) ToClientDefaultScopeArrayOutput() ClientDefaultScopeArrayOutput {
	return i.ToClientDefaultScopeArrayOutputWithContext(context.Background())
}

func (i ClientDefaultScopeArray) ToClientDefaultScopeArrayOutputWithContext(ctx context.Context) ClientDefaultScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientDefaultScopeArrayOutput)
}

func (i ClientDefaultScopeArray) ToOutput(ctx context.Context) pulumix.Output[[]*ClientDefaultScope] {
	return pulumix.Output[[]*ClientDefaultScope]{
		OutputState: i.ToClientDefaultScopeArrayOutputWithContext(ctx).OutputState,
	}
}

// ClientDefaultScopeMapInput is an input type that accepts ClientDefaultScopeMap and ClientDefaultScopeMapOutput values.
// You can construct a concrete instance of `ClientDefaultScopeMapInput` via:
//
//	ClientDefaultScopeMap{ "key": ClientDefaultScopeArgs{...} }
type ClientDefaultScopeMapInput interface {
	pulumi.Input

	ToClientDefaultScopeMapOutput() ClientDefaultScopeMapOutput
	ToClientDefaultScopeMapOutputWithContext(context.Context) ClientDefaultScopeMapOutput
}

type ClientDefaultScopeMap map[string]ClientDefaultScopeInput

func (ClientDefaultScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientDefaultScope)(nil)).Elem()
}

func (i ClientDefaultScopeMap) ToClientDefaultScopeMapOutput() ClientDefaultScopeMapOutput {
	return i.ToClientDefaultScopeMapOutputWithContext(context.Background())
}

func (i ClientDefaultScopeMap) ToClientDefaultScopeMapOutputWithContext(ctx context.Context) ClientDefaultScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientDefaultScopeMapOutput)
}

func (i ClientDefaultScopeMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClientDefaultScope] {
	return pulumix.Output[map[string]*ClientDefaultScope]{
		OutputState: i.ToClientDefaultScopeMapOutputWithContext(ctx).OutputState,
	}
}

type ClientDefaultScopeOutput struct{ *pulumi.OutputState }

func (ClientDefaultScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientDefaultScope)(nil)).Elem()
}

func (o ClientDefaultScopeOutput) ToClientDefaultScopeOutput() ClientDefaultScopeOutput {
	return o
}

func (o ClientDefaultScopeOutput) ToClientDefaultScopeOutputWithContext(ctx context.Context) ClientDefaultScopeOutput {
	return o
}

func (o ClientDefaultScopeOutput) ToOutput(ctx context.Context) pulumix.Output[*ClientDefaultScope] {
	return pulumix.Output[*ClientDefaultScope]{
		OutputState: o.OutputState,
	}
}

// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
func (o ClientDefaultScopeOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientDefaultScope) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// An array of client scope names to attach to this client.
func (o ClientDefaultScopeOutput) DefaultScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClientDefaultScope) pulumi.StringArrayOutput { return v.DefaultScopes }).(pulumi.StringArrayOutput)
}

// The realm this client and scopes exists in.
func (o ClientDefaultScopeOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientDefaultScope) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

type ClientDefaultScopeArrayOutput struct{ *pulumi.OutputState }

func (ClientDefaultScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClientDefaultScope)(nil)).Elem()
}

func (o ClientDefaultScopeArrayOutput) ToClientDefaultScopeArrayOutput() ClientDefaultScopeArrayOutput {
	return o
}

func (o ClientDefaultScopeArrayOutput) ToClientDefaultScopeArrayOutputWithContext(ctx context.Context) ClientDefaultScopeArrayOutput {
	return o
}

func (o ClientDefaultScopeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ClientDefaultScope] {
	return pulumix.Output[[]*ClientDefaultScope]{
		OutputState: o.OutputState,
	}
}

func (o ClientDefaultScopeArrayOutput) Index(i pulumi.IntInput) ClientDefaultScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClientDefaultScope {
		return vs[0].([]*ClientDefaultScope)[vs[1].(int)]
	}).(ClientDefaultScopeOutput)
}

type ClientDefaultScopeMapOutput struct{ *pulumi.OutputState }

func (ClientDefaultScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClientDefaultScope)(nil)).Elem()
}

func (o ClientDefaultScopeMapOutput) ToClientDefaultScopeMapOutput() ClientDefaultScopeMapOutput {
	return o
}

func (o ClientDefaultScopeMapOutput) ToClientDefaultScopeMapOutputWithContext(ctx context.Context) ClientDefaultScopeMapOutput {
	return o
}

func (o ClientDefaultScopeMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ClientDefaultScope] {
	return pulumix.Output[map[string]*ClientDefaultScope]{
		OutputState: o.OutputState,
	}
}

func (o ClientDefaultScopeMapOutput) MapIndex(k pulumi.StringInput) ClientDefaultScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClientDefaultScope {
		return vs[0].(map[string]*ClientDefaultScope)[vs[1].(string)]
	}).(ClientDefaultScopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientDefaultScopeInput)(nil)).Elem(), &ClientDefaultScope{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientDefaultScopeArrayInput)(nil)).Elem(), ClientDefaultScopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientDefaultScopeMapInput)(nil)).Elem(), ClientDefaultScopeMap{})
	pulumi.RegisterOutputType(ClientDefaultScopeOutput{})
	pulumi.RegisterOutputType(ClientDefaultScopeArrayOutput{})
	pulumi.RegisterOutputType(ClientDefaultScopeMapOutput{})
}
