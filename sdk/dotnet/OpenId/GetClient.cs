// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public static class GetClient
    {
        /// <summary>
        /// This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
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
        ///     var realmManagement = Keycloak.OpenId.GetClient.Invoke(new()
        ///     {
        ///         RealmId = "my-realm",
        ///         ClientId = "realm-management",
        ///     });
        /// 
        ///     // use the data source
        ///     var admin = Keycloak.GetRole.Invoke(new()
        ///     {
        ///         RealmId = "my-realm",
        ///         ClientId = realmManagement.Apply(getClientResult =&gt; getClientResult.Id),
        ///         Name = "realm-admin",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetClientResult> InvokeAsync(GetClientArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClientResult>("keycloak:openid/getClient:getClient", args ?? new GetClientArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
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
        ///     var realmManagement = Keycloak.OpenId.GetClient.Invoke(new()
        ///     {
        ///         RealmId = "my-realm",
        ///         ClientId = "realm-management",
        ///     });
        /// 
        ///     // use the data source
        ///     var admin = Keycloak.GetRole.Invoke(new()
        ///     {
        ///         RealmId = "my-realm",
        ///         ClientId = realmManagement.Apply(getClientResult =&gt; getClientResult.Id),
        ///         Name = "realm-admin",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetClientResult> Invoke(GetClientInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClientResult>("keycloak:openid/getClient:getClient", args ?? new GetClientInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
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
        ///     var realmManagement = Keycloak.OpenId.GetClient.Invoke(new()
        ///     {
        ///         RealmId = "my-realm",
        ///         ClientId = "realm-management",
        ///     });
        /// 
        ///     // use the data source
        ///     var admin = Keycloak.GetRole.Invoke(new()
        ///     {
        ///         RealmId = "my-realm",
        ///         ClientId = realmManagement.Apply(getClientResult =&gt; getClientResult.Id),
        ///         Name = "realm-admin",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetClientResult> Invoke(GetClientInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetClientResult>("keycloak:openid/getClient:getClient", args ?? new GetClientInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClientArgs : global::Pulumi.InvokeArgs
    {
        [Input("alwaysDisplayInConsole")]
        public bool? AlwaysDisplayInConsole { get; set; }

        /// <summary>
        /// The client id (not its unique ID).
        /// </summary>
        [Input("clientId", required: true)]
        public string ClientId { get; set; } = null!;

        [Input("consentScreenText")]
        public string? ConsentScreenText { get; set; }

        [Input("displayOnConsentScreen")]
        public bool? DisplayOnConsentScreen { get; set; }

        [Input("extraConfig")]
        private Dictionary<string, string>? _extraConfig;
        public Dictionary<string, string> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new Dictionary<string, string>());
            set => _extraConfig = value;
        }

        [Input("oauth2DeviceAuthorizationGrantEnabled")]
        public bool? Oauth2DeviceAuthorizationGrantEnabled { get; set; }

        [Input("oauth2DeviceCodeLifespan")]
        public string? Oauth2DeviceCodeLifespan { get; set; }

        [Input("oauth2DevicePollingInterval")]
        public string? Oauth2DevicePollingInterval { get; set; }

        /// <summary>
        /// The realm id.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetClientArgs()
        {
        }
        public static new GetClientArgs Empty => new GetClientArgs();
    }

    public sealed class GetClientInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("alwaysDisplayInConsole")]
        public Input<bool>? AlwaysDisplayInConsole { get; set; }

        /// <summary>
        /// The client id (not its unique ID).
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("consentScreenText")]
        public Input<string>? ConsentScreenText { get; set; }

        [Input("displayOnConsentScreen")]
        public Input<bool>? DisplayOnConsentScreen { get; set; }

        [Input("extraConfig")]
        private InputMap<string>? _extraConfig;
        public InputMap<string> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<string>());
            set => _extraConfig = value;
        }

        [Input("oauth2DeviceAuthorizationGrantEnabled")]
        public Input<bool>? Oauth2DeviceAuthorizationGrantEnabled { get; set; }

        [Input("oauth2DeviceCodeLifespan")]
        public Input<string>? Oauth2DeviceCodeLifespan { get; set; }

        [Input("oauth2DevicePollingInterval")]
        public Input<string>? Oauth2DevicePollingInterval { get; set; }

        /// <summary>
        /// The realm id.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GetClientInvokeArgs()
        {
        }
        public static new GetClientInvokeArgs Empty => new GetClientInvokeArgs();
    }


    [OutputType]
    public sealed class GetClientResult
    {
        public readonly string AccessTokenLifespan;
        public readonly string AccessType;
        public readonly string AdminUrl;
        public readonly bool? AlwaysDisplayInConsole;
        public readonly ImmutableArray<Outputs.GetClientAuthenticationFlowBindingOverrideResult> AuthenticationFlowBindingOverrides;
        public readonly ImmutableArray<Outputs.GetClientAuthorizationResult> Authorizations;
        public readonly bool BackchannelLogoutRevokeOfflineSessions;
        public readonly bool BackchannelLogoutSessionRequired;
        public readonly string BackchannelLogoutUrl;
        public readonly string BaseUrl;
        public readonly string ClientAuthenticatorType;
        public readonly string ClientId;
        public readonly string ClientOfflineSessionIdleTimeout;
        public readonly string ClientOfflineSessionMaxLifespan;
        public readonly string ClientSecret;
        public readonly string ClientSessionIdleTimeout;
        public readonly string ClientSessionMaxLifespan;
        public readonly bool ConsentRequired;
        public readonly string? ConsentScreenText;
        public readonly string Description;
        public readonly bool DirectAccessGrantsEnabled;
        public readonly bool? DisplayOnConsentScreen;
        public readonly bool Enabled;
        public readonly bool ExcludeIssuerFromAuthResponse;
        public readonly bool ExcludeSessionStateFromAuthResponse;
        public readonly ImmutableDictionary<string, string> ExtraConfig;
        public readonly bool FrontchannelLogoutEnabled;
        public readonly string FrontchannelLogoutUrl;
        public readonly bool FullScopeAllowed;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool ImplicitFlowEnabled;
        public readonly string LoginTheme;
        public readonly string Name;
        public readonly bool? Oauth2DeviceAuthorizationGrantEnabled;
        public readonly string? Oauth2DeviceCodeLifespan;
        public readonly string? Oauth2DevicePollingInterval;
        public readonly string PkceCodeChallengeMethod;
        public readonly string RealmId;
        public readonly string ResourceServerId;
        public readonly string RootUrl;
        public readonly string ServiceAccountUserId;
        public readonly bool ServiceAccountsEnabled;
        public readonly bool StandardFlowEnabled;
        public readonly bool UseRefreshTokens;
        public readonly bool UseRefreshTokensClientCredentials;
        public readonly ImmutableArray<string> ValidPostLogoutRedirectUris;
        public readonly ImmutableArray<string> ValidRedirectUris;
        public readonly ImmutableArray<string> WebOrigins;

        [OutputConstructor]
        private GetClientResult(
            string accessTokenLifespan,

            string accessType,

            string adminUrl,

            bool? alwaysDisplayInConsole,

            ImmutableArray<Outputs.GetClientAuthenticationFlowBindingOverrideResult> authenticationFlowBindingOverrides,

            ImmutableArray<Outputs.GetClientAuthorizationResult> authorizations,

            bool backchannelLogoutRevokeOfflineSessions,

            bool backchannelLogoutSessionRequired,

            string backchannelLogoutUrl,

            string baseUrl,

            string clientAuthenticatorType,

            string clientId,

            string clientOfflineSessionIdleTimeout,

            string clientOfflineSessionMaxLifespan,

            string clientSecret,

            string clientSessionIdleTimeout,

            string clientSessionMaxLifespan,

            bool consentRequired,

            string? consentScreenText,

            string description,

            bool directAccessGrantsEnabled,

            bool? displayOnConsentScreen,

            bool enabled,

            bool excludeIssuerFromAuthResponse,

            bool excludeSessionStateFromAuthResponse,

            ImmutableDictionary<string, string> extraConfig,

            bool frontchannelLogoutEnabled,

            string frontchannelLogoutUrl,

            bool fullScopeAllowed,

            string id,

            bool implicitFlowEnabled,

            string loginTheme,

            string name,

            bool? oauth2DeviceAuthorizationGrantEnabled,

            string? oauth2DeviceCodeLifespan,

            string? oauth2DevicePollingInterval,

            string pkceCodeChallengeMethod,

            string realmId,

            string resourceServerId,

            string rootUrl,

            string serviceAccountUserId,

            bool serviceAccountsEnabled,

            bool standardFlowEnabled,

            bool useRefreshTokens,

            bool useRefreshTokensClientCredentials,

            ImmutableArray<string> validPostLogoutRedirectUris,

            ImmutableArray<string> validRedirectUris,

            ImmutableArray<string> webOrigins)
        {
            AccessTokenLifespan = accessTokenLifespan;
            AccessType = accessType;
            AdminUrl = adminUrl;
            AlwaysDisplayInConsole = alwaysDisplayInConsole;
            AuthenticationFlowBindingOverrides = authenticationFlowBindingOverrides;
            Authorizations = authorizations;
            BackchannelLogoutRevokeOfflineSessions = backchannelLogoutRevokeOfflineSessions;
            BackchannelLogoutSessionRequired = backchannelLogoutSessionRequired;
            BackchannelLogoutUrl = backchannelLogoutUrl;
            BaseUrl = baseUrl;
            ClientAuthenticatorType = clientAuthenticatorType;
            ClientId = clientId;
            ClientOfflineSessionIdleTimeout = clientOfflineSessionIdleTimeout;
            ClientOfflineSessionMaxLifespan = clientOfflineSessionMaxLifespan;
            ClientSecret = clientSecret;
            ClientSessionIdleTimeout = clientSessionIdleTimeout;
            ClientSessionMaxLifespan = clientSessionMaxLifespan;
            ConsentRequired = consentRequired;
            ConsentScreenText = consentScreenText;
            Description = description;
            DirectAccessGrantsEnabled = directAccessGrantsEnabled;
            DisplayOnConsentScreen = displayOnConsentScreen;
            Enabled = enabled;
            ExcludeIssuerFromAuthResponse = excludeIssuerFromAuthResponse;
            ExcludeSessionStateFromAuthResponse = excludeSessionStateFromAuthResponse;
            ExtraConfig = extraConfig;
            FrontchannelLogoutEnabled = frontchannelLogoutEnabled;
            FrontchannelLogoutUrl = frontchannelLogoutUrl;
            FullScopeAllowed = fullScopeAllowed;
            Id = id;
            ImplicitFlowEnabled = implicitFlowEnabled;
            LoginTheme = loginTheme;
            Name = name;
            Oauth2DeviceAuthorizationGrantEnabled = oauth2DeviceAuthorizationGrantEnabled;
            Oauth2DeviceCodeLifespan = oauth2DeviceCodeLifespan;
            Oauth2DevicePollingInterval = oauth2DevicePollingInterval;
            PkceCodeChallengeMethod = pkceCodeChallengeMethod;
            RealmId = realmId;
            ResourceServerId = resourceServerId;
            RootUrl = rootUrl;
            ServiceAccountUserId = serviceAccountUserId;
            ServiceAccountsEnabled = serviceAccountsEnabled;
            StandardFlowEnabled = standardFlowEnabled;
            UseRefreshTokens = useRefreshTokens;
            UseRefreshTokensClientCredentials = useRefreshTokensClientCredentials;
            ValidPostLogoutRedirectUris = validPostLogoutRedirectUris;
            ValidRedirectUris = validRedirectUris;
            WebOrigins = webOrigins;
        }
    }
}
