// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.RealmArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.RealmState;
import com.pulumi.keycloak.outputs.RealmInternationalization;
import com.pulumi.keycloak.outputs.RealmOtpPolicy;
import com.pulumi.keycloak.outputs.RealmSecurityDefenses;
import com.pulumi.keycloak.outputs.RealmSmtpServer;
import com.pulumi.keycloak.outputs.RealmWebAuthnPasswordlessPolicy;
import com.pulumi.keycloak.outputs.RealmWebAuthnPolicy;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing Realms within Keycloak.
 * 
 * A realm manages a logical collection of users, credentials, roles, and groups. Users log in to realms and can be federated
 * from multiple sources.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.inputs.RealmInternationalizationArgs;
 * import com.pulumi.keycloak.inputs.RealmSecurityDefensesArgs;
 * import com.pulumi.keycloak.inputs.RealmSecurityDefensesBruteForceDetectionArgs;
 * import com.pulumi.keycloak.inputs.RealmSecurityDefensesHeadersArgs;
 * import com.pulumi.keycloak.inputs.RealmSmtpServerArgs;
 * import com.pulumi.keycloak.inputs.RealmSmtpServerAuthArgs;
 * import com.pulumi.keycloak.inputs.RealmWebAuthnPolicyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var realm = new Realm(&#34;realm&#34;, RealmArgs.builder()        
 *             .accessCodeLifespan(&#34;1h&#34;)
 *             .attributes(Map.of(&#34;mycustomAttribute&#34;, &#34;myCustomValue&#34;))
 *             .displayName(&#34;my realm&#34;)
 *             .displayNameHtml(&#34;&lt;b&gt;my realm&lt;/b&gt;&#34;)
 *             .enabled(true)
 *             .internationalization(RealmInternationalizationArgs.builder()
 *                 .defaultLocale(&#34;en&#34;)
 *                 .supportedLocales(                
 *                     &#34;en&#34;,
 *                     &#34;de&#34;,
 *                     &#34;es&#34;)
 *                 .build())
 *             .loginTheme(&#34;base&#34;)
 *             .passwordPolicy(&#34;upperCase(1) and length(8) and forceExpiredPasswordChange(365) and notUsername&#34;)
 *             .realm(&#34;my-realm&#34;)
 *             .securityDefenses(RealmSecurityDefensesArgs.builder()
 *                 .bruteForceDetection(RealmSecurityDefensesBruteForceDetectionArgs.builder()
 *                     .failureResetTimeSeconds(43200)
 *                     .maxFailureWaitSeconds(900)
 *                     .maxLoginFailures(30)
 *                     .minimumQuickLoginWaitSeconds(60)
 *                     .permanentLockout(false)
 *                     .quickLoginCheckMilliSeconds(1000)
 *                     .waitIncrementSeconds(60)
 *                     .build())
 *                 .headers(RealmSecurityDefensesHeadersArgs.builder()
 *                     .contentSecurityPolicy(&#34;frame-src &#39;self&#39;; frame-ancestors &#39;self&#39;; object-src &#39;none&#39;;&#34;)
 *                     .contentSecurityPolicyReportOnly(&#34;&#34;)
 *                     .strictTransportSecurity(&#34;max-age=31536000; includeSubDomains&#34;)
 *                     .xContentTypeOptions(&#34;nosniff&#34;)
 *                     .xFrameOptions(&#34;DENY&#34;)
 *                     .xRobotsTag(&#34;none&#34;)
 *                     .xXssProtection(&#34;1; mode=block&#34;)
 *                     .build())
 *                 .build())
 *             .smtpServer(RealmSmtpServerArgs.builder()
 *                 .auth(RealmSmtpServerAuthArgs.builder()
 *                     .password(&#34;password&#34;)
 *                     .username(&#34;tom&#34;)
 *                     .build())
 *                 .from(&#34;example@example.com&#34;)
 *                 .host(&#34;smtp.example.com&#34;)
 *                 .build())
 *             .sslRequired(&#34;external&#34;)
 *             .webAuthnPolicy(RealmWebAuthnPolicyArgs.builder()
 *                 .relyingPartyEntityName(&#34;Example&#34;)
 *                 .relyingPartyId(&#34;keycloak.example.com&#34;)
 *                 .signatureAlgorithms(                
 *                     &#34;ES256&#34;,
 *                     &#34;RS256&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Default Client Scopes
 * 
 * - `default_default_client_scopes` - (Optional) A list of default default client scopes to be used for client definitions. Defaults to `[]` or keycloak&#39;s built-in default default client-scopes.
 * - `default_optional_client_scopes` - (Optional) A list of default optional client scopes to be used for client definitions. Defaults to `[]` or keycloak&#39;s built-in default optional client-scopes.
 * 
 * ## Import
 * 
 * Realms can be imported using their name. Examplebash
 * 
 * ```sh
 *  $ pulumi import keycloak:index/realm:Realm realm my-realm
 * ```
 * 
 */
@ResourceType(type="keycloak:index/realm:Realm")
public class Realm extends com.pulumi.resources.CustomResource {
    /**
     * The maximum amount of time a client has to finish the authorization code flow.
     * 
     */
    @Export(name="accessCodeLifespan", type=String.class, parameters={})
    private Output<String> accessCodeLifespan;

    /**
     * @return The maximum amount of time a client has to finish the authorization code flow.
     * 
     */
    public Output<String> accessCodeLifespan() {
        return this.accessCodeLifespan;
    }
    /**
     * The maximum amount of time a user is permitted to stay on the login page before the authentication process must be restarted.
     * 
     */
    @Export(name="accessCodeLifespanLogin", type=String.class, parameters={})
    private Output<String> accessCodeLifespanLogin;

    /**
     * @return The maximum amount of time a user is permitted to stay on the login page before the authentication process must be restarted.
     * 
     */
    public Output<String> accessCodeLifespanLogin() {
        return this.accessCodeLifespanLogin;
    }
    /**
     * The maximum amount of time a user has to complete login related actions, such as updating a password.
     * 
     */
    @Export(name="accessCodeLifespanUserAction", type=String.class, parameters={})
    private Output<String> accessCodeLifespanUserAction;

    /**
     * @return The maximum amount of time a user has to complete login related actions, such as updating a password.
     * 
     */
    public Output<String> accessCodeLifespanUserAction() {
        return this.accessCodeLifespanUserAction;
    }
    /**
     * The amount of time an access token can be used before it expires.
     * 
     */
    @Export(name="accessTokenLifespan", type=String.class, parameters={})
    private Output<String> accessTokenLifespan;

    /**
     * @return The amount of time an access token can be used before it expires.
     * 
     */
    public Output<String> accessTokenLifespan() {
        return this.accessTokenLifespan;
    }
    /**
     * The amount of time an access token issued with the OpenID Connect Implicit Flow can be used before it expires.
     * 
     */
    @Export(name="accessTokenLifespanForImplicitFlow", type=String.class, parameters={})
    private Output<String> accessTokenLifespanForImplicitFlow;

    /**
     * @return The amount of time an access token issued with the OpenID Connect Implicit Flow can be used before it expires.
     * 
     */
    public Output<String> accessTokenLifespanForImplicitFlow() {
        return this.accessTokenLifespanForImplicitFlow;
    }
    /**
     * Used for account management pages.
     * 
     */
    @Export(name="accountTheme", type=String.class, parameters={})
    private Output</* @Nullable */ String> accountTheme;

    /**
     * @return Used for account management pages.
     * 
     */
    public Output<Optional<String>> accountTheme() {
        return Codegen.optional(this.accountTheme);
    }
    /**
     * The maximum time a user has to use an admin-generated permit before it expires.
     * 
     */
    @Export(name="actionTokenGeneratedByAdminLifespan", type=String.class, parameters={})
    private Output<String> actionTokenGeneratedByAdminLifespan;

    /**
     * @return The maximum time a user has to use an admin-generated permit before it expires.
     * 
     */
    public Output<String> actionTokenGeneratedByAdminLifespan() {
        return this.actionTokenGeneratedByAdminLifespan;
    }
    /**
     * The maximum time a user has to use a user-generated permit before it expires.
     * 
     */
    @Export(name="actionTokenGeneratedByUserLifespan", type=String.class, parameters={})
    private Output<String> actionTokenGeneratedByUserLifespan;

    /**
     * @return The maximum time a user has to use a user-generated permit before it expires.
     * 
     */
    public Output<String> actionTokenGeneratedByUserLifespan() {
        return this.actionTokenGeneratedByUserLifespan;
    }
    /**
     * Used for the admin console.
     * 
     */
    @Export(name="adminTheme", type=String.class, parameters={})
    private Output</* @Nullable */ String> adminTheme;

    /**
     * @return Used for the admin console.
     * 
     */
    public Output<Optional<String>> adminTheme() {
        return Codegen.optional(this.adminTheme);
    }
    /**
     * A map of custom attributes to add to the realm.
     * 
     */
    @Export(name="attributes", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> attributes;

    /**
     * @return A map of custom attributes to add to the realm.
     * 
     */
    public Output<Optional<Map<String,Object>>> attributes() {
        return Codegen.optional(this.attributes);
    }
    /**
     * The desired flow for browser authentication. Defaults to `browser`.
     * 
     */
    @Export(name="browserFlow", type=String.class, parameters={})
    private Output<String> browserFlow;

    /**
     * @return The desired flow for browser authentication. Defaults to `browser`.
     * 
     */
    public Output<String> browserFlow() {
        return this.browserFlow;
    }
    /**
     * The desired flow for client authentication. Defaults to `clients`.
     * 
     */
    @Export(name="clientAuthenticationFlow", type=String.class, parameters={})
    private Output<String> clientAuthenticationFlow;

    /**
     * @return The desired flow for client authentication. Defaults to `clients`.
     * 
     */
    public Output<String> clientAuthenticationFlow() {
        return this.clientAuthenticationFlow;
    }
    /**
     * The amount of time a session can be idle before it expires. Users can override it for individual clients.
     * 
     */
    @Export(name="clientSessionIdleTimeout", type=String.class, parameters={})
    private Output<String> clientSessionIdleTimeout;

    /**
     * @return The amount of time a session can be idle before it expires. Users can override it for individual clients.
     * 
     */
    public Output<String> clientSessionIdleTimeout() {
        return this.clientSessionIdleTimeout;
    }
    /**
     * The maximum amount of time before a session expires regardless of activity. Users can override it for individual clients.
     * 
     */
    @Export(name="clientSessionMaxLifespan", type=String.class, parameters={})
    private Output<String> clientSessionMaxLifespan;

    /**
     * @return The maximum amount of time before a session expires regardless of activity. Users can override it for individual clients.
     * 
     */
    public Output<String> clientSessionMaxLifespan() {
        return this.clientSessionMaxLifespan;
    }
    @Export(name="defaultDefaultClientScopes", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> defaultDefaultClientScopes;

    public Output<Optional<List<String>>> defaultDefaultClientScopes() {
        return Codegen.optional(this.defaultDefaultClientScopes);
    }
    @Export(name="defaultOptionalClientScopes", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> defaultOptionalClientScopes;

    public Output<Optional<List<String>>> defaultOptionalClientScopes() {
        return Codegen.optional(this.defaultOptionalClientScopes);
    }
    /**
     * Default algorithm used to sign tokens for the realm.
     * 
     */
    @Export(name="defaultSignatureAlgorithm", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultSignatureAlgorithm;

    /**
     * @return Default algorithm used to sign tokens for the realm.
     * 
     */
    public Output<Optional<String>> defaultSignatureAlgorithm() {
        return Codegen.optional(this.defaultSignatureAlgorithm);
    }
    /**
     * The desired flow for direct access authentication. Defaults to `direct grant`.
     * 
     */
    @Export(name="directGrantFlow", type=String.class, parameters={})
    private Output<String> directGrantFlow;

    /**
     * @return The desired flow for direct access authentication. Defaults to `direct grant`.
     * 
     */
    public Output<String> directGrantFlow() {
        return this.directGrantFlow;
    }
    /**
     * The display name for the realm that is shown when logging in to the admin console.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The display name for the realm that is shown when logging in to the admin console.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The display name for the realm that is rendered as HTML on the screen when logging in to the admin console.
     * 
     */
    @Export(name="displayNameHtml", type=String.class, parameters={})
    private Output</* @Nullable */ String> displayNameHtml;

    /**
     * @return The display name for the realm that is rendered as HTML on the screen when logging in to the admin console.
     * 
     */
    public Output<Optional<String>> displayNameHtml() {
        return Codegen.optional(this.displayNameHtml);
    }
    /**
     * The desired flow for Docker authentication. Defaults to `docker auth`.
     * 
     */
    @Export(name="dockerAuthenticationFlow", type=String.class, parameters={})
    private Output<String> dockerAuthenticationFlow;

    /**
     * @return The desired flow for Docker authentication. Defaults to `docker auth`.
     * 
     */
    public Output<String> dockerAuthenticationFlow() {
        return this.dockerAuthenticationFlow;
    }
    /**
     * When true, multiple users will be allowed to have the same email address. This argument must be set to `false` if `login_with_email_allowed` is set to `true`.
     * 
     */
    @Export(name="duplicateEmailsAllowed", type=Boolean.class, parameters={})
    private Output<Boolean> duplicateEmailsAllowed;

    /**
     * @return When true, multiple users will be allowed to have the same email address. This argument must be set to `false` if `login_with_email_allowed` is set to `true`.
     * 
     */
    public Output<Boolean> duplicateEmailsAllowed() {
        return this.duplicateEmailsAllowed;
    }
    /**
     * When true, the username field is editable.
     * 
     */
    @Export(name="editUsernameAllowed", type=Boolean.class, parameters={})
    private Output<Boolean> editUsernameAllowed;

    /**
     * @return When true, the username field is editable.
     * 
     */
    public Output<Boolean> editUsernameAllowed() {
        return this.editUsernameAllowed;
    }
    /**
     * Used for emails that are sent by Keycloak.
     * 
     */
    @Export(name="emailTheme", type=String.class, parameters={})
    private Output</* @Nullable */ String> emailTheme;

    /**
     * @return Used for emails that are sent by Keycloak.
     * 
     */
    public Output<Optional<String>> emailTheme() {
        return Codegen.optional(this.emailTheme);
    }
    /**
     * When `false`, users and clients will not be able to access this realm. Defaults to `true`.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return When `false`, users and clients will not be able to access this realm. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    @Export(name="internalId", type=String.class, parameters={})
    private Output<String> internalId;

    public Output<String> internalId() {
        return this.internalId;
    }
    @Export(name="internationalization", type=RealmInternationalization.class, parameters={})
    private Output</* @Nullable */ RealmInternationalization> internationalization;

    public Output<Optional<RealmInternationalization>> internationalization() {
        return Codegen.optional(this.internationalization);
    }
    /**
     * Used for the login, forgot password, and registration pages.
     * 
     */
    @Export(name="loginTheme", type=String.class, parameters={})
    private Output</* @Nullable */ String> loginTheme;

    /**
     * @return Used for the login, forgot password, and registration pages.
     * 
     */
    public Output<Optional<String>> loginTheme() {
        return Codegen.optional(this.loginTheme);
    }
    /**
     * When true, users may log in with their email address.
     * 
     */
    @Export(name="loginWithEmailAllowed", type=Boolean.class, parameters={})
    private Output<Boolean> loginWithEmailAllowed;

    /**
     * @return When true, users may log in with their email address.
     * 
     */
    public Output<Boolean> loginWithEmailAllowed() {
        return this.loginWithEmailAllowed;
    }
    /**
     * The maximum amount of time a client has to finish the device code flow before it expires.
     * 
     */
    @Export(name="oauth2DeviceCodeLifespan", type=String.class, parameters={})
    private Output<String> oauth2DeviceCodeLifespan;

    /**
     * @return The maximum amount of time a client has to finish the device code flow before it expires.
     * 
     */
    public Output<String> oauth2DeviceCodeLifespan() {
        return this.oauth2DeviceCodeLifespan;
    }
    /**
     * The minimum amount of time in seconds that the client should wait between polling requests to the token endpoint.
     * 
     */
    @Export(name="oauth2DevicePollingInterval", type=Integer.class, parameters={})
    private Output<Integer> oauth2DevicePollingInterval;

    /**
     * @return The minimum amount of time in seconds that the client should wait between polling requests to the token endpoint.
     * 
     */
    public Output<Integer> oauth2DevicePollingInterval() {
        return this.oauth2DevicePollingInterval;
    }
    /**
     * The amount of time an offline session can be idle before it expires.
     * 
     */
    @Export(name="offlineSessionIdleTimeout", type=String.class, parameters={})
    private Output<String> offlineSessionIdleTimeout;

    /**
     * @return The amount of time an offline session can be idle before it expires.
     * 
     */
    public Output<String> offlineSessionIdleTimeout() {
        return this.offlineSessionIdleTimeout;
    }
    /**
     * The maximum amount of time before an offline session expires regardless of activity.
     * 
     */
    @Export(name="offlineSessionMaxLifespan", type=String.class, parameters={})
    private Output<String> offlineSessionMaxLifespan;

    /**
     * @return The maximum amount of time before an offline session expires regardless of activity.
     * 
     */
    public Output<String> offlineSessionMaxLifespan() {
        return this.offlineSessionMaxLifespan;
    }
    /**
     * Enable `offline_session_max_lifespan`.
     * 
     */
    @Export(name="offlineSessionMaxLifespanEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> offlineSessionMaxLifespanEnabled;

    /**
     * @return Enable `offline_session_max_lifespan`.
     * 
     */
    public Output<Optional<Boolean>> offlineSessionMaxLifespanEnabled() {
        return Codegen.optional(this.offlineSessionMaxLifespanEnabled);
    }
    @Export(name="otpPolicy", type=RealmOtpPolicy.class, parameters={})
    private Output<RealmOtpPolicy> otpPolicy;

    public Output<RealmOtpPolicy> otpPolicy() {
        return this.otpPolicy;
    }
    /**
     * The password policy for users within the realm.
     * 
     */
    @Export(name="passwordPolicy", type=String.class, parameters={})
    private Output</* @Nullable */ String> passwordPolicy;

    /**
     * @return The password policy for users within the realm.
     * 
     */
    public Output<Optional<String>> passwordPolicy() {
        return Codegen.optional(this.passwordPolicy);
    }
    /**
     * The name of the realm. This is unique across Keycloak. This will also be used as the realm&#39;s internal ID within Keycloak.
     * 
     */
    @Export(name="realm", type=String.class, parameters={})
    private Output<String> realm;

    /**
     * @return The name of the realm. This is unique across Keycloak. This will also be used as the realm&#39;s internal ID within Keycloak.
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }
    /**
     * Maximum number of times a refresh token can be reused before they are revoked. If unspecified and &#39;revoke_refresh_token&#39; is enabled the default value is 0 and refresh tokens can not be reused.
     * 
     */
    @Export(name="refreshTokenMaxReuse", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> refreshTokenMaxReuse;

    /**
     * @return Maximum number of times a refresh token can be reused before they are revoked. If unspecified and &#39;revoke_refresh_token&#39; is enabled the default value is 0 and refresh tokens can not be reused.
     * 
     */
    public Output<Optional<Integer>> refreshTokenMaxReuse() {
        return Codegen.optional(this.refreshTokenMaxReuse);
    }
    /**
     * When true, user registration will be enabled, and a link for registration will be displayed on the login page.
     * 
     */
    @Export(name="registrationAllowed", type=Boolean.class, parameters={})
    private Output<Boolean> registrationAllowed;

    /**
     * @return When true, user registration will be enabled, and a link for registration will be displayed on the login page.
     * 
     */
    public Output<Boolean> registrationAllowed() {
        return this.registrationAllowed;
    }
    /**
     * When true, the user&#39;s email will be used as their username during registration.
     * 
     */
    @Export(name="registrationEmailAsUsername", type=Boolean.class, parameters={})
    private Output<Boolean> registrationEmailAsUsername;

    /**
     * @return When true, the user&#39;s email will be used as their username during registration.
     * 
     */
    public Output<Boolean> registrationEmailAsUsername() {
        return this.registrationEmailAsUsername;
    }
    /**
     * The desired flow for user registration. Defaults to `registration`.
     * 
     */
    @Export(name="registrationFlow", type=String.class, parameters={})
    private Output<String> registrationFlow;

    /**
     * @return The desired flow for user registration. Defaults to `registration`.
     * 
     */
    public Output<String> registrationFlow() {
        return this.registrationFlow;
    }
    /**
     * When true, a &#34;remember me&#34; checkbox will be displayed on the login page, and the user&#39;s session will not expire between browser restarts.
     * 
     */
    @Export(name="rememberMe", type=Boolean.class, parameters={})
    private Output<Boolean> rememberMe;

    /**
     * @return When true, a &#34;remember me&#34; checkbox will be displayed on the login page, and the user&#39;s session will not expire between browser restarts.
     * 
     */
    public Output<Boolean> rememberMe() {
        return this.rememberMe;
    }
    /**
     * The desired flow to use when a user attempts to reset their credentials. Defaults to `reset credentials`.
     * 
     */
    @Export(name="resetCredentialsFlow", type=String.class, parameters={})
    private Output<String> resetCredentialsFlow;

    /**
     * @return The desired flow to use when a user attempts to reset their credentials. Defaults to `reset credentials`.
     * 
     */
    public Output<String> resetCredentialsFlow() {
        return this.resetCredentialsFlow;
    }
    /**
     * When true, a &#34;forgot password&#34; link will be displayed on the login page.
     * 
     */
    @Export(name="resetPasswordAllowed", type=Boolean.class, parameters={})
    private Output<Boolean> resetPasswordAllowed;

    /**
     * @return When true, a &#34;forgot password&#34; link will be displayed on the login page.
     * 
     */
    public Output<Boolean> resetPasswordAllowed() {
        return this.resetPasswordAllowed;
    }
    /**
     * If enabled a refresh token can only be used number of times specified in &#39;refresh_token_max_reuse&#39; before they are revoked. If unspecified, refresh tokens can be reused.
     * 
     */
    @Export(name="revokeRefreshToken", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> revokeRefreshToken;

    /**
     * @return If enabled a refresh token can only be used number of times specified in &#39;refresh_token_max_reuse&#39; before they are revoked. If unspecified, refresh tokens can be reused.
     * 
     */
    public Output<Optional<Boolean>> revokeRefreshToken() {
        return Codegen.optional(this.revokeRefreshToken);
    }
    @Export(name="securityDefenses", type=RealmSecurityDefenses.class, parameters={})
    private Output</* @Nullable */ RealmSecurityDefenses> securityDefenses;

    public Output<Optional<RealmSecurityDefenses>> securityDefenses() {
        return Codegen.optional(this.securityDefenses);
    }
    @Export(name="smtpServer", type=RealmSmtpServer.class, parameters={})
    private Output</* @Nullable */ RealmSmtpServer> smtpServer;

    public Output<Optional<RealmSmtpServer>> smtpServer() {
        return Codegen.optional(this.smtpServer);
    }
    /**
     * Can be one of following values: &#39;none, &#39;external&#39; or &#39;all&#39;
     * 
     */
    @Export(name="sslRequired", type=String.class, parameters={})
    private Output</* @Nullable */ String> sslRequired;

    /**
     * @return Can be one of following values: &#39;none, &#39;external&#39; or &#39;all&#39;
     * 
     */
    public Output<Optional<String>> sslRequired() {
        return Codegen.optional(this.sslRequired);
    }
    /**
     * The amount of time a session can be idle before it expires.
     * 
     */
    @Export(name="ssoSessionIdleTimeout", type=String.class, parameters={})
    private Output<String> ssoSessionIdleTimeout;

    /**
     * @return The amount of time a session can be idle before it expires.
     * 
     */
    public Output<String> ssoSessionIdleTimeout() {
        return this.ssoSessionIdleTimeout;
    }
    @Export(name="ssoSessionIdleTimeoutRememberMe", type=String.class, parameters={})
    private Output<String> ssoSessionIdleTimeoutRememberMe;

    public Output<String> ssoSessionIdleTimeoutRememberMe() {
        return this.ssoSessionIdleTimeoutRememberMe;
    }
    /**
     * The maximum amount of time before a session expires regardless of activity.
     * 
     */
    @Export(name="ssoSessionMaxLifespan", type=String.class, parameters={})
    private Output<String> ssoSessionMaxLifespan;

    /**
     * @return The maximum amount of time before a session expires regardless of activity.
     * 
     */
    public Output<String> ssoSessionMaxLifespan() {
        return this.ssoSessionMaxLifespan;
    }
    @Export(name="ssoSessionMaxLifespanRememberMe", type=String.class, parameters={})
    private Output<String> ssoSessionMaxLifespanRememberMe;

    public Output<String> ssoSessionMaxLifespanRememberMe() {
        return this.ssoSessionMaxLifespanRememberMe;
    }
    /**
     * When `true`, users are allowed to manage their own resources. Defaults to `false`.
     * 
     */
    @Export(name="userManagedAccess", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> userManagedAccess;

    /**
     * @return When `true`, users are allowed to manage their own resources. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> userManagedAccess() {
        return Codegen.optional(this.userManagedAccess);
    }
    /**
     * When true, users are required to verify their email address after registration and after email address changes.
     * 
     */
    @Export(name="verifyEmail", type=Boolean.class, parameters={})
    private Output<Boolean> verifyEmail;

    /**
     * @return When true, users are required to verify their email address after registration and after email address changes.
     * 
     */
    public Output<Boolean> verifyEmail() {
        return this.verifyEmail;
    }
    /**
     * Configuration for WebAuthn Passwordless Policy authentication.
     * 
     */
    @Export(name="webAuthnPasswordlessPolicy", type=RealmWebAuthnPasswordlessPolicy.class, parameters={})
    private Output<RealmWebAuthnPasswordlessPolicy> webAuthnPasswordlessPolicy;

    /**
     * @return Configuration for WebAuthn Passwordless Policy authentication.
     * 
     */
    public Output<RealmWebAuthnPasswordlessPolicy> webAuthnPasswordlessPolicy() {
        return this.webAuthnPasswordlessPolicy;
    }
    /**
     * Configuration for WebAuthn Policy authentication.
     * 
     */
    @Export(name="webAuthnPolicy", type=RealmWebAuthnPolicy.class, parameters={})
    private Output<RealmWebAuthnPolicy> webAuthnPolicy;

    /**
     * @return Configuration for WebAuthn Policy authentication.
     * 
     */
    public Output<RealmWebAuthnPolicy> webAuthnPolicy() {
        return this.webAuthnPolicy;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Realm(String name) {
        this(name, RealmArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Realm(String name, RealmArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Realm(String name, RealmArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realm:Realm", name, args == null ? RealmArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Realm(String name, Output<String> id, @Nullable RealmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realm:Realm", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Realm get(String name, Output<String> id, @Nullable RealmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Realm(name, id, state, options);
    }
}
