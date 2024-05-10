// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.oidc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.oidc.IdentityProviderArgs;
import com.pulumi.keycloak.oidc.inputs.IdentityProviderState;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing OIDC Identity Providers within Keycloak.
 * 
 * OIDC (OpenID Connect) identity providers allows users to authenticate through a third party system using the OIDC standard.
 * 
 * ## Example Usage
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
 * import com.pulumi.keycloak.oidc.IdentityProvider;
 * import com.pulumi.keycloak.oidc.IdentityProviderArgs;
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
 *         var realmIdentityProvider = new IdentityProvider("realmIdentityProvider", IdentityProviderArgs.builder()        
 *             .realm(realm.id())
 *             .alias("my-idp")
 *             .authorizationUrl("https://authorizationurl.com")
 *             .clientId("clientID")
 *             .clientSecret("clientSecret")
 *             .tokenUrl("https://tokenurl.com")
 *             .extraConfig(Map.of("clientAuthMethod", "client_secret_post"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Identity providers can be imported using the format `{{realm_id}}/{{idp_alias}}`, where `idp_alias` is the identity provider alias.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:oidc/identityProvider:IdentityProvider realm_identity_provider my-realm/my-idp
 * ```
 * 
 */
@ResourceType(type="keycloak:oidc/identityProvider:IdentityProvider")
public class IdentityProvider extends com.pulumi.resources.CustomResource {
    /**
     * When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
     * 
     */
    @Export(name="acceptsPromptNoneForwardFromClient", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> acceptsPromptNoneForwardFromClient;

    /**
     * @return When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> acceptsPromptNoneForwardFromClient() {
        return Codegen.optional(this.acceptsPromptNoneForwardFromClient);
    }
    /**
     * When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     * 
     */
    @Export(name="addReadTokenRoleOnCreate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> addReadTokenRoleOnCreate;

    /**
     * @return When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> addReadTokenRoleOnCreate() {
        return Codegen.optional(this.addReadTokenRoleOnCreate);
    }
    /**
     * The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
     * 
     */
    @Export(name="alias", refs={String.class}, tree="[0]")
    private Output<String> alias;

    /**
     * @return The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }
    /**
     * Enable/disable authenticate users by default.
     * 
     */
    @Export(name="authenticateByDefault", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> authenticateByDefault;

    /**
     * @return Enable/disable authenticate users by default.
     * 
     */
    public Output<Optional<Boolean>> authenticateByDefault() {
        return Codegen.optional(this.authenticateByDefault);
    }
    /**
     * The Authorization Url.
     * 
     */
    @Export(name="authorizationUrl", refs={String.class}, tree="[0]")
    private Output<String> authorizationUrl;

    /**
     * @return The Authorization Url.
     * 
     */
    public Output<String> authorizationUrl() {
        return this.authorizationUrl;
    }
    /**
     * Does the external IDP support backchannel logout? Defaults to `true`.
     * 
     */
    @Export(name="backchannelSupported", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> backchannelSupported;

    /**
     * @return Does the external IDP support backchannel logout? Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> backchannelSupported() {
        return Codegen.optional(this.backchannelSupported);
    }
    /**
     * The client or client identifier registered within the identity provider.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output<String> clientId;

    /**
     * @return The client or client identifier registered within the identity provider.
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
     * 
     */
    @Export(name="clientSecret", refs={String.class}, tree="[0]")
    private Output<String> clientSecret;

    /**
     * @return The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
     * 
     */
    public Output<String> clientSecret() {
        return this.clientSecret;
    }
    /**
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
     * 
     */
    @Export(name="defaultScopes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultScopes;

    /**
     * @return The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
     * 
     */
    public Output<Optional<String>> defaultScopes() {
        return Codegen.optional(this.defaultScopes);
    }
    /**
     * When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     * 
     */
    @Export(name="disableUserInfo", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableUserInfo;

    /**
     * @return When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> disableUserInfo() {
        return Codegen.optional(this.disableUserInfo);
    }
    /**
     * Display name for the identity provider in the GUI.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Display name for the identity provider in the GUI.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    @Export(name="extraConfig", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> extraConfig;

    public Output<Optional<Map<String,Object>>> extraConfig() {
        return Codegen.optional(this.extraConfig);
    }
    /**
     * The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     * 
     */
    @Export(name="firstBrokerLoginFlowAlias", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> firstBrokerLoginFlowAlias;

    /**
     * @return The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     * 
     */
    public Output<Optional<String>> firstBrokerLoginFlowAlias() {
        return Codegen.optional(this.firstBrokerLoginFlowAlias);
    }
    /**
     * A number defining the order of this identity provider in the GUI.
     * 
     */
    @Export(name="guiOrder", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> guiOrder;

    /**
     * @return A number defining the order of this identity provider in the GUI.
     * 
     */
    public Output<Optional<String>> guiOrder() {
        return Codegen.optional(this.guiOrder);
    }
    /**
     * When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
     * 
     */
    @Export(name="hideOnLoginPage", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hideOnLoginPage;

    /**
     * @return When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> hideOnLoginPage() {
        return Codegen.optional(this.hideOnLoginPage);
    }
    /**
     * (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
     * 
     */
    @Export(name="internalId", refs={String.class}, tree="[0]")
    private Output<String> internalId;

    /**
     * @return (Computed) The unique ID that Keycloak assigns to the identity provider upon creation.
     * 
     */
    public Output<String> internalId() {
        return this.internalId;
    }
    /**
     * The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
     * 
     */
    @Export(name="issuer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> issuer;

    /**
     * @return The issuer identifier for the issuer of the response. If not provided, no validation will be performed.
     * 
     */
    public Output<Optional<String>> issuer() {
        return Codegen.optional(this.issuer);
    }
    /**
     * JSON Web Key Set URL.
     * 
     */
    @Export(name="jwksUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> jwksUrl;

    /**
     * @return JSON Web Key Set URL.
     * 
     */
    public Output<Optional<String>> jwksUrl() {
        return Codegen.optional(this.jwksUrl);
    }
    /**
     * When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     * 
     */
    @Export(name="linkOnly", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> linkOnly;

    /**
     * @return When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> linkOnly() {
        return Codegen.optional(this.linkOnly);
    }
    /**
     * Pass login hint to identity provider.
     * 
     */
    @Export(name="loginHint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loginHint;

    /**
     * @return Pass login hint to identity provider.
     * 
     */
    public Output<Optional<String>> loginHint() {
        return Codegen.optional(this.loginHint);
    }
    /**
     * The Logout URL is the end session endpoint to use to logout user from external identity provider.
     * 
     */
    @Export(name="logoutUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> logoutUrl;

    /**
     * @return The Logout URL is the end session endpoint to use to logout user from external identity provider.
     * 
     */
    public Output<Optional<String>> logoutUrl() {
        return Codegen.optional(this.logoutUrl);
    }
    /**
     * The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     * 
     */
    @Export(name="postBrokerLoginFlowAlias", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> postBrokerLoginFlowAlias;

    /**
     * @return The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     * 
     */
    public Output<Optional<String>> postBrokerLoginFlowAlias() {
        return Codegen.optional(this.postBrokerLoginFlowAlias);
    }
    /**
     * The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
     * 
     */
    @Export(name="providerId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> providerId;

    /**
     * @return The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
     * 
     */
    public Output<Optional<String>> providerId() {
        return Codegen.optional(this.providerId);
    }
    /**
     * The name of the realm. This is unique across Keycloak.
     * 
     */
    @Export(name="realm", refs={String.class}, tree="[0]")
    private Output<String> realm;

    /**
     * @return The name of the realm. This is unique across Keycloak.
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }
    /**
     * When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     * 
     */
    @Export(name="storeToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> storeToken;

    /**
     * @return When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> storeToken() {
        return Codegen.optional(this.storeToken);
    }
    /**
     * The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     * 
     */
    @Export(name="syncMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> syncMode;

    /**
     * @return The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     * 
     */
    public Output<Optional<String>> syncMode() {
        return Codegen.optional(this.syncMode);
    }
    /**
     * The Token URL.
     * 
     */
    @Export(name="tokenUrl", refs={String.class}, tree="[0]")
    private Output<String> tokenUrl;

    /**
     * @return The Token URL.
     * 
     */
    public Output<String> tokenUrl() {
        return this.tokenUrl;
    }
    /**
     * When `true`, email addresses for users in this provider will automatically be verified regardless of the realm&#39;s email verification policy. Defaults to `false`.
     * 
     */
    @Export(name="trustEmail", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> trustEmail;

    /**
     * @return When `true`, email addresses for users in this provider will automatically be verified regardless of the realm&#39;s email verification policy. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> trustEmail() {
        return Codegen.optional(this.trustEmail);
    }
    /**
     * Pass current locale to identity provider. Defaults to `false`.
     * 
     */
    @Export(name="uiLocales", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> uiLocales;

    /**
     * @return Pass current locale to identity provider. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> uiLocales() {
        return Codegen.optional(this.uiLocales);
    }
    /**
     * User Info URL.
     * 
     */
    @Export(name="userInfoUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userInfoUrl;

    /**
     * @return User Info URL.
     * 
     */
    public Output<Optional<String>> userInfoUrl() {
        return Codegen.optional(this.userInfoUrl);
    }
    /**
     * Enable/disable signature validation of external IDP signatures. Defaults to `false`.
     * 
     */
    @Export(name="validateSignature", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> validateSignature;

    /**
     * @return Enable/disable signature validation of external IDP signatures. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> validateSignature() {
        return Codegen.optional(this.validateSignature);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IdentityProvider(String name) {
        this(name, IdentityProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IdentityProvider(String name, IdentityProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IdentityProvider(String name, IdentityProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:oidc/identityProvider:IdentityProvider", name, args == null ? IdentityProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IdentityProvider(String name, Output<String> id, @Nullable IdentityProviderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:oidc/identityProvider:IdentityProvider", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static IdentityProvider get(String name, Output<String> id, @Nullable IdentityProviderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IdentityProvider(name, id, state, options);
    }
}
