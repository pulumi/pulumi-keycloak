// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClientAuthorization {
    /**
     * @return When `true`, resources can be managed remotely by the resource server. Defaults to `false`.
     * 
     */
    private @Nullable Boolean allowRemoteResourceManagement;
    /**
     * @return Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
     * 
     */
    private @Nullable String decisionStrategy;
    /**
     * @return When `true`, defaults set by Keycloak will be respected. Defaults to `false`.
     * 
     */
    private @Nullable Boolean keepDefaults;
    /**
     * @return Dictates how policies are enforced when evaluating authorization requests. Can be one of `ENFORCING`, `PERMISSIVE`, or `DISABLED`.
     * 
     */
    private String policyEnforcementMode;

    private ClientAuthorization() {}
    /**
     * @return When `true`, resources can be managed remotely by the resource server. Defaults to `false`.
     * 
     */
    public Optional<Boolean> allowRemoteResourceManagement() {
        return Optional.ofNullable(this.allowRemoteResourceManagement);
    }
    /**
     * @return Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
     * 
     */
    public Optional<String> decisionStrategy() {
        return Optional.ofNullable(this.decisionStrategy);
    }
    /**
     * @return When `true`, defaults set by Keycloak will be respected. Defaults to `false`.
     * 
     */
    public Optional<Boolean> keepDefaults() {
        return Optional.ofNullable(this.keepDefaults);
    }
    /**
     * @return Dictates how policies are enforced when evaluating authorization requests. Can be one of `ENFORCING`, `PERMISSIVE`, or `DISABLED`.
     * 
     */
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
            this.policyEnforcementMode = Objects.requireNonNull(policyEnforcementMode);
            return this;
        }
        public ClientAuthorization build() {
            final var o = new ClientAuthorization();
            o.allowRemoteResourceManagement = allowRemoteResourceManagement;
            o.decisionStrategy = decisionStrategy;
            o.keepDefaults = keepDefaults;
            o.policyEnforcementMode = policyEnforcementMode;
            return o;
        }
    }
}
