// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RealmWebAuthnPolicy {
    /**
     * @return A set of AAGUIDs for which an authenticator can be registered.
     * 
     */
    private @Nullable List<String> acceptableAaguids;
    /**
     * @return Either none, indirect or direct
     * 
     */
    private @Nullable String attestationConveyancePreference;
    /**
     * @return Either platform or cross-platform
     * 
     */
    private @Nullable String authenticatorAttachment;
    /**
     * @return When `true`, Keycloak will avoid registering the authenticator for WebAuthn if it has already been registered. Defaults to `false`.
     * 
     */
    private @Nullable Boolean avoidSameAuthenticatorRegister;
    /**
     * @return The timeout value for creating a user&#39;s public key credential in seconds. When set to `0`, this timeout option is not adapted. Defaults to `0`.
     * 
     */
    private @Nullable Integer createTimeout;
    /**
     * @return A set of extra origins for non-web applications.
     * 
     */
    private @Nullable List<String> extraOrigins;
    /**
     * @return A human-readable server name for the WebAuthn Relying Party. Defaults to `keycloak`.
     * 
     */
    private @Nullable String relyingPartyEntityName;
    /**
     * @return The WebAuthn relying party ID.
     * 
     */
    private @Nullable String relyingPartyId;
    /**
     * @return Either Yes or No
     * 
     */
    private @Nullable String requireResidentKey;
    /**
     * @return Keycloak lists ES256, ES384, ES512, RS256, RS384, RS512, RS1 at the time of writing
     * 
     */
    private @Nullable List<String> signatureAlgorithms;
    /**
     * @return Either required, preferred or discouraged
     * 
     */
    private @Nullable String userVerificationRequirement;

    private RealmWebAuthnPolicy() {}
    /**
     * @return A set of AAGUIDs for which an authenticator can be registered.
     * 
     */
    public List<String> acceptableAaguids() {
        return this.acceptableAaguids == null ? List.of() : this.acceptableAaguids;
    }
    /**
     * @return Either none, indirect or direct
     * 
     */
    public Optional<String> attestationConveyancePreference() {
        return Optional.ofNullable(this.attestationConveyancePreference);
    }
    /**
     * @return Either platform or cross-platform
     * 
     */
    public Optional<String> authenticatorAttachment() {
        return Optional.ofNullable(this.authenticatorAttachment);
    }
    /**
     * @return When `true`, Keycloak will avoid registering the authenticator for WebAuthn if it has already been registered. Defaults to `false`.
     * 
     */
    public Optional<Boolean> avoidSameAuthenticatorRegister() {
        return Optional.ofNullable(this.avoidSameAuthenticatorRegister);
    }
    /**
     * @return The timeout value for creating a user&#39;s public key credential in seconds. When set to `0`, this timeout option is not adapted. Defaults to `0`.
     * 
     */
    public Optional<Integer> createTimeout() {
        return Optional.ofNullable(this.createTimeout);
    }
    /**
     * @return A set of extra origins for non-web applications.
     * 
     */
    public List<String> extraOrigins() {
        return this.extraOrigins == null ? List.of() : this.extraOrigins;
    }
    /**
     * @return A human-readable server name for the WebAuthn Relying Party. Defaults to `keycloak`.
     * 
     */
    public Optional<String> relyingPartyEntityName() {
        return Optional.ofNullable(this.relyingPartyEntityName);
    }
    /**
     * @return The WebAuthn relying party ID.
     * 
     */
    public Optional<String> relyingPartyId() {
        return Optional.ofNullable(this.relyingPartyId);
    }
    /**
     * @return Either Yes or No
     * 
     */
    public Optional<String> requireResidentKey() {
        return Optional.ofNullable(this.requireResidentKey);
    }
    /**
     * @return Keycloak lists ES256, ES384, ES512, RS256, RS384, RS512, RS1 at the time of writing
     * 
     */
    public List<String> signatureAlgorithms() {
        return this.signatureAlgorithms == null ? List.of() : this.signatureAlgorithms;
    }
    /**
     * @return Either required, preferred or discouraged
     * 
     */
    public Optional<String> userVerificationRequirement() {
        return Optional.ofNullable(this.userVerificationRequirement);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RealmWebAuthnPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> acceptableAaguids;
        private @Nullable String attestationConveyancePreference;
        private @Nullable String authenticatorAttachment;
        private @Nullable Boolean avoidSameAuthenticatorRegister;
        private @Nullable Integer createTimeout;
        private @Nullable List<String> extraOrigins;
        private @Nullable String relyingPartyEntityName;
        private @Nullable String relyingPartyId;
        private @Nullable String requireResidentKey;
        private @Nullable List<String> signatureAlgorithms;
        private @Nullable String userVerificationRequirement;
        public Builder() {}
        public Builder(RealmWebAuthnPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acceptableAaguids = defaults.acceptableAaguids;
    	      this.attestationConveyancePreference = defaults.attestationConveyancePreference;
    	      this.authenticatorAttachment = defaults.authenticatorAttachment;
    	      this.avoidSameAuthenticatorRegister = defaults.avoidSameAuthenticatorRegister;
    	      this.createTimeout = defaults.createTimeout;
    	      this.extraOrigins = defaults.extraOrigins;
    	      this.relyingPartyEntityName = defaults.relyingPartyEntityName;
    	      this.relyingPartyId = defaults.relyingPartyId;
    	      this.requireResidentKey = defaults.requireResidentKey;
    	      this.signatureAlgorithms = defaults.signatureAlgorithms;
    	      this.userVerificationRequirement = defaults.userVerificationRequirement;
        }

        @CustomType.Setter
        public Builder acceptableAaguids(@Nullable List<String> acceptableAaguids) {

            this.acceptableAaguids = acceptableAaguids;
            return this;
        }
        public Builder acceptableAaguids(String... acceptableAaguids) {
            return acceptableAaguids(List.of(acceptableAaguids));
        }
        @CustomType.Setter
        public Builder attestationConveyancePreference(@Nullable String attestationConveyancePreference) {

            this.attestationConveyancePreference = attestationConveyancePreference;
            return this;
        }
        @CustomType.Setter
        public Builder authenticatorAttachment(@Nullable String authenticatorAttachment) {

            this.authenticatorAttachment = authenticatorAttachment;
            return this;
        }
        @CustomType.Setter
        public Builder avoidSameAuthenticatorRegister(@Nullable Boolean avoidSameAuthenticatorRegister) {

            this.avoidSameAuthenticatorRegister = avoidSameAuthenticatorRegister;
            return this;
        }
        @CustomType.Setter
        public Builder createTimeout(@Nullable Integer createTimeout) {

            this.createTimeout = createTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder extraOrigins(@Nullable List<String> extraOrigins) {

            this.extraOrigins = extraOrigins;
            return this;
        }
        public Builder extraOrigins(String... extraOrigins) {
            return extraOrigins(List.of(extraOrigins));
        }
        @CustomType.Setter
        public Builder relyingPartyEntityName(@Nullable String relyingPartyEntityName) {

            this.relyingPartyEntityName = relyingPartyEntityName;
            return this;
        }
        @CustomType.Setter
        public Builder relyingPartyId(@Nullable String relyingPartyId) {

            this.relyingPartyId = relyingPartyId;
            return this;
        }
        @CustomType.Setter
        public Builder requireResidentKey(@Nullable String requireResidentKey) {

            this.requireResidentKey = requireResidentKey;
            return this;
        }
        @CustomType.Setter
        public Builder signatureAlgorithms(@Nullable List<String> signatureAlgorithms) {

            this.signatureAlgorithms = signatureAlgorithms;
            return this;
        }
        public Builder signatureAlgorithms(String... signatureAlgorithms) {
            return signatureAlgorithms(List.of(signatureAlgorithms));
        }
        @CustomType.Setter
        public Builder userVerificationRequirement(@Nullable String userVerificationRequirement) {

            this.userVerificationRequirement = userVerificationRequirement;
            return this;
        }
        public RealmWebAuthnPolicy build() {
            final var _resultValue = new RealmWebAuthnPolicy();
            _resultValue.acceptableAaguids = acceptableAaguids;
            _resultValue.attestationConveyancePreference = attestationConveyancePreference;
            _resultValue.authenticatorAttachment = authenticatorAttachment;
            _resultValue.avoidSameAuthenticatorRegister = avoidSameAuthenticatorRegister;
            _resultValue.createTimeout = createTimeout;
            _resultValue.extraOrigins = extraOrigins;
            _resultValue.relyingPartyEntityName = relyingPartyEntityName;
            _resultValue.relyingPartyId = relyingPartyId;
            _resultValue.requireResidentKey = requireResidentKey;
            _resultValue.signatureAlgorithms = signatureAlgorithms;
            _resultValue.userVerificationRequirement = userVerificationRequirement;
            return _resultValue;
        }
    }
}
