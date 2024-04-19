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
    ///     var google = new Keycloak.Oidc.GoogleIdentityProvider("google", new()
    ///     {
    ///         Realm = realm.Id,
    ///         ClientId = googleIdentityProviderClientId,
    ///         ClientSecret = googleIdentityProviderClientSecret,
    ///         TrustEmail = true,
    ///         HostedDomain = "example.com",
    ///         SyncMode = "IMPORT",
    ///         ExtraConfig = 
    ///         {
    ///             { "myCustomConfigKey", "myValue" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Google Identity providers can be imported using the format {{realm_id}}/{{idp_alias}}, where idp_alias is the identity provider alias.
    /// 
    /// Example:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider google_identity_provider my-realm/my-google-idp
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider")]
    public partial class GoogleIdentityProvider : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
        /// </summary>
        [Output("acceptsPromptNoneForwardFromClient")]
        public Output<bool?> AcceptsPromptNoneForwardFromClient { get; private set; } = null!;

        /// <summary>
        /// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        /// </summary>
        [Output("addReadTokenRoleOnCreate")]
        public Output<bool?> AddReadTokenRoleOnCreate { get; private set; } = null!;

        /// <summary>
        /// (Computed) The alias for the Google identity provider.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Enable/disable authenticate users by default.
        /// </summary>
        [Output("authenticateByDefault")]
        public Output<bool?> AuthenticateByDefault { get; private set; } = null!;

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
        /// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
        /// </summary>
        [Output("defaultScopes")]
        public Output<string?> DefaultScopes { get; private set; } = null!;

        /// <summary>
        /// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
        /// </summary>
        [Output("disableUserInfo")]
        public Output<bool?> DisableUserInfo { get; private set; } = null!;

        /// <summary>
        /// (Computed) Display name for the Google identity provider in the GUI.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

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
        /// When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
        /// </summary>
        [Output("hideOnLoginPage")]
        public Output<bool?> HideOnLoginPage { get; private set; } = null!;

        /// <summary>
        /// Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
        /// </summary>
        [Output("hostedDomain")]
        public Output<string?> HostedDomain { get; private set; } = null!;

        /// <summary>
        /// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
        /// </summary>
        [Output("internalId")]
        public Output<string> InternalId { get; private set; } = null!;

        /// <summary>
        /// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        /// </summary>
        [Output("linkOnly")]
        public Output<bool?> LinkOnly { get; private set; } = null!;

        /// <summary>
        /// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
        /// </summary>
        [Output("postBrokerLoginFlowAlias")]
        public Output<string?> PostBrokerLoginFlowAlias { get; private set; } = null!;

        /// <summary>
        /// The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
        /// </summary>
        [Output("providerId")]
        public Output<string?> ProviderId { get; private set; } = null!;

        /// <summary>
        /// The name of the realm. This is unique across Keycloak.
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// Sets the "access_type" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
        /// </summary>
        [Output("requestRefreshToken")]
        public Output<bool?> RequestRefreshToken { get; private set; } = null!;

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
        /// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        /// </summary>
        [Output("trustEmail")]
        public Output<bool?> TrustEmail { get; private set; } = null!;

        /// <summary>
        /// Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
        /// </summary>
        [Output("useUserIpParam")]
        public Output<bool?> UseUserIpParam { get; private set; } = null!;


        /// <summary>
        /// Create a GoogleIdentityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GoogleIdentityProvider(string name, GoogleIdentityProviderArgs args, CustomResourceOptions? options = null)
            : base("keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider", name, args ?? new GoogleIdentityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GoogleIdentityProvider(string name, Input<string> id, GoogleIdentityProviderState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GoogleIdentityProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GoogleIdentityProvider Get(string name, Input<string> id, GoogleIdentityProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new GoogleIdentityProvider(name, id, state, options);
        }
    }

    public sealed class GoogleIdentityProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
        /// </summary>
        [Input("acceptsPromptNoneForwardFromClient")]
        public Input<bool>? AcceptsPromptNoneForwardFromClient { get; set; }

        /// <summary>
        /// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        /// </summary>
        [Input("addReadTokenRoleOnCreate")]
        public Input<bool>? AddReadTokenRoleOnCreate { get; set; }

        /// <summary>
        /// Enable/disable authenticate users by default.
        /// </summary>
        [Input("authenticateByDefault")]
        public Input<bool>? AuthenticateByDefault { get; set; }

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
        /// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
        /// </summary>
        [Input("defaultScopes")]
        public Input<string>? DefaultScopes { get; set; }

        /// <summary>
        /// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
        /// </summary>
        [Input("disableUserInfo")]
        public Input<bool>? DisableUserInfo { get; set; }

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
        /// When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
        /// </summary>
        [Input("hideOnLoginPage")]
        public Input<bool>? HideOnLoginPage { get; set; }

        /// <summary>
        /// Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
        /// </summary>
        [Input("hostedDomain")]
        public Input<string>? HostedDomain { get; set; }

        /// <summary>
        /// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        /// </summary>
        [Input("linkOnly")]
        public Input<bool>? LinkOnly { get; set; }

        /// <summary>
        /// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
        /// </summary>
        [Input("postBrokerLoginFlowAlias")]
        public Input<string>? PostBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The name of the realm. This is unique across Keycloak.
        /// </summary>
        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        /// <summary>
        /// Sets the "access_type" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
        /// </summary>
        [Input("requestRefreshToken")]
        public Input<bool>? RequestRefreshToken { get; set; }

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
        /// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        /// </summary>
        [Input("trustEmail")]
        public Input<bool>? TrustEmail { get; set; }

        /// <summary>
        /// Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
        /// </summary>
        [Input("useUserIpParam")]
        public Input<bool>? UseUserIpParam { get; set; }

        public GoogleIdentityProviderArgs()
        {
        }
        public static new GoogleIdentityProviderArgs Empty => new GoogleIdentityProviderArgs();
    }

    public sealed class GoogleIdentityProviderState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
        /// </summary>
        [Input("acceptsPromptNoneForwardFromClient")]
        public Input<bool>? AcceptsPromptNoneForwardFromClient { get; set; }

        /// <summary>
        /// When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
        /// </summary>
        [Input("addReadTokenRoleOnCreate")]
        public Input<bool>? AddReadTokenRoleOnCreate { get; set; }

        /// <summary>
        /// (Computed) The alias for the Google identity provider.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Enable/disable authenticate users by default.
        /// </summary>
        [Input("authenticateByDefault")]
        public Input<bool>? AuthenticateByDefault { get; set; }

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
        /// The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
        /// </summary>
        [Input("defaultScopes")]
        public Input<string>? DefaultScopes { get; set; }

        /// <summary>
        /// When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
        /// </summary>
        [Input("disableUserInfo")]
        public Input<bool>? DisableUserInfo { get; set; }

        /// <summary>
        /// (Computed) Display name for the Google identity provider in the GUI.
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
        /// When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
        /// </summary>
        [Input("hideOnLoginPage")]
        public Input<bool>? HideOnLoginPage { get; set; }

        /// <summary>
        /// Sets the "hd" query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
        /// </summary>
        [Input("hostedDomain")]
        public Input<string>? HostedDomain { get; set; }

        /// <summary>
        /// (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
        /// </summary>
        [Input("internalId")]
        public Input<string>? InternalId { get; set; }

        /// <summary>
        /// When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
        /// </summary>
        [Input("linkOnly")]
        public Input<bool>? LinkOnly { get; set; }

        /// <summary>
        /// The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
        /// </summary>
        [Input("postBrokerLoginFlowAlias")]
        public Input<string>? PostBrokerLoginFlowAlias { get; set; }

        /// <summary>
        /// The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The name of the realm. This is unique across Keycloak.
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// Sets the "access_type" query parameter to "offline" when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
        /// </summary>
        [Input("requestRefreshToken")]
        public Input<bool>? RequestRefreshToken { get; set; }

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
        /// When `true`, email addresses for users in this provider will automatically be verified regardless of the realm's email verification policy. Defaults to `false`.
        /// </summary>
        [Input("trustEmail")]
        public Input<bool>? TrustEmail { get; set; }

        /// <summary>
        /// Sets the "userIp" query parameter when querying Google's User Info service. This will use the user's IP address. This is useful if Google is throttling Keycloak's access to the User Info service.
        /// </summary>
        [Input("useUserIpParam")]
        public Input<bool>? UseUserIpParam { get; set; }

        public GoogleIdentityProviderState()
        {
        }
        public static new GoogleIdentityProviderState Empty => new GoogleIdentityProviderState();
    }
}
