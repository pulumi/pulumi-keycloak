// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    /// <summary>
    /// ## # keycloak.openid.Client
    /// 
    /// Allows for creating and managing Keycloak clients that use the OpenID Connect protocol.
    /// 
    /// Clients are entities that can use Keycloak for user authentication. Typically,
    /// clients are applications that redirect users to Keycloak for authentication
    /// in order to take advantage of Keycloak's user sessions for SSO.
    /// 
    /// ### Example Usage
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
    ///     var openidClient = new Keycloak.OpenId.Client("openid_client", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "test-client",
    ///         Name = "test client",
    ///         Enabled = true,
    ///         AccessType = "CONFIDENTIAL",
    ///         ValidRedirectUris = new[]
    ///         {
    ///             "http://localhost:8080/openid-callback",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this client is attached to.
    /// - `client_id` - (Required) The unique ID of this client, referenced in the URI during authentication and in issued tokens.
    /// - `name` - (Optional) The display name of this client in the GUI.
    /// - `enabled` - (Optional) When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
    /// - `description` - (Optional) The description of this client in the GUI.
    /// - `access_type` - (Required) Specifies the type of client, which can be one of the following:
    ///     - `CONFIDENTIAL` - Used for server-side clients that require both client ID and secret when authenticating.
    ///       This client should be used for applications using the Authorization Code or Client Credentials grant flows.
    ///     - `PUBLIC` - Used for browser-only applications that do not require a client secret, and instead rely only on authorized redirect
    ///       URIs for security. This client should be used for applications using the Implicit grant flow.
    ///     - `BEARER-ONLY` - Used for services that never initiate a login. This client will only allow bearer token requests.
    /// - `client_secret` - (Optional) The secret for clients with an `access_type` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and
    ///   should be treated with the same care as a password. If omitted, Keycloak will generate a GUID for this attribute.
    /// - `standard_flow_enabled` - (Optional) When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
    /// - `implicit_flow_enabled` - (Optional) When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
    /// - `direct_access_grants_enabled` - (Optional) When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
    /// - `service_accounts_enabled` - (Optional) When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
    /// - `valid_redirect_uris` - (Optional) A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
    ///   wildcards in the form of an asterisk can be used here. This attribute must be set if either `standard_flow_enabled` or `implicit_flow_enabled`
    ///   is set to `true`.
    /// - `web_origins` - (Optional) A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
    /// - `admin_url` - (Optional) URL to the admin interface of the client.
    /// - `base_url` - (Optional) Default URL to use when the auth server needs to redirect or link back to the client.
    /// - `pkce_code_challenge_method` - (Optional) The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
    /// - `full_scope_allowed` - (Optional) - Allow to include all roles mappings in the access token.
    /// 
    /// ### Attributes Reference
    /// 
    /// In addition to the arguments listed above, the following computed attributes are exported:
    /// 
    /// - `service_account_user_id` - When service accounts are enabled for this client, this attribute is the unique ID for the Keycloak user that represents this service account.
    /// 
    /// ### Import
    /// 
    /// Clients can be imported using the format `{{realm_id}}/{{client_keycloak_id}}`, where `client_keycloak_id` is the unique ID that Keycloak
    /// assigns to the client upon creation. This value can be found in the URI when editing this client in the GUI, and is typically a GUID.
    /// 
    /// Example:
    /// </summary>
    [KeycloakResourceType("keycloak:openid/client:Client")]
    public partial class Client : global::Pulumi.CustomResource
    {
        [Output("accessTokenLifespan")]
        public Output<string> AccessTokenLifespan { get; private set; } = null!;

        [Output("accessType")]
        public Output<string> AccessType { get; private set; } = null!;

        [Output("adminUrl")]
        public Output<string> AdminUrl { get; private set; } = null!;

        [Output("authenticationFlowBindingOverrides")]
        public Output<Outputs.ClientAuthenticationFlowBindingOverrides?> AuthenticationFlowBindingOverrides { get; private set; } = null!;

        [Output("authorization")]
        public Output<Outputs.ClientAuthorization?> Authorization { get; private set; } = null!;

        [Output("backchannelLogoutRevokeOfflineSessions")]
        public Output<bool?> BackchannelLogoutRevokeOfflineSessions { get; private set; } = null!;

        [Output("backchannelLogoutSessionRequired")]
        public Output<bool?> BackchannelLogoutSessionRequired { get; private set; } = null!;

        [Output("backchannelLogoutUrl")]
        public Output<string?> BackchannelLogoutUrl { get; private set; } = null!;

        [Output("baseUrl")]
        public Output<string> BaseUrl { get; private set; } = null!;

        [Output("clientAuthenticatorType")]
        public Output<string?> ClientAuthenticatorType { get; private set; } = null!;

        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        [Output("clientOfflineSessionIdleTimeout")]
        public Output<string> ClientOfflineSessionIdleTimeout { get; private set; } = null!;

        [Output("clientOfflineSessionMaxLifespan")]
        public Output<string> ClientOfflineSessionMaxLifespan { get; private set; } = null!;

        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        [Output("clientSessionIdleTimeout")]
        public Output<string> ClientSessionIdleTimeout { get; private set; } = null!;

        [Output("clientSessionMaxLifespan")]
        public Output<string> ClientSessionMaxLifespan { get; private set; } = null!;

        [Output("consentRequired")]
        public Output<bool> ConsentRequired { get; private set; } = null!;

        [Output("consentScreenText")]
        public Output<string> ConsentScreenText { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("directAccessGrantsEnabled")]
        public Output<bool> DirectAccessGrantsEnabled { get; private set; } = null!;

        [Output("displayOnConsentScreen")]
        public Output<bool> DisplayOnConsentScreen { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("excludeSessionStateFromAuthResponse")]
        public Output<bool> ExcludeSessionStateFromAuthResponse { get; private set; } = null!;

        [Output("extraConfig")]
        public Output<ImmutableDictionary<string, object>?> ExtraConfig { get; private set; } = null!;

        [Output("frontchannelLogoutEnabled")]
        public Output<bool> FrontchannelLogoutEnabled { get; private set; } = null!;

        [Output("frontchannelLogoutUrl")]
        public Output<string?> FrontchannelLogoutUrl { get; private set; } = null!;

        [Output("fullScopeAllowed")]
        public Output<bool?> FullScopeAllowed { get; private set; } = null!;

        [Output("implicitFlowEnabled")]
        public Output<bool> ImplicitFlowEnabled { get; private set; } = null!;

        [Output("import")]
        public Output<bool?> Import { get; private set; } = null!;

        [Output("loginTheme")]
        public Output<string?> LoginTheme { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("oauth2DeviceAuthorizationGrantEnabled")]
        public Output<bool?> Oauth2DeviceAuthorizationGrantEnabled { get; private set; } = null!;

        [Output("oauth2DeviceCodeLifespan")]
        public Output<string?> Oauth2DeviceCodeLifespan { get; private set; } = null!;

        [Output("oauth2DevicePollingInterval")]
        public Output<string?> Oauth2DevicePollingInterval { get; private set; } = null!;

        [Output("pkceCodeChallengeMethod")]
        public Output<string?> PkceCodeChallengeMethod { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("resourceServerId")]
        public Output<string> ResourceServerId { get; private set; } = null!;

        [Output("rootUrl")]
        public Output<string> RootUrl { get; private set; } = null!;

        [Output("serviceAccountUserId")]
        public Output<string> ServiceAccountUserId { get; private set; } = null!;

        [Output("serviceAccountsEnabled")]
        public Output<bool> ServiceAccountsEnabled { get; private set; } = null!;

        [Output("standardFlowEnabled")]
        public Output<bool> StandardFlowEnabled { get; private set; } = null!;

        [Output("useRefreshTokens")]
        public Output<bool?> UseRefreshTokens { get; private set; } = null!;

        [Output("useRefreshTokensClientCredentials")]
        public Output<bool?> UseRefreshTokensClientCredentials { get; private set; } = null!;

        [Output("validPostLogoutRedirectUris")]
        public Output<ImmutableArray<string>> ValidPostLogoutRedirectUris { get; private set; } = null!;

        [Output("validRedirectUris")]
        public Output<ImmutableArray<string>> ValidRedirectUris { get; private set; } = null!;

        [Output("webOrigins")]
        public Output<ImmutableArray<string>> WebOrigins { get; private set; } = null!;


        /// <summary>
        /// Create a Client resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Client(string name, ClientArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/client:Client", name, args ?? new ClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Client(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/client:Client", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Client resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Client Get(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
        {
            return new Client(name, id, state, options);
        }
    }

    public sealed class ClientArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessTokenLifespan")]
        public Input<string>? AccessTokenLifespan { get; set; }

        [Input("accessType", required: true)]
        public Input<string> AccessType { get; set; } = null!;

        [Input("adminUrl")]
        public Input<string>? AdminUrl { get; set; }

        [Input("authenticationFlowBindingOverrides")]
        public Input<Inputs.ClientAuthenticationFlowBindingOverridesArgs>? AuthenticationFlowBindingOverrides { get; set; }

        [Input("authorization")]
        public Input<Inputs.ClientAuthorizationArgs>? Authorization { get; set; }

        [Input("backchannelLogoutRevokeOfflineSessions")]
        public Input<bool>? BackchannelLogoutRevokeOfflineSessions { get; set; }

        [Input("backchannelLogoutSessionRequired")]
        public Input<bool>? BackchannelLogoutSessionRequired { get; set; }

        [Input("backchannelLogoutUrl")]
        public Input<string>? BackchannelLogoutUrl { get; set; }

        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        [Input("clientAuthenticatorType")]
        public Input<string>? ClientAuthenticatorType { get; set; }

        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientOfflineSessionIdleTimeout")]
        public Input<string>? ClientOfflineSessionIdleTimeout { get; set; }

        [Input("clientOfflineSessionMaxLifespan")]
        public Input<string>? ClientOfflineSessionMaxLifespan { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientSessionIdleTimeout")]
        public Input<string>? ClientSessionIdleTimeout { get; set; }

        [Input("clientSessionMaxLifespan")]
        public Input<string>? ClientSessionMaxLifespan { get; set; }

        [Input("consentRequired")]
        public Input<bool>? ConsentRequired { get; set; }

        [Input("consentScreenText")]
        public Input<string>? ConsentScreenText { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("directAccessGrantsEnabled")]
        public Input<bool>? DirectAccessGrantsEnabled { get; set; }

        [Input("displayOnConsentScreen")]
        public Input<bool>? DisplayOnConsentScreen { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("excludeSessionStateFromAuthResponse")]
        public Input<bool>? ExcludeSessionStateFromAuthResponse { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        [Input("frontchannelLogoutEnabled")]
        public Input<bool>? FrontchannelLogoutEnabled { get; set; }

        [Input("frontchannelLogoutUrl")]
        public Input<string>? FrontchannelLogoutUrl { get; set; }

        [Input("fullScopeAllowed")]
        public Input<bool>? FullScopeAllowed { get; set; }

        [Input("implicitFlowEnabled")]
        public Input<bool>? ImplicitFlowEnabled { get; set; }

        [Input("import")]
        public Input<bool>? Import { get; set; }

        [Input("loginTheme")]
        public Input<string>? LoginTheme { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("oauth2DeviceAuthorizationGrantEnabled")]
        public Input<bool>? Oauth2DeviceAuthorizationGrantEnabled { get; set; }

        [Input("oauth2DeviceCodeLifespan")]
        public Input<string>? Oauth2DeviceCodeLifespan { get; set; }

        [Input("oauth2DevicePollingInterval")]
        public Input<string>? Oauth2DevicePollingInterval { get; set; }

        [Input("pkceCodeChallengeMethod")]
        public Input<string>? PkceCodeChallengeMethod { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("rootUrl")]
        public Input<string>? RootUrl { get; set; }

        [Input("serviceAccountsEnabled")]
        public Input<bool>? ServiceAccountsEnabled { get; set; }

        [Input("standardFlowEnabled")]
        public Input<bool>? StandardFlowEnabled { get; set; }

        [Input("useRefreshTokens")]
        public Input<bool>? UseRefreshTokens { get; set; }

        [Input("useRefreshTokensClientCredentials")]
        public Input<bool>? UseRefreshTokensClientCredentials { get; set; }

        [Input("validPostLogoutRedirectUris")]
        private InputList<string>? _validPostLogoutRedirectUris;
        public InputList<string> ValidPostLogoutRedirectUris
        {
            get => _validPostLogoutRedirectUris ?? (_validPostLogoutRedirectUris = new InputList<string>());
            set => _validPostLogoutRedirectUris = value;
        }

        [Input("validRedirectUris")]
        private InputList<string>? _validRedirectUris;
        public InputList<string> ValidRedirectUris
        {
            get => _validRedirectUris ?? (_validRedirectUris = new InputList<string>());
            set => _validRedirectUris = value;
        }

        [Input("webOrigins")]
        private InputList<string>? _webOrigins;
        public InputList<string> WebOrigins
        {
            get => _webOrigins ?? (_webOrigins = new InputList<string>());
            set => _webOrigins = value;
        }

        public ClientArgs()
        {
        }
        public static new ClientArgs Empty => new ClientArgs();
    }

    public sealed class ClientState : global::Pulumi.ResourceArgs
    {
        [Input("accessTokenLifespan")]
        public Input<string>? AccessTokenLifespan { get; set; }

        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        [Input("adminUrl")]
        public Input<string>? AdminUrl { get; set; }

        [Input("authenticationFlowBindingOverrides")]
        public Input<Inputs.ClientAuthenticationFlowBindingOverridesGetArgs>? AuthenticationFlowBindingOverrides { get; set; }

        [Input("authorization")]
        public Input<Inputs.ClientAuthorizationGetArgs>? Authorization { get; set; }

        [Input("backchannelLogoutRevokeOfflineSessions")]
        public Input<bool>? BackchannelLogoutRevokeOfflineSessions { get; set; }

        [Input("backchannelLogoutSessionRequired")]
        public Input<bool>? BackchannelLogoutSessionRequired { get; set; }

        [Input("backchannelLogoutUrl")]
        public Input<string>? BackchannelLogoutUrl { get; set; }

        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        [Input("clientAuthenticatorType")]
        public Input<string>? ClientAuthenticatorType { get; set; }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientOfflineSessionIdleTimeout")]
        public Input<string>? ClientOfflineSessionIdleTimeout { get; set; }

        [Input("clientOfflineSessionMaxLifespan")]
        public Input<string>? ClientOfflineSessionMaxLifespan { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientSessionIdleTimeout")]
        public Input<string>? ClientSessionIdleTimeout { get; set; }

        [Input("clientSessionMaxLifespan")]
        public Input<string>? ClientSessionMaxLifespan { get; set; }

        [Input("consentRequired")]
        public Input<bool>? ConsentRequired { get; set; }

        [Input("consentScreenText")]
        public Input<string>? ConsentScreenText { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("directAccessGrantsEnabled")]
        public Input<bool>? DirectAccessGrantsEnabled { get; set; }

        [Input("displayOnConsentScreen")]
        public Input<bool>? DisplayOnConsentScreen { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("excludeSessionStateFromAuthResponse")]
        public Input<bool>? ExcludeSessionStateFromAuthResponse { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        [Input("frontchannelLogoutEnabled")]
        public Input<bool>? FrontchannelLogoutEnabled { get; set; }

        [Input("frontchannelLogoutUrl")]
        public Input<string>? FrontchannelLogoutUrl { get; set; }

        [Input("fullScopeAllowed")]
        public Input<bool>? FullScopeAllowed { get; set; }

        [Input("implicitFlowEnabled")]
        public Input<bool>? ImplicitFlowEnabled { get; set; }

        [Input("import")]
        public Input<bool>? Import { get; set; }

        [Input("loginTheme")]
        public Input<string>? LoginTheme { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("oauth2DeviceAuthorizationGrantEnabled")]
        public Input<bool>? Oauth2DeviceAuthorizationGrantEnabled { get; set; }

        [Input("oauth2DeviceCodeLifespan")]
        public Input<string>? Oauth2DeviceCodeLifespan { get; set; }

        [Input("oauth2DevicePollingInterval")]
        public Input<string>? Oauth2DevicePollingInterval { get; set; }

        [Input("pkceCodeChallengeMethod")]
        public Input<string>? PkceCodeChallengeMethod { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("resourceServerId")]
        public Input<string>? ResourceServerId { get; set; }

        [Input("rootUrl")]
        public Input<string>? RootUrl { get; set; }

        [Input("serviceAccountUserId")]
        public Input<string>? ServiceAccountUserId { get; set; }

        [Input("serviceAccountsEnabled")]
        public Input<bool>? ServiceAccountsEnabled { get; set; }

        [Input("standardFlowEnabled")]
        public Input<bool>? StandardFlowEnabled { get; set; }

        [Input("useRefreshTokens")]
        public Input<bool>? UseRefreshTokens { get; set; }

        [Input("useRefreshTokensClientCredentials")]
        public Input<bool>? UseRefreshTokensClientCredentials { get; set; }

        [Input("validPostLogoutRedirectUris")]
        private InputList<string>? _validPostLogoutRedirectUris;
        public InputList<string> ValidPostLogoutRedirectUris
        {
            get => _validPostLogoutRedirectUris ?? (_validPostLogoutRedirectUris = new InputList<string>());
            set => _validPostLogoutRedirectUris = value;
        }

        [Input("validRedirectUris")]
        private InputList<string>? _validRedirectUris;
        public InputList<string> ValidRedirectUris
        {
            get => _validRedirectUris ?? (_validRedirectUris = new InputList<string>());
            set => _validRedirectUris = value;
        }

        [Input("webOrigins")]
        private InputList<string>? _webOrigins;
        public InputList<string> WebOrigins
        {
            get => _webOrigins ?? (_webOrigins = new InputList<string>());
            set => _webOrigins = value;
        }

        public ClientState()
        {
        }
        public static new ClientState Empty => new ClientState();
    }
}
