// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.oidc;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IdentityProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final IdentityProviderArgs Empty = new IdentityProviderArgs();

    /**
     * When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
     * 
     */
    @Import(name="acceptsPromptNoneForwardFromClient")
    private @Nullable Output<Boolean> acceptsPromptNoneForwardFromClient;

    /**
     * @return When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> acceptsPromptNoneForwardFromClient() {
        return Optional.ofNullable(this.acceptsPromptNoneForwardFromClient);
    }

    /**
     * When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     * 
     */
    @Import(name="addReadTokenRoleOnCreate")
    private @Nullable Output<Boolean> addReadTokenRoleOnCreate;

    /**
     * @return When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> addReadTokenRoleOnCreate() {
        return Optional.ofNullable(this.addReadTokenRoleOnCreate);
    }

    /**
     * The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
     * 
     */
    @Import(name="alias", required=true)
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
    @Import(name="authenticateByDefault")
    private @Nullable Output<Boolean> authenticateByDefault;

    /**
     * @return Enable/disable authenticate users by default.
     * 
     */
    public Optional<Output<Boolean>> authenticateByDefault() {
        return Optional.ofNullable(this.authenticateByDefault);
    }

    /**
     * The Authorization Url.
     * 
     */
    @Import(name="authorizationUrl", required=true)
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
    @Import(name="backchannelSupported")
    private @Nullable Output<Boolean> backchannelSupported;

    /**
     * @return Does the external IDP support backchannel logout? Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> backchannelSupported() {
        return Optional.ofNullable(this.backchannelSupported);
    }

    /**
     * The client or client identifier registered within the identity provider.
     * 
     */
    @Import(name="clientId", required=true)
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
    @Import(name="clientSecret", required=true)
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
    @Import(name="defaultScopes")
    private @Nullable Output<String> defaultScopes;

    /**
     * @return The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
     * 
     */
    public Optional<Output<String>> defaultScopes() {
        return Optional.ofNullable(this.defaultScopes);
    }

    /**
     * When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     * 
     */
    @Import(name="disableUserInfo")
    private @Nullable Output<Boolean> disableUserInfo;

    /**
     * @return When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> disableUserInfo() {
        return Optional.ofNullable(this.disableUserInfo);
    }

    /**
     * Display name for the identity provider in the GUI.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Display name for the identity provider in the GUI.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="extraConfig")
    private @Nullable Output<Map<String,Object>> extraConfig;

    public Optional<Output<Map<String,Object>>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    /**
     * The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     * 
     */
    @Import(name="firstBrokerLoginFlowAlias")
    private @Nullable Output<String> firstBrokerLoginFlowAlias;

    /**
     * @return The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
     * 
     */
    public Optional<Output<String>> firstBrokerLoginFlowAlias() {
        return Optional.ofNullable(this.firstBrokerLoginFlowAlias);
    }

    /**
     * A number defining the order of this identity provider in the GUI.
     * 
     */
    @Import(name="guiOrder")
    private @Nullable Output<String> guiOrder;

    /**
     * @return A number defining the order of this identity provider in the GUI.
     * 
     */
    public Optional<Output<String>> guiOrder() {
        return Optional.ofNullable(this.guiOrder);
    }

    /**
     * When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
     * 
     */
    @Import(name="hideOnLoginPage")
    private @Nullable Output<Boolean> hideOnLoginPage;

    /**
     * @return When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> hideOnLoginPage() {
        return Optional.ofNullable(this.hideOnLoginPage);
    }

    /**
     * JSON Web Key Set URL.
     * 
     */
    @Import(name="jwksUrl")
    private @Nullable Output<String> jwksUrl;

    /**
     * @return JSON Web Key Set URL.
     * 
     */
    public Optional<Output<String>> jwksUrl() {
        return Optional.ofNullable(this.jwksUrl);
    }

    /**
     * When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     * 
     */
    @Import(name="linkOnly")
    private @Nullable Output<Boolean> linkOnly;

    /**
     * @return When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> linkOnly() {
        return Optional.ofNullable(this.linkOnly);
    }

    /**
     * Pass login hint to identity provider.
     * 
     */
    @Import(name="loginHint")
    private @Nullable Output<String> loginHint;

    /**
     * @return Pass login hint to identity provider.
     * 
     */
    public Optional<Output<String>> loginHint() {
        return Optional.ofNullable(this.loginHint);
    }

    /**
     * The Logout URL is the end session endpoint to use to logout user from external identity provider.
     * 
     */
    @Import(name="logoutUrl")
    private @Nullable Output<String> logoutUrl;

    /**
     * @return The Logout URL is the end session endpoint to use to logout user from external identity provider.
     * 
     */
    public Optional<Output<String>> logoutUrl() {
        return Optional.ofNullable(this.logoutUrl);
    }

    /**
     * The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     * 
     */
    @Import(name="postBrokerLoginFlowAlias")
    private @Nullable Output<String> postBrokerLoginFlowAlias;

    /**
     * @return The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
     * 
     */
    public Optional<Output<String>> postBrokerLoginFlowAlias() {
        return Optional.ofNullable(this.postBrokerLoginFlowAlias);
    }

    /**
     * The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
     * 
     */
    @Import(name="providerId")
    private @Nullable Output<String> providerId;

    /**
     * @return The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
     * 
     */
    public Optional<Output<String>> providerId() {
        return Optional.ofNullable(this.providerId);
    }

    /**
     * The name of the realm. This is unique across Keycloak.
     * 
     */
    @Import(name="realm", required=true)
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
    @Import(name="storeToken")
    private @Nullable Output<Boolean> storeToken;

    /**
     * @return When `true`, tokens will be stored after authenticating users. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> storeToken() {
        return Optional.ofNullable(this.storeToken);
    }

    /**
     * The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     * 
     */
    @Import(name="syncMode")
    private @Nullable Output<String> syncMode;

    /**
     * @return The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
     * 
     */
    public Optional<Output<String>> syncMode() {
        return Optional.ofNullable(this.syncMode);
    }

    /**
     * The Token URL.
     * 
     */
    @Import(name="tokenUrl", required=true)
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
    @Import(name="trustEmail")
    private @Nullable Output<Boolean> trustEmail;

    /**
     * @return When `true`, email addresses for users in this provider will automatically be verified regardless of the realm&#39;s email verification policy. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> trustEmail() {
        return Optional.ofNullable(this.trustEmail);
    }

    /**
     * Pass current locale to identity provider. Defaults to `false`.
     * 
     */
    @Import(name="uiLocales")
    private @Nullable Output<Boolean> uiLocales;

    /**
     * @return Pass current locale to identity provider. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> uiLocales() {
        return Optional.ofNullable(this.uiLocales);
    }

    /**
     * User Info URL.
     * 
     */
    @Import(name="userInfoUrl")
    private @Nullable Output<String> userInfoUrl;

    /**
     * @return User Info URL.
     * 
     */
    public Optional<Output<String>> userInfoUrl() {
        return Optional.ofNullable(this.userInfoUrl);
    }

    /**
     * Enable/disable signature validation of external IDP signatures. Defaults to `false`.
     * 
     */
    @Import(name="validateSignature")
    private @Nullable Output<Boolean> validateSignature;

    /**
     * @return Enable/disable signature validation of external IDP signatures. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> validateSignature() {
        return Optional.ofNullable(this.validateSignature);
    }

    private IdentityProviderArgs() {}

    private IdentityProviderArgs(IdentityProviderArgs $) {
        this.acceptsPromptNoneForwardFromClient = $.acceptsPromptNoneForwardFromClient;
        this.addReadTokenRoleOnCreate = $.addReadTokenRoleOnCreate;
        this.alias = $.alias;
        this.authenticateByDefault = $.authenticateByDefault;
        this.authorizationUrl = $.authorizationUrl;
        this.backchannelSupported = $.backchannelSupported;
        this.clientId = $.clientId;
        this.clientSecret = $.clientSecret;
        this.defaultScopes = $.defaultScopes;
        this.disableUserInfo = $.disableUserInfo;
        this.displayName = $.displayName;
        this.enabled = $.enabled;
        this.extraConfig = $.extraConfig;
        this.firstBrokerLoginFlowAlias = $.firstBrokerLoginFlowAlias;
        this.guiOrder = $.guiOrder;
        this.hideOnLoginPage = $.hideOnLoginPage;
        this.jwksUrl = $.jwksUrl;
        this.linkOnly = $.linkOnly;
        this.loginHint = $.loginHint;
        this.logoutUrl = $.logoutUrl;
        this.postBrokerLoginFlowAlias = $.postBrokerLoginFlowAlias;
        this.providerId = $.providerId;
        this.realm = $.realm;
        this.storeToken = $.storeToken;
        this.syncMode = $.syncMode;
        this.tokenUrl = $.tokenUrl;
        this.trustEmail = $.trustEmail;
        this.uiLocales = $.uiLocales;
        this.userInfoUrl = $.userInfoUrl;
        this.validateSignature = $.validateSignature;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IdentityProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IdentityProviderArgs $;

        public Builder() {
            $ = new IdentityProviderArgs();
        }

        public Builder(IdentityProviderArgs defaults) {
            $ = new IdentityProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acceptsPromptNoneForwardFromClient When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder acceptsPromptNoneForwardFromClient(@Nullable Output<Boolean> acceptsPromptNoneForwardFromClient) {
            $.acceptsPromptNoneForwardFromClient = acceptsPromptNoneForwardFromClient;
            return this;
        }

        /**
         * @param acceptsPromptNoneForwardFromClient When `true`, the IDP will accept forwarded authentication requests that contain the `prompt=none` query parameter. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder acceptsPromptNoneForwardFromClient(Boolean acceptsPromptNoneForwardFromClient) {
            return acceptsPromptNoneForwardFromClient(Output.of(acceptsPromptNoneForwardFromClient));
        }

        /**
         * @param addReadTokenRoleOnCreate When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder addReadTokenRoleOnCreate(@Nullable Output<Boolean> addReadTokenRoleOnCreate) {
            $.addReadTokenRoleOnCreate = addReadTokenRoleOnCreate;
            return this;
        }

        /**
         * @param addReadTokenRoleOnCreate When `true`, new users will be able to read stored tokens. This will automatically assign the `broker.read-token` role. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder addReadTokenRoleOnCreate(Boolean addReadTokenRoleOnCreate) {
            return addReadTokenRoleOnCreate(Output.of(addReadTokenRoleOnCreate));
        }

        /**
         * @param alias The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
         * 
         * @return builder
         * 
         */
        public Builder alias(Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param authenticateByDefault Enable/disable authenticate users by default.
         * 
         * @return builder
         * 
         */
        public Builder authenticateByDefault(@Nullable Output<Boolean> authenticateByDefault) {
            $.authenticateByDefault = authenticateByDefault;
            return this;
        }

        /**
         * @param authenticateByDefault Enable/disable authenticate users by default.
         * 
         * @return builder
         * 
         */
        public Builder authenticateByDefault(Boolean authenticateByDefault) {
            return authenticateByDefault(Output.of(authenticateByDefault));
        }

        /**
         * @param authorizationUrl The Authorization Url.
         * 
         * @return builder
         * 
         */
        public Builder authorizationUrl(Output<String> authorizationUrl) {
            $.authorizationUrl = authorizationUrl;
            return this;
        }

        /**
         * @param authorizationUrl The Authorization Url.
         * 
         * @return builder
         * 
         */
        public Builder authorizationUrl(String authorizationUrl) {
            return authorizationUrl(Output.of(authorizationUrl));
        }

        /**
         * @param backchannelSupported Does the external IDP support backchannel logout? Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder backchannelSupported(@Nullable Output<Boolean> backchannelSupported) {
            $.backchannelSupported = backchannelSupported;
            return this;
        }

        /**
         * @param backchannelSupported Does the external IDP support backchannel logout? Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder backchannelSupported(Boolean backchannelSupported) {
            return backchannelSupported(Output.of(backchannelSupported));
        }

        /**
         * @param clientId The client or client identifier registered within the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder clientId(Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The client or client identifier registered within the identity provider.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param clientSecret The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
         * 
         * @return builder
         * 
         */
        public Builder clientSecret(Output<String> clientSecret) {
            $.clientSecret = clientSecret;
            return this;
        }

        /**
         * @param clientSecret The client or client secret registered within the identity provider. This field is able to obtain its value from vault, use $${vault.ID} format.
         * 
         * @return builder
         * 
         */
        public Builder clientSecret(String clientSecret) {
            return clientSecret(Output.of(clientSecret));
        }

        /**
         * @param defaultScopes The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
         * 
         * @return builder
         * 
         */
        public Builder defaultScopes(@Nullable Output<String> defaultScopes) {
            $.defaultScopes = defaultScopes;
            return this;
        }

        /**
         * @param defaultScopes The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to `openid`.
         * 
         * @return builder
         * 
         */
        public Builder defaultScopes(String defaultScopes) {
            return defaultScopes(Output.of(defaultScopes));
        }

        /**
         * @param disableUserInfo When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder disableUserInfo(@Nullable Output<Boolean> disableUserInfo) {
            $.disableUserInfo = disableUserInfo;
            return this;
        }

        /**
         * @param disableUserInfo When `true`, disables the usage of the user info service to obtain additional user information. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder disableUserInfo(Boolean disableUserInfo) {
            return disableUserInfo(Output.of(disableUserInfo));
        }

        /**
         * @param displayName Display name for the identity provider in the GUI.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Display name for the identity provider in the GUI.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param enabled When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled When `true`, users will be able to log in to this realm using this identity provider. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder extraConfig(@Nullable Output<Map<String,Object>> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        public Builder extraConfig(Map<String,Object> extraConfig) {
            return extraConfig(Output.of(extraConfig));
        }

        /**
         * @param firstBrokerLoginFlowAlias The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
         * 
         * @return builder
         * 
         */
        public Builder firstBrokerLoginFlowAlias(@Nullable Output<String> firstBrokerLoginFlowAlias) {
            $.firstBrokerLoginFlowAlias = firstBrokerLoginFlowAlias;
            return this;
        }

        /**
         * @param firstBrokerLoginFlowAlias The authentication flow to use when users log in for the first time through this identity provider. Defaults to `first broker login`.
         * 
         * @return builder
         * 
         */
        public Builder firstBrokerLoginFlowAlias(String firstBrokerLoginFlowAlias) {
            return firstBrokerLoginFlowAlias(Output.of(firstBrokerLoginFlowAlias));
        }

        /**
         * @param guiOrder A number defining the order of this identity provider in the GUI.
         * 
         * @return builder
         * 
         */
        public Builder guiOrder(@Nullable Output<String> guiOrder) {
            $.guiOrder = guiOrder;
            return this;
        }

        /**
         * @param guiOrder A number defining the order of this identity provider in the GUI.
         * 
         * @return builder
         * 
         */
        public Builder guiOrder(String guiOrder) {
            return guiOrder(Output.of(guiOrder));
        }

        /**
         * @param hideOnLoginPage When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder hideOnLoginPage(@Nullable Output<Boolean> hideOnLoginPage) {
            $.hideOnLoginPage = hideOnLoginPage;
            return this;
        }

        /**
         * @param hideOnLoginPage When `true`, this provider will be hidden on the login page, and is only accessible when requested explicitly. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder hideOnLoginPage(Boolean hideOnLoginPage) {
            return hideOnLoginPage(Output.of(hideOnLoginPage));
        }

        /**
         * @param jwksUrl JSON Web Key Set URL.
         * 
         * @return builder
         * 
         */
        public Builder jwksUrl(@Nullable Output<String> jwksUrl) {
            $.jwksUrl = jwksUrl;
            return this;
        }

        /**
         * @param jwksUrl JSON Web Key Set URL.
         * 
         * @return builder
         * 
         */
        public Builder jwksUrl(String jwksUrl) {
            return jwksUrl(Output.of(jwksUrl));
        }

        /**
         * @param linkOnly When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder linkOnly(@Nullable Output<Boolean> linkOnly) {
            $.linkOnly = linkOnly;
            return this;
        }

        /**
         * @param linkOnly When `true`, users cannot login using this provider, but their existing accounts will be linked when possible. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder linkOnly(Boolean linkOnly) {
            return linkOnly(Output.of(linkOnly));
        }

        /**
         * @param loginHint Pass login hint to identity provider.
         * 
         * @return builder
         * 
         */
        public Builder loginHint(@Nullable Output<String> loginHint) {
            $.loginHint = loginHint;
            return this;
        }

        /**
         * @param loginHint Pass login hint to identity provider.
         * 
         * @return builder
         * 
         */
        public Builder loginHint(String loginHint) {
            return loginHint(Output.of(loginHint));
        }

        /**
         * @param logoutUrl The Logout URL is the end session endpoint to use to logout user from external identity provider.
         * 
         * @return builder
         * 
         */
        public Builder logoutUrl(@Nullable Output<String> logoutUrl) {
            $.logoutUrl = logoutUrl;
            return this;
        }

        /**
         * @param logoutUrl The Logout URL is the end session endpoint to use to logout user from external identity provider.
         * 
         * @return builder
         * 
         */
        public Builder logoutUrl(String logoutUrl) {
            return logoutUrl(Output.of(logoutUrl));
        }

        /**
         * @param postBrokerLoginFlowAlias The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
         * 
         * @return builder
         * 
         */
        public Builder postBrokerLoginFlowAlias(@Nullable Output<String> postBrokerLoginFlowAlias) {
            $.postBrokerLoginFlowAlias = postBrokerLoginFlowAlias;
            return this;
        }

        /**
         * @param postBrokerLoginFlowAlias The authentication flow to use after users have successfully logged in, which can be used to perform additional user verification (such as OTP checking). Defaults to an empty string, which means no post login flow will be used.
         * 
         * @return builder
         * 
         */
        public Builder postBrokerLoginFlowAlias(String postBrokerLoginFlowAlias) {
            return postBrokerLoginFlowAlias(Output.of(postBrokerLoginFlowAlias));
        }

        /**
         * @param providerId The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
         * 
         * @return builder
         * 
         */
        public Builder providerId(@Nullable Output<String> providerId) {
            $.providerId = providerId;
            return this;
        }

        /**
         * @param providerId The ID of the identity provider to use. Defaults to `oidc`, which should be used unless you have extended Keycloak and provided your own implementation.
         * 
         * @return builder
         * 
         */
        public Builder providerId(String providerId) {
            return providerId(Output.of(providerId));
        }

        /**
         * @param realm The name of the realm. This is unique across Keycloak.
         * 
         * @return builder
         * 
         */
        public Builder realm(Output<String> realm) {
            $.realm = realm;
            return this;
        }

        /**
         * @param realm The name of the realm. This is unique across Keycloak.
         * 
         * @return builder
         * 
         */
        public Builder realm(String realm) {
            return realm(Output.of(realm));
        }

        /**
         * @param storeToken When `true`, tokens will be stored after authenticating users. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder storeToken(@Nullable Output<Boolean> storeToken) {
            $.storeToken = storeToken;
            return this;
        }

        /**
         * @param storeToken When `true`, tokens will be stored after authenticating users. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder storeToken(Boolean storeToken) {
            return storeToken(Output.of(storeToken));
        }

        /**
         * @param syncMode The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
         * 
         * @return builder
         * 
         */
        public Builder syncMode(@Nullable Output<String> syncMode) {
            $.syncMode = syncMode;
            return this;
        }

        /**
         * @param syncMode The default sync mode to use for all mappers attached to this identity provider. Can be once of `IMPORT`, `FORCE`, or `LEGACY`.
         * 
         * @return builder
         * 
         */
        public Builder syncMode(String syncMode) {
            return syncMode(Output.of(syncMode));
        }

        /**
         * @param tokenUrl The Token URL.
         * 
         * @return builder
         * 
         */
        public Builder tokenUrl(Output<String> tokenUrl) {
            $.tokenUrl = tokenUrl;
            return this;
        }

        /**
         * @param tokenUrl The Token URL.
         * 
         * @return builder
         * 
         */
        public Builder tokenUrl(String tokenUrl) {
            return tokenUrl(Output.of(tokenUrl));
        }

        /**
         * @param trustEmail When `true`, email addresses for users in this provider will automatically be verified regardless of the realm&#39;s email verification policy. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder trustEmail(@Nullable Output<Boolean> trustEmail) {
            $.trustEmail = trustEmail;
            return this;
        }

        /**
         * @param trustEmail When `true`, email addresses for users in this provider will automatically be verified regardless of the realm&#39;s email verification policy. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder trustEmail(Boolean trustEmail) {
            return trustEmail(Output.of(trustEmail));
        }

        /**
         * @param uiLocales Pass current locale to identity provider. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder uiLocales(@Nullable Output<Boolean> uiLocales) {
            $.uiLocales = uiLocales;
            return this;
        }

        /**
         * @param uiLocales Pass current locale to identity provider. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder uiLocales(Boolean uiLocales) {
            return uiLocales(Output.of(uiLocales));
        }

        /**
         * @param userInfoUrl User Info URL.
         * 
         * @return builder
         * 
         */
        public Builder userInfoUrl(@Nullable Output<String> userInfoUrl) {
            $.userInfoUrl = userInfoUrl;
            return this;
        }

        /**
         * @param userInfoUrl User Info URL.
         * 
         * @return builder
         * 
         */
        public Builder userInfoUrl(String userInfoUrl) {
            return userInfoUrl(Output.of(userInfoUrl));
        }

        /**
         * @param validateSignature Enable/disable signature validation of external IDP signatures. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder validateSignature(@Nullable Output<Boolean> validateSignature) {
            $.validateSignature = validateSignature;
            return this;
        }

        /**
         * @param validateSignature Enable/disable signature validation of external IDP signatures. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder validateSignature(Boolean validateSignature) {
            return validateSignature(Output.of(validateSignature));
        }

        public IdentityProviderArgs build() {
            $.alias = Objects.requireNonNull($.alias, "expected parameter 'alias' to be non-null");
            $.authorizationUrl = Objects.requireNonNull($.authorizationUrl, "expected parameter 'authorizationUrl' to be non-null");
            $.clientId = Objects.requireNonNull($.clientId, "expected parameter 'clientId' to be non-null");
            $.clientSecret = Objects.requireNonNull($.clientSecret, "expected parameter 'clientSecret' to be non-null");
            $.realm = Objects.requireNonNull($.realm, "expected parameter 'realm' to be non-null");
            $.tokenUrl = Objects.requireNonNull($.tokenUrl, "expected parameter 'tokenUrl' to be non-null");
            return $;
        }
    }

}
