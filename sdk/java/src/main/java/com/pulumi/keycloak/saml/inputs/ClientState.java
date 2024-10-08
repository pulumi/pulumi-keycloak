// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.keycloak.saml.inputs.ClientAuthenticationFlowBindingOverridesArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientState extends com.pulumi.resources.ResourceArgs {

    public static final ClientState Empty = new ClientState();

    @Import(name="assertionConsumerPostUrl")
    private @Nullable Output<String> assertionConsumerPostUrl;

    public Optional<Output<String>> assertionConsumerPostUrl() {
        return Optional.ofNullable(this.assertionConsumerPostUrl);
    }

    @Import(name="assertionConsumerRedirectUrl")
    private @Nullable Output<String> assertionConsumerRedirectUrl;

    public Optional<Output<String>> assertionConsumerRedirectUrl() {
        return Optional.ofNullable(this.assertionConsumerRedirectUrl);
    }

    @Import(name="authenticationFlowBindingOverrides")
    private @Nullable Output<ClientAuthenticationFlowBindingOverridesArgs> authenticationFlowBindingOverrides;

    public Optional<Output<ClientAuthenticationFlowBindingOverridesArgs>> authenticationFlowBindingOverrides() {
        return Optional.ofNullable(this.authenticationFlowBindingOverrides);
    }

    @Import(name="baseUrl")
    private @Nullable Output<String> baseUrl;

    public Optional<Output<String>> baseUrl() {
        return Optional.ofNullable(this.baseUrl);
    }

    @Import(name="canonicalizationMethod")
    private @Nullable Output<String> canonicalizationMethod;

    public Optional<Output<String>> canonicalizationMethod() {
        return Optional.ofNullable(this.canonicalizationMethod);
    }

    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    @Import(name="clientSignatureRequired")
    private @Nullable Output<Boolean> clientSignatureRequired;

    public Optional<Output<Boolean>> clientSignatureRequired() {
        return Optional.ofNullable(this.clientSignatureRequired);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="encryptAssertions")
    private @Nullable Output<Boolean> encryptAssertions;

    public Optional<Output<Boolean>> encryptAssertions() {
        return Optional.ofNullable(this.encryptAssertions);
    }

    @Import(name="encryptionCertificate")
    private @Nullable Output<String> encryptionCertificate;

    public Optional<Output<String>> encryptionCertificate() {
        return Optional.ofNullable(this.encryptionCertificate);
    }

    @Import(name="encryptionCertificateSha1")
    private @Nullable Output<String> encryptionCertificateSha1;

    public Optional<Output<String>> encryptionCertificateSha1() {
        return Optional.ofNullable(this.encryptionCertificateSha1);
    }

    @Import(name="extraConfig")
    private @Nullable Output<Map<String,String>> extraConfig;

    public Optional<Output<Map<String,String>>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    @Import(name="forceNameIdFormat")
    private @Nullable Output<Boolean> forceNameIdFormat;

    public Optional<Output<Boolean>> forceNameIdFormat() {
        return Optional.ofNullable(this.forceNameIdFormat);
    }

    @Import(name="forcePostBinding")
    private @Nullable Output<Boolean> forcePostBinding;

    public Optional<Output<Boolean>> forcePostBinding() {
        return Optional.ofNullable(this.forcePostBinding);
    }

    @Import(name="frontChannelLogout")
    private @Nullable Output<Boolean> frontChannelLogout;

    public Optional<Output<Boolean>> frontChannelLogout() {
        return Optional.ofNullable(this.frontChannelLogout);
    }

    @Import(name="fullScopeAllowed")
    private @Nullable Output<Boolean> fullScopeAllowed;

    public Optional<Output<Boolean>> fullScopeAllowed() {
        return Optional.ofNullable(this.fullScopeAllowed);
    }

    @Import(name="idpInitiatedSsoRelayState")
    private @Nullable Output<String> idpInitiatedSsoRelayState;

    public Optional<Output<String>> idpInitiatedSsoRelayState() {
        return Optional.ofNullable(this.idpInitiatedSsoRelayState);
    }

    @Import(name="idpInitiatedSsoUrlName")
    private @Nullable Output<String> idpInitiatedSsoUrlName;

    public Optional<Output<String>> idpInitiatedSsoUrlName() {
        return Optional.ofNullable(this.idpInitiatedSsoUrlName);
    }

    @Import(name="includeAuthnStatement")
    private @Nullable Output<Boolean> includeAuthnStatement;

    public Optional<Output<Boolean>> includeAuthnStatement() {
        return Optional.ofNullable(this.includeAuthnStatement);
    }

    @Import(name="loginTheme")
    private @Nullable Output<String> loginTheme;

    public Optional<Output<String>> loginTheme() {
        return Optional.ofNullable(this.loginTheme);
    }

    @Import(name="logoutServicePostBindingUrl")
    private @Nullable Output<String> logoutServicePostBindingUrl;

    public Optional<Output<String>> logoutServicePostBindingUrl() {
        return Optional.ofNullable(this.logoutServicePostBindingUrl);
    }

    @Import(name="logoutServiceRedirectBindingUrl")
    private @Nullable Output<String> logoutServiceRedirectBindingUrl;

    public Optional<Output<String>> logoutServiceRedirectBindingUrl() {
        return Optional.ofNullable(this.logoutServiceRedirectBindingUrl);
    }

    @Import(name="masterSamlProcessingUrl")
    private @Nullable Output<String> masterSamlProcessingUrl;

    public Optional<Output<String>> masterSamlProcessingUrl() {
        return Optional.ofNullable(this.masterSamlProcessingUrl);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="nameIdFormat")
    private @Nullable Output<String> nameIdFormat;

    public Optional<Output<String>> nameIdFormat() {
        return Optional.ofNullable(this.nameIdFormat);
    }

    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    @Import(name="rootUrl")
    private @Nullable Output<String> rootUrl;

    public Optional<Output<String>> rootUrl() {
        return Optional.ofNullable(this.rootUrl);
    }

    @Import(name="signAssertions")
    private @Nullable Output<Boolean> signAssertions;

    public Optional<Output<Boolean>> signAssertions() {
        return Optional.ofNullable(this.signAssertions);
    }

    @Import(name="signDocuments")
    private @Nullable Output<Boolean> signDocuments;

    public Optional<Output<Boolean>> signDocuments() {
        return Optional.ofNullable(this.signDocuments);
    }

    @Import(name="signatureAlgorithm")
    private @Nullable Output<String> signatureAlgorithm;

    public Optional<Output<String>> signatureAlgorithm() {
        return Optional.ofNullable(this.signatureAlgorithm);
    }

    @Import(name="signatureKeyName")
    private @Nullable Output<String> signatureKeyName;

    public Optional<Output<String>> signatureKeyName() {
        return Optional.ofNullable(this.signatureKeyName);
    }

    @Import(name="signingCertificate")
    private @Nullable Output<String> signingCertificate;

    public Optional<Output<String>> signingCertificate() {
        return Optional.ofNullable(this.signingCertificate);
    }

    @Import(name="signingCertificateSha1")
    private @Nullable Output<String> signingCertificateSha1;

    public Optional<Output<String>> signingCertificateSha1() {
        return Optional.ofNullable(this.signingCertificateSha1);
    }

    @Import(name="signingPrivateKey")
    private @Nullable Output<String> signingPrivateKey;

    public Optional<Output<String>> signingPrivateKey() {
        return Optional.ofNullable(this.signingPrivateKey);
    }

    @Import(name="signingPrivateKeySha1")
    private @Nullable Output<String> signingPrivateKeySha1;

    public Optional<Output<String>> signingPrivateKeySha1() {
        return Optional.ofNullable(this.signingPrivateKeySha1);
    }

    @Import(name="validRedirectUris")
    private @Nullable Output<List<String>> validRedirectUris;

    public Optional<Output<List<String>>> validRedirectUris() {
        return Optional.ofNullable(this.validRedirectUris);
    }

    private ClientState() {}

    private ClientState(ClientState $) {
        this.assertionConsumerPostUrl = $.assertionConsumerPostUrl;
        this.assertionConsumerRedirectUrl = $.assertionConsumerRedirectUrl;
        this.authenticationFlowBindingOverrides = $.authenticationFlowBindingOverrides;
        this.baseUrl = $.baseUrl;
        this.canonicalizationMethod = $.canonicalizationMethod;
        this.clientId = $.clientId;
        this.clientSignatureRequired = $.clientSignatureRequired;
        this.description = $.description;
        this.enabled = $.enabled;
        this.encryptAssertions = $.encryptAssertions;
        this.encryptionCertificate = $.encryptionCertificate;
        this.encryptionCertificateSha1 = $.encryptionCertificateSha1;
        this.extraConfig = $.extraConfig;
        this.forceNameIdFormat = $.forceNameIdFormat;
        this.forcePostBinding = $.forcePostBinding;
        this.frontChannelLogout = $.frontChannelLogout;
        this.fullScopeAllowed = $.fullScopeAllowed;
        this.idpInitiatedSsoRelayState = $.idpInitiatedSsoRelayState;
        this.idpInitiatedSsoUrlName = $.idpInitiatedSsoUrlName;
        this.includeAuthnStatement = $.includeAuthnStatement;
        this.loginTheme = $.loginTheme;
        this.logoutServicePostBindingUrl = $.logoutServicePostBindingUrl;
        this.logoutServiceRedirectBindingUrl = $.logoutServiceRedirectBindingUrl;
        this.masterSamlProcessingUrl = $.masterSamlProcessingUrl;
        this.name = $.name;
        this.nameIdFormat = $.nameIdFormat;
        this.realmId = $.realmId;
        this.rootUrl = $.rootUrl;
        this.signAssertions = $.signAssertions;
        this.signDocuments = $.signDocuments;
        this.signatureAlgorithm = $.signatureAlgorithm;
        this.signatureKeyName = $.signatureKeyName;
        this.signingCertificate = $.signingCertificate;
        this.signingCertificateSha1 = $.signingCertificateSha1;
        this.signingPrivateKey = $.signingPrivateKey;
        this.signingPrivateKeySha1 = $.signingPrivateKeySha1;
        this.validRedirectUris = $.validRedirectUris;
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

        public Builder assertionConsumerPostUrl(@Nullable Output<String> assertionConsumerPostUrl) {
            $.assertionConsumerPostUrl = assertionConsumerPostUrl;
            return this;
        }

        public Builder assertionConsumerPostUrl(String assertionConsumerPostUrl) {
            return assertionConsumerPostUrl(Output.of(assertionConsumerPostUrl));
        }

        public Builder assertionConsumerRedirectUrl(@Nullable Output<String> assertionConsumerRedirectUrl) {
            $.assertionConsumerRedirectUrl = assertionConsumerRedirectUrl;
            return this;
        }

        public Builder assertionConsumerRedirectUrl(String assertionConsumerRedirectUrl) {
            return assertionConsumerRedirectUrl(Output.of(assertionConsumerRedirectUrl));
        }

        public Builder authenticationFlowBindingOverrides(@Nullable Output<ClientAuthenticationFlowBindingOverridesArgs> authenticationFlowBindingOverrides) {
            $.authenticationFlowBindingOverrides = authenticationFlowBindingOverrides;
            return this;
        }

        public Builder authenticationFlowBindingOverrides(ClientAuthenticationFlowBindingOverridesArgs authenticationFlowBindingOverrides) {
            return authenticationFlowBindingOverrides(Output.of(authenticationFlowBindingOverrides));
        }

        public Builder baseUrl(@Nullable Output<String> baseUrl) {
            $.baseUrl = baseUrl;
            return this;
        }

        public Builder baseUrl(String baseUrl) {
            return baseUrl(Output.of(baseUrl));
        }

        public Builder canonicalizationMethod(@Nullable Output<String> canonicalizationMethod) {
            $.canonicalizationMethod = canonicalizationMethod;
            return this;
        }

        public Builder canonicalizationMethod(String canonicalizationMethod) {
            return canonicalizationMethod(Output.of(canonicalizationMethod));
        }

        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        public Builder clientSignatureRequired(@Nullable Output<Boolean> clientSignatureRequired) {
            $.clientSignatureRequired = clientSignatureRequired;
            return this;
        }

        public Builder clientSignatureRequired(Boolean clientSignatureRequired) {
            return clientSignatureRequired(Output.of(clientSignatureRequired));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder encryptAssertions(@Nullable Output<Boolean> encryptAssertions) {
            $.encryptAssertions = encryptAssertions;
            return this;
        }

        public Builder encryptAssertions(Boolean encryptAssertions) {
            return encryptAssertions(Output.of(encryptAssertions));
        }

        public Builder encryptionCertificate(@Nullable Output<String> encryptionCertificate) {
            $.encryptionCertificate = encryptionCertificate;
            return this;
        }

        public Builder encryptionCertificate(String encryptionCertificate) {
            return encryptionCertificate(Output.of(encryptionCertificate));
        }

        public Builder encryptionCertificateSha1(@Nullable Output<String> encryptionCertificateSha1) {
            $.encryptionCertificateSha1 = encryptionCertificateSha1;
            return this;
        }

        public Builder encryptionCertificateSha1(String encryptionCertificateSha1) {
            return encryptionCertificateSha1(Output.of(encryptionCertificateSha1));
        }

        public Builder extraConfig(@Nullable Output<Map<String,String>> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        public Builder extraConfig(Map<String,String> extraConfig) {
            return extraConfig(Output.of(extraConfig));
        }

        public Builder forceNameIdFormat(@Nullable Output<Boolean> forceNameIdFormat) {
            $.forceNameIdFormat = forceNameIdFormat;
            return this;
        }

        public Builder forceNameIdFormat(Boolean forceNameIdFormat) {
            return forceNameIdFormat(Output.of(forceNameIdFormat));
        }

        public Builder forcePostBinding(@Nullable Output<Boolean> forcePostBinding) {
            $.forcePostBinding = forcePostBinding;
            return this;
        }

        public Builder forcePostBinding(Boolean forcePostBinding) {
            return forcePostBinding(Output.of(forcePostBinding));
        }

        public Builder frontChannelLogout(@Nullable Output<Boolean> frontChannelLogout) {
            $.frontChannelLogout = frontChannelLogout;
            return this;
        }

        public Builder frontChannelLogout(Boolean frontChannelLogout) {
            return frontChannelLogout(Output.of(frontChannelLogout));
        }

        public Builder fullScopeAllowed(@Nullable Output<Boolean> fullScopeAllowed) {
            $.fullScopeAllowed = fullScopeAllowed;
            return this;
        }

        public Builder fullScopeAllowed(Boolean fullScopeAllowed) {
            return fullScopeAllowed(Output.of(fullScopeAllowed));
        }

        public Builder idpInitiatedSsoRelayState(@Nullable Output<String> idpInitiatedSsoRelayState) {
            $.idpInitiatedSsoRelayState = idpInitiatedSsoRelayState;
            return this;
        }

        public Builder idpInitiatedSsoRelayState(String idpInitiatedSsoRelayState) {
            return idpInitiatedSsoRelayState(Output.of(idpInitiatedSsoRelayState));
        }

        public Builder idpInitiatedSsoUrlName(@Nullable Output<String> idpInitiatedSsoUrlName) {
            $.idpInitiatedSsoUrlName = idpInitiatedSsoUrlName;
            return this;
        }

        public Builder idpInitiatedSsoUrlName(String idpInitiatedSsoUrlName) {
            return idpInitiatedSsoUrlName(Output.of(idpInitiatedSsoUrlName));
        }

        public Builder includeAuthnStatement(@Nullable Output<Boolean> includeAuthnStatement) {
            $.includeAuthnStatement = includeAuthnStatement;
            return this;
        }

        public Builder includeAuthnStatement(Boolean includeAuthnStatement) {
            return includeAuthnStatement(Output.of(includeAuthnStatement));
        }

        public Builder loginTheme(@Nullable Output<String> loginTheme) {
            $.loginTheme = loginTheme;
            return this;
        }

        public Builder loginTheme(String loginTheme) {
            return loginTheme(Output.of(loginTheme));
        }

        public Builder logoutServicePostBindingUrl(@Nullable Output<String> logoutServicePostBindingUrl) {
            $.logoutServicePostBindingUrl = logoutServicePostBindingUrl;
            return this;
        }

        public Builder logoutServicePostBindingUrl(String logoutServicePostBindingUrl) {
            return logoutServicePostBindingUrl(Output.of(logoutServicePostBindingUrl));
        }

        public Builder logoutServiceRedirectBindingUrl(@Nullable Output<String> logoutServiceRedirectBindingUrl) {
            $.logoutServiceRedirectBindingUrl = logoutServiceRedirectBindingUrl;
            return this;
        }

        public Builder logoutServiceRedirectBindingUrl(String logoutServiceRedirectBindingUrl) {
            return logoutServiceRedirectBindingUrl(Output.of(logoutServiceRedirectBindingUrl));
        }

        public Builder masterSamlProcessingUrl(@Nullable Output<String> masterSamlProcessingUrl) {
            $.masterSamlProcessingUrl = masterSamlProcessingUrl;
            return this;
        }

        public Builder masterSamlProcessingUrl(String masterSamlProcessingUrl) {
            return masterSamlProcessingUrl(Output.of(masterSamlProcessingUrl));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder nameIdFormat(@Nullable Output<String> nameIdFormat) {
            $.nameIdFormat = nameIdFormat;
            return this;
        }

        public Builder nameIdFormat(String nameIdFormat) {
            return nameIdFormat(Output.of(nameIdFormat));
        }

        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public Builder rootUrl(@Nullable Output<String> rootUrl) {
            $.rootUrl = rootUrl;
            return this;
        }

        public Builder rootUrl(String rootUrl) {
            return rootUrl(Output.of(rootUrl));
        }

        public Builder signAssertions(@Nullable Output<Boolean> signAssertions) {
            $.signAssertions = signAssertions;
            return this;
        }

        public Builder signAssertions(Boolean signAssertions) {
            return signAssertions(Output.of(signAssertions));
        }

        public Builder signDocuments(@Nullable Output<Boolean> signDocuments) {
            $.signDocuments = signDocuments;
            return this;
        }

        public Builder signDocuments(Boolean signDocuments) {
            return signDocuments(Output.of(signDocuments));
        }

        public Builder signatureAlgorithm(@Nullable Output<String> signatureAlgorithm) {
            $.signatureAlgorithm = signatureAlgorithm;
            return this;
        }

        public Builder signatureAlgorithm(String signatureAlgorithm) {
            return signatureAlgorithm(Output.of(signatureAlgorithm));
        }

        public Builder signatureKeyName(@Nullable Output<String> signatureKeyName) {
            $.signatureKeyName = signatureKeyName;
            return this;
        }

        public Builder signatureKeyName(String signatureKeyName) {
            return signatureKeyName(Output.of(signatureKeyName));
        }

        public Builder signingCertificate(@Nullable Output<String> signingCertificate) {
            $.signingCertificate = signingCertificate;
            return this;
        }

        public Builder signingCertificate(String signingCertificate) {
            return signingCertificate(Output.of(signingCertificate));
        }

        public Builder signingCertificateSha1(@Nullable Output<String> signingCertificateSha1) {
            $.signingCertificateSha1 = signingCertificateSha1;
            return this;
        }

        public Builder signingCertificateSha1(String signingCertificateSha1) {
            return signingCertificateSha1(Output.of(signingCertificateSha1));
        }

        public Builder signingPrivateKey(@Nullable Output<String> signingPrivateKey) {
            $.signingPrivateKey = signingPrivateKey;
            return this;
        }

        public Builder signingPrivateKey(String signingPrivateKey) {
            return signingPrivateKey(Output.of(signingPrivateKey));
        }

        public Builder signingPrivateKeySha1(@Nullable Output<String> signingPrivateKeySha1) {
            $.signingPrivateKeySha1 = signingPrivateKeySha1;
            return this;
        }

        public Builder signingPrivateKeySha1(String signingPrivateKeySha1) {
            return signingPrivateKeySha1(Output.of(signingPrivateKeySha1));
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

        public ClientState build() {
            return $;
        }
    }

}
