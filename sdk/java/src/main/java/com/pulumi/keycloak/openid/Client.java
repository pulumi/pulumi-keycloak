// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.ClientArgs;
import com.pulumi.keycloak.openid.inputs.ClientState;
import com.pulumi.keycloak.openid.outputs.ClientAuthenticationFlowBindingOverrides;
import com.pulumi.keycloak.openid.outputs.ClientAuthorization;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # keycloak.openid.Client
 * 
 * Allows for creating and managing Keycloak clients that use the OpenID Connect protocol.
 * 
 * Clients are entities that can use Keycloak for user authentication. Typically,
 * clients are applications that redirect users to Keycloak for authentication
 * in order to take advantage of Keycloak&#39;s user sessions for SSO.
 * 
 * ### Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.openid.Client;
 * import com.pulumi.keycloak.openid.ClientArgs;
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
 *         var realm = new Realm("realm", RealmArgs.builder()
 *             .realm("my-realm")
 *             .enabled(true)
 *             .build());
 * 
 *         var openidClient = new Client("openidClient", ClientArgs.builder()
 *             .realmId(realm.id())
 *             .clientId("test-client")
 *             .name("test client")
 *             .enabled(true)
 *             .accessType("CONFIDENTIAL")
 *             .validRedirectUris("http://localhost:8080/openid-callback")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realm_id` - (Required) The realm this client is attached to.
 * - `client_id` - (Required) The unique ID of this client, referenced in the URI during authentication and in issued tokens.
 * - `name` - (Optional) The display name of this client in the GUI.
 * - `enabled` - (Optional) When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
 * - `description` - (Optional) The description of this client in the GUI.
 * - `access_type` - (Required) Specifies the type of client, which can be one of the following:
 *     - `CONFIDENTIAL` - Used for server-side clients that require both client ID and secret when authenticating.
 *       This client should be used for applications using the Authorization Code or Client Credentials grant flows.
 *     - `PUBLIC` - Used for browser-only applications that do not require a client secret, and instead rely only on authorized redirect
 *       URIs for security. This client should be used for applications using the Implicit grant flow.
 *     - `BEARER-ONLY` - Used for services that never initiate a login. This client will only allow bearer token requests.
 * - `client_secret` - (Optional) The secret for clients with an `access_type` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and
 *   should be treated with the same care as a password. If omitted, Keycloak will generate a GUID for this attribute.
 * - `standard_flow_enabled` - (Optional) When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
 * - `implicit_flow_enabled` - (Optional) When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
 * - `direct_access_grants_enabled` - (Optional) When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
 * - `service_accounts_enabled` - (Optional) When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
 * - `valid_redirect_uris` - (Optional) A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
 *   wildcards in the form of an asterisk can be used here. This attribute must be set if either `standard_flow_enabled` or `implicit_flow_enabled`
 *   is set to `true`.
 * - `web_origins` - (Optional) A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
 * - `admin_url` - (Optional) URL to the admin interface of the client.
 * - `base_url` - (Optional) Default URL to use when the auth server needs to redirect or link back to the client.
 * - `pkce_code_challenge_method` - (Optional) The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
 * - `full_scope_allowed` - (Optional) - Allow to include all roles mappings in the access token.
 * 
 * ### Attributes Reference
 * 
 * In addition to the arguments listed above, the following computed attributes are exported:
 * 
 * - `service_account_user_id` - When service accounts are enabled for this client, this attribute is the unique ID for the Keycloak user that represents this service account.
 * 
 * ### Import
 * 
 * Clients can be imported using the format `{{realm_id}}/{{client_keycloak_id}}`, where `client_keycloak_id` is the unique ID that Keycloak
 * assigns to the client upon creation. This value can be found in the URI when editing this client in the GUI, and is typically a GUID.
 * 
 * Example:
 * 
 */
@ResourceType(type="keycloak:openid/client:Client")
public class Client extends com.pulumi.resources.CustomResource {
    @Export(name="accessTokenLifespan", refs={String.class}, tree="[0]")
    private Output<String> accessTokenLifespan;

    public Output<String> accessTokenLifespan() {
        return this.accessTokenLifespan;
    }
    @Export(name="accessType", refs={String.class}, tree="[0]")
    private Output<String> accessType;

    public Output<String> accessType() {
        return this.accessType;
    }
    @Export(name="adminUrl", refs={String.class}, tree="[0]")
    private Output<String> adminUrl;

    public Output<String> adminUrl() {
        return this.adminUrl;
    }
    @Export(name="authenticationFlowBindingOverrides", refs={ClientAuthenticationFlowBindingOverrides.class}, tree="[0]")
    private Output</* @Nullable */ ClientAuthenticationFlowBindingOverrides> authenticationFlowBindingOverrides;

    public Output<Optional<ClientAuthenticationFlowBindingOverrides>> authenticationFlowBindingOverrides() {
        return Codegen.optional(this.authenticationFlowBindingOverrides);
    }
    @Export(name="authorization", refs={ClientAuthorization.class}, tree="[0]")
    private Output</* @Nullable */ ClientAuthorization> authorization;

    public Output<Optional<ClientAuthorization>> authorization() {
        return Codegen.optional(this.authorization);
    }
    @Export(name="backchannelLogoutRevokeOfflineSessions", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> backchannelLogoutRevokeOfflineSessions;

    public Output<Optional<Boolean>> backchannelLogoutRevokeOfflineSessions() {
        return Codegen.optional(this.backchannelLogoutRevokeOfflineSessions);
    }
    @Export(name="backchannelLogoutSessionRequired", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> backchannelLogoutSessionRequired;

    public Output<Optional<Boolean>> backchannelLogoutSessionRequired() {
        return Codegen.optional(this.backchannelLogoutSessionRequired);
    }
    @Export(name="backchannelLogoutUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backchannelLogoutUrl;

    public Output<Optional<String>> backchannelLogoutUrl() {
        return Codegen.optional(this.backchannelLogoutUrl);
    }
    @Export(name="baseUrl", refs={String.class}, tree="[0]")
    private Output<String> baseUrl;

    public Output<String> baseUrl() {
        return this.baseUrl;
    }
    @Export(name="clientAuthenticatorType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientAuthenticatorType;

    public Output<Optional<String>> clientAuthenticatorType() {
        return Codegen.optional(this.clientAuthenticatorType);
    }
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output<String> clientId;

    public Output<String> clientId() {
        return this.clientId;
    }
    @Export(name="clientOfflineSessionIdleTimeout", refs={String.class}, tree="[0]")
    private Output<String> clientOfflineSessionIdleTimeout;

    public Output<String> clientOfflineSessionIdleTimeout() {
        return this.clientOfflineSessionIdleTimeout;
    }
    @Export(name="clientOfflineSessionMaxLifespan", refs={String.class}, tree="[0]")
    private Output<String> clientOfflineSessionMaxLifespan;

    public Output<String> clientOfflineSessionMaxLifespan() {
        return this.clientOfflineSessionMaxLifespan;
    }
    @Export(name="clientSecret", refs={String.class}, tree="[0]")
    private Output<String> clientSecret;

    public Output<String> clientSecret() {
        return this.clientSecret;
    }
    @Export(name="clientSessionIdleTimeout", refs={String.class}, tree="[0]")
    private Output<String> clientSessionIdleTimeout;

    public Output<String> clientSessionIdleTimeout() {
        return this.clientSessionIdleTimeout;
    }
    @Export(name="clientSessionMaxLifespan", refs={String.class}, tree="[0]")
    private Output<String> clientSessionMaxLifespan;

    public Output<String> clientSessionMaxLifespan() {
        return this.clientSessionMaxLifespan;
    }
    @Export(name="consentRequired", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> consentRequired;

    public Output<Boolean> consentRequired() {
        return this.consentRequired;
    }
    @Export(name="consentScreenText", refs={String.class}, tree="[0]")
    private Output<String> consentScreenText;

    public Output<String> consentScreenText() {
        return this.consentScreenText;
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    public Output<String> description() {
        return this.description;
    }
    @Export(name="directAccessGrantsEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> directAccessGrantsEnabled;

    public Output<Boolean> directAccessGrantsEnabled() {
        return this.directAccessGrantsEnabled;
    }
    @Export(name="displayOnConsentScreen", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> displayOnConsentScreen;

    public Output<Boolean> displayOnConsentScreen() {
        return this.displayOnConsentScreen;
    }
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    @Export(name="excludeSessionStateFromAuthResponse", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> excludeSessionStateFromAuthResponse;

    public Output<Boolean> excludeSessionStateFromAuthResponse() {
        return this.excludeSessionStateFromAuthResponse;
    }
    @Export(name="extraConfig", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> extraConfig;

    public Output<Optional<Map<String,Object>>> extraConfig() {
        return Codegen.optional(this.extraConfig);
    }
    @Export(name="frontchannelLogoutEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> frontchannelLogoutEnabled;

    public Output<Boolean> frontchannelLogoutEnabled() {
        return this.frontchannelLogoutEnabled;
    }
    @Export(name="frontchannelLogoutUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> frontchannelLogoutUrl;

    public Output<Optional<String>> frontchannelLogoutUrl() {
        return Codegen.optional(this.frontchannelLogoutUrl);
    }
    @Export(name="fullScopeAllowed", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> fullScopeAllowed;

    public Output<Optional<Boolean>> fullScopeAllowed() {
        return Codegen.optional(this.fullScopeAllowed);
    }
    @Export(name="implicitFlowEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> implicitFlowEnabled;

    public Output<Boolean> implicitFlowEnabled() {
        return this.implicitFlowEnabled;
    }
    @Export(name="import", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> import_;

    public Output<Optional<Boolean>> import_() {
        return Codegen.optional(this.import_);
    }
    @Export(name="loginTheme", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loginTheme;

    public Output<Optional<String>> loginTheme() {
        return Codegen.optional(this.loginTheme);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="oauth2DeviceAuthorizationGrantEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> oauth2DeviceAuthorizationGrantEnabled;

    public Output<Optional<Boolean>> oauth2DeviceAuthorizationGrantEnabled() {
        return Codegen.optional(this.oauth2DeviceAuthorizationGrantEnabled);
    }
    @Export(name="oauth2DeviceCodeLifespan", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> oauth2DeviceCodeLifespan;

    public Output<Optional<String>> oauth2DeviceCodeLifespan() {
        return Codegen.optional(this.oauth2DeviceCodeLifespan);
    }
    @Export(name="oauth2DevicePollingInterval", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> oauth2DevicePollingInterval;

    public Output<Optional<String>> oauth2DevicePollingInterval() {
        return Codegen.optional(this.oauth2DevicePollingInterval);
    }
    @Export(name="pkceCodeChallengeMethod", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pkceCodeChallengeMethod;

    public Output<Optional<String>> pkceCodeChallengeMethod() {
        return Codegen.optional(this.pkceCodeChallengeMethod);
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="resourceServerId", refs={String.class}, tree="[0]")
    private Output<String> resourceServerId;

    public Output<String> resourceServerId() {
        return this.resourceServerId;
    }
    @Export(name="rootUrl", refs={String.class}, tree="[0]")
    private Output<String> rootUrl;

    public Output<String> rootUrl() {
        return this.rootUrl;
    }
    @Export(name="serviceAccountUserId", refs={String.class}, tree="[0]")
    private Output<String> serviceAccountUserId;

    public Output<String> serviceAccountUserId() {
        return this.serviceAccountUserId;
    }
    @Export(name="serviceAccountsEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> serviceAccountsEnabled;

    public Output<Boolean> serviceAccountsEnabled() {
        return this.serviceAccountsEnabled;
    }
    @Export(name="standardFlowEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> standardFlowEnabled;

    public Output<Boolean> standardFlowEnabled() {
        return this.standardFlowEnabled;
    }
    @Export(name="useRefreshTokens", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useRefreshTokens;

    public Output<Optional<Boolean>> useRefreshTokens() {
        return Codegen.optional(this.useRefreshTokens);
    }
    @Export(name="useRefreshTokensClientCredentials", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useRefreshTokensClientCredentials;

    public Output<Optional<Boolean>> useRefreshTokensClientCredentials() {
        return Codegen.optional(this.useRefreshTokensClientCredentials);
    }
    @Export(name="validPostLogoutRedirectUris", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> validPostLogoutRedirectUris;

    public Output<List<String>> validPostLogoutRedirectUris() {
        return this.validPostLogoutRedirectUris;
    }
    @Export(name="validRedirectUris", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> validRedirectUris;

    public Output<List<String>> validRedirectUris() {
        return this.validRedirectUris;
    }
    @Export(name="webOrigins", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> webOrigins;

    public Output<List<String>> webOrigins() {
        return this.webOrigins;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Client(java.lang.String name) {
        this(name, ClientArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Client(java.lang.String name, ClientArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Client(java.lang.String name, ClientArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/client:Client", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Client(java.lang.String name, Output<java.lang.String> id, @Nullable ClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/client:Client", name, state, makeResourceOptions(options, id), false);
    }

    private static ClientArgs makeArgs(ClientArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ClientArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "clientSecret"
            ))
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
    public static Client get(java.lang.String name, Output<java.lang.String> id, @Nullable ClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Client(name, id, state, options);
    }
}
