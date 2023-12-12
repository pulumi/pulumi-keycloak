// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAuthenticationFlowResult {
    private String alias;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String realmId;

    private GetAuthenticationFlowResult() {}
    public String alias() {
        return this.alias;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String realmId() {
        return this.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAuthenticationFlowResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String alias;
        private String id;
        private String realmId;
        public Builder() {}
        public Builder(GetAuthenticationFlowResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alias = defaults.alias;
    	      this.id = defaults.id;
    	      this.realmId = defaults.realmId;
        }

        @CustomType.Setter
        public Builder alias(String alias) {
            this.alias = Objects.requireNonNull(alias);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder realmId(String realmId) {
            this.realmId = Objects.requireNonNull(realmId);
            return this;
        }
        public GetAuthenticationFlowResult build() {
            final var _resultValue = new GetAuthenticationFlowResult();
            _resultValue.alias = alias;
            _resultValue.id = id;
            _resultValue.realmId = realmId;
            return _resultValue;
        }
    }
}
