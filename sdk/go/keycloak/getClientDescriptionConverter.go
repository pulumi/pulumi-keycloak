// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source uses the [ClientDescriptionConverter](https://www.keycloak.org/docs-api/6.0/javadocs/org/keycloak/exportimport/ClientDescriptionConverter.html) API to convert a generic client description into a Keycloak
// client. This data can then be used to manage the client within Keycloak.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
//	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak/saml"
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
//			samlClient := keycloak.GetClientDescriptionConverterOutput(ctx, keycloak.GetClientDescriptionConverterOutputArgs{
//				RealmId: realm.ID(),
//				Body: pulumi.String(`	<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" validUntil="2021-04-17T12:41:46Z" cacheDuration="PT604800S" entityID="FakeEntityId">
//	    <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
//	        <md:KeyDescriptor use="signing">
//				<ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
//					<ds:X509Data>
//						<ds:X509Certificate>MIICyDCCAjGgAwIBAgIBADANBgkqhkiG9w0BAQ0FADCBgDELMAkGA1UEBhMCdXMx
//						CzAJBgNVBAgMAklBMSQwIgYDVQQKDBt0ZXJyYWZvcm0tcHJvdmlkZXIta2V5Y2xv
//						YWsxHDAaBgNVBAMME21ycGFya2Vycy5naXRodWIuaW8xIDAeBgkqhkiG9w0BCQEW
//						EW1pY2hhZWxAcGFya2VyLmdnMB4XDTE5MDEwODE0NDYzNloXDTI5MDEwNTE0NDYz
//						NlowgYAxCzAJBgNVBAYTAnVzMQswCQYDVQQIDAJJQTEkMCIGA1UECgwbdGVycmFm
//						b3JtLXByb3ZpZGVyLWtleWNsb2FrMRwwGgYDVQQDDBNtcnBhcmtlcnMuZ2l0aHVi
//						LmlvMSAwHgYJKoZIhvcNAQkBFhFtaWNoYWVsQHBhcmtlci5nZzCBnzANBgkqhkiG
//						9w0BAQEFAAOBjQAwgYkCgYEAxuZny7uyYxGVPtpie14gNQC4tT9sAvO2sVNDhuoe
//						qIKLRpNwkHnwQmwe5OxSh9K0BPHp/DNuuVWUqvo4tniEYn3jBr7FwLYLTKojQIxj
//						53S1UTT9EXq3eP5HsHMD0QnTuca2nlNYUDBm6ud2fQj0Jt5qLx86EbEC28N56IRv
//						GX8CAwEAAaNQME4wHQYDVR0OBBYEFMLnbQh77j7vhGTpAhKpDhCrBsPZMB8GA1Ud
//						IwQYMBaAFMLnbQh77j7vhGTpAhKpDhCrBsPZMAwGA1UdEwQFMAMBAf8wDQYJKoZI
//						hvcNAQENBQADgYEAB8wGrAQY0pAfwbnYSyBt4STbebeRTu1/q1ucfrtc3qsegcd5
//						n01xTR+T2uZJwqHFPpFjr4IPORiHx3+4BWCweslPD53qBjKUPXcbMO1Revjef6Tj
//						K3K0AuJ94fxgXVoT61Nzu/a6Lj6RhzU/Dao9mlSbJY+YSbm+ZBpsuRUQ84s=</ds:X509Certificate>
//					</ds:X509Data>
//				</ds:KeyInfo>
//			</md:KeyDescriptor>
//			<md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
//	        <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://localhost/acs/saml/" index="1"/>
//	    </md:SPSSODescriptor>
//
// </md:EntityDescriptor>
// `),
//
//			}, nil)
//			_, err = saml.NewClient(ctx, "saml_client", &saml.ClientArgs{
//				RealmId: realm.ID(),
//				ClientId: pulumi.String(samlClient.ApplyT(func(samlClient keycloak.GetClientDescriptionConverterResult) (*string, error) {
//					return &samlClient.ClientId, nil
//				}).(pulumi.StringPtrOutput)),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetClientDescriptionConverter(ctx *pulumi.Context, args *GetClientDescriptionConverterArgs, opts ...pulumi.InvokeOption) (*GetClientDescriptionConverterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetClientDescriptionConverterResult
	err := ctx.Invoke("keycloak:index/getClientDescriptionConverter:getClientDescriptionConverter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClientDescriptionConverter.
type GetClientDescriptionConverterArgs struct {
	// The body of the request to convert.
	Body string `pulumi:"body"`
	// The realm to use for the client description converter API call.
	RealmId string `pulumi:"realmId"`
}

// A collection of values returned by getClientDescriptionConverter.
type GetClientDescriptionConverterResult struct {
	Access                             map[string]string `pulumi:"access"`
	AdminUrl                           string            `pulumi:"adminUrl"`
	Attributes                         map[string]string `pulumi:"attributes"`
	AuthenticationFlowBindingOverrides map[string]string `pulumi:"authenticationFlowBindingOverrides"`
	AuthorizationServicesEnabled       bool              `pulumi:"authorizationServicesEnabled"`
	AuthorizationSettings              map[string]string `pulumi:"authorizationSettings"`
	BaseUrl                            string            `pulumi:"baseUrl"`
	BearerOnly                         bool              `pulumi:"bearerOnly"`
	Body                               string            `pulumi:"body"`
	ClientAuthenticatorType            string            `pulumi:"clientAuthenticatorType"`
	ClientId                           string            `pulumi:"clientId"`
	ConsentRequired                    string            `pulumi:"consentRequired"`
	DefaultClientScopes                []string          `pulumi:"defaultClientScopes"`
	DefaultRoles                       []string          `pulumi:"defaultRoles"`
	Description                        string            `pulumi:"description"`
	DirectAccessGrantsEnabled          bool              `pulumi:"directAccessGrantsEnabled"`
	Enabled                            bool              `pulumi:"enabled"`
	FrontchannelLogout                 bool              `pulumi:"frontchannelLogout"`
	FullScopeAllowed                   bool              `pulumi:"fullScopeAllowed"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string                                        `pulumi:"id"`
	ImplicitFlowEnabled     bool                                          `pulumi:"implicitFlowEnabled"`
	Name                    string                                        `pulumi:"name"`
	NotBefore               int                                           `pulumi:"notBefore"`
	OptionalClientScopes    []string                                      `pulumi:"optionalClientScopes"`
	Origin                  string                                        `pulumi:"origin"`
	Protocol                string                                        `pulumi:"protocol"`
	ProtocolMappers         []GetClientDescriptionConverterProtocolMapper `pulumi:"protocolMappers"`
	PublicClient            bool                                          `pulumi:"publicClient"`
	RealmId                 string                                        `pulumi:"realmId"`
	RedirectUris            []string                                      `pulumi:"redirectUris"`
	RegisteredNodes         map[string]string                             `pulumi:"registeredNodes"`
	RegistrationAccessToken string                                        `pulumi:"registrationAccessToken"`
	RootUrl                 string                                        `pulumi:"rootUrl"`
	Secret                  string                                        `pulumi:"secret"`
	ServiceAccountsEnabled  bool                                          `pulumi:"serviceAccountsEnabled"`
	StandardFlowEnabled     bool                                          `pulumi:"standardFlowEnabled"`
	SurrogateAuthRequired   bool                                          `pulumi:"surrogateAuthRequired"`
	WebOrigins              []string                                      `pulumi:"webOrigins"`
}

func GetClientDescriptionConverterOutput(ctx *pulumi.Context, args GetClientDescriptionConverterOutputArgs, opts ...pulumi.InvokeOption) GetClientDescriptionConverterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClientDescriptionConverterResult, error) {
			args := v.(GetClientDescriptionConverterArgs)
			r, err := GetClientDescriptionConverter(ctx, &args, opts...)
			var s GetClientDescriptionConverterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClientDescriptionConverterResultOutput)
}

// A collection of arguments for invoking getClientDescriptionConverter.
type GetClientDescriptionConverterOutputArgs struct {
	// The body of the request to convert.
	Body pulumi.StringInput `pulumi:"body"`
	// The realm to use for the client description converter API call.
	RealmId pulumi.StringInput `pulumi:"realmId"`
}

func (GetClientDescriptionConverterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientDescriptionConverterArgs)(nil)).Elem()
}

// A collection of values returned by getClientDescriptionConverter.
type GetClientDescriptionConverterResultOutput struct{ *pulumi.OutputState }

func (GetClientDescriptionConverterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientDescriptionConverterResult)(nil)).Elem()
}

func (o GetClientDescriptionConverterResultOutput) ToGetClientDescriptionConverterResultOutput() GetClientDescriptionConverterResultOutput {
	return o
}

func (o GetClientDescriptionConverterResultOutput) ToGetClientDescriptionConverterResultOutputWithContext(ctx context.Context) GetClientDescriptionConverterResultOutput {
	return o
}

func (o GetClientDescriptionConverterResultOutput) Access() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) map[string]string { return v.Access }).(pulumi.StringMapOutput)
}

func (o GetClientDescriptionConverterResultOutput) AdminUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.AdminUrl }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) map[string]string { return v.Attributes }).(pulumi.StringMapOutput)
}

func (o GetClientDescriptionConverterResultOutput) AuthenticationFlowBindingOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) map[string]string {
		return v.AuthenticationFlowBindingOverrides
	}).(pulumi.StringMapOutput)
}

func (o GetClientDescriptionConverterResultOutput) AuthorizationServicesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.AuthorizationServicesEnabled }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) AuthorizationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) map[string]string { return v.AuthorizationSettings }).(pulumi.StringMapOutput)
}

func (o GetClientDescriptionConverterResultOutput) BaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.BaseUrl }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) BearerOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.BearerOnly }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.Body }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) ClientAuthenticatorType() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.ClientAuthenticatorType }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) ConsentRequired() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.ConsentRequired }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) DefaultClientScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) []string { return v.DefaultClientScopes }).(pulumi.StringArrayOutput)
}

func (o GetClientDescriptionConverterResultOutput) DefaultRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) []string { return v.DefaultRoles }).(pulumi.StringArrayOutput)
}

func (o GetClientDescriptionConverterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) DirectAccessGrantsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.DirectAccessGrantsEnabled }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) FrontchannelLogout() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.FrontchannelLogout }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) FullScopeAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.FullScopeAllowed }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClientDescriptionConverterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) ImplicitFlowEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.ImplicitFlowEnabled }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) NotBefore() pulumi.IntOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) int { return v.NotBefore }).(pulumi.IntOutput)
}

func (o GetClientDescriptionConverterResultOutput) OptionalClientScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) []string { return v.OptionalClientScopes }).(pulumi.StringArrayOutput)
}

func (o GetClientDescriptionConverterResultOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.Origin }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) ProtocolMappers() GetClientDescriptionConverterProtocolMapperArrayOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) []GetClientDescriptionConverterProtocolMapper {
		return v.ProtocolMappers
	}).(GetClientDescriptionConverterProtocolMapperArrayOutput)
}

func (o GetClientDescriptionConverterResultOutput) PublicClient() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.PublicClient }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.RealmId }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) RedirectUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) []string { return v.RedirectUris }).(pulumi.StringArrayOutput)
}

func (o GetClientDescriptionConverterResultOutput) RegisteredNodes() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) map[string]string { return v.RegisteredNodes }).(pulumi.StringMapOutput)
}

func (o GetClientDescriptionConverterResultOutput) RegistrationAccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.RegistrationAccessToken }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) RootUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.RootUrl }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) string { return v.Secret }).(pulumi.StringOutput)
}

func (o GetClientDescriptionConverterResultOutput) ServiceAccountsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.ServiceAccountsEnabled }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) StandardFlowEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.StandardFlowEnabled }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) SurrogateAuthRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) bool { return v.SurrogateAuthRequired }).(pulumi.BoolOutput)
}

func (o GetClientDescriptionConverterResultOutput) WebOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetClientDescriptionConverterResult) []string { return v.WebOrigins }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClientDescriptionConverterResultOutput{})
}
