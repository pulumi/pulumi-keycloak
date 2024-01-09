// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRealmWebAuthnPolicy {
    private List<String> acceptableAaguids;
    private String attestationConveyancePreference;
    private String authenticatorAttachment;
    private Boolean avoidSameAuthenticatorRegister;
    private Integer createTimeout;
    private String relyingPartyEntityName;
    private String relyingPartyId;
    private String requireResidentKey;
    private List<String> signatureAlgorithms;
    private String userVerificationRequirement;

    private GetRealmWebAuthnPolicy() {}
    public List<String> acceptableAaguids() {
        return this.acceptableAaguids;
    }
    public String attestationConveyancePreference() {
        return this.attestationConveyancePreference;
    }
    public String authenticatorAttachment() {
        return this.authenticatorAttachment;
    }
    public Boolean avoidSameAuthenticatorRegister() {
        return this.avoidSameAuthenticatorRegister;
    }
    public Integer createTimeout() {
        return this.createTimeout;
    }
    public String relyingPartyEntityName() {
        return this.relyingPartyEntityName;
    }
    public String relyingPartyId() {
        return this.relyingPartyId;
    }
    public String requireResidentKey() {
        return this.requireResidentKey;
    }
    public List<String> signatureAlgorithms() {
        return this.signatureAlgorithms;
    }
    public String userVerificationRequirement() {
        return this.userVerificationRequirement;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRealmWebAuthnPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> acceptableAaguids;
        private String attestationConveyancePreference;
        private String authenticatorAttachment;
        private Boolean avoidSameAuthenticatorRegister;
        private Integer createTimeout;
        private String relyingPartyEntityName;
        private String relyingPartyId;
        private String requireResidentKey;
        private List<String> signatureAlgorithms;
        private String userVerificationRequirement;
        public Builder() {}
        public Builder(GetRealmWebAuthnPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acceptableAaguids = defaults.acceptableAaguids;
    	      this.attestationConveyancePreference = defaults.attestationConveyancePreference;
    	      this.authenticatorAttachment = defaults.authenticatorAttachment;
    	      this.avoidSameAuthenticatorRegister = defaults.avoidSameAuthenticatorRegister;
    	      this.createTimeout = defaults.createTimeout;
    	      this.relyingPartyEntityName = defaults.relyingPartyEntityName;
    	      this.relyingPartyId = defaults.relyingPartyId;
    	      this.requireResidentKey = defaults.requireResidentKey;
    	      this.signatureAlgorithms = defaults.signatureAlgorithms;
    	      this.userVerificationRequirement = defaults.userVerificationRequirement;
        }

        @CustomType.Setter
        public Builder acceptableAaguids(List<String> acceptableAaguids) {
            if (acceptableAaguids == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "acceptableAaguids");
            }
            this.acceptableAaguids = acceptableAaguids;
            return this;
        }
        public Builder acceptableAaguids(String... acceptableAaguids) {
            return acceptableAaguids(List.of(acceptableAaguids));
        }
        @CustomType.Setter
        public Builder attestationConveyancePreference(String attestationConveyancePreference) {
            if (attestationConveyancePreference == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "attestationConveyancePreference");
            }
            this.attestationConveyancePreference = attestationConveyancePreference;
            return this;
        }
        @CustomType.Setter
        public Builder authenticatorAttachment(String authenticatorAttachment) {
            if (authenticatorAttachment == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "authenticatorAttachment");
            }
            this.authenticatorAttachment = authenticatorAttachment;
            return this;
        }
        @CustomType.Setter
        public Builder avoidSameAuthenticatorRegister(Boolean avoidSameAuthenticatorRegister) {
            if (avoidSameAuthenticatorRegister == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "avoidSameAuthenticatorRegister");
            }
            this.avoidSameAuthenticatorRegister = avoidSameAuthenticatorRegister;
            return this;
        }
        @CustomType.Setter
        public Builder createTimeout(Integer createTimeout) {
            if (createTimeout == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "createTimeout");
            }
            this.createTimeout = createTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder relyingPartyEntityName(String relyingPartyEntityName) {
            if (relyingPartyEntityName == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "relyingPartyEntityName");
            }
            this.relyingPartyEntityName = relyingPartyEntityName;
            return this;
        }
        @CustomType.Setter
        public Builder relyingPartyId(String relyingPartyId) {
            if (relyingPartyId == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "relyingPartyId");
            }
            this.relyingPartyId = relyingPartyId;
            return this;
        }
        @CustomType.Setter
        public Builder requireResidentKey(String requireResidentKey) {
            if (requireResidentKey == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "requireResidentKey");
            }
            this.requireResidentKey = requireResidentKey;
            return this;
        }
        @CustomType.Setter
        public Builder signatureAlgorithms(List<String> signatureAlgorithms) {
            if (signatureAlgorithms == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "signatureAlgorithms");
            }
            this.signatureAlgorithms = signatureAlgorithms;
            return this;
        }
        public Builder signatureAlgorithms(String... signatureAlgorithms) {
            return signatureAlgorithms(List.of(signatureAlgorithms));
        }
        @CustomType.Setter
        public Builder userVerificationRequirement(String userVerificationRequirement) {
            if (userVerificationRequirement == null) {
              throw new MissingRequiredPropertyException("GetRealmWebAuthnPolicy", "userVerificationRequirement");
            }
            this.userVerificationRequirement = userVerificationRequirement;
            return this;
        }
        public GetRealmWebAuthnPolicy build() {
            final var _resultValue = new GetRealmWebAuthnPolicy();
            _resultValue.acceptableAaguids = acceptableAaguids;
            _resultValue.attestationConveyancePreference = attestationConveyancePreference;
            _resultValue.authenticatorAttachment = authenticatorAttachment;
            _resultValue.avoidSameAuthenticatorRegister = avoidSameAuthenticatorRegister;
            _resultValue.createTimeout = createTimeout;
            _resultValue.relyingPartyEntityName = relyingPartyEntityName;
            _resultValue.relyingPartyId = relyingPartyId;
            _resultValue.requireResidentKey = requireResidentKey;
            _resultValue.signatureAlgorithms = signatureAlgorithms;
            _resultValue.userVerificationRequirement = userVerificationRequirement;
            return _resultValue;
        }
    }
}
