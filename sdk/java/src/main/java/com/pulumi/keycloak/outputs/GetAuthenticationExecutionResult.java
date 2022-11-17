// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAuthenticationExecutionResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String parentFlowAlias;
    private String providerId;
    private String realmId;

    private GetAuthenticationExecutionResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String parentFlowAlias() {
        return this.parentFlowAlias;
    }
    public String providerId() {
        return this.providerId;
    }
    public String realmId() {
        return this.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAuthenticationExecutionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String parentFlowAlias;
        private String providerId;
        private String realmId;
        public Builder() {}
        public Builder(GetAuthenticationExecutionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.parentFlowAlias = defaults.parentFlowAlias;
    	      this.providerId = defaults.providerId;
    	      this.realmId = defaults.realmId;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder parentFlowAlias(String parentFlowAlias) {
            this.parentFlowAlias = Objects.requireNonNull(parentFlowAlias);
            return this;
        }
        @CustomType.Setter
        public Builder providerId(String providerId) {
            this.providerId = Objects.requireNonNull(providerId);
            return this;
        }
        @CustomType.Setter
        public Builder realmId(String realmId) {
            this.realmId = Objects.requireNonNull(realmId);
            return this;
        }
        public GetAuthenticationExecutionResult build() {
            final var o = new GetAuthenticationExecutionResult();
            o.id = id;
            o.parentFlowAlias = parentFlowAlias;
            o.providerId = providerId;
            o.realmId = realmId;
            return o;
        }
    }
}