// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.keycloak.openid.inputs.ClientAuthenticationFlowBindingOverridesArgs;
import com.pulumi.keycloak.openid.inputs.ClientAuthorizationArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientState extends com.pulumi.resources.ResourceArgs {

    public static final ClientState Empty = new ClientState();

    @Import(name="accessTokenLifespan")
    private @Nullable Output<String> accessTokenLifespan;

    public Optional<Output<String>> accessTokenLifespan() {
        return Optional.ofNullable(this.accessTokenLifespan);
    }

    @Import(name="accessType")
    private @Nullable Output<String> accessType;

    public Optional<Output<String>> accessType() {
        return Optional.ofNullable(this.accessType);
    }

    @Import(name="adminUrl")
    private @Nullable Output<String> adminUrl;

    public Optional<Output<String>> adminUrl() {
        return Optional.ofNullable(this.adminUrl);
    }

    @Import(name="authenticationFlowBindingOverrides")
    private @Nullable Output<ClientAuthenticationFlowBindingOverridesArgs> authenticationFlowBindingOverrides;

    public Optional<Output<ClientAuthenticationFlowBindingOverridesArgs>> authenticationFlowBindingOverrides() {
        return Optional.ofNullable(this.authenticationFlowBindingOverrides);
    }

    @Import(name="authorization")
    private @Nullable Output<ClientAuthorizationArgs> authorization;

    public Optional<Output<ClientAuthorizationArgs>> authorization() {
        return Optional.ofNullable(this.authorization);
    }

    @Import(name="backchannelLogoutRevokeOfflineSessions")
    private @Nullable Output<Boolean> backchannelLogoutRevokeOfflineSessions;

    public Optional<Output<Boolean>> backchannelLogoutRevokeOfflineSessions() {
        return Optional.ofNullable(this.backchannelLogoutRevokeOfflineSessions);
    }

    @Import(name="backchannelLogoutSessionRequired")
    private @Nullable Output<Boolean> backchannelLogoutSessionRequired;

    public Optional<Output<Boolean>> backchannelLogoutSessionRequired() {
        return Optional.ofNullable(this.backchannelLogoutSessionRequired);
    }

    @Import(name="backchannelLogoutUrl")
    private @Nullable Output<String> backchannelLogoutUrl;

    public Optional<Output<String>> backchannelLogoutUrl() {
        return Optional.ofNullable(this.backchannelLogoutUrl);
    }

    @Import(name="baseUrl")
    private @Nullable Output<String> baseUrl;

    public Optional<Output<String>> baseUrl() {
        return Optional.ofNullable(this.baseUrl);
    }

    @Import(name="clientAuthenticatorType")
    private @Nullable Output<String> clientAuthenticatorType;

    public Optional<Output<String>> clientAuthenticatorType() {
        return Optional.ofNullable(this.clientAuthenticatorType);
    }

    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    @Import(name="clientOfflineSessionIdleTimeout")
    private @Nullable Output<String> clientOfflineSessionIdleTimeout;

    public Optional<Output<String>> clientOfflineSessionIdleTimeout() {
        return Optional.ofNullable(this.clientOfflineSessionIdleTimeout);
    }

    @Import(name="clientOfflineSessionMaxLifespan")
    private @Nullable Output<String> clientOfflineSessionMaxLifespan;

    public Optional<Output<String>> clientOfflineSessionMaxLifespan() {
        return Optional.ofNullable(this.clientOfflineSessionMaxLifespan);
    }

    @Import(name="clientSecret")
    private @Nullable Output<String> clientSecret;

    public Optional<Output<String>> clientSecret() {
        return Optional.ofNullable(this.clientSecret);
    }

    @Import(name="clientSessionIdleTimeout")
    private @Nullable Output<String> clientSessionIdleTimeout;

    public Optional<Output<String>> clientSessionIdleTimeout() {
        return Optional.ofNullable(this.clientSessionIdleTimeout);
    }

    @Import(name="clientSessionMaxLifespan")
    private @Nullable Output<String> clientSessionMaxLifespan;

    public Optional<Output<String>> clientSessionMaxLifespan() {
        return Optional.ofNullable(this.clientSessionMaxLifespan);
    }

    @Import(name="consentRequired")
    private @Nullable Output<Boolean> consentRequired;

    public Optional<Output<Boolean>> consentRequired() {
        return Optional.ofNullable(this.consentRequired);
    }

    @Import(name="consentScreenText")
    private @Nullable Output<String> consentScreenText;

    public Optional<Output<String>> consentScreenText() {
        return Optional.ofNullable(this.consentScreenText);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="directAccessGrantsEnabled")
    private @Nullable Output<Boolean> directAccessGrantsEnabled;

    public Optional<Output<Boolean>> directAccessGrantsEnabled() {
        return Optional.ofNullable(this.directAccessGrantsEnabled);
    }

    @Import(name="displayOnConsentScreen")
    private @Nullable Output<Boolean> displayOnConsentScreen;

    public Optional<Output<Boolean>> displayOnConsentScreen() {
        return Optional.ofNullable(this.displayOnConsentScreen);
    }

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="excludeSessionStateFromAuthResponse")
    private @Nullable Output<Boolean> excludeSessionStateFromAuthResponse;

    public Optional<Output<Boolean>> excludeSessionStateFromAuthResponse() {
        return Optional.ofNullable(this.excludeSessionStateFromAuthResponse);
    }

    @Import(name="extraConfig")
    private @Nullable Output<Map<String,String>> extraConfig;

    public Optional<Output<Map<String,String>>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    @Import(name="frontchannelLogoutEnabled")
    private @Nullable Output<Boolean> frontchannelLogoutEnabled;

    public Optional<Output<Boolean>> frontchannelLogoutEnabled() {
        return Optional.ofNullable(this.frontchannelLogoutEnabled);
    }

    @Import(name="frontchannelLogoutUrl")
    private @Nullable Output<String> frontchannelLogoutUrl;

    public Optional<Output<String>> frontchannelLogoutUrl() {
        return Optional.ofNullable(this.frontchannelLogoutUrl);
    }

    @Import(name="fullScopeAllowed")
    private @Nullable Output<Boolean> fullScopeAllowed;

    public Optional<Output<Boolean>> fullScopeAllowed() {
        return Optional.ofNullable(this.fullScopeAllowed);
    }

    @Import(name="implicitFlowEnabled")
    private @Nullable Output<Boolean> implicitFlowEnabled;

    public Optional<Output<Boolean>> implicitFlowEnabled() {
        return Optional.ofNullable(this.implicitFlowEnabled);
    }

    @Import(name="import")
    private @Nullable Output<Boolean> import_;

    public Optional<Output<Boolean>> import_() {
        return Optional.ofNullable(this.import_);
    }

    @Import(name="loginTheme")
    private @Nullable Output<String> loginTheme;

    public Optional<Output<String>> loginTheme() {
        return Optional.ofNullable(this.loginTheme);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="oauth2DeviceAuthorizationGrantEnabled")
    private @Nullable Output<Boolean> oauth2DeviceAuthorizationGrantEnabled;

    public Optional<Output<Boolean>> oauth2DeviceAuthorizationGrantEnabled() {
        return Optional.ofNullable(this.oauth2DeviceAuthorizationGrantEnabled);
    }

    @Import(name="oauth2DeviceCodeLifespan")
    private @Nullable Output<String> oauth2DeviceCodeLifespan;

    public Optional<Output<String>> oauth2DeviceCodeLifespan() {
        return Optional.ofNullable(this.oauth2DeviceCodeLifespan);
    }

    @Import(name="oauth2DevicePollingInterval")
    private @Nullable Output<String> oauth2DevicePollingInterval;

    public Optional<Output<String>> oauth2DevicePollingInterval() {
        return Optional.ofNullable(this.oauth2DevicePollingInterval);
    }

    @Import(name="pkceCodeChallengeMethod")
    private @Nullable Output<String> pkceCodeChallengeMethod;

    public Optional<Output<String>> pkceCodeChallengeMethod() {
        return Optional.ofNullable(this.pkceCodeChallengeMethod);
    }

    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    @Import(name="resourceServerId")
    private @Nullable Output<String> resourceServerId;

    public Optional<Output<String>> resourceServerId() {
        return Optional.ofNullable(this.resourceServerId);
    }

    @Import(name="rootUrl")
    private @Nullable Output<String> rootUrl;

    public Optional<Output<String>> rootUrl() {
        return Optional.ofNullable(this.rootUrl);
    }

    @Import(name="serviceAccountUserId")
    private @Nullable Output<String> serviceAccountUserId;

    public Optional<Output<String>> serviceAccountUserId() {
        return Optional.ofNullable(this.serviceAccountUserId);
    }

    @Import(name="serviceAccountsEnabled")
    private @Nullable Output<Boolean> serviceAccountsEnabled;

    public Optional<Output<Boolean>> serviceAccountsEnabled() {
        return Optional.ofNullable(this.serviceAccountsEnabled);
    }

    @Import(name="standardFlowEnabled")
    private @Nullable Output<Boolean> standardFlowEnabled;

    public Optional<Output<Boolean>> standardFlowEnabled() {
        return Optional.ofNullable(this.standardFlowEnabled);
    }

    @Import(name="useRefreshTokens")
    private @Nullable Output<Boolean> useRefreshTokens;

    public Optional<Output<Boolean>> useRefreshTokens() {
        return Optional.ofNullable(this.useRefreshTokens);
    }

    @Import(name="useRefreshTokensClientCredentials")
    private @Nullable Output<Boolean> useRefreshTokensClientCredentials;

    public Optional<Output<Boolean>> useRefreshTokensClientCredentials() {
        return Optional.ofNullable(this.useRefreshTokensClientCredentials);
    }

    @Import(name="validPostLogoutRedirectUris")
    private @Nullable Output<List<String>> validPostLogoutRedirectUris;

    public Optional<Output<List<String>>> validPostLogoutRedirectUris() {
        return Optional.ofNullable(this.validPostLogoutRedirectUris);
    }

    @Import(name="validRedirectUris")
    private @Nullable Output<List<String>> validRedirectUris;

    public Optional<Output<List<String>>> validRedirectUris() {
        return Optional.ofNullable(this.validRedirectUris);
    }

    @Import(name="webOrigins")
    private @Nullable Output<List<String>> webOrigins;

    public Optional<Output<List<String>>> webOrigins() {
        return Optional.ofNullable(this.webOrigins);
    }

    private ClientState() {}

    private ClientState(ClientState $) {
        this.accessTokenLifespan = $.accessTokenLifespan;
        this.accessType = $.accessType;
        this.adminUrl = $.adminUrl;
        this.authenticationFlowBindingOverrides = $.authenticationFlowBindingOverrides;
        this.authorization = $.authorization;
        this.backchannelLogoutRevokeOfflineSessions = $.backchannelLogoutRevokeOfflineSessions;
        this.backchannelLogoutSessionRequired = $.backchannelLogoutSessionRequired;
        this.backchannelLogoutUrl = $.backchannelLogoutUrl;
        this.baseUrl = $.baseUrl;
        this.clientAuthenticatorType = $.clientAuthenticatorType;
        this.clientId = $.clientId;
        this.clientOfflineSessionIdleTimeout = $.clientOfflineSessionIdleTimeout;
        this.clientOfflineSessionMaxLifespan = $.clientOfflineSessionMaxLifespan;
        this.clientSecret = $.clientSecret;
        this.clientSessionIdleTimeout = $.clientSessionIdleTimeout;
        this.clientSessionMaxLifespan = $.clientSessionMaxLifespan;
        this.consentRequired = $.consentRequired;
        this.consentScreenText = $.consentScreenText;
        this.description = $.description;
        this.directAccessGrantsEnabled = $.directAccessGrantsEnabled;
        this.displayOnConsentScreen = $.displayOnConsentScreen;
        this.enabled = $.enabled;
        this.excludeSessionStateFromAuthResponse = $.excludeSessionStateFromAuthResponse;
        this.extraConfig = $.extraConfig;
        this.frontchannelLogoutEnabled = $.frontchannelLogoutEnabled;
        this.frontchannelLogoutUrl = $.frontchannelLogoutUrl;
        this.fullScopeAllowed = $.fullScopeAllowed;
        this.implicitFlowEnabled = $.implicitFlowEnabled;
        this.import_ = $.import_;
        this.loginTheme = $.loginTheme;
        this.name = $.name;
        this.oauth2DeviceAuthorizationGrantEnabled = $.oauth2DeviceAuthorizationGrantEnabled;
        this.oauth2DeviceCodeLifespan = $.oauth2DeviceCodeLifespan;
        this.oauth2DevicePollingInterval = $.oauth2DevicePollingInterval;
        this.pkceCodeChallengeMethod = $.pkceCodeChallengeMethod;
        this.realmId = $.realmId;
        this.resourceServerId = $.resourceServerId;
        this.rootUrl = $.rootUrl;
        this.serviceAccountUserId = $.serviceAccountUserId;
        this.serviceAccountsEnabled = $.serviceAccountsEnabled;
        this.standardFlowEnabled = $.standardFlowEnabled;
        this.useRefreshTokens = $.useRefreshTokens;
        this.useRefreshTokensClientCredentials = $.useRefreshTokensClientCredentials;
        this.validPostLogoutRedirectUris = $.validPostLogoutRedirectUris;
        this.validRedirectUris = $.validRedirectUris;
        this.webOrigins = $.webOrigins;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientState $;

        public Builder() {
            $ = new ClientState();
        }

        public Builder(ClientState defaults) {
            $ = new ClientState(Objects.requireNonNull(defaults));
        }

        public Builder accessTokenLifespan(@Nullable Output<String> accessTokenLifespan) {
            $.accessTokenLifespan = accessTokenLifespan;
            return this;
        }

        public Builder accessTokenLifespan(String accessTokenLifespan) {
            return accessTokenLifespan(Output.of(accessTokenLifespan));
        }

        public Builder accessType(@Nullable Output<String> accessType) {
            $.accessType = accessType;
            return this;
        }

        public Builder accessType(String accessType) {
            return accessType(Output.of(accessType));
        }

        public Builder adminUrl(@Nullable Output<String> adminUrl) {
            $.adminUrl = adminUrl;
            return this;
        }

        public Builder adminUrl(String adminUrl) {
            return adminUrl(Output.of(adminUrl));
        }

        public Builder authenticationFlowBindingOverrides(@Nullable Output<ClientAuthenticationFlowBindingOverridesArgs> authenticationFlowBindingOverrides) {
            $.authenticationFlowBindingOverrides = authenticationFlowBindingOverrides;
            return this;
        }

        public Builder authenticationFlowBindingOverrides(ClientAuthenticationFlowBindingOverridesArgs authenticationFlowBindingOverrides) {
            return authenticationFlowBindingOverrides(Output.of(authenticationFlowBindingOverrides));
        }

        public Builder authorization(@Nullable Output<ClientAuthorizationArgs> authorization) {
            $.authorization = authorization;
            return this;
        }

        public Builder authorization(ClientAuthorizationArgs authorization) {
            return authorization(Output.of(authorization));
        }

        public Builder backchannelLogoutRevokeOfflineSessions(@Nullable Output<Boolean> backchannelLogoutRevokeOfflineSessions) {
            $.backchannelLogoutRevokeOfflineSessions = backchannelLogoutRevokeOfflineSessions;
            return this;
        }

        public Builder backchannelLogoutRevokeOfflineSessions(Boolean backchannelLogoutRevokeOfflineSessions) {
            return backchannelLogoutRevokeOfflineSessions(Output.of(backchannelLogoutRevokeOfflineSessions));
        }

        public Builder backchannelLogoutSessionRequired(@Nullable Output<Boolean> backchannelLogoutSessionRequired) {
            $.backchannelLogoutSessionRequired = backchannelLogoutSessionRequired;
            return this;
        }

        public Builder backchannelLogoutSessionRequired(Boolean backchannelLogoutSessionRequired) {
            return backchannelLogoutSessionRequired(Output.of(backchannelLogoutSessionRequired));
        }

        public Builder backchannelLogoutUrl(@Nullable Output<String> backchannelLogoutUrl) {
            $.backchannelLogoutUrl = backchannelLogoutUrl;
            return this;
        }

        public Builder backchannelLogoutUrl(String backchannelLogoutUrl) {
            return backchannelLogoutUrl(Output.of(backchannelLogoutUrl));
        }

        public Builder baseUrl(@Nullable Output<String> baseUrl) {
            $.baseUrl = baseUrl;
            return this;
        }

        public Builder baseUrl(String baseUrl) {
            return baseUrl(Output.of(baseUrl));
        }

        public Builder clientAuthenticatorType(@Nullable Output<String> clientAuthenticatorType) {
            $.clientAuthenticatorType = clientAuthenticatorType;
            return this;
        }

        public Builder clientAuthenticatorType(String clientAuthenticatorType) {
            return clientAuthenticatorType(Output.of(clientAuthenticatorType));
        }

        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        public Builder clientOfflineSessionIdleTimeout(@Nullable Output<String> clientOfflineSessionIdleTimeout) {
            $.clientOfflineSessionIdleTimeout = clientOfflineSessionIdleTimeout;
            return this;
        }

        public Builder clientOfflineSessionIdleTimeout(String clientOfflineSessionIdleTimeout) {
            return clientOfflineSessionIdleTimeout(Output.of(clientOfflineSessionIdleTimeout));
        }

        public Builder clientOfflineSessionMaxLifespan(@Nullable Output<String> clientOfflineSessionMaxLifespan) {
            $.clientOfflineSessionMaxLifespan = clientOfflineSessionMaxLifespan;
            return this;
        }

        public Builder clientOfflineSessionMaxLifespan(String clientOfflineSessionMaxLifespan) {
            return clientOfflineSessionMaxLifespan(Output.of(clientOfflineSessionMaxLifespan));
        }

        public Builder clientSecret(@Nullable Output<String> clientSecret) {
            $.clientSecret = clientSecret;
            return this;
        }

        public Builder clientSecret(String clientSecret) {
            return clientSecret(Output.of(clientSecret));
        }

        public Builder clientSessionIdleTimeout(@Nullable Output<String> clientSessionIdleTimeout) {
            $.clientSessionIdleTimeout = clientSessionIdleTimeout;
            return this;
        }

        public Builder clientSessionIdleTimeout(String clientSessionIdleTimeout) {
            return clientSessionIdleTimeout(Output.of(clientSessionIdleTimeout));
        }

        public Builder clientSessionMaxLifespan(@Nullable Output<String> clientSessionMaxLifespan) {
            $.clientSessionMaxLifespan = clientSessionMaxLifespan;
            return this;
        }

        public Builder clientSessionMaxLifespan(String clientSessionMaxLifespan) {
            return clientSessionMaxLifespan(Output.of(clientSessionMaxLifespan));
        }

        public Builder consentRequired(@Nullable Output<Boolean> consentRequired) {
            $.consentRequired = consentRequired;
            return this;
        }

        public Builder consentRequired(Boolean consentRequired) {
            return consentRequired(Output.of(consentRequired));
        }

        public Builder consentScreenText(@Nullable Output<String> consentScreenText) {
            $.consentScreenText = consentScreenText;
            return this;
        }

        public Builder consentScreenText(String consentScreenText) {
            return consentScreenText(Output.of(consentScreenText));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder directAccessGrantsEnabled(@Nullable Output<Boolean> directAccessGrantsEnabled) {
            $.directAccessGrantsEnabled = directAccessGrantsEnabled;
            return this;
        }

        public Builder directAccessGrantsEnabled(Boolean directAccessGrantsEnabled) {
            return directAccessGrantsEnabled(Output.of(directAccessGrantsEnabled));
        }

        public Builder displayOnConsentScreen(@Nullable Output<Boolean> displayOnConsentScreen) {
            $.displayOnConsentScreen = displayOnConsentScreen;
            return this;
        }

        public Builder displayOnConsentScreen(Boolean displayOnConsentScreen) {
            return displayOnConsentScreen(Output.of(displayOnConsentScreen));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder excludeSessionStateFromAuthResponse(@Nullable Output<Boolean> excludeSessionStateFromAuthResponse) {
            $.excludeSessionStateFromAuthResponse = excludeSessionStateFromAuthResponse;
            return this;
        }

        public Builder excludeSessionStateFromAuthResponse(Boolean excludeSessionStateFromAuthResponse) {
            return excludeSessionStateFromAuthResponse(Output.of(excludeSessionStateFromAuthResponse));
        }

        public Builder extraConfig(@Nullable Output<Map<String,String>> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        public Builder extraConfig(Map<String,String> extraConfig) {
            return extraConfig(Output.of(extraConfig));
        }

        public Builder frontchannelLogoutEnabled(@Nullable Output<Boolean> frontchannelLogoutEnabled) {
            $.frontchannelLogoutEnabled = frontchannelLogoutEnabled;
            return this;
        }

        public Builder frontchannelLogoutEnabled(Boolean frontchannelLogoutEnabled) {
            return frontchannelLogoutEnabled(Output.of(frontchannelLogoutEnabled));
        }

        public Builder frontchannelLogoutUrl(@Nullable Output<String> frontchannelLogoutUrl) {
            $.frontchannelLogoutUrl = frontchannelLogoutUrl;
            return this;
        }

        public Builder frontchannelLogoutUrl(String frontchannelLogoutUrl) {
            return frontchannelLogoutUrl(Output.of(frontchannelLogoutUrl));
        }

        public Builder fullScopeAllowed(@Nullable Output<Boolean> fullScopeAllowed) {
            $.fullScopeAllowed = fullScopeAllowed;
            return this;
        }

        public Builder fullScopeAllowed(Boolean fullScopeAllowed) {
            return fullScopeAllowed(Output.of(fullScopeAllowed));
        }

        public Builder implicitFlowEnabled(@Nullable Output<Boolean> implicitFlowEnabled) {
            $.implicitFlowEnabled = implicitFlowEnabled;
            return this;
        }

        public Builder implicitFlowEnabled(Boolean implicitFlowEnabled) {
            return implicitFlowEnabled(Output.of(implicitFlowEnabled));
        }

        public Builder import_(@Nullable Output<Boolean> import_) {
            $.import_ = import_;
            return this;
        }

        public Builder import_(Boolean import_) {
            return import_(Output.of(import_));
        }

        public Builder loginTheme(@Nullable Output<String> loginTheme) {
            $.loginTheme = loginTheme;
            return this;
        }

        public Builder loginTheme(String loginTheme) {
            return loginTheme(Output.of(loginTheme));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder oauth2DeviceAuthorizationGrantEnabled(@Nullable Output<Boolean> oauth2DeviceAuthorizationGrantEnabled) {
            $.oauth2DeviceAuthorizationGrantEnabled = oauth2DeviceAuthorizationGrantEnabled;
            return this;
        }

        public Builder oauth2DeviceAuthorizationGrantEnabled(Boolean oauth2DeviceAuthorizationGrantEnabled) {
            return oauth2DeviceAuthorizationGrantEnabled(Output.of(oauth2DeviceAuthorizationGrantEnabled));
        }

        public Builder oauth2DeviceCodeLifespan(@Nullable Output<String> oauth2DeviceCodeLifespan) {
            $.oauth2DeviceCodeLifespan = oauth2DeviceCodeLifespan;
            return this;
        }

        public Builder oauth2DeviceCodeLifespan(String oauth2DeviceCodeLifespan) {
            return oauth2DeviceCodeLifespan(Output.of(oauth2DeviceCodeLifespan));
        }

        public Builder oauth2DevicePollingInterval(@Nullable Output<String> oauth2DevicePollingInterval) {
            $.oauth2DevicePollingInterval = oauth2DevicePollingInterval;
            return this;
        }

        public Builder oauth2DevicePollingInterval(String oauth2DevicePollingInterval) {
            return oauth2DevicePollingInterval(Output.of(oauth2DevicePollingInterval));
        }

        public Builder pkceCodeChallengeMethod(@Nullable Output<String> pkceCodeChallengeMethod) {
            $.pkceCodeChallengeMethod = pkceCodeChallengeMethod;
            return this;
        }

        public Builder pkceCodeChallengeMethod(String pkceCodeChallengeMethod) {
            return pkceCodeChallengeMethod(Output.of(pkceCodeChallengeMethod));
        }

        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public Builder resourceServerId(@Nullable Output<String> resourceServerId) {
            $.resourceServerId = resourceServerId;
            return this;
        }

        public Builder resourceServerId(String resourceServerId) {
            return resourceServerId(Output.of(resourceServerId));
        }

        public Builder rootUrl(@Nullable Output<String> rootUrl) {
            $.rootUrl = rootUrl;
            return this;
        }

        public Builder rootUrl(String rootUrl) {
            return rootUrl(Output.of(rootUrl));
        }

        public Builder serviceAccountUserId(@Nullable Output<String> serviceAccountUserId) {
            $.serviceAccountUserId = serviceAccountUserId;
            return this;
        }

        public Builder serviceAccountUserId(String serviceAccountUserId) {
            return serviceAccountUserId(Output.of(serviceAccountUserId));
        }

        public Builder serviceAccountsEnabled(@Nullable Output<Boolean> serviceAccountsEnabled) {
            $.serviceAccountsEnabled = serviceAccountsEnabled;
            return this;
        }

        public Builder serviceAccountsEnabled(Boolean serviceAccountsEnabled) {
            return serviceAccountsEnabled(Output.of(serviceAccountsEnabled));
        }

        public Builder standardFlowEnabled(@Nullable Output<Boolean> standardFlowEnabled) {
            $.standardFlowEnabled = standardFlowEnabled;
            return this;
        }

        public Builder standardFlowEnabled(Boolean standardFlowEnabled) {
            return standardFlowEnabled(Output.of(standardFlowEnabled));
        }

        public Builder useRefreshTokens(@Nullable Output<Boolean> useRefreshTokens) {
            $.useRefreshTokens = useRefreshTokens;
            return this;
        }

        public Builder useRefreshTokens(Boolean useRefreshTokens) {
            return useRefreshTokens(Output.of(useRefreshTokens));
        }

        public Builder useRefreshTokensClientCredentials(@Nullable Output<Boolean> useRefreshTokensClientCredentials) {
            $.useRefreshTokensClientCredentials = useRefreshTokensClientCredentials;
            return this;
        }

        public Builder useRefreshTokensClientCredentials(Boolean useRefreshTokensClientCredentials) {
            return useRefreshTokensClientCredentials(Output.of(useRefreshTokensClientCredentials));
        }

        public Builder validPostLogoutRedirectUris(@Nullable Output<List<String>> validPostLogoutRedirectUris) {
            $.validPostLogoutRedirectUris = validPostLogoutRedirectUris;
            return this;
        }

        public Builder validPostLogoutRedirectUris(List<String> validPostLogoutRedirectUris) {
            return validPostLogoutRedirectUris(Output.of(validPostLogoutRedirectUris));
        }

        public Builder validPostLogoutRedirectUris(String... validPostLogoutRedirectUris) {
            return validPostLogoutRedirectUris(List.of(validPostLogoutRedirectUris));
        }

        public Builder validRedirectUris(@Nullable Output<List<String>> validRedirectUris) {
            $.validRedirectUris = validRedirectUris;
            return this;
        }

        public Builder validRedirectUris(List<String> validRedirectUris) {
            return validRedirectUris(Output.of(validRedirectUris));
        }

        public Builder validRedirectUris(String... validRedirectUris) {
            return validRedirectUris(List.of(validRedirectUris));
        }

        public Builder webOrigins(@Nullable Output<List<String>> webOrigins) {
            $.webOrigins = webOrigins;
            return this;
        }

        public Builder webOrigins(List<String> webOrigins) {
            return webOrigins(Output.of(webOrigins));
        }

        public Builder webOrigins(String... webOrigins) {
            return webOrigins(List.of(webOrigins));
        }

        public ClientState build() {
            return $;
        }
    }

}
