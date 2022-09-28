// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClientInstallationProviderResult {
    private String clientId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String providerId;
    private String realmId;
    /**
     * @return (Computed) The returned document needed for SAML installation.
     * 
     */
    private String value;

    private GetClientInstallationProviderResult() {}
    public String clientId() {
        return this.clientId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String providerId() {
        return this.providerId;
    }
    public String realmId() {
        return this.realmId;
    }
    /**
     * @return (Computed) The returned document needed for SAML installation.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClientInstallationProviderResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clientId;
        private String id;
        private String providerId;
        private String realmId;
        private String value;
        public Builder() {}
        public Builder(GetClientInstallationProviderResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientId = defaults.clientId;
    	      this.id = defaults.id;
    	      this.providerId = defaults.providerId;
    	      this.realmId = defaults.realmId;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder clientId(String clientId) {
            this.clientId = Objects.requireNonNull(clientId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
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
        @CustomType.Setter
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public GetClientInstallationProviderResult build() {
            final var o = new GetClientInstallationProviderResult();
            o.clientId = clientId;
            o.id = id;
            o.providerId = providerId;
            o.realmId = realmId;
            o.value = value;
            return o;
        }
    }
}
