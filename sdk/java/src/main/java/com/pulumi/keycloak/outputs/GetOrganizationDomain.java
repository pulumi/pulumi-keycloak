// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOrganizationDomain {
    /**
     * @return The organization name.
     * 
     */
    private String name;
    private Boolean verified;

    private GetOrganizationDomain() {}
    /**
     * @return The organization name.
     * 
     */
    public String name() {
        return this.name;
    }
    public Boolean verified() {
        return this.verified;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrganizationDomain defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private Boolean verified;
        public Builder() {}
        public Builder(GetOrganizationDomain defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.verified = defaults.verified;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetOrganizationDomain", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder verified(Boolean verified) {
            if (verified == null) {
              throw new MissingRequiredPropertyException("GetOrganizationDomain", "verified");
            }
            this.verified = verified;
            return this;
        }
        public GetOrganizationDomain build() {
            final var _resultValue = new GetOrganizationDomain();
            _resultValue.name = name;
            _resultValue.verified = verified;
            return _resultValue;
        }
    }
}
