// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Saml
{
    /// <summary>
    /// Allows for creating and managing SAML Identity Providers within Keycloak.
    /// 
    /// SAML (Security Assertion Markup Language) identity providers allows users to authenticate through a third-party system using the SAML protocol.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var realm = new Keycloak.Realm("realm", new Keycloak.RealmArgs
    ///         {
    ///             Realm = "my-realm",
    ///             Enabled = true,
    ///         });
    ///         var realmSamlIdentityProvider = new Keycloak.Saml.IdentityProvider("realmSamlIdentityProvider", new Keycloak.Saml.IdentityProviderArgs
    ///         {
    ///             Realm = realm.Id,
    ///             Alias = "my-saml-idp",
    ///             EntityId = "https://domain.com/entity_id",
    ///             SingleSignOnServiceUrl = "https://domain.com/adfs/ls/",
    ///             SingleLogoutServiceUrl = "https://domain.com/adfs/ls/?wa=wsignout1.0",
    ///             BackchannelSupported = true,
    ///             PostBindingResponse = true,
    ///             PostBindingLogout = true,
    ///             PostBindingAuthnRequest = true,
    ///             StoreToken = false,
    ///             TrustEmail = true,
    ///             ForceAuthn = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Identity providers can be imported using the format `{{realm_id}}/{{idp_alias}}`, where `idp_alias` is the identity provider alias. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:saml/identityProvider:IdentityProvider realm_saml_identity_provider my-realm/my-saml-idp
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:saml/identityProvider:IdentityProvider")]
    public partial class IdentityProvider : Pulumi.CustomResource
    {
        /// <summary>
        /// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        /// </summary>
        [Output("addReadTokenRoleOnCreate")]
        public Output<bool?> AddReadTokenRoleOnCreate { get; private set; } = null!;

        /// <summary>
        /// The unique name of identity provider.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Authenticate users by default. Defaults to `false`.
        /// </summary>
        [Output("authenticateByDefault")]
        public Output<bool?> AuthenticateByDefault { get; private set; } = null!;

        /// <summary>
        /// Does the external IDP support back-channel logout ?.
        /// </summary>
        [Output("backchannelSupported")]
        public Output<bool?> BackchannelSupported { get; private set; } = null!;

        /// <summary>
        /// The display name for the realm that is shown when logging in to the admin console.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// When `false`, users and clients will not be able to access this realm. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The Entity ID that will be used to uniquely identify this SAML Service Provider.
        /// </summary>
        [Output("entityId")]
        public Output<string> EntityId { get; private set; } = null!;

        /// <summary>
        /// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
        /// </summary>
        [Output("firstBrokerLoginFlowAlias")]
        public Output<string?> FirstBrokerLoginFlowAlias { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
        /// </summary>
        [Output("forceAuthn")]
        public Output<bool?> ForceAuthn { get; private set; } = null!;

        /// <summary>
        /// GUI Order
        /// </summary>
        [Output("guiOrder")]
        public Output<string?> GuiOrder { get; private set; } = null!;

        /// <summary>
        /// If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
        /// </summary>
        [Output("hideOnLoginPage")]
        public Output<bool?> HideOnLoginPage { get; private set; } = null!;

        /// <summary>
        /// Internal Identity Provider Id
        /// </summary>
        [Output("internalId")]
        public Output<string> InternalId { get; private set; } = null!;

        /// <summary>
        /// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        /// </summary>
        [Output("linkOnly")]
        public Output<bool?> LinkOnly { get; private set; } = null!;

        /// <summary>
        /// Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
        /// </summary>
        [Output("nameIdPolicyFormat")]
        public Output<string?> NameIdPolicyFormat { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
        /// </summary>
        [Output("postBindingAuthnRequest")]
        public Output<bool?> PostBindingAuthnRequest { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
        /// </summary>
        [Output("postBindingLogout")]
        public Output<bool?> PostBindingLogout { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
        /// </summary>
        [Output("postBindingResponse")]
        public Output<bool?> PostBindingResponse { get; private set; } = null!;

        /// <summary>
        /// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
        /// </summary>
        [Output("postBrokerLoginFlowAlias")]
        public Output<string?> PostBrokerLoginFlowAlias { get; private set; } = null!;

        /// <summary>
        /// Principal Attribute
        /// </summary>
        [Output("principalAttribute")]
        public Output<string?> PrincipalAttribute { get; private set; } = null!;

        /// <summary>
        /// Principal Type
        /// </summary>
        [Output("principalType")]
        public Output<string?> PrincipalType { get; private set; } = null!;

        /// <summary>
        /// The name of the realm. This is unique across Keycloak.
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// Signing Algorithm. Defaults to empty.
        /// </summary>
        [Output("signatureAlgorithm")]
        public Output<string?> SignatureAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Signing Certificate.
        /// </summary>
        [Output("signingCertificate")]
        public Output<string?> SigningCertificate { get; private set; } = null!;

        /// <summary>
        /// The Url that must be used to send logout requests.
        /// </summary>
        [Output("singleLogoutServiceUrl")]
        public Output<string?> SingleLogoutServiceUrl { get; private set; } = null!;

        /// <summary>
        /// The Url that must be used to send authentication requests (SAML AuthnRequest).
        /// </summary>
        [Output("singleSignOnServiceUrl")]
        public Output<string> SingleSignOnServiceUrl { get; private set; } = null!;

        /// <summary>
        /// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
        /// </summary>
        [Output("storeToken")]
        public Output<bool?> StoreToken { get; private set; } = null!;

        /// <summary>
        /// Sync Mode
        /// </summary>
        [Output("syncMode")]
        public Output<string?> SyncMode { get; private set; } = null!;

        /// <summary>
        /// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        /// </summary>
        [Output("trustEmail")]
        public Output<bool?> TrustEmail { get; private set; } = null!;

        /// <summary>
        /// Enable/disable signature validation of SAML responses.
        /// </summary>
        [Output("validateSignature")]
        public Output<bool?> ValidateSignature { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this service provider expects an encrypted Assertion.
        /// </summary>
        [Output("wantAssertionsEncrypted")]
        public Output<bool?> WantAssertionsEncrypted { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this service provider expects a signed Assertion.
        /// </summary>
        [Output("wantAssertionsSigned")]
        public Output<bool?> WantAssertionsSigned { get; private set; } = null!;

        /// <summary>
        /// Sign Key Transformer. Defaults to empty.
        /// </summary>
        [Output("xmlSignKeyInfoKeyNameTransformer")]
        public Output<string?> XmlSignKeyInfoKeyNameTransformer { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityProvider(string name, IdentityProviderArgs args, CustomResourceOptions? options = null)
            : base("keycloak:saml/identityProvider:IdentityProvider", name, args ?? new IdentityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityProvider(string name, Input<string> id, IdentityProviderState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:saml/identityProvider:IdentityProvider", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IdentityProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityProvider Get(string name, Input<string> id, IdentityProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityProvider(name, id, state, options);
        }
    }

    public sealed class IdentityProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        /// </summary>
        [Input("addReadTokenRoleOnCreate")]
        public Input<bool>? AddReadTokenRoleOnCreate { get; set; }

        /// <summary>
        /// The unique name of identity provider.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// Authenticate users by default. Defaults to `false`.
        /// </summary>
        [Input("authenticateByDefault")]
        public Input<bool>? AuthenticateByDefault { get; set; }

        /// <summary>
        /// Does the external IDP support back-channel logout ?.
        /// </summary>
        [Input("backchannelSupported")]
        public Input<bool>? BackchannelSupported { get; set; }

        /// <summary>
        /// The display name for the realm that is shown when logging in to the admin console.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// When `false`, users and clients will not be able to access this realm. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Entity ID that will be used to uniquely identify this SAML Service Provider.
        /// </summary>
        [Input("entityId", required: true)]
        public Input<string> EntityId { get; set; } = null!;

        /// <summary>
        /// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
        /// </summary>
        [Input("firstBrokerLoginFlowAlias")]
        public Input<string>? FirstBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
        /// </summary>
        [Input("forceAuthn")]
        public Input<bool>? ForceAuthn { get; set; }

        /// <summary>
        /// GUI Order
        /// </summary>
        [Input("guiOrder")]
        public Input<string>? GuiOrder { get; set; }

        /// <summary>
        /// If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
        /// </summary>
        [Input("hideOnLoginPage")]
        public Input<bool>? HideOnLoginPage { get; set; }

        /// <summary>
        /// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        /// </summary>
        [Input("linkOnly")]
        public Input<bool>? LinkOnly { get; set; }

        /// <summary>
        /// Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
        /// </summary>
        [Input("nameIdPolicyFormat")]
        public Input<string>? NameIdPolicyFormat { get; set; }

        /// <summary>
        /// Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
        /// </summary>
        [Input("postBindingAuthnRequest")]
        public Input<bool>? PostBindingAuthnRequest { get; set; }

        /// <summary>
        /// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
        /// </summary>
        [Input("postBindingLogout")]
        public Input<bool>? PostBindingLogout { get; set; }

        /// <summary>
        /// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
        /// </summary>
        [Input("postBindingResponse")]
        public Input<bool>? PostBindingResponse { get; set; }

        /// <summary>
        /// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
        /// </summary>
        [Input("postBrokerLoginFlowAlias")]
        public Input<string>? PostBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// Principal Attribute
        /// </summary>
        [Input("principalAttribute")]
        public Input<string>? PrincipalAttribute { get; set; }

        /// <summary>
        /// Principal Type
        /// </summary>
        [Input("principalType")]
        public Input<string>? PrincipalType { get; set; }

        /// <summary>
        /// The name of the realm. This is unique across Keycloak.
        /// </summary>
        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        /// <summary>
        /// Signing Algorithm. Defaults to empty.
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        /// <summary>
        /// Signing Certificate.
        /// </summary>
        [Input("signingCertificate")]
        public Input<string>? SigningCertificate { get; set; }

        /// <summary>
        /// The Url that must be used to send logout requests.
        /// </summary>
        [Input("singleLogoutServiceUrl")]
        public Input<string>? SingleLogoutServiceUrl { get; set; }

        /// <summary>
        /// The Url that must be used to send authentication requests (SAML AuthnRequest).
        /// </summary>
        [Input("singleSignOnServiceUrl", required: true)]
        public Input<string> SingleSignOnServiceUrl { get; set; } = null!;

        /// <summary>
        /// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
        /// </summary>
        [Input("storeToken")]
        public Input<bool>? StoreToken { get; set; }

        /// <summary>
        /// Sync Mode
        /// </summary>
        [Input("syncMode")]
        public Input<string>? SyncMode { get; set; }

        /// <summary>
        /// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        /// </summary>
        [Input("trustEmail")]
        public Input<bool>? TrustEmail { get; set; }

        /// <summary>
        /// Enable/disable signature validation of SAML responses.
        /// </summary>
        [Input("validateSignature")]
        public Input<bool>? ValidateSignature { get; set; }

        /// <summary>
        /// Indicates whether this service provider expects an encrypted Assertion.
        /// </summary>
        [Input("wantAssertionsEncrypted")]
        public Input<bool>? WantAssertionsEncrypted { get; set; }

        /// <summary>
        /// Indicates whether this service provider expects a signed Assertion.
        /// </summary>
        [Input("wantAssertionsSigned")]
        public Input<bool>? WantAssertionsSigned { get; set; }

        /// <summary>
        /// Sign Key Transformer. Defaults to empty.
        /// </summary>
        [Input("xmlSignKeyInfoKeyNameTransformer")]
        public Input<string>? XmlSignKeyInfoKeyNameTransformer { get; set; }

        public IdentityProviderArgs()
        {
        }
    }

    public sealed class IdentityProviderState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        /// </summary>
        [Input("addReadTokenRoleOnCreate")]
        public Input<bool>? AddReadTokenRoleOnCreate { get; set; }

        /// <summary>
        /// The unique name of identity provider.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Authenticate users by default. Defaults to `false`.
        /// </summary>
        [Input("authenticateByDefault")]
        public Input<bool>? AuthenticateByDefault { get; set; }

        /// <summary>
        /// Does the external IDP support back-channel logout ?.
        /// </summary>
        [Input("backchannelSupported")]
        public Input<bool>? BackchannelSupported { get; set; }

        /// <summary>
        /// The display name for the realm that is shown when logging in to the admin console.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// When `false`, users and clients will not be able to access this realm. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Entity ID that will be used to uniquely identify this SAML Service Provider.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means that there is not yet existing Keycloak account linked with the authenticated identity provider account. Defaults to `first broker login`.
        /// </summary>
        [Input("firstBrokerLoginFlowAlias")]
        public Input<string>? FirstBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// Indicates whether the identity provider must authenticate the presenter directly rather than rely on a previous security context.
        /// </summary>
        [Input("forceAuthn")]
        public Input<bool>? ForceAuthn { get; set; }

        /// <summary>
        /// GUI Order
        /// </summary>
        [Input("guiOrder")]
        public Input<string>? GuiOrder { get; set; }

        /// <summary>
        /// If hidden, then login with this provider is possible only if requested explicitly, e.g. using the 'kc_idp_hint' parameter.
        /// </summary>
        [Input("hideOnLoginPage")]
        public Input<bool>? HideOnLoginPage { get; set; }

        /// <summary>
        /// Internal Identity Provider Id
        /// </summary>
        [Input("internalId")]
        public Input<string>? InternalId { get; set; }

        /// <summary>
        /// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        /// </summary>
        [Input("linkOnly")]
        public Input<bool>? LinkOnly { get; set; }

        /// <summary>
        /// Specifies the URI reference corresponding to a name identifier format. Defaults to empty.
        /// </summary>
        [Input("nameIdPolicyFormat")]
        public Input<string>? NameIdPolicyFormat { get; set; }

        /// <summary>
        /// Indicates whether the AuthnRequest must be sent using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
        /// </summary>
        [Input("postBindingAuthnRequest")]
        public Input<bool>? PostBindingAuthnRequest { get; set; }

        /// <summary>
        /// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used.
        /// </summary>
        [Input("postBindingLogout")]
        public Input<bool>? PostBindingLogout { get; set; }

        /// <summary>
        /// Indicates whether to respond to requests using HTTP-POST binding. If false, HTTP-REDIRECT binding will be used..
        /// </summary>
        [Input("postBindingResponse")]
        public Input<bool>? PostBindingResponse { get; set; }

        /// <summary>
        /// Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that authenticator implementations must assume that user is already set in ClientSession as identity provider already set it. Defaults to empty.
        /// </summary>
        [Input("postBrokerLoginFlowAlias")]
        public Input<string>? PostBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// Principal Attribute
        /// </summary>
        [Input("principalAttribute")]
        public Input<string>? PrincipalAttribute { get; set; }

        /// <summary>
        /// Principal Type
        /// </summary>
        [Input("principalType")]
        public Input<string>? PrincipalType { get; set; }

        /// <summary>
        /// The name of the realm. This is unique across Keycloak.
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// Signing Algorithm. Defaults to empty.
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        /// <summary>
        /// Signing Certificate.
        /// </summary>
        [Input("signingCertificate")]
        public Input<string>? SigningCertificate { get; set; }

        /// <summary>
        /// The Url that must be used to send logout requests.
        /// </summary>
        [Input("singleLogoutServiceUrl")]
        public Input<string>? SingleLogoutServiceUrl { get; set; }

        /// <summary>
        /// The Url that must be used to send authentication requests (SAML AuthnRequest).
        /// </summary>
        [Input("singleSignOnServiceUrl")]
        public Input<string>? SingleSignOnServiceUrl { get; set; }

        /// <summary>
        /// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
        /// </summary>
        [Input("storeToken")]
        public Input<bool>? StoreToken { get; set; }

        /// <summary>
        /// Sync Mode
        /// </summary>
        [Input("syncMode")]
        public Input<string>? SyncMode { get; set; }

        /// <summary>
        /// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        /// </summary>
        [Input("trustEmail")]
        public Input<bool>? TrustEmail { get; set; }

        /// <summary>
        /// Enable/disable signature validation of SAML responses.
        /// </summary>
        [Input("validateSignature")]
        public Input<bool>? ValidateSignature { get; set; }

        /// <summary>
        /// Indicates whether this service provider expects an encrypted Assertion.
        /// </summary>
        [Input("wantAssertionsEncrypted")]
        public Input<bool>? WantAssertionsEncrypted { get; set; }

        /// <summary>
        /// Indicates whether this service provider expects a signed Assertion.
        /// </summary>
        [Input("wantAssertionsSigned")]
        public Input<bool>? WantAssertionsSigned { get; set; }

        /// <summary>
        /// Sign Key Transformer. Defaults to empty.
        /// </summary>
        [Input("xmlSignKeyInfoKeyNameTransformer")]
        public Input<string>? XmlSignKeyInfoKeyNameTransformer { get; set; }

        public IdentityProviderState()
        {
        }
    }
}
