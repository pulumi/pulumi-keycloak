// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public static class GetRealm
    {
        /// <summary>
        /// ## # keycloak.Realm data source
        /// 
        /// This data source can be used to fetch properties of a Keycloak realm for
        /// usage with other resources.
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
        ///     var realm = Keycloak.GetRealm.Invoke(new()
        ///     {
        ///         Realm = "my-realm",
        ///     });
        /// 
        ///     // use the data source
        ///     var @group = new Keycloak.Role("group", new()
        ///     {
        ///         RealmId = id,
        ///         Name = "group",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### Argument Reference
        /// 
        /// The following arguments are supported:
        /// 
        /// - `realm` - (Required) The realm name.
        /// 
        /// ### Attributes Reference
        /// 
        /// See the docs for the `keycloak.Realm` resource for details on the exported attributes.
        /// </summary>
        public static Task<GetRealmResult> InvokeAsync(GetRealmArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRealmResult>("keycloak:index/getRealm:getRealm", args ?? new GetRealmArgs(), options.WithDefaults());

        /// <summary>
        /// ## # keycloak.Realm data source
        /// 
        /// This data source can be used to fetch properties of a Keycloak realm for
        /// usage with other resources.
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
        ///     var realm = Keycloak.GetRealm.Invoke(new()
        ///     {
        ///         Realm = "my-realm",
        ///     });
        /// 
        ///     // use the data source
        ///     var @group = new Keycloak.Role("group", new()
        ///     {
        ///         RealmId = id,
        ///         Name = "group",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ### Argument Reference
        /// 
        /// The following arguments are supported:
        /// 
        /// - `realm` - (Required) The realm name.
        /// 
        /// ### Attributes Reference
        /// 
        /// See the docs for the `keycloak.Realm` resource for details on the exported attributes.
        /// </summary>
        public static Output<GetRealmResult> Invoke(GetRealmInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRealmResult>("keycloak:index/getRealm:getRealm", args ?? new GetRealmInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRealmArgs : global::Pulumi.InvokeArgs
    {
        [Input("attributes")]
        private Dictionary<string, object>? _attributes;
        public Dictionary<string, object> Attributes
        {
            get => _attributes ?? (_attributes = new Dictionary<string, object>());
            set => _attributes = value;
        }

        [Input("defaultDefaultClientScopes")]
        private List<string>? _defaultDefaultClientScopes;
        public List<string> DefaultDefaultClientScopes
        {
            get => _defaultDefaultClientScopes ?? (_defaultDefaultClientScopes = new List<string>());
            set => _defaultDefaultClientScopes = value;
        }

        [Input("defaultOptionalClientScopes")]
        private List<string>? _defaultOptionalClientScopes;
        public List<string> DefaultOptionalClientScopes
        {
            get => _defaultOptionalClientScopes ?? (_defaultOptionalClientScopes = new List<string>());
            set => _defaultOptionalClientScopes = value;
        }

        [Input("displayNameHtml")]
        public string? DisplayNameHtml { get; set; }

        [Input("internationalizations")]
        private List<Inputs.GetRealmInternationalizationArgs>? _internationalizations;
        public List<Inputs.GetRealmInternationalizationArgs> Internationalizations
        {
            get => _internationalizations ?? (_internationalizations = new List<Inputs.GetRealmInternationalizationArgs>());
            set => _internationalizations = value;
        }

        [Input("otpPolicy")]
        public Inputs.GetRealmOtpPolicyArgs? OtpPolicy { get; set; }

        [Input("realm", required: true)]
        public string Realm { get; set; } = null!;

        [Input("securityDefenses")]
        private List<Inputs.GetRealmSecurityDefenseArgs>? _securityDefenses;
        public List<Inputs.GetRealmSecurityDefenseArgs> SecurityDefenses
        {
            get => _securityDefenses ?? (_securityDefenses = new List<Inputs.GetRealmSecurityDefenseArgs>());
            set => _securityDefenses = value;
        }

        [Input("smtpServers")]
        private List<Inputs.GetRealmSmtpServerArgs>? _smtpServers;
        public List<Inputs.GetRealmSmtpServerArgs> SmtpServers
        {
            get => _smtpServers ?? (_smtpServers = new List<Inputs.GetRealmSmtpServerArgs>());
            set => _smtpServers = value;
        }

        [Input("webAuthnPasswordlessPolicy")]
        public Inputs.GetRealmWebAuthnPasswordlessPolicyArgs? WebAuthnPasswordlessPolicy { get; set; }

        [Input("webAuthnPolicy")]
        public Inputs.GetRealmWebAuthnPolicyArgs? WebAuthnPolicy { get; set; }

        public GetRealmArgs()
        {
        }
        public static new GetRealmArgs Empty => new GetRealmArgs();
    }

    public sealed class GetRealmInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("attributes")]
        private InputMap<object>? _attributes;
        public InputMap<object> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<object>());
            set => _attributes = value;
        }

        [Input("defaultDefaultClientScopes")]
        private InputList<string>? _defaultDefaultClientScopes;
        public InputList<string> DefaultDefaultClientScopes
        {
            get => _defaultDefaultClientScopes ?? (_defaultDefaultClientScopes = new InputList<string>());
            set => _defaultDefaultClientScopes = value;
        }

        [Input("defaultOptionalClientScopes")]
        private InputList<string>? _defaultOptionalClientScopes;
        public InputList<string> DefaultOptionalClientScopes
        {
            get => _defaultOptionalClientScopes ?? (_defaultOptionalClientScopes = new InputList<string>());
            set => _defaultOptionalClientScopes = value;
        }

        [Input("displayNameHtml")]
        public Input<string>? DisplayNameHtml { get; set; }

        [Input("internationalizations")]
        private InputList<Inputs.GetRealmInternationalizationInputArgs>? _internationalizations;
        public InputList<Inputs.GetRealmInternationalizationInputArgs> Internationalizations
        {
            get => _internationalizations ?? (_internationalizations = new InputList<Inputs.GetRealmInternationalizationInputArgs>());
            set => _internationalizations = value;
        }

        [Input("otpPolicy")]
        public Input<Inputs.GetRealmOtpPolicyInputArgs>? OtpPolicy { get; set; }

        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        [Input("securityDefenses")]
        private InputList<Inputs.GetRealmSecurityDefenseInputArgs>? _securityDefenses;
        public InputList<Inputs.GetRealmSecurityDefenseInputArgs> SecurityDefenses
        {
            get => _securityDefenses ?? (_securityDefenses = new InputList<Inputs.GetRealmSecurityDefenseInputArgs>());
            set => _securityDefenses = value;
        }

        [Input("smtpServers")]
        private InputList<Inputs.GetRealmSmtpServerInputArgs>? _smtpServers;
        public InputList<Inputs.GetRealmSmtpServerInputArgs> SmtpServers
        {
            get => _smtpServers ?? (_smtpServers = new InputList<Inputs.GetRealmSmtpServerInputArgs>());
            set => _smtpServers = value;
        }

        [Input("webAuthnPasswordlessPolicy")]
        public Input<Inputs.GetRealmWebAuthnPasswordlessPolicyInputArgs>? WebAuthnPasswordlessPolicy { get; set; }

        [Input("webAuthnPolicy")]
        public Input<Inputs.GetRealmWebAuthnPolicyInputArgs>? WebAuthnPolicy { get; set; }

        public GetRealmInvokeArgs()
        {
        }
        public static new GetRealmInvokeArgs Empty => new GetRealmInvokeArgs();
    }


    [OutputType]
    public sealed class GetRealmResult
    {
        public readonly string AccessCodeLifespan;
        public readonly string AccessCodeLifespanLogin;
        public readonly string AccessCodeLifespanUserAction;
        public readonly string AccessTokenLifespan;
        public readonly string AccessTokenLifespanForImplicitFlow;
        public readonly string AccountTheme;
        public readonly string ActionTokenGeneratedByAdminLifespan;
        public readonly string ActionTokenGeneratedByUserLifespan;
        public readonly string AdminTheme;
        public readonly ImmutableDictionary<string, object> Attributes;
        public readonly string BrowserFlow;
        public readonly string ClientAuthenticationFlow;
        public readonly string ClientSessionIdleTimeout;
        public readonly string ClientSessionMaxLifespan;
        public readonly ImmutableArray<string> DefaultDefaultClientScopes;
        public readonly ImmutableArray<string> DefaultOptionalClientScopes;
        public readonly string DefaultSignatureAlgorithm;
        public readonly string DirectGrantFlow;
        public readonly string DisplayName;
        public readonly string? DisplayNameHtml;
        public readonly string DockerAuthenticationFlow;
        public readonly bool DuplicateEmailsAllowed;
        public readonly bool EditUsernameAllowed;
        public readonly string EmailTheme;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InternalId;
        public readonly ImmutableArray<Outputs.GetRealmInternationalizationResult> Internationalizations;
        public readonly string LoginTheme;
        public readonly bool LoginWithEmailAllowed;
        public readonly string Oauth2DeviceCodeLifespan;
        public readonly int Oauth2DevicePollingInterval;
        public readonly string OfflineSessionIdleTimeout;
        public readonly string OfflineSessionMaxLifespan;
        public readonly bool OfflineSessionMaxLifespanEnabled;
        public readonly Outputs.GetRealmOtpPolicyResult OtpPolicy;
        public readonly string PasswordPolicy;
        public readonly string Realm;
        public readonly int RefreshTokenMaxReuse;
        public readonly bool RegistrationAllowed;
        public readonly bool RegistrationEmailAsUsername;
        public readonly string RegistrationFlow;
        public readonly bool RememberMe;
        public readonly string ResetCredentialsFlow;
        public readonly bool ResetPasswordAllowed;
        public readonly bool RevokeRefreshToken;
        public readonly ImmutableArray<Outputs.GetRealmSecurityDefenseResult> SecurityDefenses;
        public readonly ImmutableArray<Outputs.GetRealmSmtpServerResult> SmtpServers;
        public readonly string SslRequired;
        public readonly string SsoSessionIdleTimeout;
        public readonly string SsoSessionIdleTimeoutRememberMe;
        public readonly string SsoSessionMaxLifespan;
        public readonly string SsoSessionMaxLifespanRememberMe;
        public readonly bool UserManagedAccess;
        public readonly bool VerifyEmail;
        public readonly Outputs.GetRealmWebAuthnPasswordlessPolicyResult WebAuthnPasswordlessPolicy;
        public readonly Outputs.GetRealmWebAuthnPolicyResult WebAuthnPolicy;

        [OutputConstructor]
        private GetRealmResult(
            string accessCodeLifespan,

            string accessCodeLifespanLogin,

            string accessCodeLifespanUserAction,

            string accessTokenLifespan,

            string accessTokenLifespanForImplicitFlow,

            string accountTheme,

            string actionTokenGeneratedByAdminLifespan,

            string actionTokenGeneratedByUserLifespan,

            string adminTheme,

            ImmutableDictionary<string, object> attributes,

            string browserFlow,

            string clientAuthenticationFlow,

            string clientSessionIdleTimeout,

            string clientSessionMaxLifespan,

            ImmutableArray<string> defaultDefaultClientScopes,

            ImmutableArray<string> defaultOptionalClientScopes,

            string defaultSignatureAlgorithm,

            string directGrantFlow,

            string displayName,

            string? displayNameHtml,

            string dockerAuthenticationFlow,

            bool duplicateEmailsAllowed,

            bool editUsernameAllowed,

            string emailTheme,

            bool enabled,

            string id,

            string internalId,

            ImmutableArray<Outputs.GetRealmInternationalizationResult> internationalizations,

            string loginTheme,

            bool loginWithEmailAllowed,

            string oauth2DeviceCodeLifespan,

            int oauth2DevicePollingInterval,

            string offlineSessionIdleTimeout,

            string offlineSessionMaxLifespan,

            bool offlineSessionMaxLifespanEnabled,

            Outputs.GetRealmOtpPolicyResult otpPolicy,

            string passwordPolicy,

            string realm,

            int refreshTokenMaxReuse,

            bool registrationAllowed,

            bool registrationEmailAsUsername,

            string registrationFlow,

            bool rememberMe,

            string resetCredentialsFlow,

            bool resetPasswordAllowed,

            bool revokeRefreshToken,

            ImmutableArray<Outputs.GetRealmSecurityDefenseResult> securityDefenses,

            ImmutableArray<Outputs.GetRealmSmtpServerResult> smtpServers,

            string sslRequired,

            string ssoSessionIdleTimeout,

            string ssoSessionIdleTimeoutRememberMe,

            string ssoSessionMaxLifespan,

            string ssoSessionMaxLifespanRememberMe,

            bool userManagedAccess,

            bool verifyEmail,

            Outputs.GetRealmWebAuthnPasswordlessPolicyResult webAuthnPasswordlessPolicy,

            Outputs.GetRealmWebAuthnPolicyResult webAuthnPolicy)
        {
            AccessCodeLifespan = accessCodeLifespan;
            AccessCodeLifespanLogin = accessCodeLifespanLogin;
            AccessCodeLifespanUserAction = accessCodeLifespanUserAction;
            AccessTokenLifespan = accessTokenLifespan;
            AccessTokenLifespanForImplicitFlow = accessTokenLifespanForImplicitFlow;
            AccountTheme = accountTheme;
            ActionTokenGeneratedByAdminLifespan = actionTokenGeneratedByAdminLifespan;
            ActionTokenGeneratedByUserLifespan = actionTokenGeneratedByUserLifespan;
            AdminTheme = adminTheme;
            Attributes = attributes;
            BrowserFlow = browserFlow;
            ClientAuthenticationFlow = clientAuthenticationFlow;
            ClientSessionIdleTimeout = clientSessionIdleTimeout;
            ClientSessionMaxLifespan = clientSessionMaxLifespan;
            DefaultDefaultClientScopes = defaultDefaultClientScopes;
            DefaultOptionalClientScopes = defaultOptionalClientScopes;
            DefaultSignatureAlgorithm = defaultSignatureAlgorithm;
            DirectGrantFlow = directGrantFlow;
            DisplayName = displayName;
            DisplayNameHtml = displayNameHtml;
            DockerAuthenticationFlow = dockerAuthenticationFlow;
            DuplicateEmailsAllowed = duplicateEmailsAllowed;
            EditUsernameAllowed = editUsernameAllowed;
            EmailTheme = emailTheme;
            Enabled = enabled;
            Id = id;
            InternalId = internalId;
            Internationalizations = internationalizations;
            LoginTheme = loginTheme;
            LoginWithEmailAllowed = loginWithEmailAllowed;
            Oauth2DeviceCodeLifespan = oauth2DeviceCodeLifespan;
            Oauth2DevicePollingInterval = oauth2DevicePollingInterval;
            OfflineSessionIdleTimeout = offlineSessionIdleTimeout;
            OfflineSessionMaxLifespan = offlineSessionMaxLifespan;
            OfflineSessionMaxLifespanEnabled = offlineSessionMaxLifespanEnabled;
            OtpPolicy = otpPolicy;
            PasswordPolicy = passwordPolicy;
            Realm = realm;
            RefreshTokenMaxReuse = refreshTokenMaxReuse;
            RegistrationAllowed = registrationAllowed;
            RegistrationEmailAsUsername = registrationEmailAsUsername;
            RegistrationFlow = registrationFlow;
            RememberMe = rememberMe;
            ResetCredentialsFlow = resetCredentialsFlow;
            ResetPasswordAllowed = resetPasswordAllowed;
            RevokeRefreshToken = revokeRefreshToken;
            SecurityDefenses = securityDefenses;
            SmtpServers = smtpServers;
            SslRequired = sslRequired;
            SsoSessionIdleTimeout = ssoSessionIdleTimeout;
            SsoSessionIdleTimeoutRememberMe = ssoSessionIdleTimeoutRememberMe;
            SsoSessionMaxLifespan = ssoSessionMaxLifespan;
            SsoSessionMaxLifespanRememberMe = ssoSessionMaxLifespanRememberMe;
            UserManagedAccess = userManagedAccess;
            VerifyEmail = verifyEmail;
            WebAuthnPasswordlessPolicy = webAuthnPasswordlessPolicy;
            WebAuthnPolicy = webAuthnPolicy;
        }
    }
}
