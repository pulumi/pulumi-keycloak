// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClientAuthorization {
    private @Nullable Boolean allowRemoteResourceManagement;
    private @Nullable String decisionStrategy;
    private @Nullable Boolean keepDefaults;
    private String policyEnforcementMode;

    private ClientAuthorization() {}
    public Optional<Boolean> allowRemoteResourceManagement() {
        return Optional.ofNullable(this.allowRemoteResourceManagement);
    }
    public Optional<String> decisionStrategy() {
        return Optional.ofNullable(this.decisionStrategy);
    }
    public Optional<Boolean> keepDefaults() {
        return Optional.ofNullable(this.keepDefaults);
    }
    public String policyEnforcementMode() {
        return this.policyEnforcementMode;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClientAuthorization defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean allowRemoteResourceManagement;
        private @Nullable String decisionStrategy;
        private @Nullable Boolean keepDefaults;
        private String policyEnforcementMode;
        public Builder() {}
        public Builder(ClientAuthorization defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowRemoteResourceManagement = defaults.allowRemoteResourceManagement;
    	      this.decisionStrategy = defaults.decisionStrategy;
    	      this.keepDefaults = defaults.keepDefaults;
    	      this.policyEnforcementMode = defaults.policyEnforcementMode;
        }

        @CustomType.Setter
        public Builder allowRemoteResourceManagement(@Nullable Boolean allowRemoteResourceManagement) {

            this.allowRemoteResourceManagement = allowRemoteResourceManagement;
            return this;
        }
        @CustomType.Setter
        public Builder decisionStrategy(@Nullable String decisionStrategy) {

            this.decisionStrategy = decisionStrategy;
            return this;
        }
        @CustomType.Setter
        public Builder keepDefaults(@Nullable Boolean keepDefaults) {

            this.keepDefaults = keepDefaults;
            return this;
        }
        @CustomType.Setter
        public Builder policyEnforcementMode(String policyEnforcementMode) {
            if (policyEnforcementMode == null) {
              throw new MissingRequiredPropertyException("ClientAuthorization", "policyEnforcementMode");
            }
            this.policyEnforcementMode = policyEnforcementMode;
            return this;
        }
        public ClientAuthorization build() {
            final var _resultValue = new ClientAuthorization();
            _resultValue.allowRemoteResourceManagement = allowRemoteResourceManagement;
            _resultValue.decisionStrategy = decisionStrategy;
            _resultValue.keepDefaults = keepDefaults;
            _resultValue.policyEnforcementMode = policyEnforcementMode;
            return _resultValue;
        }
    }
}
