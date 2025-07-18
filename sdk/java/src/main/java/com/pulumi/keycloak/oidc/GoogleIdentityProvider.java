// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.oidc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.oidc.GoogleIdentityProviderArgs;
import com.pulumi.keycloak.oidc.inputs.GoogleIdentityProviderState;
import java.lang.Boolean;
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
 * import com.pulumi.keycloak.oidc.GoogleIdentityProvider;
 * import com.pulumi.keycloak.oidc.GoogleIdentityProviderArgs;
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
 *         var google = new GoogleIdentityProvider("google", GoogleIdentityProviderArgs.builder()
 *             .realm(realm.id())
 *             .clientId(googleIdentityProviderClientId)
 *             .clientSecret(googleIdentityProviderClientSecret)
 *             .trustEmail(true)
 *             .hostedDomain("example.com")
 *             .syncMode("IMPORT")
 *             .extraConfig(Map.of("myCustomConfigKey", "myValue"))
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
 * Google Identity providers can be imported using the format {{realm_id}}/{{idp_alias}}, where idp_alias is the identity provider alias.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider google_identity_provider my-realm/my-google-idp
 * ```
 * 
 */
@ResourceType(type="keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider")
public class GoogleIdentityProvider extends com.pulumi.resources.CustomResource {
    /**
     * When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
     * 
     */
    @Export(name="acceptsPromptNoneForwardFromClient", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> acceptsPromptNoneForwardFromClient;

    /**
     * @return When `true`, unauthenticated requests with `prompt=none` will be forwarded to Google instead of returning an error. Defaults to `false`.
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
     * The alias for the Google identity provider.
     * 
     */
    @Export(name="alias", refs={String.class}, tree="[0]")
    private Output<String> alias;

    /**
     * @return The alias for the Google identity provider.
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
     * The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
     * 
     */
    @Export(name="defaultScopes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultScopes;

    /**
     * @return The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid profile email`.
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
     * Display name for the Google identity provider in the GUI.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Display name for the Google identity provider in the GUI.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
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
    @Export(name="extraConfig", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> extraConfig;

    public Output<Optional<Map<String,String>>> extraConfig() {
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
     * When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
     * 
     */
    @Export(name="hideOnLoginPage", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hideOnLoginPage;

    /**
     * @return When `true`, this identity provider will be hidden on the login page. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> hideOnLoginPage() {
        return Codegen.optional(this.hideOnLoginPage);
    }
    /**
     * Sets the &#34;hd&#34; query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
     * 
     */
    @Export(name="hostedDomain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hostedDomain;

    /**
     * @return Sets the &#34;hd&#34; query parameter when logging in with Google. Google will only list accounts for this domain. Keycloak will validate that the returned identity token has a claim for this domain. When `*` is entered, an account from any domain can be used.
     * 
     */
    public Output<Optional<String>> hostedDomain() {
        return Codegen.optional(this.hostedDomain);
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
     * When `true`, users cannot sign-in using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     * 
     */
    @Export(name="linkOnly", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> linkOnly;

    /**
     * @return When `true`, users cannot sign-in using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> linkOnly() {
        return Codegen.optional(this.linkOnly);
    }
    @Export(name="orgDomain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> orgDomain;

    public Output<Optional<String>> orgDomain() {
        return Codegen.optional(this.orgDomain);
    }
    @Export(name="orgRedirectModeEmailMatches", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> orgRedirectModeEmailMatches;

    public Output<Optional<Boolean>> orgRedirectModeEmailMatches() {
        return Codegen.optional(this.orgRedirectModeEmailMatches);
    }
    /**
     * ID of organization with which this identity is linked.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> organizationId;

    /**
     * @return ID of organization with which this identity is linked.
     * 
     */
    public Output<Optional<String>> organizationId() {
        return Codegen.optional(this.organizationId);
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
     * The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
     * 
     */
    @Export(name="providerId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> providerId;

    /**
     * @return The ID of the identity provider to use. Defaults to `google`, which should be used unless you have extended Keycloak and provided your own implementation.
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
     * Sets the &#34;access_type&#34; query parameter to &#34;offline&#34; when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
     * 
     */
    @Export(name="requestRefreshToken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> requestRefreshToken;

    /**
     * @return Sets the &#34;access_type&#34; query parameter to &#34;offline&#34; when redirecting to google authorization endpoint,to get a refresh token back. This is useful for using Token Exchange to retrieve a Google token to access Google APIs when the user is offline.
     * 
     */
    public Output<Optional<Boolean>> requestRefreshToken() {
        return Codegen.optional(this.requestRefreshToken);
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
     * Sets the &#34;userIp&#34; query parameter when querying Google&#39;s User Info service. This will use the user&#39;s IP address. This is useful if Google is throttling Keycloak&#39;s access to the User Info service.
     * 
     */
    @Export(name="useUserIpParam", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useUserIpParam;

    /**
     * @return Sets the &#34;userIp&#34; query parameter when querying Google&#39;s User Info service. This will use the user&#39;s IP address. This is useful if Google is throttling Keycloak&#39;s access to the User Info service.
     * 
     */
    public Output<Optional<Boolean>> useUserIpParam() {
        return Codegen.optional(this.useUserIpParam);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GoogleIdentityProvider(java.lang.String name) {
        this(name, GoogleIdentityProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GoogleIdentityProvider(java.lang.String name, GoogleIdentityProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GoogleIdentityProvider(java.lang.String name, GoogleIdentityProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GoogleIdentityProvider(java.lang.String name, Output<java.lang.String> id, @Nullable GoogleIdentityProviderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:oidc/googleIdentityProvider:GoogleIdentityProvider", name, state, makeResourceOptions(options, id), false);
    }

    private static GoogleIdentityProviderArgs makeArgs(GoogleIdentityProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GoogleIdentityProviderArgs.Empty : args;
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
    public static GoogleIdentityProvider get(java.lang.String name, Output<java.lang.String> id, @Nullable GoogleIdentityProviderState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GoogleIdentityProvider(name, id, state, options);
    }
}
