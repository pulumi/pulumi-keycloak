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
        /// using Pulumi;
        /// using Keycloak = Pulumi.Keycloak;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var realm = Output.Create(Keycloak.GetRealm.InvokeAsync(new Keycloak.GetRealmArgs
        ///         {
        ///             Realm = "my-realm",
        ///         }));
        ///         var @group = new Keycloak.Role("group", new Keycloak.RoleArgs
        ///         {
        ///             RealmId = data.Keycloak_realm.Id,
        ///         });
        ///     }
        /// 
        /// }
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
            => Pulumi.Deployment.Instance.InvokeAsync<GetRealmResult>("keycloak:index/getRealm:getRealm", args ?? new GetRealmArgs(), options.WithVersion());
    }


    public sealed class GetRealmArgs : Pulumi.InvokeArgs
    {
        [Input("attributes")]
        private Dictionary<string, object>? _attributes;
        public Dictionary<string, object> Attributes
        {
            get => _attributes ?? (_attributes = new Dictionary<string, object>());
            set => _attributes = value;
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

        public GetRealmArgs()
        {
        }
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
        public readonly string OfflineSessionIdleTimeout;
        public readonly string OfflineSessionMaxLifespan;
        public readonly string PasswordPolicy;
        public readonly string Realm;
        public readonly int RefreshTokenMaxReuse;
        public readonly bool RegistrationAllowed;
        public readonly bool RegistrationEmailAsUsername;
        public readonly string RegistrationFlow;
        public readonly bool RememberMe;
        public readonly string ResetCredentialsFlow;
        public readonly bool ResetPasswordAllowed;
        public readonly ImmutableArray<Outputs.GetRealmSecurityDefenseResult> SecurityDefenses;
        public readonly ImmutableArray<Outputs.GetRealmSmtpServerResult> SmtpServers;
        public readonly string SslRequired;
        public readonly string SsoSessionIdleTimeout;
        public readonly string SsoSessionMaxLifespan;
        public readonly bool UserManagedAccess;
        public readonly bool VerifyEmail;

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

            string offlineSessionIdleTimeout,

            string offlineSessionMaxLifespan,

            string passwordPolicy,

            string realm,

            int refreshTokenMaxReuse,

            bool registrationAllowed,

            bool registrationEmailAsUsername,

            string registrationFlow,

            bool rememberMe,

            string resetCredentialsFlow,

            bool resetPasswordAllowed,

            ImmutableArray<Outputs.GetRealmSecurityDefenseResult> securityDefenses,

            ImmutableArray<Outputs.GetRealmSmtpServerResult> smtpServers,

            string sslRequired,

            string ssoSessionIdleTimeout,

            string ssoSessionMaxLifespan,

            bool userManagedAccess,

            bool verifyEmail)
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
            OfflineSessionIdleTimeout = offlineSessionIdleTimeout;
            OfflineSessionMaxLifespan = offlineSessionMaxLifespan;
            PasswordPolicy = passwordPolicy;
            Realm = realm;
            RefreshTokenMaxReuse = refreshTokenMaxReuse;
            RegistrationAllowed = registrationAllowed;
            RegistrationEmailAsUsername = registrationEmailAsUsername;
            RegistrationFlow = registrationFlow;
            RememberMe = rememberMe;
            ResetCredentialsFlow = resetCredentialsFlow;
            ResetPasswordAllowed = resetPasswordAllowed;
            SecurityDefenses = securityDefenses;
            SmtpServers = smtpServers;
            SslRequired = sslRequired;
            SsoSessionIdleTimeout = ssoSessionIdleTimeout;
            SsoSessionMaxLifespan = ssoSessionMaxLifespan;
            UserManagedAccess = userManagedAccess;
            VerifyEmail = verifyEmail;
        }
    }
}
