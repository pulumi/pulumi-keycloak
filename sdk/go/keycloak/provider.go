// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the keycloak package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	BasePath     pulumi.StringPtrOutput `pulumi:"basePath"`
	ClientId     pulumi.StringOutput    `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	Password     pulumi.StringPtrOutput `pulumi:"password"`
	Realm        pulumi.StringPtrOutput `pulumi:"realm"`
	// Allows x509 calls using an unknown CA certificate (for development purposes)
	RootCaCertificate pulumi.StringPtrOutput `pulumi:"rootCaCertificate"`
	// The base URL of the Keycloak instance, before `/auth`
	Url      pulumi.StringOutput    `pulumi:"url"`
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if isZero(args.ClientTimeout) {
		args.ClientTimeout = pulumi.IntPtr(getEnvOrDefault(5, parseEnvInt, "KEYCLOAK_CLIENT_TIMEOUT").(int))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:keycloak", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	AdditionalHeaders map[string]string `pulumi:"additionalHeaders"`
	BasePath          *string           `pulumi:"basePath"`
	ClientId          string            `pulumi:"clientId"`
	ClientSecret      *string           `pulumi:"clientSecret"`
	// Timeout (in seconds) of the Keycloak client
	ClientTimeout *int `pulumi:"clientTimeout"`
	// Whether or not to login to Keycloak instance on provider initialization
	InitialLogin *bool   `pulumi:"initialLogin"`
	Password     *string `pulumi:"password"`
	Realm        *string `pulumi:"realm"`
	// Allows x509 calls using an unknown CA certificate (for development purposes)
	RootCaCertificate *string `pulumi:"rootCaCertificate"`
	// Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
	// should be avoided.
	TlsInsecureSkipVerify *bool `pulumi:"tlsInsecureSkipVerify"`
	// The base URL of the Keycloak instance, before `/auth`
	Url      string  `pulumi:"url"`
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	AdditionalHeaders pulumi.StringMapInput
	BasePath          pulumi.StringPtrInput
	ClientId          pulumi.StringInput
	ClientSecret      pulumi.StringPtrInput
	// Timeout (in seconds) of the Keycloak client
	ClientTimeout pulumi.IntPtrInput
	// Whether or not to login to Keycloak instance on provider initialization
	InitialLogin pulumi.BoolPtrInput
	Password     pulumi.StringPtrInput
	Realm        pulumi.StringPtrInput
	// Allows x509 calls using an unknown CA certificate (for development purposes)
	RootCaCertificate pulumi.StringPtrInput
	// Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
	// should be avoided.
	TlsInsecureSkipVerify pulumi.BoolPtrInput
	// The base URL of the Keycloak instance, before `/auth`
	Url      pulumi.StringInput
	Username pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
