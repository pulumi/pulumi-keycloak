// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClientAuthenticationFlowBindingOverrides {
    /**
     * @return Browser flow id, (flow needs to exist)
     * 
     */
    private @Nullable String browserId;
    /**
     * @return Direct grant flow id (flow needs to exist)
     * 
     */
    private @Nullable String directGrantId;

    private ClientAuthenticationFlowBindingOverrides() {}
    /**
     * @return Browser flow id, (flow needs to exist)
     * 
     */
    public Optional<String> browserId() {
        return Optional.ofNullable(this.browserId);
    }
    /**
     * @return Direct grant flow id (flow needs to exist)
     * 
     */
    public Optional<String> directGrantId() {
        return Optional.ofNullable(this.directGrantId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClientAuthenticationFlowBindingOverrides defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String browserId;
        private @Nullable String directGrantId;
        public Builder() {}
        public Builder(ClientAuthenticationFlowBindingOverrides defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.browserId = defaults.browserId;
    	      this.directGrantId = defaults.directGrantId;
        }

        @CustomType.Setter
        public Builder browserId(@Nullable String browserId) {
            this.browserId = browserId;
            return this;
        }
        @CustomType.Setter
        public Builder directGrantId(@Nullable String directGrantId) {
            this.directGrantId = directGrantId;
            return this;
        }
        public ClientAuthenticationFlowBindingOverrides build() {
            final var _resultValue = new ClientAuthenticationFlowBindingOverrides();
            _resultValue.browserId = browserId;
            _resultValue.directGrantId = directGrantId;
            return _resultValue;
        }
    }
}
