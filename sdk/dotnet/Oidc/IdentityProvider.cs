// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Oidc
{
    /// <summary>
    /// Allows for creating and managing OIDC Identity Providers within Keycloak.
    /// 
    /// OIDC (OpenID Connect) identity providers allows users to authenticate through a third party system using the OIDC standard.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var realm = new Keycloak.Realm("realm", new()
    ///     {
    ///         RealmName = "my-realm",
    ///         Enabled = true,
    ///     });
    /// 
    ///     var realmIdentityProvider = new Keycloak.Oidc.IdentityProvider("realm_identity_provider", new()
    ///     {
    ///         Realm = realm.Id,
    ///         Alias = "my-idp",
    ///         AuthorizationUrl = "https://authorizationurl.com",
    ///         ClientId = "clientID",
    ///         ClientSecret = "clientSecret",
    ///         TokenUrl = "https://tokenurl.com",
    ///         ExtraConfig = 
    ///         {
    ///             { "clientAuthMethod", "client_secret_post" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Identity providers can be imported using the format `{{realm_id}}/{{idp_alias}}`, where `idp_alias` is the identity provider alias.
    /// 
    /// Example:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:oidc/identityProvider:IdentityProvider realm_identity_provider my-realm/my-idp
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:oidc/identityProvider:IdentityProvider")]
    public partial class IdentityProvider : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
        /// </summary>
        [Output("acceptsPromptNoneForwardFromClient")]
        public Output<bool?> AcceptsPromptNoneForwardFromClient { get; private set; } = null!;

        /// <summary>
        /// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        /// </summary>
        [Output("addReadTokenRoleOnCreate")]
        public Output<bool?> AddReadTokenRoleOnCreate { get; private set; } = null!;

        /// <summary>
        /// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Enable/disable authenticate users by default.
        /// </summary>
        [Output("authenticateByDefault")]
        public Output<bool?> AuthenticateByDefault { get; private set; } = null!;

        /// <summary>
        /// The Authorization Url.
        /// </summary>
        [Output("authorizationUrl")]
        public Output<string> AuthorizationUrl { get; private set; } = null!;

        /// <summary>
        /// Does the external IDP support backchannel logout? Defaults to `true`.
        /// </summary>
        [Output("backchannelSupported")]
        public Output<bool?> BackchannelSupported { get; private set; } = null!;

        /// <summary>
        /// The client or client identifier registered within the identity provider.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
        /// </summary>
        [Output("defaultScopes")]
        public Output<string?> DefaultScopes { get; private set; } = null!;

        /// <summary>
        /// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
        /// </summary>
        [Output("disableUserInfo")]
        public Output<bool?> DisableUserInfo { get; private set; } = null!;

        /// <summary>
        /// Display name for the identity provider in the GUI.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("extraConfig")]
        public Output<ImmutableDictionary<string, object>?> ExtraConfig { get; private set; } = null!;

        /// <summary>
        /// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
        /// </summary>
        [Output("firstBrokerLoginFlowAlias")]
        public Output<string?> FirstBrokerLoginFlowAlias { get; private set; } = null!;

        /// <summary>
        /// A number defining the order of this identity provider in the GUI.
        /// </summary>
        [Output("guiOrder")]
        public Output<string?> GuiOrder { get; private set; } = null!;

        /// <summary>
        /// When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
        /// </summary>
        [Output("hideOnLoginPage")]
        public Output<bool?> HideOnLoginPage { get; private set; } = null!;

        /// <summary>
        /// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
        /// </summary>
        [Output("internalId")]
        public Output<string> InternalId { get; private set; } = null!;

        /// <summary>
        /// The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
        /// </summary>
        [Output("issuer")]
        public Output<string?> Issuer { get; private set; } = null!;

        /// <summary>
        /// JSON Web Key Set URL.
        /// </summary>
        [Output("jwksUrl")]
        public Output<string?> JwksUrl { get; private set; } = null!;

        /// <summary>
        /// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        /// </summary>
        [Output("linkOnly")]
        public Output<bool?> LinkOnly { get; private set; } = null!;

        /// <summary>
        /// Pass login hint to identity provider.
        /// </summary>
        [Output("loginHint")]
        public Output<string?> LoginHint { get; private set; } = null!;

        /// <summary>
        /// The Logout URL is the end session endpoint to use to logout user from external identity provider.
        /// </summary>
        [Output("logoutUrl")]
        public Output<string?> LogoutUrl { get; private set; } = null!;

        /// <summary>
        /// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
        /// </summary>
        [Output("postBrokerLoginFlowAlias")]
        public Output<string?> PostBrokerLoginFlowAlias { get; private set; } = null!;

        /// <summary>
        /// The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
        /// </summary>
        [Output("providerId")]
        public Output<string?> ProviderId { get; private set; } = null!;

        /// <summary>
        /// The name of the realm. This is unique across Keycloak.
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
        /// </summary>
        [Output("storeToken")]
        public Output<bool?> StoreToken { get; private set; } = null!;

        /// <summary>
        /// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
        /// </summary>
        [Output("syncMode")]
        public Output<string?> SyncMode { get; private set; } = null!;

        /// <summary>
        /// The Token URL.
        /// </summary>
        [Output("tokenUrl")]
        public Output<string> TokenUrl { get; private set; } = null!;

        /// <summary>
        /// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        /// </summary>
        [Output("trustEmail")]
        public Output<bool?> TrustEmail { get; private set; } = null!;

        /// <summary>
        /// Pass current locale to identity provider. Defaults to `false`.
        /// </summary>
        [Output("uiLocales")]
        public Output<bool?> UiLocales { get; private set; } = null!;

        /// <summary>
        /// User Info URL.
        /// </summary>
        [Output("userInfoUrl")]
        public Output<string?> UserInfoUrl { get; private set; } = null!;

        /// <summary>
        /// Enable/disable signature validation of external IDP signatures. Defaults to `false`.
        /// </summary>
        [Output("validateSignature")]
        public Output<bool?> ValidateSignature { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityProvider(string name, IdentityProviderArgs args, CustomResourceOptions? options = null)
            : base("keycloak:oidc/identityProvider:IdentityProvider", name, args ?? new IdentityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityProvider(string name, Input<string> id, IdentityProviderState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:oidc/identityProvider:IdentityProvider", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "clientSecret",
                },
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

    public sealed class IdentityProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
        /// </summary>
        [Input("acceptsPromptNoneForwardFromClient")]
        public Input<bool>? AcceptsPromptNoneForwardFromClient { get; set; }

        /// <summary>
        /// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        /// </summary>
        [Input("addReadTokenRoleOnCreate")]
        public Input<bool>? AddReadTokenRoleOnCreate { get; set; }

        /// <summary>
        /// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// Enable/disable authenticate users by default.
        /// </summary>
        [Input("authenticateByDefault")]
        public Input<bool>? AuthenticateByDefault { get; set; }

        /// <summary>
        /// The Authorization Url.
        /// </summary>
        [Input("authorizationUrl", required: true)]
        public Input<string> AuthorizationUrl { get; set; } = null!;

        /// <summary>
        /// Does the external IDP support backchannel logout? Defaults to `true`.
        /// </summary>
        [Input("backchannelSupported")]
        public Input<bool>? BackchannelSupported { get; set; }

        /// <summary>
        /// The client or client identifier registered within the identity provider.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSecret", required: true)]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
        /// </summary>
        [Input("defaultScopes")]
        public Input<string>? DefaultScopes { get; set; }

        /// <summary>
        /// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
        /// </summary>
        [Input("disableUserInfo")]
        public Input<bool>? DisableUserInfo { get; set; }

        /// <summary>
        /// Display name for the identity provider in the GUI.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
        /// </summary>
        [Input("firstBrokerLoginFlowAlias")]
        public Input<string>? FirstBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// A number defining the order of this identity provider in the GUI.
        /// </summary>
        [Input("guiOrder")]
        public Input<string>? GuiOrder { get; set; }

        /// <summary>
        /// When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
        /// </summary>
        [Input("hideOnLoginPage")]
        public Input<bool>? HideOnLoginPage { get; set; }

        /// <summary>
        /// The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// JSON Web Key Set URL.
        /// </summary>
        [Input("jwksUrl")]
        public Input<string>? JwksUrl { get; set; }

        /// <summary>
        /// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        /// </summary>
        [Input("linkOnly")]
        public Input<bool>? LinkOnly { get; set; }

        /// <summary>
        /// Pass login hint to identity provider.
        /// </summary>
        [Input("loginHint")]
        public Input<string>? LoginHint { get; set; }

        /// <summary>
        /// The Logout URL is the end session endpoint to use to logout user from external identity provider.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        /// <summary>
        /// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
        /// </summary>
        [Input("postBrokerLoginFlowAlias")]
        public Input<string>? PostBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The name of the realm. This is unique across Keycloak.
        /// </summary>
        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        /// <summary>
        /// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
        /// </summary>
        [Input("storeToken")]
        public Input<bool>? StoreToken { get; set; }

        /// <summary>
        /// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
        /// </summary>
        [Input("syncMode")]
        public Input<string>? SyncMode { get; set; }

        /// <summary>
        /// The Token URL.
        /// </summary>
        [Input("tokenUrl", required: true)]
        public Input<string> TokenUrl { get; set; } = null!;

        /// <summary>
        /// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        /// </summary>
        [Input("trustEmail")]
        public Input<bool>? TrustEmail { get; set; }

        /// <summary>
        /// Pass current locale to identity provider. Defaults to `false`.
        /// </summary>
        [Input("uiLocales")]
        public Input<bool>? UiLocales { get; set; }

        /// <summary>
        /// User Info URL.
        /// </summary>
        [Input("userInfoUrl")]
        public Input<string>? UserInfoUrl { get; set; }

        /// <summary>
        /// Enable/disable signature validation of external IDP signatures. Defaults to `false`.
        /// </summary>
        [Input("validateSignature")]
        public Input<bool>? ValidateSignature { get; set; }

        public IdentityProviderArgs()
        {
        }
        public static new IdentityProviderArgs Empty => new IdentityProviderArgs();
    }

    public sealed class IdentityProviderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
        /// </summary>
        [Input("acceptsPromptNoneForwardFromClient")]
        public Input<bool>? AcceptsPromptNoneForwardFromClient { get; set; }

        /// <summary>
        /// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        /// </summary>
        [Input("addReadTokenRoleOnCreate")]
        public Input<bool>? AddReadTokenRoleOnCreate { get; set; }

        /// <summary>
        /// The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Enable/disable authenticate users by default.
        /// </summary>
        [Input("authenticateByDefault")]
        public Input<bool>? AuthenticateByDefault { get; set; }

        /// <summary>
        /// The Authorization Url.
        /// </summary>
        [Input("authorizationUrl")]
        public Input<string>? AuthorizationUrl { get; set; }

        /// <summary>
        /// Does the external IDP support backchannel logout? Defaults to `true`.
        /// </summary>
        [Input("backchannelSupported")]
        public Input<bool>? BackchannelSupported { get; set; }

        /// <summary>
        /// The client or client identifier registered within the identity provider.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
        /// </summary>
        [Input("defaultScopes")]
        public Input<string>? DefaultScopes { get; set; }

        /// <summary>
        /// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
        /// </summary>
        [Input("disableUserInfo")]
        public Input<bool>? DisableUserInfo { get; set; }

        /// <summary>
        /// Display name for the identity provider in the GUI.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
        /// </summary>
        [Input("firstBrokerLoginFlowAlias")]
        public Input<string>? FirstBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// A number defining the order of this identity provider in the GUI.
        /// </summary>
        [Input("guiOrder")]
        public Input<string>? GuiOrder { get; set; }

        /// <summary>
        /// When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
        /// </summary>
        [Input("hideOnLoginPage")]
        public Input<bool>? HideOnLoginPage { get; set; }

        /// <summary>
        /// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
        /// </summary>
        [Input("internalId")]
        public Input<string>? InternalId { get; set; }

        /// <summary>
        /// The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// JSON Web Key Set URL.
        /// </summary>
        [Input("jwksUrl")]
        public Input<string>? JwksUrl { get; set; }

        /// <summary>
        /// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        /// </summary>
        [Input("linkOnly")]
        public Input<bool>? LinkOnly { get; set; }

        /// <summary>
        /// Pass login hint to identity provider.
        /// </summary>
        [Input("loginHint")]
        public Input<string>? LoginHint { get; set; }

        /// <summary>
        /// The Logout URL is the end session endpoint to use to logout user from external identity provider.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        /// <summary>
        /// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
        /// </summary>
        [Input("postBrokerLoginFlowAlias")]
        public Input<string>? PostBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The name of the realm. This is unique across Keycloak.
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// When `true`, tokens will be stored after authenticating users. Defaults to `true`.
        /// </summary>
        [Input("storeToken")]
        public Input<bool>? StoreToken { get; set; }

        /// <summary>
        /// The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
        /// </summary>
        [Input("syncMode")]
        public Input<string>? SyncMode { get; set; }

        /// <summary>
        /// The Token URL.
        /// </summary>
        [Input("tokenUrl")]
        public Input<string>? TokenUrl { get; set; }

        /// <summary>
        /// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        /// </summary>
        [Input("trustEmail")]
        public Input<bool>? TrustEmail { get; set; }

        /// <summary>
        /// Pass current locale to identity provider. Defaults to `false`.
        /// </summary>
        [Input("uiLocales")]
        public Input<bool>? UiLocales { get; set; }

        /// <summary>
        /// User Info URL.
        /// </summary>
        [Input("userInfoUrl")]
        public Input<string>? UserInfoUrl { get; set; }

        /// <summary>
        /// Enable/disable signature validation of external IDP signatures. Defaults to `false`.
        /// </summary>
        [Input("validateSignature")]
        public Input<bool>? ValidateSignature { get; set; }

        public IdentityProviderState()
        {
        }
        public static new IdentityProviderState Empty => new IdentityProviderState();
    }
}
