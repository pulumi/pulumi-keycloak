// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClientAuthorization {
    private final Boolean allowRemoteResourceManagement;
    private final String decisionStrategy;
    private final Boolean keepDefaults;
    private final String policyEnforcementMode;

    @CustomType.Constructor
    private GetClientAuthorization(
        @CustomType.Parameter("allowRemoteResourceManagement") Boolean allowRemoteResourceManagement,
        @CustomType.Parameter("decisionStrategy") String decisionStrategy,
        @CustomType.Parameter("keepDefaults") Boolean keepDefaults,
        @CustomType.Parameter("policyEnforcementMode") String policyEnforcementMode) {
        this.allowRemoteResourceManagement = allowRemoteResourceManagement;
        this.decisionStrategy = decisionStrategy;
        this.keepDefaults = keepDefaults;
        this.policyEnforcementMode = policyEnforcementMode;
    }

    public Boolean allowRemoteResourceManagement() {
        return this.allowRemoteResourceManagement;
    }
    public String decisionStrategy() {
        return this.decisionStrategy;
    }
    public Boolean keepDefaults() {
        return this.keepDefaults;
    }
    public String policyEnforcementMode() {
        return this.policyEnforcementMode;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClientAuthorization defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Boolean allowRemoteResourceManagement;
        private String decisionStrategy;
        private Boolean keepDefaults;
        private String policyEnforcementMode;

        public Builder() {
    	      // Empty
        }

        public Builder(GetClientAuthorization defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowRemoteResourceManagement = defaults.allowRemoteResourceManagement;
    	      this.decisionStrategy = defaults.decisionStrategy;
    	      this.keepDefaults = defaults.keepDefaults;
    	      this.policyEnforcementMode = defaults.policyEnforcementMode;
        }

        public Builder allowRemoteResourceManagement(Boolean allowRemoteResourceManagement) {
            this.allowRemoteResourceManagement = Objects.requireNonNull(allowRemoteResourceManagement);
            return this;
        }
        public Builder decisionStrategy(String decisionStrategy) {
            this.decisionStrategy = Objects.requireNonNull(decisionStrategy);
            return this;
        }
        public Builder keepDefaults(Boolean keepDefaults) {
            this.keepDefaults = Objects.requireNonNull(keepDefaults);
            return this;
        }
        public Builder policyEnforcementMode(String policyEnforcementMode) {
            this.policyEnforcementMode = Objects.requireNonNull(policyEnforcementMode);
            return this;
        }        public GetClientAuthorization build() {
            return new GetClientAuthorization(allowRemoteResourceManagement, decisionStrategy, keepDefaults, policyEnforcementMode);
        }
    }
}
